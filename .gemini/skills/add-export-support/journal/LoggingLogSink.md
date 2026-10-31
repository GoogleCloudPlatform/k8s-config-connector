# LoggingLogSink Export Implementation Journal

Implementing export support for `LoggingLogSink` went extremely smoothly. Below are the key design choices and observations that are helpful for other developers.

## Observations & Best Practices

1. **Setting the GVK and Name:**
   We assigned `u.Object = uObj` before setting the Name and GVK. This prevents any metadata fields from being wiped, as noted in the critical instructions of `add-export-support/SKILL.md`.

2. **Handling Parent References:**
   `LoggingLogSinkSpec` supports three alternative parent binders:
   - `projectRef`
   - `folderRef`
   - `organizationRef`

   We inspected the resolved parent from the resource identity (`a.id.Project`, `a.id.Folder`, `a.id.Organization`) and mapped it to the corresponding `.Ref` field in the Spec dynamically. Because `projectRef` is a reference-bound field, we did not call `export.SetProjectID(...)` on the unstructured object to prevent setting the deprecated project-id annotation.

3. **Handling Resource ID:**
   We mapped `spec.resourceID` using `direct.LazyPtr(a.id.Sink)`.

4. **Handling Write-Only / Non-Proto Fields:**
   The `uniqueWriterIdentity` boolean field is not returned in the standard `loggingpb.LogSink` GCP representation (it is only passed as a query param on creation/update).
   However, we were able to infer its desired value during export:
   - If the returned `WriterIdentity` is not empty and is not the default shared logging service account (`serviceAccount:cloud-logs@system.gserviceaccount.com`), then `uniqueWriterIdentity` must have been set to `true`.
   - This matched the legacy Terraform controller's expectation perfectly, resulting in `0` alignment differences in E2E tests.
