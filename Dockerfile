
# Build the Svelte frontend ------------------------------------
FROM node:22 AS ui
WORKDIR /app/frontend

# Install deps first (cached if package files don't change)
COPY frontend/package*.json ./
RUN npm install

# Copy the rest of the frontend and build it (outputs to /app/frontend/dist)
COPY frontend/ .
RUN npm run build


# Build the Go backend and embed the UI -------------------------
FROM golang:1.25 AS gobuild
WORKDIR /app/backend

# Copy Go module files first (better cache on go mod download)
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy backend source
COPY backend/ .

# Copy built frontend into backend/static so go:embed can pick it up
COPY --from=ui /app/frontend/dist ./static

# Build a static-ish Linux binary. CGO disabled = fewer surprises in the final image.
ENV CGO_ENABLED=0
RUN go build -o /server ./cmd/server


# --- Stage 3: tiny runtime image that just runs the binary -------------------
FROM alpine:3.20
WORKDIR /app

# Copy the compiled server from the previous stage
COPY --from=gobuild /server /server

# Cloud Run uses $PORT; default to 8080 for local runs
ENV PORT=8080
EXPOSE 8080

# Start the server
CMD ["/server"]

