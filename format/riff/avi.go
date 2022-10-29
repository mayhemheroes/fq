package riff

// TODO:
// refactor wav?
// assert
// refactor riff? common audio/video headers?
// DV handler https://learn.microsoft.com/en-us/windows/win32/directshow/dv-data-in-the-avi-file-format
// palette change
// rec groups
// 64 bit size?
// AVIX, multiple RIFF headers?

// https://learn.microsoft.com/en-us/windows/win32/directshow/avi-riff-file-reference
// http://www.jmcgowan.com/odmlff2.pdf
// https://github.com/FFmpeg/FFmpeg/blob/master/libavformat/avidec.c

import (
	"strconv"

	"github.com/wader/fq/format"
	"github.com/wader/fq/pkg/decode"
	"github.com/wader/fq/pkg/interp"
	"github.com/wader/fq/pkg/ranges"
	"github.com/wader/fq/pkg/scalar"
)

var aviMp3FrameFormat decode.Group
var aviMpegAVCAUFormat decode.Group
var aviMpegHEVCAUFormat decode.Group
var aviFLACFrameFormat decode.Group

func init() {
	interp.RegisterFormat(decode.Format{
		Name:        format.AVI,
		Description: "Audio Video Interleaved",
		DecodeFn:    aviDecode,
		DecodeInArg: format.AviIn{
			DecodeSamples: true,
		},
		Dependencies: []decode.Dependency{
			{Names: []string{format.AVC_AU}, Group: &aviMpegAVCAUFormat},
			{Names: []string{format.HEVC_AU}, Group: &aviMpegHEVCAUFormat},
			{Names: []string{format.MP3_FRAME}, Group: &aviMp3FrameFormat},
			{Names: []string{format.FLAC_FRAME}, Group: &aviFLACFrameFormat},
		},
		Groups: []string{format.PROBE},
	})
}

var aviListTypeDescriptions = scalar.StrToDescription{
	"hdrl": "AVI main list",
	"strl": "Stream list",
}

var aviStrhTypeDescriptions = scalar.StrToDescription{
	"auds": "Audio stream",
	"mids": "MIDI stream",
	"txts": "Text stream",
	"vids": "Video stream",
}

var aviIndexTypeNames = scalar.UToSymStr{
	0: "indexes",
	1: "chunks",
}

const (
	aviStreamChunkTypeUncompressedVideo = "db"
	aviStreamChunkTypeCompressedVideo   = "dc"
	aviStreamChunkTypePaletteChange     = "pc"
	aviStreamChunkTypeAudio             = "wb"
	aviStreamChunkTypeIndex             = "ix"
)

var aviStreamChunkTypeDescriptions = scalar.StrToDescription{
	aviStreamChunkTypeUncompressedVideo: "Uncompressed video frame",
	aviStreamChunkTypeCompressedVideo:   "Compressed video frame",
	aviStreamChunkTypePaletteChange:     "Palette change",
	aviStreamChunkTypeAudio:             "Audio data",
	aviStreamChunkTypeIndex:             "Index",
}

const aviRiffType = "AVI "

type aviStrl struct {
	typ     string
	handler string
	stream  *aviStream
}

type idx1Sample struct {
	offset     int64
	size       int64
	streamNr   int
	streamType string
}

type aviStream struct {
	hasFormat   bool
	format      decode.Group
	formatInArg any
	indexes     []ranges.Range
}

func aviParseChunkID(id string) (string, int, bool) {
	if len(id) != 4 {
		return "", 0, false
	}

	isDigits := func(s string) bool {
		for _, c := range s {
			if !(c >= '0' && c <= '9') {
				return false
			}
		}
		return true
	}

	var typ string
	var indexStr string
	switch {
	case isDigits(id[0:2]):
		// ##dc media etc
		indexStr, typ = id[0:2], id[2:4]
	case isDigits(id[2:4]):
		// ix## index etc
		typ, indexStr = id[0:2], id[2:4]
	default:
		return "", 0, false
	}

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		panic("unreachable")
	}

	return typ, index, true

}

func aviIsStreamType(typ string) bool {
	switch typ {
	case aviStreamChunkTypeUncompressedVideo,
		aviStreamChunkTypeCompressedVideo,
		aviStreamChunkTypeAudio:
		return true
	default:
		return false
	}
}

func aviDecorateStreamID(d *decode.D, id string) (string, int) {
	typ, index, ok := aviParseChunkID(id)
	if ok && aviIsStreamType(typ) {
		d.FieldValueStr("stream_type", typ, aviStreamChunkTypeDescriptions)
		d.FieldValueU("stream_nr", uint64(index))
		return typ, index
	}
	return "", 0
}

func aviDecode(d *decode.D, in any) any {
	ai, _ := in.(format.AviIn)

	d.Endian = decode.LittleEndian

	var streams []*aviStream
	var idx1Samples []idx1Sample
	var moviListPos int64 // point to first bit after type

	var riffType string
	riffDecode(
		d,
		nil,
		func(d *decode.D, path path) (string, int64) {
			id := d.FieldUTF8("id", 4, chunkIDDescriptions)
			aviDecorateStreamID(d, id)
			size := d.FieldU32("size")
			return id, int64(size)
		},
		func(d *decode.D, id string, path path) (bool, any) {
			switch id {
			case "RIFF":
				riffType = d.FieldUTF8("type", 4, d.AssertStr(aviRiffType))
				return true, nil

			case "LIST":
				typ := d.FieldUTF8("type", 4, aviListTypeDescriptions)
				switch typ {
				case "strl":
					return true, &aviStrl{}
				case "movi":
					moviListPos = d.Pos()
				}

				return true, nil

			case "idx1":
				// TODO: limit not on eof? or trunc align?
				// TODO: seems there are files with weird tailing extra index entries?
				d.FieldArray("indexes", func(d *decode.D) {
					for !d.End() {
						d.FieldStruct("index", func(d *decode.D) {
							id := d.FieldUTF8("id", 4)
							typ, index := aviDecorateStreamID(d, id)
							d.FieldU32("flags") // TODO: key frame flag etc
							offset := int64(d.FieldU32("offset"))
							length := int64(d.FieldU32("length"))

							idx1Samples = append(idx1Samples, idx1Sample{
								offset:     offset * 8,
								size:       length * 8,
								streamNr:   index,
								streamType: typ,
							})
						})
					}
				})
				return false, nil

			case "avih":
				d.FieldU32("micro_sec_per_frame")
				d.FieldU32("max_bytes_per_sec")
				d.FieldU32("padding_granularity")
				d.FieldU32("flags") // TODO
				d.FieldU32("total_frames")
				d.FieldU32("initial_frames")
				d.FieldU32("streams")
				d.FieldU32("suggested_buffer_size")
				d.FieldU32("width")
				d.FieldU32("height")
				d.FieldRawLen("reserved", 32*4)
				return false, nil

			case "dmlh":
				d.FieldU32("total_frames")
				return false, nil

			case "strh":
				typ := d.FieldUTF8("type", 4, aviStrhTypeDescriptions)
				handler := d.FieldUTF8("handler", 4)
				d.FieldU32("flags")
				d.FieldU16("priority")
				d.FieldU16("language")
				d.FieldU32("initial_frames")
				d.FieldU32("scale")
				d.FieldU32("rate")
				d.FieldU32("start")
				d.FieldU32("length")
				d.FieldU32("suggested_buffer_size")
				d.FieldU32("quality")
				d.FieldU32("sample_size")
				// TODO: ?
				d.FieldArray("frame", func(d *decode.D) {
					d.FieldU16("x")
					d.FieldU16("y")
					d.FieldU16("width")
					d.FieldU16("height")
				})

				if aviStrl, aviStrlOk := path.topData().(*aviStrl); aviStrlOk {
					aviStrl.typ = typ
					aviStrl.handler = handler
				}

				return false, nil

			case "strf":
				s := &aviStream{}

				typ := ""
				if aviStrl, aviStrlOk := path.topData().(*aviStrl); aviStrlOk {
					typ = aviStrl.typ
					aviStrl.stream = s
				}

				switch typ {
				case "vids":
					// BITMAPINFOHEADER
					size := d.BitsLeft()
					biSize := d.FieldU32("bi_size")
					d.FieldU32("width")
					d.FieldU32("height")
					d.FieldU16("planes")
					d.FieldU16("bit_count")
					compression := d.FieldUTF8("compression", 4)
					d.FieldU32("size_image")
					d.FieldU32("x_pels_per_meter")
					d.FieldU32("y_pels_per_meter")
					d.FieldU32("clr_used")
					d.FieldU32("clr_important")
					extraSize := size - int64(biSize)*8 - 2*32
					if extraSize > 0 {
						d.FieldRawLen("extra", extraSize)
					}

					// TODO: if dvsd handler and extraSize >= 32 then DVINFO?

					switch compression {
					case format.BMPTagH264,
						format.BMPTagH264_h264,
						format.BMPTagH264_X264,
						format.BMPTagH264_x264,
						format.BMPTagH264_avc1,
						format.BMPTagH264_DAVC,
						format.BMPTagH264_SMV2,
						format.BMPTagH264_VSSH,
						format.BMPTagH264_Q264,
						format.BMPTagH264_V264,
						format.BMPTagH264_GAVC,
						format.BMPTagH264_UMSV,
						format.BMPTagH264_tshd,
						format.BMPTagH264_INMC:
						s.format = aviMpegAVCAUFormat
						s.hasFormat = true
					case format.BMPTagHEVC,
						format.BMPTagHEVC_H265:
						s.format = aviMpegHEVCAUFormat
						s.hasFormat = true
					}

				case "auds":
					// WAVEFORMATEX
					formatTag := d.FieldU16("format_tag", format.WAVTagNames)
					d.FieldU16("channels")
					d.FieldU32("samples_per_sec")
					d.FieldU32("avg_bytes_per_sec")
					d.FieldU16("block_align")
					d.FieldU16("bits_per_sample")
					// seems to be optional
					if d.BitsLeft() >= 16 {
						cbSize := d.FieldU16("cb_size")
						if cbSize > 0 {
							d.FieldRawLen("extra", int64(cbSize)*8)
						}
					}

					switch formatTag {
					case format.WAVTagMP3:
						s.format = aviMp3FrameFormat
						s.hasFormat = true
					case format.WAVTagFLAC:
						// TODO: can flac in avi have streaminfo somehow?
						s.format = aviFLACFrameFormat
						s.hasFormat = true
					}
				case "iavs":
					// DVINFO
					d.FieldU32("dva_aux_src")
					d.FieldU32("dva_aux_ctl")
					d.FieldU32("dva_aux_src1")
					d.FieldU32("dva_aux_ctl1")
					d.FieldU32("dvv_aux_src")
					d.FieldU32("dvv_aux_ctl")
					d.FieldRawLen("dvv_reserved", 32*2)
				}

				streams = append(streams, s)

				return false, nil

			case "indx":
				var stream *aviStream
				if aviStrl, aviStrlOk := path.topData().(*aviStrl); aviStrlOk {
					stream = aviStrl.stream
				}

				d.FieldU16("longs_per_entry") // TODO: use?
				d.FieldU8("index_subtype")
				d.FieldU8("index_type")
				nEntriesInUse := d.FieldU32("entries_in_use")
				chunkID := d.FieldUTF8("chunk_id", 4)
				aviDecorateStreamID(d, chunkID)
				d.FieldU64("base")
				d.FieldU32("unused")
				d.FieldArray("index", func(d *decode.D) {
					for i := 0; i < int(nEntriesInUse); i++ {
						d.FieldStruct("index", func(d *decode.D) {
							offset := int64(d.FieldU64("offset"))
							size := int64(d.FieldU32("size"))
							d.FieldU32("duration")

							if stream != nil {
								stream.indexes = append(stream.indexes, ranges.Range{
									Start: offset * 8,
									Len:   size * 8,
								})
							}
						})
					}
				})

				// DWORD cb;
				// WORD wLongsPerEntry; // size of each entry in aIndex array
				// BYTE bIndexSubType; // future use. must be 0
				// BYTE bIndexType; // one of AVI_INDEX_* codes
				// DWORD nEntriesInUse; // index of first unused member in aIndex array
				// DWORD dwChunkId; // fcc of what is indexed
				// DWORD dwReserved[3]; // meaning differs for each index
				// // type/subtype. 0 if unused
				// struct _aviindex_entry {
				// DWORD adw[wLongsPerEntry];
				// } aIndex[ ];

				return false, nil

			case "vprp":
				d.FieldU32("video_format_token")
				d.FieldU32("video_standard")
				d.FieldU32("vertical_refresh_rate")
				d.FieldU32("h_total_in_t")
				d.FieldU32("v_total_in_lines")
				d.FieldU32("frame_aspect_ratio")
				d.FieldU32("frame_width_in_pixels")
				d.FieldU32("frame_height_in_lines")
				nbFieldPerFrame := d.FieldU32("nb_field_per_frame")
				d.FieldArray("field_info", func(d *decode.D) {
					for i := 0; i < int(nbFieldPerFrame); i++ {
						d.FieldStruct("field_info", func(d *decode.D) {
							d.FieldU32("compressed_bm_height")
							d.FieldU32("compressed_bm_width")
							d.FieldU32("valid_bm_height")
							d.FieldU32("valid_bm_width")
							d.FieldU32("valid_bmx_offset")
							d.FieldU32("valid_bmy_offset")
							d.FieldU32("video_x_offset_in_t")
							d.FieldU32("video_y_valid_start_line")
						})
					}
				})
				return false, nil

			default:
				if riffIsStringChunkID(id) {
					d.FieldUTF8NullFixedLen("value", int(d.BitsLeft())/8)
					return false, nil
				}

				// TODO: zero length samples?
				// TODO: palette?

				typ, index, _ := aviParseChunkID(id)
				switch {
				case typ == "ix":

					// TODO: index for stream index ##
					// TODO: share index code? if longs_per_entry == 2 skip duration?
					// TODO: share with below index decode

					d.FieldU16("longs_per_entry") // TODO: use?
					d.FieldU8("index_subtype")
					d.FieldU8("index_type", aviIndexTypeNames)
					nEntriesInUse := d.FieldU32("entries_in_use")
					chunkID := d.FieldUTF8("chunk_id", 4)
					aviDecorateStreamID(d, chunkID)
					d.FieldU64("base")
					d.FieldU32("unused")
					d.FieldArray("index", func(d *decode.D) {
						for i := 0; i < int(nEntriesInUse); i++ {
							d.FieldStruct("index", func(d *decode.D) {
								offset := int64(d.FieldU32("offset"))
								size := int64(d.FieldU32("size"))

								// d.FieldU32("duration"
								_ = offset
								_ = size

								// if stream != nil {
								// 	stream.indexes = append(stream.indexes, ranges.Range{
								// 		Start: offset * 8,
								// 		Len:   size * 8,
								// 	})
								// }
							})
						}
					})

				case d.BitsLeft() > 0 &&
					ai.DecodeSamples &&
					aviIsStreamType(typ) &&
					index < len(streams) &&
					streams[index].hasFormat:
					s := streams[index]
					d.FieldFormatLen("data", d.BitsLeft(), s.format, s.formatInArg)
				default:
					d.FieldRawLen("data", d.BitsLeft())
				}

				return false, nil
			}
		},
	)

	if riffType != aviRiffType {
		d.Errorf("wrong or no AVI riff type found (%s)", riffType)
	}

	d.FieldArray("streams", func(d *decode.D) {
		for si, s := range streams {
			d.FieldStruct("stream", func(d *decode.D) {

				var sampleRanges []ranges.Range

				d.FieldArray("indexes", func(d *decode.D) {
					for _, i := range s.indexes {
						d.FieldStruct("index", func(d *decode.D) {
							d.RangeFn(i.Start, i.Len, func(d *decode.D) {
								d.FieldUTF8("type", 4)
								d.FieldU32("cb")
								d.FieldU16("longs_per_entry") // TODO: use?
								d.FieldU8("index_subtype")
								d.FieldU8("index_type", aviIndexTypeNames)
								nEntriesInUse := d.FieldU32("entries_in_use")
								chunkID := d.FieldUTF8("chunk_id", 4)
								aviDecorateStreamID(d, chunkID)
								baseOffset := int64(d.FieldU64("base_offset"))
								d.FieldU32("unused")
								d.FieldArray("index", func(d *decode.D) {
									for i := 0; i < int(nEntriesInUse); i++ {
										d.FieldStruct("index", func(d *decode.D) {
											offset := int64(d.FieldU32("offset"))
											size := int64(d.FieldU32("size"))
											sampleRanges = append(sampleRanges, ranges.Range{
												// TODO: wrong, seem avix files has chunk header on samples?
												Start: baseOffset*8 + offset*8 - 64,
												Len:   size*8 + 64,
											})
										})
									}
								})
							})
						})
					}
				})

				d.FieldArray("samples", func(d *decode.D) {
					for _, sr := range sampleRanges {
						d.RangeFn(sr.Start, sr.Len, func(d *decode.D) {
							if sr.Len > 0 && ai.DecodeSamples && s.hasFormat {
								d.FieldFormat("sample", s.format, s.formatInArg)
							} else {
								d.FieldRawLen("sample", d.BitsLeft())
							}
						})
					}
				})

				_ = si
				_ = moviListPos

				d.FieldArray("idx_samples", func(d *decode.D) {
					for _, is := range idx1Samples {
						if is.streamNr != si {
							continue
						}

						// TODO: palette?
						// TODO: +32? nicer
						d.RangeFn(moviListPos+is.offset+32, is.size, func(d *decode.D) {
							if is.size > 0 && ai.DecodeSamples && s.hasFormat {
								d.FieldFormat("sample", s.format, s.formatInArg)
							} else {
								d.FieldRawLen("sample", d.BitsLeft())
							}
						})
					}
				})
			})
		}
	})

	return nil
}
