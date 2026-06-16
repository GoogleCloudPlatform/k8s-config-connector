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
// proto.message: google.cloud.iap.v1.Brand

package iap

import (
	pb "cloud.google.com/go/iap/apiv1/iappb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(iapBrandFuzzer())
}

func iapBrandFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.Brand{},
		IAPBrandSpec_v1beta1_FromProto, IAPBrandSpec_v1beta1_ToProto,
	)

	// Spec fields
	f.SpecFields.Insert(".application_title")
	f.SpecFields.Insert(".support_email")

	f.IdentityField(".name")
	f.Unimplemented_NotYetTriaged(".org_internal_only")

	return f
}
