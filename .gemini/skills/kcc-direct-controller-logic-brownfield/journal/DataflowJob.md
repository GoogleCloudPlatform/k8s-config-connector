# DataflowJob Migration Journal

## Observation & Design Decisions

- **Region Resolution Fallback:** `DataflowJobSpec` allows `Region` to be optional, but the GCP API regional templates service always requires a region (regional endpoint) to launch a job. To solve this without breaking existing user manifests where only `Zone` is specified (e.g., `zone: us-central1-a`), we implemented a fallback in our direct controller. If `Region` is not specified, we parse `Zone` (which usually has form `<region>-<letter>`) to extract the parent region name (e.g. `us-central1-a` -> `us-central1`).
- **Create vs. Update APIs:** We aligned precisely with Terraform's backend behavior. Standard creations are dispatched via `CreateJobFromTemplate` on `api.TemplatesClient`. For mutable fields (like parameter/metadata generation updates on Streaming/Stream jobs), updates are reconciliation-driven and dispatched via `LaunchTemplate` on `api.TemplatesClient` with `Update = true` set on `LaunchTemplateParameters`.
- **Batch Jobs Immutability:** Batch jobs cannot be updated. Any metadata generation change on an active Batch job raises a descriptive error to alert the user that updates are not permitted on batch templates.
