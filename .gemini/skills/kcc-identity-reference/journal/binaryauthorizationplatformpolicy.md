# BinaryAuthorizationPlatformPolicy Identity and Reference Journal

## Observations

- `BinaryAuthorizationPlatformPolicy` has a `platform` field in its GCP URL: `projects/{project}/platforms/{platform}/policies/{policy}`.
- The `platform` field is a custom field in the `Spec`, not a standard `location`.
- The `getIdentityFromBinaryAuthorizationPlatformPolicySpec` function must handle `*unstructured.Unstructured` objects because it is called by `Normalize` with an unstructured object.
- When accessing custom fields from the spec in `getIdentityFrom...Spec`, use `unstructured.NestedString(u.Object, "spec", "fieldName")` after ensuring the object is converted to unstructured if it's not already.

## Shortcomings in SKILL.md

- The skill assumes that all fields in the identity can be resolved using `refs` helpers (e.g. `ResolveProjectID`, `GetLocation`). It should mention that for custom fields, explicit handling of `*unstructured.Unstructured` is required in `getIdentityFrom...Spec`.
