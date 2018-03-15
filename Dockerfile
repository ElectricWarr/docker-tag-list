FROM alpine:3.7

RUN apk --no-cache add \
      bash \
      curl \
      jq \
      util-linux

COPY tag-list.sh tag-list

ENTRYPOINT ["/bin/bash","./tag-list"]
