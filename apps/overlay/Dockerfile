FROM node:14.15.0-alpine3.12

WORKDIR /workspace

COPY jest.config.js \
    jest.preset.js \
    nx.json \
    package-lock.json \
    package.json \
    tsconfig.base.json \
    tslint.json \
    workspace.json \
    /workspace/

RUN npm install
