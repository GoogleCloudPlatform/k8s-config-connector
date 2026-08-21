# CertificateManagerCertificateMapEntry Journal

## Observations & Learnings

### 1. Handling Name Collisions in Protobuf/Go Oneof Fields
In Go protobuf generated code, if there is a `oneof` field name that collides with an enum type (e.g. both are named `Matcher` or `matcher`), the protobuf generator avoids the collision by adding a trailing underscore `_` to the wrapper struct name. 
In our case:
- The enum type was `pb.CertificateMapEntry_Matcher`
- The wrapper struct type was `pb.CertificateMapEntry_Matcher_` (with a trailing underscore)

When manually instantiating or matching oneof types, always inspect the generated `pb` package symbols (or write a quick reflection helper to list/compile-check) to verify if trailing underscores are present on wrapper structures.

### 2. Fields in GCP API URL vs. Message Payload
Some fields in the baseline CRD (like `projectRef` and `mapRef`) correspond to the GCP resource identity components in the API URL (e.g. `projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}`).
These fields do not exist as properties inside the protobuf message `CertificateMapEntry` itself.
- When implementing `ToProto` and `FromProto`, these URL-scoped fields should not be assigned to the proto payload. They are resolved and handled by the direct controller reconciler/adapter.

### 3. Enum Roundtripping during Fuzz Testing
Protobuf 3 enum fields with a default value of `MATCHER_UNSPECIFIED` (or 0) are mapped in KRM to strings. During fuzzing/roundtripping, if KRM contains `"MATCHER_UNSPECIFIED"`, we must map it back to `pb.CertificateMapEntry_MATCHER_UNSPECIFIED` and vice-versa. Failing to map `pb.CertificateMapEntry_MATCHER_UNSPECIFIED` in `FromProto` causes the `oneof` field to be nil in the reconstructed proto, leading to fuzz test failures.
