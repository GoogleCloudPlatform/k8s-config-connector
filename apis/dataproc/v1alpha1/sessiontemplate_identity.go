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
	_ identity.IdentityV2 = &DataprocSessionTemplateIdentity{}
	_ identity.Resource   = &DataprocSessionTemplate{}
)

var DataprocSessionTemplateIdentityFormat = gcpurls.Template[DataprocSessionTemplateIdentity]("dataproc.googleapis.com", "projects/{project}/locations/{location}/sessionTemplates/{template}")

// +k8s:deepcopy-gen=false
// DataprocSessionTemplateIdentity is the identity of a GCP DataprocSessionTemplate resource.
type DataprocSessionTemplateIdentity struct {
	Project  string
	Location string
	Template string
}

func (i *DataprocSessionTemplateIdentity) String() string {
	return DataprocSessionTemplateIdentityFormat.ToString(*i)
}

func (i *DataprocSessionTemplateIdentity) FromExternal(ref string) error {
	parsed, match, err := DataprocSessionTemplateIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataprocSessionTemplate external=%q was not known (use %s): %w", ref, DataprocSessionTemplateIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataprocSessionTemplate external=%q was not known (use %s)", ref, DataprocSessionTemplateIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DataprocSessionTemplateIdentity) Host() string {
	return DataprocSessionTemplateIdentityFormat.Host()
}

func getIdentityFromDataprocSessionTemplateSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DataprocSessionTemplateIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &DataprocSessionTemplateIdentity{
		Project:  projectID,
		Location: location,
		Template: resourceID,
	}
	return identity, nil
}

func (obj *DataprocSessionTemplate) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDataprocSessionTemplateSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &DataprocSessionTemplateIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DataprocSessionTemplate identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
