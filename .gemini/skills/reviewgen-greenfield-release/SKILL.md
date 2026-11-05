---
name: reviewgen-greenfield-release
description: Review criteria for the final KCC Greenfield Release PR, which removes a kind from BLOCKED_KINDS so its CRD ships in the release bundle.
---

# Review Guide for KCC Greenfield Release PRs

Use this guide for a `Greenfield: Release <Kind> (remove it from BLOCKED_KINDS)` PR. It is the last step of the greenfield pipeline: Step 5 of `.agents/workflows/kcc-greenfield.txt`, or the `finally` step of the `.agents/graphs/kcc-greenfield*.yaml` graphs.

## Checks

*   **Only the blocklist changed**: The PR changes only `dev/tasks/build-release-bundle`, removing exactly `"<Kind>"` from `BLOCKED_KINDS`. No other entry moved or was removed.
*   **Fixture evidence**: Every fixture of `<Kind>` under `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind_lower>/` has a real-GCP `_http.log`, an `_audit_probe.log` and an `_http_mock.log`. If any is missing, the kind isn't ready: fail the PR.
*   **Release note**: `New Alpha Resources (Direct Reconciler): <Kind>` (the kind in backticks).

# Review Comment Template

```markdown
### KCC Greenfield Release Auto-Review Results
* **Only `BLOCKED_KINDS` Changed**: [Pass/Fail]
* **Fixture Evidence (`_http.log`, `_audit_probe.log`, `_http_mock.log`)**: [Pass/Fail]
* **Release Note**: [Pass/Fail]

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
