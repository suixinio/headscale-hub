FROM registry.cn-qingdao.aliyuncs.com/ohmymc/node:20.18.0-bullseye AS builder
WORKDIR /hub-ui/
COPY ./package.json ./package.json
RUN npm --registry https://registry.npmmirror.com install
COPY ./ ./
RUN npm run build:prod

FROM registry.cn-qingdao.aliyuncs.com/ohmymc/nginx:1.25.3-alpine
COPY --from=builder /hub-ui/dist/ /usr/share/nginx/html/
