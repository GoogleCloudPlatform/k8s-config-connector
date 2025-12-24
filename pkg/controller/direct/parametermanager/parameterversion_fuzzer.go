// Copyright 2025 Google LLC
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
// proto.message: google.cloud.parametermanager.v1.ParameterVersion
// api.group: parametermanager.cnrm.cloud.google.com

package parametermanager

import (
	pb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(parameterVersionFuzzer())
}

func parameterVersionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ParameterVersion{},
		ParameterManagerParameterVersionSpec_FromProto, ParameterManagerParameterVersionSpec_ToProto,
		ParameterManagerParameterVersionObservedState_FromProto, ParameterManagerParameterVersionObservedState_ToProto,
	)

	f.SpecFields.Insert(".disabled")
	f.SpecFields.Insert(".payload")

	f.StatusFields.Insert(".name")            // Output Only
	f.StatusFields.Insert(".create_time")     // Output Only
	f.StatusFields.Insert(".update_time")     // Output Only
	f.StatusFields.Insert(".kms_key_version") // Output Only

	return f
}
