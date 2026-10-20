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
	_ identity.IdentityV2 = &LoggingLogScopeIdentity{}
)

var (
	LoggingLogScopeIdentityFormat = gcpurls.Template[LoggingLogScopeIdentity]("logging.googleapis.com", "projects/{project}/locations/{location}/logScopes/{logScope}")
)

// +k8s:deepcopy-gen=false
type LoggingLogScopeIdentity struct {
	Project  string
	Location string
	LogScope string
}

func (i *LoggingLogScopeIdentity) String() string {
	return LoggingLogScopeIdentityFormat.ToString(*i)
}

func (i *LoggingLogScopeIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *LoggingLogScopeIdentity) FromExternal(ref string) error {
	parsed, match, err := LoggingLogScopeIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of LoggingLogScope external=%q was not known (use %s): %w", ref, LoggingLogScopeIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of LoggingLogScope external=%q was not known (use %s)", ref, LoggingLogScopeIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *LoggingLogScopeIdentity) Host() string {
	return LoggingLogScopeIdentityFormat.Host()
}

func getIdentityFromLoggingLogScopeSpec(ctx context.Context, reader client.Reader, obj *LoggingLogScope) (*LoggingLogScopeIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location := obj.Spec.Location
	if location == "" {
		location = "global"
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &LoggingLogScopeIdentity{
		Project:  projectID,
		Location: location,
		LogScope: resourceID,
	}

	return identity, nil
}

func (obj *LoggingLogScope) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromLoggingLogScopeSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
