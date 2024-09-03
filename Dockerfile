FROM golang:1.22.3-alpine AS builder

WORKDIR /go/src/github.com/kenoboya/backend_university_system

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o ./app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/src/github.com/kenoboya/backend_university_system/app .
COPY --from=builder /go/src/github.com/kenoboya/backend_university_system/configs /configs
COPY .env .

CMD ["./app", "-config", "/configs/server.yml"]
