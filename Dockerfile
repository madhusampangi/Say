FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app

COPY ./saytext/server .
COPY ./go.mod .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sayserver


FROM alpine:latest

RUN apk update && apk add flite

RUN mkdir /app
WORKDIR /app
EXPOSE 8080

COPY --from=builder /app/sayserver .

CMD ["./sayserver"]