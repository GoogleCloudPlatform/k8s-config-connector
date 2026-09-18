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
// proto.message: google.cloud.numberregistry.v1alpha.IpamAdminScope
// api.group: cloudnumberregistry.cnrm.cloud.google.com

package cloudnumberregistry

import (
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/cloudnumberregistry/pb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(CloudNumberRegistryIpamAdminScopeFuzzer())
}

func CloudNumberRegistryIpamAdminScopeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.IpamAdminScope{},
		CloudNumberRegistryIpamAdminScopeSpec_FromProto, CloudNumberRegistryIpamAdminScopeSpec_ToProto,
		CloudNumberRegistryIpamAdminScopeObservedState_FromProto, CloudNumberRegistryIpamAdminScopeObservedState_ToProto,
	)

	f.SpecField(".projectRef")
	f.SpecField(".location")
	f.SpecField(".resourceID")
	f.SpecField(".enabledAddonPlatforms")
	f.SpecField(".scopeRefs")
	f.SpecField(".labels")

	f.StatusField(".conditions")
	f.StatusField(".observedGeneration")
	f.StatusField(".externalRef")
	f.StatusField(".observedState")

	// Proto fields not in Spec
	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".state")

	// Proto fields not in ObservedState
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".enabled_addon_platforms")
	f.Unimplemented_NotYetTriaged(".scopes")

	return f
}
