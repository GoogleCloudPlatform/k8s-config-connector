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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const serviceDomain = "//alloydb.googleapis.com"

// InstanceIdentity defines the resource reference to AlloyDBInstance, which "External" field
// holds the GCP identifier for the KRM object.
type InstanceIdentity struct {
	parent *InstanceParent
	id     string
}

func (i *InstanceIdentity) String() string {
	return i.parent.String() + "/instances/" + i.id
}

func (i *InstanceIdentity) ID() string {
	return i.id
}

func (i *InstanceIdentity) Parent() *InstanceParent {
	return i.parent
}

// AsExternalRef builds a externalRef from a PrivilegedAccessManagerEntitlement.
func (i *InstanceIdentity) AsExternalRef() *string {
	er := serviceDomain + "/" + i.String()
	return &er
}

type InstanceParent struct {
	clusterName string
}

func (p *InstanceParent) String() string {
	return p.clusterName
}

// New builds a InstanceIdentity from the Config Connector Instance object.
func NewInstanceIdentity(ctx context.Context, reader client.Reader, obj *AlloyDBInstance) (*InstanceIdentity, error) {

	// Get Parent
	clusterRef, err := refsv1beta1.ResolveAlloyDBCluster(ctx, reader, obj, obj.Spec.ClusterRef)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve AlloyDBCluster ref: %w", err)
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved ExternalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseInstanceExternalRef(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.clusterName != clusterRef.String() {
			return nil, fmt.Errorf("spec.clusterRef changed, expect %s, got %s", actualParent.clusterName, clusterRef)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &InstanceIdentity{
		parent: &InstanceParent{
			clusterName: clusterRef.String(),
		},
		id: resourceID,
	}, nil
}

func ParseInstanceExternalRef(externalRef string) (parent *InstanceParent, resourceID string, err error) {
	if !strings.HasPrefix(externalRef, serviceDomain) {
		return nil, "", fmt.Errorf("externalRef should have prefix %s, got %s", serviceDomain, externalRef)
	}
	path := strings.TrimPrefix(externalRef, serviceDomain+"/")
	return ParseInstanceExternal(path)
}

func ParseInstanceExternal(external string) (parent *InstanceParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "clusters" || tokens[6] != "instances" {
		return nil, "", fmt.Errorf("format of AlloyDBInstance external=%q was not known (use projects/<projectId>/locations/<location>/clusters/<clusterID>/instances/<instanceID>)", external)
	}
	parent = &InstanceParent{
		clusterName: fmt.Sprintf("%s/%s/%s/%s/%s/%s", tokens[0], tokens[1], tokens[2], tokens[3], tokens[4], tokens[5]),
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
