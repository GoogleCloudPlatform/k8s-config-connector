# Pausing Re-Reconciliation on Non-Retryable Errors

## Context

Traditionally, when a Kubernetes Config Connector (KCC) controller encounters a reconciliation failure, it returns an error to the underlying `controller-runtime` workqueue. This triggers automatic exponential backoff retries.

While this behavior is appropriate for transient errors (such as network glitches, temporary server errors, or eventually consistent states), it is highly inefficient for permanent or client-side non-retryable errors (such as `InvalidArgument` or `Unimplemented`). For these errors, the resource is stuck in a loop of failing and retrying indefinitely, which wastes computing resources, creates excessive API traffic to GCP, and floods the operator logs with repetitive errors.

## Design Decision

To solve this problem, we are introducing a mechanism to **pause actuation** specifically for resources managed by **Direct Controllers** when they encounter a verified non-retryable error.

### 1. Direct Controllers Only

This optimization is restricted exclusively to KCC's newer **direct controllers**. Legacy controllers (Terraform and DCL-based) and hand-written IAM controllers will continue using their existing retry mechanisms.

### 2. Identifying Non-Retryable Errors

We classify an error as non-retryable using high-confidence criteria that indicate client-side or permanent API issues. Currently, we start with a minimal, high-confidence set of error types:
*   **gRPC Status Codes**: `codes.InvalidArgument` and `codes.Unimplemented`.
*   **HTTP Status Codes**: `400` (Bad Request) and `501` (Not Implemented).

Other errors (such as `FailedPrecondition`, authentication/permission issues `401`/`403`, rate limits/quota, or transient server-side errors) are treated as retryable since they might resolve over time or through external actions.

### 3. Actuation-Paused Annotation

Instead of swallowing errors or continuously failing, when a direct controller encounters a non-retryable error during `Create` or `Update`:
1.  The resource's `status.conditions` are updated to reflect the error (via the standard error-handling path).
2.  The controller updates the resource's metadata to add the annotation:
    `cnrm.cloud.google.com/actuation-paused: "true"`
3.  The controller then returns `reconcile.Result{}, nil` to `controller-runtime` to prevent further automatic retries.

Updating the resource on the server triggers a new reconciliation event. In this (and all future) reconciliation runs:
*   The controller inspects the annotations.
*   If `cnrm.cloud.google.com/actuation-paused` is set to `"true"`, it checks if the resource has been updated by the user since the pause.
*   If `metadata.generation > status.observedGeneration`, the user has updated the spec. The controller automatically clears the `cnrm.cloud.google.com/actuation-paused` annotation and proceeds with reconciliation.
*   If the generation is unchanged, the controller skips actuation and immediately exits with `reconcile.Result{}, nil`.

### 4. Retriggering Reconciliation

Users can easily retrigger actuation of a paused resource in two ways:
1.  **Updating the `spec`**: Any change to the resource's configuration increments `metadata.generation`, which clears the pause annotation and initiates actuation.
2.  **Manually clearing/updating the annotation**: Setting `cnrm.cloud.google.com/actuation-paused: "false"` or removing the annotation entirely will retrigger actuation.
