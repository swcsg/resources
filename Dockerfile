FROM alpine:latest

ADD main /

EXPOSE 8080

ENTRYPOINT ["/main"]