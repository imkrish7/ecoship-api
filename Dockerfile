FROM golang:1.22.3 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/web

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api .

EXPOSE 4567

CMD ["./api"]