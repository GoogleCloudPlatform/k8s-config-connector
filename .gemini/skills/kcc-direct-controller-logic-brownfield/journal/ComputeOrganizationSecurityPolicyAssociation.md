# ComputeOrganizationSecurityPolicyAssociation Direct Controller Migration Journal

## Overview
This document records learnings and implementation details from migrating `ComputeOrganizationSecurityPolicyAssociation` from Terraform to the direct controller.

## Learnings & Observations

### 1. API Architecture & Identity
- `ComputeOrganizationSecurityPolicyAssociation` represents an association between an organization security policy and a resource (typically an organization or folder).
- The REST API endpoint uses methods on the `OrganizationSecurityPolicies` service:
  - Create: `AddAssociation` (POST `locations/global/securityPolicies/{security_policy}/addAssociation`)
  - Read: `GetAssociation` (GET `locations/global/securityPolicies/{security_policy}/getAssociation?name={name}`)
  - Delete: `RemoveAssociation` (POST `locations/global/securityPolicies/{security_policy}/removeAssociation?name={name}`)
- In GCP Compute API, querying a non-existent association via `GetAssociation` returns an HTTP 400 with the message `Invalid value for field 'name': '<name>'. An association with that name does not exist.` instead of an HTTP 404.
- Handled this error in `isAssociationNotFound` by matching `gerr.Code == 400` and `strings.Contains(gerr.Message, "An association with that name does not exist")`, treating it identically to a 404 Not Found error in `Find` and `Delete`.

### 2. Immutability
- Associations in GCP cannot be updated in-place (the attachment ID and policy ID cannot be modified after creation).
- In `Update`, top-level field diffing is performed; if any diff is detected, the adapter returns an error stating that the resource is immutable.

### 3. Dual-Controller Support and Dynamic Testing
- Added `k8s.ReconcilerTypeDirect` to `SupportedControllers` for `ComputeOrganizationSecurityPolicyAssociation` in `pkg/controller/resourceconfig/static_config.go`.
- Ran unified tests against MockGCP, verifying both legacy controller and direct controller creation, reconciliation, and deletion flows.
