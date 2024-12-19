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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type OrganizationRef struct {
	/* The Organization selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of an `Organization` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of an `Organization` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type Organization struct {
	OrganizationName string
}

func (s *Organization) FullyQualifiedName() string {
	return "organizations/" + s.OrganizationName
}
func ResolveOrganization(ctx context.Context, reader client.Reader, src client.Object, ref *OrganizationRef) (*Organization, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on organization reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 1 {
			return &Organization{OrganizationName: tokens[0]}, nil
		}
		if len(tokens) == 2 && tokens[0] == "organizations" {
			return &Organization{OrganizationName: tokens[1]}, nil
		}
		return nil, fmt.Errorf("format of organization external=%q was not known (use organization/<orgId> or <orgId>)", ref.External)

	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on organization reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	organization := &unstructured.Unstructured{}
	organization.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "apigee.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "ApigeeOrganization",
	})

	if err := reader.Get(ctx, key, organization); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(organization.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced Organization %v: %w", key, err)
	}

	orgID, err := GetResourceID(organization)

	if err != nil {
		return nil, err
	}

	return &Organization{OrganizationName: orgID}, nil
}
