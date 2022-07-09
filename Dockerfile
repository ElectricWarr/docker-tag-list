FROM golang:1.18-alpine3.16 AS build
RUN apk --no-cache add \
      curl \
      jq
WORKDIR /go/src/tag-list
# RUN go mod init && go mod tidy
COPY golang/go.mod golang/go.sum ./
RUN go mod download
COPY golang/ ./
# RUN go mod init && go mod tidy
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -o /go/bin/tag-list

# FROM scratch AS run
FROM golang:1.18-alpine3.16 AS run
WORKDIR /
COPY --from=build /go/bin/tag-list /bin/tag-list
ENTRYPOINT ["tag-list"]
CMD ["--help"]
