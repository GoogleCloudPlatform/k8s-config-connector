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
// proto.message: google.cloud.privilegedaccessmanager.v1.Entitlement
// api.group: privilegedaccessmanager.cnrm.cloud.google.com

package privilegedaccessmanager

import (
	pb "cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzEntitlement())
}

func fuzzEntitlement() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Entitlement{},
		PrivilegedAccessManagerEntitlementSpec_FromProto, privilegedAccessManagerEntitlementSpec_ToProto,
		PrivilegedAccessManagerEntitlementObservedState_FromProto, PrivilegedAccessManagerEntitlementObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")                                           // special field
	f.Unimplemented_Internal(".privileged_access.gcp_iam_access.resource_type") // hidden duplicate field
	f.Unimplemented_Internal(".privileged_access.gcp_iam_access.resource")      // hidden duplicate field

	f.SpecField(".eligible_users")
	f.SpecField(".approval_workflow")
	f.SpecField(".max_request_duration")
	f.SpecField(".privileged_access")
	f.SpecField(".requester_justification_config")
	f.SpecField(".additional_notification_targets")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".etag")
	f.StatusField(".state")
	return f
}

func privilegedAccessManagerEntitlementSpec_ToProto(ctx *direct.MapContext, k *krm.PrivilegedAccessManagerEntitlementSpec) *pb.Entitlement {
	return PrivilegedAccessManagerEntitlementSpec_ToProto(ctx, k, gcpIAMAccessResource{})
}
