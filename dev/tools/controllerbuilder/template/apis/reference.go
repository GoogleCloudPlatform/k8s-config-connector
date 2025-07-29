package apis

const ReferenceTemplate = `// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package {{ .Version }}

import (
	"context"

    refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// {{.ProtoMessageName}}Ref defines the resource reference to {{.Kind}}, which "External" field
// holds the GCP identifier for the KRM object.
type {{.ProtoMessageName}}Ref struct {
	// A reference to an externally managed {{.Kind}} resource.
    // TODO(contributor): Update format to the actual valid format of your case
	// Should be in the format "projects/{{"{{"}}projectID{{"}}"}}/locations/{{"{{"}}location{{"}}"}}/{{.ProtoMessageName | ToLower }}s/{{"{{"}}{{.ProtoMessageName | ToLower }}ID{{"}}"}}".
	External string ` + "`" + `json:"external,omitempty"` + "`" + `

	// The name of a {{.Kind}} resource.
	Name string ` + "`" + `json:"name,omitempty"` + "`" + `

	// The namespace of a {{.Kind}} resource.
	Namespace string ` + "`" + `json:"namespace,omitempty"` + "`" + `
}

// TODO(contributor): Choose one of the two options below.

// --------------------------------------------------------------------------------
// Option 1: For legacy resources requiring backward compatibility.
// Your resource is migrated from legacy controller, you must implement custom normalization logic.
// --------------------------------------------------------------------------------
var _ refsv1beta1.ExternalNormalizer = &{{.ProtoMessageName}}Ref{}

// NormalizedExternal provision the "External" value for other resource that depends on {{.Kind}}.
// If the "External" is given in the other resource's spec.{{.Kind}}Ref, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual {{.Kind}} object from the cluster.
func (r *{{.ProtoMessageName}}Ref) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	// TODO(contributor): write your custom normalization logic here. Please consult with the KCC team if you need guidance on the implementation.
}

// --------------------------------------------------------------------------------
// Option 2: For pure direct resources.
// Your resource is a brand new resource that only managed by direct controller
// --------------------------------------------------------------------------------
var _ reference.Reference = &{{.ProtoMessageName}}Ref{}

func (r *{{.ProtoMessageName}}Ref) GetGVK() schema.GroupVersionKind {
	return {{.Kind}}GVK
}

func (r *{{.ProtoMessageName}}Ref) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *{{.ProtoMessageName}}Ref) GetExternal() string {
	return r.External
}

func (r *{{.ProtoMessageName}}Ref) SetExternal(ref string) {
	r.External = ref
}

func (r *{{.ProtoMessageName}}Ref) ValidateExternal() error {
	if _, _, err := Parse{{.ProtoMessageName}}External(r.External); err != nil {
		return err
	}
	return nil
}

func (r *{{.ProtoMessageName}}Ref) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return reference.Normalize(ctx, reader, r, defaultNamespace)
}
`
