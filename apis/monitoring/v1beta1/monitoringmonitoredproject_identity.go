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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &MonitoringMonitoredProjectIdentity{}
	_ identity.Resource   = &MonitoringMonitoredProject{}
)

var MonitoringMonitoredProjectIdentityFormat = gcpurls.Template[MonitoringMonitoredProjectIdentity]("monitoring.googleapis.com", "locations/global/metricsScopes/{metricsScope}/projects/{project}")

// MonitoringMonitoredProjectIdentity is the identity of a Google Cloud MonitoringMonitoredProject resource.
// +k8s:deepcopy-gen=false
type MonitoringMonitoredProjectIdentity struct {
	MetricsScope string
	Project      string
}

func (i *MonitoringMonitoredProjectIdentity) String() string {
	return MonitoringMonitoredProjectIdentityFormat.ToString(*i)
}

func (i *MonitoringMonitoredProjectIdentity) FromExternal(ref string) error {
	parsed, match, err := MonitoringMonitoredProjectIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MonitoringMonitoredProject external=%q was not known (use %s): %w", ref, MonitoringMonitoredProjectIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MonitoringMonitoredProject external=%q was not known (use %s)", ref, MonitoringMonitoredProjectIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MonitoringMonitoredProjectIdentity) Host() string {
	return MonitoringMonitoredProjectIdentityFormat.Host()
}

func (i *MonitoringMonitoredProjectIdentity) ParentString() string {
	return fmt.Sprintf("locations/global/metricsScopes/%s", i.MetricsScope)
}

func getIdentityFromMonitoringMonitoredProjectSpec(ctx context.Context, reader client.Reader, obj *MonitoringMonitoredProject) (*MonitoringMonitoredProjectIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	metricsScope := obj.Spec.MetricsScope
	if metricsScope == "" {
		return nil, fmt.Errorf("metricsScope is required")
	}

	// Normalize metricsScope to get the metricsScope ID.
	// It can be of the form "locations/global/metricsScopes/{metricsScope}" or "location/global/metricsScopes/{metricsScope}"
	// We want to extract just the {metricsScope} part.
	var metricsScopeID string
	tokens := strings.Split(metricsScope, "/")
	if len(tokens) == 4 {
		if tokens[0] != "locations" && tokens[0] != "location" {
			return nil, fmt.Errorf("invalid metricsScope format %q: first segment must be \"locations\" or \"location\"", metricsScope)
		}
		if tokens[1] != "global" {
			return nil, fmt.Errorf("invalid metricsScope format %q: second segment must be \"global\"", metricsScope)
		}
		if tokens[2] != "metricsScopes" {
			return nil, fmt.Errorf("invalid metricsScope format %q: third segment must be \"metricsScopes\"", metricsScope)
		}
		if tokens[3] == "" {
			return nil, fmt.Errorf("invalid metricsScope format %q: metricsScope ID cannot be empty", metricsScope)
		}
		metricsScopeID = tokens[3]
	} else if len(tokens) == 1 {
		if tokens[0] == "" {
			return nil, fmt.Errorf("metricsScope cannot be empty")
		}
		metricsScopeID = tokens[0]
	} else {
		return nil, fmt.Errorf("invalid metricsScope format %q: expected \"locations/global/metricsScopes/{metricsScopeID}\" or just \"{metricsScopeID}\"", metricsScope)
	}

	identity := &MonitoringMonitoredProjectIdentity{
		MetricsScope: metricsScopeID,
		Project:      resourceID,
	}
	return identity, nil
}

func (obj *MonitoringMonitoredProject) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromMonitoringMonitoredProjectSpec(ctx, reader, obj)
}
