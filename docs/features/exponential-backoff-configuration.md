# Exponential Backoff Configuration

Config Connector supports configuring the maximum exponential backoff retry delay for individual resources via the `cnrm.cloud.google.com/backoff-max-delay-in-seconds` annotation.

When a resource fails to reconcile, Config Connector retries reconciliation using an exponential backoff schedule (`250ms` base delay, doubling on each consecutive failure: `0.25s, 0.5s, 1s, 2s, 4s, ...`) up to a default maximum delay ceiling of `120` seconds (`2` minutes). By setting `cnrm.cloud.google.com/backoff-max-delay-in-seconds` on a resource, you can customize this retry ceiling or halt automatic error retries altogether.

> [!NOTE]
> **Scope**: The `cnrm.cloud.google.com/backoff-max-delay-in-seconds` annotation is currently supported **exclusively for resources managed by Direct controllers**.

## Use Cases

* **Reduce GCP API quota consumption on failing resources**: Increase the maximum backoff delay ceiling (for example, `"600"` or `"3600"`) for resources that are blocked on slow external dependencies or persistent configuration issues to avoid frequent failed API calls.
* **Halt automated retries during debugging**: Set the annotation to `"0"` so that a failing resource stops automatically retrying until you update its `.spec` or annotations.
* **Accelerate recovery for time-sensitive resources**: Lower the maximum backoff delay ceiling (for example, `"10"` or `"30"`) so that transient errors are retried more frequently.

## Supported Values and Behavior

The annotation accepts a string representing a non-negative 32-bit integer (in seconds):

| Annotation Value | Behavior |
| :--- | :--- |
| **Unset or `""` (Default)** | Uses the default exponential backoff (`250ms` base delay, doubling per consecutive failure) capped at a maximum delay of **`120` seconds**. |
| **Positive integer (e.g., `"30"`, `"600"`)** | Uses exponential backoff (`250ms * 2^(failures - 1)`) capped at the specified number of seconds. When a reconciliation attempt succeeds, the failure counter is reset (`Forget`). |
| **`"0"`** | **Halts automated error retries.** If a reconciliation attempt fails, Config Connector records the failure condition on the resource's `.status` and does **not** schedule automatic retry attempts. |
| **Negative or non-integer value** | Logs an invalid annotation error and falls back to the default `120` second maximum backoff delay. |

### Resuming Reconciliation When Set to `"0"`

When `cnrm.cloud.google.com/backoff-max-delay-in-seconds: "0"` is set and a resource enters a failed state, Config Connector will only re-trigger reconciliation when a relevant change is made to the Kubernetes object, such as:

* Updating `.spec` (which increments `metadata.generation`)
* Updating `metadata.labels`
* Changing or removing `cnrm.cloud.google.com/backoff-max-delay-in-seconds`
* Changing other watched annotations (`cnrm.cloud.google.com/reconcile-interval-in-seconds`, `cnrm.cloud.google.com/actuation-mode`, or project/folder/organization container annotations)

Once a subsequent reconciliation succeeds and the resource reaches `Ready` (`UpToDate`), normal periodic drift-correction reconciliation (`reconcile-interval-in-seconds`) resumes.

### Behavior During Resource Deletion

To prevent resources from becoming stuck in a `Terminating` state due to transient deletion errors, **custom backoff ceilings and `"0"` retry halting do not apply when a resource is being deleted** (`metadata.deletionTimestamp` is set). Deletion failures always use standard controller-runtime exponential backoff retries until deletion succeeds.

## Configuration Examples

### Example 1: Increase the Maximum Backoff Delay Ceiling

The following example configures a `LoggingLogMetric` resource to cap exponential backoff retries at `600` seconds (`10` minutes) instead of the default `120` seconds when reconciliation fails:

```yaml
apiVersion: logging.cnrm.cloud.google.com/v1beta1
kind: LoggingLogMetric
metadata:
  name: logginglogmetric-sample
  annotations:
    cnrm.cloud.google.com/backoff-max-delay-in-seconds: "600"
spec:
  filter: "resource.type=gae_app AND severity>=ERROR"
  projectRef:
    external: "projects/my-project-id"
```

With this configuration, consecutive failures will retry after `0.25s`, `0.5s`, `1s`, `2s`, `4s`, `8s`, `16s`, `32s`, `64s`, `128s`, `256s`, `512s`, and then cap at `600s` for subsequent retries until reconciliation succeeds.

### Example 2: Halt Automated Error Retries (`"0"`)

The following example sets `cnrm.cloud.google.com/backoff-max-delay-in-seconds: "0"` so that if reconciliation fails, Config Connector will not automatically retry:

```yaml
apiVersion: logging.cnrm.cloud.google.com/v1beta1
kind: LoggingLogMetric
metadata:
  name: logginglogmetric-no-retry
  annotations:
    cnrm.cloud.google.com/backoff-max-delay-in-seconds: "0"
spec:
  filter: "resource.type=gae_app AND severity>=ERROR"
  projectRef:
    external: "projects/my-project-id"
```

If this resource fails reconciliation:
1. Its `.status.conditions` will reflect the failure (`Ready: False`, `Reason: UpdateFailed` or `CreateFailed`).
2. Config Connector will log `Halting automated error retries as backoff-max-delay-in-seconds is set to 0` and will not requeue the resource.
3. To trigger a retry after resolving the issue, either update `.spec` or change/remove the `cnrm.cloud.google.com/backoff-max-delay-in-seconds` annotation.

---

## Further Reading

* [Choose Controller Implementations (Direct vs. Terraform)](controller-configuration.md) to verify or configure whether a resource uses a Direct controller.
* [Configure Rate Limiting](ratelimit.md) for configuring controller-wide Kubernetes API client QPS and burst limits.
* [Pause Reconciliation](pause.md) for pausing all actuation on a resource, namespace, or cluster.
