FROM golang:1.25.7-alpine AS builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN go build ./cmd/api/

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/api /app/
WORKDIR /app
CMD ["./api"]
