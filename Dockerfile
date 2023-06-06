FROM golang:1.19-alpine as builder

WORKDIR /project

ENV DB_HOST=host.docker.internal \
    DB_PORT=3303 \
    DB_USER=root \
    DB_NAME=db-demo \
    DB_PASS=root

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./cmd/main.go

FROM alpine:3.17

ENV DB_HOST=host.docker.internal \
    DB_PORT=3303 \
    DB_USER=root \
    DB_NAME=db-demo \
    DB_PASS=root
COPY --from=builder ./project/bin/server /server
COPY --from=builder ./project/migration /migration

CMD ["/server"]