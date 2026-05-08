# Implementer's Journal

- **Update testdata/exceptions/alpha-missingfields.txt**: When a new resource is implemented, its tests should cover the necessary fields. Consequently, the apichecks will start failing because the resource is still listed in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`. To resolve this, run `UPDATE_GOLDEN_FILES=1 go test ./tests/apichecks/...` or manually remove the resource's missing fields from the exception list. Make sure you rebase with `upstream/master` first to avoid modifying unrelated changes.

# Observations: add-missing-field

1. **WRITE_GOLDEN_OUTPUT=1**: When running `hack/compare-mock`, you must set `WRITE_GOLDEN_OUTPUT=1` to update the golden files (like `_http.log` and `_generated_object...golden.yaml`). The test will fail even if it successfully updates the files (because it detects a diff from the previous run). This is expected.
2. **Directory Naming**: Test directories under `pkg/test/resourcefixture/testdata/basic/<group>/<version>/` might not match the exact resource name `kind`. For example, for `MemorystoreInstance`, the directory is `memorystoreinstance` instead of `instance`.
3. **Mappers Symmetrical Rule**: Even for `ObservedState` (output-only) fields, if there is a manual `_ToProto` function (e.g., `MemorystoreInstanceObservedState_ToProto`), you must add the field there as well so that the types continue to round-trip properly inside the controller's internal logic.
4. **Pointer Assignments**: For optional proto fields (like proto3 `optional` string), the generated Go proto type might be a `*string`. You can assign it directly to a KRM `*string` field without needing `direct.ValueOf` or `direct.LazyPtr`.
