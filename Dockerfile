ARG GO_VERSION=1.15

FROM golang:${GO_VERSION}-alpine As builder

ENV PORT=3002

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api

WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./app ./src/main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .
COPY --from=builder /api/src/.env .

EXPOSE $PORT

ENTRYPOINT ["./app"]
