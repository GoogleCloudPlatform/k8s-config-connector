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
// proto.message: google.cloud.apigeeregistry.v1.Instance
// api.group: apigeeregistry.cnrm.cloud.google.com

package apigeeregistryinstance

import (
	pb "cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb"
	mappers "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apigeeregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apigeeRegistryInstanceFuzzer())
}

func apigeeRegistryInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		mappers.ApigeeRegistryInstanceSpec_FromProto, mappers.ApigeeRegistryInstanceSpec_ToProto,
		mappers.ApigeeRegistryInstanceObservedState_FromProto, mappers.ApigeeRegistryInstanceObservedState_ToProto,
	)

	f.SpecField(".config")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".state_message")

	f.Unimplemented_NotYetTriaged(".config.location")
	f.Unimplemented_Identity(".name")

	return f
}
