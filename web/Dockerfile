FROM node:21-alpine AS builder
LABEL maintainer="arelive <me@are.moe> (are.moe)"

WORKDIR /build

COPY . .
RUN npm install

RUN npm run build

FROM nginx:1.25-alpine

COPY --from=builder /build/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
