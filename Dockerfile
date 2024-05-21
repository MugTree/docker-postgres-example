FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o bin/app

FROM scratch

COPY --from=builder /app/bin/app /
COPY --from=builder /app/html /html

EXPOSE 8000

CMD ["/app"]