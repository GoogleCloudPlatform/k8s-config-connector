# BackupDRBackupPlan Export Support Journal

## Overview
Implemented export support for the `BackupDRBackupPlan` resource.

## Key Observations and Learnings
1. **Reference-Bound Project Binding**:
   - `BackupDRBackupPlan` uses `spec.projectRef` in KRM instead of the `cnrm.cloud.google.com/project-id` annotation.
   - Per the export guidelines, for reference-bound resources, we map `obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}` and completely avoid calling `export.SetProjectID(...)` to prevent adding the redundant annotation.

2. **Control Plane Assets for Headless Testing**:
   - Running the E2E export tests under `E2E_KUBE_TARGET=mock` caused CRD polling timeouts (as the mock Kubernetes API server does not establish the CRDs' status).
   - We resolved this by installing and configuring the `envtest` control plane assets using:
     `go run sigs.k8s.io/controller-runtime/tools/setup-envtest@latest use -p path`
   - This provided the dynamically determined local assets path (`/root/.local/share/kubebuilder-envtest/k8s/1.36.0-linux-amd64`).
   - Running the tests with `E2E_KUBE_TARGET=envtest` and `KUBEBUILDER_ASSETS=/root/.local/share/kubebuilder-envtest/k8s/1.36.0-linux-amd64` allowed the CRD installation to succeed instantly and all fixtures to reconcile/export successfully.
