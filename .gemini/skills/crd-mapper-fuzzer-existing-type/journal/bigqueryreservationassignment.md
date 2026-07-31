# BigQueryReservationAssignment Journal

## Context
When transitioning/configuring the existing `BigQueryReservationAssignment` resource to support direct controller type generation, we investigated the existing KRM types and restructured the generator scripts.

## Key Learnings & Resolutions
1. **Split generate.sh Pattern**:
   - Previously, the directory had a single top-level `apis/bigqueryreservation/generate.sh`. 
   - In accordance with KCC conventions and standard directory structures, we replaced it with two separate version-specific scripts: `apis/bigqueryreservation/v1alpha1/generate.sh` and `apis/bigqueryreservation/v1beta1/generate.sh`.
   - The `generate-mapper` invocation resides under `v1beta1/generate.sh` with the `--multiversion` flag to correctly handle both versions without collision.

2. **Strict Schema Compatibility**:
   - The generated types were verified to be strictly schema-compatible with the baseline CRD.
   - We ran `dev/tasks/diff-crds`, which returned an empty diff, confirming that no fields or schemas were modified, preserving compatibility perfectly.
   - We ran the full generator pipeline, Go formatting (`make fmt`), type checking (`go vet`), and generated the static configs (`python3 dev/tasks/generate_static_config.py`).
