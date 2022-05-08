FROM node:16.15.0-alpine3.14

WORKDIR /workspace

# Build argument for fontawesome npm token
ARG FONTAWESOME_NPM_AUTH_TOKEN

COPY jest.config.ts \
    jest.preset.ts \
    nx.json \
    package-lock.json \
    package.json \
    tsconfig.base.json \
    tslint.json \
    workspace.json \
    .npmrc \
    /workspace/

RUN npm install
