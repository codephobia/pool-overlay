# Use nxgo/cli as the base image to do the build
FROM nxgo/cli as builder

# Create app directory
WORKDIR /workspace

# Build argument for fontawesome npm token
ARG FONTAWESOME_NPM_AUTH_TOKEN

# Copy package.json and the lock file
COPY package.json package-lock.json .npmrc ./

# Install app dependencies
RUN npm ci

# Copy go mod files
COPY go.mod go.sum ./

# Download go modules
RUN go mod download

# Copy go api source files
COPY apps/api apps/api

# Copy go seed source files
COPY apps/seed apps/seed

# Copy go lib source files
COPY libs/go libs/go

# Copy Nx files
COPY nx.json workspace.json tsconfig.base.json ./

# Build api app
# RUN nx build api
RUN go build -o dist/apps/api/api ./apps/api

# Build seed app
# RUN nx build seed
RUN go build -o dist/apps/seed/seed ./apps/seed

# This is the stage where the final production image is built
FROM golang:1.20-alpine as final

# Set environment variables
ENV PORT=1268
ENV HOST=0.0.0.0

# Expose default port
EXPOSE 1268

# Create app directory
WORKDIR /workspace

# Copy over api from builder image
COPY --from=builder /workspace/dist/apps/api/api ./api

# Copy over seed from builder image
COPY --from=builder /workspace/dist/apps/seed/seed ./seed

# Copy seed data
COPY apps/seed/data ./data

# Copy entrypoint bash script
COPY apps/api/entrypoint.sh ./

# Start server
ENTRYPOINT [ "sh", "entrypoint.sh" ]
