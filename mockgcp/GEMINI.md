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

The project uses a "golden file" testing approach. This involves running tests against real GCP to capture ground truth, and then ensuring the mock implementation produces identical results.

**Expert Workflow: Aligning Mock with Real GCP**

Whenever you add a new field or modify GCP communication, you MUST align the mock logs.

1.  **Establish a Baseline**: 
    - Identify the `<testname>` (final folder name) and `<fixture-path>` (full relative path).
    - Run `hack/record-gcp "fixtures/^<testname>$"` to capture real GCP behavior (writes `_http.log` and `_generated_object_*.golden.yaml`).
    - **CRITICAL**: `git add <fixture-path>` and `git commit -m "Update real GCP golden logs"` to establish a clean git baseline.

2.  **Identify Discrepancies**:
    - Run `hack/compare-mock "fixtures/^<testname>$"` to record current mock behavior in the working tree.
    - Run `git diff <fixture-path>` to see the exact discrepancies between Mock and Real GCP.

3.  **Diagnose and Fix (Incremental)**:
    - **Output-Only Fields/IDs**: If IDs or URLs are wrong, search for `populate<Resource>Defaults` in `mockgcp/mock<service>/<resource>.go`. Extract the exact pattern from real logs (look for prefixes, trimming, or hashing) and implement it precisely.
    - **Volatile Fields**: If data like timestamps or etags differ but are functionally equivalent, update `mockgcp/mock<service>/normalize.go`.
    - **Controller Mappings**: If field names differ from the proto, check the controller mapping:
        - TF-Based: Search `third_party/github.com/hashicorp/terraform-provider-google` for `flatten<FieldName>` functions.
        - Direct: Check `pkg/controller/direct/<service>/mapper.generated.go`.
    - **Rule**: Fix only **ONE small point** at a time and verify immediately. Revert (`git reset --hard`) if you get stuck in a loop.

4.  **Verify**: Re-run `hack/compare-mock "fixtures/^<testname>$"` and check `git diff <fixture-path>`. Repeat until aligned.

### Normalization and Scoping

Mock services often need to normalize responses to ensure stable golden logs (e.g., replacing timestamps or IDs with placeholders). This is typically implemented in `normalize.go` using the `Previsit` and `ConfigureVisitor` methods.

**Purpose of Normalization**
Normalization is STRICTLY for removing randomness and non-reproducible values (like timestamps, UUIDs, generated IP addresses, etc.) to ensure tests are stable. It MUST NOT be used to paper over behavioral differences between mockgcp and real GCP. For example, if a real GCP API returns a project number in a specific field but the mock returns a project ID, you should fix the mock implementation to correctly return the project number rather than writing a normalizer to hide the discrepancy. Our goal is to make mockgcp accurately reflect real GCP behavior so controllers are forced to handle real-world API responses properly.

**Critical Rule: Previsit Scoping**

The `Previsit` method is **globally executed** for every event across all registered services. To prevent one service's normalization rules from affecting another service's tests, you MUST explicitly scope the `Previsit` logic to the relevant service URL.

Example of correct scoping:

```go
func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
    // Only apply normalization if the request is for this service
    if !strings.Contains(event.URL(), "myservice.googleapis.com") {
        return
    }
    // ... normalization logic ...
}
```

Failure to scope `Previsit` can lead to unintended log changes in unrelated services, causing "log corruption" and massive diffs in unrelated test data.
