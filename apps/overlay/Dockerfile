FROM node:16.15.0-alpine3.14

WORKDIR /workspace

COPY jest.config.ts \
    jest.preset.ts \
    nx.json \
    package-lock.json \
    package.json \
    tsconfig.base.json \
    tslint.json \
    workspace.json \
    /workspace/

RUN npm install
