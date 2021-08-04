FROM node:14.15.0-alpine3.12

WORKDIR /workspace

# Note: tsconfig.base.json holds paths for libs, so may want this being loaded
# via a volume instead.
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
