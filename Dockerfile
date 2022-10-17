FROM golang:1.19.1-buster as go-target
ADD . /fq
WORKDIR /fq
RUN go build

FROM golang:1.19.1-buster
COPY --from=go-target /fq/fq /
COPY --from=go-target /fq/fq /testsuite/
COPY --from=go-target /usr/bin/whoami /testsuite/
ENTRYPOINT []
CMD ["/fq", ".",  "@@"]
