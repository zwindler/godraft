FROM golang:1.18-alpine AS builder
ARG VERSION

WORKDIR /bin

COPY / .
RUN apk add build-base && go mod download && go build -o godraft \
    -ldflags "-X main.Version=$VERSION" main.go

FROM golang:1.18-alpine

WORKDIR /app

RUN mkdir -p /app/templates /app/static
COPY --from=builder /bin/godraft /app/
COPY static/ /app/static/
COPY templates /app/templates

EXPOSE 3000

CMD [ "/app/godraft" ]
