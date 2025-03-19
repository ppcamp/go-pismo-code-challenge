# go-pismo-code-challenge


## Tools

To create migrations on your machine, or use other extra tools, you can just
install them by typing:

```sh
make setup_dev
```

> See `make help` or type `make` to see available commands.

## How to run?

1. Type `make up`. this will startup the database in a docker environment.
2. Type `make migrate`, this will run the `cli app` that executes the migrations
stored in `migrations` folder.
3. Type `make run` to run app in development mode.

Alternativaly, you can run the production config: `make up_env`

