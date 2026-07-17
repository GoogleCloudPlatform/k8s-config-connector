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
	_ identity.IdentityV2 = &MonitoringDashboardIdentity{}
	_ identity.Resource   = &MonitoringDashboard{}
)

var MonitoringDashboardIdentityFormat = gcpurls.Template[MonitoringDashboardIdentity]("monitoring.googleapis.com", "projects/{project}/dashboards/{dashboard}")

// +k8s:deepcopy-gen=false
type MonitoringDashboardIdentity struct {
	Project   string
	Dashboard string
}

func (i *MonitoringDashboardIdentity) String() string {
	return MonitoringDashboardIdentityFormat.ToString(*i)
}

func (i *MonitoringDashboardIdentity) FromExternal(ref string) error {
	parsed, match, err := MonitoringDashboardIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MonitoringDashboard external=%q was not known (use %s): %w", ref, MonitoringDashboardIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MonitoringDashboard external=%q was not known (use %s)", ref, MonitoringDashboardIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MonitoringDashboardIdentity) Host() string {
	return MonitoringDashboardIdentityFormat.Host()
}

func (i *MonitoringDashboardIdentity) ParentString() string {
	return "projects/" + i.Project
}

func getIdentityFromMonitoringDashboardSpec(ctx context.Context, reader client.Reader, obj *MonitoringDashboard) (*MonitoringDashboardIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &MonitoringDashboardIdentity{
		Project:   projectID,
		Dashboard: resourceID,
	}
	return identity, nil
}

func (obj *MonitoringDashboard) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromMonitoringDashboardSpec(ctx, reader, obj)
}
