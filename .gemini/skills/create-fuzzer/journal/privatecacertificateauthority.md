# Fuzzer Implementation Journal: PrivateCACertificateAuthority

## Observations and Learnings
1. **Missing Status Time Mappings**: While transitioning time fields (`create_time`, `delete_time`, `expire_time`, `update_time`) from unimplemented to actively fuzzed status fields, we discovered that `UpdateTime` was completely missing in the `PrivateCACertificateAuthorityStatus_FromProto` mapper logic. This was fixed by mapping it using `direct.StringTimestamp_FromProto`.
2. **Promoting Time Fields**: By changing `create_time`, `delete_time`, `expire_time`, and `update_time` to active `StatusField` configurations, the fuzzer now actively asserts that GCP's timestamp fields map back and forth to KRM Status fields perfectly.
3. **Thorough Spec and Status Field Comparison**: Adding detailed documented comparisons for all spec and status fields directly in the fuzzer makes it easy for reviewers to audit the mapping coverage without looking at multiple files.
