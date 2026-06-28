# Journal: KRM Fuzzer for RedisInstance

## Problem: 1-to-Many Nil vs Empty Struct Mapping Dilemma
During fuzz testing of `RedisInstance`, we encountered a round-trip failure on `weekly_maintenance_window[].start_time`.

The KRM struct defines `StartTime` as an embedded struct, not a pointer:
```go
type InstanceWeeklyMaintenanceWindow struct {
	StartTime TimeOfDay `json:"startTime"`
}
```
Because `StartTime` is not a pointer, there is no way in KRM to distinguish between a proto `start_time` of `nil` (unset) vs an empty `start_time: {}` (all subfields unset). Both map to the same zero-value KRM `StartTime{}` struct.

When mapping from KRM back to proto:
1. If the standard `ToProto` mapper is used, it calls `TimeOfDay_ToProto`, which populates `p2` with an empty but non-nil proto message `start_time: {}`.
2. This creates a mismatch if the original proto `p1` was `nil`.
3. If we change `ToProto` to only emit `start_time` if any of its subfields are set (leaving it as `nil` otherwise), then `p2` gets `nil`. But if the original `p1` had `{}`, we get a mismatch where `p1` is `{}` and `p2` is `nil`.

## Solution: FilterSpec and Custom ToProto Check
To solve this robustly:
1. We implement a custom check in the `ToProto` mapper (`InstanceWeeklyMaintenanceWindow_ToProto`) to only convert and populate `StartTime` in the proto if at least one subfield is non-nil/non-zero:
   ```go
   if in.StartTime.Hours != nil || in.StartTime.Minutes != nil || in.StartTime.Seconds != nil || in.StartTime.Nanos != nil {
       out.StartTime = TimeOfDay_ToProto(mapCtx, &in.StartTime)
   }
   ```
2. In the fuzzer configuration, we do NOT mark the field as unimplemented because it is actually fully supported.
3. Instead, we use `f.FilterSpec` to normalize the original proto object `p1` before mapping. If `p1` has an empty `StartTime` (`hours == 0 && minutes == 0 && seconds == 0 && nanos == 0`), we set it to `nil` in `p1`:
   ```go
   f.FilterSpec = func(in *redispb.Instance) {
       if in.MaintenancePolicy != nil {
           for _, w := range in.MaintenancePolicy.WeeklyMaintenanceWindow {
               if w != nil && w.StartTime != nil {
                   if w.StartTime.Hours == 0 && w.StartTime.Minutes == 0 && w.StartTime.Seconds == 0 && w.StartTime.Nanos == 0 {
                       w.StartTime = nil
                   }
               }
           }
       }
   }
   ```
4. This ensures that both `p1` and `p2` represent `nil` for empty start times, and round-tripping works seamlessly for all random inputs.
