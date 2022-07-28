# Build stage
FROM  golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# --/-- For Windows users
RUN apk add dos2unix --update-cache --repository http://dl-3.alpinelinux.org/alpine/edge/community/ --allow-untrusted
RUN dos2unix wait-for.sh start.sh
# --/--

RUN chmod +x wait-for.sh start.sh
RUN go build -o simple ./cmd/app/main.go
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/simple .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY --from=builder /app/start.sh .
COPY --from=builder /app/wait-for.sh .
COPY ./configs ./configs
COPY .env .
COPY ./internal/migrations ./migrations

EXPOSE 8080
CMD [ "/app/simple" ]
ENTRYPOINT [ "/app/start.sh" ]