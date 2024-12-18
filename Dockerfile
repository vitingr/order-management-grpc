FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.work go.work.sum ./

COPY ./common ./common
COPY ./gateway ./gateway
COPY ./kitchen ./kitchen
COPY ./orders ./orders
COPY ./payments ./payments
COPY ./stock ./stock

RUN go mod tidy

RUN go build -o /app/bin/gateway ./gateway
RUN go build -o /app/bin/orders ./orders
RUN go build -o /app/bin/payments ./payments
RUN go build -o /app/bin/kitchen ./kitchen
RUN go build -o /app/bin/stock ./stock

FROM debian:bullseye-slim AS runtime

WORKDIR /app

COPY --from=builder /app/bin /app/bin
