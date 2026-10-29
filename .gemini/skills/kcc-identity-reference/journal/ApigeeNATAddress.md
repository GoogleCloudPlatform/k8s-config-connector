### [2026-10-21] ApigeeNATAddress Identity & Reference
- **Context**: Implementing `identity.IdentityV2` and `refs.Ref` for `ApigeeNATAddress` following the canonical `gcpurls.Template` pattern.
- **Observations & Decisions**:
  1. **Schema Preservation with String-typed Parent (`spec.instanceId`)**:
     - The `ApigeeNATAddressSpec` defines `instanceId string` (referencing an Apigee instance in the format `organizations/{{org_name}}/instances/{{instance_name}}`) rather than an `InstanceRef` struct.
     - Per the critical rule to maintain zero schema diff against existing CRDs, we preserved `spec.instanceId`.
     - In `getIdentityFromApigeeNATAddressSpec`, we handle both external instance paths (e.g., `organizations/{org}/instances/{instance}`, full URIs with scheme/host, or `{org}/instances/{instance}`) and local KRM `ApigeeInstance` names (resolved via `client.Reader` and `ApigeeInstanceRef.Normalize`).
  2. **URL Template & Discovery Alignment**:
     - URL template: `var ApigeeNATAddressIdentityFormat = gcpurls.Template[ApigeeNATAddressIdentity]("apigee.googleapis.com", "organizations/{organization}/instances/{instance}/natAddresses/{natAddress}")`.
     - Although absent from `cloudassetinventory_names.jsonl`, `discovery_names.jsonl` defines `//apigee.googleapis.com/organizations/{{organizationsId}}/instances/{{instancesId}}/natAddresses/{{natAddressesId}}`, which matches this canonical template. `TestRegisteredTemplatesMatchCAI` in `pkg/gcpurls` passes automatically without requiring an ignore entry.
  3. **Authoritative Status Cross-Check in `GetIdentity`**:
     - `ApigeeNATAddressStatus` contains no `externalRef` or `name` field. Following the skill rules, no status cross-check is performed and the spec identity is returned directly.
  4. **Legacy Fallback in `ApigeeNATAddressRef.Normalize`**:
     - As `ApigeeNATAddress` is managed by the Terraform controller and lacks `status.externalRef`, `ApigeeNATAddressRef.Normalize` uses `refs.NormalizeWithFallback`.
     - The fallback inspects `status.conditions` for `Ready == True`. If not ready, it returns empty string to trigger `k8s.NewReferenceNotReadyError`. If ready, it constructs the identity using `getIdentityFromApigeeNATAddressSpec`.
