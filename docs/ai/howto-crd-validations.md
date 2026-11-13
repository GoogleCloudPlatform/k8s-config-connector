# How to Apply CRD Validations

This guide explains how to apply CRD-level validations to Config Connector resource fields using `+kubebuilder` markers in the resource's `_types.go` file. These validations are fast, client-side checks that provide immediate feedback to the user.

---

## Rule 1: Required vs. Optional Fields

-   **Required fields** must have the `// +required` marker.
-   **Optional fields** should not have the `+required` marker and must include `omitempty` in their JSON tag.

**Example:** From `apis/cloudbuild/v1beta1/workerpool_types.go`:

```go
// +kcc:proto=google.devtools.cloudbuild.v1.WorkerPool
type CloudBuildWorkerPoolSpec struct {
	// ...

	// Legacy Private Pool configuration.
	// +required
	PrivatePoolConfig *PrivatePoolV1Config `json:"privatePoolV1Config,omitempty"`

    // A user-specified, human-readable name for the `WorkerPool`. If provided,
	//  this value must be 1-63 characters.
	DisplayName string `json:"displayName,omitempty"`

	// ...
}
```

## Rule 2: Immutable Fields

For **required** fields, you can enforce immutability at the CRD level using a CEL rule. Do not use this for optional fields, as it prevents setting a value if it was previously unset.

**Example:** From `apis/cloudbuild/v1beta1/workerpool_types.go`:

```go
// +kcc:spec:proto=google.devtools.cloudbuild.v1.PrivatePoolV1Config.NetworkConfig
type PrivatePoolV1Config_NetworkConfigSpec struct {
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
	// Immutable. The network definition that the workers are peered to.
	PeeredNetworkRef refv1beta1.ComputeNetworkRef `json:"peeredNetworkRef,omitempty"`

    // ...
}
```
*Note: For most immutable fields, especially optional ones, the validation is handled by the controller at runtime. This CRD-level validation is only for required fields where immediate feedback is desired.*

## Rule 3: Parent Fields

Parent fields (like `projectRef`, `location`, etc.) identify the parent resource in GCP. They should be **required** and are considered immutable by the controller.

-   **Required:** Group parent fields in an inlined struct with the `+required` marker.
-   **Immutable:** Immutability is enforced by the controller, not at the CRD level.

**Example:** From `apis/securesourcemanager/v1beta1/instance_types.go`:

```go
type SecureSourceManagerInstanceSpec struct {
	Parent `json:",inline"`
    // ...
}

type Parent struct {
	/* Immutable. The Project that this resource belongs to. */
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	/* Immutable. Location of the instance. */
	// +required
	Location string `json:"location"`
}
```

**Nested parent:**
