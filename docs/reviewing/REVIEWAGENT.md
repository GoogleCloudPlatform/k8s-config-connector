### Use pointer types in CRD types

When creating the go type for a CRD, use `*string` instead of `string`, `*int32` instead of `int32` etc,
so we can tell the difference between the zero value and unset fields.
We could think about it on a case-by-case basis, but it's easier not to.
Where a field is required, you can annotate that with `// +required` rather than using the non-pointer type.
