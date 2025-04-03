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
// proto.message: google.cloud.securesourcemanager.v1.Repository
// api.group: securesourcemanager.cnrm.cloud.google.com

package securesourcemanager

import (
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(secureSourceManagerRepositoryFuzzer())
}

func secureSourceManagerRepositoryFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Repository{},
		SecureSourceManagerRepositorySpec_FromProto, SecureSourceManagerRepositorySpec_ToProto,
		SecureSourceManagerRepositoryObservedState_FromProto, SecureSourceManagerRepositoryObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // Identifier
	f.UnimplementedFields.Insert(".description")
	f.UnimplementedFields.Insert(".uid")
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".update_time")
	f.UnimplementedFields.Insert(".etag")
	f.UnimplementedFields.Insert(".uris")

	f.SpecFields.Insert(".instance")
	f.SpecFields.Insert(".initial_config")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".uris")

	return f
}
