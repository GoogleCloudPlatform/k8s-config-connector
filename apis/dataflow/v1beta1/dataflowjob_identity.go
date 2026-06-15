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

package v1beta1

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
	_ identity.IdentityV2 = &DataflowJobIdentity{}
	_ identity.Resource   = &DataflowJob{}
)

var DataflowJobIdentityFormat = gcpurls.Template[DataflowJobIdentity]("dataflow.googleapis.com", "projects/{project}/locations/{location}/jobs/{job}")

// +k8s:deepcopy-gen=false
type DataflowJobIdentity struct {
	Project  string
	Location string
	Job      string
}

func (i *DataflowJobIdentity) String() string {
	return DataflowJobIdentityFormat.ToString(*i)
}

func (i *DataflowJobIdentity) FromExternal(ref string) error {
	parsed, match, err := DataflowJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataflowJob external=%q was not known (use %s): %w", ref, DataflowJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataflowJob external=%q was not known (use %s)", ref, DataflowJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DataflowJobIdentity) Host() string {
	return DataflowJobIdentityFormat.Host()
}

func getIdentityFromDataflowJobSpec(ctx context.Context, reader client.Reader, obj *DataflowJob) (*DataflowJobIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	region := common.ValueOf(obj.Spec.Region)
	if region == "" {
		return nil, fmt.Errorf("region is required but not found in spec")
	}

	identity := &DataflowJobIdentity{
		Project:  projectID,
		Location: region,
		Job:      resourceID,
	}
	return identity, nil
}

func (obj *DataflowJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDataflowJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
