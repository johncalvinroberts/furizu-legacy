# furizu

> A simple way to back up your data to cold storage, forever.


## Tech Stack

* The backend is created with Golang, using [the gin HTTP framework](https://github.com/gin-gonic/gin) and [gqlgen](https://gqlgen.com/)
* The frontend is created with [Svelte](https://svelte.dev/)
* This application has a preference for boring AWS services
  * Archives are stored on AWS Glacier
  * User info and data describing the archives is stored on Dynamo DB
  * AWS's SES is used for sending emails
  * AWS's SNS is used as a messaged queue

## Using This App

This app uses mostly `make` commands.

* `make install` to install deps for both frontend and backend
* `make dev` to run backend and frontend in parallel. Frontend will be on localhost:3000, backend on localhost:4000
* `make clean` cleans out build artifacts.
* `make build` to build the production binary. The output will be in `tmp/furizu`.
* `make gql` generates graphql scaffolding in the backend codebase.

For more, see the `Makefile`.

## Documentation

Sequence Diagrams, database schemas, and other diagrams are created using mermaid syntax. Use [mermaid.live](https://mermaid.live/) for editing.

Mermaid docs and any other documentation lives in [docs](./docs).

