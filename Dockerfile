FROM golang:1.21.1-alpine AS BUILDER

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o api cmd/api/main.go

FROM alpine:latest

WORKDIR /

COPY --from=BUILDER /app/api /api

RUN mkdir /db/

COPY --from=BUILDER /app/db/migrations /db/migrations/

EXPOSE 3000

ENTRYPOINT ["/api"]
