# syntax=docker/dockerfile:1

FROM golang:1.22 AS builder

# Set destination for COPY
WORKDIR /app

# Copy the source code
COPY . ./

# Build
RUN GOOS=linux go build -o /app/app ./cmd/inventory/main.go

FROM gcr.io/distroless/base-debian12

EXPOSE 8080

WORKDIR /app

COPY . ./
COPY --from=builder /app/app .

CMD ["/app/app"]