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
	fuzztesting.RegisterKRMFuzzer(PrivateCACertificateTemplateFuzzer())
}

func PrivateCACertificateTemplateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CertificateTemplate{},
		PrivateCACertificateTemplateSpec_FromProto, PrivateCACertificateTemplateSpec_ToProto,
		PrivateCACertificateTemplateStatus_FromProto, PrivateCACertificateTemplateStatus_ToProto,
	)

	// Explicitly compare the KRM Spec fields with fuzzer/proto fields:
	// - Spec.Description maps to .description
	// - Spec.IdentityConstraints maps to .identity_constraints
	// - Spec.PassthroughExtensions maps to .passthrough_extensions
	// - Spec.PredefinedValues maps to .predefined_values
	// - Spec.ProjectRef is handled externally as identity
	// - Spec.Location is handled externally as identity
	// - Spec.ResourceID is handled externally as identity

	// Spec fields
	f.SpecField(".description")

	f.SpecField(".identity_constraints")
	f.SpecField(".identity_constraints.allow_subject_alt_names_passthrough")
	f.SpecField(".identity_constraints.allow_subject_passthrough")
	f.SpecField(".identity_constraints.cel_expression")
	f.SpecField(".identity_constraints.cel_expression.expression")
	f.SpecField(".identity_constraints.cel_expression.title")
	f.SpecField(".identity_constraints.cel_expression.description")
	f.SpecField(".identity_constraints.cel_expression.location")

	f.SpecField(".passthrough_extensions")
	f.SpecField(".passthrough_extensions.known_extensions")
	f.SpecField(".passthrough_extensions.additional_extensions")
	f.SpecField(".passthrough_extensions.additional_extensions.object_id_path")

	f.SpecField(".predefined_values")
	f.SpecField(".predefined_values.key_usage")
	f.SpecField(".predefined_values.key_usage.base_key_usage")
	f.SpecField(".predefined_values.key_usage.base_key_usage.digital_signature")
	f.SpecField(".predefined_values.key_usage.base_key_usage.content_commitment")
	f.SpecField(".predefined_values.key_usage.base_key_usage.key_encipherment")
	f.SpecField(".predefined_values.key_usage.base_key_usage.data_encipherment")
	f.SpecField(".predefined_values.key_usage.base_key_usage.key_agreement")
	f.SpecField(".predefined_values.key_usage.base_key_usage.cert_sign")
	f.SpecField(".predefined_values.key_usage.base_key_usage.crl_sign")
	f.SpecField(".predefined_values.key_usage.base_key_usage.encipher_only")
	f.SpecField(".predefined_values.key_usage.base_key_usage.decipher_only")

	f.SpecField(".predefined_values.key_usage.extended_key_usage")
	f.SpecField(".predefined_values.key_usage.extended_key_usage.server_auth")
	f.SpecField(".predefined_values.key_usage.extended_key_usage.client_auth")
	f.SpecField(".predefined_values.key_usage.extended_key_usage.code_signing")
	f.SpecField(".predefined_values.key_usage.extended_key_usage.email_protection")
	f.SpecField(".predefined_values.key_usage.extended_key_usage.time_stamping")
	f.SpecField(".predefined_values.key_usage.extended_key_usage.ocsp_signing")

	f.SpecField(".predefined_values.key_usage.unknown_extended_key_usages")
	f.SpecField(".predefined_values.key_usage.unknown_extended_key_usages.object_id_path")

	f.SpecField(".predefined_values.ca_options")
	f.SpecField(".predefined_values.ca_options.is_ca")
	f.SpecField(".predefined_values.ca_options.max_issuer_path_length")

	f.SpecField(".predefined_values.policy_ids")
	f.SpecField(".predefined_values.policy_ids.object_id_path")

	f.SpecField(".predefined_values.aia_ocsp_servers")

	f.SpecField(".predefined_values.additional_extensions")
	f.SpecField(".predefined_values.additional_extensions.object_id")
	f.SpecField(".predefined_values.additional_extensions.object_id.object_id_path")
	f.SpecField(".predefined_values.additional_extensions.critical")
	f.SpecField(".predefined_values.additional_extensions.value")

	// Status fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	// Identity/unimplemented fields
	f.Unimplemented_Identity(".name")

	f.Unimplemented_NotYetTriaged(".maximum_lifetime")
	f.Unimplemented_NotYetTriaged(".predefined_values.name_constraints")
	f.Unimplemented_NotYetTriaged(".labels")

	return f
}
