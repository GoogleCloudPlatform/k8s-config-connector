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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SpannerInstanceIdentity struct {
	id     string
	parent *SpannerInstanceParent
}

func (i *SpannerInstanceIdentity) String() string {
	return i.parent.String() + "/instances/" + i.id
}

func (r *SpannerInstanceIdentity) Parent() *SpannerInstanceParent {
	return r.parent
}

func (r *SpannerInstanceIdentity) ID() string {
	return r.id
}

type SpannerInstanceParent struct {
	ProjectID string
}

func (p *SpannerInstanceParent) String() string {
	return "projects/" + p.ProjectID
}

func NewSpannerInstanceIdentity(ctx context.Context, reader client.Reader, obj *SpannerInstance, u *unstructured.Unstructured) (*SpannerInstanceIdentity, error) {
	// Get Parent
	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
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
		actualIdentity, err := ParseSpannerInstanceExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualIdentity.parent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.parent.ProjectID, projectID)
		}
		if actualIdentity.id != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.id)
		}
	}
	return &SpannerInstanceIdentity{
		parent: &SpannerInstanceParent{ProjectID: projectID},
		id:     resourceID,
	}, nil
}

func (r *SpannerInstanceIdentity) SpannerInstanceConfigPrefix() string {
	return fmt.Sprintf("projects/%s/instanceConfigs/", r.parent.ProjectID)
}
