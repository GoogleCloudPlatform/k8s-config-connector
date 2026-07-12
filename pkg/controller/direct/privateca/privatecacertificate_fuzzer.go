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
// See the License for the_license.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.security.privateca.v1.Certificate
// api.group: privateca.cnrm.cloud.google.com

package privateca

import (
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/types/known/durationpb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(PrivateCACertificateFuzzer())
}

func PrivateCACertificateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Certificate{},
		PrivateCACertificateSpec_FromProto, PrivateCACertificateSpec_ToProto,
		PrivateCACertificateStatus_FromProto, PrivateCACertificateStatus_ToProto,
	)

	f.FilterSpec = func(in *pb.Certificate) {
		if in.Lifetime == nil {
			in.Lifetime = &durationpb.Duration{Seconds: 3600}
		}
		if cfg, ok := in.CertificateConfig.(*pb.Certificate_Config); ok && cfg != nil && cfg.Config != nil {
			if cfg.Config.SubjectConfig == nil {
				cfg.Config.SubjectConfig = &pb.CertificateConfig_SubjectConfig{
					Subject: &pb.Subject{},
				}
			}
		}
	}

	// Documented Field Comparison (KRM Spec -> Proto):
	// - projectRef (refs.ProjectRef): Mapped to parent project context, not directly in Certificate proto.
	// - location (string): Mapped to parent location context, not directly in Certificate proto.
	// - caPoolRef (privatecarefs.PrivateCACAPoolRef): Mapped to parent CA Pool context, not directly in Certificate proto.
	// - certificateAuthorityRef (privatecarefs.PrivateCACertificateAuthorityRef): Mapped to parent CA context, not directly in Certificate proto.
	// - certificateTemplateRef (*PrivateCACertificateTemplateRef): Mapped to .certificate_template string.
	// - resourceID (*string): Mapped to the last segment of the .name field in the Certificate proto.
	// - config (*Certificate_Config): Mapped to .config (under oneof .certificate_config).
	// - pemCsr (*string): Mapped to .pem_csr (under oneof .certificate_config).
	// - lifetime (string): Mapped to .lifetime duration.
	// - subjectMode (string): Mapped to .subject_mode enum.

	// Spec fields:
	f.SpecField(".pem_csr")
	f.SpecField(".config")
	f.SpecField(".config.subject_config")
	f.SpecField(".config.subject_config.subject")
	f.SpecField(".config.subject_config.subject.common_name")
	f.SpecField(".config.subject_config.subject.country_code")
	f.SpecField(".config.subject_config.subject.locality")
	f.SpecField(".config.subject_config.subject.organization")
	f.SpecField(".config.subject_config.subject.organizational_unit")
	f.SpecField(".config.subject_config.subject.postal_code")
	f.SpecField(".config.subject_config.subject.province")
	f.SpecField(".config.subject_config.subject.street_address")

	f.SpecField(".config.subject_config.subject_alt_name")
	f.SpecField(".config.subject_config.subject_alt_name.dns_names")
	f.SpecField(".config.subject_config.subject_alt_name.email_addresses")
	f.SpecField(".config.subject_config.subject_alt_name.ip_addresses")
	f.SpecField(".config.subject_config.subject_alt_name.uris")

	f.SpecField(".config.x509_config")
	f.SpecField(".config.x509_config.key_usage")
	f.SpecField(".config.x509_config.key_usage.base_key_usage")
	f.SpecField(".config.x509_config.key_usage.base_key_usage.digital_signature")
	f.SpecField(".config.x509_config.key_usage.base_key_usage.content_commitment")
	f.SpecField(".config.x509_config.key_usage.base_key_usage.key_encipherment")
	f.SpecField(".config.x509_config.key_usage.base_key_usage.data_encipherment")
	f.SpecField(".config.x509_config.key_usage.base_key_usage.key_agreement")
	f.SpecField(".config.x509_config.key_usage.base_key_usage.cert_sign")
	f.SpecField(".config.x509_config.key_usage.base_key_usage.crl_sign")
	f.SpecField(".config.x509_config.key_usage.base_key_usage.encipher_only")
	f.SpecField(".config.x509_config.key_usage.base_key_usage.decipher_only")

	f.SpecField(".config.x509_config.key_usage.extended_key_usage")
	f.SpecField(".config.x509_config.key_usage.extended_key_usage.server_auth")
	f.SpecField(".config.x509_config.key_usage.extended_key_usage.client_auth")
	f.SpecField(".config.x509_config.key_usage.extended_key_usage.code_signing")
	f.SpecField(".config.x509_config.key_usage.extended_key_usage.email_protection")
	f.SpecField(".config.x509_config.key_usage.extended_key_usage.time_stamping")
	f.SpecField(".config.x509_config.key_usage.extended_key_usage.ocsp_signing")

	f.SpecField(".config.x509_config.key_usage.unknown_extended_key_usages")
	f.SpecField(".config.x509_config.key_usage.unknown_extended_key_usages.object_id_path")

	f.SpecField(".config.x509_config.ca_options")
	f.SpecField(".config.x509_config.ca_options.is_ca")
	f.SpecField(".config.x509_config.ca_options.max_issuer_path_length")

	f.SpecField(".config.x509_config.policy_ids")
	f.SpecField(".config.x509_config.policy_ids.object_id_path")

	f.SpecField(".config.x509_config.aia_ocsp_servers")

	f.SpecField(".config.x509_config.additional_extensions")
	f.SpecField(".config.x509_config.additional_extensions.critical")
	f.SpecField(".config.x509_config.additional_extensions.value")

	f.SpecField(".config.public_key")
	f.SpecField(".config.public_key.key")
	f.SpecField(".config.public_key.format")

	f.SpecField(".lifetime")
	f.SpecField(".certificate_template")
	f.SpecField(".subject_mode")

	// Documented Field Comparison (KRM Status -> Proto):
	// - conditions ([]v1alpha1.Condition): Metadata, not mapped to proto.
	// - observedGeneration (*int64): Metadata, not mapped to proto.
	// - certificateDescription (*CertificateCertificateDescriptionStatus): Mapped to .certificate_description.
	// - createTime (*string): Mapped to .create_time.
	// - issuerCertificateAuthority (*string): Mapped to .issuer_certificate_authority.
	// - pemCertificate (*string): Mapped to .pem_certificate.
	// - pemCertificateChain ([]string): Mapped to .pem_certificate_chain.
	// - revocationDetails (*CertificateRevocationDetailsStatus): Mapped to .revocation_details.
	// - updateTime (*string): Mapped to .update_time.

	// Status fields:
	f.StatusField(".certificate_description")
	f.StatusField(".certificate_description.aia_issuing_certificate_urls")
	f.StatusField(".certificate_description.authority_key_id")
	f.StatusField(".certificate_description.authority_key_id.key_id")
	f.StatusField(".certificate_description.cert_fingerprint")
	f.StatusField(".certificate_description.cert_fingerprint.sha256_hash")
	f.StatusField(".certificate_description.crl_distribution_points")
	f.StatusField(".certificate_description.public_key")
	f.StatusField(".certificate_description.public_key.format")
	f.StatusField(".certificate_description.public_key.key")

	f.StatusField(".certificate_description.subject_description")
	f.StatusField(".certificate_description.subject_description.hex_serial_number")
	f.StatusField(".certificate_description.subject_description.lifetime")
	f.StatusField(".certificate_description.subject_description.not_before_time")
	f.StatusField(".certificate_description.subject_description.not_after_time")

	f.StatusField(".certificate_description.subject_description.subject")
	f.StatusField(".certificate_description.subject_description.subject.common_name")
	f.StatusField(".certificate_description.subject_description.subject.country_code")
	f.StatusField(".certificate_description.subject_description.subject.locality")
	f.StatusField(".certificate_description.subject_description.subject.organization")
	f.StatusField(".certificate_description.subject_description.subject.organizational_unit")
	f.StatusField(".certificate_description.subject_description.subject.postal_code")
	f.StatusField(".certificate_description.subject_description.subject.province")
	f.StatusField(".certificate_description.subject_description.subject.street_address")

	f.StatusField(".certificate_description.subject_description.subject_alt_name")
	f.StatusField(".certificate_description.subject_description.subject_alt_name.dns_names")
	f.StatusField(".certificate_description.subject_description.subject_alt_name.email_addresses")
	f.StatusField(".certificate_description.subject_description.subject_alt_name.ip_addresses")
	f.StatusField(".certificate_description.subject_description.subject_alt_name.uris")
	f.StatusField(".certificate_description.subject_description.subject_alt_name.custom_sans")
	f.StatusField(".certificate_description.subject_description.subject_alt_name.custom_sans.object_id")
	f.StatusField(".certificate_description.subject_description.subject_alt_name.custom_sans.object_id.object_id_path")
	f.StatusField(".certificate_description.subject_description.subject_alt_name.custom_sans.critical")
	f.StatusField(".certificate_description.subject_description.subject_alt_name.custom_sans.value")

	f.StatusField(".certificate_description.subject_key_id")
	f.StatusField(".certificate_description.subject_key_id.key_id")

	f.StatusField(".certificate_description.x509_description")
	f.StatusField(".certificate_description.x509_description.aia_ocsp_servers")

	f.StatusField(".certificate_description.x509_description.ca_options")
	f.StatusField(".certificate_description.x509_description.ca_options.is_ca")
	f.StatusField(".certificate_description.x509_description.ca_options.max_issuer_path_length")

	f.StatusField(".certificate_description.x509_description.key_usage")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage.digital_signature")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage.content_commitment")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage.key_encipherment")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage.data_encipherment")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage.key_agreement")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage.cert_sign")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage.crl_sign")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage.encipher_only")
	f.StatusField(".certificate_description.x509_description.key_usage.base_key_usage.decipher_only")

	f.StatusField(".certificate_description.x509_description.key_usage.extended_key_usage")
	f.StatusField(".certificate_description.x509_description.key_usage.extended_key_usage.server_auth")
	f.StatusField(".certificate_description.x509_description.key_usage.extended_key_usage.client_auth")
	f.StatusField(".certificate_description.x509_description.key_usage.extended_key_usage.code_signing")
	f.StatusField(".certificate_description.x509_description.key_usage.extended_key_usage.email_protection")
	f.StatusField(".certificate_description.x509_description.key_usage.extended_key_usage.time_stamping")
	f.StatusField(".certificate_description.x509_description.key_usage.extended_key_usage.ocsp_signing")

	f.StatusField(".certificate_description.x509_description.key_usage.unknown_extended_key_usages")
	f.StatusField(".certificate_description.x509_description.key_usage.unknown_extended_key_usages.object_id_path")

	f.StatusField(".certificate_description.x509_description.policy_ids")
	f.StatusField(".certificate_description.x509_description.policy_ids.object_id_path")

	f.StatusField(".certificate_description.x509_description.additional_extensions")
	f.StatusField(".certificate_description.x509_description.additional_extensions.object_id")
	f.StatusField(".certificate_description.x509_description.additional_extensions.object_id.object_id_path")
	f.StatusField(".certificate_description.x509_description.additional_extensions.critical")
	f.StatusField(".certificate_description.x509_description.additional_extensions.value")

	f.StatusField(".create_time")
	f.StatusField(".issuer_certificate_authority")
	f.StatusField(".pem_certificate")
	f.StatusField(".pem_certificate_chain")
	f.StatusField(".revocation_details")
	f.StatusField(".revocation_details.revocation_state")
	f.StatusField(".revocation_details.revocation_time")
	f.StatusField(".update_time")

	// Identity and labels/annotations
	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".config.subject_key_id")
	f.Unimplemented_NotYetTriaged(".config.subject_config.subject.rdn_sequence")
	f.Unimplemented_NotYetTriaged(".config.subject_config.subject_alt_name.custom_sans")
	f.Unimplemented_NotYetTriaged(".config.x509_config.name_constraints")
	f.Unimplemented_NotYetTriaged(".config.x509_config.additional_extensions[].object_id")
	f.Unimplemented_NotYetTriaged(".certificate_description.subject_description.subject.rdn_sequence")
	f.Unimplemented_NotYetTriaged(".certificate_description.x509_description.name_constraints")
	f.Unimplemented_NotYetTriaged(".certificate_description.tbs_certificate_digest")

	return f
}
