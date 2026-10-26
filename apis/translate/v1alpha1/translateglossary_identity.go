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
	_ identity.IdentityV2 = &TranslateGlossaryIdentity{}
	_ identity.Resource   = &TranslateGlossary{}
)

var TranslateGlossaryIdentityFormat = gcpurls.Template[TranslateGlossaryIdentity]("translate.googleapis.com", "projects/{project}/locations/{location}/glossaries/{glossary}")

// TranslateGlossaryIdentity is the identity of a GCP TranslateGlossary resource.
// +k8s:deepcopy-gen=false
type TranslateGlossaryIdentity struct {
	Project  string
	Location string
	Glossary string
}

func (i *TranslateGlossaryIdentity) String() string {
	return TranslateGlossaryIdentityFormat.ToString(*i)
}

func (i *TranslateGlossaryIdentity) FromExternal(ref string) error {
	parsed, match, err := TranslateGlossaryIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of TranslateGlossary external=%q was not known (use %s): %w", ref, TranslateGlossaryIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of TranslateGlossary external=%q was not known (use %s)", ref, TranslateGlossaryIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *TranslateGlossaryIdentity) Host() string {
	return TranslateGlossaryIdentityFormat.Host()
}

func (i *TranslateGlossaryIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromTranslateGlossarySpec(ctx context.Context, reader client.Reader, obj *TranslateGlossary) (*TranslateGlossaryIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	return &TranslateGlossaryIdentity{
		Project:  projectID,
		Location: location,
		Glossary: resourceID,
	}, nil
}

func (obj *TranslateGlossary) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromTranslateGlossarySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &TranslateGlossaryIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change TranslateGlossary identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
