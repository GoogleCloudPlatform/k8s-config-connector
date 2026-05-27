# Update GCP Dependencies Skill

This skill guides the process of updating GCP dependencies and handling associated complexities.

## Workflow

1.  **Run the update script:** Execute `dev/tasks/update-gcp-dependencies`.
2.  **Handle compilation errors:** If `controller-gen` or `go build` fails after updates, it often means that manually managed types in `apis/` are out of sync with newly generated types in `types.generated.go`.
3.  **Identify conflicting types:** Look for types in `types.generated.go` that lack the `ObservedState` suffix but are used in fields that should be observed state. This usually happens when a parent type is manually managed and missing new fields.
4.  **Update manually managed types:** Add the missing fields to the manually managed types (e.g., `connectivitytest_types.go`) and define any new nested types with the appropriate `ObservedState` suffix.
5.  **Re-run generation:** Run the `generate.sh` script for the affected API.
6.  **Verify CRDs:** Run `dev/tasks/generate-crds` to ensure all CRDs are valid.

## Common Issues

*   **Script bugs:** Ensure `dev/tasks/update-gcp-dependencies` writes `git.versions` to `apis/git.versions`, not the root directory.
*   **Missing Docker:** If `make ready-pr` fails due to missing Docker, use `go work sync` and `dev/tasks/run-golangci-lint` (if available) as alternatives for local validation.
