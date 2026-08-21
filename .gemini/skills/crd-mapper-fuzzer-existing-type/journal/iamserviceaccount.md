# Journal: IAMServiceAccount Migration to Direct Types

## GVK Redeclaration Issue
When migrating `IAMServiceAccount`, we discovered that `serviceaccountkey_reference.go` had already declared `IAMServiceAccountGVK`. Adding `IAMServiceAccountGVK` in `serviceaccount_types.go` as part of the new direct types led to a redeclaration build error during `go vet`.
- **Mitigation:** We surgically removed the duplicate declaration of `IAMServiceAccountGVK` from `serviceaccountkey_reference.go` and kept the canonical definition in `serviceaccount_types.go`.

## Fully Automated Mapper Matching
Since the fields in the handcoded `IAMServiceAccountSpec` structure match the proto schema of `google.iam.admin.v1.ServiceAccount` exactly (`Description`, `Disabled`, `DisplayName`), the controller generator successfully produced automatic mapping functions (`IAMServiceAccountSpec_FromProto` and `IAMServiceAccountSpec_ToProto`) inside `mapper.generated.go`.
- No handcoded mappers under a separate `mappers.go` were required for this resource, illustrating the benefit of matching struct and proto field names exactly.
