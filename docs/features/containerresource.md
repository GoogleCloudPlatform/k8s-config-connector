# Customizing Container Resources

You can customize the CPU and memory resource requests and limits for Config Connector pods. You can also customize the number of replicas for the `cnrm-webhook-manager`.

There are two ways to configure container resources, depending on whether you are running Config Connector in cluster mode or namespaced mode.

- **Cluster Mode:** Use a `ControllerResource` object to configure resources for all Config Connector pods in the cluster.
- **Namespaced Mode:** Use a `NamespacedControllerResource` object to configure resources for Config Connector pods in a specific namespace.

## Configurable Pods and Containers

The following table lists the workloads (may be a statefulSet or Deployment) and containers that can be configured.

| Pod                            | Container Name(s)        |
| ------------------------------ | ------------------------ |
| `cnrm-controller-manager`      | `manager`, `prom-to-sd`  |
| `cnrm-webhook-manager`         | `webhook`                |
| `cnrm-deletion-defender`       | `deletiondefender`       |
| `cnrm-resource-stats-recorder` | `recorder`, `prom-to-sd` |
| `cnrm-unmanaged-detector`      | `unmanageddetector`      |

## Configuring Resources in Cluster Mode

To customize resources in cluster mode, create a `ControllerResource` object. The `metadata.name` of the `ControllerResource` object must match the name of the pod you are configuring.

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

To customize resources in namespaced mode, create a `NamespacedControllerResource` object. The `metadata.name` of the `NamespacedControllerResource` object must match the name of the pod you are configuring, and the `metadata.namespace` must match the namespace of the Config Connector installation.

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

| Field      | Description                                                                                                                            |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `containers` | A list of containers whose resource requirements to be customized.                                                                     |
| `replicas`   | The number of desired replicas of the config connector controller. This field takes effect only if the controller name is "cnrm-webhook-manager". |

### NamespacedControllerResourceSpec

| Field      | Description                                                                                                                            |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `containers` | A list of containers whose resource requirements to be customized.                                                                     |

### ContainerResourceSpec

| Field       | Description                                                              |
| ----------- | ------------------------------------------------------------------------ |
| `name`      | The name of the container whose resource requirements will be customized. |
| `resources` | Specifies the resource customization of this container.                  |

### ResourceRequirements

| Field      | Description                                                                                                                                                                                          |
| ---------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `limits`   | Describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/                                                     |
| `requests` | Describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. |
