FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o orders ./cmd

FROM scratch
WORKDIR /app
COPY --from=builder /app/orders .
COPY --from=builder /app/cmd/.env .
ENTRYPOINT ["./orders"]