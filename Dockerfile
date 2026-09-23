FROM golang:1.27-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
COPY internal/api ./internal/api
COPY internal/db ./internal/db
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o football-teams-creator .

FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache sqlite-libs

COPY --from=builder /app/football-teams-creator .
COPY --from=builder /app/web ./web

EXPOSE 8080

CMD ["./football-teams-creator"]
