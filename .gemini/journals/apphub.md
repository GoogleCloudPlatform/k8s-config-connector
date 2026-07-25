# AppHub Service Journal

### 2026-07-24 AppHubServiceProjectAttachment Greenfield Controller Implementation
- **Context**: Greenfield direct controller implementation for AppHubServiceProjectAttachment (GVK: `apphub.cnrm.cloud.google.com/v1alpha1`, Kind: `AppHubServiceProjectAttachment`).
- **Problem**:
  1. `AppHubServiceProjectAttachment` is a completely immutable resource once created (no update RPC exists on the AppHub GCP Client). Standard direct controller templates expect mutable spec fields.
  2. The service project in the incoming spec can be a project ID or project number. However, the output returned from GCP always maps the service project to the project number (e.g., `projects/<number>`). This can lead to false positives in diff checks when reconciling an existing resource.
  3. During `AdapterForObject` operation, the standard identity parser `ParseAppHubServiceProjectAttachmentIdentity` fails for newly created resources because `status.externalRef` is initially empty.
- **Solution**:
  1. In `Update`, we performed a check on spec fields using a structured diff. If a diff was detected on any spec field, we returned a descriptive error stating that the resource is completely immutable and cannot be updated.
  2. We implemented a robust project-comparison function `compareServiceProjectAttachment` that uses `refs.ResolveProject` to resolve both the desired and actual projects to their canonical metadata. This allows safely comparing the resolved project IDs, preventing any false positive diffs due to ID vs number representations.
  3. We resolved the identity in `AdapterForObject` by utilizing the Resource-level `GetIdentity` method (`obj.GetIdentity(ctx, reader)`), which gracefully defaults to constructing the identity from the spec if `status.externalRef` is empty.
- **Impact**: Ensures that immutable service project attachments are reconciled seamlessly, handles project ID/number variations correctly without false diffs, and prevents creation failure for new resources.
