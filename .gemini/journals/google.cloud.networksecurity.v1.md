## NetworkSecurityInterceptEndpointGroupAssociation

*   **API Version:** `networksecurity.cnrm.cloud.google.com/v1alpha1`
*   **Proto Quirks:**
    *   The proto `intercept.proto` containing `InterceptEndpointGroupAssociation` was missing from the pinned `googleapis` SHA (`731d7f2ab6e4e2ea15030c95039e2cb66174d4fb`). It was extracted from the `master` branch and placed into `mockgcp/apis/google/cloud/networksecurity/v1/intercept.proto` to allow generation.
    *   Added `dev/tools/controllerbuilder/generate-proto.sh` entry for `mockgcp/apis/google/cloud/networksecurity/*/*.proto` to compile it.
*   **Reference Quirks:**
    *   Created `ComputeNetworkRef` in `apis/refs/v1beta1/computerefs.go` for the `networkRef` field.
    *   Created `NetworkSecurityInterceptEndpointGroupRef` in `apis/refs/v1beta1/networksecurity_refs.go` for the `interceptEndpointGroupRef` field.

