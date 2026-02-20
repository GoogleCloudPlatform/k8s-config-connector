# Boolean Field Handling in TF-based Resources

## Context

When we patch a new field in TF (Terraform), we should use `d.Get()` for expansion it this is an optional boolean field.

This doc captures the details about why we made this design decision.

Please note that we should minimize TF patches if possible.

## KCC-Specific Behavior

Understanding KCC's mapping logic is prerequisite to making the correct decision in the Terraform Provider.

1.  **Zero Values (e.g., `false`, `0`, `""`):** KCC treats setting a primitive field to its zero value (e.g. setting a
    boolean field to `false`) as an "unset" intention in the Terraform configuration. It does **not** pass the zero
    value (`false`, `0`, `""`, etc) to Terraform; the field appears as unset in `schema.ResourceData`.
2.  **Unset (Omitted in KCC):**
    *   **Creation:** The field is unset in Terraform.
    *   **Update:** KCC treats this as "stop managing." It passes the **current live value from GCP** to Terraform to
        prevent drift/unintended changes.

## TF-Specific Behavior for Boolean Fields

In the context of Terraform providers, handling boolean fields (type `schema.TypeBool`) that are marked as `Optional`
can be tricky due to how `schema.ResourceData`'s `Get` method works.

When a boolean field is defined with `Optional: true` in the Terraform schema, calling `d.Get("field_name")` will return
`false` in two distinct scenarios:
*   The user explicitly sets `field_name = false` in their Terraform configuration.
*   The user entirely omits `field_name` from their configuration.
    Since false is the zero value for booleans in Go, `d.Get()` returns this zero value when the field isn't present, making
    it impossible to distinguish between an explicit false and an omitted field using `d.Get()` alone.

If `d.Get()` is used, and the field is omitted by the user, the provider will still see `false` and might mistakenly
send `false` to the API, thus "enforcing" `false` even when the user intended to rely on the API's default. This can lead
to unexpected behavior or prevent the API's actual default from taking effect.

Alternatively, `d.GetOk("field_name")` can be used to explicitly differentiate between an unset field (`ok` is `false`)
and a field explicitly set to `false` (`ok` is `true`, value is `false`). TF provider can then choose to send the value
to the API only when `ok` is `true`, thus honoring the user's intent to omit the field and let the API decide the
default.

## Analysis of `d.Get(...)` and `d.GetOk(...)`

### Approach A: `d.Get(...)` (Current/Chosen)

*   **Mechanism:** `d.Get` returns the zero value (`false`) if the field is unset in the Terraform configuration.
*   **Scenario 1: User sets `true` in KCC:**
    *   KCC passes `true` to TF.
    *   `d.Get` returns `true`.
    *   **Result:** Sends `true` to API. (Correct)
*   **Scenario 2: User sets `false` in KCC:**
    *   KCC **unsets** the field in TF (Zero Value behavior).
    *   `d.Get` returns `false`.
    *   **Result:** Sends `false` to API. (Correct).
*   **Scenario 3: User omits field (Update):**
    *   KCC passes **live value** to TF.
    *   If live is `true` -> `d.Get` returns `true` -> Sends `true`.
    *   If live is `false` -> `d.Get` returns `false` -> Sends `false`.
    *   **Result:** Preserves live state. (**May confuse user if their intention is to unset the field in API**).
*   **Scenario 4: User omits field (Create):**
    *   Field is unset or is set to the default value in TF.
    *   `d.Get` returns `false` or the default value.
    *   **Result:** Sends `false` or the default value to API. (Correct).

### Approach B: `d.GetOk(...)` (Rejected)

*   **Mechanism:** `d.GetOk` returns `ok=false` if the field is unset in the Terraform configuration.
*   **Scenario 1: User sets `true` in KCC:**
    *   KCC passes `true` to TF.
    *   `d.GetOk` returns `ok=true`, `value=true`.
    *   **Result:** Sends `true` to API. (Correct)
*   **Scenario 2: User sets `false` in KCC:**
    *   KCC **unsets** the field in TF (Zero Value behavior).
    *   `d.GetOk` returns `ok=false`.
    *   Block is skipped. Field is **omitted** from API request.
    *   **Result:** **INCORRECT.** The user explicitly asked for `false`, but we sent nothing. If the API default is
        `true` (or if we needed to overwrite a previous `true`), this fails.
*   **Scenario 3: User omits field (Update):**
    *   KCC passes **live value** to TF.
    *   If live is `true` -> `d.GetOk` returns `ok=true`, `value=true` -> Sends `true`.
    *   If live is `false` -> `d.GetOk` returns `ok=true`, `value=false` -> Sends `false`.
    *   **Result:** Preserves live state. (**May confuse user if their intention is to unset the field in API**).
*   **Scenario 4: User omits field (Create):**
    *   Field is unset or is set to the default value in TF.
    *   If field is unset -> `d.GetOk` returns `ok=false` -> Block is skipped. Field is **omitted** from API request.
    *   If field has a default value -> `d.GetOk` returns `ok=true`, `value=default_value` -> Sends `default_value`.
    *   **Result:** Sends default value to API or let API defaults it. (Correct).

## Decision

**Use `d.Get(...)` for optional boolean field expansion.**

### Rationale

1.  **Correctness for Zero Values:** Due to KCC's behavior of mapping `false` to "unset" in Terraform, `d.Get(...)` is
    the **only** mechanism that allows us to correctly enforce a user's explicit request for `false`. `d.GetOk(...)`
    would silently ignore the user's intent.
2.  **Config Management Practice:** Users managing infrastructure via KCC often define their desired state
    declaratively. If they want a feature off, they set it to `false`. It is more common to actively manage the state of
    critical flags than to leave them unmanaged and rely on opaque API behaviors that might change over time. Therefore,
    ensuring `false` (zero value) works reliably covers the primary use case for managing these resources.
3.  **Security:** In the creation scenario where the user omits the field, defaulting to sending `false` (via `d.Get`'s
    default) is usually a safer behavior for a security-sensitive field.