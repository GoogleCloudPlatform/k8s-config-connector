# ManagedKafka Journal

### [2026-06-15] Underscores in URL Template Variable Matching
- **Context**: Implementing IdentityV2 for `ManagedKafkaConnectCluster`.
- **Problem**: The GCP URL path contains the placeholder `{connect_cluster}`, which `gcpurls.Template` parses. It then attempts to find a matching field in the identity struct by converting both the template placeholder and Go struct fields to lower case and matching them. A Go field named `ConnectCluster` is lower-cased to `"connectcluster"`, which does not match the lower-cased template placeholder `"connect_cluster"`, leading to a runtime panic during initialization.
- **Solution**: Named the struct field exactly `Connect_cluster string`. The lower-cased field name is `"connect_cluster"`, which matches the placeholder `"connect_cluster"` perfectly, resolving the initialization panic.
- **Impact**: When implementing IdentityV2 for other resources with underscores in their URL template path parameters, ensure that the Go struct field retains the underscore (e.g. `Field_name`) to match the placeholder.
