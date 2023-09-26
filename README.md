# GopherCon 2023 Buf CSR Demo

## Requirements

- Go toolchain 1.21+
- Git
- Docker
- [Buf CLI](https://buf.build/docs/installation)

## Setup steps

### Getting started

1. Clone this repo

   ```
   git clone https://github.com/bufbuild/gophercon-2023-workshop.git
   cd gophercon-2023-workshop
   ```

1. Start downloading Docker images (~860MB total)

   ```
   docker-compose -f ./docker/docker-compose.yml pull
   ```

1. [Login to the BSR](https://buf-gophercon.buf.dev/) using demo credentials (provided by presenter).

1. [Create a user token](https://buf-gophercon.buf.dev/settings/user) on the BSR to log into the Buf CLI. *Note*: Save this token somewhere you can easily retrieve it, we'll be using it in another step later.

   ```
   buf registry login buf-gophercon.buf.dev --username gophercon-csr-demo
   ```

1. Set `GOPRIVATE` environment variable to skip the Go module proxy for remote package code generation.

   ```
   export GOPRIVATE="${GOPRIVATE},buf-gophercon.buf.dev/gen/go"
   ```

### Creating the CSR instance

1. Navigate to the admin panel to [create a CSR instance](https://buf-gophercon.buf.dev/admin/csr). Name it something unique to you.

1. Update `CSRInstanceName` in [`config.go`](config.go) with the instance name you just created.

1. Update `BSRToken` in [`config.go`](config.go) with the token you created earlier (found in `$HOME/.netrc`).

### Creating the BSR Repository and CSR Subject/Schema

1. Update [`proto/emails/v1/events.proto`](proto/emails/v1/events.proto) with the CSR instance name you created previously.

1. Update [`proto/buf.yaml`](proto/buf.yaml) with a unique `name` for the repository.

1. Generate the app code:

   ```
   buf generate
   ```

1. Create and push the repository up to the BSR (and associated CSR instance):

   ```
   buf push --create --create-visibility public proto
   ```

### Run the demo app

1. Boot up the Kafka broker and AKHQ server:

   ```
   docker-compose -f ./docker/docker-compose.yml up -d
   ```

1. Run the demo app:

   ```
   go run .
   ```

1. In another terminal window, send a request to update an email address:

   ```
   curl -XPOST -H"content-type:application/json" \
       -d '{ "user_id": 123, "new_address": "gophercon@buf.build" }' \
       localhost:8888/emails.v1.EmailService/UpdateEmail
   ```

1. The app should log both the publish of the event on the producer side and 
   shortly after the verification of the email address. You can see the email 
   is marked as verified:

   ```
   curl -XPOST -H"content-type:application/json" \
       -d '{ "user_id": 123 }' \
       localhost:8888/emails.v1.EmailService/GetEmail
   ```
