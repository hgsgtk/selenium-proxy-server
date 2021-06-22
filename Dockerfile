FROM golang:1.16-alpine as deploy-builder

WORKDIR /app

RUN apk add --no-cache git

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

COPY . .
ARG REVISION="default"

RUN go mod download

RUN go build -o server -ldflags "-X main.revision=${REVISION}"

FROM alpine as deploy

COPY --from=deploy-builder /app/server .

CMD ["./server"]
