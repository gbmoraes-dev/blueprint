FROM golang:1.26.1-alpine AS builder

RUN apk add --no-cache ca-certificates tzdata

RUN adduser -D -g '' nonroot

WORKDIR /app

COPY go.mod go.sum* ./

RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/server/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /etc/group /etc/group

WORKDIR /app

COPY --from=builder /app/server .

USER nonroot

EXPOSE 8080

CMD ["./server"]
