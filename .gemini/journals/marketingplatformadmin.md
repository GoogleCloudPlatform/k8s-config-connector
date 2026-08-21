### 2026-07-02 Implementing MarketingPlatformAdminAnalyticsAccountLink
- **Context**: Greenfield implementation of initial types, CRD, and IdentityV2 for `MarketingPlatformAdminAnalyticsAccountLink` (PR [#11247](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11247)).
- **Problem**: 
  1. The proto package `google.marketingplatform.admin.v1alpha` is not located under `google/cloud/`, so it was not compiled by `generate-proto.sh` by default.
  2. The resource is organization-scoped (`organizations/{organization}/analyticsAccountLinks/{analytics_account_link}`) rather than the standard project-scoped layout, which required mapping `organizationRef` instead of `projectRef`.
  3. The resource references an `AnalyticsAccount` using the format `analyticsadmin.googleapis.com/accounts/{account_id}`.
- **Solution**:
  1. Updated `dev/tools/controllerbuilder/generate-proto.sh` to compile `${THIRD_PARTY}/googleapis/google/marketingplatform/admin/*/*.proto` so that the message metadata is included in the descriptor.
  2. Updated the spec to map `organizationRef` using `refsv1beta1.OrganizationRef`.
  3. Mapped the `analyticsAccount` reference to the canonical `analyticsv1alpha1.AccountRef` type from the analytics package, maintaining clean references.
- **Impact**: Ensures that any future work on marketingplatformadmin (like writing the controller or mockgcp layer) works out of the box because the proto definitions and types are perfectly aligned.
