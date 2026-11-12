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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.Resource = &DialogflowCXPage{}
)

var flowParentFormat = gcpurls.Template[flowParent]("dialogflow.googleapis.com", "projects/{project}/locations/{location}/agents/{agent}/flows/{flow}")

type flowParent struct {
	Project  string
	Location string
	Agent    string
	Flow     string
}

func getIdentityFromDialogflowCXPageSpec(ctx context.Context, reader client.Reader, obj *DialogflowCXPage) (*DialogflowCXPageIdentity, error) {
	if obj.Spec.Parent == nil || *obj.Spec.Parent == "" {
		return nil, fmt.Errorf("spec.parent is required")
	}

	parent, match, err := flowParentFormat.Parse(*obj.Spec.Parent)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.parent %q: %w", *obj.Spec.Parent, err)
	}
	if !match {
		return nil, fmt.Errorf("format of DialogflowCXPage spec.parent=%q was not known (use %s)", *obj.Spec.Parent, flowParentFormat.CanonicalForm())
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)

	return &DialogflowCXPageIdentity{
		Project:  parent.Project,
		Location: parent.Location,
		Agent:    parent.Agent,
		Flow:     parent.Flow,
		Page:     resourceID,
	}, nil
}

func (obj *DialogflowCXPage) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDialogflowCXPageSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.Name)
	}
	if externalRef != "" {
		statusIdentity := &DialogflowCXPageIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity.Page == "" {
			specIdentity.Page = statusIdentity.Page
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DialogflowCXPage identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
