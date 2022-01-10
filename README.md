# furizu

> A simple way to back up your data to cold storage, forever.


## Tech Stack

* The backend is created with Golang, using [the gin HTTP framework](https://github.com/gin-gonic/gin)
* The frontend is created with [Svelte](https://svelte.dev/)
* This application has a preference for boring AWS services
  * Archives are stored on AWS Glacier
  * User info and data describing the archives is stored on Dynamo DB
    * [guregu/dynamo](https://github.com/guregu/dynamo) is used for a GO dynamo sdk
  * AWS's SES is used for sending emails

## Using This App

This app uses mostly `make` commands.

* `make install` to install deps for both frontend and beckend
* Run `make dev` to run backend and frontend in parallel. Frontend will be on localhost:3000, backend on localhost:4000
* Run `make clean` and `make build` to build the production binary. The output will be in `tmp/furizu`

For more, see the `Makefile`.

## Documentation

* [API Documentation](./docs/api.md)
* [Database Schema](./docs/mermaid/dynamo.mmd)

