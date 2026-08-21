---
name: crd-mcp-server
description: Provides instructions and a tool for analyzing Custom Resource Definition (CRD) changes, including checking for KRM API structural equivalence and backward compatibility.
---

# CRD Analyzer

## Overview

The `crd-mcp-server` is a binary tool under `dev/tools/crd-mcp-server/` designed to check and validate Custom Resource Definition (CRD) schema changes in the Config Connector repository. It has two main subcommands:

1. **`equivalent`**: Checks whether a CRD change is equivalent to its previous git-committed version (no fields added or removed under spec, no type or CEL validation changes, etc.).
2. **`compatible`**: Checks whether a CRD change is backward compatible with its previous git-committed version (no fields removed or renamed, no incompatible type changes, allowed additions).

This tool is used in presubmits (e.g., `dev/ci/presubmits/crd-equivalence-check`) to ensure that KRM schemas remain compatible or equivalent during migration from legacy to direct reconcilers.

## How to Run

Before running, ensure you run from the repository root. You can build and run the tool directly using `go run`:

### Checking for Equivalence (e.g., during promotion or refactoring)

To check if a changed CRD is equivalent to its version on `master`:

```bash
go run ./dev/tools/crd-mcp-server equivalent --file config/crds/resources/services/v1beta1/service_kind.yaml --ref master
```

### Checking for Backward Compatibility

To check if a changed CRD is backward-compatible with `master`:

```bash
go run ./dev/tools/crd-mcp-server compatible --file config/crds/resources/services/v1beta1/service_kind.yaml --ref master
```
