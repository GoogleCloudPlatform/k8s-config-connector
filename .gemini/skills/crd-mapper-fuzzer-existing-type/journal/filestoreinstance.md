# Journal: FilestoreInstance Direct KRM Types Migration

During the migration of `FilestoreInstance` to direct KRM types, we made several key observations and technical choices to ensure 100% schema-compatibility:

1. **Protocol and Proto mapping**:
   - The GCP API proto package is `google.cloud.filestore.v1`.
   - The GCP resource `Instance` maps to the KCC `FilestoreInstance` kind.
   - We configured `apis/filestore/v1beta1/generate.sh` to run `generate-types` and `generate-mapper` with the flag `--include-skipped-output`.

2. **Schema Matching & Custom References**:
   - **`projectRef`**: The baseline CRD schema does not contain the `kind` field under `projectRef`. To match this exactly, we imported and used `refs.ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` instead of a local reference structure or the standard `v1beta1.ProjectRef`.
   - **`sourceBackupRef`**: Under `fileShares[]`, the reference to `FilestoreBackup` is defined as `SourceBackupRef` in the baseline, and did not include a `.kind` field. We defined a local custom reference type `FilestoreBackupRef` without a `Kind` field to avoid adding `.kind` to the OpenAPI schema.
   - **Format Validation**: The status field `createTime` in the baseline has format `date-time`. We explicitly annotated the generated `CreateTime *string` with `// +kubebuilder:validation:Format=date-time` to achieve absolute parity.

3. **Validation**:
   - Running `dev/tasks/diff-crds` returned an empty diff, confirming 100% schema compatibility with the baseline.
