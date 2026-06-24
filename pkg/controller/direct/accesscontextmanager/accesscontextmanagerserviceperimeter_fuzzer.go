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
// proto.message: google.identity.accesscontextmanager.v1.ServicePerimeter
// api.group: accesscontextmanager.cnrm.cloud.google.com

package accesscontextmanager

import (
	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(accessContextManagerServicePerimeterFuzzer())
}

func accessContextManagerServicePerimeterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ServicePerimeter{},
		AccessContextManagerServicePerimeterSpec_FromProto, AccessContextManagerServicePerimeterSpec_ToProto,
		AccessContextManagerServicePerimeterStatus_FromProto, AccessContextManagerServicePerimeterStatus_ToProto,
	)

	f.IdentityField(".name")

	f.SpecField(".title")
	f.SpecField(".description")
	f.SpecField(".perimeter_type")
	f.SpecField(".status")
	f.SpecField(".spec")
	f.SpecField(".use_explicit_dry_run_spec")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
