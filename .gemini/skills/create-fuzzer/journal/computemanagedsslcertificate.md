# ComputeManagedSSLCertificate Fuzzer Journal

## Implementation & Alignment
- **Fuzzer Location:** `pkg/controller/direct/compute/computemanagedsslcertificate_fuzzer.go`
- **Registered Function:** `computeManagedSSLCertificateFuzzer` in `init()` via `fuzztesting.RegisterKRMFuzzer`
- **Package Registration:** `pkg/controller/direct/compute` is imported anonymously in `pkg/controller/direct/register/register.go`.

## Proto & KRM Mapping Details
`ComputeManagedSSLCertificate` maps to the protobuf message `google.cloud.compute.v1.SslCertificate` (`cloud.google.com/go/compute/apiv1/computepb`).
Unlike `ComputeSSLCertificate` (which handles self-managed certificates with certificate and private key strings), `ComputeManagedSSLCertificate` manages Google-managed SSL certificates via the nested `.managed` proto message.

### Spec Fields
- `description` maps to proto `.description` (`f.SpecField(".description")`)
- `managed` maps to proto `.managed` (`f.SpecField(".managed")`)
- `managed.domains` maps to proto `.managed.domains` (`f.SpecField(".managed.domains")`)
- `type` maps to proto `.type` (`f.SpecField(".type")`)

### Status Fields (ObservedState)
- `certificateID` maps to proto `.id` (`f.StatusField(".id")`, handcoded in mappers because OpenAPI integer without format maps to `*int` in KRM vs `*uint64` in protobuf)
- `creationTimestamp` maps to proto `.creation_timestamp` (`f.StatusField(".creation_timestamp")`)
- `expireTime` maps to proto `.expire_time` (`f.StatusField(".expire_time")`)
- `selfLink` maps to proto `.self_link` (`f.StatusField(".self_link")`)
- `subjectAlternativeNames` maps to proto `.subject_alternative_names` (`f.StatusField(".subject_alternative_names")`)

### Identity & Special Fields
- `resourceID` / `.name` maps to Identity (`f.Unimplemented_Identity(".name")`)

### Unimplemented / Not Yet Triaged Fields
- `.certificate`: Not in KRM (for self-managed certificates)
- `.kind`: Internal GCP resource type indicator
- `.managed.domain_status`: Map of domain statuses not yet triaged in KRM
- `.managed.status`: Certificate provisioning status not yet triaged in KRM
- `.private_key`: Not in KRM (for self-managed certificates)
- `.region`: Regional certificate field
- `.self_managed`: Not in KRM (for self-managed certificates)

## Verification
1. Focused fuzz test suite:
   ```bash
   FOCUS=ComputeManagedSSLCertificate go test -count=100 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers
   ```
   Both Spec and Status round-trip tests passed 100 iterations.
2. Verified that `ComputeManagedSSLCertificate` is no longer listed in `TestListTypesWithoutFuzzers`:
   ```bash
   go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestListTypesWithoutFuzzers
   ```
