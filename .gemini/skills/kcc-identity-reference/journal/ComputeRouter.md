# ComputeRouter Identity and Reference Journal

## Observations & Learnings

- **Standard Regional Identity Template:** ComputeRouter utilizes a standard regional path format: `projects/{project}/regions/{region}/routers/{router}`. This maps cleanly to a single `gcpurls.Template`.
- **Compute URL Prefix Trimming:** Since this is a Compute resource, we integrated `apirefs.TrimComputeURIPrefix` in `FromExternal` to robustly strip standard GCP Compute Engine API endpoints (e.g. `https://compute.googleapis.com/...`), ensuring maximum compatibility with old/new formats.
- **Reference Backwards Compatibility with Fallback:** In `ComputeRouterRef.Normalize`, we implemented a fallback pattern checking `status.selfLink` on the Unstructured object, before resolving via spec. This preserves backward compatibility for resources that do not utilize `status.externalRef`.
- **Golden Identity File Updates:** Implementing this pattern automatically resolved CAIS URL generation for ComputeRouter, correcting `caisURL: unknown` to `caisURL: //compute.googleapis.com/projects/${projectId}/regions/${region}/routers/computerouter-${uniqueId}` in multiple test fixtures during golden file validation.
- **Zero Schema Schema Modifications:** Followed the strict rule to not change the schema of `ComputeRouter` or any other resource.
