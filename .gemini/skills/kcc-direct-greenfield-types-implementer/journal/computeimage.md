# Greenfield Types Implementation Journal - ComputeImage

## Kind: ComputeImage
- **GCP Resource:** Image
- **GCP Service:** `google.cloud.compute.v1`
- **KRM Package:** `apis/compute/v1beta1`

## Observations and Learnings
- **Pre-existing Identity & Reference:** The identity and reference files (`computeimage_identity.go` and `computeimage_reference.go`) were already hand-written in the target package. We reused and integrated them.
- **Reference Dependencies:** `ComputeImage` references several other types: `ComputeDisk`, `ComputeSnapshot`, `KMSCryptoKey`, and `IAMServiceAccount`. `KMSCryptoKeyRef` and `IAMServiceAccountRef` were already defined under `apis/refs/v1beta1/`. `ComputeDiskRef` and `ComputeSnapshotRef` were not defined, so we implemented them under `apis/refs/v1beta1/computerefs.go` with respective resolver methods to maintain structural type integrity.
- **Root Status Fields:** For backward compatibility, the four status fields `archiveSizeBytes`, `creationTimestamp`, `labelFingerprint`, and `selfLink` must be kept at the root of `ComputeImageStatus`. We achieved this using explicit `// +kcc:proto:field` annotations directly under the root status struct.
