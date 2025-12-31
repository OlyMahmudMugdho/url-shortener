FROM golang:1.22.5-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o url-shortener main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/url-shortener .
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/db ./db
EXPOSE 8080
CMD ["./url-shortener"]