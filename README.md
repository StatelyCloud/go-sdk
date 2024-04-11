# StatelyDB SDK for Go

This is the Go client for the Stately Cloud API. We're still in an invite-only
preview mode - if you're interested, please reach out to preview@stately.cloud.

This client is currently meant for use from Go server applications.

The client library can be installed as such:

```
go get github.com/StatelyCloud/go-sdk
```

When you join the preview program, we'll set you up with a few bits of information:

1. `STATELY_CLIENT_ID` - a client identifier so we know what client you are.
2. `STATELY_CLIENT_SECRET` - a sensitive secret that lets your applications authenticate with the API.
3. A store ID that identifies which store in your organization you're using.
4. A link to more in-depth documentation than this README.

To use the client from a Go application:

```go
import (
  "github.com/StatelyCloud/go-sdk/data"
  "github.com/StatelyCloud/go-sdk/client"
)

ctx := context.Background() // TODO: Use a real context please

// Create a client. This will use the environment variables
// STATELY_CLIENT_ID and STATELY_CLIENT_SECRET for your client.
client, err := data.NewClient(ctx, 12345, nil)
```

Now, you can call the Data API:

```go
data := map[string][string]{
  "name": "Stiley",
}
client.Put(ctx, "/user-1", data);
```
