FROM golang:1.23.4 AS go-build


WORKDIR /go/src/github.com/Armatorix/mojepole/be
COPY ./be/go.mod \
    ./be/go.sum \
    ./
RUN go mod download
COPY ./be ./
RUN CGO_ENABLED=0 go build -o apibin

FROM oven/bun:1.1  AS node-build
WORKDIR /app/
COPY ./web/package.json \
    ./web/bun.lockb ./

RUN bun i

COPY ./web/public public
COPY ./web/index.html \
    ./web/tsconfig.json \
    ./web/tsconfig.node.json \
    ./web/vite.config.ts \
    ./web/tailwind.config.cjs \
    ./web/postcss.config.cjs \
    ./
COPY ./web/src src


RUN bun run build

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=go-build \
    /go/src/github.com/Armatorix/mojepole/be/apibin \
    /app/api

COPY --from=node-build \
    /app/dist \
    /app/public

RUN apt-get update -y && apt-get install ca-certificates -y

CMD ["/app/api"]