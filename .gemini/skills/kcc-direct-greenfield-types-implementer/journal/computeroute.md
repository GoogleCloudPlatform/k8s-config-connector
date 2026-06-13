# ComputeRoute Direct Greenfield Types Journal

## Observations & Learnings

1. **Global vs Regional/Zonal Resource Scope**:
   - `ComputeRoute` is a global VPC-scoped resource in GCP. The automatic type generator scaffolded a `Location` field because it is a common pattern, but this is irrelevant for global resources. We manually removed `Location` from `ComputeRouteSpec`.

2. **GCP Unsigned Integer Field Types (`uint32`)**:
   - The protobuf field `Route.priority` is of type `uint32`. While Kubernetes integers are often mapped to `*int64` in Go, declaring `Priority *int64` in KRM types caused a type mismatch compile-time error in the generated mapper.
   - Changing the Go field type to `*uint32` resolved the compiler issue and matched the API definition perfectly.

3. **Legacy Reconciler Labels (`tf2crd`)**:
   - To prevent breaking existing controller routing during this transitional types-only implementation phase, we explicitly retained the `cnrm.cloud.google.com/tf2crd=true` label on the `ComputeRoute` struct. This ensures reconciliation continues to route to the Terraform-based legacy controller.

4. **Reference Implementation for Legacy Resources**:
   - `ComputeRoute` references `ComputeVPNTunnel` via `nextHopVPNTunnelRef`. Since `ComputeVPNTunnel` does not have a direct controller or reference types implemented yet, we created `computevpntunnel_reference.go` and `computevpntunnel_identity.go` in the `v1beta1` package to provide standard reference parsing capabilities using `gcpurls.Template` and `refs.NormalizeWithFallback`. This keeps the API checks clean and avoids adding to the exception list.
