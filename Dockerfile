From golang:1.15-alpine3.12 as builder
COPY main /app/
COPY config.yaml /app/config/config.yaml
CMD ["/app/main"]