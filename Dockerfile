FROM golang:1.18-alpine AS builder

ARG GOPROXY=https://proxy.golang.org,direct

ENV GO111MODULE=on \
  GOPROXY=${GOPROXY}

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

FROM alpine:3.7
WORKDIR /app
COPY --from=builder /app/public ./public/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
