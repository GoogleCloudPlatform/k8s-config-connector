### [2026-05-26] ApigeeRegistryInstance Scaffolding
- **Context**: Implementing ApigeeRegistryInstance KRM types and IdentityV2.
- **Problem**: Need to replace `cmek_key_name` primitive with a reference in the spec.
- **Solution**: Defined `Instance_Config` in `instance_types.go` overriding the generated type, replacing `CmekKeyName *string` with `CmekKeyNameRef *refsv1beta1.KMSCryptoKeyRef`, successfully hiding the generated `Instance_Config` from `types.generated.go`. Also ensured `Location` is a pointer `*string` to pass the primitive pointer invariants.
- **Impact**: Demonstrates that replacing generated structures in `instance_types.go` properly prunes them from `types.generated.go` upon running `generate-types`.

### [2026-05-19] ApigeeRegistry API Client Quirk and Identity
- **Context**: Implementing direct types for `ApigeeRegistryAPI`.
- **Problem**: The standard `apigeeregistry.NewRegistryRESTClient` does not exist in `cloud.google.com/go/apigeeregistry/apiv1`. Additionally, creating `UpdateApiRequest` requires a specific `UpdateMask`. 
- **Solution**: Use `apigeeregistry.NewRegistryClient(ctx, opts...)` instead of `NewRegistryRESTClient`. For the `UpdateMask`, use `common.CompareProtoMessage` to build a `fieldmaskpb.FieldMask`.
- **Impact**: The next agent working on `ApigeeRegistry` resources (like `ApigeeRegistryInstance`) should also use `NewRegistryClient` and take care to populate the `UpdateMask` dynamically if required.
