
FROM golang:1.20-alpine as builder
RUN apk add --no-cache ca-certificates git
RUN apk add build-base

WORKDIR /app
COPY . .
RUN go mod download

ARG WHO
RUN if [ -z "$WHO" ]; then echo "WHO is not set"; exit 1; fi
RUN go build -ldflags "-X main.who=${WHO}" -o hello .

FROM alpine:3.17.0 as runner
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/hello /hello

EXPOSE 8080
ENTRYPOINT ["/hello"]
