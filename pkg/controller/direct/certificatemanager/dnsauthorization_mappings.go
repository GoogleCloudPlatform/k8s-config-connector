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
	certificatemanagerpb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CertificateManagerDNSAuthorizationStatusObservedState_FromProto(mapCtx *direct.MapContext, created *certificatemanagerpb.DnsAuthorization) *krm.CertificateManagerDNSAuthorizationStatus {
	status := &krm.CertificateManagerDNSAuthorizationStatus{
		DnsResourceRecord: []krm.DnsAuthorization_DnsResourceRecord{direct.ValueOf(DnsAuthorization_DnsResourceRecord_FromProto(mapCtx, created.GetDnsResourceRecord()))},
	}
	return status
}
