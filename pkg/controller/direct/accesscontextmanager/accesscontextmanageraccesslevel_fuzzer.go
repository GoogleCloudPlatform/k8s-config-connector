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
// proto.message: google.identity.accesscontextmanager.v1.AccessLevel
// api.group: accesscontextmanager.cnrm.cloud.google.com

package accesscontextmanager

import (
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/genproto/googleapis/type/expr"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(AccessContextManagerAccessLevelFuzzer())
}

func AccessContextManagerAccessLevelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AccessLevel{},
		AccessContextManagerAccessLevelSpec_FromProto, AccessContextManagerAccessLevelSpec_ToProto,
		AccessContextManagerAccessLevelObservedState_FromProto, AccessContextManagerAccessLevelObservedState_ToProto,
	)

	f.SpecField(".title")
	f.SpecField(".description")
	f.SpecField(".basic")
	f.SpecField(".custom")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")

	f.Unimplemented_NotYetTriaged(".basic.conditions[].ip_subnetworks")
	f.Unimplemented_NotYetTriaged(".basic.conditions[].device_policy.require_screenlock")
	f.Unimplemented_NotYetTriaged(".basic.conditions[].device_policy.os_constraints")

	f.FilterSpec = func(in *pb.AccessLevel) {
		if in.GetCustom() != nil {
			if in.GetCustom().Expr == nil {
				in.GetCustom().Expr = &expr.Expr{}
			}
		}
	}

	return f
}
