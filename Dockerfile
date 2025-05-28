# Build stage
FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Final image
FROM debian:bookworm-slim
COPY --from=builder /app/main /main
EXPOSE 4444
CMD ["/main"]

