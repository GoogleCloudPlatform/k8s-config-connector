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
// proto.message: google.cloud.metastore.v1.Federation
// api.group: metastore.cnrm.cloud.google.com

package metastore

import (
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(metastoreFederationFuzzer())
}

func metastoreFederationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Federation{},
		MetastoreFederationSpec_FromProto, MetastoreFederationSpec_ToProto,
		MetastoreFederationObservedState_FromProto, MetastoreFederationObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".version")
	f.SpecFields.Insert(".backend_metastores")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".endpoint_uri")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".state_message")
	f.StatusFields.Insert(".uid")

	f.UnimplementedFields.Insert(".name") // special field
	return f
}
