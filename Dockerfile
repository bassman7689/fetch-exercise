FROM golang:1.21.1-alpine AS BUILDER

WORKDIR /app

COPY . .

RUN go build -o api cmd/api/main.go

FROM alpine:latest

WORKDIR /

COPY --from=BUILDER /app/api /api

EXPOSE 3000

ENTRYPOINT ["/api"]
