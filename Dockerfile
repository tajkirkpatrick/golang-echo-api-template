FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

COPY *.go .

RUN go build -o ./api-server

EXPOSE 8080

CMD ["/api-server"]

# Minimize Container

FROM alpine:latest

WORKDIR /go/bin/api-server

RUN apk --no-cache add curl

COPY --from=builder /app/api-server .

HEALTHCHECK CMD curl --fail http://localhost:8080/healthcheck || exit 1

ENTRYPOINT ["./api-server"]