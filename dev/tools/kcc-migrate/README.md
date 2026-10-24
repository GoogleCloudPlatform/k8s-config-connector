# `kcc-migrate`: Zero-Downtime Migration CLI for Config Connector

`kcc-migrate` is an enterprise-grade migration automation CLI designed to transition Google Cloud infrastructure-as-code from Google-Managed Config Controller (ACP) to Self-Managed Standalone Config Connector on GKE Standard or Autopilot without workload downtime or cloud resource recreation.

## Core Capabilities

1. **`export`**: Discovers all active KCC CRDs and concurrently exports live resources across specified namespaces using a configurable worker pool (20 workers default).
2. **`sanitize`**: Strips runtime-managed metadata (`uid`, `resourceVersion`, `generation`, `managedFields`, `status`, and source `finalizers`) while automatically injecting required safety invariants:
   - `cnrm.cloud.google.com/deletion-policy: abandon`
   - `cnrm.cloud.google.com/management-conflict-prevention-policy: none`
   - `cnrm.cloud.google.com/state-into-spec: absent`
3. **`lint`**: Performs static linting against target KCC schema standards, enforcing safety annotations and failing fast on unmanaged runtime fields.
4. **`adopt`**: Executes server-side dry-run validation (`--dry-run=true`) followed by live phased adoption into the target Standalone cluster, tracking resource status until `Ready: True (UpToDate)`.
5. **`detach`**: Safely clears finalizers from the source Config Controller cluster, strictly filtering out operator CRDs to prevent controller pod termination while workload finalizers are resolving.

## Quick Start

```bash
# 1. Export live resources concurrently from source cluster
kcc-migrate export --context=source-cc-ctx --namespaces=tenant-prod,tenant-shared --output-dir=./raw

# 2. Sanitize manifests and inject safety invariants
kcc-migrate sanitize --input-dir=./raw --output-dir=./sanitized

# 3. Strictly validate manifests prior to mutation
kcc-migrate lint --input-dir=./sanitized --strict=true

# 4. Dry-run and adopt into target Standalone KCC cluster
kcc-migrate adopt --context=target-kcc-ctx --input-dir=./sanitized --dry-run=true
kcc-migrate adopt --context=target-kcc-ctx --input-dir=./sanitized --dry-run=false

# 5. Non-destructively detach from legacy source cluster
kcc-migrate detach --context=source-cc-ctx --namespaces=tenant-prod,tenant-shared
```
