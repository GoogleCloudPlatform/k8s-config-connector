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
	_ identity.IdentityV2 = &MonitoringGroupIdentity{}
	_ identity.Resource   = &MonitoringGroup{}
)

var MonitoringGroupIdentityFormat = gcpurls.Template[MonitoringGroupIdentity]("monitoring.googleapis.com", "projects/{project}/groups/{group}")

// MonitoringGroupIdentity is the identity of a GCP MonitoringGroup resource.
// +k8s:deepcopy-gen=false
type MonitoringGroupIdentity struct {
	Project string
	Group   string
}

func (i *MonitoringGroupIdentity) String() string {
	return MonitoringGroupIdentityFormat.ToString(*i)
}

func (i *MonitoringGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := MonitoringGroupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MonitoringGroup external=%q was not known (use %s): %w", ref, MonitoringGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MonitoringGroup external=%q was not known (use %s)", ref, MonitoringGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MonitoringGroupIdentity) Host() string {
	return MonitoringGroupIdentityFormat.Host()
}

func getIdentityFromMonitoringGroupSpec(ctx context.Context, reader client.Reader, obj *MonitoringGroup) (*MonitoringGroupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &MonitoringGroupIdentity{
		Project: projectID,
		Group:   resourceID,
	}
	return identity, nil
}

func (obj *MonitoringGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromMonitoringGroupSpec(ctx, reader, obj)
}
