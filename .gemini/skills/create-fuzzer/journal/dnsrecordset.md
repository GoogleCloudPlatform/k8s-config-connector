# Fuzzer Journal: DNSRecordSet (NoProto Fuzzing of Nested Slices)

This journal captures learnings from implementing the round-trip KRM fuzzer for `DNSRecordSet` (`api.ResourceRecordSet`).

## Key Observations and Gotchas

### 1. Traversal Path Reset in Slices
In KCC's `ClearNonProtoFields` and `RandomFiller` (`pkg/test/fuzz/krmgen.go`), slice element traversal completely resets the tracking path:
```go
	case reflect.Slice:
		for j := 0; j < field.Len(); j++ {
			rf.fillWithClear(t, "", field.Index(j))
		}
```
Passing `""` as the `fieldName` parameter resets the path inside slice items. Consequently, any nested fields of slice elements must be declared relative to the containing slice element itself.

**Example:**
For a nested field path like `.RoutingPolicy.Geo.Items[].HealthCheckedTargets.ExternalEndpoints`, the fuzzer lookup actually matches:
- `".HealthCheckedTargets.ExternalEndpoints"`

Declaring the fully qualified path (`.RoutingPolicy.Geo.Items[].HealthCheckedTargets.ExternalEndpoints`) will **not** match or zero out the field, causing round-trip comparison failures.

### 2. Conflicting Field Names under Slices
If a nested field inside a slice element has the same name as a top-level mapped field (for example, `.Rrdatas`), they both resolve to `".Rrdatas"`.
- If you mark `".Rrdatas"` as unimplemented, the fuzzer zeroes it out at the top level too, losing validation of your mapped top-level field.
- If you do not mark it, the fuzzer randomizes the nested field under the slice, which isn't mapped, causing a mismatch during diff comparison.

### 3. Solution: Using Custom `FilterSpec` / `FilterStatus`
The most robust and elegant way to resolve conflicting slice element fields is to hook into `f.FilterSpec` and `f.FilterStatus` to clear the unmapped nested fields post-randomization, while preserving the top-level spec fields:

```go
	filter := func(in *api.ResourceRecordSet) {
		if in.RoutingPolicy != nil {
			if in.RoutingPolicy.Geo != nil {
				for _, item := range in.RoutingPolicy.Geo.Items {
					item.Rrdatas = nil
					item.SignatureRrdatas = nil
				}
			}
			if in.RoutingPolicy.Wrr != nil {
				for _, item := range in.RoutingPolicy.Wrr.Items {
					item.Rrdatas = nil
					item.SignatureRrdatas = nil
				}
			}
			if in.RoutingPolicy.PrimaryBackup != nil {
				if in.RoutingPolicy.PrimaryBackup.BackupGeoTargets != nil {
					for _, item := range in.RoutingPolicy.PrimaryBackup.BackupGeoTargets.Items {
						item.Rrdatas = nil
						item.SignatureRrdatas = nil
					}
				}
			}
		}
	}
	f.FilterSpec = filter
	f.FilterStatus = filter
```
This keeps the fuzzer declarations extremely clean and ensures 100% round-trip accuracy.
