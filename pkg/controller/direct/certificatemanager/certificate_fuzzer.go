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
// proto.message: google.cloud.certificatemanager.v1.Certificate
// api.group: certificatemanager.cnrm.cloud.google.com

package certificatemanager

import (
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(certificateFuzzer())
}

func certificateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Certificate{},
		CertificateManagerCertificateSpec_FromProto, CertificateManagerCertificateSpec_ToProto,
		CertificateObservedStateStatus_FromProto, CertificateObservedStateStatus_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".scope")
	f.SpecField(".self_managed")

	f.SpecField(".self_managed.pem_certificate")
	f.SpecField(".self_managed.pem_private_key")

	f.SpecField(".managed.domains")
	f.SpecField(".managed.dns_authorizations")
	f.SpecField(".managed.issuance_config")

	f.StatusField(".managed.state")
	f.StatusField(".managed.provisioning_issue")
	f.StatusField(".managed.authorization_attempt_info")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_Internal(".create_time")
	f.Unimplemented_Internal(".update_time")
	f.UnimplementedFields.Insert(".san_dnsnames")
	f.UnimplementedFields.Insert(".pem_certificate")
	f.UnimplementedFields.Insert(".expire_time")
	f.UnimplementedFields.Insert(".type")

	f.FilterSpec = func(in *pb.Certificate) {
		if in.GetSelfManaged() != nil {
			// Clear private key since it's write-only and not round-tripped (not returned by GET)
			in.GetSelfManaged().PemPrivateKey = ""
		}
	}

	return f
}
