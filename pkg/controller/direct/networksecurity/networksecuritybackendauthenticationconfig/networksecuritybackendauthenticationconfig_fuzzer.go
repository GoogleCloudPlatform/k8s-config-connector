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

package networksecuritybackendauthenticationconfig

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networksecurity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(NetworkSecurityBackendAuthenticationConfigFuzzer())
}

func NetworkSecurityBackendAuthenticationConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackendAuthenticationConfig{},
		NetworkSecurityBackendAuthenticationConfigSpec_FromProto, NetworkSecurityBackendAuthenticationConfigSpec_ToProto,
		NetworkSecurityBackendAuthenticationConfigObservedState_FromProto, NetworkSecurityBackendAuthenticationConfigObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".client_certificate")
	f.SpecFields.Insert(".trust_config")
	f.SpecFields.Insert(".well_known_roots")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".etag")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
