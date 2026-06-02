### [2026-06-02] ParameterManagerParameterVersion Identity & Refs
- **Context**: Moving ParameterManagerParameterVersion to `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**: 
    1. ParameterManagerParameterVersion was using the older `identity.Identity` pattern with nested parent structs (`ParameterIdentity`), which required multiple steps to access `location` or `project`.
    2. No reference file (`parameterversion_reference.go`) existed for this resource.
- **Solution**:
    1. Reimplemented `apis/parametermanager/v1alpha1/parameterversion_identity.go` using `identity.IdentityV2` backed by the canonical `gcpurls.Template`.
    2. Created `apis/parametermanager/v1alpha1/parameterversion_reference.go` to implement `ParameterManagerParameterVersionRef` conforming to `refs.Ref`.
    3. Updated the direct controller `parameterversion_controller.go` to parse, resolve, and export using the simplified `ParameterManagerParameterVersionIdentity` structure.
- **Impact**: Code complexity in both the identity definition and the direct controller was greatly reduced. Clean, modular, and fully aligned with modern KCC patterns.
