// Copyright 2024 Google LLC
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
// proto.message: google.cloud.securesourcemanager.v1.Instance
// api.group: securesourcemanager.cnrm.cloud.google.com

package securesourcemanager

import (
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(secureSourceManagerInstanceFuzzer())
}

func secureSourceManagerInstanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		SecureSourceManagerInstanceSpec_FromProto, SecureSourceManagerInstanceSpec_ToProto,
		SecureSourceManagerInstanceObservedState_FromProto, SecureSourceManagerInstanceObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")        // Identifier
	f.UnimplementedFields.Insert(".create_time") // Output only
	f.UnimplementedFields.Insert(".update_time") // Output only
	f.UnimplementedFields.Insert(".labels")      // NOTYET
	f.UnimplementedFields.Insert(".private_config.ssh_service_attachment")
	f.UnimplementedFields.Insert(".private_config.http_service_attachment")

	f.SpecFields.Insert(".private_config")
	f.SpecFields.Insert(".kms_key")

	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_note")
	f.StatusFields.Insert(".host_config")

	return f
}
