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
COPY --from=builder /app/sql/migrations ./sql/migrations
COPY Makefile ./
RUN apk add --no-cache mysql-client make tar curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin    
ENTRYPOINT ["./orders"]