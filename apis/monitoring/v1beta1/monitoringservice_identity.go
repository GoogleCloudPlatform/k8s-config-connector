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
	_ identity.IdentityV2 = &MonitoringServiceIdentity{}
	_ identity.Resource   = &MonitoringService{}
)

var MonitoringServiceIdentityFormat = gcpurls.Template[MonitoringServiceIdentity]("monitoring.googleapis.com", "projects/{project}/services/{service}")

// MonitoringServiceIdentity is the identity of a GCP MonitoringService resource.
// +k8s:deepcopy-gen=false
type MonitoringServiceIdentity struct {
	Project string
	Service string
}

func (i *MonitoringServiceIdentity) String() string {
	return MonitoringServiceIdentityFormat.ToString(*i)
}

func (i *MonitoringServiceIdentity) FromExternal(ref string) error {
	normalized := stripMonitoringPrefixes(ref)

	parsed, match, err := MonitoringServiceIdentityFormat.Parse(normalized)
	if err != nil {
		return fmt.Errorf("format of MonitoringService external=%q was not known (use %s): %w", ref, MonitoringServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MonitoringService external=%q was not known (use %s)", ref, MonitoringServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func stripMonitoringPrefixes(ref string) string {
	normalized := ref
	normalized = strings.TrimPrefix(normalized, "https:")
	normalized = strings.TrimPrefix(normalized, "http:")
	normalized = strings.TrimPrefix(normalized, "//")
	normalized = strings.TrimPrefix(normalized, "monitoring.googleapis.com/")
	normalized = strings.TrimPrefix(normalized, "v3/")
	return normalized
}

func (i *MonitoringServiceIdentity) Host() string {
	return MonitoringServiceIdentityFormat.Host()
}

func getIdentityFromMonitoringServiceSpec(ctx context.Context, reader client.Reader, obj client.Object) (*MonitoringServiceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &MonitoringServiceIdentity{
		Project: projectID,
		Service: resourceID,
	}
	return identity, nil
}

func (obj *MonitoringService) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromMonitoringServiceSpec(ctx, reader, obj)
}
