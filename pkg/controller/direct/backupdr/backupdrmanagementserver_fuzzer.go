// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.backupdr.v1.ManagementServer
// api.group: backupdr.cnrm.cloud.google.com

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(backupDRManagementServerFuzzer())
}

func backupDRManagementServerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ManagementServer{},
		BackupDRManagementServerSpec_v1alpha1_FromProto, BackupDRManagementServerSpec_v1alpha1_ToProto,
		BackupDRManagementServerObservedState_v1alpha1_FromProto, BackupDRManagementServerObservedState_v1alpha1_ToProto,
	)

	// Detailed field comparison:
	// - Spec.Description -> .description
	// - Spec.Labels -> .labels (handled via Unimplemented_LabelsAnnotations)
	// - Spec.Type -> .type
	// - Spec.Networks -> .networks (NetworkConfig contains NetworkRef and PeeringMode)
	f.SpecField(".description")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.SpecField(".type")
	f.SpecField(".networks")

	// Detailed status field comparison:
	// - Status.ObservedState.CreateTime -> .create_time
	// - Status.ObservedState.UpdateTime -> .update_time
	// - Status.ObservedState.ManagementURI -> .management_uri (ManagementURI contains WebUI and API)
	// - Status.ObservedState.WorkforceIdentityBasedManagementURI -> .workforce_identity_based_management_uri (WorkforceIdentityBasedManagementURI contains FirstPartyManagementURI and ThirdPartyManagementURI)
	// - Status.ObservedState.State -> .state
	// - Status.ObservedState.OAuth2ClientID -> .oauth2_client_id
	// - Status.ObservedState.WorkforceIdentityBasedOAuth2ClientID -> .workforce_identity_based_oauth2_client_id (WorkforceIdentityBasedOAuth2ClientID contains FirstPartyOAuth2ClientID and ThirdPartyOAuth2ClientID)
	// - Status.ObservedState.BAProxyURIs -> .ba_proxy_uri
	f.StatusField(".management_uri")
	f.StatusField(".workforce_identity_based_management_uri")
	f.StatusField(".state")
	f.StatusField(".oauth2_client_id")
	f.StatusField(".workforce_identity_based_oauth2_client_id")
	f.StatusField(".ba_proxy_uri")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	// Special and Unimplemented fields:
	// - .name -> Identity field (maps to externalRef/metadata.name)
	// - .etag -> Server-specified ETag
	// - .satisfies_pzs -> Reserved for future use (not yet triaged)
	// - .satisfies_pzi -> Reserved for future use (not yet triaged)
	f.Unimplemented_Identity(".name") // special field
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_Etag()

	return f
}
