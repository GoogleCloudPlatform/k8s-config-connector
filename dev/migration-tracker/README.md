# KCC Direct Migration Tracker

This folder tracks the migration of Config Connector (KCC) resources from legacy controllers (`Terraform` or `DCL`) to `Direct` controllers.

The primary status document is **[`MIGRATION_STATUS.md`](./MIGRATION_STATUS.md)**.

## How to View Migration Status

- **Primary (Markdown Dashboard)**: View **[`MIGRATION_STATUS.md`](./MIGRATION_STATUS.md)** directly for live progress metrics, dependency priority rankings, step implementation status, and full resource registries.
- **Web Dashboard**: View the hosted interactive dashboard at **[https://googlecloudplatform.github.io/k8s-config-connector/migration-tracker/#](https://googlecloudplatform.github.io/k8s-config-connector/migration-tracker/#)** (may be outdated). Alternatively, you can serve it locally using `python3 -m http.server 8000` in this directory.

## Scope of Brownfield Migration

The migration tracker specifically targets **v1beta1 Brownfield resources**:
- **Included**: Legacy resources registered in [`pkg/controller/resourceconfig/static_config.go`](../../pkg/controller/resourceconfig/static_config.go) whose `SupportedControllers` includes legacy reconcilers (`Terraform` or `DCL`) and have reached `v1beta1` API maturity.
- **Excluded**:
  - **Pure Greenfield Resources**: Resources built from scratch using the direct controller architecture (where `SupportedControllers` contains only `Direct`).
  - **`v1alpha1`-Only Resources**: Experimental or alpha resources that have not yet been promoted to `v1beta1`.

## Data Generation & Architecture

The migration tracking workflow uses a clean two-stage pipeline:

1. **[`generate_data.py`](./generate_data.py)**: Performs disk scanning across the repository:
   - Scans `apis/<group>/<version>/` for Go API types, strictly checking for standard `<kind_lower>_identity.go` and `<kind_lower>_reference.go` files (while detecting shared/non-standard file placements).
   - Scans `pkg/controller/direct/<group>/` for direct controller adapters (`*_controller.go`), mappers (`*_mapper.go`), and fuzzers (`*_fuzzer.go`).
   - Analyzes CRD schema dependencies in `config/crds/resources/` to calculate topological ordering.
   - Outputs the complete dataset to [`data.json`](./data.json).

2. **[`generate_md.py`](./generate_md.py)**: Pure Markdown generator:
   - Invokes `generate_data.py` to ensure dataset freshness.
   - Reads `data.json` and renders [`MIGRATION_STATUS.md`](./MIGRATION_STATUS.md), including progress metrics, step completion breakdown, topologically sorted priority list, complete resource registry, and architectural gap analysis.

## Files

- [`MIGRATION_STATUS.md`](./MIGRATION_STATUS.md): Primary auto-generated Markdown dashboard visualizing progress metrics, dependency priority rankings, resource registries, and architectural gap analysis.
- [`generate_md.py`](./generate_md.py): Python script to regenerate `MIGRATION_STATUS.md` from `data.json`.
- [`generate_data.py`](./generate_data.py): Python script to scan the codebase and update `data.json`.
- [`data.json`](./data.json): Structured dataset containing in-scope v1beta1 brownfield KCC resources, migration steps, artifact metadata, and dependencies.
- [`list_top_unmigrated.py`](./list_top_unmigrated.py): CLI script to list unmigrated resources sorted by topological dependency priority.
- `index.html`, `app.js`, `style.css`: Optional static web UI dashboard.

## Updating Data & Dashboard

To update resource status data and regenerate the markdown dashboard, run:

```bash
python3 dev/migration-tracker/generate_md.py
```
