# ApigeeNATAddress Migration Journal

## Overview
Migrated `ApigeeNATAddress` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **API Client & Endpoints:**
   - Uses `google.golang.org/api/apigee/v1` (`OrganizationsInstancesNatAddressesService`).
   - Added `instancesNatAddressesClient()` accessor to `gcpClient` in `pkg/controller/direct/apigee/client.go`.
   - The GCP API supports `Create`, `Delete`, and `Get` operations, which return `GoogleLongrunningOperation` for async creation and deletion.
   - Used `WaitForApigeeOp` for long-running operations.

2. **Resource Immutability:**
   - In GCP Apigee API, NAT addresses are immutable once created (only `Create`, `Delete`, and `Get` methods are supported).
   - In the KRM spec, `instanceId` and `resourceID` are immutable.
   - The `Update` method updates KRM status (`ipAddress`, `state`) without triggering unsupported API mutations.

3. **Status Fields:**
   - Maps `ipAddress` and `state` (`RESERVED`, `ACTIVE`, etc.) from `GoogleCloudApigeeV1NatAddress` to `ApigeeNATAddressStatus`.

4. **Brownfield Dual-Testing Alignment:**
   - Updated `pkg/controller/resourceconfig/static_config.go` to add `k8s.ReconcilerTypeDirect` to `SupportedControllers` while retaining `DefaultController: k8s.ReconcilerTypeTerraform`.
   - Verified that both legacy and direct reconcilers pass mock testing with clean golden alignment and 0-write steady-state re-reconciliation.
