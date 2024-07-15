FROM golang:1.22-alpine

ARG NAME
ARG PORT
ARG PROXY_PATH

ENV APP_ENV=development
ENV APP_NAME=$NAME
ENV APP_PORT=$PORT
ENV APP_PROXY_PATH=$PROXY_PATH

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app/${APP_NAME}

COPY . .

RUN go install github.com/cespare/reflex@latest
RUN go install github.com/daarlabs/hrx@v0.1.4

ENTRYPOINT reflex -r '(\.go$|go\.mod|\.yaml$)' -s go run ./app/${APP_NAME}/ & hrx dev