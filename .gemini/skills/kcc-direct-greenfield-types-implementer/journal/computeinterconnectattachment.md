# ComputeInterconnectAttachment Migration Journal

Observations and lessons learned during the migration of the existing `ComputeInterconnectAttachment` resource to direct KRM types:

## 1. Strict Schema Compatibility
- The baseline CRD for `ComputeInterconnectAttachment` did not include standard direct controller fields such as `projectRef` or `location` in its `spec`.
- Instead, the baseline `spec` contains the fields `region` (string) and `routerRef` (referencing `ComputeRouter`).
- To maintain 100% strict backward compatibility, we omitted `projectRef` and `location` from our Go `ComputeInterconnectAttachmentSpec` definition, keeping `region` and `routerRef` exactly as they were in the original TF-based resource configuration. This achieved a completely clean/empty output from `dev/tasks/diff-crds`.

## 2. Handcoding Reference Types without `kind` Field
- In the baseline CRD, the references `routerRef` and `ipsecInternalAddresses` do not include a `kind` field inside their schemas.
- Using the standard `k8sv1alpha1.ResourceRef` introduces the `kind` property, which violates strict compatibility.
- To resolve this, we handcoded local custom reference types (`ComputeInterconnectAttachmentRouterRef` and `ComputeInterconnectAttachmentIpsecInternalAddressesRef`) in `interconnectattachment_types.go` that only define `external`, `name`, and `namespace`.
- The KCC CRD builder `scripts/add-validation-to-crds` automatically detects types with `"external,name,namespace"` signature and appends the standard OpenAPI `oneOf` reference constraints to them automatically.

## 3. Type Mismatch in `mtu` Field
- In KRM, the `mtu` field is specified as a string (`*string`), whereas in the Google Cloud Compute protobuf definition, it is represented as a 32-bit integer (`*int32`).
- To map between these different representations, custom translation logic must be implemented in the direct reconciler adapter.
