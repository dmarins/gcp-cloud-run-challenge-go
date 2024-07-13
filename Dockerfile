FROM golang:1.22 as builder
WORKDIR /app
RUN apt-get install -y ca-certificates
COPY . .
WORKDIR /app/cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun main.go wire_gen.go && mkdir ../../dist && mv cloudrun .env ../../dist

FROM scratch as runner
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/dist/.env .
COPY --from=builder /app/dist/cloudrun .
ENTRYPOINT ["./cloudrun"]