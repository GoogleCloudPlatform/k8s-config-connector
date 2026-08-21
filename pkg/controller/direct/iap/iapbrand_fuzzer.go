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
	fuzztesting.RegisterKRMFuzzer(iapBrandFuzzer())
}

func iapBrandFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Brand{},
		IAPBrandSpec_v1beta1_FromProto, IAPBrandSpec_v1beta1_ToProto,
		IAPBrandStatus_v1beta1_FromProto, IAPBrandStatus_v1beta1_ToProto,
	)

	// Comparison of KRM fields vs GCP Brand fields:
	// - .supportEmail maps to .support_email
	// - .applicationTitle maps to .application_title
	f.SpecField(".support_email")
	f.SpecField(".application_title")

	// - .orgInternalOnly maps to .org_internal_only
	f.StatusField(".org_internal_only")

	// - .name represents the unique resource identifier (URL) of the Brand
	f.Unimplemented_Identity(".name")

	return f
}
