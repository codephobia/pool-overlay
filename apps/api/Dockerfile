# Use nxgo/cli as the base image to do the build
FROM nxgo/cli as builder

# Create app directory
WORKDIR /workspace

# Copy package.json and the lock file
COPY package.json package-lock.json /workspace/

# Install app dependencies
RUN npm i

# Copy source files
COPY . .

# Download go modules
RUN go mod download

# Build apps
# RUN nx build api
RUN go build -o /workspace/dist/apps/api ./apps/api/src/

# This is the stage where the final production image is built
FROM golang:1.14-alpine as final

# Copy over artifacts from builder image
COPY --from=builder /workspace/dist/apps/api /workspace/api

# Set environment variables
ENV PORT=3000
ENV HOST=0.0.0.0

# Expose default port
EXPOSE 3000

# Start server
CMD [ "/workspace/api" ]