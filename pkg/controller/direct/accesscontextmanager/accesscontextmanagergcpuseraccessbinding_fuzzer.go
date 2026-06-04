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
// proto.message: google.identity.accesscontextmanager.v1.GcpUserAccessBinding
// api.group: accesscontextmanager.cnrm.cloud.google.com

package accesscontextmanager

import (
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(accessContextManagerGCPUserAccessBindingFuzzer())
}

func accessContextManagerGCPUserAccessBindingFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.GcpUserAccessBinding{},
		AccessContextManagerGCPUserAccessBindingSpec_FromProto, AccessContextManagerGCPUserAccessBindingSpec_ToProto,
		AccessContextManagerGCPUserAccessBindingObservedState_FromProto, AccessContextManagerGCPUserAccessBindingObservedState_ToProto,
	)

	f.SpecFields.Insert(".group_key")
	f.SpecFields.Insert(".access_levels")

	f.UnimplementedFields.Insert(".name")

	return f
}
