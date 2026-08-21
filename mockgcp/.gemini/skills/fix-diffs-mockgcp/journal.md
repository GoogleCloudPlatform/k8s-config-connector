# Notes on MockGCP Diffs

* `populateDefaultsFor<Resource>`: Most resources in `mockgcp` have this pattern to inject server-side defaults. Call this at the end of resource creation (POST/Insert) and before returning the resource (GET).
* Normalization wildcard paths: Using `NormalizingVisitor`, you can match values using `event.VisitResponseStringValues(func(path, value string) { ... })`. You can use `strings.HasSuffix(path, ".ipAddress")` to target deep paths inside arrays without specifying `[]` index values.
* If testing operations, mock operations return `RUNNING` by default in `mockgcp/mockcompute/operations.go`, whereas some real GCP operations might start as `PENDING`. This is an accepted difference for now unless strict parity is required.