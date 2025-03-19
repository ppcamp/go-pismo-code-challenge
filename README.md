# go-pismo-code-challenge


## Tools

To create migrations on your machine, or use other extra tools, you can just
install them by typing:

```sh
make setup_dev
```

## How to run?

1. Type `docker compose up`. This will:
    - Build service image
    - Startup the database
    - Startup the service

To run in development mode, just type `make run`, see `make help` or type `make`
to see available commands.
