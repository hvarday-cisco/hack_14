# Hack_14

### Scripts

There are 2 scripts -

- **Register webhook** - call Google API to register our webhook endpoint for a specific org
- **Deregister webhook** - call Google API to cancel an existing webhook subscription.

#### How to run?

```shell script
go run scripts/register_webhook.go
go run scripts/deregister_webhook.go
```

### Webhook Server

We run the server for incoming webhook requests with `server.go` file. It processes the information
and sends a 200 - Status OK response back.

#### How to run?
```shell script
go run server.go
```
