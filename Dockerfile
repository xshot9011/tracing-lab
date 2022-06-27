FROM golang:1.18 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o main .

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 80
ENTRYPOINT [ "/app/main" ]
