FROM golang:1.21.6-alpine AS builder

RUN adduser -D -g '' elf

WORKDIR /opt/app/
COPY go.mod go.sum ./

RUN go mod download && \
    go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/stucoin-rewards ./cmd/app


FROM alpine:3.19.0
LABEL language="golang"
LABEL org.opencontainers.image.source https://github.com/cpprian/stucoin-backend

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder --chown=elf:1000 /go/bin/stucoin-rewards /stucoin-rewards

USER elf

ENTRYPOINT ["./stucoin-rewards"]