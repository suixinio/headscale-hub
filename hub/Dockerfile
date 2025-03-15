FROM registry.cn-qingdao.aliyuncs.com/ohmymc/golang:1.23.0-alpine AS builder
WORKDIR /hub
RUN apk update
RUN apk add --no-cache build-base
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o headscale-hub .

FROM registry.cn-qingdao.aliyuncs.com/ohmymc/alpine:3.21.3
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
WORKDIR /hub
COPY --from=builder /hub/headscale-hub /bin/headscale-hub
EXPOSE 8088/tcp
CMD ["headscale-hub", "-c", "/hub/config.yml"]
