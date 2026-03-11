FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache git make bash

COPY . .

RUN go mod download

RUN make bin-deps

EXPOSE 8080

CMD ["make", "run"]