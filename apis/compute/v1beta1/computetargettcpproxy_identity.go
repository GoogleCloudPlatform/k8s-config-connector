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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type TargetTCPProxyIdentity struct {
	id     string
	parent *parent.ComputeParent
}

func (i *TargetTCPProxyIdentity) String() string {
	return i.parent.String() + "/targetTcpProxies/" + i.id
}

func (i *TargetTCPProxyIdentity) Parent() *parent.ComputeParent {
	return i.parent
}

func (i *TargetTCPProxyIdentity) ID() string {
	return i.id
}

func NewTargetTCPProxyIdentity(ctx context.Context, reader client.Reader, obj *ComputeTargetTCPProxy, u *unstructured.Unstructured) (*TargetTCPProxyIdentity, error) {
	// Get projectID
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, err
	}
	// Get Location
	var location string
	if obj.Spec.Location == nil {
		location = "global"
	} else {
		location = common.ValueOf(obj.Spec.Location)
	}

	// Get resourceID
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
		actualIdentity, err := parseTargetTCPProxyExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualIdentity.parent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.parent.ProjectID, projectID)
		}
		if actualIdentity.parent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.parent.Location, location)
		}
		if actualIdentity.id != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.id)
		}
	}

	return &TargetTCPProxyIdentity{
		parent: &parent.ComputeParent{ProjectID: projectID, Location: location},
		id:     resourceID,
	}, nil
}
