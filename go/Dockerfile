FROM golang:1.13-alpine as builder

COPY . /go
WORKDIR /go/src/rocks

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rocks .

FROM scratch

COPY --from=builder /go/src/rocks/rocks /

EXPOSE 8000
CMD ["/rocks"]