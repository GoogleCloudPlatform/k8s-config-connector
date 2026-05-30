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
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &SecurityProfileRef{}

// SecurityProfileRef defines the resource reference to NetworkSecuritySecurityProfile, which "External" field
// holds the GCP identifier for the KRM object.
type SecurityProfileRef struct {
	// A reference to an externally managed NetworkSecuritySecurityProfile resource.
	// Should be in the format `projects/{{projectID}}/locations/{{location}}/securityProfiles/{{securityProfile}}`.
	External string `json:"external,omitempty"`

	// The `name` field of a NetworkSecuritySecurityProfile resource.
	Name string `json:"name,omitempty"`

	// The `namespace` field of a NetworkSecuritySecurityProfile resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provisions the "External" value for other resource that depends on SecurityProfile.
// If the "External" is given in the other resource's spec.SecurityProfileRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual SecurityProfile object from the cluster.
func (r *SecurityProfileRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", NetworkSecuritySecurityProfileGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, match, _ := securityProfileProjectURL.Parse(r.External); !match {
			if _, matchOrg, _ := securityProfileOrganizationURL.Parse(r.External); !matchOrg {
				return "", fmt.Errorf("format of NetworkSecuritySecurityProfile external=%q was not known", r.External)
			}
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &NetworkSecuritySecurityProfile{}
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", NetworkSecuritySecurityProfileGVK.Kind, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef := u.Status.ExternalRef
	if actualExternalRef == nil || *actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	return *actualExternalRef, nil
}
