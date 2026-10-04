# Terminal Error Mode

> `Feature State`: `alpha` as of version v1.159 (Direct Controllers)

Config Connector direct controllers can be configured to halt infinite exponential-backoff retries when encountering permanent, deterministic client errors from Google Cloud APIs (such as malformed configurations, invalid arguments, or attempts to mutate immutable fields).

By default, Kubernetes controllers retry failed reconciliations indefinitely using exponential backoff. While this ensures eventual consistency for transient errors or out-of-order resource creation, permanent client errors (such as HTTP 400 / gRPC `InvalidArgument`) can exhaust customer API quotas and generate noisy logs.

`terminal-error-mode` provides status-driven error pausing to stop continuous retries until the resource specification is corrected.

---

## Configuring Terminal Error Mode

Terminal error pausing is controlled via the `cnrm.cloud.google.com/terminal-error-mode` annotation on individual Direct resources.

### Supported Modes

| Value | Description |
| :--- | :--- |
| `"none"` *(default / unset)* | Standard Config Connector behavior. The controller continuously retries reconciliation using exponential backoff on all errors. |
| `"builtin"` | Halts reconciliation retries on verified terminal client errors. Sets condition `Ready=False` with `reason: UpdateFailedTerminalError` and halts workqueue requeueing. |

### Example

```yaml
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: my-bucket
  namespace: config-connector
  annotations:
    cnrm.cloud.google.com/terminal-error-mode: "builtin"
spec:
  # ...
```

---

## Error Qualification Criteria

When `terminal-error-mode: "builtin"` is enabled, Config Connector evaluates errors from the Google Cloud API:

### Qualified Terminal Errors
The controller qualifies an error as terminal if it meets all of the following criteria:
1. Returns HTTP `400 Bad Request` or gRPC `codes.InvalidArgument (3)`.
2. Matches verified invalid argument payload details (e.g. `google.rpc.BadRequest` field violations, `google.rpc.ErrorInfo` with reasons like `INVALID_ARGUMENT`, `FIELD_VIOLATION`, `IMMUTABLE_FIELD`, `MISSING_REQUIRED_FIELD`, or syntax errors).

### Transient Exception Exclusions
Some GCP services return HTTP 400 for transient dependency states or concurrency locks. These are explicitly excluded from terminal error pausing and continue to retry normally:
- `RESOURCE_IN_USE_BY_ANOTHER_RESOURCE`
- `RESOURCE_NOT_READY`
- `OPERATION_IN_PROGRESS`
- `CONCURRENT_MUTATION`
- `LOCKED`

Other non-400 errors (HTTP 403, 404, 409, 429, 5xx) and network connectivity issues are also treated as retryable to preserve eventual consistency.

---

## Lifecycle & Resume Behavior

### Status Reporting
When a terminal error occurs:
- `status.conditions[Ready]` is set to `status: "False"` with `reason: "UpdateFailedTerminalError"`.
- `status.conditions[Ready].message` contains the detailed error message from the GCP API.
- `status.observedGeneration` is updated to match `metadata.generation`.
- Requeueing is halted (`reconcile.Result{}, nil`), stopping API calls and controller log spam.

### Auto-Resume on Spec Updates
When a user updates the resource manifest `.spec`:
1. Kubernetes increments `metadata.generation`.
2. Because `metadata.generation > status.observedGeneration`, the terminal error pause is automatically bypassed.
3. The controller immediately re-attempts reconciliation against Google Cloud APIs.

### Manual Force Retry
To manually force reconciliation without changing `.spec`:
- Change the annotation to `cnrm.cloud.google.com/terminal-error-mode: "none"`, or
- Remove the `cnrm.cloud.google.com/terminal-error-mode` annotation.

The controller's event predicate detects annotation changes and immediately enqueues the resource for reconciliation.

### Deletion Precedence
Deletion requests (`metadata.deletionTimestamp != nil`) always bypass terminal error pausing. If a resource in a terminal error state is deleted (`kubectl delete`), the controller immediately executes deletion reconciliation and finalizer removal.

---

## Further Reading

* [Controller Configuration](controller-configuration.md): Choosing between Direct and Legacy controllers.
* [Pause Reconciliation](pause.md): Global and namespace-level actuation pausing.
