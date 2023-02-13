FROM node:18.14.0-alpine3.17

WORKDIR /workspace

# Build argument for fontawesome npm token
ARG FONTAWESOME_NPM_AUTH_TOKEN

# Note: tsconfig.base.json holds paths for libs, so may want this being loaded
# via a volume instead.
COPY jest.config.ts \
    jest.preset.js \
    nx.json \
    package-lock.json \
    package.json \
    tsconfig.base.json \
    tslint.json \
    workspace.json \
    .npmrc \
    /workspace/

RUN npm ci
