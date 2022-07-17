```
# go-spiget

go-spiget is a Go client library for accessing the [Spiget API v2](https://spiget.org/documentation).
## Installation

go-spiget is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/sunxyw/go-spiget
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/sunxyw/go-spiget"
```

and run `go get` without parameters.

## Usage/Examples

Construct a new Spiget client, then use the various services on the client to access different parts of the Spiget API.
For example:

```go
client := spiget.NewClient(nil)

// Get resource with ID 6245, PlaceholderAPI
resource, _, err := client.Resources.Get(context.Background(), 6245)
```

Some API methods have optional parameters that can be passed. For example:

```go
opt := &spiget.ResourceSearchOptions{
    Field: "name",
}
resources, _, err := client.Resources.Search(context.Background(), "PlaceholderAPI", opt)
```

The services of a client divide the API ito logical chunks and correspond to the structure of the Spiget API documentation at https://spiget.org/documentation .

NOTE: Using the [context](https://godoc.org/context) package, one can easily pass cancelation signals and deadlines to various services of the client for handling a request. In case there is no context available, then context.Background() can be used as a starting point.

For more sample code snippets, head over to the [example](https://github.com/sunxyw/go-spiget/tree/master/example) directory.

## License

[MIT](https://choosealicense.com/licenses/mit/)

```

```
