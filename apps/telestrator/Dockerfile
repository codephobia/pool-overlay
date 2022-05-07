FROM node:16.15.0-alpine3.14

WORKDIR /workspace

# Note: tsconfig.base.json holds paths for libs, so may want this being loaded
# via a volume instead.
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
