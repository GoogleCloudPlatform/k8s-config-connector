// Copyright 2026 Google LLC
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

package v1alpha1

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ServiceDirectoryServiceRef struct {
	// A reference to an externally managed ServiceDirectoryService resource.
	// Should be in the format "projects/*/locations/*/namespaces/*/services/*".
	External string `json:"external,omitempty"`

	// The name of a ServiceDirectoryService resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ServiceDirectoryService resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ServiceDirectoryServiceRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	ref := v1alpha1.ServiceDirectoryServiceRef{
		External:  r.External,
		Name:      r.Name,
		Namespace: r.Namespace,
	}
	ext, err := ref.NormalizedExternal(ctx, reader, otherNamespace)
	r.External = ref.External
	return ext, err
}
