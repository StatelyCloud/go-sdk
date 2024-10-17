# StatelyDB SDK for Go
[![Go Reference](https://pkg.go.dev/badge/github.com/StatelyCloud/go-sdk.svg)](https://pkg.go.dev/github.com/StatelyCloud/go-sdk)

This is the Go SDK for [StatelyDB](https://stately.cloud).

### Getting started:

##### Disclaimer:

We're still in an invite-only preview mode - if you're interested, please reach out to [preview@stately.cloud](mailto:preview@stately.cloud?subject=Early%20Access%20Program).

When you join the preview program, we'll set you up with a few bits of information:

1. `STATELY_CLIENT_ID` - a client identifier so we know what client you are.
2. `STATELY_CLIENT_SECRET` - a sensitive secret that lets your applications authenticate with the API.
3. A store ID that identifies which store in your organization you're using.
4. Access to our in-depth [Getting Started Guide].

Begin by following our [Getting Started Guide] which will help you define, generate, and publish a DB schema so that it can be used.

##### Install the SDK

```sh
go get -u github.com/StatelyCloud/go-sdk
```

### Usage:

After [Defining Schema], be sure to generate a schema in a go module where it can be referenced. For the purposes of this readme ths will be `github.com/Project/Package/schema`.

###### Instantiate an Authenticated Client

To use the client from a Go application:

```go
import (
	"github.com/StatelyCloud/go-sdk/stately"

	"github.com/Project/Package/schema"
)

func main() {
	ctx := context.Background() // TODO: Use a real context please
	// Create a client. This will use the environment variables
	// STATELY_CLIENT_ID and STATELY_CLIENT_SECRET for your client.
	client, err := schema.NewClient(ctx, 12345)
	if err != nil { ... }

	// Alternatively:
	client, err := schema.NewClient(ctx, *stately.Options{
		ClientID: "myClientID",
		ClientSecret: "myClientSecret",
	})
	if err != nil { ... }
}
```

###### Using a client:

Once you have an authenticated client, you can use reference your item types (where they live in your schema module) and use the client!

```go
func PutMyItem(ctx context.Context, client stately.Client) error {
	item := &schema.Person{
		Handle: "i_am_jane",
		Name: "Jane Doe",
	}
	putResult, err := client.Put(ctx, item)
	if err != nil { ... }

	getResult, err := client.Get(ctx, "/user-i_am_jane")
    if err != nil { ... }

	// etc.
}
```

---

[Getting Started Guide]: https://docs.stately.cloud/guides/getting-started/
[Defining Schema]: https://docs.stately.cloud/guides/schema/