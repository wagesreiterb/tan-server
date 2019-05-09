FROM golang:1.8-alpine
ADD . /go/src/tan-server
RUN go install tan-server

FROM alpine:latest
COPY --from=0 /go/bin/tan-server .
ENV PORT 8080
CMD ["./tan-server"]
