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
	_ identity.IdentityV2 = &MonitoringUptimeCheckConfigIdentity{}
	_ identity.Resource   = &MonitoringUptimeCheckConfig{}
)

var MonitoringUptimeCheckConfigIdentityFormat = gcpurls.Template[MonitoringUptimeCheckConfigIdentity]("monitoring.googleapis.com", "projects/{project}/uptimeCheckConfigs/{uptimecheckconfig}")

// MonitoringUptimeCheckConfigIdentity is the identity of a Google Cloud MonitoringUptimeCheckConfig resource.
// +k8s:deepcopy-gen=false
type MonitoringUptimeCheckConfigIdentity struct {
	Project           string
	UptimeCheckConfig string
}

func (i *MonitoringUptimeCheckConfigIdentity) String() string {
	return MonitoringUptimeCheckConfigIdentityFormat.ToString(*i)
}

func (i *MonitoringUptimeCheckConfigIdentity) FromExternal(ref string) error {
	parsed, match, err := MonitoringUptimeCheckConfigIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MonitoringUptimeCheckConfig external=%q was not known (use %s): %w", ref, MonitoringUptimeCheckConfigIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MonitoringUptimeCheckConfig external=%q was not known (use %s)", ref, MonitoringUptimeCheckConfigIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MonitoringUptimeCheckConfigIdentity) Host() string {
	return MonitoringUptimeCheckConfigIdentityFormat.Host()
}

func getIdentityFromMonitoringUptimeCheckConfigSpec(ctx context.Context, reader client.Reader, obj *MonitoringUptimeCheckConfig) (*MonitoringUptimeCheckConfigIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &MonitoringUptimeCheckConfigIdentity{
		Project:           projectID,
		UptimeCheckConfig: resourceID,
	}
	return identity, nil
}

func (obj *MonitoringUptimeCheckConfig) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromMonitoringUptimeCheckConfigSpec(ctx, reader, obj)
}
