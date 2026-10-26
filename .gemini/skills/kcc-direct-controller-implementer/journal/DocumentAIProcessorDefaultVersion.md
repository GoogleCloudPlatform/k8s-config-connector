# DocumentAIProcessorDefaultVersion Controller Implementation Journal

## Observations & Quirks

- **Non-Deletable Resource Behavior**: DocumentAI processor default versions cannot be deleted from Google Cloud (there is no deletion API in GCP DocumentAI). The model is registered using `registry.CannotBeDeleted()`, and the `Delete` operation in `ProcessorDefaultVersionAdapter` is a no-op returning `true, nil` to cleanly remove the Kubernetes object without getting stuck in finalization.
- **Regional Endpoint Routing**: DocumentAI is a regional service (`us`, `eu`, etc.). The GAPIC REST client constructor must specify `option.WithEndpoint(fmt.Sprintf("%s-documentai.googleapis.com", location))` based on the processor location extracted from the identity.
- **Diff Suppression**: In GCP, the processor default version can be returned using either project numbers or project IDs. Using regex-based project number diff suppression (matching legacy Terraform behavior) prevents unnecessary network calls to Cloud Resource Manager during diff calculation and prevents spurious `SetDefaultProcessorVersion` calls.
- **Alias Resolution**: DocumentAI processors can have processor version aliases (e.g., `rc`, `stable`). The diff logic checks `actual.ProcessorVersionAliases` to recognize when a requested alias already points to the current default processor version.
