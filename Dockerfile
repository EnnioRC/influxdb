FROM golang:1.10.6 as builder
RUN go get -u github.com/golang/dep/...
WORKDIR /go/src/github.com/EnnioRC/influxdb
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . /go/src/github.com/EnnioRC/influxdb
RUN go install ./cmd/...

FROM debian:stretch
COPY --from=builder /go/bin/* /usr/bin/
COPY --from=builder /go/src/github.com/EnnioRC/influxdb/etc/config.sample.toml /etc/influxdb/influxdb.conf

EXPOSE 8086
VOLUME /var/lib/influxdb

COPY docker/entrypoint.sh /entrypoint.sh
COPY docker/init-influxdb.sh /init-influxdb.sh
ENTRYPOINT ["/entrypoint.sh"]
CMD ["influxd"]
