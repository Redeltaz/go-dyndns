FROM golang:alpine AS builder

WORKDIR /builder

COPY . .

RUN go mod download
RUN go build -o go-dyndns

FROM alpine AS runner

RUN apk add --no-cache coreutils

ARG SCHEDULER_CRON

COPY --from=builder /builder/go-dyndns /usr/bin/
RUN mkdir /var/log/go-dyndns

RUN echo "${SCHEDULER_CRON} /usr/bin/go-dyndns 2>> /var/log/go-dyndns/go-dyndns.log" > /var/spool/cron/crontabs/root

CMD crond -f
