FROM golang:1.22-alpine AS builder
LABEL maintainer="arelive <me@are.moe> (are.moe)"

WORKDIR /build

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/web ./cmd/web/

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/ /app/

EXPOSE ${PORT}

ENTRYPOINT ["./web"]
