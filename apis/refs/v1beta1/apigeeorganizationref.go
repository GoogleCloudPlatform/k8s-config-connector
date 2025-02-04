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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ApigeeGroupVersion    = schema.GroupVersion{Group: "apigeeorganizations.apigee.cnrm.cloud.google.com", Version: "v1beta1"}
	ApigeeOrganizationGVK = ApigeeGroupVersion.WithKind("ApigeeOrganization")
)

var _ ExternalNormalizer = &ApigeeOrganizationRef{}

type ApigeeOrganizationRef struct {
	/* The ApigeeOrganization selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of an `ApigeeOrganization` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of an `ApigeeOrganization` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type ApigeeOrganization struct {
	Name string
}

func (o *ApigeeOrganization) FullyQualifiedName() string {
	return "organizations/" + o.Name
}

func (r *ApigeeOrganizationRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ApigeeOrganizationGVK.Kind)
	}

	// From given External
	if r.External != "" {
		external := strings.TrimPrefix(r.External, "/")
		tokens := strings.Split(external, "/")

		if len(tokens) == 2 && tokens[0] == "organizations" {
			return r.External, nil
		}
		return "", fmt.Errorf("format of Organization external=%q was not known (use organizations/{{orgId}})", external)
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ApigeeOrganizationGVK)

	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("error reading referenced %s %s: %w", ApigeeOrganizationGVK, key, err)
	}

	// get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}

	r.External = actualExternalRef
	return r.External, nil
}

func ParseApigeeOrganizationExternal(external string) (resourceID string, err error) {
	tokens := strings.Split(external, "/")

	if len(tokens) != 2 || tokens[0] != "organizations" {
		return "", fmt.Errorf("format of ApigeeOrganization external=%q was not known (use organizations/{{organization}})",
			external)
	}
	resourceID = tokens[1]
	return resourceID, nil
}
