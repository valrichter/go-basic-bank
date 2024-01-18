# Build stage
FROM golang:1.21.6-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start_db.sh .
COPY db/migration ./db/migration
RUN ls -la

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start_db.sh" ]