# balance-sync

Simple gRPC client for balance sync service

Run with

`go run client/client.go`

You can run multiple clients to get updates between them.
After you run the second one - first one will be updated on changes.

If you run more - previous clients will be updated also,
feel free to run 2-3 clients at same time to test it.
