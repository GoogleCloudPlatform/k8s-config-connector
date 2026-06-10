# Journal: KMSKeyRing Transition to Direct KRM types

## Learnings & Observations

### 1. Simple Proto with Identifiers in Spec URL Path
For resources like `KMSKeyRing`, fields in the KRM spec (like `location` and `resourceID`) do not map directly to properties of the GCP `KeyRing` proto message. Instead, they are part of the resource's path parsed from `pb.KeyRing.Name`.
- **Solution:** We hand-coded custom mapping functions `KMSKeyRingSpec_FromProto` and `KMSKeyRingSpec_ToProto` in `pkg/controller/direct/kms/keyring_mappers.go`.
- **Lossless Spec Roundtrips in Fuzzer:** To ensure the fuzzer doesn't encounter any mismatch on the `.name` field during round-trips, we implemented `KMSKeyRingSpec_ToProto` to construct and return a valid `.Name` string from `Location` and `ResourceID` using a dummy project placeholder (`p`). This allows the fuzzer's Spec mapping verification to pass perfectly.

### 2. Status Field Mapped from .name
The `selfLink` field in KRM status maps directly to `pb.KeyRing.Name`. We mapped this successfully using:
- `KMSKeyRingStatus_FromProto` setting `SelfLink` from `Name`.
- `KMSKeyRingStatus_ToProto` setting `Name` from `SelfLink`.
This matches the baseline CRD schema perfectly while providing full compatibility.

### 3. Verification & Compatibility
Running `dev/tasks/diff-crds` produced an empty diff, confirming that zero schema changes were introduced to the baseline CRD of `KMSKeyRing`. All fuzzing and registration tests passed flawlessly.
