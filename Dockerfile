# Build Stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY wait-for.sh .
COPY app.env .
COPY start.sh .


EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT ["/app/start.sh"]