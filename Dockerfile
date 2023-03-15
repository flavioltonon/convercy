FROM golang:1.18-alpine3.17 as builder

RUN apk update && apk add make

WORKDIR /app

COPY . .

RUN make build

FROM scratch

COPY --from=builder /app/build/bin /bin
COPY --from=builder /app/config.yaml /config.yaml

CMD [ "convercy" ]