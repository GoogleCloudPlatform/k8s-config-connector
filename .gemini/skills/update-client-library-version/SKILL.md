---
name: update-client-library-version
description: Guides through updating a Google Cloud Go client library to the latest version across the Config Connector repository.
---

# Upgrading Google Cloud Client Libraries in Config Connector

This skill outlines the process for updating Google Cloud client libraries (e.g., `cloud.google.com/go/alloydb` or `cloud.google.com/go/spanner`) in Config Connector (KCC).

## Workflow

### 1. Identify the Latest Version

Find the latest version of the target Go client library. You can run:
```bash
go list -m -versions cloud.google.com/go/<service>
```

### 2. Update the Root Module

Update the dependency in the root `go.mod` directory to the target version and tidy up.
```bash
go get cloud.google.com/go/<service>@<version>
go mod tidy
```

### 3. Update the MockGCP Module

Update the dependency in the `mockgcp/go.mod` directory to align version configuration and tidy up.
```bash
cd mockgcp
go get cloud.google.com/go/<service>@<version>
go mod tidy
cd ..
```

### 4. Verify Fuzzing & Identify Missing Fields

Run KRM mapping fuzz tests to ensure the updated types do not introduce serialization or validation regressions.
```bash
go test -v ./pkg/fuzztesting/fuzztests/
```

If any new fields are introduced in the updated library, the roundtrip verification might fail. The failure messages will output helpful recommendations such as:
`Add f.Unimplemented_NotYetTriaged(".new_field") to the fuzzer for the proto type <proto_type> to mark this field as not yet triaged.`

### 5. Mark New Fields as Untriaged

Locate the corresponding fuzzer file for the service (typically under `pkg/controller/direct/<service>/<kind>_fuzzer.go`).
Add the recommended field ignore using `f.Unimplemented_NotYetTriaged(".new_field")` in the initialization of the fuzzer.

### 6. Format and Validate

Run formatting and verify the entire package compiled without static analysis issues:
```bash
make fmt
go vet ./pkg/controller/direct/<service>/... ./pkg/fuzztesting/fuzztests/...
```

Run the fuzz tests again to ensure all tests pass:
```bash
go test -v ./pkg/fuzztesting/fuzztests/
```
