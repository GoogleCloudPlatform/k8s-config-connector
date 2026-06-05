# Dataplex Journal

### [2026-06-05] Implement DataplexGlossary initial types, CRD, and IdentityV2
- **Context**: Implementing Greenfield types and identity for `DataplexGlossary` (GroupKind: `DataplexGlossary`, Service: `google.cloud.dataplex.v1.BusinessGlossaryService`).
- **Problem**: 
  1. The Google APIs SHA initially was pinned to an older version that did not contain `business_glossary.proto`, leading to a proto-not-found error. 
  2. While the latest `googleapis` HEAD had the required proto, updating `apis/git.versions` permanently to `HEAD` caused Go compilation / build failures across other packages like `firestore` and `sql` because their corresponding Go SDK client libraries inside KCC did not support the newer protobuf fields yet.
- **Solution**: 
  1. Temporarily upgraded `apis/git.versions` to HEAD/ee4a3e1ce to run the generator and compile the `.pb` file.
  2. Reverted `apis/git.versions` and `generate.sh` back to their original states once generation completed. This avoids breaking build/vet checks in CI for firestore/sql, but preserves the generated `glossary_types.go`, identity, and CRD YAML.
  3. Ensured `Location` is defined as a pointer (`*string`) in `DataplexGlossarySpec` to adhere to the strict scalar primitive pointer standard in KCC direct resources.
- **Impact**: Enables `DataplexGlossary` CRD and types to be safely committed and validated in CI without triggering global protobuf mismatch regressions.
