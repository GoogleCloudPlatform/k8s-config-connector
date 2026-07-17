# Fuzzer Journal: IAPBrand

## Observations
- `IAPBrand` contains a `Spec` (`IAPBrandSpec` mapping to `pb.Brand`) and a read-only Status field `OrgInternalOnly` (`IAPBrandStatus` mapping to `pb.Brand.org_internal_only`).
- We implemented custom Status `FromProto` and `ToProto` mappings in `iapbrand_mappings.go` to successfully round-trip the status.
- We registered `IAPBrand` with the full `NewKRMTypedFuzzer` and `RegisterKRMFuzzer` in `iapbrand_fuzzer.go` to validate both Spec and Status fields.
- Registered `.support_email` and `.application_title` as Spec fields, and `.org_internal_only` as a Status field.
- Registered `.name` as an identity field using `f.Unimplemented_Identity`.
- All focused and central fuzz tests compiled and passed successfully on the first try.
