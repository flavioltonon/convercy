FROM golang:1.18-alpine3.17 as builder

RUN apk update && apk add ca-certificates

WORKDIR /app

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/bin/convercy ./cmd/main.go

FROM scratch

COPY --from=builder /app/build/bin /bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "convercy" ]