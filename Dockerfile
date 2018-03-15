FROM alpine:3.7

RUN apk --no-cache add \
      curl \
      jq

COPY tag-list.sh tag-list

ENTRYPOINT /tag-list
