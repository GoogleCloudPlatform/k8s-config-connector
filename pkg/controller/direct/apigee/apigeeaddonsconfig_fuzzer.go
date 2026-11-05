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

// +tool:fuzz-gen
// proto.message: google.golang.org/api/apigee/v1.GoogleCloudApigeeV1AddonsConfig
// api.group: apigee.cnrm.cloud.google.com

package apigee

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/apigee/v1"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer_NoProto(apigeeAddonsConfigFuzzer())
}

func apigeeAddonsConfigFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedSpecFuzzer_NoProto(&api.GoogleCloudApigeeV1AddonsConfig{},
		ApigeeAddonsConfigSpec_FromAPI, ApigeeAddonsConfigSpec_ToAPI,
	)

	// Field Comparison & Mapping Status:
	// KRM ApigeeAddonsConfigSpec Fields:
	// - org -> Parent organization ID in GCP REST URL path (organizations/{org}/addonsConfig), not part of GoogleCloudApigeeV1AddonsConfig struct.
	// - resourceID -> KCC resource ID, not part of GoogleCloudApigeeV1AddonsConfig struct.
	// - addonsConfig.advancedApiOpsConfig.enabled -> .AdvancedApiOpsConfig.Enabled (SpecField)
	// - addonsConfig.apiSecurityConfig.enabled -> .ApiSecurityConfig.Enabled (SpecField)
	// - addonsConfig.apiSecurityConfig.expiresAt -> .ApiSecurityConfig.ExpiresAt (SpecField)
	// - addonsConfig.connectorsPlatformConfig.enabled -> .ConnectorsPlatformConfig.Enabled (SpecField)
	// - addonsConfig.connectorsPlatformConfig.expiresAt -> .ConnectorsPlatformConfig.ExpiresAt (SpecField)
	// - addonsConfig.integrationConfig.enabled -> .IntegrationConfig.Enabled (SpecField)
	// - addonsConfig.monetizationConfig.enabled -> .MonetizationConfig.Enabled (SpecField)
	//
	// GCP API Fields:
	// - .AnalyticsConfig -> Only used for environment addons (organizations.environments.addonsConfig), not organization addons. Marked as Unimplemented_NotYetTriaged.
	f.SpecField(".AdvancedApiOpsConfig")
	f.SpecField(".ApiSecurityConfig")
	f.SpecField(".ConnectorsPlatformConfig")
	f.SpecField(".IntegrationConfig")
	f.SpecField(".MonetizationConfig")

	f.Unimplemented_NotYetTriaged(".AnalyticsConfig")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")
	f.Ignore_JSONBookkeeping(".AdvancedApiOpsConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".AdvancedApiOpsConfig.NullFields")
	f.Ignore_JSONBookkeeping(".ApiSecurityConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".ApiSecurityConfig.NullFields")
	f.Ignore_JSONBookkeeping(".ConnectorsPlatformConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".ConnectorsPlatformConfig.NullFields")
	f.Ignore_JSONBookkeeping(".IntegrationConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".IntegrationConfig.NullFields")
	f.Ignore_JSONBookkeeping(".MonetizationConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".MonetizationConfig.NullFields")

	return f
}
