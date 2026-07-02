# Parallelstore Journal

### [2026-07-02] Scaffolding ParallelstoreInstance
- **Context**: Scaffolded the KRM types, CRD, and IdentityV2 for `ParallelstoreInstance` under `v1alpha1`.
- **Problem**: Enums and required fields defined in parallelstore.proto needed manual validation annotations and alignment with the direct controller pattern.
- **Solution**: Defined `computev1beta1.ComputeNetworkRef` as a reference for VPC `networkRef` instead of using basic strings. Correctly mapped all enums (FileStripeLevel, DirectoryStripeLevel, DeploymentType, and State) and added kubebuilder validations. Added necessary labels and registered with the scheme.
- **Impact**: The next agent implementing Phase 2 (the controller logic and E2E tests) can immediately use these types with proper KRM reference resolution and validation schemas.
