FROM golang:1.14.7-alpine3.12 as builder

RUN mkdir /app
WORKDIR /app
RUN apk add --update --no-cache ca-certificates git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /go/bin/test cmd/nginx_test/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/test /go/bin/test
ENTRYPOINT ["/go/bin/test"]