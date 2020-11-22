# Use nxgo/cli as the base image to do the build
FROM nxgo/cli as builder

# Create app directory
WORKDIR /app

# Copy package.json and the lock file
COPY package.json package-lock.json /app/

# Install app dependencies
RUN npm i

# Copy source files
COPY . .

# Build apps
RUN nx build api

# This is the stage where the final production image is built
FROM golang:1.14-alpine as final

# Copy over artifacts from builder image
COPY --from=builder /app/dist/apps/api /app/api

# Set environment variables
ENV PORT=3000
ENV HOST=0.0.0.0

# Expose default port
EXPOSE 3000

# Start server
CMD [ "/app/api" ]