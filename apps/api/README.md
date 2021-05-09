# API

The REST API and Websocket connections used by the Overlay.

## Swagger API

You can generate a swagger spec for the API endpoints. The generated file will be output to `dist/apps/api/swagger.json`. This can then be imported into tools like Insomnia or Postman.

This uses go-swagger to generate the spec file via annotations in the code. You need to install go-swagger and add it to your path before being able to generate the spec file.

Download go-swagger [here](https://github.com/go-swagger/go-swagger/releases).

`npm run api:swagger`

You can also view the swagger spec via swagger-ui with the following command (also requires go-swagger):

`npm run api:swagger:serve`

## Last Done

- Adding /game GET to api.
- Be able to add a player to a game.
- Be able to add a team to a game.
- Remove players from game when adding a team and vice versa.

## TODO NEXT

- Unset a player.
- Create a player.
- Edit a player.
- Unset a team.
- Create a team.
- Add a player to a team.
