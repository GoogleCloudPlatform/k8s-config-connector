# Journal: ComputeDiskResourcePolicyAttachment Migration

## Findings

1. **Unresolvable Dependency Placeholder**:
   The dependency resource `ComputeDisk` in both `diskresourcepolicyattachmentautogen` and `diskresourcepolicyattachmentautogen-direct` had the `image` field set to `${data.google_compute_image.my_image.self_link}`. Since this placeholder is not resolved by the KCC E2E Go test runner, it caused real GCP runs to fail with:
   `Error resolving image name '${data.google_compute_image.my_image.self_link}': Could not find image or family ${data.google_compute_image.my_image.self_link}`

2. **Schema Mismatch on ComputeDisk v1beta1**:
   The `v1beta1` schema of `ComputeDisk` defines `imageRef` rather than `image` as a field. Using `.spec.image` resulted in client-side / API schema validation failure:
   `.spec.image: field not declared in schema`

3. **Resolution**:
   Changed the `image` field in `dependencies.yaml` to `imageRef.external: debian-cloud/debian-11` in both test fixtures. This successfully reconciled the dependency `ComputeDisk` on real GCP, allowing the direct migration E2E tests and `record-gcp` to run and record the HTTP traffic cassettes.
