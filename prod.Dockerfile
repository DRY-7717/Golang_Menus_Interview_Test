FROM golang:1.25.3-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o go_menus .

FROM alpine:3.22


RUN apk add --no-cache ca-certificates tzdata curl


ENV TZ=Asia/Jakarta

RUN adduser -D appuser

WORKDIR /app

COPY --from=builder /app/go_menus .
COPY --from=builder /app/docs ./docs

RUN chmod +x go_menus

USER appuser

EXPOSE 8000

HEALTHCHECK --interval=30s --timeout=5s --start-period=30s --retries=3 \
    CMD curl -f http://localhost:8000/api/check || exit 1



CMD ["./go_menus"]