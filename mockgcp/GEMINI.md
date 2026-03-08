# Gemini Code Understanding

## Project Overview

This project, MockGCP, is a Go-based testing tool that provides mock implementations of Google Cloud Platform (GCP) APIs. Its primary purpose is to facilitate testing for the Config Connector project by simulating GCP services. This allows for consistent and reliable testing without the need for actual GCP resources, enabling fault injection and faster test execution.

The core of MockGCP is an HTTP round-tripper that intercepts requests and directs them to the appropriate mock service. These services are implemented using GCP's published protocol buffer (proto) definitions, and `grpc-gateway` is used to create HTTP servers from these gRPC definitions. This approach allows for in-process testing without needing to run a separate web server.

## Building and Running

### Building

The primary build process involves generating Go code from protocol buffer definitions. This is managed through the `Makefile`.

To build the project, run:

```bash
make all
```

This command will download necessary tools, sync the required version of the `googleapis` repository, and generate the Go code from the `.proto` files.

### Running Tests

Tests are standard `go test` tests, but some helper scripts are provided:

* `dev/tools/record-gcp [<testpath>]` will record the results of running the tests against (real) GCP; writing the golden output logs.
* `dev/tools/compare-mock [<testpath>]` will record the results of running the tests against our mocks of GCP; ideally the golden output will not change.

## Development Conventions

### Adding a New Mock Service

The `README.md` file provides a detailed guide for adding a new mock service. The general process is as follows:

1.  **Add the Proto Definition**: Locate the relevant `.proto` file for the GCP service in the `googleapis` repository and add it to the `Makefile`.
2.  **Generate Go Code**: Run `make gen-proto` to generate the necessary Go files from the proto definition.
3.  **Create a Mock Service Directory**: Create a new directory named `mock<servicename>`.
4.  **Implement the basic service**: In your service directory, there should be a service.go and a normalize.go file.  The service should also be registered in `/register.go`
4.  **Implement the Service**: Implement the core CRUD (Create, Read, Update, Delete) methods for the service in a `.go` file within the new directory. This typically involves creating a `<resource>>.go` file and implementing the proto methods, and ensuring that service is registered in `service.go` for both GRPC and HTTP.
5.  **Run and Debug**: Run the tests for the new service and use the provided debugging tools to ensure it behaves as expected.

### Golden File Testing

The project uses a "golden file" testing approach. This involves:

1.  Running tests against the real GCP services to capture the expected HTTP requests and responses.
2.  Saving this captured data as "golden files" (e.g., `_http.log`).
3.  Running the tests against the mock implementation and comparing the actual output to the golden files.

This ensures that the mock implementation accurately reflects the behavior of the real GCP APIs.
