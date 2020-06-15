FROM golang:1.14-alpine as builder
COPY . /src
WORKDIR /src

RUN adduser --disabled-password --uid 10001 app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o /app ./cmd


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

USER app
COPY --from=builder /app /app

EXPOSE 8080
CMD ["/app", "run"]
