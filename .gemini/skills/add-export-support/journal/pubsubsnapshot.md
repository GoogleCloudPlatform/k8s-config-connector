# PubSubSnapshot Export Support Implementation Journal

## Observations
1. **Reference-Bound Project Binding**: Since `PubSubSnapshot` uses the `ProjectRef` field for its project binding, we set `obj.Spec.ProjectRef.External = a.id.Project` in the direct `Export` method and did not call `export.SetProjectID(...)` to avoid adding the deprecated `cnrm.cloud.google.com/project-id` annotation.
2. **KRM-only Spec Fields**: The `pubSubSubscriptionRef` field is required when creating a snapshot, but it is not part of the standard GCP `pb.Snapshot` returned from GET requests. Therefore, during export, this field is set to `nil`.
3. **No Legacy Controller Alignment**: `PubSubSnapshot` is registered solely as a direct controller in Config Connector. This means there were no old controller/legacy golden files to align with, making the direct exporter the sole source of truth for golden traffic and object output files.
