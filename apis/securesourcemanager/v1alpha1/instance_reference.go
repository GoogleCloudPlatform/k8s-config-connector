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
	"strconv"
	"strings"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &SecureSourceManagerInstanceRef{}

// SecureSourceManagerInstanceRef defines the resource reference to SecureSourceManagerInstance, which "External" field
// holds the GCP identifier for the KRM object.
type SecureSourceManagerInstanceRef struct {
	// A reference to an externally managed SecureSourceManagerInstance resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	External string `json:"external,omitempty"`

	// The name of a SecureSourceManagerInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SecureSourceManagerInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

func ParseSecureSourceManagerInstanceUrl(url string) (*InstanceIdentity, error) {
	id, err := parseSecureSourceManagerInstanceExternal(url)
	if err != nil {
		return nil, err
	}
	return &InstanceIdentity{
		parent: id.parent,
		id:     id.id,
	}, nil
}

func parseSecureSourceManagerInstanceExternal(external string) (*InstanceIdentity, error) {
	s := external
	s = strings.TrimPrefix(s, "/")
	s = strings.TrimPrefix(s, "securesourcemanager.googleapis.com/")

	tokens := strings.Split(s, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		return &InstanceIdentity{
			parent: &InstanceParent{ProjectID: tokens[1], Location: tokens[3]},
			id:     tokens[5],
		}, nil
	}

	return nil, fmt.Errorf("format of SecureSourceManagerInstance external=%q was not known (use projects/{{projectId}}/locations/{{location}}/instances/{{instanceID}})", external)
}

// ConvertToProjectNumber converts the external reference to use a project number.
func (r *SecureSourceManagerInstanceRef) ConvertToProjectNumber(ctx context.Context, projectsClient *resourcemanager.ProjectsClient) error {
	if r == nil {
		return nil
	}
	instanceIdentity, err := parseSecureSourceManagerInstanceExternal(r.External)
	if err != nil {
		return err
	}
	id := instanceIdentity.id
	parent := instanceIdentity.parent

	// Check if the project number is already a valid integer
	// If not, we need to look it up
	projectNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		req := &resourcemanagerpb.GetProjectRequest{
			Name: "projects/" + parent.ProjectID,
		}
		project, err := projectsClient.GetProject(ctx, req)
		if err != nil {
			return fmt.Errorf("error getting project %q: %w", req.Name, err)
		}
		n, err := strconv.ParseInt(strings.TrimPrefix(project.Name, "projects/"), 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing project number for %q: %w", project.Name, err)
		}
		projectNumber = n
	}
	parent.ProjectID = strconv.FormatInt(projectNumber, 10)
	r.External = fmt.Sprintf("%s/instances/%s", parent.String(), id)
	return nil
}

// NormalizedExternal provision the "External" value for other resource that depends on SecureSourceManagerInstance.
// If the "External" is given in the other resource's spec.SecureSourceManagerInstanceRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual SecureSourceManagerInstance object from the cluster.
func (r *SecureSourceManagerInstanceRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", SecureSourceManagerInstanceGVK.Kind)
	}
	// From a referenced config connector object
	if r.External == "" {
		if r.Namespace == "" {
			r.Namespace = otherNamespace
		}
		key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(SecureSourceManagerInstanceGVK)
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
			}
			return "", fmt.Errorf("reading referenced %s %s: %w", SecureSourceManagerInstanceGVK, key, err)
		}
		// Get external from status.externalRef. This is the most trustworthy place.
		actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
		if err != nil {
			return "", fmt.Errorf("reading status.externalRef: %w", err)
		}
		if actualExternalRef == "" {
			return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		r.External = actualExternalRef
	}
	return r.External, nil
}
