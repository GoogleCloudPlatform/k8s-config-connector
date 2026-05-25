### [2026-05-25] Types Implementer Best Practices
- **Acronyms**: Go field names for acronyms must use all caps (e.g., `RawDER`, not `RawDer`).
- **References**: When a field expects a resource name, you MUST use KCC style reference fields (e.g., `ServiceDirectoryServiceRef`). You are NOT allowed to add exceptions to `tests/apichecks/testdata/exceptions/missingrefs.txt`.
- **Unsupported Delete API**: If the GCP API does not support a Delete method for a resource, the KRM object can still be deleted. The direct controller `Delete` method should return an error but MUST include a clear `klog` message indicating that the resource itself won't be deleted and users should use the `abandon` policy. Do not leave "thought comments" in the code.
- **Client Generation**: When running `make generate-go-client ensure fmt`, ensure `// +genclient` is placed correctly without separating blank lines to the type definition.
