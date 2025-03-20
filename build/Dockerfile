FROM golang:1.24.1-alpine as builder

RUN mkdir /build
WORKDIR /build

COPY . . 

RUN apk update && apk upgrade && \
    apk add --no-cache \
    musl-dev \
    gcc \
    build-base \
    libc-dev \
    ca-certificates \
    wget


RUN ls /build
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /build/application cmd/*.go


# ------------------------------------------------------------------------------

FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    tzdata ca-certificates && \
    cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime; \
    echo "America/Sao_Paulo" > /etc/timezone

COPY --from=builder /build/application .

# executable
ENTRYPOINT [ "./application" ]

# arguments that can be overridden
CMD [ "./application" ]
