# BigQueryDatasetAccess Direct Controller Journal

- **Resource Kind:** `BigQueryDatasetAccess`
- **GCP Service:** `bigquery.googleapis.com` (`google.golang.org/api/bigquery/v2`)
- **Package Path:** `pkg/controller/direct/bigquery`

## Implementation Details
1. **Client Setup:**
   - Used `api.NewService(ctx, opts...)` with `m.config.RESTClientOptions()` to access `gcpService.Datasets`.

2. **Matching & Immutability:**
   - Access entries are matched against `dataset.Access` entries by role (with basic role canonicalization) and member type.
   - Handled `IamMember` conversion: if GCP translated `iamMember` to `userByEmail`, `groupByEmail`, `domain`, or `specialGroup`, it is identified as a match and `apiUpdatedMember` is recorded on the status.
   - Immutability check implemented in `Update()`.

3. **Validation:**
   - Fuzzer passes: `FOCUS=BigQueryDatasetAccess go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`
   - Real GCP logs recorded with `./hack/record-gcp`
   - Mock GCP tests pass with `./hack/compare-mock`
   - Golden log alignment passes: `go test -v ./pkg/test/resourcefixture -run TestGoldenLogAlignment/bigquery/v1alpha1/bigquerydatasetaccess`
   - API checks pass: `go test -v ./tests/apichecks/...`
