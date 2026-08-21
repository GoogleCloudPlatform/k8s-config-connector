### 2026-06-12 Implement DNSResponsePolicy direct controller and MockGCP support
- **Context**: Implementing direct controller and E2E fixtures for `DNSResponsePolicy` ([Issue #9772](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9772)).
- **Problem**: `DNSResponsePolicy` was not yet supported by MockGCP. We needed to:
  1. Add MockGCP endpoints for `dns.responsePolicies` in `mockgcp/mockdns/responsepolicy.go` with `Create`, `Get`, `Update`, `Patch`, `Delete`, and `List` methods.
  2. Map these methods and register them in `mockgcp/mockdns/service.go`.
  3. Register `DNSResponsePolicy` in `config/tests/samples/create/harness.go` under `dns.cnrm.cloud.google.com`.
  4. Ensure that volatile auto-generated ID strings from the mock API response were properly normalized in `mockgcp/mockdns/normalize.go` for the kind `dns#responsePolicy`.
- **Solution**:
  1. Implemented a fully functional `responsePoliciesService` mock in `mockgcp/mockdns/responsepolicy.go` modeling the exact JSON-based Google SDK API structure.
  2. Registered the service with both gRPC and HTTP routing inside `mockgcp/mockdns/service.go`.
  3. Added the `dns#responsePolicy` normalization rule to map random policy IDs to `${dnsResponsePolicyId}` within `mockgcp/mockdns/normalize.go`.
- **Impact**: Enables complete hermetic E2E testing of the `DNSResponsePolicy` direct controller using the newly captured and recorded `dnsresponsepolicy-minimal` fixture.
