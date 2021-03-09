FROM golang:1.11 as builder
WORKDIR /go/src/payment-service
COPY . .
RUN go test --cover ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o paymentsvc ./cmd/paymentsvc/main.go

FROM alpine:3.6
WORKDIR /
RUN apk update && apk add ca-certificates
COPY --from=builder /go/src/payment-service/paymentsvc .
COPY migrations /go/src/payment-service/migrations
EXPOSE 80
CMD ["/paymentsvc"]
