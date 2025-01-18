FROM golang:1.23.5 AS go-build

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

COPY ./web/.env.production .env 
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

FROM alpine:latest AS certs
RUN apk --update add ca-certificates

FROM scratch

WORKDIR /app
COPY --from=certs /etc/ssl/certs/ca-certificates.crt \
    /etc/ssl/certs/ca-certificates.crt

COPY --from=go-build \
    /go/src/github.com/Armatorix/mojepole/be/apibin \
    /app/api

COPY --from=node-build \
    /app/dist \
    /app/public

CMD ["/app/api"]