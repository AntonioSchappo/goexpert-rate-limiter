FROM golang:1.22.2 as builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rate-limiter .

FROM scratch
COPY --from=builder /app/rate-limiter .
COPY .env .
EXPOSE 8080
CMD ["./rate-limiter"]