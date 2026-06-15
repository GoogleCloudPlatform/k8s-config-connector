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

Implement it by parsing the URL to the resource's Identity struct and returning a populated Adapter.

1. Check if the URL matches the expected prefix (e.g., host / protocol) and trim it.
2. Initialize an instance of the resource's Identity struct and parse the trimmed URL using `FromExternal(url)`. If parsing fails, return `nil, nil` (not recognized).
3. Construct the GCP client using the model's client factory.
4. Return a populated Adapter with the parsed identity and the GCP client.

**Example Implementation:**
```go
func (m *modelArtifactRegistryRepository) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	if !strings.HasPrefix(url, "//artifactregistry.googleapis.com/") {
		return nil, nil
	}

	url = strings.TrimPrefix(url, "//artifactregistry.googleapis.com/")

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
2. Use the CAIS package (`github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cais`) to dynamically resolve the export URI based on the resource's CAIS identity.
3. If CAIS successfully resolves the identity URL (i.e., `CAISURL != "unknown"`), assign it to `exportURI`.

**Example integration:**
```go
	case schema.GroupKind{Group: "artifactregistry.cnrm.cloud.google.com", Kind: "ArtifactRegistryRepository"}:
		caisScheme := cais.NewScheme()
		caisResults, err := cais.GetCAISIdentities(h.Ctx, caisScheme, h.GetClient(), []*unstructured.Unstructured{obj})
		if err == nil && len(caisResults) > 0 && caisResults[0].CAISURL != "unknown" {
			exportURI = caisResults[0].CAISURL
		} else {
			h.T.Errorf("failed to get CAIS identity for ArtifactRegistryRepository: %v", err)
		}
```

### Step 3: Enable Export Diff Verification in `unified_test.go`

In `tests/e2e/unified_test.go`, ensure that if the test is running with `forceDirect` (or `ForceDirectController = true`), `createDiffs` validates that the direct-exported and fallback-exported YAML files match and outputs a diff file `_exported_object.diff`.

Update `createDiffs` with:
```go
	// _exported_object.diff
	{
		oldPath := filepath.Join(dir, "_exported_old_controller.golden.yaml")
		newPath := filepath.Join(dir, "_generated_export_"+fixture.Name+".golden")

		if fileExists(oldPath) && fileExists(newPath) {
			diff := computeDiff(oldPath, newPath)
			h.CompareGoldenFile(filepath.Join(dir, "_exported_object.diff"), diff)
		} else {
			h.AssertGoldenFileNotFound(filepath.Join(dir, "_exported_object.diff"))
		}
	}
```

### Step 4: Run E2E Test and Generate Golden Export Files

Run the target test fixture using:
```bash
WRITE_GOLDEN_OUTPUT=1 go test -v ./tests/e2e -run "TestAllInSeries/.*<test_name>.*"
```
Ensure that `_generated_export_<test_name>.golden` and `_exported_object.diff` files are created, verified, and correctly captured.
