# Configuring the Rate Limit

Config Connector uses a token bucket to rate limit requests to the Kubernetes API server. You can customize the rate limit settings to better suit your needs.

Increasing the rate limit values can reduce the time it takes for your resources to be reconciled, especially if you have a large number of IAM resources.

**Warning:** Increasing the rate limit can cause Config Connector to consume more of your GCP quota, particularly your read quota. Monitor your quota usage carefully after changing these settings.

## Namespaced Mode

Namespaced mode is the preferred installation mode for Config Connector.

In namespaced mode, you can configure the rate limit for each namespace by creating a `NamespacedControllerReconciler` resource. This allows you to have different rate limit settings for different namespaces.

To properly reason about quota usage with this configuration, it is valuable to keep namespaces and GCP projects on a 1:1 relationship.

Here is an example of a `NamespacedControllerReconciler` resource that sets the `qps` and `burst` values to 80:

```yaml
apiVersion: customize.core.cnrm.cloud.google.com/v1beta1
kind: NamespacedControllerReconciler
metadata:
  name: cnrm-controller-manager
  namespace: YOUR_NAMESPACE # Replace with your namespace
spec:
  rateLimit:
    qps: 80
    burst: 80
```

## Cluster Mode

In cluster mode, you can configure the rate limit for the entire cluster by creating a `ControllerReconciler` resource.

Here is an example of a `ControllerReconciler` resource that sets the `qps` and `burst` values to 80:

```yaml
apiVersion: customize.core.cnrm.cloud.google.com/v1beta1
kind: ControllerReconciler
metadata:
  name: cnrm-controller-manager
spec:
  rateLimit:
    qps: 80
    burst: 80
```
