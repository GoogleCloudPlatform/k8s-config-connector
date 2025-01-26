// Copyright 2025 Google LLC
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

// +tool:krm-identity
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.message: google.cloud.filestore.v1.Instance
// crd.type: FilestoreInstance
// crd.version: v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// InstanceIdentity defines the full identity for a filestore Instance
//
// +k8s:deepcopy-gen=false
type InstanceIdentity struct {
	parent.ProjectAndLocationParent
	Instance string
}

func (l *InstanceIdentity) String() string {
	return l.FullyQualifiedName()
}

func (l *InstanceIdentity) FullyQualifiedName() string {
	return l.ProjectAndLocationParent.String() + "/instances/" + l.Instance
}

// InstanceIdentityForObject builds an InstanceIdentity from the Config Connector object.
func InstanceIdentityForObject(ctx context.Context, reader client.Reader, obj *FilestoreInstance) (*InstanceIdentity, error) {
	// Get Parent
	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
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

	// Get desired ID
	resourceID := valueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &InstanceIdentity{
		ProjectAndLocationParent: parent.ProjectAndLocationParent{
			ProjectID: projectID,
			Location:  location,
		},
		Instance: resourceID,
	}

	// Validate the status.externalRef, if set
	externalRef := valueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusID, err := ParseInstanceIdentityExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if statusID.String() != id.String() {
			return nil, fmt.Errorf("cannot change object identity after creation; status=%q, new=%q",
				statusID.String(), id.String())
		}
		id = statusID
	}
	return id, nil
}

// Should match https://cloud.google.com/asset-inventory/docs/asset-names format
func ParseInstanceIdentityExternal(external string) (*InstanceIdentity, error) {
	s := strings.TrimPrefix(external, "//file.googleapis.com/")
	s = strings.TrimPrefix(s, "/")
	tokens := strings.Split(s, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		projectAndLocation := parent.ProjectAndLocationParent{
			ProjectID: tokens[1],
			Location:  tokens[3],
		}

		link := &InstanceIdentity{
			ProjectAndLocationParent: projectAndLocation,
			Instance:                 tokens[5],
		}
		return link, nil
	}
	return nil, fmt.Errorf("format of FilestoreInstance external=%q was not known (use projects/{{projectId}}/locations/{{location}}/instances/{{instanceID}})", external)
}

func valueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}
