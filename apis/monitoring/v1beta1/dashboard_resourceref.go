// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

type LogsPanelResourceRef struct {
	/* The external name of the referenced resource */
	External string `json:"external,omitempty"`
	/* Kind of the referent. */
	Kind string `json:"kind,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

var _ refs.Ref = &LogsPanelResourceRef{}

func (r *LogsPanelResourceRef) GetExternal() string {
	return r.External
}

func (r *LogsPanelResourceRef) SetExternal(ref string) {
	r.External = ref
}

// GetGVK returns the schema.GroupVersionKind of the reference type
func (r *LogsPanelResourceRef) GetGVK() schema.GroupVersionKind {
	ref, err := r.resolveRef()
	if err != nil {
		klog.Warningf("unable to resolve GVK for LogsPanelResourceRef: %v", err)
		return schema.GroupVersionKind{}
	}
	return ref.GetGVK()
}

func (r *LogsPanelResourceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *LogsPanelResourceRef) ValidateExternal(external string) error {
	ref, err := r.resolveRef()
	if err != nil {
		return err
	}

	if err := ref.ValidateExternal(external); err != nil {
		return err
	}
	return nil
}

func (r *LogsPanelResourceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	ref, err := r.resolveRef()
	if err != nil {
		return err
	}

	if err := ref.Normalize(ctx, reader, defaultNamespace); err != nil {
		return err
	}
	r.External = ref.GetExternal()
	return nil
}

// resolveRef resolves the LogsPanelResourceRef into a specific Ref implementation.
func (r *LogsPanelResourceRef) resolveRef() (refs.Ref, error) {
	kind := r.Kind

	// For backwards compatibility, infer "Project" kind
	if kind == "" && r.External != "" {
		tokens := strings.Split(r.External, "/")
		if len(tokens) == 2 && tokens[0] == "projects" {
			kind = "Project"
		}
	}

	if kind == "" {
		return nil, fmt.Errorf("must specify kind on reference (%+v)", r)
	}

	switch kind {
	case "Project":
		return &refs.ProjectRef{
			Name:      r.Name,
			Namespace: r.Namespace,
			External:  r.External,
			Kind:      kind,
		}, nil

	default:
		return nil, fmt.Errorf("references to kind %q are not supported", kind)
	}
}
