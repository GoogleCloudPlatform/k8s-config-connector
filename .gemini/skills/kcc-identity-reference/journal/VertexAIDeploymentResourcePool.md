When implementing VertexAIDeploymentResourcePool, I found:
1. The resource was missing from `docs/ai/metadata/cloudassetinventory_names.jsonl` (not supported by CAIS).
2. I added an exception to `pkg/gcpurls/registry_test.go`'s `ignoredTemplates`: `"//aiplatform.googleapis.com/projects/{}/locations/{}/deploymentResourcePools/{}"`.
3. In `deploymentresourcepool_identity.go`, the template segment MUST map exactly to the struct field, lowercased. Since the field is `DeploymentResourcePool`, the template string must use `{deploymentresourcepool}`, NOT `{deployment_resource_pool}`. Using the underscore version causes a panic at initialization: `field "deployment_resource_pool" not found in struct`.
