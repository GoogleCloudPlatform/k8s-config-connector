# Journal - ComputeRouterInterface Direct Types Scaffolding

## Observations and Learnings

- **Hierarchical/Nested Resource Representation:**
  In GCP Compute API, a Router Interface is represented as a nested structure inside `Router` (`routers.interfaces`) rather than as a top-level standalone resource.
  In Config Connector, it is exposed as a separate `ComputeRouterInterface` CRD. To bridge this architectural model, we require the parent `RouterRef` in `ComputeRouterInterfaceSpec` to establish ownership and identify the parent `ComputeRouter`.

- **Unmigrated Parent/Target Resources References:**
  Several resources referenced by `ComputeRouterInterface` (e.g., `ComputeRouter`, `ComputeVPNTunnel`, and `ComputeInterconnectAttachment`) have not been migrated to direct controllers yet. They lacked standard Go reference and identity structures (unlike `ComputeSubnetwork` which had `ComputeSubnetworkRef`).
  We cleanly solved this by introducing localized references and identities under `apis/compute/v1alpha1/` (e.g., `computerouter_reference.go`, `computerouter_identity.go`, etc.) and registered them with the `refs` package. This allows `ComputeRouterInterface` to resolve and parse those resources smoothly.

- **Status & Observed State:**
  For `ComputeRouterInterfaceObservedState`, we mapped `managementType` to the corresponding `management_type` field in the GCP `RouterInterface` proto, keeping all other specs and references strictly aligned with direct guidelines.
