FROM golang:1.21.3-alpine

RUN apk add --no-cache make

RUN mkdir /build
WORKDIR /build

COPY ../../go.* .
RUN go mod download

COPY ../../cmd/. cmd/
COPY ../../Makefile .
COPY ../../internal/. internal/

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make swag-generate-api

CMD ["make", "run-api"]