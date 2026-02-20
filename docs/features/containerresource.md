# Customizing Container Resources

You can customize the CPU and memory resource requests and limits for Config Connector pods. You can also customize the number of replicas for the `cnrm-webhook-manager` and enable Vertical Pod Autoscaler (VPA) for supported pods.

There are two ways to configure container resources, depending on whether you are running Config Connector in cluster mode or namespaced mode.

- **Cluster Mode:** Use a `ControllerResource` object to configure resources for all Config Connector pods in the cluster.
- **Namespaced Mode:** Use a `NamespacedControllerResource` object to configure resources for Config Connector pods in a specific namespace.

## Configurable Pods and Containers

The following table lists the workloads (StatefulSet or Deployment) and containers that can be configured.

| Pod Name                       | Mode Support            | Workload Type | Container Name(s)        |
| ------------------------------ | ----------------------- | ------------- | ------------------------ |
| `cnrm-controller-manager`      | Cluster and Namespaced  | StatefulSet   | `manager`, `prom-to-sd`  |
| `cnrm-webhook-manager`         | Cluster only            | Deployment    | `webhook`                |
| `cnrm-deletiondefender`        | Cluster only            | StatefulSet   | `deletiondefender`       |
| `cnrm-resource-stats-recorder` | Cluster only            | Deployment    | `recorder`, `prom-to-sd` |
| `cnrm-unmanaged-detector`      | Cluster only            | StatefulSet   | `unmanageddetector`      |

Note: In Namespaced Mode, only the `cnrm-controller-manager` pod can be customized using `NamespacedControllerResource` because other components are shared across the cluster.

## Configuring Resources in Cluster Mode

To customize resources in cluster mode, create a `ControllerResource` object. The `metadata.name` of the `ControllerResource` object must match the name of the pod you are configuring (e.g., `cnrm-controller-manager`).

### Example: Customizing `cnrm-controller-manager`

The following example shows how to customize the CPU and memory for the `manager` container in the `cnrm-controller-manager` pod.

```yaml
apiVersion: customize.core.cnrm.cloud.google.com/v1beta1
kind: ControllerResource
metadata:
  name: cnrm-controller-manager
spec:
  containers:
    - name: manager
      resources:
        limits:
          cpu: 200m
          memory: 512Mi
        requests:
          cpu: 100m
          memory: 256Mi
```

### Example: Customizing `cnrm-webhook-manager` and Replicas

The following example shows how to customize the memory for the `webhook` container in the `cnrm-webhook-manager` pod and set the number of replicas to 4.

```yaml
apiVersion: customize.core.cnrm.cloud.google.com/v1beta1
kind: ControllerResource
metadata:
  name: cnrm-webhook-manager
spec:
  replicas: 4
  containers:
    - name: webhook
      resources:
        limits:
          memory: 512Mi
        requests:
          memory: 256Mi
```

## Configuring Resources in Namespaced Mode

To customize resources in namespaced mode, create a `NamespacedControllerResource` object. The `metadata.name` of the `NamespacedControllerResource` object must match the name of the pod you are configuring, and the `metadata.namespace` must match the namespace where Config Connector is installed for that namespace.

### Example: Customizing `cnrm-controller-manager` in a specific namespace

The following example shows how to customize the CPU and memory for the `manager` container in the `cnrm-controller-manager` pod in the `config-control` namespace.

```yaml
apiVersion: customize.core.cnrm.cloud.google.com/v1beta1
kind: NamespacedControllerResource
metadata:
  name: cnrm-controller-manager # name should not contain the namespace ID suffix
  namespace: config-control
spec:
  containers:
    - name: manager
      resources:
        limits:
          cpu: 200m
          memory: 512Mi
        requests:
          cpu: 100m
          memory: 256Mi
```

## Configuration Options

### ControllerResourceSpec

| Field                       | Type                     | Description                                                                                                                                                                     |
| --------------------------- | ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `containers`                | `[]ContainerResourceSpec`| A list of containers whose resource requirements are to be customized. Mutually exclusive with `verticalPodAutoscalerMode: Enabled`.                                           |
| `replicas`                  | `int64`                  | The number of desired replicas. This field only takes effect if the `metadata.name` is `cnrm-webhook-manager`.                                                                  |
| `verticalPodAutoscalerMode` | `string`                 | Mode of Vertical Pod Autoscaler for the controller. Allowed values: `Enabled`, `Disabled`. Defaults to `Disabled`. Mutually exclusive with `containers`.                        |
| `metadataHost`              | `string`                 | Overrides the GCP metadata server hostname (injected as `GCE_METADATA_HOST`). Useful for IPv6-only clusters where the default `169.254.169.254` is unreachable.                |

### NamespacedControllerResourceSpec

| Field                       | Type                     | Description                                                                                                                                                                     |
| --------------------------- | ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `containers`                | `[]ContainerResourceSpec`| A list of containers whose resource requirements are to be customized. Mutually exclusive with `verticalPodAutoscalerMode: Enabled`.                                           |
| `verticalPodAutoscalerMode` | `string`                 | Mode of Vertical Pod Autoscaler for the controller. Allowed values: `Enabled`, `Disabled`. Defaults to `Disabled`. Mutually exclusive with `containers`.                        |

### ContainerResourceSpec

| Field       | Type                   | Description                                                              |
| ----------- | ---------------------- | ------------------------------------------------------------------------ |
| `name`      | `string`               | The name of the container (e.g., `manager`, `webhook`, `prom-to-sd`).    |
| `resources` | `ResourceRequirements` | Specifies the resource customization (requests and limits) for this container. |

### ResourceRequirements

| Field      | Type           | Description                                                                                                                                                                                          |
| ---------- | -------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `limits`   | `ResourceList` | Describes the maximum amount of compute resources allowed. More info: [Kubernetes Resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)                        |
| `requests` | `ResourceList` | Describes the minimum amount of compute resources required. If omitted, it defaults to `limits` if specified. More info: [Kubernetes Resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) |

## Enabling Vertical Pod Autoscaler

You can enable the Vertical Pod Autoscaler (VPA) for Config Connector pods by setting the `verticalPodAutoscalerMode` field to `Enabled`. 

When VPA is enabled, Config Connector automatically creates a `VerticalPodAutoscaler` resource for the target workload with `updateMode: Auto`. The operator will also periodically fetch recommendations from the VPA and apply them to the pod specifications to ensure they remain in sync with VPA's suggestions.

**Important:** `verticalPodAutoscalerMode: Enabled` is mutually exclusive with specifying `containers` in the same object. When VPA is enabled, the `containers` field should be empty or omitted.

### Example: Enabling VPA for `cnrm-controller-manager` in Cluster Mode

```yaml
apiVersion: customize.core.cnrm.cloud.google.com/v1beta1
kind: ControllerResource
metadata:
  name: cnrm-controller-manager
spec:
  verticalPodAutoscalerMode: Enabled
  containers: []
```

### Example: Enabling VPA for `cnrm-controller-manager` in Namespaced Mode

```yaml
apiVersion: customize.core.cnrm.cloud.google.com/v1beta1
kind: NamespacedControllerResource
metadata:
  name: cnrm-controller-manager # name should not contain the namespace ID suffix
  namespace: config-control
spec:
  verticalPodAutoscalerMode: Enabled
  containers: []
```
