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
// proto.message: google.cloud.security.privateca.v1.CertificateAuthority

package privateca

import (
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/types/known/durationpb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(PrivateCACertificateAuthorityFuzzer())
}

func PrivateCACertificateAuthorityFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CertificateAuthority{},
		PrivateCACertificateAuthoritySpec_FromProto, PrivateCACertificateAuthoritySpec_ToProto,
		PrivateCACertificateAuthorityStatus_FromProto, PrivateCACertificateAuthorityStatus_ToProto,
	)

	f.FilterSpec = func(in *pb.CertificateAuthority) {
		if in.Config == nil {
			in.Config = &pb.CertificateConfig{
				SubjectConfig: &pb.CertificateConfig_SubjectConfig{
					Subject: &pb.Subject{},
				},
				X509Config: &pb.X509Parameters{},
			}
		}
		if in.KeySpec == nil {
			in.KeySpec = &pb.CertificateAuthority_KeyVersionSpec{}
		}
		if in.Lifetime == nil {
			in.Lifetime = &durationpb.Duration{Seconds: 3600}
		}
		if keySpec := in.GetKeySpec(); keySpec != nil {
			if alg, ok := keySpec.GetKeyVersion().(*pb.CertificateAuthority_KeyVersionSpec_Algorithm); ok {
				if alg.Algorithm == pb.CertificateAuthority_SIGN_HASH_ALGORITHM_UNSPECIFIED {
					alg.Algorithm = pb.CertificateAuthority_RSA_PSS_2048_SHA256
				}
			}
		}
	}

	f.Unimplemented_Identity(".name")

	// Documented Field Comparison (KRM Spec -> Proto):
	// - projectRef (refs.ProjectRef): Mapped to parent project context, not directly in CertificateAuthority proto.
	// - location (string): Mapped to parent location context, not directly in CertificateAuthority proto.
	// - resourceID (*string): Mapped to the last segment of the .name field in the CertificateAuthority proto.
	// - config (CertificateConfig): Mapped to .config.
	// - keySpec (CertificateAuthority_KeyVersionSpec): Mapped to .key_spec.
	// - lifetime (string): Mapped to .lifetime.
	// - type (string): Mapped to .type.
	// - gcsBucketRef (*storagev1beta1.StorageBucketRef): Mapped to .gcs_bucket string.
	// - caPoolRef (privatecarefs.PrivateCACAPoolRef): Mapped to parent CA Pool context, not directly in CertificateAuthority proto.
	f.SpecField(".config")
	f.SpecField(".key_spec")
	f.SpecField(".lifetime")
	f.SpecField(".type")
	f.SpecField(".gcs_bucket")

	// Documented Field Comparison (KRM Status -> Proto):
	// - conditions ([]v1alpha1.Condition): Metadata, not mapped to proto.
	// - observedGeneration (*int64): Metadata, not mapped to proto.
	// - accessUrls (*CertificateAuthority_AccessUrls): Mapped to .access_urls.
	// - caCertificateDescriptions ([]CertificateDescription): Mapped to .ca_certificate_descriptions.
	// - config (*CertificateAuthority_ConfigStatus): Mapped to .config.
	// - createTime (*string): Mapped to .create_time.
	// - deleteTime (*string): Mapped to .delete_time.
	// - expireTime (*string): Mapped to .expire_time.
	// - pemCaCertificates ([]string): Mapped to .pem_ca_certificates.
	// - state (*string): Mapped to .state.
	// - subordinateConfig (*SubordinateConfig): Mapped to .subordinate_config.
	// - tier (*string): Mapped to .tier.
	// - updateTime (*string): Mapped to .update_time.
	f.StatusField(".access_urls")
	f.StatusField(".ca_certificate_descriptions")
	f.StatusField(".config")
	f.StatusField(".pem_ca_certificates")
	f.StatusField(".state")
	f.StatusField(".subordinate_config")
	f.StatusField(".tier")
	f.StatusField(".create_time")
	f.StatusField(".delete_time")
	f.StatusField(".expire_time")
	f.StatusField(".update_time")

	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".user_defined_access_urls")

	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].subject_description.subject.rdn_sequence")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].subject_key_id")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].authority_key_id")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].tbs_certificate_digest")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].x509_description.name_constraints")

	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].subject_description.lifetime.nanos")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].subject_description.not_after_time")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].subject_description.not_after_time.seconds")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].subject_description.not_after_time.nanos")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].subject_description.not_before_time")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].subject_description.not_before_time.seconds")
	f.Unimplemented_NotYetTriaged(".ca_certificate_descriptions[].subject_description.not_before_time.nanos")

	return f
}
