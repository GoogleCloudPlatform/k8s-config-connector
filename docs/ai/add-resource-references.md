This document provides instructions on how to handle API resource references in KCC.

## Background

In KCC, a resource can reference another resource. For example, a `StorageBucket` might have a `kmsKeyName` field that references a `KMSCryptoKey` resource. When KCC sees such a reference, it needs to resolve the reference to the fully-qualified GCP resource name.

The way KCC handles this is by defining a "reference" object. This object is a struct that implements the `refsv1beta1.Ref` interface. This interface has methods that allow the generic normalization logic to resolve the reference.

The reference object is typically placed in its own file, `apis/<service>/<version>/*_reference.go`. For example, the reference object for `BigQueryBigLakeTable` is in `apis/bigquerybiglake/v1alpha1/table_reference.go`.

## Task

Your task is to identify fields in KCC's API that are references to other GCP resources, and change them to use the KCC reference object, following the established conventions.

### Identifying Reference Fields

There are two cases to consider:

1.  **Greenfield resources:** These are resources that are being added to KCC for the first time. In this case, you should look at the comments for each field in the API. If a field's comment indicates that it is a reference to another GCP resource, then you should mark it as a reference by adding the `//+kcc:ref={Your guess of its KCC kind}` annotation.

2.  **Terraform/DCL-migrated resources:** These are resources that are being migrated from Terraform or DCL to KCC. In this case, you should look at the `config/crds` differences. The field may be different by ending with `*Ref`, and is a structure with fields like "name", "namespace", and/or "external", rather than a string.

### Changing a Field to a Reference

Once you have identified a reference field, you need to change its definition in the Go struct to follow the naming convention required by the `mapper-generator`.

-   **Field Name:** The KRM field name **must** be the `UpperCamelCase` version of the original GCP proto field name, with a `Ref` suffix. For a proto field `kms_key_name`, the KRM field name must be `KmsKeyNameRef`.
-   **Field Type:** The Go type for this field must be a pointer to the reference struct, e.g., `*KMSCryptoKeyRef`.

For example, if you have a proto field `kms_key_name` that is a reference to a `KMSCryptoKey` resource, you should change the Go struct definition from `KmsKeyName *string` to:

```go
KmsKeyNameRef *KMSKeyKeyRef `json:"kmsKeyNameRef,omitempty"`
```

If the reference object (e.g., `KMSCryptoKeyRef`) does not exist, you should create one. You can use `apis/bigquerybiglake/v1alpha1/table_reference.go` as a template.

Here is an example of a reference object:

```go
// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &BigQueryBigLakeTableRef{}

// BigQueryBigLakeTableRef is a reference to a BigQueryBigLakeTable resource.
type BigQueryBigLakeTableRef struct {
	// A reference to an externally managed BigQueryBigLakeTable resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/catalogs/{{catalogID}}/databases/{{databaseID}}/tables/{{tableID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryBigLakeTable resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryBigLakeTable resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *BigQueryBigLakeTableRef) GetGVK() schema.GroupVersionKind {
	return BigLakeTableGVK
}

func (r *BigQueryBigLakeTableRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BigQueryBigLakeTableRef) GetExternal() string {
	return r.External
}

func (r *BigQueryBigLakeTableRef) SetExternal(ref string) {
	r.External = ref
}

func (r *BigQueryBigLakeTableRef) ValidateExternal(ref string) error {
	id := &TableIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *BigQueryBigLakeTableRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}
```