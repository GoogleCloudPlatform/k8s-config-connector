# Notes on Kcc Identity Reference

*   When checking the format for the `External` docstring in `_reference.go` or the `gcpurls.Template`, be aware that path segments are sometimes camelCase in GCP URLs (e.g., `entryGroups` instead of `entrygroups`). Match the case used in existing manual implementations if they exist.
*   When updating a resource's identity struct to `IdentityV2`, be sure to check for existing usages of the struct and its old methods (e.g. `.Parent()`, `.ID()`) in dependent identity files and direct controllers, and update them to use the new fields (e.g. `.Project`, `.Location`, etc.). The compiler is your friend: remove the functions, then run `go vet ./...` or `go build ./...` to look for references to functions that no longer exist.
* `getIdentityFrom<Kind>Spec` takes `client.Object`, so you can pass `*unstructured.Unstructured` directly to it in `Normalize` instead of converting it first.
* In `gcpurls.Template`, when you define a template variable like `{entrytype}`, it maps to the struct field `EntryType` case-insensitively. Avoid underscores like `{entry_type}` in the template variable if the struct field is camel case without underscores, as it will cause a panic during initialization.
* `ParseExternalToIdentity` takes no arguments and returns `(identity.Identity, error)`.
