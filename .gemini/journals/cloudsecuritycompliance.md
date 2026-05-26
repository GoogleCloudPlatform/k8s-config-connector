### [2026-05-26] Proto Package Discrepancy for CloudSecurityFramework
- **Context**: Implementing Greenfield Step 1 types for `CloudSecurityFramework` in `google.cloud.security.compliance.v1`.
- **Problem**: The proto package was actually named `google.cloud.cloudsecuritycompliance.v1` instead of `google.cloud.security.compliance.v1`. Additionally, the proto was not present in the default `googleapis` SHA and required a bump.
- **Solution**: Updated `generate.sh` to use `google.cloud.cloudsecuritycompliance.v1` and bumped `apis/git.versions` to a newer SHA that included the API.
- **Impact**: Future agents implementing `cloudsecuritycompliance` resources should use the correct proto package and ensure they have a recent enough `googleapis` SHA.
