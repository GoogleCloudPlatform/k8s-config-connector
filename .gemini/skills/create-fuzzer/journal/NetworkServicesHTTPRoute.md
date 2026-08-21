# Journal: KRM Fuzzer for NetworkServicesHTTPRoute

## Observations

1. **Boolean zero-value mapping loss in `oneof`**:
   The generated mapper uses `direct.LazyPtr(in.GetPresentMatch())` for mapping `MatchType` inside `HttprouteQueryParameters` and `HttprouteHeaders`. Since `PresentMatch` is a boolean, and `direct.LazyPtr(false)` returns `nil`, any `PresentMatch: false` (which was explicitly randomized in the fuzz tests) became `nil` after the round-trip conversion.
   - **Fix**: We hand-coded custom mappers `HttprouteQueryParameters_FromProto`/`ToProto` and `HttprouteHeaders_FromProto`/`ToProto` to directly inspect the proto `oneof` variant (e.g., checking `in.MatchType.(type)`), preserving both `true` and `false` correctly without losing data.

2. **Unimplemented fields**:
   Several nested fields in the `HttpRoute` proto were identified during randomized fuzz testing as having no KRM equivalent, such as:
   - `rules[].action.direct_response`
   - `rules[].action.idle_timeout`
   - `rules[].action.stateful_session_affinity`
   - `rules[].action.destinations[].request_header_modifier`
   - `rules[].action.destinations[].response_header_modifier`
   - `rules[].action.request_mirror_policy.destination.request_header_modifier`
   - `rules[].action.request_mirror_policy.destination.response_header_modifier`
   - `rules[].action.request_mirror_policy.mirror_percent`

   These were safely marked as `Unimplemented_NotYetTriaged` in the fuzzer to bypass round-trip comparisons on those fields.
