# DNSResponsePolicyRule Direct KRM Type & generate.sh Migration

## Overview
DNSResponsePolicyRule is a `v1alpha1` resource that has been transitioned to support direct KRM types and automatic mapper generation.

## Key Decisions

### 1. `generate.sh` Configuration
Since the DNS service API definitions are managed using Google Cloud DNS OpenAPI/Discovery specifications (`dns-api.json`) instead of gRPC/Protobuf protocols, we utilize the `openapi-to-krm` and `openapi-to-krm/cmd/generate-mapper` tools. We created a local `generate.sh` file under `apis/dns/v1alpha1/generate.sh` to:
- Generate OpenAPI types into `apis/dns/v1alpha1/types.generated.go`.
- Automatically generate mapper files to `pkg/controller/direct/dns/zz_generated.v1alpha1.mappers.go`.

### 2. Strict Schema Compatibility
To preserve perfect schema compatibility with the baseline CRD, we defined the KRM Go structures in `apis/dns/v1alpha1/dnsresponsepolicyrule_types.go` matching the original client-go representations exactly:
- `localData.localDatas` uses a simplified representation of `ResourceRecordSet` with only four fields: `name`, `rrdatas`, `ttl`, and `type`.
- We used standard annotations like `// +openapi:ResponsePolicyRule`, `// +openapi:ResponsePolicyRuleLocalData`, and `// +openapi:ResourceRecordSet` on the structs to map to their OpenAPI components.
- We ran `dev/tasks/diff-crds` and confirmed that there are **absolutely zero schema changes or deviations** compared to the baseline CRD.

### 3. Fuzzer Setup
We registered a new type-safe fuzzer in `pkg/controller/direct/dns/dnsresponsepolicyrule_fuzzer.go` using `fuzztesting.RegisterKRMFuzzer_NoProto(...)` which successfully validated round-trip capability under extensive randomized testing.
