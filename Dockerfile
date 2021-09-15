FROM golang:1.17-alpine AS builder

COPY . /app
WORKDIR /app
RUN set -x && \
    go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]
CMD "serve"
CMD ["serve"]
