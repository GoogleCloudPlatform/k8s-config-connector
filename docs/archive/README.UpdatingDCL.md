# Updating the DCL (declarative-resource-client-library)

Approximately 25% of Config Connector resources leverage the [DCL](https://github.com/GoogleCloudPlatform/declarative-resource-client-library) library. Periodically, we need to update this dependency to incorporate new fields for existing resources and address bug fixes.

Upgrade Process:

1. In the repository root directory, run `make upgrade-dcl` to update the DCL dependency.
1. Run `make ready-pr` to ensure proper updates for all generated files.
1. Check both the DCL [release notes](https://github.com/GoogleCloudPlatform/declarative-resource-client-library/releases) and the generated files, especially diff in CRDs.
   * Create a Pull Request titled "Update DCL to v1.58.0" (replace version as needed).
   * In the PR description, list the new CRD fields and relevant bug fixes introduced by the DCL update.
