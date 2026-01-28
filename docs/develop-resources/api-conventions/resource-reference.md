# Direct Resource Reference Guide

## TL;DR

Referencing GCP objects should define an API field that ends with `Ref` or `Refs` to indicate a dependency on another GCP resource. This document outlines the conventions for these reference fields, ensuring consistency and enabling automation via tooling.

## Naming Conventions

The naming of resource reference fields follows a clear and consistent pattern, mapping from the GCP proto field name to the KCC KRM field name and type. This convention is **required** for the `mapper-generator` to function correctly.

### From GCP Proto to KRM

-   **GCP Proto Field:** The original field name in the GCP proto definition (e.g., `crypto_key_name`). This is typically in `snake_case`.
-   **Go Field Name:** The field name in the Go struct for the KCC CRD. This **must** be the `UpperCamelCase` version of the GCP proto field name, with a `Ref` suffix (e.g., `CryptoKeyNameRef`). This strict convention is required for the `mapper-generator` to work correctly.
-   **Go Field Type:** The Go type of the KRM field. This should be a pointer to a reference struct, named `*<Kind>Ref` (e.g., `*KMSCryptoKeyRef`), where `Kind` is the KCC Kind of the referenced resource.
-   **JSON Tag:** The `json` tag defines the field name in the Kubernetes resource's YAML/JSON representation. This name can differ from the Go field name. This is particularly useful for maintaining backward compatibility when migrating existing Beta resources. For new resources, the JSON tag should be the `camelCase` version of the Go field name (e.g., `cryptoKeyNameRef`).

**Example:**

A GCP proto has a field `crypto_key_name`. In a previous version of the KCC resource, this was exposed as a string field named `cryptoKeyName`. To migrate this to a proper reference object while maintaining backward compatibility, the new field would be:

```go
// CryptoKeyNameRef is a reference to a KMSCryptoKey.
// The Go field name 'CryptoKeyNameRef' follows the convention for the mapper-generator.
// The json tag 'cryptoKeyName' is kept for backward compatibility with the old CRD.
CryptoKeyNameRef *KMSCryptoKeyRef `json:"cryptoKeyName,omitempty"`
```

Here is a table summarizing the convention:

| GCP Proto Field (`snake_case`) | Go Field Name (`UpperCamelCase` + `Ref`) | Go Field Type (`*<Kind>Ref`) | Referenced KCC Kind | JSON tag (`camelCase`)                               |
| ------------------------------ | ---------------------------------------- | ---------------------------- | ------------------- | ---------------------------------------------------- |
| `crypto_key_name`              | `CryptoKeyNameRef`                       | `*KMSCryptoKeyRef`           | `KMSCryptoKey`      | `cryptoKeyNameRef` (new) or `cryptoKeyName` (compat) |
| `network`                      | `NetworkRef`                             | `*ComputeNetworkRef`         | `ComputeNetwork`    | `networkRef` (new) or `network` (compat)             |

### Pluralization

If a resource can reference multiple resources of the same Kind, the field name should be pluralized by changing the `Ref` suffix to `Refs` (e.g., `ProjectRefs`). The type will be a slice of pointers to the reference struct (e.g., `[]*ProjectRef`).

## API Rule

The struct for a reference field type (e.g., `KMSCryptoKeyRef`) must implement the `refsv1beta1.Ref` interface. This enables consistent handling of reference resolution and validation.

```go
// apis/refs/v1beta1/interface.go

type Ref interface {
	GetGVK() schema.GroupVersionKind
	GetNamespacedName() types.NamespacedName
	GetExternal() string
	SetExternal(ref string)
	ValidateExternal(ref string) error
	Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error
}
```

The reference struct itself should be defined as follows:

```go
type <Kind>Ref struct {
  // +optional
  External string   `json:"external,omitempty"`
  // +optional
  Name string       `json:"name,omitempty"`
  // +optional
  Namespace string  `json:"namespace,omitempty"`
}
```

## Type and Mapper Generation

The `type-generator` and `mapper-generator` tools enforce these conventions and automate the creation of mapper functions. These tools will:

1.  **Identify reference fields** in the GCP protos.
2.  **Determine the corresponding KCC Kind** for the referenced resource (often via a `//+kcc:ref` annotation).
3.  **Validate that the KRM field name** matches the `UpperCamelCase(proto_field_name) + "Ref"` convention.
4.  **Generate the mapper function** that uses the `refsv1beta1.Ref` interface to resolve the reference.

This automation reduces boilerplate code and ensures that all reference fields are handled consistently.

## Validation

### Rule 1: Config Connector Level Only

Reference validation can be either CRD validation or a Config Connector controller check. It should not require GCP calls.

### Rule 2: Required Fields

- If the reference does not have a corresponding Config Connector Kind yet, the `.<Kind>Ref.external` is required. Note: the `.<Kind>Ref` itself can be optional.
- If the reference has a corresponding Config Connector Kind, the `<Kind>Ref.external` and `<Kind>Ref.name` are `oneOf` required. Note: the `.<Kind>Ref` itself can be optional.

### Rule 3: External

The `external` field should be in the format of the asset inventory without the service domain (e.g., `projects/<projectID>/global/networks/<networkID>`).

### Rule 4: Namespace

- If the referenced Config Connector object is cluster-scoped or in the same namespace, the `namespace` field is optional.
- If the referenced Config Connector object is in a different namespace, the `namespace` field is **required**.

### Rule 5: Errors

Use the `k8s.ReferenceNotFoundError` when the referenced Config Connector object is not found.

## Same Kind References

For lists of references of the same kind:

- **Rule 1: Avoid Mixed Forms:** Use either `external` or `name` for all references in a list, not a mix of both.
- **Rule 2: Form Switching:** Users should be able to switch between using `external` and `name`.
- **Rule 3: Unique `external`:** Validate the uniqueness of `external` string values.
- **Rule 4: Non-unique `name`:** Do not validate the uniqueness of `name`/`namespace` pairs. Instead, check the uniqueness of the GCP resources referenced by the `externalRef` field of the corresponding Config Connector objects.
- **Rule 5: Ordering:** Preserve the order of references as defined in the CR, unless the GCP service prefers a sorted order.

## Code Style

Place the code for a reference struct in a `<kind>_reference.go` file under `apis/<service>/<version>/`.

## Backward Compatibility

For TF-based or DCL-based Beta resources, maintain the original CRD and behavior when migrating to a Direct Resource. Address ambiguous `external` usage and generic references with `Kind` on a case-by-case basis, with a preference for migrating to the new reference structure.
