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

package agentregistry

import (
	pb "cloud.google.com/go/agentregistry/apiv1/agentregistrypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzAgentRegistryBinding())
}

func fuzzAgentRegistryBinding() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Binding{},
		AgentRegistryBindingSpec_FromProto, AgentRegistryBindingSpec_ToProto,
		AgentRegistryBindingObservedState_FromProto, AgentRegistryBindingObservedState_ToProto,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".source")
	f.SpecField(".target")
	f.SpecField(".auth_provider_binding")

	// Identity field
	f.Unimplemented_Identity(".name")

	// Status fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
