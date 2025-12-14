# build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app

COPY go.mod go.sum ./ 
RUN go mod download

COPY . .
# RUN go get -d -v ./...
RUN mkdir -p /go/bin/app
RUN go build -v -o /go/bin/app ./...

# deploy stage
FROM alpine:latest
#RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app/breadzilla"]