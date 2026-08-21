# BigtableGCPolicy direct types migration journal

## Observations and Learnings
1. **Custom Resource Schema Mapping:** `BigtableGCPolicy` is a KCC-only resource that maps a GCP Table GC policy (represented as `GcRule` nested inside a column family of a Bigtable table). Since it represents a nested concept, we mapped Kind `BigtableGCPolicy` to the Proto message `GcRule` in `apis/bigtable/v1beta1/generate.sh`.
2. **Types File Naming:** According to the skill instructions, the types file must match the lowercase proto message name. So `gcrule_types.go` is the correct types file name.
3. **Strict Schema Compatibility:** To guarantee 100% compatibility with the original baseline CRD schema, we handcoded `gcrule_types.go` and omitted standard direct controller fields such as `projectRef`, `location`, `observedState`, and `externalRef` because they were not present in the original baseline CRD.
4. **Validation via `diff-crds`:** We validated our changes using `dev/tasks/diff-crds` which showed zero diffs, guaranteeing that the KRM schema has remained completely unmodified.
