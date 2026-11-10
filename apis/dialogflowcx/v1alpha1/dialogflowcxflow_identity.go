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
	_ identity.Resource = &DialogflowCXFlow{}
)

type dialogflowCXAgentIdentity struct {
	Project  string
	Location string
	Agent    string
}

var dialogflowCXAgentIdentityFormat = gcpurls.Template[dialogflowCXAgentIdentity]("dialogflow.googleapis.com", "projects/{project}/locations/{location}/agents/{agent}")

func getIdentityFromDialogflowCXFlowSpec(ctx context.Context, reader client.Reader, obj *DialogflowCXFlow) (*DialogflowCXFlowIdentity, error) {
	parent := common.ValueOf(obj.Spec.Parent)
	if parent == "" {
		return nil, fmt.Errorf("spec.parent is required")
	}

	parsedParent, match, err := dialogflowCXAgentIdentityFormat.Parse(parent)
	if err != nil || !match {
		return nil, fmt.Errorf("spec.parent %q is invalid, expected format: projects/{project}/locations/{location}/agents/{agent}", parent)
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)

	return &DialogflowCXFlowIdentity{
		Project:  parsedParent.Project,
		Location: parsedParent.Location,
		Agent:    parsedParent.Agent,
		Flow:     resourceID,
	}, nil
}

func (obj *DialogflowCXFlow) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDialogflowCXFlowSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.Name)
	}
	if externalRef != "" {
		statusIdentity := &DialogflowCXFlowIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity.Flow == "" {
			specIdentity.Flow = statusIdentity.Flow
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DialogflowCXFlow identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
