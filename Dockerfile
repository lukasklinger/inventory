# syntax=docker/dockerfile:1

FROM golang:1.22 as builder

# Set destination for COPY
WORKDIR /app

# Copy the source code
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./cmd/inventory/main.go

RUN ls

FROM scratch

EXPOSE 8080

COPY --from=builder /app/* .

CMD ["/app"]