# Journal: MonitoringServiceLevelObjective Greenfield Types Implementation

## Overview
We have implemented the direct KRM types and reference support for the `MonitoringServiceLevelObjective` resource, promoting it to a Direct-types-capable model in `v1beta1`.

## Key Observations and Implementation Details

1. **Global Scope vs Regional Scope**:
   - GCP's `ServiceLevelObjective` resource is global, representing the SLO of a parent `Service` under `projects/{project}/services/{service}/serviceLevelObjectives/{servicelevelobjective}`.
   - Consequently, the KCC `Location` spec property was omitted, keeping the resource definition aligned with GCP's URL patterns.

2. **Parent Reference (`MonitoringServiceRef`)**:
   - `MonitoringServiceLevelObjective` belongs directly to a parent `MonitoringService`.
   - To follow Config Connector's strict reference guidelines, we fully implemented a proper reference type `MonitoringServiceRef` (implementing `refs.Ref`) and the corresponding parsed identity type `MonitoringServiceIdentity` (implementing `identity.IdentityV2`).
   - We updated `MonitoringServiceLevelObjectiveSpec` to require `serviceRef *MonitoringServiceRef`.

3. **Reachable Go Types Generation**:
   - By specifying nested fields (such as `ServiceLevelIndicator`) directly in `MonitoringServiceLevelObjectiveSpec`, the code generator's `prune-unused-types` pass automatically identified these nested types as reachable.
   - This kept the types in `types.generated.go` uncommented and active, removing the need to manually clone or rewrite massive nested proto-derived structures.

4. **Preserved Schema Compatibility**:
   - The stability level `stable` was carefully preserved for the `v1beta1` resource, preventing stability level regression.
   - The resulting CRD OpenAPI schema matches all previous DCL-based properties perfectly, ensuring backward compatibility for all users.
