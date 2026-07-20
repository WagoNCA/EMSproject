# Compiler (stage 1)

FROM golang:1.26.4 AS builder

WORKDIR /app

COPY .env .
COPY ./EMS-API .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -mod=vendor \
    -o EMS-API .

# Final image (stage 2)

FROM alpine

WORKDIR /app

COPY --from=builder /app/.env .
COPY --from=builder /app/EMS-API .

EXPOSE 8000

CMD ["./EMS-API"]