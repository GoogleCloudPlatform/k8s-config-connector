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
	_ identity.IdentityV2 = &DataLabelingEvaluationJobIdentity{}
	_ identity.Resource   = &DataLabelingEvaluationJob{}
)

// DataLabelingEvaluationJobIdentityFormat is the template for constructing and parsing evaluation job GCP URLs.
var DataLabelingEvaluationJobIdentityFormat = gcpurls.Template[DataLabelingEvaluationJobIdentity]("datalabeling.googleapis.com", "projects/{project}/evaluationJobs/{evaluationjob}")

// DataLabelingEvaluationJobIdentity is the identity of a GCP DataLabelingEvaluationJob resource.
// +k8s:deepcopy-gen=false
type DataLabelingEvaluationJobIdentity struct {
	Project       string
	EvaluationJob string
}

func (i *DataLabelingEvaluationJobIdentity) String() string {
	return DataLabelingEvaluationJobIdentityFormat.ToString(*i)
}

func (i *DataLabelingEvaluationJobIdentity) FromExternal(ref string) error {
	parsed, match, err := DataLabelingEvaluationJobIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataLabelingEvaluationJob external=%q was not known (use %s): %w", ref, DataLabelingEvaluationJobIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataLabelingEvaluationJob external=%q was not known (use %s)", ref, DataLabelingEvaluationJobIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DataLabelingEvaluationJobIdentity) Host() string {
	return DataLabelingEvaluationJobIdentityFormat.Host()
}

func getIdentityFromDataLabelingEvaluationJobSpec(ctx context.Context, reader client.Reader, obj *DataLabelingEvaluationJob) (*DataLabelingEvaluationJobIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &DataLabelingEvaluationJobIdentity{
		Project:       projectID,
		EvaluationJob: resourceID,
	}
	return identity, nil
}

func (obj *DataLabelingEvaluationJob) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDataLabelingEvaluationJobSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DataLabelingEvaluationJobIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DataLabelingEvaluationJob identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
