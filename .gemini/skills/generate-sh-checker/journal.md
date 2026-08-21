# Generate.sh Checker Journal

## 2026-06-06: Fully-qualified Proto Names for Beta/Alpha resources

When backfilling `generate.sh` for `apis/compute/v1alpha1`, we found that `ComputeFutureReservation` maps to `google.cloud.compute.v1beta.FutureReservation` because it only exists in GCE's `v1beta` proto package, whereas other compute v1alpha1 KRM resources like `ComputeInterconnect` map to GCE's `v1` package (`google.cloud.compute.v1.Interconnect`).

If we use the default resource mapping in `generate-types` (e.g. `--resource ComputeFutureReservation:FutureReservation`), it assumes the resource is in the `--service` package (`google.cloud.compute.v1`), which fails.

**Solution:**
We can specify the fully-qualified proto message name (including the full package name) in the `--resource` argument:
`--resource ComputeFutureReservation:google.cloud.compute.v1beta.FutureReservation`

Since the proto compilation under `.build/` includes both `v1` and `v1beta` packages for compute, the generator can find the fully-qualified descriptor and successfully generate the types.
