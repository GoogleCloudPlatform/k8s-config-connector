---
name: add-export-support
description: Guides the implementation of export support for a direct controller, including implementing AdapterForURL and verifying it through the unified test runner.
---

# Add Export Support

## Overview

This skill guides you through implementing export support for a "direct" controller in Config Connector. Direct controllers reconcile resources using `controller-runtime` and direct SDK calls. To support the `config-connector export` command, each direct controller must implement `AdapterForURL` in its `Model` implementation.

## Workflow

### Step 1: Implement `AdapterForURL` in the Direct Controller

In the direct controller file (e.g. `pkg/controller/direct/<service>/<kind>_controller.go`), locate the `AdapterForURL` method on the model struct. Currently, it typically has:
```go
func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support AdapterForURL
	return nil, nil
}
```

Implement it by parsing the URL directly to the resource's Identity struct and returning a populated Adapter.

`gcpurls.Template` / `FromExternal` natively supports parsing full CAI URLs starting with `//{host}/` (e.g. `//artifactregistry.googleapis.com/...`), so you can pass the URL directly to `FromExternal(url)` without manual trimming.

**Example Implementation:**
```go
func (m *modelArtifactRegistryRepository) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ArtifactRegistryRepositoryIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ArtifactRegistryRepositoryAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
```

### Step 2: Register/Integrate with E2E Export Test Harness

To test your exporter, integrate it with the E2E test harness:

1. Update `tests/e2e/export.go`. In the `exportResource` function, add a case to the `switch gvk.GroupKind()` block for your GVK.
2. Use the `resolveCAISURI` helper to resolve the export URI based on the resource's CAIS identity.

**Example integration:**
```go
	case schema.GroupKind{Group: "artifactregistry.cnrm.cloud.google.com", Kind: "ArtifactRegistryRepository"}:
		exportURI = resolveCAISURI(h, obj)
```

### Step 3: Run E2E Test and Generate Golden Export Files

Run the target test fixture using:
```bash
WRITE_GOLDEN_OUTPUT=1 go test -v ./tests/e2e -run "TestAllInSeries/.*<test_name>.*"
```
Ensure that `_generated_export_<test_name>.golden` files are created, verified, and correctly captured.
