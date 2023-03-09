FROM golang:alpine AS builder

WORKDIR /builder

COPY . .

RUN go build -o go-dyndns

FROM alpine AS runner

COPY --from=builder /builder/go-dyndns /usr/bin/

ENTRYPOINT ["go-dyndns"]
