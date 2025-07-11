      
# =======================================================
#               STAGE 1: Build API Binary
# =======================================================
FROM golang:1.21-alpine AS api-builder

WORKDIR /app

# Salin hanya apa yang diperlukan untuk dependensi
COPY go.mod go.sum ./
RUN go mod download

# Salin semua kode
COPY . .

# Build API binary
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/api ./cmd/api/main.go


# =======================================================
#               STAGE 2: Build Worker Binary
# =======================================================
FROM golang:1.21-alpine AS worker-builder

WORKDIR /app

# Salin hanya apa yang diperlukan untuk dependensi
COPY go.mod go.sum ./
RUN go mod download

# Salin semua kode
COPY . .

# Build Worker binary
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/worker ./cmd/worker/main.go


# =======================================================
#               FINAL STAGE: Create Final Image
# =======================================================
FROM alpine:latest

WORKDIR /app

# Salin biner dari stage build yang relevan
COPY --from=api-builder /app/api .
COPY --from=worker-builder /app/worker .

# Verifikasi bahwa file benar-benar ada di dalam image final
RUN ls -la /app

# Expose port untuk API
EXPOSE 8080

# Command akan ditentukan di docker-compose.yml

    