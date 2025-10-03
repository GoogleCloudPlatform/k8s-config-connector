// Copyright 2024 Google LLC
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

package certificatemanager

import (
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(dnsAuthorizationFuzzer())
}

func dnsAuthorizationFuzzer() fuzztesting.KRMFuzzer {
	fuzzer := fuzztesting.NewKRMTypedSpecFuzzer(
		&pb.DnsAuthorization{},
		CertificateManagerDNSAuthorizationSpec_FromProto,
		CertificateManagerDNSAuthorizationSpec_ToProto,
	)

	fuzzer.UnimplementedFields.Insert(".dns_resource_record")
	fuzzer.UnimplementedFields.Insert(".name")
	fuzzer.UnimplementedFields.Insert(".labels")
	fuzzer.UnimplementedFields.Insert(".type")

	fuzzer.StatusFields.Insert(".create_time")
	fuzzer.StatusFields.Insert(".update_time")

	return fuzzer
}
