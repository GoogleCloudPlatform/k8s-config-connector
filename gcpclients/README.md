Sometimes the google proto client has not yet been generated.

In these cases, we can generate our own GRPC types.

It must be its own library because these GRPC types must be shared with mockgcp and KCC itself.

We generate a go.mod, and use the expected paths for when the libraries are generated.