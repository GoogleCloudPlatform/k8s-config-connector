### [2026-06-13] Implementation of ComputeTargetGRPCProxy KRM Types
- **Context**: Implementing `ComputeTargetGRPCProxy` under `apis/compute/v1alpha1/`.
- **Problem**: `TargetGrpcProxy` is a global resource in GCP (using url pattern `projects/*/global/targetGrpcProxies/*`) but contains no location/region segment in its API. However, Config Connector's general CRD spec conventions require `location` in the Spec.
- **Solution**: We implemented `ComputeTargetGRPCProxySpec` to require `location` to match CRD expectations, while keeping the identity structure and `gcpurls.Template` scoped to `projects/{project}/global/targetGrpcProxies/{targetGrpcProxy}`.
- **Impact**: Provides a blueprint for direct global compute resources that might still declare `location` in their spec while mapping to global paths.
