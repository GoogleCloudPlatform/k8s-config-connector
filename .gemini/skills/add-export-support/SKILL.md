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
1. **Set the project ID annotation (if applicable):** If the resource spec does not have a `projectRef` field, use the helper `export.SetProjectID(u, projectID)` from `github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export` to set the project-id annotation. If `spec.projectRef` is explicitly set on the spec, do not set the `cnrm.cloud.google.com/project-id` annotation since it is redundant.
2. **Set the labels:** Use the helper `export.SetLabels(u, labels)` from `github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export`.
3. **Set the identity fields on spec:** In general, fields that are part of the identity (like `location`, `region`, or `resourceID`) are not mapped automatically in the spec by the from-proto mappers, so you must set them manually on the struct before converting to Unstructured.
4. **Use the short name:** Make sure you use the short name of the resource (e.g. `a.id.Repository`) when calling `u.SetName()`, instead of the full GCP name.

**Example Implementation:**
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
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Repository)
	u.SetGroupVersionKind(krm.ArtifactRegistryRepositoryGVK)

	// Set standard metadata such as labels (do not set project-id annotation if projectRef is on spec).
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
