# Migration Journal: IAPBrand

**Current Step:** Step 1: Direct API Types

| Step | Name | Issue | Pull Request | Status | Date Started | Date Completed |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | Direct API Types | [Issue #10375](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/10375) | [PR #10379](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/10379) | PR Created | 2026-06-16 | - |
| 2 | Identity and Reference Types Pattern | - | - | - | - | - |
| 3 | Create a Round-Trip KRM Fuzzer | - | - | - | - | - |
| 4 | Implement Direct Controller & E2E Fixtures | - | - | - | - | - |

## Status Updates
- **2026-06-17**: Checked progress. Identified and fixed failing CI checks by regenerating out-of-date reference documentation and updating `missingfields.txt` and `IAPSettings.diff` test exceptions. Added missing labels and configured PR #10379 and PR #10381. Fixes are ready for commit/push. Step 1 remains in progress.
- **2026-06-17**: Checked progress. Pull Request #10379 has failed CI checks (`unit-tests`, `validate-generated-files`, and `validations`). Assigned the PR to `factorybot-robot` to request automated correction.
- **2026-06-17**: Checked progress. Pull Request #10379 has been opened by `lovelace-coder-bot` to address Issue #10375. CI checks were previously failing on `validate-generated-files` and `unit-tests`, but a new CI run has been triggered and is currently in progress.
- **2026-06-16**: Checked progress. No PR has been opened yet. Step 1 remains in progress.
- **2026-06-16**: AI Factory (argus-watcher-bot) started fixing Issue #10375 in a sandbox.
- **2026-06-16**: Started migration for IAPBrand. Opened Issue #10375 to implement direct KRM types and generate.sh for IAPBrand.
- **2026-06-16**: Assigned Issue #10375 to codebot-robot to begin implementation of Step 1.
- **2026-06-16**: Added the `overseer` label to Issue #10375 to trigger the AI Factory (argus-watcher-bot) for Step 1 implementation.
