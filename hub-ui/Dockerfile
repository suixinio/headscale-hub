FROM node:20.18.0-bullseye as builder
WORKDIR /hub-ui/
COPY ./package-lock.json ./package-lock.json
COPY ./package.json ./package.json
RUN npm install
COPY ./ ./
RUN npm run build:prod

FROM nginx:1.25.3-alpine
COPY --from=builder /hub-ui/dist/ /usr/share/nginx/html/
