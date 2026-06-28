# Export Support Journal for KMSCryptoKey

## Observations
- Implementing `AdapterForURL` for `KMSCryptoKey` was highly straightforward because `KMSCryptoKeyIdentity` already had `FromExternal()` implemented and natively supports parsing full GCP / CAI URLs.
- The `Export()` method was already implemented on `kmsCryptoKeyAdapter`, but lacked setting the `resourceID` field on `Spec`, and did not invoke `export.SetProjectID` or `export.SetLabels`. Adding these ensures proper metadata/field alignment on the exported manifest.
- Since standard `envtest` binaries were not pre-configured, we used `setup-envtest` to locate/download local control plane assets, and coupled with `ONLY_TEST_APIGROUPS=kms.cnrm.cloud.google.com`, we ran the targeted `kmscryptokey` test fixture in ~25 seconds.
- Running E2E tests with `WRITE_GOLDEN_OUTPUT=1` against MockGCP rewrites the dynamic mock HTTP logs/diffs, so discarding changes to `_http.log` / `_http_old_controller.log` and committing only `_exported.yaml` preserves the canonical real GCP baseline.
