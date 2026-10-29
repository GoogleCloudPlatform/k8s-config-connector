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
	_ identity.IdentityV2 = &DialogflowConversationProfileIdentity{}
	_ identity.Resource   = &DialogflowConversationProfile{}
)

var (
	DialogflowConversationProfileGlobalIdentityFormat   = gcpurls.Template[DialogflowConversationProfileIdentityGlobal]("dialogflow.googleapis.com", "projects/{project}/conversationProfiles/{conversationProfile}")
	DialogflowConversationProfileRegionalIdentityFormat = gcpurls.Template[DialogflowConversationProfileIdentityRegional]("dialogflow.googleapis.com", "projects/{project}/locations/{location}/conversationProfiles/{conversationProfile}")
)

// +k8s:deepcopy-gen=false
type DialogflowConversationProfileIdentityGlobal struct {
	Project             string
	ConversationProfile string
}

// +k8s:deepcopy-gen=false
type DialogflowConversationProfileIdentityRegional struct {
	Project             string
	Location            string
	ConversationProfile string
}

// DialogflowConversationProfileIdentity is the identity of a GCP DialogflowConversationProfile resource.
// +k8s:deepcopy-gen=false
type DialogflowConversationProfileIdentity struct {
	Project             string
	Location            string // empty if global
	ConversationProfile string
}

func (i *DialogflowConversationProfileIdentity) String() string {
	if i.Location != "" {
		return DialogflowConversationProfileRegionalIdentityFormat.ToString(DialogflowConversationProfileIdentityRegional{
			Project:             i.Project,
			Location:            i.Location,
			ConversationProfile: i.ConversationProfile,
		})
	}
	return DialogflowConversationProfileGlobalIdentityFormat.ToString(DialogflowConversationProfileIdentityGlobal{
		Project:             i.Project,
		ConversationProfile: i.ConversationProfile,
	})
}

func (i *DialogflowConversationProfileIdentity) FromExternal(ref string) error {
	// Try parsing as regional first
	if parsedRegional, match, _ := DialogflowConversationProfileRegionalIdentityFormat.Parse(ref); match {
		i.Project = parsedRegional.Project
		i.Location = parsedRegional.Location
		i.ConversationProfile = parsedRegional.ConversationProfile
		return nil
	}

	// Try parsing as global
	if parsedGlobal, match, _ := DialogflowConversationProfileGlobalIdentityFormat.Parse(ref); match {
		i.Project = parsedGlobal.Project
		i.Location = ""
		i.ConversationProfile = parsedGlobal.ConversationProfile
		return nil
	}

	return fmt.Errorf("format of DialogflowConversationProfile external=%q was not known (use %s or %s)",
		ref, DialogflowConversationProfileRegionalIdentityFormat.CanonicalForm(), DialogflowConversationProfileGlobalIdentityFormat.CanonicalForm())
}

func (i *DialogflowConversationProfileIdentity) Host() string {
	return DialogflowConversationProfileRegionalIdentityFormat.Host()
}

func (i *DialogflowConversationProfileIdentity) ParentString() string {
	if i.Location != "" {
		return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
	}
	return fmt.Sprintf("projects/%s", i.Project)
}

func getIdentityFromDialogflowConversationProfileSpec(ctx context.Context, reader client.Reader, obj *DialogflowConversationProfile) (*DialogflowConversationProfileIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	location := ""
	if obj.Spec.Location != nil {
		location = *obj.Spec.Location
	}

	identity := &DialogflowConversationProfileIdentity{
		Project:             projectID,
		Location:            location,
		ConversationProfile: resourceID,
	}
	return identity, nil
}

func (obj *DialogflowConversationProfile) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDialogflowConversationProfileSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DialogflowConversationProfileIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DialogflowConversationProfile identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *DialogflowConversationProfile) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
