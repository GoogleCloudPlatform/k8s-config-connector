
# How to Implement GCP Label Support in KCC

## Background

In KCC, we need to distinguish between two types of labels:

1.  **Kubernetes Labels:** Standard Kubernetes metadata attached to any KCC resource via the `.metadata.labels` field.
2.  **GCP Labels:** Labels applied to the underlying GCP resource.

Historically, KCC controllers have treated Kubernetes labels as GCP labels. This has some benefits, but also several drawbacks, especially because GCP label support can differ significantly from one GCP service to another, whereas Kubernetes label support is uniform.

The modern, direct-controller approach in KCC is to treat GCP labels as a part of the resource's specification, specifically in a field like `.spec.labels`. This avoids ambiguity and provides a clearer user experience.

## The Goal: Coherently Support Both `.spec.labels` and `.metadata.labels`

To provide a coherent and consistent user experience, any KCC resource that supports GCP labels MUST support them in two ways: `.spec.labels` and `.metadata.labels`. This applies to all resources, including newly added ones, not just those with pre-existing usage of `.metadata.labels`.

*   `.spec.labels`: This is the preferred, modern way for users to specify GCP labels. It offers a clear distinction from Kubernetes metadata.
*   `.metadata.labels`: This is the legacy method. We continue to support it for backward compatibility and to allow users to manage GCP labels using standard Kubernetes tooling like `kubectl` and `kustomize`.

## Implementation Guide

To correctly implement GCP label support in a direct controller, follow these steps.

### 1. Check if the GCP Resource Supports Labels

First, determine if the target GCP resource supports labels. You can do this by examining its api definition under `apis/<service>/<version>/<resource>_types.go` (or apis/<service>/<version>/types.generated.go). If the resource contains a `labels` field (or a field with a similar name and purpose), then the resource supports GCP labels.

### 2. Use the `label.ComputeLabels` Function

The core of the implementation is the `label.ComputeLabels()` function, located in `pkg/label/label.go`.

This function intelligently computes the final set of GCP labels that should be applied to the resource, handling the logic of prioritizing `.spec.labels` while falling back to `.metadata.labels` for backward compatibility.

**How it works:**

1.  It inspects the unstructured resource object.
2.  **If `spec.labels` is present:** It uses the value of this field as the definitive set of GCP labels.
3.  **If `spec.labels` is NOT present:** It falls back to using the labels from `.metadata.labels`. In this case, it automatically filters out any labels that have a Kubernetes-specific prefix (i.e., contain a `/`) to prevent conflicts with GCP's more restrictive label syntax.
4.  It adds the `cnrm.cloud.google.com/managed-by-kcc: "true"` label to the final set.
5.  It temporarily writes the final, computed set of labels into the `spec.labels` field of the in-memory, copied unstructured object. This modification is NOT persisted to the `spec` of the actual resource in the Kubernetes cluster.

### 3. Call `ComputeLabels` in Your Controller

You must call `label.ComputeLabels()` within your controller's `AdapterForObject` function, before you convert the unstructured object into your typed Go struct.

This ensures that the `spec.labels` field is correctly populated with the final, intended set of GCP labels before any further processing, creation, or update logic is executed.

### Example Implementation

You can find a reference implementation in the APIGatewayAPI controller.

**File:** `pkg/controller/direct/apigateway/api_controller.go`

```go
// pkg/controller/direct/apigateway/api_controller.go

func (m *apiModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.APIGatewayAPI{}

	copied := u.DeepCopy()
	if err := label.ComputeLabels(copied); err != nil {
		return nil, err
	}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(copied.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

    // ... rest of the function
}
```

By following this pattern, you ensure that your controller correctly and consistently handles GCP labels, providing a clear path for users while maintaining backward compatibility.
