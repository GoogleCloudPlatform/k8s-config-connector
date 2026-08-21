### [2026-06-05] DataprocAutoscalingPolicy Identity & Refs
- **Context**: Moving DataprocAutoscalingPolicy to `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**: 
    1. DataprocAutoscalingPolicy did not have `IdentityV2` or `refs.Ref` implementations.
    2. The resource is managed as a DCL resource, but lacks modern identity and reference support.
- **Solution**:
    1. Implemented `apis/dataproc/v1beta1/dataprocautoscalingpolicy_identity.go` utilizing `identity.IdentityV2` and backed by `gcpurls.Template`.
    2. Implemented `apis/dataproc/v1beta1/dataprocautoscalingpolicy_reference.go` utilizing standard `refs.Ref` matching modern controller standards.
    3. Added comprehensive unit tests in `dataprocautoscalingpolicy_identity_test.go` and `dataprocautoscalingpolicy_reference_test.go`.
    4. Regenerated deepcopy functions for package `v1beta1` which added `DeepCopyInto` and `DeepCopy` for `DataprocAutoscalingPolicyRef` while keeping `DataprocAutoscalingPolicyIdentity` without deepcopy generation using `+k8s:deepcopy-gen=false`.
- **Impact**: DataprocAutoscalingPolicy is fully modernized and ready for direct controller migration or advanced identity resolving.
