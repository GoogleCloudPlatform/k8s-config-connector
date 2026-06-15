# DNSPolicy Journal

## Observations
- Generated KRM types for DNSPolicy from OpenAPI schema and merged it with DNSManagedZone.
- Implemented support for multiple resources in the same service version in `dev/tools/openapi-to-krm/main.go` so that both resources can generate types in a single execution.
- Configured `--resource "DNSPolicy:Policy"` in `apis/dns/v1beta1/generate.sh` to generate the new fields and structures successfully.
- Resolved reference `networkRef` of type `computev1beta1.ComputeNetworkRef` instead of raw `v1alpha1.ResourceRef`.
- Validated CRD schema compatibility using `dev/tasks/diff-crds` resulting in 100% exact match of OpenAPI definitions, ensuring zero disruption or breakage for existing KCC users.
- Clean compilation checked via `go vet` and custom linters.
