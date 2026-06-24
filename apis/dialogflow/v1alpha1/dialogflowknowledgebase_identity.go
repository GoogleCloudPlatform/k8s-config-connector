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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DialogflowKnowledgeBaseIdentity{}
	_ identity.Resource   = &DialogflowKnowledgeBase{}
)

var (
	DialogflowKnowledgeBaseGlobalIdentityFormat   = gcpurls.Template[DialogflowKnowledgeBaseIdentityGlobal]("dialogflow.googleapis.com", "projects/{project}/knowledgeBases/{knowledge_base}")
	DialogflowKnowledgeBaseRegionalIdentityFormat = gcpurls.Template[DialogflowKnowledgeBaseIdentityRegional]("dialogflow.googleapis.com", "projects/{project}/locations/{location}/knowledgeBases/{knowledge_base}")
)

// +k8s:deepcopy-gen=false
type DialogflowKnowledgeBaseIdentityGlobal struct {
	Project        string
	Knowledge_base string
}

// +k8s:deepcopy-gen=false
type DialogflowKnowledgeBaseIdentityRegional struct {
	Project        string
	Location       string
	Knowledge_base string
}

// +k8s:deepcopy-gen=false
type DialogflowKnowledgeBaseIdentity struct {
	Project        string
	Location       string // empty if global
	Knowledge_base string
}

func (i *DialogflowKnowledgeBaseIdentity) String() string {
	if i.Location != "" {
		return DialogflowKnowledgeBaseRegionalIdentityFormat.ToString(DialogflowKnowledgeBaseIdentityRegional{
			Project:        i.Project,
			Location:       i.Location,
			Knowledge_base: i.Knowledge_base,
		})
	}
	return DialogflowKnowledgeBaseGlobalIdentityFormat.ToString(DialogflowKnowledgeBaseIdentityGlobal{
		Project:        i.Project,
		Knowledge_base: i.Knowledge_base,
	})
}

func (i *DialogflowKnowledgeBaseIdentity) FromExternal(ref string) error {
	// Try parsing as regional first
	if parsedRegional, match, _ := DialogflowKnowledgeBaseRegionalIdentityFormat.Parse(ref); match {
		i.Project = parsedRegional.Project
		i.Location = parsedRegional.Location
		i.Knowledge_base = parsedRegional.Knowledge_base
		return nil
	}

	// Try parsing as global
	if parsedGlobal, match, _ := DialogflowKnowledgeBaseGlobalIdentityFormat.Parse(ref); match {
		i.Project = parsedGlobal.Project
		i.Location = ""
		i.Knowledge_base = parsedGlobal.Knowledge_base
		return nil
	}

	return fmt.Errorf("format of DialogflowKnowledgeBase external=%q was not known (use %s or %s)",
		ref, DialogflowKnowledgeBaseRegionalIdentityFormat.CanonicalForm(), DialogflowKnowledgeBaseGlobalIdentityFormat.CanonicalForm())
}

func (i *DialogflowKnowledgeBaseIdentity) Host() string {
	return DialogflowKnowledgeBaseRegionalIdentityFormat.Host()
}

func getIdentityFromDialogflowKnowledgeBaseSpec(ctx context.Context, reader client.Reader, obj *DialogflowKnowledgeBase) (*DialogflowKnowledgeBaseIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location := ""
	if obj.Spec.Location != nil {
		location = *obj.Spec.Location
	}

	identity := &DialogflowKnowledgeBaseIdentity{
		Project:        projectID,
		Location:       location,
		Knowledge_base: resourceID,
	}
	return identity, nil
}

func (obj *DialogflowKnowledgeBase) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDialogflowKnowledgeBaseSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DialogflowKnowledgeBaseIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DialogflowKnowledgeBase identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *DialogflowKnowledgeBase) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
