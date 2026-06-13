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
	_ identity.IdentityV2 = &MonitoringServiceLevelObjectiveIdentity{}
	_ identity.Resource   = &MonitoringServiceLevelObjective{}
)

var MonitoringServiceLevelObjectiveIdentityFormat = gcpurls.Template[MonitoringServiceLevelObjectiveIdentity]("monitoring.googleapis.com", "projects/{project}/services/{service}/serviceLevelObjectives/{servicelevelobjective}")

// MonitoringServiceLevelObjectiveIdentity is the identity of a GCP MonitoringServiceLevelObjective resource.
// +k8s:deepcopy-gen=false
type MonitoringServiceLevelObjectiveIdentity struct {
	Project               string
	Service               string
	ServiceLevelObjective string
}

func (i *MonitoringServiceLevelObjectiveIdentity) String() string {
	return MonitoringServiceLevelObjectiveIdentityFormat.ToString(*i)
}

func (i *MonitoringServiceLevelObjectiveIdentity) FromExternal(ref string) error {
	parsed, match, err := MonitoringServiceLevelObjectiveIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MonitoringServiceLevelObjective external=%q was not known (use %s): %w", ref, MonitoringServiceLevelObjectiveIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MonitoringServiceLevelObjective external=%q was not known (use %s)", ref, MonitoringServiceLevelObjectiveIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MonitoringServiceLevelObjectiveIdentity) Host() string {
	return MonitoringServiceLevelObjectiveIdentityFormat.Host()
}

func getIdentityFromMonitoringServiceLevelObjectiveSpec(ctx context.Context, reader client.Reader, obj *MonitoringServiceLevelObjective) (*MonitoringServiceLevelObjectiveIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	if obj.Spec.ServiceRef == nil {
		return nil, fmt.Errorf("spec.serviceRef is required")
	}

	serviceRef := obj.Spec.ServiceRef.DeepCopy()
	if err := serviceRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("cannot resolve serviceRef: %w", err)
	}
	serviceIDRaw, err := serviceRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("cannot parse serviceRef: %w", err)
	}
	serviceID := serviceIDRaw.(*MonitoringServiceIdentity)

	if serviceID.Project != projectID {
		return nil, fmt.Errorf("serviceRef.project must match spec.projectRef")
	}

	identity := &MonitoringServiceLevelObjectiveIdentity{
		Project:               projectID,
		Service:               serviceID.Service,
		ServiceLevelObjective: resourceID,
	}
	return identity, nil
}

func (obj *MonitoringServiceLevelObjective) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromMonitoringServiceLevelObjectiveSpec(ctx, reader, obj)
}
