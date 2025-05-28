FROM golang:1.24.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -o main ./cmd/server

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
