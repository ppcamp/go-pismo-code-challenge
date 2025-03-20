# go-pismo-code-challenge


## TODO

- implement test for transactions
- implement httptest layer


## Development

Type `make lint` or `make generate` (required to generate mocks for tests)

## Tools

To create migrations on your machine, or use other extra tools, you can just
install them by typing:

```sh
make setup_dev
```

> See `make help` or type `make` to see available commands.

## How to run?

1. Type `make up`. this will startup the database and the swagger service
in a docker environment.
2. Type `make migrate`, this will run the `cli app` that executes the migrations
stored in `migrations` folder.
3. Type `make run` to run app in development mode.

Alternativaly, you can run the production config: `make up_env`

> You can check the API doc at `./build` folder or, after setup the local env,
> at http://localhost:9090. See also [swagger-editor]

The server also contains an otel stack configured, that can be accessed at
http://localhost:3000

[swagger-editor]: https://editor.swagger.io/
