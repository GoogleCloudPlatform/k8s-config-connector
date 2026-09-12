# Boolean and Zero-Value Handling in TF-based Resources

## Context

Handling optional boolean fields (and other primitive zero-values like `0`, `""`) in Terraform-based controllers presents fundamental challenges due to the design of the legacy Terraform Plugin SDK v2 and KCC's zero-value mapping.

This document details the end-to-end reconciliation mechanics, the "Non-Zero Default Trap", why `d.Get()` is chosen over `d.GetOk()`, and the long-term solution via Direct Controllers.

Please note that we should minimize TF patches if possible.

---

## 1. KCC & Terraform SDK v2 Zero-Value Mechanics

In Go, primitive types do not have `nil` states—their zero values are `false`, `0`, `""`, `[]`, `{}`.

### The Collapse of Unset and Zero-Value:
* **In Terraform Plugin SDK v2:**
  * `d.Get("bool_field")` returns `bool` (never `nil`). It returns `false` for both unset fields and explicit `false`.
  * `d.GetOk("bool_field")` checks `val != nil && !isZero(val)`. Because `isZero(false) == true`, **`d.GetOk()` returns `ok = false` for both unset fields and explicit `false`**.
* **In KCC KRM Mapping:**
  * Setting a primitive field to its zero value (e.g., `false`) appears as "unset/zero" in Terraform's configuration reader.

---

## 2. Creation vs. Update Lifecycle Matrix

> **Note:** This lifecycle behavior assumes that the field either has no default or defaults to `false` (zero value) in the Terraform schema. If a non-zero default (`Default: true`) is specified in the schema, see Section 3 for the resulting failure mode.

| Stage & User Input | KCC SSA Management | Value in TF Desired Config | Result on GCP |
| :--- | :--- | :--- | :--- |
| **Zero value at Creation** (`field: false`) | User-managed | `false` | Created with `false` (or GCP default). |
| **Unset at Creation** (omitted in YAML) | Unmanaged | `false` | Created with `false` (or GCP default). |
| **Zero value at Update** (`field: false`) | **User-managed** (`managedFields`) | `false` | **Actively enforces `false`** on GCP. |
| **Unset at Update** (omitted in YAML) | **Unmanaged** | **Copied from live state** | **Preserves live state** ("stop managing"). |

---

## 3. The "Non-Zero Default" Trap in TF Schemas

A critical trap occurs when a Terraform schema sets a **non-zero default** (e.g. `Default: true` on `schema.TypeBool`):

1. **User specifies `false`:** The user writes `spec.field: false` in YAML.
2. **Terraform evaluates the input:** Because Terraform SDK v2 collapses `false` into the zero-value/unset category, Terraform's field reader treats it as an omitted field.
3. **Terraform applies the default:** Because the field appears omitted, Terraform overwrites it with `Default: true`.
4. **Fatal Result (Loss of Opt-Out):** **It becomes impossible for any user to configure, maintain, or opt out of `false` on GCP**, because Terraform permanently forces it to `true`. Setting `field: false` or omitting the field both result in `true`.

> **Rule:** Never add `Default: true` (or any non-zero default) to optional primitive fields in Terraform provider schemas.

## 4. Consequences of Updating a Field to Default to True

When an optional boolean field in a Terraform schema is changed to `Default: true`, it introduces severe backward-compatibility and drift issues for existing KCC workloads:

1. **Active Overwrite of Explicit User Intent (`spec.field: false`):**
   * Any user who explicitly specifies `false` in their KCC YAML will find that KCC silently flips the field to `true` on GCP during reconciliation.
   * Because TF SDK v2 collapses `false` into the zero-value/omitted bucket, the schema default overrides the user configuration unconditionally.
2. **Unintended In-Place Mutations on Existing Resources (Even if Unmanaged):**
   * Existing resources on GCP that currently have the feature disabled (`false`) will be mutated to `true` during the next reconciliation.
   * **Why this happens despite KCC SSA overlay:**
     * For unmanaged fields (omitted in YAML), KCC's SSA layer copies the live state (`false`) into `config` to prevent diffs.
     * However, when Terraform SDK v2 calculates the plan via `TFResource.Diff()`, it evaluates `false` in `config` as an omitted zero-value and applies `Default: true` to the desired state.
     * Terraform then detects a diff between `liveState` (`false`) and desired state (`true`), marking `d.HasChange("field") == true`.
     * The controller proceeds to `TFResource.Apply()`, issuing an update to GCP to enable the feature (`true`). This can trigger service disruption, configuration changes, or costly node pool updates.

---

## 5. Analysis of `d.Get(...)` and `d.GetOk(...)`

### Approach A: `d.Get(...)` (Current/Chosen)

* **Mechanism:** `d.Get` returns the zero value (`false`) if the field is unset in the Terraform configuration.
* **Scenario 1: User sets `true` in KCC:**
  * KCC passes `true` to TF.
  * `d.Get` returns `true`.
  * **Result:** **CORRECT.** Sends `true` to API.
* **Scenario 2: User sets `false` in KCC:**
  * KCC unsets the field in TF (Zero Value behavior).
  * `d.Get` returns `false`.
  * **Result:** **CORRECT.** The user explicitly asked for `false`, and we send `false` to the API.
* **Scenario 3: User omits field (Update):**
  * KCC passes live value to TF via SSA overlay.
  * If live is `true` -> `d.Get` returns `true` -> Sends `true`.
  * If live is `false` -> `d.Get` returns `false` -> Sends `false`.
  * **Result:** **CORRECT.** Preserves live state.
* **Scenario 4: User omits field (Create):**
  * Field is unset in TF.
  * `d.Get` returns `false`.
  * **Result:** **INCORRECT.** Sends `false` to API. If the API default is `true`, this forces `false` on GCP instead of letting the API apply its default.

### Approach B: `d.GetOk(...)` (Alternative)

* **Mechanism:** `d.GetOk` returns `ok=false` if the field is unset in the Terraform configuration or is a zero-value (`false`).
* **Scenario 1: User sets `true` in KCC:**
  * KCC passes `true` to TF.
  * `d.GetOk` returns `ok=true`, `value=true`.
  * **Result:** **CORRECT.** Sends `true` to API.
* **Scenario 2: User sets `false` in KCC:**
  * KCC unsets the field in TF (Zero Value behavior).
  * `d.GetOk` returns `ok=false`.
  * Block is skipped. Field is omitted from API request.
  * **Result:** **INCORRECT.** The user explicitly asked for `false`, but we sent nothing. If the API default is `true` (or if we needed to overwrite a previous `true`), this fails.
* **Scenario 3: User omits field (Update):**
  * KCC passes live value to TF.
  * If live is `true` -> `d.GetOk` returns `ok=true`, `value=true` -> Sends `true`.
  * If live is `false` -> `d.GetOk` returns `ok=false` -> Block is skipped.
  * **Result:** **CORRECT.** Preserves live state.
* **Scenario 4: User omits field (Create):**
  * Field is unset.
  * `d.GetOk` returns `ok=false` -> Block is skipped. Field is omitted from API request.
  * **Result:** **CORRECT.** Field is omitted, allowing the GCP API to apply its default.

---

## 6. Decision

**Use `d.Get(...)` (along with `d.HasChange(...)` for updates) for optional boolean field expansion.**

### Rationale:
1. **Correctness for Zero Values:** Due to KCC's behavior of mapping `false` to "unset" in Terraform, `d.Get(...)` is the **only** mechanism that allows us to correctly enforce a user's explicit request for `false`. `d.GetOk(...)` silently ignores the user's intent.
2. **Config Management Practice:** Users managing infrastructure via KCC often define their desired state declaratively. If they want a feature off, they set it to `false`. Ensuring `false` (zero value) works reliably covers the primary use case for managing these resources.
3. **Security:** In the creation scenario where the user omits the field, defaulting to sending `false` (via `d.Get`'s default) is usually a safer behavior for a security-sensitive field.

---

## 7. Long-Term Architectural Solution: Direct Controllers

The permanent solution to the zero-value ambiguity is KCC's migration to **Direct Controllers**:

* Direct controllers define API fields using Go pointers (`*bool`, `*int64`, `*string`):
  * **`nil`:** Unset / unmanaged (let GCP server-default without asserting a value).
  * **`&false`:** Explicitly managed `false`.
  * **`&true`:** Explicitly managed `true`.
* This eliminates the zero-value collapse, avoids schema defaulting traps, and provides clean 3-state declarative management.