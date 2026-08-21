# Change Cookies (Stateful Reconciliation)

## Context

KCC traditionally uses a stateless reconciliation model, comparing the user's `.spec` against the live GCP resource state. While simple, this approach has limitations when dealing with server-defaulted fields and unmanaged fields, especially in complex resources.

To address these limitations, we are introducing a stateful reconciliation mechanism using "change cookies". A change cookie is a hash stored in the `cnrm.cloud.google.com/last-changed-cookie` annotation that captures the state of the resource after a successful reconciliation.

## Design Decision

We will use change cookies sparingly to improve the efficiency and reliability of reconciliation for complex resources.

### 1. Direct Controllers Only

Change cookies are only supported for resources managed by the "direct" reconciliation path. Using them with legacy Terraform or DCL-based controllers could break established "safer upgrade" and "dry run" workflows that rely on the existing stateless behavior.

### 2. Handling False-Positive Reconciliations

We acknowledge that change cookies may occasionally lead to false-positive reconciliations. For example, if the underlying GCP API starts returning additional fields or changes its normalization logic, the observed state hash may change even if the user has not modified their spec.

However, we believe this is acceptable because:
*   It should only happen once per API or KCC change.
*   Once the reconciliation completes, a new, updated cookie will be stored, and subsequent reconciliations will return to the "fast path."
*   This is preferable to continuous reconciliation (the "always-diff" problem) that can occur when defaulting logic is imperfect.

### 3. Sparingly Used

Change cookies are not intended to be a universal replacement for good defaulting and normalization logic. They will be used sparingly, primarily for the most complicated resources where maintaining perfect stateless defaulting logic is prohibitively difficult or error-prone.

If a resource can be reliably reconciled using the stateless model, it should continue to do so.

### 4. Safer Upgrades Integration

The "safer upgrades" process, which verifies that a new version of KCC does not produce unintended changes, may need to be updated. Currently, it might run a reconciliation twice (e.g., once with the legacy controller and once with the direct controller). With change cookies, we may need to verify reconciliation behavior both with and without an existing cookie to ensure consistency.
