FROM golang:1.26.1 AS builder

WORKDIR /src

COPY go.work go.work.sum ./
COPY common/go.mod ./common/go.mod
COPY common/go.sum ./common/go.sum
COPY http/go.mod ./http/go.mod
COPY http/go.sum ./http/go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /out/server ./http/cmd/main

FROM alpine:3.20 AS runtime

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /out/server /usr/local/bin/server

EXPOSE 9090

CMD ["server"]
