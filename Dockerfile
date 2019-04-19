FROM golang:1.8-alpine
ADD . /go/src/tan_server
RUN go install tan_server

FROM alpine:latest
COPY --from=0 /go/bin/tan_server .
ENV PORT 8080
CMD ["./tan_server"]
