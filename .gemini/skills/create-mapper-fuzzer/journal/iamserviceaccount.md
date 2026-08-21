# Journal: IAMServiceAccount KRM Fuzzer Implementation

Implemented the KRM round-trip fuzzer and status mapping logic for `IAMServiceAccount`.

## Observations & Findings
- **Proto vs. KRM Fields**: The `adminpb.ServiceAccount` protobuf field `unique_id` mapped directly to KRM Status `UniqueId`. The protobuf's `email` field is mapped both to KRM Status `Email` and used to derive KRM Status `Member` in the format `serviceAccount:{email}`.
- **Identity Fields**: `.project_id` represents part of the GCP URI and was marked as `Unimplemented_Identity`.
- **Etag and Oauth2ClientId**: These fields are not currently mapped in KRM, so they were marked as `Unimplemented_NotYetTriaged`.
- **Pre-existing Mappers**: Pre-existing `IAMServiceAccountSpec_FromProto` and `IAMServiceAccountSpec_ToProto` were already generated/available in the package, but status mappings had to be added manually.
