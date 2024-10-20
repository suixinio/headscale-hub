FROM golang:1.23.0-alpine AS builder

WORKDIR /hub
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o headscale-hub .

FROM alpine:latest

RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /hub

COPY --from=builder /hub/headscale-hub /bin/headscale-hub
EXPOSE 8088/tcp
CMD ["headscale-hub", "-c", "/hub/config.yml"]
