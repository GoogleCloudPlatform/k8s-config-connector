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

package filestore

// +tool:crd-fuzzer
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.resource: Instance
// crd.type: FilestoreInstance

import (
	pb "cloud.google.com/go/filestore/apiv1/filestorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(instanceFuzzer())
}

func instanceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		FilestoreInstanceSpec_FromProto, FilestoreInstanceSpec_ToProto,
		FilestoreInstanceObservedState_FromProto, FilestoreInstanceObservedState_ToProto,
	)

	// Special/system fields
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".etag")

	// Not yet implemented
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".file_shares[].source_backup")
	f.UnimplementedFields.Insert(".networks[].ip_addresses") // Unclear why this wasn't generated

	// Input fields
	f.SpecFields.Insert(".tier")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".file_shares")
	f.SpecFields.Insert(".networks")
	f.SpecFields.Insert(".kms_key_name")

	// Output fields
	f.StatusFields.Insert(".satisfies_pzs")
	f.StatusFields.Insert(".satisfies_pzi")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".status_message")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".suspension_reasons")

	return f
}
