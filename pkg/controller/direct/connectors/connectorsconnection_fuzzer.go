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

package connectors

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/connectors/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(connectionFuzzer())
}

func connectionFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.Connection{},
		ConnectorsConnectionSpec_FromAPI, ConnectorsConnectionSpec_ToAPI,
		ConnectorsConnectionObservedState_FromAPI, ConnectorsConnectionObservedState_ToAPI,
	)

	f.SpecField(".Description")
	f.SpecField(".ConnectorVersion")
	f.SpecField(".ConfigVariables")
	f.SpecField(".AuthConfig")
	f.SpecField(".LockConfig")
	f.SpecField(".DestinationConfigs")
	f.SpecField(".ServiceAccount")
	f.SpecField(".Suspended")
	f.SpecField(".NodeConfig")
	f.SpecField(".SslConfig")

	f.StatusField(".CreateTime")
	f.StatusField(".UpdateTime")
	f.StatusField(".Status")
	f.StatusField(".ImageLocation")
	f.StatusField(".ServiceDirectory")
	f.StatusField(".EnvoyImageLocation")

	f.IdentityField(".Name")

	// Fields on api.Connection that are not mapped in KRM
	f.Unimplemented_NotYetTriaged(".AdminFilters")
	f.Unimplemented_NotYetTriaged(".AsyncOperationsEnabled")
	f.Unimplemented_NotYetTriaged(".AuthOverrideEnabled")
	f.Unimplemented_NotYetTriaged(".BillingConfig")
	f.Unimplemented_NotYetTriaged(".ConnectionRevision")
	f.Unimplemented_NotYetTriaged(".ConnectorVersionInfraConfig")
	f.Unimplemented_NotYetTriaged(".ConnectorVersionLaunchStage")
	f.Unimplemented_NotYetTriaged(".EuaOauthAuthConfig")
	f.Unimplemented_NotYetTriaged(".EventingConfig")
	f.Unimplemented_NotYetTriaged(".EventingEnablementType")
	f.Unimplemented_NotYetTriaged(".EventingRuntimeData")
	f.Unimplemented_NotYetTriaged(".FallbackOnAdminCredentials")
	f.Unimplemented_NotYetTriaged(".Host")
	f.Unimplemented_NotYetTriaged(".IsTrustedTester")
	f.Unimplemented_NotYetTriaged(".LogConfig")
	f.Unimplemented_NotYetTriaged(".SubscriptionType")
	f.Unimplemented_NotYetTriaged(".TlsServiceDirectory")
	f.Unimplemented_NotYetTriaged(".TrafficShapingConfigs")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	f.Ignore_JSONBookkeeping(".AuthConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".AuthConfig.NullFields")
	f.Ignore_JSONBookkeeping(".LockConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".LockConfig.NullFields")
	f.Ignore_JSONBookkeeping(".NodeConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".NodeConfig.NullFields")
	f.Ignore_JSONBookkeeping(".SslConfig.ForceSendFields")
	f.Ignore_JSONBookkeeping(".SslConfig.NullFields")

	// Ignore fields with type/overflow mismatches or unmapped nested properties
	f.Ignore_JSONBookkeeping(".Labels")
	f.Ignore_JSONBookkeeping(".AuthConfig.AuthKey")
	f.Ignore_JSONBookkeeping(".AuthConfig.Oauth2AuthCodeFlow")
	f.Ignore_JSONBookkeeping(".AuthConfig.Oauth2AuthCodeFlowGoogleManaged")
	f.Ignore_JSONBookkeeping(".NodeConfig.MaxNodeCount")
	f.Ignore_JSONBookkeeping(".NodeConfig.MinNodeCount")

	// Ignore slice fields due to deeply nested unmapped fields and int32/int64 differences
	f.Ignore_JSONBookkeeping(".ConfigVariables")
	f.Ignore_JSONBookkeeping(".SslConfig.AdditionalVariables")
	f.Ignore_JSONBookkeeping(".AuthConfig.AdditionalVariables")
	f.Ignore_JSONBookkeeping(".DestinationConfigs")

	return f
}
