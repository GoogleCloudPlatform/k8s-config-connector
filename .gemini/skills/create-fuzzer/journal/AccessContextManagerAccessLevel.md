# Fuzzer Journal: AccessContextManagerAccessLevel

## Learnings & Observations

### 1. Handling Struct Value Fields in KRM (e.g., Expr)
In KCC's KRM Spec, some fields are defined as non-pointer struct values. For example:
```go
type AccessLevelCustom struct {
	Expr AccessLevelExpr `json:"expr"`
}
```
Because `Expr` is a non-pointer struct, the `FromProto` converter will map a `nil` proto `Expr` to a zero-valued `AccessLevelExpr` in KRM. During the `ToProto` phase, the `ToProto` converter will output a non-nil, empty proto `Expr` struct (`&expr.Expr{}`). This difference results in a roundtrip mismatch (`nil` vs. `{}`).

**Solution:**
Use the `FilterSpec` function to normalize the generated random proto before roundtripping, ensuring that if the parent field (e.g., `custom`) is present, its non-pointer KRM child fields (e.g., `Expr`) are pre-populated so they match after the round-trip conversion:
```go
	f.FilterSpec = func(in *pb.AccessLevel) {
		if in.GetCustom() != nil {
			if in.GetCustom().Expr == nil {
				in.GetCustom().Expr = &expr.Expr{}
			}
		}
	}
```

### 2. Spellcheck Proto Field Paths
Always pay close attention to `<hint_for_agent>` and proto message structure. For example, KRM might refer to `requireScreenLock`, but the proto field might be spelled `require_screenlock` instead of `require_screen_lock`. Using the exact proto field name in `Unimplemented_NotYetTriaged` is necessary for the fuzzer to correctly ignore the field.

### 3. Ensure Custom Mappers are Safe
During fuzzing, the fuzzer generates completely random proto structures, which may contain fields/pointers that are normally omitted. Ensure that custom mappers (such as `Condition_Members_FromProto`) safely allocate nested structs (like `refs.IAMServiceAccountRef`) instead of performing direct dereferences of nil pointers.
