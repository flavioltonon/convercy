FROM golang:1.18-alpine3.17 as builder

RUN apk update && apk add \
    make \
    ca-certificates

WORKDIR /app

COPY . .

RUN make build

FROM scratch

COPY --from=builder /app/build/bin /bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "convercy" ]