FROM alpine:latest

# Instal dependensi runtime
RUN apk add --no-cache \
    ca-certificates \
    sqlite-libs \
    procps

WORKDIR /app

CMD ["./xiaozhi-server"]
