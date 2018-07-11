FROM alpine:3.7

RUN apk --no-cache add \
      bash \
      curl \
      jq \
      util-linux

COPY tag-list.sh /bin/tag-list

ENTRYPOINT ["tag-list"]
