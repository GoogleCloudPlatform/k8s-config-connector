# ComputeTargetVPNGateway Direct KRM Type Journal

## Observations & Implementation Decisions

1. **Handling ID Type Mismatch (`*uint64` vs `*int64`)**:
   - The GCP Compute API representation of `TargetVpnGateway` utilizes `*uint64` for the `id` field.
   - The KCC `ComputeTargetVPNGateway` KRM schema represents `gatewayId` as an integer (`*int64` in Go).
   - Because the pointer scalar types differ (`*uint64` vs `*int64`), automatic mapping is not generated. We implemented a dedicated hand-coded mapper in `targetvpngateway_mapper.go` to handle type conversion during the `FromProto` and `ToProto` translation.

2. **Ensuring Fuzzer Round-trip Stability for `Region` (`*string` vs `string`)**:
   - The baseline CRD represents `region` as `type: string`.
   - In KRM, we initially scaffolded `Region string`. However, because the proto field is `Region *string`, this introduced round-trip fuzzer mismatches when `Region` was `nil` in the original proto (mapping to `""` in KRM, and back to a pointer to `""` in the output proto).
   - Changing the KRM Go struct field to `Region *string` matches the original OpenAPI schema exactly (it is still `type: string` in the OpenAPI schema), but allows the fuzzer to round-trip `nil` pointers perfectly.
