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

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ProcessorVersionIdentity defines the resource reference to DocumentAI, which "External" field
// holds the GCP identifier for the KRM object.
type ProcessorVersionIdentity struct {
	parent *ProcessorVersionParent
	id     string
}

func (i *ProcessorVersionIdentity) String() string {
	return i.parent.String() + "/processorVersions/" + i.id
}

func (i *ProcessorVersionIdentity) ID() string {
	return i.id
}

func (i *ProcessorVersionIdentity) Parent() *ProcessorVersionParent {
	return i.parent
}

type ProcessorVersionParent struct {
	Processor string
}

func (p *ProcessorVersionParent) String() string {
	return p.Processor
}

// NewProcessorVersionIdentity builds a ProcessorVersionIdentity from the Config Connector ProcessorVersion object.
func NewProcessorVersionIdentity(ctx context.Context, reader client.Reader, obj *DocumentAIProcessorVersion) (*ProcessorVersionIdentity, error) {
	//Get parent
	processorRef := &ProcessorRef{}
	processor, err := processorRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
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
		actualParent, actualResourceID, err := ParseProcessorVersionExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.Processor != processor {
			return nil, fmt.Errorf("spec.processorRef changed, expect %s, got %s", actualParent.Processor, processor)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ProcessorVersionIdentity{
		parent: &ProcessorVersionParent{
			Processor: processor,
		},
		id: resourceID,
	}, nil
}

func ParseProcessorVersionExternal(external string) (parent *ProcessorVersionParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "processors" || tokens[6] != "processorVersions" {
		return nil, "", fmt.Errorf("format of DocumentAI external=%q was not known (use projects/{{projectID}}/locations/{{location}}/processors/{{processorID}}/processorVersions/{{processorversionID}})", external)
	}
	processor := strings.Join(tokens[:len(tokens)-2], "/")
	parent = &ProcessorVersionParent{
		processor,
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
