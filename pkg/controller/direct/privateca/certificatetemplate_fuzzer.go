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
// proto.message: google.cloud.security.privateca.v1.CertificateTemplate
// api.group: privateca.cnrm.cloud.google.com

package privateca

import (
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(privatecaCertificateTemplateFuzzer())
}

func privatecaCertificateTemplateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CertificateTemplate{},
		PrivateCACertificateTemplateSpec_FromProto, PrivateCACertificateTemplateSpec_ToProto,
		PrivateCACertificateTemplateStatus_FromProto, PrivateCACertificateTemplateStatus_ToProto,
	)

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".identity_constraints")
	f.SpecField(".passthrough_extensions")
	f.SpecField(".predefined_values")

	// Status fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	// Unimplemented / identity fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".maximum_lifetime")
	f.Unimplemented_NotYetTriaged(".predefined_values.name_constraints")

	return f
}
