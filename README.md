# Pool Overlay

An OBS browser source overlay that will show which players are currently playing and what the score is.

![Screenshot of overlay](screenshots/screenshot1.png "Screenshot of overlay")

## Setup

1. Clone the repo.

    `git clone github.com/codephobia/pool-overlay.git`

2. Go to new directory.

    `cd pool-overlay`

3. Install dependencies.

    `npm install`

4. Launch application.

    `docker compose up`

## Apps

The following applications exist in the Nx monorepo:

| Name    | Description                                                 | Readme                           |
| ------- | ----------------------------------------------------------- | -------------------------------- |
| Api     | The REST API and Websocket connections used by the Overlay. | [README](apps/api/README.md)     |
| Seed    | Seeds the database with some basic information.             | [README](apps/seed/README.md)    |
| Overlay | An Angular application that displays the scoreboard in OBS. | [README](apps/overlay/README.md) |

## Libs

The following libraries exist in the Nx monorepo:

| Name | Description                                               | Readme                      |
| ---- | --------------------------------------------------------- | --------------------------- |
| go   | Contains all internal packages for the api and seed apps. | [README](libs/go/README.md) |

## Docker

Everything is setup to run with docker containers and docker compose. The following services exist:

| Name     | Description                                                    | Ports | Dockerfile                            |
| -------- | -------------------------------------------------------------- | ----- | ------------------------------------- |
| db       | The postgres database storing all player and game information. | 5432  |                                       |
| db-admin | An admin interface for the database using pgadmin.             | 5050  |                                       |
| api      | The REST API and Websocket connections used by the Overlay.    | 1268  | [Dockerfile](apps/api/Dockerfile)     |
| overlay  | An Angular application that displays the scoreboard in OBS.    | 4200  | [Dockerfile](apps/overlay/Dockerfile) |

## Github Actions

Github actions are setup on the repo to test linting.

Lints the following:

- [x] Go
- [ ] Angular
