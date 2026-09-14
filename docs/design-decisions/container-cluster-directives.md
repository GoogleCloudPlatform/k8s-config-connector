# Container Cluster Directives

## Context

The `ContainerCluster` resource supports specific annotations (directives) to
control the behavior of the default node pool created by GKE during cluster
creation.

### `cnrm.cloud.google.com/remove-default-node-pool`

If set to `true`, this directive removes the default node pool created by GKE
immediately after cluster creation. This maps to Terraform's
[`remove_default_node_pool`](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_node_pool#example-usage---using-a-separately-managed-node-pool-recommended)
field.

KCC recommends the following best practices:

1.  To manage node settings, enable `remove-default-node-pool` during cluster
    creation and define nodes using the `ContainerNodePool` resource instead.
2.  If `remove-default-node-pool` is set to `true`, `nodeVersion` and
    `nodeConfig` should generally be omitted from the `ContainerCluster` spec,
    as they are associated with the default pool intended for deletion.
    Including them can lead to reconciliation conflicts or Terraform errors.

## Design Decision: `cnrm.cloud.google.com/remove-default-node-pool-allow-node-config`

To accommodate users who need to remove the default pool while still including
`nodeConfig` in their spec, e.g., due to Organization Policies that require
specific configurations even on the default pool during its brief existence at
creation
(https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7274), the
opt-in annotation
`cnrm.cloud.google.com/remove-default-node-pool-allow-node-config` was
introduced in KCC
[v1.153.0](https://github.com/GoogleCloudPlatform/k8s-config-connector/releases/tag/v1.153.0).

### Behavior

1.  **Initial Creation**: If both `remove-default-node-pool` and
    `remove-default-node-pool-allow-node-config` are set to `true`, KCC
    preserves `nodeConfig` in the configuration sent to GKE.
2.  **Subsequent Reconciliations**: Once the cluster exists and the default pool
    is removed, KCC automatically removes `nodeConfig` from the desired state
    during reconciliation. This prevents permanent diffs and errors that would
    occur if KCC tried to manage a configuration for a deleted node pool.

## Direct Controller Migration Checklist

1.  **Maintain Backward Compatibility**: Ensure both `remove-default-node-pool`
    and `remove-default-node-pool-allow-node-config` annotations remain fully
    functional in the direct controller. Existing test cases should continue to
    pass:

    -   [removedefaultnodepool resource fixture](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/pkg/test/resourcefixture/testdata/directives/removedefaultnodepool)
    -   [removedefaultnodepool e2e scenario](https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/master/tests/e2e/testdata/scenarios/containercluster/removedefaultnodepool)

2.  **Support Field Conflict Validation**: In alignment with KCC best practices
    mentioned above, ensure that `nodeVersion` and `nodeConfig` still strictly
    conflict with the `remove-default-node-pool` annotation (unless the
    allow-node-config override is present) and cannot be set simultaneously in a
    way that causes errors.

3.  **Improve Error Logging**: Surface clear error messages when conflicts
    occur. Currently, Terraform uses the same message for both `nodeVersion` and
    `nodeConfig` conflicts, which can be confusing for users.
