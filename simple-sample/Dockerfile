FROM golang:1.14.4-buster AS builder

WORKDIR /go/src/github.com/cive/simple-sample-neo4j-visnetwork
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app app.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/cive/simple-sample-neo4j-visnetwork .
CMD ["./app"]
