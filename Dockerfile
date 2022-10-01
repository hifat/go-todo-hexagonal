# Build stage
FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
ENV TZ=Asia/Bangkok
RUN go build -o main main.go

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY start.sh .
COPY wait-for.sh .

EXPOSE 3000

CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]
