FROM golang:1.16-alpine as builder
WORKDIR /stats

COPY . .

RUN go build -o output/stats cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /stats/output/stats .

ENTRYPOINT ["./stats"]