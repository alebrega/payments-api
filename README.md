# Payments API

## Running the API

### Prerequisites

Prior running the API, make sure you have GO and Docker installed in your machine.

To able to perform integration tests, make sure to install `newman`

### Terminal

1. `config.toml` is used to provide container-like environment variables. The application will load your custom settings when you export the variable `ENVIRONMENT=dev` on your terminal:

```
export ENVIRONMENT=dev
```

2. The dependencies are managed using `dep`. Make sure to load all necessary to your `$GOPATH`:

```
make install
```

3. The application uses a `Postgres` database, make sure to have a running db instance (e.g. `docker-compose up db`) with the right connection parameters updated in `config.toml` before running t:

```
make run
```

### Docker

Please note the the `Dockerfile` uses the binary built locally. `Docker-compose` defines services for the Payments API, Postgres and `newman` testing:

```
make docker-compose-build # builds the payments-api and the custom newman image
make docker-compose-up
```

You can interact with the API on port `:8080`. The observabilty metrics and the heath checks are exposed on port `:8081`.
`newman` generates a html report in the reports folder.

## Testing

Unit tests and code coverage could be performaed using these commands:

```
make unit-test
make test-cover
```

> Please note that any changes to interfaces need to regenerate new mocks using the command `make update-mocks`.

Run the integration tests using:

```
make integration-tests
```
