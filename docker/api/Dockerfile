FROM alpine:3.6 as alpine

RUN apk add -U --no-cache ca-certificates

FROM scratch

MAINTAINER meguro@code-house.jp

COPY ./api /api
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8888
ENTRYPOINT ["/api"]
