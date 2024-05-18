FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o bin/app

FROM scratch

COPY --from=builder /app/bin/app /

EXPOSE 8000

CMD ["/app"]