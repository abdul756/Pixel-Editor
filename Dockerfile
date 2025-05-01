FROM golang:1.22-alpine AS builder

# Install required system dependencies for Fyne
RUN apk add --no-cache gcc g++ musl-dev xorg-server-dev \
    libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev \
    libxi-dev mesa-dev

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o pixeleditor ./src/pixel.go

# Use a lightweight Alpine image for the final image
FROM alpine:latest

# Install required runtime dependencies for Fyne applications
RUN apk add --no-cache xorg-server mesa-gl mesa-dri-gallium \
    libx11 libxcursor libxrandr libxinerama libxi \
    dbus-x11 ttf-dejavu

# Set environment variables for X11
ENV DISPLAY=:0
ENV XAUTHORITY=/tmp/.Xauthority

WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/pixeleditor .

# Set the entry point
ENTRYPOINT ["./pixeleditor"] 