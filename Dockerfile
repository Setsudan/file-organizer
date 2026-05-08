FROM golang:1.23.4-alpine AS builder

WORKDIR /src
COPY go.mod ./
RUN go mod download

COPY . .
ARG TARGETOS=linux
ARG TARGETARCH=amd64
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/file-organizer ./cmd/file-organizer

FROM alpine:3.20
RUN adduser -D -u 10001 appuser
USER appuser
WORKDIR /app
COPY --from=builder /out/file-organizer /usr/local/bin/file-organizer

ENTRYPOINT ["file-organizer"]
CMD ["--mode=once"]
