### [2026-05-12] KMSEKMConnection Delete Operation
- **Context**: Implementing the direct controller for `KMSEKMConnection`.
- **Problem**: The `google.cloud.kms.v1.EkmService` API does not expose a `DeleteEkmConnection` method.
- **Solution**: Implemented `Delete` in the direct controller Adapter to return an error stating that deletion is not supported. Users must use `cnrm.cloud.google.com/deletion-policy: abandon` to remove the resource from the Kubernetes cluster.
- **Impact**: The next agent should know that `EKMConnection` resources cannot be automatically cleaned up in GCP via the API.
