FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git bash
RUN go get -d -v ./...
#RUN go get -d -v

RUN pwd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main github.com/huynhdunguyen/boilerplate-golang-samplelang/cmd/api

#FROM scratch
FROM alpine:latest
COPY --from=builder /build/main /app/
WORKDIR /app
COPY . .
EXPOSE 8080
CMD ["./main"]
