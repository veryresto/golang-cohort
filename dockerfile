FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download \
    && go build -o ./app cmd/api/main.go

# Build from scratch
FROM scratch

COPY --from=builder /build/app /
ADD .env .

ENTRYPOINT ["/app"]