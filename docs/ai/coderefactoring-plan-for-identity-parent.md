# Code Refactoring Plan: Standardizing Resource Identity and Parentage

This document outlines a plan to refactor all resources under the `apis/` directory to adopt a new, standardized pattern for handling resource identity, parentage, and references. The goal is to improve consistency, maintainability, and clarity across the codebase.

The `apis/bigquerybiglake/v1alpha1/` implementation serves as the canonical example for this new pattern.

## High-Level Goal

The primary objective is to refactor every resource to follow a consistent pattern for:
1.  **Identity Management**: How a resource's unique GCP identity is constructed and parsed.
2.  **Parent Hierarchy**: How a resource declares its parent in the GCP resource hierarchy.
3.  **Resource References**: How other resources can create dependencies and references.

## The Refactoring Pattern

For each resource (e.g., `FooBar`), the following file structure and components will be implemented:

1.  **`foo_bar_identity.go`**:
    *   Defines a `FooBarIdentity` struct that implements the `identity.Identity` interface.
    *   This struct is responsible for parsing and constructing the full GCP resource identifier string (e.g., `projects/{{p}}/locations/{{l}}/fooBars/{{id}}`).
    *   The main `FooBar` resource struct (in `foo_bar_types.go`) will implement the `identity.Resource` interface by providing a `GetIdentity()` method. This method resolves the resource's identity from its `spec` and `status`.

2.  **`foo_bar_reference.go`**:
    *   Defines a `FooBarRef` struct that implements the `refsv1beta1.Ref` interface.
    *   This struct provides a standardized way for other resources to reference a `FooBar` resource, supporting both internal (K8s `name`/`namespace`) and external (full GCP ID) references.

3.  **`foo_bar_types.go`**:
    *   The `FooBarSpec` within this file will be updated to use the new reference structs for defining its parent.
    *   For top-level resources (e.g., parent is a Project), the spec will embed a standard parent type like `parent.ProjectRef` or `parent.ProjectAndLocationRef`.
    *   For nested resources, the spec will contain a field for its parent's reference struct (e.g., `ParentFooBarRef *v1beta1.FooBarRef`).
    *   The `resourceID` field will be used consistently for the user-specified short name of the resource, falling back to `metadata.name`.

## Step-by-Step Refactoring Process

The refactoring will be applied to each resource, one at a time, following these steps:

1.  **Select a Resource**: Choose a resource to refactor (e.g., `IAMServiceAccount`).
2.  **Create `_identity.go`**: Create a new file (e.g., `apis/iam/v1beta1/iamserviceaccount_identity.go`). Implement the `IAMServiceAccountIdentity` struct and the `GetIdentity()` method for the `IAMServiceAccount` resource.
3.  **Create `_reference.go`**: Create a new file (e.g., `apis/iam/v1beta1/iamserviceaccount_reference.go`). Implement the `IAMServiceAccountRef` struct.
4.  **Update `_types.go`**:
    *   Locate the existing `..._types.go` file for the resource. If it's in a generated location (e.g., `pkg/clients/generated/`), copy it to the correct `apis/` path (e.g., `apis/iam/v1beta1/iamserviceaccount_types.go`).
    *   Modify the resource's `Spec` to replace the old parent declaration with the new `Ref` struct (e.g., add `ProjectRef *parent.ProjectRef`).
    *   Ensure the file has the necessary imports for the `parent` and `refs` packages.
5.  **Cleanup**: Once the new, manually crafted files are in place, any corresponding auto-generated files for that resource should be removed to avoid conflicts.
6.  **Repeat**: Apply this process systematically to all other resources in the `apis/` directory.

## Concrete Example: Refactoring `BigQueryBigLakeCatalog`

Here is how the plan would be applied to the `BigQueryBigLakeCatalog` resource.

### 1. Create `catalog_identity.go`

```go
// in apis/bigquerybiglake/v1alpha1/catalog_identity.go
package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &CatalogIdentity{}

const (
	CatalogIDURL = parent.ProjectAndLocationURL + "/catalogs/{{catalogID}}"
)

type CatalogIdentity struct {
	parent *parent.ProjectAndLocationParent
	id     string
}

func (i *CatalogIdentity) String() string {
	return i.parent.String() + "/catalogs/" + i.id
}

// ... (rest of the implementation) ...

var _ identity.Resource = &BigLakeCatalog{}

func (obj *BigLakeCatalog) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// ... (implementation to resolve identity from spec) ...
}
```

### 2. Create `catalog_reference.go`

```go
// in apis/bigquerybiglake/v1alpha1/catalog_reference.go
package v1alpha1

import (
	"context"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &BigQueryBigLakeCatalogRef{}

// BigQueryBigLakeCatalogRef is a reference to a BigQueryBigLakeCatalog resource.
type BigQueryBigLakeCatalogRef struct {
	// A reference to an externally managed BigQueryBigLakeCatalog resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/catalogs/{{catalogID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryBigLakeCatalog resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryBigLakeCatalog resource.
	Namespace string `json:"namespace,omitempty"`
}

// ... (implementation of refsv1beta1.Ref interface) ...
```

### 3. Update `catalog_types.go`

The `BigLakeCatalogSpec` would be defined with an explicit parent reference:

```go
// in apis/bigquerybiglake/v1alpha1/catalog_types.go
import (
    "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
    // ... other imports
)

// BigLakeCatalogSpec defines the desired state of BigLakeCatalog
type BigLakeCatalogSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The BigLakeCatalog name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}
```
