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
	_ identity.IdentityV2 = &LoggingLogMetricIdentity{}
	_ identity.Resource   = &LoggingLogMetric{}
)

var LoggingLogMetricIdentityFormat = gcpurls.Template[LoggingLogMetricIdentity]("logging.googleapis.com", "projects/{project}/metrics/{metric}")

// +k8s:deepcopy-gen=false
type LoggingLogMetricIdentity struct {
	Project string
	Metric  string
}

func (i *LoggingLogMetricIdentity) String() string {
	return LoggingLogMetricIdentityFormat.ToString(*i)
}

func (i *LoggingLogMetricIdentity) ParentString() string {
	return "projects/" + i.Project
}

func (i *LoggingLogMetricIdentity) FromExternal(ref string) error {
	parsed, match, err := LoggingLogMetricIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of LoggingLogMetric external=%q was not known (use %s): %w", ref, LoggingLogMetricIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of LoggingLogMetric external=%q was not known (use %s)", ref, LoggingLogMetricIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *LoggingLogMetricIdentity) Host() string {
	return LoggingLogMetricIdentityFormat.Host()
}

func getIdentityFromLoggingLogMetricSpec(ctx context.Context, reader client.Reader, obj *LoggingLogMetric) (*LoggingLogMetricIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &LoggingLogMetricIdentity{
		Project: projectID,
		Metric:  resourceID,
	}
	return identity, nil
}

func (obj *LoggingLogMetric) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromLoggingLogMetricSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
