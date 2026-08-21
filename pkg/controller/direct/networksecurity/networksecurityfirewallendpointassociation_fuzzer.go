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

package networksecurity

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/networksecurity/v1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(networksecurityFirewallEndpointAssociationFuzzer())
}

func networksecurityFirewallEndpointAssociationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.FirewallEndpointAssociation{},
		NetworkSecurityFirewallEndpointAssociationSpec_FromAPI, NetworkSecurityFirewallEndpointAssociationSpec_ToAPI,
		NetworkSecurityFirewallEndpointAssociationObservedState_FromAPI, NetworkSecurityFirewallEndpointAssociationObservedState_ToAPI,
	)

	f.SpecField(".Disabled")
	f.SpecField(".Labels")

	f.StatusField(".CreateTime")
	f.StatusField(".UpdateTime")
	f.StatusField(".Reconciling")
	f.StatusField(".State")

	f.UnimplementedFields.Insert(".Name")
	f.UnimplementedFields.Insert(".Network")
	f.UnimplementedFields.Insert(".FirewallEndpoint")
	f.UnimplementedFields.Insert(".TlsInspectionPolicy")

	f.Ignore_JSONBookkeeping(".ServerResponse")
	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")

	f.FilterSpec = func(in *api.FirewallEndpointAssociation) {
		in.CreateTime = ""
		in.UpdateTime = ""
		in.State = ""
		in.Reconciling = false
	}

	f.FilterStatus = func(in *api.FirewallEndpointAssociation) {
		in.Disabled = false
		in.Labels = nil
	}

	return f
}
