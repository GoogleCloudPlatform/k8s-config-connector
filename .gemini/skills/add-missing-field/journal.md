# Implementer's Journal

- **Update testdata/exceptions/alpha-missingfields.txt**: When a new resource is implemented, its tests should cover the necessary fields. Consequently, the apichecks will start failing because the resource is still listed in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`. To resolve this, run `UPDATE_GOLDEN_FILES=1 go test ./tests/apichecks/...` or manually remove the resource's missing fields from the exception list. Make sure you rebase with `upstream/master` first to avoid modifying unrelated changes.
