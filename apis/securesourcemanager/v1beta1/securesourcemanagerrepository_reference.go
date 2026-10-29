// Copyright 2024 Google LLC
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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ refsv1beta1.ExternalNormalizer = &SecureSourceManagerRepositoryRef{}
	_ refsv1beta1.Ref                = &SecureSourceManagerRepositoryRef{}
	_ refsv1beta1.ExternalRef        = &SecureSourceManagerRepositoryRef{}
)

func init() {
	refsv1beta1.Register(&SecureSourceManagerRepositoryRef{}, &SecureSourceManagerRepository{})
}

// SecureSourceManagerRepositoryRef is a reference to a SecureSourceManagerRepository.
type SecureSourceManagerRepositoryRef struct {
	// A reference to an externally managed SecureSourceManagerRepository resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/repositories/{{repositoryID}}".
	External string `json:"external,omitempty"`

	// The name of a SecureSourceManagerRepository resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SecureSourceManagerRepository resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *SecureSourceManagerRepositoryRef) GetGVK() schema.GroupVersionKind {
	return SecureSourceManagerRepositoryGVK
}

func (r *SecureSourceManagerRepositoryRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SecureSourceManagerRepositoryRef) GetExternal() string {
	return r.External
}

func (r *SecureSourceManagerRepositoryRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *SecureSourceManagerRepositoryRef) ValidateExternal(ref string) error {
	_, err := parseSecureSourceManagerRepositoryExternal(ref)
	return err
}

func (r *SecureSourceManagerRepositoryRef) ParseExternalToIdentity() (identity.Identity, error) {
	id, err := parseSecureSourceManagerRepositoryExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SecureSourceManagerRepositoryRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

// NormalizedExternal provision the "External" value for other resource that depends on SecureSourceManagerRepository.
// If the "External" is given in the other resource's spec.SecureSourceManagerRepositoryRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual SecureSourceManagerRepository object from the cluster.
func (r *SecureSourceManagerRepositoryRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", SecureSourceManagerRepositoryGVK.Kind)
	}
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

func parseSecureSourceManagerRepositoryExternal(external string) (*RepositoryIdentity, error) {
	external = strings.TrimPrefix(external, "//securesourcemanager.googleapis.com/")
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "repositories" {
		return nil, fmt.Errorf("format of SecureSourceManagerRepository external=%q was not known (use projects/{{projectId}}/locations/{{location}}/repositories/{{repositoryID}})", external)
	}
	return &RepositoryIdentity{
		parent: &RepositoryParent{ProjectID: tokens[1], Location: tokens[3]},
		id:     tokens[5],
	}, nil
}
