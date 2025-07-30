# Builder
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app ./cmd/app/

# Runner
FROM gcr.io/distroless/static

COPY --from=builder /app/app /

USER nonroot

EXPOSE 8080

ENTRYPOINT ["/app"]
