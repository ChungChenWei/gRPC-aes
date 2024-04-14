FROM golang:latest as builder

WORKDIR /go/src

COPY src .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/server

FROM alpine:3.19.1 as release

WORKDIR /

COPY --from=builder /go/bin/server /server

EXPOSE 3000

CMD [ "./server" ]