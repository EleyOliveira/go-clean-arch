FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o orders ./cmd

FROM alpine:latest
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=builder /app/orders .
COPY --from=builder /app/cmd/.env .
COPY --from=builder /app/wait-for-it.sh .
ENTRYPOINT ["./orders"]