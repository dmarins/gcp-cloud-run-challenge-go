FROM golang:1.22 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=LINUX GOARCH=amd64 go build -o cloudrun

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
ENTRYPOINT ["./cloudrun"]