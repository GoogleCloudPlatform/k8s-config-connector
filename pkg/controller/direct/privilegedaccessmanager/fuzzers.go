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

	f.UnimplementedFields.Insert(".name")                                           // special field
	f.UnimplementedFields.Insert(".privileged_access.gcp_iam_access.resource_type") // hidden duplicate field
	f.UnimplementedFields.Insert(".privileged_access.gcp_iam_access.resource")      // hidden duplicate field

	f.SpecFields.Insert(".eligible_users")
	f.SpecFields.Insert(".approval_workflow")
	f.SpecFields.Insert(".max_request_duration")
	f.SpecFields.Insert(".privileged_access")
	f.SpecFields.Insert(".requester_justification_config")
	f.SpecFields.Insert(".additional_notification_targets")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".state")
	return f
}

func privilegedAccessManagerEntitlementSpec_ToProto(ctx *direct.MapContext, k *krm.PrivilegedAccessManagerEntitlementSpec) *pb.Entitlement {
	return PrivilegedAccessManagerEntitlementSpec_ToProto(ctx, k, gcpIAMAccessResource{})
}
