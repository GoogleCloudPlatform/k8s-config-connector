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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DataprocAutoscalingPolicyIdentity{}
)

var DataprocAutoscalingPolicyIdentityFormat = gcpurls.Template[DataprocAutoscalingPolicyIdentity]("dataproc.googleapis.com", "projects/{project}/regions/{region}/autoscalingPolicies/{autoscalingPolicy}")

// DataprocAutoscalingPolicyIdentity defines the identity for a DataprocAutoscalingPolicy resource.
// +k8s:deepcopy-gen=false
type DataprocAutoscalingPolicyIdentity struct {
	Project           string
	Region            string
	AutoscalingPolicy string
}

// String returns the external reference for the DataprocAutoscalingPolicy resource.
func (i *DataprocAutoscalingPolicyIdentity) String() string {
	return DataprocAutoscalingPolicyIdentityFormat.ToString(*i)
}

// FromExternal parses the external reference string into the DataprocAutoscalingPolicyIdentity.
func (i *DataprocAutoscalingPolicyIdentity) FromExternal(ref string) error {
	parsed, match, err := DataprocAutoscalingPolicyIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DataprocAutoscalingPolicy external=%q was not known (use %s): %w", ref, DataprocAutoscalingPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DataprocAutoscalingPolicy external=%q was not known (use %s)", ref, DataprocAutoscalingPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

// Host returns the host for the DataprocAutoscalingPolicy resource.
func (i *DataprocAutoscalingPolicyIdentity) Host() string {
	return DataprocAutoscalingPolicyIdentityFormat.Host()
}

func getIdentityFromDataprocAutoscalingPolicySpec(ctx context.Context, reader client.Reader, obj client.Object) (*DataprocAutoscalingPolicyIdentity, error) {
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

	identity := &DataprocAutoscalingPolicyIdentity{
		Project:           projectID,
		Region:            location,
		AutoscalingPolicy: resourceID,
	}
	return identity, nil
}
