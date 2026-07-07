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
	_ identity.IdentityV2 = &ErrorReportingNotificationSettingsIdentity{}
	_ identity.Resource   = &ErrorReportingNotificationSettings{}
)

var ErrorReportingNotificationSettingsIdentityFormat = gcpurls.Template[ErrorReportingNotificationSettingsIdentity]("clouderrorreporting.googleapis.com", "projects/{project}/locations/global/notificationSettings")

// +k8s:deepcopy-gen=false
type ErrorReportingNotificationSettingsIdentity struct {
	Project string
}

func (i *ErrorReportingNotificationSettingsIdentity) String() string {
	return ErrorReportingNotificationSettingsIdentityFormat.ToString(*i)
}

func (i *ErrorReportingNotificationSettingsIdentity) FromExternal(ref string) error {
	parsed, match, err := ErrorReportingNotificationSettingsIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ErrorReportingNotificationSettings external=%q was not known (use %s): %w", ref, ErrorReportingNotificationSettingsIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ErrorReportingNotificationSettings external=%q was not known (use %s)", ref, ErrorReportingNotificationSettingsIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ErrorReportingNotificationSettingsIdentity) Host() string {
	return ErrorReportingNotificationSettingsIdentityFormat.Host()
}

func getIdentityFromErrorReportingNotificationSettings(ctx context.Context, reader client.Reader, obj *ErrorReportingNotificationSettings) (*ErrorReportingNotificationSettingsIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ErrorReportingNotificationSettingsIdentity{
		Project: projectID,
	}
	return identity, nil
}

func (obj *ErrorReportingNotificationSettings) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromErrorReportingNotificationSettings(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	return specIdentity, nil
}
