# Use golang:1.22-bookworm as the base image
FROM golang:1.22-bookworm AS builder
WORKDIR /app

# Copy the source code into the container
COPY .. .

# Build the project
RUN make buildprod

# Stage 2: Use a lightweight image to serve the content
FROM nginx:alpine

# Copy the built contents from the previous stage
COPY --from=builder /app/build /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start nginx to serve content
CMD ["nginx", "-g", "daemon off;"]