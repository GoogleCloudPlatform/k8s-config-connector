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

package privilegedaccessmanager

import (
	pb "cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
	"k8s.io/apimachinery/pkg/util/sets"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterFuzzer(entitlementSpecFuzzer().FuzzSpec)
	fuzztesting.RegisterFuzzer(entitlementObservedStateFuzzer().FuzzObservedState)
}

var entitlementKrmFields = fuzztesting.KRMFields{
	UnimplementedFields: sets.New(
		".name", // special field
		".privileged_access.gcp_iam_access.resource_type", // hidden duplicate field
		".privileged_access.gcp_iam_access.resource",      // hidden duplicate field
	),
	SpecFields: sets.New(".eligible_users",
		".approval_workflow",
		".max_request_duration",
		".privileged_access",
		".requester_justification_config",
		".additional_notification_targets"),
	ObservedStateFields: sets.New(".create_time",
		".update_time",
		".etag",
		".state"),
}

func entitlementSpecFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Entitlement{},
		PrivilegedAccessManagerEntitlementSpec_FromProto, privilegedAccessManagerEntitlementSpec_ToProto,
	)
	f.KRMFields = entitlementKrmFields
	return f
}

func entitlementObservedStateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Entitlement{},
		PrivilegedAccessManagerEntitlementObservedState_FromProto, PrivilegedAccessManagerEntitlementObservedState_ToProto,
	)
	f.KRMFields = entitlementKrmFields

	return f
}

func privilegedAccessManagerEntitlementSpec_ToProto(ctx *direct.MapContext, k *krm.PrivilegedAccessManagerEntitlementSpec) *pb.Entitlement {
	return PrivilegedAccessManagerEntitlementSpec_ToProto(ctx, k, gcpIAMAccessResource{})
}
