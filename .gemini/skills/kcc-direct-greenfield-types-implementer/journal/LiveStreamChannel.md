# LiveStreamChannel Journal

## Observations & Patterns

### 1. Resource References
- `InputAttachment.input` is a reference to `LiveStreamInput`, resolved as `InputRef *LiveStreamInputRef`.
- `StaticOverlay.asset` is a reference to `LiveStreamAsset`, resolved as `AssetRef *LiveStreamAssetRef`. `LiveStreamAssetRef` was created in `apis/livestream/v1alpha1/livestreamasset_reference.go`.
- `Encryption.SecretManagerSource.secret_version` references SecretManager SecretVersion, resolved as `SecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef`.

### 2. Empty Structs / Oneof markers in CRD Validation
- `Encryption` contains empty message structs (`Aes128Encryption`, `Clearkey`, `Fairplay`, `Playready`, `SampleAesEncryption`, `Widevine`).
- To satisfy OpenAPI / Kubernetes CRD validation without schema errors about empty object properties, these types are defined in `livestreamchannel_types.go` annotated with:
  ```go
  // +kubebuilder:pruning:PreserveUnknownFields
  // +kubebuilder:validation:XPreserveUnknownFields
  ```

### 3. Custom Protobuf Types
- `TimecodeConfig.TimeZone` uses `google.type.TimeZone`. Custom mapper functions `TimeZone_FromProto` and `TimeZone_ToProto` were implemented in `pkg/controller/direct/livestream/livestreamchannel_mappings.go`.
