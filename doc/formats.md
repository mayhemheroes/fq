## Supported formats

[fq -rn -L . 'include "formats"; formats_table']: sh-start

|Name                              |Description                                                                              |Dependencies|
|-                                 |-                                                                                        |-|
|[`aac_frame`](#aac_frame)         |Advanced&nbsp;Audio&nbsp;Coding&nbsp;frame                                               |<sub></sub>|
|`adts`                            |Audio&nbsp;Data&nbsp;Transport&nbsp;Stream                                               |<sub>`adts_frame`</sub>|
|`adts_frame`                      |Audio&nbsp;Data&nbsp;Transport&nbsp;Stream&nbsp;frame                                    |<sub>`aac_frame`</sub>|
|`amf0`                            |Action&nbsp;Message&nbsp;Format&nbsp;0                                                   |<sub></sub>|
|`apev2`                           |APEv2&nbsp;metadata&nbsp;tag                                                             |<sub>`image`</sub>|
|`ar`                              |Unix&nbsp;archive                                                                        |<sub>`probe`</sub>|
|[`asn1_ber`](#asn1_ber)           |ASN1&nbsp;BER&nbsp;(basic&nbsp;encoding&nbsp;rules,&nbsp;also&nbsp;CER&nbsp;and&nbsp;DER)|<sub></sub>|
|`av1_ccr`                         |AV1&nbsp;Codec&nbsp;Configuration&nbsp;Record                                            |<sub></sub>|
|`av1_frame`                       |AV1&nbsp;frame                                                                           |<sub>`av1_obu`</sub>|
|`av1_obu`                         |AV1&nbsp;Open&nbsp;Bitstream&nbsp;Unit                                                   |<sub></sub>|
|`avc_annexb`                      |H.264/AVC&nbsp;Annex&nbsp;B                                                              |<sub>`avc_nalu`</sub>|
|[`avc_au`](#avc_au)               |H.264/AVC&nbsp;Access&nbsp;Unit                                                          |<sub>`avc_nalu`</sub>|
|`avc_dcr`                         |H.264/AVC&nbsp;Decoder&nbsp;Configuration&nbsp;Record                                    |<sub>`avc_nalu`</sub>|
|`avc_nalu`                        |H.264/AVC&nbsp;Network&nbsp;Access&nbsp;Layer&nbsp;Unit                                  |<sub>`avc_sps` `avc_pps` `avc_sei`</sub>|
|`avc_pps`                         |H.264/AVC&nbsp;Picture&nbsp;Parameter&nbsp;Set                                           |<sub></sub>|
|`avc_sei`                         |H.264/AVC&nbsp;Supplemental&nbsp;Enhancement&nbsp;Information                            |<sub></sub>|
|`avc_sps`                         |H.264/AVC&nbsp;Sequence&nbsp;Parameter&nbsp;Set                                          |<sub></sub>|
|[`avi`](#avi)                     |Audio&nbsp;Video&nbsp;Interleave                                                         |<sub>`avc_au` `mp3_frame` `flac_frame`</sub>|
|[`avro_ocf`](#avro_ocf)           |Avro&nbsp;object&nbsp;container&nbsp;file                                                |<sub></sub>|
|[`bencode`](#bencode)             |BitTorrent&nbsp;bencoding                                                                |<sub></sub>|
|`bitcoin_blkdat`                  |Bitcoin&nbsp;blk.dat                                                                     |<sub>`bitcoin_block`</sub>|
|[`bitcoin_block`](#bitcoin_block) |Bitcoin&nbsp;block                                                                       |<sub>`bitcoin_transaction`</sub>|
|`bitcoin_script`                  |Bitcoin&nbsp;script                                                                      |<sub></sub>|
|`bitcoin_transaction`             |Bitcoin&nbsp;transaction                                                                 |<sub>`bitcoin_script`</sub>|
|[`bplist`](#bplist)               |Apple&nbsp;Binary&nbsp;Property&nbsp;List                                                |<sub></sub>|
|`bsd_loopback_frame`              |BSD&nbsp;loopback&nbsp;frame                                                             |<sub>`inet_packet`</sub>|
|[`bson`](#bson)                   |Binary&nbsp;JSON                                                                         |<sub></sub>|
|`bzip2`                           |bzip2&nbsp;compression                                                                   |<sub>`probe`</sub>|
|[`cbor`](#cbor)                   |Concise&nbsp;Binary&nbsp;Object&nbsp;Representation                                      |<sub></sub>|
|[`csv`](#csv)                     |Comma&nbsp;separated&nbsp;values                                                         |<sub></sub>|
|`dns`                             |DNS&nbsp;packet                                                                          |<sub></sub>|
|`dns_tcp`                         |DNS&nbsp;packet&nbsp;(TCP)                                                               |<sub></sub>|
|`elf`                             |Executable&nbsp;and&nbsp;Linkable&nbsp;Format                                            |<sub></sub>|
|`ether8023_frame`                 |Ethernet&nbsp;802.3&nbsp;frame                                                           |<sub>`inet_packet`</sub>|
|`exif`                            |Exchangeable&nbsp;Image&nbsp;File&nbsp;Format                                            |<sub></sub>|
|`fairplay_spc`                    |FairPlay&nbsp;Server&nbsp;Playback&nbsp;Context                                          |<sub></sub>|
|`flac`                            |Free&nbsp;Lossless&nbsp;Audio&nbsp;Codec&nbsp;file                                       |<sub>`flac_metadatablocks` `flac_frame`</sub>|
|[`flac_frame`](#flac_frame)       |FLAC&nbsp;frame                                                                          |<sub></sub>|
|`flac_metadatablock`              |FLAC&nbsp;metadatablock                                                                  |<sub>`flac_streaminfo` `flac_picture` `vorbis_comment`</sub>|
|`flac_metadatablocks`             |FLAC&nbsp;metadatablocks                                                                 |<sub>`flac_metadatablock`</sub>|
|`flac_picture`                    |FLAC&nbsp;metadatablock&nbsp;picture                                                     |<sub>`image`</sub>|
|`flac_streaminfo`                 |FLAC&nbsp;streaminfo                                                                     |<sub></sub>|
|`gif`                             |Graphics&nbsp;Interchange&nbsp;Format                                                    |<sub></sub>|
|`gzip`                            |gzip&nbsp;compression                                                                    |<sub>`probe`</sub>|
|`hevc_annexb`                     |H.265/HEVC&nbsp;Annex&nbsp;B                                                             |<sub>`hevc_nalu`</sub>|
|[`hevc_au`](#hevc_au)             |H.265/HEVC&nbsp;Access&nbsp;Unit                                                         |<sub>`hevc_nalu`</sub>|
|`hevc_dcr`                        |H.265/HEVC&nbsp;Decoder&nbsp;Configuration&nbsp;Record                                   |<sub>`hevc_nalu`</sub>|
|`hevc_nalu`                       |H.265/HEVC&nbsp;Network&nbsp;Access&nbsp;Layer&nbsp;Unit                                 |<sub>`hevc_vps` `hevc_pps` `hevc_sps`</sub>|
|`hevc_pps`                        |H.265/HEVC&nbsp;Picture&nbsp;Parameter&nbsp;Set                                          |<sub></sub>|
|`hevc_sps`                        |H.265/HEVC&nbsp;Sequence&nbsp;Parameter&nbsp;Set                                         |<sub></sub>|
|`hevc_vps`                        |H.265/HEVC&nbsp;Video&nbsp;Parameter&nbsp;Set                                            |<sub></sub>|
|[`html`](#html)                   |HyperText&nbsp;Markup&nbsp;Language                                                      |<sub></sub>|
|`icc_profile`                     |International&nbsp;Color&nbsp;Consortium&nbsp;profile                                    |<sub></sub>|
|`icmp`                            |Internet&nbsp;Control&nbsp;Message&nbsp;Protocol                                         |<sub></sub>|
|`icmpv6`                          |Internet&nbsp;Control&nbsp;Message&nbsp;Protocol&nbsp;v6                                 |<sub></sub>|
|`id3v1`                           |ID3v1&nbsp;metadata                                                                      |<sub></sub>|
|`id3v11`                          |ID3v1.1&nbsp;metadata                                                                    |<sub></sub>|
|`id3v2`                           |ID3v2&nbsp;metadata                                                                      |<sub>`image`</sub>|
|`ipv4_packet`                     |Internet&nbsp;protocol&nbsp;v4&nbsp;packet                                               |<sub>`ip_packet`</sub>|
|`ipv6_packet`                     |Internet&nbsp;protocol&nbsp;v6&nbsp;packet                                               |<sub>`ip_packet`</sub>|
|`jpeg`                            |Joint&nbsp;Photographic&nbsp;Experts&nbsp;Group&nbsp;file                                |<sub>`exif` `icc_profile`</sub>|
|`json`                            |JavaScript&nbsp;Object&nbsp;Notation                                                     |<sub></sub>|
|`jsonl`                           |JavaScript&nbsp;Object&nbsp;Notation&nbsp;Lines                                          |<sub></sub>|
|[`macho`](#macho)                 |Mach-O&nbsp;macOS&nbsp;executable                                                        |<sub></sub>|
|`macho_fat`                       |Fat&nbsp;Mach-O&nbsp;macOS&nbsp;executable&nbsp;(multi-architecture)                     |<sub>`macho`</sub>|
|[`markdown`](#markdown)           |Markdown                                                                                 |<sub></sub>|
|[`matroska`](#matroska)           |Matroska&nbsp;file                                                                       |<sub>`aac_frame` `av1_ccr` `av1_frame` `avc_au` `avc_dcr` `flac_frame` `flac_metadatablocks` `hevc_au` `hevc_dcr` `image` `mp3_frame` `mpeg_asc` `mpeg_pes_packet` `mpeg_spu` `opus_packet` `vorbis_packet` `vp8_frame` `vp9_cfm` `vp9_frame`</sub>|
|[`mp3`](#mp3)                     |MP3&nbsp;file                                                                            |<sub>`id3v2` `id3v1` `id3v11` `apev2` `mp3_frame`</sub>|
|`mp3_frame`                       |MPEG&nbsp;audio&nbsp;layer&nbsp;3&nbsp;frame                                             |<sub>`xing`</sub>|
|[`mp4`](#mp4)                     |ISOBMFF,&nbsp;QuickTime&nbsp;and&nbsp;similar                                            |<sub>`aac_frame` `av1_ccr` `av1_frame` `avc_au` `avc_dcr` `flac_frame` `flac_metadatablocks` `hevc_au` `hevc_dcr` `icc_profile` `id3v2` `image` `jpeg` `mp3_frame` `mpeg_es` `mpeg_pes_packet` `opus_packet` `prores_frame` `protobuf_widevine` `pssh_playready` `vorbis_packet` `vp9_frame` `vpx_ccr`</sub>|
|`mpeg_asc`                        |MPEG-4&nbsp;Audio&nbsp;Specific&nbsp;Config                                              |<sub></sub>|
|`mpeg_es`                         |MPEG&nbsp;Elementary&nbsp;Stream                                                         |<sub>`mpeg_asc` `vorbis_packet`</sub>|
|`mpeg_pes`                        |MPEG&nbsp;Packetized&nbsp;elementary&nbsp;stream                                         |<sub>`mpeg_pes_packet` `mpeg_spu`</sub>|
|`mpeg_pes_packet`                 |MPEG&nbsp;Packetized&nbsp;elementary&nbsp;stream&nbsp;packet                             |<sub></sub>|
|`mpeg_spu`                        |Sub&nbsp;Picture&nbsp;Unit&nbsp;(DVD&nbsp;subtitle)                                      |<sub></sub>|
|`mpeg_ts`                         |MPEG&nbsp;Transport&nbsp;Stream                                                          |<sub></sub>|
|[`msgpack`](#msgpack)             |MessagePack                                                                              |<sub></sub>|
|`ogg`                             |OGG&nbsp;file                                                                            |<sub>`ogg_page` `vorbis_packet` `opus_packet` `flac_metadatablock` `flac_frame`</sub>|
|`ogg_page`                        |OGG&nbsp;page                                                                            |<sub></sub>|
|`opus_packet`                     |Opus&nbsp;packet                                                                         |<sub>`vorbis_comment`</sub>|
|[`pcap`](#pcap)                   |PCAP&nbsp;packet&nbsp;capture                                                            |<sub>`link_frame` `tcp_stream` `ipv4_packet`</sub>|
|`pcapng`                          |PCAPNG&nbsp;packet&nbsp;capture                                                          |<sub>`link_frame` `tcp_stream` `ipv4_packet`</sub>|
|`png`                             |Portable&nbsp;Network&nbsp;Graphics&nbsp;file                                            |<sub>`icc_profile` `exif`</sub>|
|`prores_frame`                    |Apple&nbsp;ProRes&nbsp;frame                                                             |<sub></sub>|
|[`protobuf`](#protobuf)           |Protobuf                                                                                 |<sub></sub>|
|`protobuf_widevine`               |Widevine&nbsp;protobuf                                                                   |<sub>`protobuf`</sub>|
|`pssh_playready`                  |PlayReady&nbsp;PSSH                                                                      |<sub></sub>|
|`raw`                             |Raw&nbsp;bits                                                                            |<sub></sub>|
|[`rtmp`](#rtmp)                   |Real-Time&nbsp;Messaging&nbsp;Protocol                                                   |<sub>`amf0` `mpeg_asc`</sub>|
|`sll2_packet`                     |Linux&nbsp;cooked&nbsp;capture&nbsp;encapsulation&nbsp;v2                                |<sub>`inet_packet`</sub>|
|`sll_packet`                      |Linux&nbsp;cooked&nbsp;capture&nbsp;encapsulation                                        |<sub>`inet_packet`</sub>|
|`tar`                             |Tar&nbsp;archive                                                                         |<sub>`probe`</sub>|
|`tcp_segment`                     |Transmission&nbsp;control&nbsp;protocol&nbsp;segment                                     |<sub></sub>|
|`tiff`                            |Tag&nbsp;Image&nbsp;File&nbsp;Format                                                     |<sub>`icc_profile`</sub>|
|`toml`                            |Tom's&nbsp;Obvious,&nbsp;Minimal&nbsp;Language                                           |<sub></sub>|
|`udp_datagram`                    |User&nbsp;datagram&nbsp;protocol                                                         |<sub>`udp_payload`</sub>|
|`vorbis_comment`                  |Vorbis&nbsp;comment                                                                      |<sub>`flac_picture`</sub>|
|`vorbis_packet`                   |Vorbis&nbsp;packet                                                                       |<sub>`vorbis_comment`</sub>|
|`vp8_frame`                       |VP8&nbsp;frame                                                                           |<sub></sub>|
|`vp9_cfm`                         |VP9&nbsp;Codec&nbsp;Feature&nbsp;Metadata                                                |<sub></sub>|
|`vp9_frame`                       |VP9&nbsp;frame                                                                           |<sub></sub>|
|`vpx_ccr`                         |VPX&nbsp;Codec&nbsp;Configuration&nbsp;Record                                            |<sub></sub>|
|[`wasm`](#wasm)                   |WebAssembly&nbsp;Binary&nbsp;Format                                                      |<sub></sub>|
|`wav`                             |WAV&nbsp;file                                                                            |<sub>`id3v2` `id3v1` `id3v11`</sub>|
|`webp`                            |WebP&nbsp;image                                                                          |<sub>`vp8_frame`</sub>|
|`xing`                            |Xing&nbsp;header                                                                         |<sub></sub>|
|[`xml`](#xml)                     |Extensible&nbsp;Markup&nbsp;Language                                                     |<sub></sub>|
|`yaml`                            |YAML&nbsp;Ain't&nbsp;Markup&nbsp;Language                                                |<sub></sub>|
|[`zip`](#zip)                     |ZIP&nbsp;archive                                                                         |<sub>`probe`</sub>|
|`image`                           |Group                                                                                    |<sub>`gif` `jpeg` `mp4` `png` `tiff` `webp`</sub>|
|`inet_packet`                     |Group                                                                                    |<sub>`ipv4_packet` `ipv6_packet`</sub>|
|`ip_packet`                       |Group                                                                                    |<sub>`icmp` `icmpv6` `tcp_segment` `udp_datagram`</sub>|
|`link_frame`                      |Group                                                                                    |<sub>`bsd_loopback_frame` `ether8023_frame` `sll2_packet` `sll_packet`</sub>|
|`probe`                           |Group                                                                                    |<sub>`adts` `ar` `avi` `avro_ocf` `bitcoin_blkdat` `bplist` `bzip2` `elf` `flac` `gif` `gzip` `jpeg` `json` `jsonl` `macho` `macho_fat` `matroska` `mp3` `mp4` `mpeg_ts` `ogg` `pcap` `pcapng` `png` `tar` `tiff` `toml` `wasm` `wav` `webp` `xml` `yaml` `zip`</sub>|
|`tcp_stream`                      |Group                                                                                    |<sub>`dns_tcp` `rtmp`</sub>|
|`udp_payload`                     |Group                                                                                    |<sub>`dns`</sub>|

[#]: sh-end

## Global format options

Currently the only global option is `force` and is used to ignore some format assertion errors. It can be used as a decode option or as a CLI `-o` option:

```
fq -d mp4 -o force=true file.mp4
fq -d raw 'mp4({force: true})' file.mp4
```

## Format details

[fq -rn -L . 'include "formats"; formats_sections']: sh-start

## aac_frame

### Options

|Name         |Default|Description|
|-            |-      |-|
|`object_type`|1      |Audio object type|

### Examples

Decode file using aac_frame options
```
$ fq -d aac_frame -o object_type=1 . file
```

Decode value as aac_frame
```
... | aac_frame({object_type:1})
```

## asn1_ber

Supports decoding BER, CER and DER (X.690).

- Currently no extra validation is done for CER and DER.
- Does not support specifying a schema.
- Supports `torepr` but without schema all sequences and sets will be arrays.

### Can be used to decode certificates etc

```sh
$ fq -d raw 'frompem | asn1_ber | d' cert.pem
```

### Can decode nested values

```sh
$ fq -d asn1_ber '.constructed[1].value | asn1_ber' file.ber
```

### Manual schema

```sh
$ fq -d asn1_ber 'torepr as $r | ["version", "modulus", "private_exponent", "private_exponen", "prime1", "prime2", "exponent1", "exponent2", "coefficient"] | with_entries({key: .value, value: $r[.key]})' pkcs1.der
```

### References
- https://www.itu.int/ITU-T/studygroups/com10/languages/X.690_1297.pdf
- https://en.wikipedia.org/wiki/X.690
- https://letsencrypt.org/docs/a-warm-welcome-to-asn1-and-der/
- https://lapo.it/asn1js/

## avc_au

### Options

|Name         |Default|Description|
|-            |-      |-|
|`length_size`|0      |Length value size|

### Examples

Decode file using avc_au options
```
$ fq -d avc_au -o length_size=0 . file
```

Decode value as avc_au
```
... | avc_au({length_size:0})
```

## avi

### Options

|Name            |Default|Description|
|-               |-      |-|
|`decode_samples`|true   |Decode supported media samples|

### Examples

Decode file using avi options
```
$ fq -d avi -o decode_samples=true . file
```

Decode value as avi
```
... | avi({decode_samples:true})
```

## avro_ocf

Supports reading Avro Object Container Format (OCF) files based on the 1.11.0 specification.

Capable of handling null, deflate, and snappy codecs for data compression.

Limitations:

- Schema does not support self-referential types, only built-in types.
- Decimal logical types are not supported for decoding, will just be treated as their primitive type

### References
- https://avro.apache.org/docs/current/spec.html#Object+Container+Files

### Authors
- Xentripetal
xentripetal@fastmail.com
[@xentripetal](https://github.com/xentripetal)

## bencode

### Convert represented value to JSON

```
$ fq -d bencode torepr file.torrent
```

### References
- https://wiki.theory.org/BitTorrentSpecification#Bencoding

## bitcoin_block

### Options

|Name        |Default|Description|
|-           |-      |-|
|`has_header`|false  |Has blkdat header|

### Examples

Decode file using bitcoin_block options
```
$ fq -d bitcoin_block -o has_header=false . file
```

Decode value as bitcoin_block
```
... | bitcoin_block({has_header:false})
```

## bplist

### Show full decoding
```sh
$ fq d Info.plist
```

### Timestamps
Timestamps in Apple Binary Property Lists are encoded as Cocoa Core Data
timestamps, where the raw value is the floating point number of seconds since
January 1, 2001. By default, `fq` will render the raw floating point value. In
order to get the raw value or string description, use the `todescription`
function, you can use the `tovalue` and `todescription` functions:

```sh
$ fq 'torepr.SomeTimeStamp | tovalue' Info.plist
685135328

$ fq 'torepr.SomeTimeStamp | todescription' Info.plist
"2022-09-17T19:22:08Z"
```


### Get JSON representation
```sh
$ fq torepr com.apple.UIAutomation.plist
{
  "UIAutomationEnabled": true
}
```

### Authors
- David McDonald
[@dgmcdona](https://github.com/dgmcdona)

### References
- http://fileformats.archiveteam.org/wiki/Property_List/Binary
- https://medium.com/@karaiskc/understanding-apples-binary-property-list-format-281e6da00dbd
- https://opensource.apple.com/source/CF/CF-550/CFBinaryPList.c

## bson

### Convert represented value to JSON

```
$ fq -d bson torepr file.bson
```

### Filter represented value

```
$ fq -d bson 'torepr | select(.name=="bob")' file.bson
```

### References
- https://bsonspec.org/spec.html

## cbor

### Convert represented value to JSON

```
$ fq -d cbor torepr file.cbor
```

### References
- https://en.wikipedia.org/wiki/CBOR
- https://www.rfc-editor.org/rfc/rfc8949.html

## csv

### Options

|Name     |Default|Description|
|-        |-      |-|
|`comma`  |,      |Separator character|
|`comment`|#      |Comment line character|

### Examples

Decode file using csv options
```
$ fq -d csv -o comma="," -o comment="#" . file
```

Decode value as csv
```
... | csv({comma:",",comment:"#"})
```

### TSV to CSV

```sh
$ fq -d csv -o comma="\t" tocsv file.tsv
```

### Convert rows to objects based on header row

```sh
$ fq -d csv '.[0] as $t | .[1:] | map(with_entries(.key = $t[.key]))' file.csv
```

## flac_frame

### Options

|Name             |Default|Description|
|-                |-      |-|
|`bits_per_sample`|16     |Bits per sample|

### Examples

Decode file using flac_frame options
```
$ fq -d flac_frame -o bits_per_sample=16 . file
```

Decode value as flac_frame
```
... | flac_frame({bits_per_sample:16})
```

## hevc_au

### Options

|Name         |Default|Description|
|-            |-      |-|
|`length_size`|4      |Length value size|

### Examples

Decode file using hevc_au options
```
$ fq -d hevc_au -o length_size=4 . file
```

Decode value as hevc_au
```
... | hevc_au({length_size:4})
```

## html

### Options

|Name              |Default|Description|
|-                 |-      |-|
|`array`           |false  |Decode as nested arrays|
|`attribute_prefix`|@      |Prefix for attribute keys|
|`seq`             |false  |Use seq attribute to preserve element order|

### Examples

Decode file using html options
```
$ fq -d html -o array=false -o attribute_prefix="@" -o seq=false . file
```

Decode value as html
```
... | html({array:false,attribute_prefix:"@",seq:false})
```

HTML is decoded in HTML5 mode and will always include `<html>`, `<body>` and `<head>` element.

See xml format for more examples and how to preserve element order and how to encode to xml.

There is no `tohtml` function, see `toxml` instead.

### Element as object

```sh
# decode as object is the default
$ echo '<a href="url">text</a>' | fq -d html
{
  "html": {
    "body": {
      "a": {
        "#text": "text",
        "@href": "url"
      }
    },
    "head": ""
  }
}
```

### Element as array

```sh
$ '<a href="url">text</a>' | fq -d html -o array=true
[
  "html",
  null,
  [
    [
      "head",
      null,
      []
    ],
    [
      "body",
      null,
      [
        [
          "a",
          {
            "#text": "text",
            "href": "url"
          },
          []
        ]
      ]
    ]
  ]
]

# decode html files to a {file: "title", ...} object
$ fq -n -d html '[inputs | {key: input_filename, value: .html.head.title?}] | from_entries' *.html

# <a> href:s in file
$ fq -r -o array=true -d html '.. | select(.[0] == "a" and .[1].href)?.[1].href' file.html
```

## macho

Supports decoding vanilla and FAT Mach-O binaries.

### Select 64bit load segments

```sh
$ fq '.load_commands[] | select(.cmd=="segment_64")' file
```

### References
- https://github.com/aidansteele/osx-abi-macho-file-format-reference

### Authors
- Sıddık AÇIL
acils@itu.edu.tr
[@Akaame](https://github.com/Akaame)

## markdown

### Array with all level 1 and 2 headers
```sh
$ fq -d markdown '[.. | select(.type=="heading" and .level<=2)?.children[0]]' file.md
```
## matroska

### Lookup element using path

```sh
$ fq 'matroska_path(".Segment.Tracks[0)")' file.mkv
```

### Get path to element

```sh
$ fq 'grep_by(.id == "Tracks") | matroska_path' file.mkv
```

### References
- https://tools.ietf.org/html/draft-ietf-cellar-ebml-00
- https://matroska.org/technical/specs/index.html
- https://www.matroska.org/technical/basics.html
- https://www.matroska.org/technical/codec_specs.html
- https://wiki.xiph.org/MatroskaOpus

## mp3

### Options

|Name                       |Default|Description|
|-                          |-      |-|
|`max_sync_seek`            |32768  |Max byte distance to next sync|
|`max_unique_header_configs`|5      |Max number of unique frame header configs allowed|

### Examples

Decode file using mp3 options
```
$ fq -d mp3 -o max_sync_seek=32768 -o max_unique_header_configs=5 . file
```

Decode value as mp3
```
... | mp3({max_sync_seek:32768,max_unique_header_configs:5})
```

## mp4

### Options

|Name             |Default|Description|
|-                |-      |-|
|`allow_truncated`|false  |Allow box to be truncated|
|`decode_samples` |true   |Decode supported media samples|

### Examples

Decode file using mp4 options
```
$ fq -d mp4 -o allow_truncated=false -o decode_samples=true . file
```

Decode value as mp4
```
... | mp4({allow_truncated:false,decode_samples:true})
```

### Lookup mp4 box using a mp4 box path.

```sh
# <decode value box> | mp4_path($path) -> <decode value box>
$ fq 'mp4_path(".moov.trak[1]")' file.mp4
```

### Get mp4 box path for a decode value box.

```sh
# <decode value box> | mp4_path -> string
$ fq 'grep_by(.type == "trak") | mp4_path' file.mp4
```

### Force decode a single box

```sh
$ fq -n '"AAAAHGVsc3QAAAAAAAAAAQAAADIAAAQAAAEAAA==" | frombase64 | mp4({force:true}) | d'
```

### Speed up decoding by not decoding samples

```sh
# manually decode first sample as a aac_frame
$ fq -o decode_samples=false '.tracks[0].samples[0] | aac_frame | d' file.mp4
```

### Entries for first edit list as values

```sh
$ fq 'first(grep_by(.type=="elst").entries) | tovalue' file.mp4
```

### References

- [ISO/IEC base media file format (MPEG-4 Part 12)](https://en.wikipedia.org/wiki/ISO/IEC_base_media_file_format)
- [Quicktime file format](https://developer.apple.com/standards/qtff-2001.pdf)

## msgpack

### Convert represented value to JSON

```
$ fq -d msgpack torepr file.msgpack
```

### References
- https://github.com/msgpack/msgpack/blob/master/spec.md

## pcap

### Build object with number of (reassembled) TCP bytes sent to/from client IP
```sh
# for a pcapng file you would use .[0].tcp_connections for first section
$ fq '.tcp_connections | group_by(.client.ip) | map({key: .[0].client.ip, value: map(.client.stream, .server.stream | tobytes.size) | add}) | from_entries'
{
  "10.1.0.22": 15116,
  "10.99.12.136": 234,
  "10.99.12.150": 218
}
```
## protobuf

### Can decode sub messages

```sh
$ fq -d protobuf '.fields[6].wire_value | protobuf | d' file
```

### References
- https://developers.google.com/protocol-buffers/docs/encoding

## rtmp

Current only supports plain RTMP (not RTMPT or encrypted variants etc) with AMF0 (not AMF3).

### Show rtmp streams in PCAP file
```sh
fq '.tcp_connections[] | select(.server.port=="rtmp") | d' file.cap
```

### References
- https://rtmp.veriskope.com/docs/spec/
- https://rtmp.veriskope.com/pdf/video_file_format_spec_v10.pdf

## wasm

### Count opcode usage
```sh
$ fq '.sections[] | select(.id == "code_section") | [.. | .opcode? // empty] | count | map({key: .[0], value: .[1]}) | from_entries' file.wasm
```

### List exports and imports
```sh
$ fq '.sections | {import: map(select(.id == "import_section").content.im.x[].nm.b), export: map(select(.id == "export_section").content.ex.x[].nm.b)}' file.wasm
```

### Authors
- Takashi Oguma
[@bitbears-dev](https://github.com/bitbears-dev)
[@0xb17bea125](https://twitter.com/0xb17bea125)

### References
- https://webassembly.github.io/spec/core/

## xml

### Options

|Name              |Default|Description|
|-                 |-      |-|
|`array`           |false  |Decode as nested arrays|
|`attribute_prefix`|@      |Prefix for attribute keys|
|`seq`             |false  |Use seq attribute to preserve element order|

### Examples

Decode file using xml options
```
$ fq -d xml -o array=false -o attribute_prefix="@" -o seq=false . file
```

Decode value as xml
```
... | xml({array:false,attribute_prefix:"@",seq:false})
```

XML can be decoded and encoded into jq values in two ways, elements as object or array.
Which variant to use depends a bit what you want to do. The object variant might be easier
to query for a specific value but array might be easier to use to generate xml or to query
after all elements of some kind etc.

Encoding is done using the `toxml` function and it will figure what variant that is used based on the input value.
Is has two optional options `indent` and `attribute_prefix`.

### Elements as object

Element can have different shapes depending on body text, attributes and children:

- `<a key="value">text</a>` is `{"a":{"#text":"text","@key":"value"}}`, has text (`#text`) and attributes (`@key`)
- `<a>text</a>` is `{"a":"text"}`
- `<a><b>text</b></a>` is `{"a":{"b":"text"}}` one child with only text and no attributes
- `<a><b/><b>text</b></a>` is `{"a":{"b":["","text"]}}` two children with same name end up in an array
- `<a><b/><b key="value">text</b></a>` is `{"a":{"b":["",{"#text":"text","@key":"value"}]}}`

If there is `#seq` attribute it encodes the child element order. Use `-o seq=true` to include sequence number when decoding,
otherwise order might be lost.

```sh
# decode as object is the default
$ echo '<a><b/><b>bbb</b><c attr="value">ccc</c></a>' | fq -d xml -o seq=true
{
  "a": {
    "b": [
      {
        "#seq": 0
      },
      {
        "#seq": 1,
        "#text": "bbb"
      }
    ],
    "c": {
      "#seq": 2,
      "#text": "ccc",
      "@attr": "value"
    }
  }
}

# access text of the <c> element
$ echo '<a><b/><b>bbb</b><c attr="value">ccc</c></a>' | fq '.a.c["#text"]'
"ccc"

# decode to object and encode to xml
$ echo '<a><b/><b>bbb</b><c attr="value">ccc</c></a>' | fq -r -d xml -o seq=true 'toxml({indent:2})'
<a>
  <b></b>
  <b>bbb</b>
  <c attr="value">ccc</c>
</a>
```

### Elements as array

Elements are arrays of the shape `["#text": "body text", "attr_name", {key: "attr value"}|null, [<child element>, ...]]`.

```sh
# decode as array
$ echo '<a><b/><b>bbb</b><c attr="value">ccc</c></a>' | fq -d xml -o array=true
[
  "a",
  null,
  [
    [
      "b",
      null,
      []
    ],
    [
      "b",
      {
        "#text": "bbb"
      },
      []
    ],
    [
      "c",
      {
        "#text": "ccc",
        "attr": "value"
      },
      []
    ]
  ]
]

# decode to array and encode to xml
$ echo '<a><b/><b>bbb</b><c attr="value">ccc</c></a>' | fq -r -d xml -o array=true -o seq=true 'toxml({indent:2})'
<a>
  <b></b>
  <b>bbb</b>
  <c attr="value">ccc</c>
</a>

# access text of the <c> element, the object variant above is probably easier to use
$ echo '<a><b/><b>bbb</b><c attr="value">ccc</c></a>' | fq -o array=true '.[2][2][1]["#text"]'
"ccc"
```

### References
- [xml.com's Converting Between XML and JSON](https://www.xml.com/pub/a/2006/05/31/converting-between-xml-and-json.html)

## zip

### Options

|Name        |Default|Description|
|-           |-      |-|
|`uncompress`|true   |Uncompress and probe files|

### Examples

Decode file using zip options
```
$ fq -d zip -o uncompress=true . file
```

Decode value as zip
```
... | zip({uncompress:true})
```

Supports ZIP64.

### References
- https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT
- https://opensource.apple.com/source/zip/zip-6/unzip/unzip/proginfo/extra.fld


[#]: sh-end

## Dependency graph

![alt text](formats.svg "Format diagram")
