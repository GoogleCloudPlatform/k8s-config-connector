// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.api.serviceusage.v1beta1.ServiceIdentity
// api.group: serviceusage.cnrm.cloud.google.com
// crd.kind: ServiceIdentity

package serviceusage

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "google.golang.org/genproto/googleapis/api/serviceusage/v1beta1"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(serviceIdentityFuzzer())
}

func serviceIdentityFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ServiceIdentity{},
		ServiceIdentitySpec_FromProto, ServiceIdentitySpec_ToProto,
		ServiceIdentityObservedState_FromProto, ServiceIdentityObservedState_ToProto,
	)

	f.StatusFields.Insert(".email")
	f.StatusFields.Insert(".unique_id")

	return f
}
