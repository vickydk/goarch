############################
# STEP 1 build executable binary
############################
FROM golang:1.18-alpine AS builder

RUN apk --update --no-cache add \
    openssl \
    git \
    curl \
    tzdata \
    ca-certificates \
    && update-ca-certificates

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -v -o goarch ./cmd/app/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /app/resources /resources
COPY --from=builder /app/goarch /app/goarch

CMD ["/app/goarch"]
