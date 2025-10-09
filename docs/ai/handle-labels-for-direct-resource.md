# How to Handle Labels for Direct Resources

This document provides instructions on how to handle GCP labels for direct resources in KCC. The standard approach is to use the Kubernetes `metadata.labels` field as the single source of truth for the labels on the GCP resource. This means the `labels` field should not be part of the resource's `Spec`.

To properly handle labels for a direct resource, follow these steps:

### 1. Remove Labels from the API Specification

First, ensure that the `labels` field is not part of the resource's `Spec`. If it exists, comment it out.

**File to edit:** `apis/backupdr/v1alpha1/backupdrbackupplan_types.go`

In `apis/backupdr/v1alpha1/backupdrbackupplan_types.go`, the `Labels` field in the `BackupDRBackupPlanSpec` struct should be commented out:

```go
type BackupDRBackupPlanSpec struct {
    // ... other fields ...

    // KRM-style labels for the resource.
    // +optional
    // Labels map[string]string `json:"labels,omitempty"`

    // ... other fields ...
}
```

After commenting out the field, regenerate the CRDs by running:

```bash
dev/tasks/generate-crds
```

### 2. Update the Fuzzer Test

Next, update the fuzzer test to acknowledge that labels are intentionally unimplemented in the spec.

**File to edit:** `pkg/controller/direct/backupdr/backupdrbackupplan_fuzzer.go`

Add the following line to your fuzzer configuration:

```go
f.Unimplemented_LabelsAnnotations(".labels")
```

This tells the fuzzer to ignore the `.labels` field during testing, preventing false positives.

### 3. Update the Controller to Handle Labels

Finally, modify the controller to correctly map Kubernetes metadata labels to the GCP resource labels. The recommended pattern is to handle all the logic for constructing the desired GCP resource state, including labels, within the `AdapterForObject` function.

**File to edit:** `pkg/controller/direct/backupdr/backupdrbackupplan_controller.go`

In `pkg/controller/direct/backupdr/backupdrbackupplan_controller.go`, the `AdapterForObject` function would be updated as follows:

```go
import (
    // ... other imports
    "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
)

// ...

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
    obj := &krm.BackupDRBackupPlan{}
    if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
        return nil, fmt.Errorf("error converting to %T: %w", obj, err)
    }

    // ...

    // Convert KCC resource spec to GCP proto message
    mapCtx := &direct.MapContext{}
    desiredProto := BackupDRBackupPlanSpec_ToProto(mapCtx, &obj.Spec)
    if mapCtx.Err() != nil {
        return nil, mapCtx.Err()
    }

    // Handle GCP Labels
    desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

    // ...

    return &Adapter{
        // ...
        desired: desiredProto,
        // ...
    }, nil
}
```

By preparing the desired state in `AdapterForObject`, the `Create` and `Update` methods become simpler and can use the `desired` proto message stored in the adapter directly.

### 4. Update Fixture Tests

After implementing the label handling in the controller, update the resource fixture tests to reflect the changes and verify the new behavior.

**Path to edit:** `pkg/test/resourcefixture/testdata/basic/backupdr/v1alpha1/backupdrbackupplan/`

1.  **Remove `spec.labels` from existing tests:**
    If `create.yaml` or `update.yaml` in existing test fixtures contain a `labels` field within the `spec`, remove it. The labels are now managed via `metadata.labels`.

2.  **Add a new test case for labels:**
    Create a new test fixture directory (e.g., `backupdrbackupplan-labels`). This test will specifically verify that `metadata.labels` are correctly propagated to the GCP resource.

    **`create.yaml` for `backupdrbackupplan-labels` test:**
    ```yaml
    apiVersion: backupdr.cnrm.cloud.google.com/v1alpha1
    kind: BackupDRBackupPlan
    metadata:
      name: backupdrbackupplan-labels
    labels: 
        app.kubernetes.io/name: "mock-app"
        app.kubernetes.io/instance: "mock-instance"
        app.kubernetes.io/version: "v1.0.0"
        app.kubernetes.io/component: "mock-component"
        app.kubernetes.io/part-of: "mock-part-of"
        app.kubernetes.io/managed-by: "configmanagement.gke.io"
        applyset.kubernetes.io/id: mock-applyset-id
        configmanagement.gke.io/sync-name: mock-sync-name
        configmanagement.gke.io/sync-namespace: mock-sync-namespace
        custom-label: "foo" # valid label
    spec:
      location: us-central1
      # ... other required spec fields
    ```

    **`update.yaml` for `backupdrbackupplan-labels` test:**
    ```yaml
    apiVersion: backupdr.cnrm.cloud.google.com/v1alpha1
    kind: BackupDRBackupPlan
    metadata:
      name: backupdrbackupplan-labels
    labels: 
        app.kubernetes.io/name: "mock-app"
        app.kubernetes.io/instance: "mock-instance"
        app.kubernetes.io/version: "v1.0.0"
        app.kubernetes.io/component: "mock-component"
        app.kubernetes.io/part-of: "mock-part-of"
        app.kubernetes.io/managed-by: "configmanagement.gke.io"
        applyset.kubernetes.io/id: mock-applyset-id
        configmanagement.gke.io/sync-name: mock-sync-name
        configmanagement.gke.io/sync-namespace: mock-sync-namespace
        custom-label: "bar" # updated valid label
    spec:
      location: us-central1
      # ... other required spec fields
    ```
    This `update.yaml` tests updating one label (`custom-label`), and implicitly removing the invalid labels.

By following these steps, you ensure that labels are handled consistently and correctly for direct resources, leveraging the standard Kubernetes `metadata.labels` field as the source of truth.
