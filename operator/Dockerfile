# -------- Builder stage --------
FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 👇 The magic that makes it work on ARM64 + distroless
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64

RUN go build -o beanlord-operator ./operator/main.go

# -------- Final image --------
FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/beanlord-operator .

USER nonroot:nonroot

ENTRYPOINT ["/beanlord-operator"]