FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o lu_api ./cmd/api

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/lu_api .

EXPOSE 8000

CMD ["./lu_api"]
