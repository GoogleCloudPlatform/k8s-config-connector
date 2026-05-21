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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DataflowFlexTemplateJobIdentity{}
	_ identity.Resource   = &DataflowFlexTemplateJob{}
)

var DataflowFlexTemplateJobIdentityFormat = gcpurls.Template[DataflowFlexTemplateJobIdentity]("dataflow.googleapis.com", "projects/{project}/locations/{location}/jobs/{job}")

// +k8s:deepcopy-gen=false
type DataflowFlexTemplateJobIdentity struct {
	Project  string
	Location string
	Job      string
}

func (i *DataflowFlexTemplateJobIdentity) String() string {
	return DataflowFlexTemplateJobIdentityFormat.ToString(*i)
}

func (i *DataflowFlexTemplateJobIdentity) FromExternal(ref string) error {
	parsed, match, err := DataflowFlexTemplateJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataflowFlexTemplateJob external=%q was not known (use %s): %w", ref, DataflowFlexTemplateJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataflowFlexTemplateJob external=%q was not known (use %s)", ref, DataflowFlexTemplateJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DataflowFlexTemplateJobIdentity) Host() string {
	return DataflowFlexTemplateJobIdentityFormat.Host()
}

func getIdentityFromDataflowFlexTemplateJobSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DataflowFlexTemplateJobIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	var location string
	var jobID string

	if typed, ok := obj.(*DataflowFlexTemplateJob); ok {
		location = common.ValueOf(typed.Spec.Region)
		jobID = typed.Status.JobID
	} else {
		u := obj.(*unstructured.Unstructured)
		location, _, _ = unstructured.NestedString(u.Object, "spec", "region")
		jobID, _, _ = unstructured.NestedString(u.Object, "status", "jobId")
	}

	if location == "" {
		return nil, fmt.Errorf("cannot resolve region")
	}

	// Job ID is server-generated. Use JobID from status if it exists.
	if jobID == "" {
		return nil, nil
	}

	return &DataflowFlexTemplateJobIdentity{
		Project:  projectID,
		Location: location,
		Job:      jobID,
	}, nil
}

func (obj *DataflowFlexTemplateJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDataflowFlexTemplateJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Use Status.ExternalRef if it exists
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualIdentity := &DataflowFlexTemplateJobIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if specIdentity != nil {
			if actualIdentity.Project != specIdentity.Project {
				return nil, fmt.Errorf("project changed, expect %s, got %s", actualIdentity.Project, specIdentity.Project)
			}
			if actualIdentity.Location != specIdentity.Location {
				return nil, fmt.Errorf("location changed, expect %s, got %s", actualIdentity.Location, specIdentity.Location)
			}
		}

		return actualIdentity, nil
	}

	if specIdentity == nil {
		return nil, nil
	}

	return specIdentity, nil
}
