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
	_ identity.ServerGeneratedIdentity = &MonitoringAlertPolicyIdentity{}
	_ identity.Resource                = &MonitoringAlertPolicy{}
)

var MonitoringAlertPolicyIdentityFormat = gcpurls.Template[MonitoringAlertPolicyIdentity]("monitoring.googleapis.com", "projects/{project}/alertPolicies/{alertpolicy}")

// MonitoringAlertPolicyIdentity is the identity of a GCP MonitoringAlertPolicy resource.
// +k8s:deepcopy-gen=false
type MonitoringAlertPolicyIdentity struct {
	Project     string
	AlertPolicy string
}

func (i *MonitoringAlertPolicyIdentity) String() string {
	return MonitoringAlertPolicyIdentityFormat.ToString(*i)
}

func (i *MonitoringAlertPolicyIdentity) ParentString() string {
	return "projects/" + i.Project
}

func (obj *MonitoringAlertPolicy) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromMonitoringAlertPolicySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	statusName := common.ValueOf(obj.Status.Name)
	if statusName != "" {
		statusIdentity := &MonitoringAlertPolicyIdentity{}
		if err := statusIdentity.FromExternal(statusName); err != nil {
			return nil, err
		}

		if specIdentity.AlertPolicy == "" {
			specIdentity.AlertPolicy = statusIdentity.AlertPolicy
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change MonitoringAlertPolicy identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (i *MonitoringAlertPolicyIdentity) FromExternal(ref string) error {
	normalized := ref
	normalized = strings.TrimPrefix(normalized, "https:")
	normalized = strings.TrimPrefix(normalized, "http:")
	normalized = strings.TrimPrefix(normalized, "//")
	normalized = strings.TrimPrefix(normalized, "monitoring.googleapis.com/")
	normalized = strings.TrimPrefix(normalized, "v3/")

	// This relative format is a very special case and unusual for a GCP API.
	// For example, the IncidentList resource under Monitoring Dashboard explicitly
	// asks that we do NOT pass the project in the policy name (e.g. use alertPolicies/utilization).
	// Reference: https://docs.cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards#incidentlist
	//
	// NOTE: This is a design mistake in the KCC API, but we have to support it for legacy compatibility.
	// The mistake is this: Just because the GCP API requires a particular format, that does not mean
	// that we should mirror that format in the Ref. KCC is supposed to be consistent, regardless of
	// the underlying GCP API.
	if strings.HasPrefix(normalized, "alertPolicies/") {
		parts := strings.Split(normalized, "/")
		if len(parts) == 2 && parts[1] != "" {
			i.Project = ""
			i.AlertPolicy = parts[1]
			return nil
		}
	}

	parsed, match, err := MonitoringAlertPolicyIdentityFormat.Parse(normalized)
	if err != nil {
		return fmt.Errorf("format of MonitoringAlertPolicy external=%q was not known (use %s): %w", ref, MonitoringAlertPolicyIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MonitoringAlertPolicy external=%q was not known (use %s)", ref, MonitoringAlertPolicyIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MonitoringAlertPolicyIdentity) Host() string {
	return MonitoringAlertPolicyIdentityFormat.Host()
}

func (i *MonitoringAlertPolicyIdentity) HasIdentitySpecified() bool {
	return i.AlertPolicy != ""
}

func getIdentityFromMonitoringAlertPolicySpec(ctx context.Context, reader client.Reader, obj *MonitoringAlertPolicy) (*MonitoringAlertPolicyIdentity, error) {
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &MonitoringAlertPolicyIdentity{
		Project:     projectID,
		AlertPolicy: resourceID,
	}
	return identity, nil
}
