# ComputeSSLPolicy Direct Types Journal

- **Short Name Pluralization**: The generator scaffolded `gcpcomputesslpolicys` by default. However, to maintain backwards compatibility, this was manually updated to `gcpcomputesslpolicies` (semicolon-separated). `controller-gen` correctly parses and generates multiple short names if they are semicolon-separated in the `shortName` tag.
- **MinTLSVersion Field Alignment**: The protobuf has `min_tls_version` mapping to `MinTLSVersion` in Go. However, the existing KCC CRD represents this as `minTlsVersion`. We hand-wrote the Go field as `MinTlsVersion` and annotated it with `// +kcc:proto:field=google.cloud.compute.v1.SslPolicy.min_tls_version` to maintain exact backward compatibility in the CRD schema. This works seamlessly.
