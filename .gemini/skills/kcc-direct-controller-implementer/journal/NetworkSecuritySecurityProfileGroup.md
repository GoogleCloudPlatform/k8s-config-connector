# NetworkSecuritySecurityProfileGroup Direct Controller Journal

## Service Client Disambiguation
During implementation, we encountered a collision in `mockgcp`'s HTTP-to-gRPC gateway/mux when mapping `*networksecuritypb.securityProfileGroupServiceClient`. Both `SecurityProfileGroupService` and `OrganizationSecurityProfileGroupService` expose identical methods, causing the reflex matcher to throw a "found multiple matching service" fatal error.
We resolved this by enhancing the service matching logic in `mockgcp/common/httptogrpc/service.go`. When multiple matches are found, we filter them using a substring check: we check if the lowercase version of the service descriptor name is a substring of the lowercase Go client's name. This cleanly disambiguated `SecurityProfileGroupService` from `OrganizationSecurityProfileGroupService` without breaking any existing services.

## External References
Because `NetworkSecuritySecurityProfile` itself is not yet implemented as a Config Connector controller, we designed our maximal E2E test fixture to use `external` references (e.g., `projects/${projectId}/locations/global/securityProfiles/dummy-threat-${uniqueId}`) to point to a simulated/dummy external security profile. This allowed us to successfully test and verify reference fields and their resolution within `NetworkSecuritySecurityProfileGroupSpec` without requiring a parent/dependent KCC-managed `NetworkSecuritySecurityProfile` resource.
