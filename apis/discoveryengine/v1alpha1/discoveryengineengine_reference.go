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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &DiscoveryEngineEngineRef{}
var _ refsv1beta1.ExternalRef = &DiscoveryEngineEngineRef{}
var _ refsv1beta1.ExternalNormalizer = &DiscoveryEngineEngineRef{}

// DiscoveryEngineEngineRef is a reference to a DiscoveryEngineEngine.
type DiscoveryEngineEngineRef struct {
	// A reference to an externally managed DiscoveryEngineEngine resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/collections/{{collectionID}}/engines/{{engineID}}".
	External string `json:"external,omitempty"`

	// The name of a DiscoveryEngineEngine resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DiscoveryEngineEngine resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&DiscoveryEngineEngineRef{}, &DiscoveryEngineEngine{})
}

func (r *DiscoveryEngineEngineRef) GetGVK() schema.GroupVersionKind {
	return DiscoveryEngineEngineGVK
}

func (r *DiscoveryEngineEngineRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DiscoveryEngineEngineRef) GetExternal() string {
	return r.External
}

func (r *DiscoveryEngineEngineRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *DiscoveryEngineEngineRef) ValidateExternal(ref string) error {
	_, err := ParseDiscoveryEngineEngineExternal(ref)
	return err
}

func (r *DiscoveryEngineEngineRef) ParseExternalToIdentity() (identity.Identity, error) {
	return ParseDiscoveryEngineEngineExternal(r.External)
}

// Normalize ensures the "External" reference (in string format) is set.
func (r *DiscoveryEngineEngineRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if err := refsv1beta1.Normalize(ctx, reader, r, defaultNamespace); err != nil {
		return err
	}
	id, err := ParseDiscoveryEngineEngineExternal(r.External)
	if err != nil {
		return err
	}
	r.External = id.String()
	return nil
}

// NormalizedExternal provision the "External" value for other resource that depends on DiscoveryEngineEngine.
// If the "External" is given in the other resource's spec.DiscoveryEngineEngineRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual DiscoveryEngineEngine object from the cluster.
func (r *DiscoveryEngineEngineRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", DiscoveryEngineEngineGVK.Kind)
	}
	// From given External
	if r.External != "" {
		id, err := ParseDiscoveryEngineEngineExternal(r.External)
		if err != nil {
			return "", err
		}
		r.External = id.String()
		return r.External, nil
	}

	// From the Config Connector object
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

// New builds a DiscoveryEngineEngineRef from the Config Connector DiscoveryEngineEngine object.
func NewDiscoveryEngineEngineRef(ctx context.Context, reader client.Reader, obj *DiscoveryEngineEngine) (*DiscoveryEngineEngineID, error) {
	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	collectionID := obj.Spec.Collection
	if collectionID == "" {
		return nil, fmt.Errorf("cannot resolve collection")
	}

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &DiscoveryEngineEngineID{
		CollectionLink: &CollectionLink{
			ProjectAndLocation: &ProjectAndLocation{
				ProjectID: projectID,
				Location:  location,
			},
			Collection: collectionID,
		},
		Engine: resourceID,
	}

	// Validate the status.externalRef, if set
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusID, err := ParseDiscoveryEngineEngineExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if statusID.String() != id.String() {
			return nil, fmt.Errorf("cannot change object key after creation; status=%q, new=%q",
				statusID.String(), id.String())
		}
	}
	return id, nil
}

// func (r *DiscoveryEngineEngineRef) Parent() (*DiscoveryEngineEngineParent, error) {
// 	if r.parent != nil {
// 		return r.parent, nil
// 	}
// 	if r.External != "" {
// 		parent, _, err := ParseDiscoveryEngineEngineExternal(r.External)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return parent, nil
// 	}
// 	return nil, fmt.Errorf("DiscoveryEngineEngineRef not initialized from `NewDiscoveryEngineEngineRef` or `NormalizedExternal`")
// }

// DiscoveryEngineEngineID is the resolved identifier for a DiscoveryEngineEngine
type DiscoveryEngineEngineID struct {
	*CollectionLink
	Engine string
}

func (p *DiscoveryEngineEngineID) String() string {
	return p.CollectionLink.String() + "/engines/" + p.Engine
}

func (p *DiscoveryEngineEngineID) FromExternal(ref string) error {
	id, err := ParseDiscoveryEngineEngineExternal(ref)
	if err != nil {
		return err
	}
	*p = *id
	return nil
}

// func (p *DiscoveryEngineEngineParent) String() string {
// 	return "projects/" + p.ProjectID + "/locations/" + p.Location
// }

// func asDiscoveryEngineEngineExternal(parent *DiscoveryEngineEngineParent, resourceID string) (external string) {
// 	return parent.String() + "/engines/" + resourceID
// }

func ParseDiscoveryEngineEngineExternal(external string) (*DiscoveryEngineEngineID, error) {
	s := strings.TrimPrefix(external, "//discoveryengine.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	tokens := strings.Split(s, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "engines" {
		projectAndLocation := &ProjectAndLocation{
			ProjectID: tokens[1],
			Location:  tokens[3],
		}
		collection := &CollectionLink{
			ProjectAndLocation: projectAndLocation,
			Collection:         tokens[5],
		}
		return &DiscoveryEngineEngineID{
			CollectionLink: collection,
			Engine:         tokens[7],
		}, nil
	}
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "engines" {
		projectAndLocation := &ProjectAndLocation{
			ProjectID: tokens[1],
			Location:  tokens[3],
		}
		collection := &CollectionLink{
			ProjectAndLocation: projectAndLocation,
			Collection:         "default_collection",
		}
		return &DiscoveryEngineEngineID{
			CollectionLink: collection,
			Engine:         tokens[5],
		}, nil
	}
	return nil, fmt.Errorf("format of DiscoveryEngineEngine external=%q was not known (use projects/{{projectId}}/locations/{{location}}/collections/{{collectionID}}/engines/{{engineID}})", external)
}
