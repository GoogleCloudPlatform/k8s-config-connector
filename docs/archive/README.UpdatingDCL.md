# [DEPRECATED] Updating the DCL (declarative-resource-client-library)

> [!WARNING]
> **This document has been deprecated and moved.**
> Please refer to the updated playbook: **[docs/develop-resources/scenarios/dcl-field-and-resource.md](../develop-resources/scenarios/dcl-field-and-resource.md)**.
>
> Config Connector is actively prioritizing the [Direct reconciliation approach](../develop-resources/README.md). New resources and fields should be developed as Direct controllers wherever possible.

---

### Archived Content (For Reference Only)

Approximately 25% of Config Connector resources leverage the [DCL](https://github.com/GoogleCloudPlatform/declarative-resource-client-library) library. Periodically, we need to update this dependency to incorporate new fields for existing resources and address bug fixes.

Upgrade Process:

1. In the repository root directory, run `make upgrade-dcl` to update the DCL dependency.
1. Run `make ready-pr` to ensure proper updates for all generated files.
1. Check both the DCL [release notes](https://github.com/GoogleCloudPlatform/declarative-resource-client-library/releases) and the generated files, especially diff in CRDs.
   * Create a Pull Request titled "Update DCL to v1.58.0" (replace version as needed).
   * In the PR description, list the new CRD fields and relevant bug fixes introduced by the DCL update.
