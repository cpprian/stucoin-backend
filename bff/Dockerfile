FROM golang:1.21.6-alpine AS builder

RUN adduser -D -g '' appuser

WORKDIR /opt/app

COPY go.mod go.sum ./

RUN go mod download && \
    go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/stucoin-bff ./cmd/app


FROM alpine:3.19.0
LABEL language="golang"
LABEL org.opencontainers.image.source https://github.com/cpprian/blog_posts_with_markdown

COPY --from=builder /etc/passwd /etc/passwd

WORKDIR /opt/app

COPY --from=builder --chown=appuser:1000 /go/bin/stucoin-bff stucoin-bff

ENTRYPOINT ["./stucoin-bff"]