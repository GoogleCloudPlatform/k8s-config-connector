When implementing MemorystoreInstanceEndpoint, I found:
1. `MemorystoreInstanceEndpoint` does not represent an independent resource with its own URL path in GCP. Instead, it is a proxy/sub-resource that configures the `endpoints` field of a `MemorystoreInstance`.
2. Therefore, its external identity template matches the template of `MemorystoreInstance`: `"projects/{project}/locations/{location}/instances/{instance}"` (host `"memorystore.googleapis.com"`).
3. The format `"//memorystore.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/instances/{{INSTANCE}}"` is already listed in `cloudassetinventory_names.jsonl` (for `MemorystoreInstance`), so no exception was needed in `pkg/gcpurls/registry_test.go`.
4. In `getIdentityFromMemorystoreInstanceEndpointSpec`, we extract the identity by resolving/normalizing the required `spec.instanceRef` field to get its external reference URL and then parsing it.
