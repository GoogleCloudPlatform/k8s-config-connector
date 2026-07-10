# Export Support for LoggingLogExclusion Journal

## Findings and Observations

During the implementation of export support for `LoggingLogExclusion`, several key observations were made:

1. **Flexible Parent Bindings Resolution:**
   `LoggingLogExclusion` supports project-level, folder-level, organization-level, and billing account-level scopes (via `projectRef`, `folderRef`, `organizationRef`, and `billingAccountRef` respectively). During export, we determined the scope by checking which field is populated in the `LoggingLogExclusionIdentity` parsed from the resource URI.

2. **ResourceID Mapping in Spec:**
   Like other direct controllers, identity fields (such as `resourceID` / the exclusion ID) must be manually populated on the exported KRM `Spec` during translation since the from-proto mapping doesn't automatically map these fields.

3. **Avoiding standard SetProjectID helper:**
   As `LoggingLogExclusion` is a reference-bound resource (it uses `projectRef`/`folderRef`/etc.), we avoided using `export.SetProjectID(...)` which would otherwise inject a `cnrm.cloud.google.com/project-id` annotation. Instead, we correctly mapped the parent via the reference fields under `spec`.
