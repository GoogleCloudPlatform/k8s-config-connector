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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &MonitoringMetricDescriptorIdentity{}
	_ identity.Resource   = &MonitoringMetricDescriptor{}
)

var MonitoringMetricDescriptorIdentityFormat = gcpurls.Template[MonitoringMetricDescriptorIdentity]("monitoring.googleapis.com", "projects/{project}/metricDescriptors/{metricid}")

// MonitoringMetricDescriptorIdentity is the identity of a Google Cloud MonitoringMetricDescriptor resource.
// +k8s:deepcopy-gen=false
type MonitoringMetricDescriptorIdentity struct {
	Project  string
	MetricID string
}

func (i *MonitoringMetricDescriptorIdentity) String() string {
	return MonitoringMetricDescriptorIdentityFormat.ToString(*i)
}

func (i *MonitoringMetricDescriptorIdentity) FromExternal(ref string) error {
	s := ref
	s = strings.TrimPrefix(s, "https:")
	s = strings.TrimPrefix(s, "http:")
	s = strings.TrimPrefix(s, "//")
	s = strings.TrimPrefix(s, "monitoring.googleapis.com/")
	s = strings.Trim(s, "/")

	tokens := strings.Split(s, "/")
	if len(tokens) >= 4 && tokens[0] == "projects" && tokens[2] == "metricDescriptors" {
		i.Project = tokens[1]
		i.MetricID = strings.Join(tokens[3:], "/")
		return nil
	}

	return fmt.Errorf("format of MonitoringMetricDescriptor external=%q was not known (use %s)", ref, MonitoringMetricDescriptorIdentityFormat.CanonicalForm())
}

func (i *MonitoringMetricDescriptorIdentity) Host() string {
	return MonitoringMetricDescriptorIdentityFormat.Host()
}

func getIdentityFromMonitoringMetricDescriptorSpec(ctx context.Context, reader client.Reader, obj *MonitoringMetricDescriptor) (*MonitoringMetricDescriptorIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	metricID := obj.Spec.Type
	if metricID == "" {
		return nil, fmt.Errorf("spec.type is empty or not set")
	}

	identity := &MonitoringMetricDescriptorIdentity{
		Project:  projectID,
		MetricID: metricID,
	}
	return identity, nil
}

func (obj *MonitoringMetricDescriptor) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMonitoringMetricDescriptorSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &MonitoringMetricDescriptorIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change MonitoringMetricDescriptor identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
