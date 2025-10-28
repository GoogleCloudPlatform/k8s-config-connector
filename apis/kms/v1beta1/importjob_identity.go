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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ImportJobIdentity defines the resource reference to KMSImportJob, which "External" field
// holds the GCP identifier for the KRM object.
type ImportJobIdentity struct {
	parent *KMSKeyRingIdentity
	id     string
}

func (i *ImportJobIdentity) String() string {
	return i.parent.String() + "/importJobs/" + i.id
}

func (i *ImportJobIdentity) ID() string {
	return i.id
}

func (i *ImportJobIdentity) Parent() *KMSKeyRingIdentity {
	return i.parent
}

// New builds a ImportJobIdentity from the Config Connector ImportJob object.
func NewImportJobIdentity(ctx context.Context, reader client.Reader, obj *KMSImportJob) (*ImportJobIdentity, error) {

	// Get Parent
	kmsKeyRingExternal, err := obj.Spec.KMSKeyRingRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, err
	}
	kmsKeyRing, err := ParseKMSKeyRingExternal(kmsKeyRingExternal)
	if err != nil {
		return nil, err
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseImportJobExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.String() != kmsKeyRing.String() {
			return nil, fmt.Errorf("spec.kmsKeyRingRef changed, expect %s, got %s", actualParent.String(), kmsKeyRing.String())
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ImportJobIdentity{
		parent: kmsKeyRing,
		id:     resourceID,
	}, nil
}

func ParseImportJobExternal(external string) (*KMSKeyRingIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "keyRings" || tokens[6] != "importJobs" {
		return nil, "", fmt.Errorf("format of KMSImportJob external=%q was not known (use projects/{{projectID}}/locations/{{location}}/keyRings/{{keyRingID}}/importJobs/{{importJobID}})", external)
	}
	p := &KMSKeyRingIdentity{
		Parent: &parent.ProjectAndLocationParent{
			ProjectID: tokens[1], Location: tokens[3],
		}, ID: tokens[5],
	}
	resourceID := tokens[7]
	return p, resourceID, nil
}
