FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o lu_api ./cmd/api

FROM golang:1.24-alpine

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]
