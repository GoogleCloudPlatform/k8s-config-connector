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

### Step 2: Implement the `Export` method on the Adapter

In the Adapter implementation (e.g. `ArtifactRegistryRepositoryAdapter`), implement the `Export(ctx context.Context)` method. This method is responsible for translating the retrieved GCP state (`a.actual`) back to KRM format.

During export, you must:
1. **Set the project ID annotation / spec.projectRef:**
   - If the resource uses reference-based project binding (i.e. has a `spec.projectRef` field), set `obj.Spec.ProjectRef.External = a.id.Project`. **CRITICAL**: Do NOT call `export.SetProjectID(u, ...)` if you are setting `spec.projectRef` on the resource, as reference-bound resources should not have the `cnrm.cloud.google.com/project-id` annotation.
   - If the resource does not support `spec.projectRef`, use the helper `export.SetProjectID(u, projectID)` from `github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export` to set the project-id annotation.
2. **Set the labels:** Use the helper `export.SetLabels(u, labels)` from `github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export`.
3. **Set the identity fields on spec:** In general, fields that are part of the identity (like `location`, `region`, or `resourceID`) are not mapped automatically in the spec by the from-proto mappers, so you must set them manually on the struct before converting to Unstructured.
4. **Use the short name:** Make sure you use the short name of the resource (e.g. `a.id.Repository`) when calling `u.SetName()`, instead of the full GCP name.

**Example Implementation (for a reference-bound resource):**
```go
import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	...
)

func (a *ArtifactRegistryRepositoryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ArtifactRegistryRepository{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ArtifactRegistryRepositorySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Identity fields not mapped from proto must be set manually on the Spec struct.
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Repository)
	obj.Spec.ProjectRef.External = a.id.Project // Using spec.projectRef, so SetProjectID is omitted.

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Repository)
	u.SetGroupVersionKind(krm.ArtifactRegistryRepositoryGVK)

	// Set standard metadata such as labels (SetProjectID is NOT called because of projectRef).
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}
```

### Step 3: Register/Integrate with E2E Export Test Harness

To test your exporter, integrate it with the E2E test harness:

1. Update `tests/e2e/export.go`. In the `exportResource` function, add a case to the `switch gvk.GroupKind()` block for your GVK.
2. Use the `resolveCAISURI` helper to resolve the export URI based on the resource's CAIS identity.

**Example integration:**
```go
	case schema.GroupKind{Group: "artifactregistry.cnrm.cloud.google.com", Kind: "ArtifactRegistryRepository"}:
		exportURI = resolveCAISURI(h, obj)
```

### Step 4: Run E2E Test and Generate Golden Export Files

Run the target test fixture using:
```bash
WRITE_GOLDEN_OUTPUT=1 go test -v ./tests/e2e -run "TestAllInSeries/.*<test_name>.*"
```
Ensure that `_generated_export_<test_name>.golden` files are created, verified, and correctly captured.

### Alignment & avoiding `_exported_object.diff`

The E2E test harness compares the direct exporter's output (`_exported.yaml`) with the old/legacy exporter's output (`_exported_old_controller.golden.yaml`). If there are differences, an `_exported_object.diff` file will be generated.

We aim to **avoid generating any `_exported_object.diff` files** or at least minimize them to the absolute minimum of unresolvable differences. To achieve perfect alignment and prevent the creation of `_exported_object.diff` files:
1. **Always use the export helpers (when applicable):** Call `export.SetLabels(u, actual.Labels)`. Also call `export.SetProjectID(u, id.Project)` **ONLY** if the resource does not support `spec.projectRef` (i.e. if it is NOT a reference-bound resource). If the resource has `spec.projectRef`, do **NOT** call `export.SetProjectID(u, ...)`.
2. **Expose immutable fields:** Set the `ResourceID` and other identity fields on the exported Spec so they match the configuration exactly.
3. If an `_exported_object.diff` is generated during your tests, inspect the diff and modify your direct `Export()` mapping logic to match the output format (such as default values or standard formats) of the legacy controller.
