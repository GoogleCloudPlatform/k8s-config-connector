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
// proto.message: google.cloud.beyondcorp.appconnections.v1.AppConnection
// api.group: beyondcorp.cnrm.cloud.google.com

package beyondcorp

import (
	pb "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(beyondCorpAppConnectionFuzzer())
}

func beyondCorpAppConnectionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.AppConnection{},
		BeyondCorpAppConnectionSpec_FromProto, BeyondCorpAppConnectionSpec_ToProto,
	)

	// Field comparison comments:
	// KRM BeyondCorpAppConnectionSpec fields:
	// - ApplicationEndpoint.Host -> maps to proto .application_endpoint.host
	// - ApplicationEndpoint.Port -> maps to proto .application_endpoint.port
	// - Connectors               -> maps to proto .connectors
	// - DisplayName              -> maps to proto .display_name
	// - Gateway.AppGateway       -> maps to proto .gateway.app_gateway
	// - Gateway.IngressPort      -> maps to proto .gateway.ingress_port
	// - Gateway.Type             -> maps to proto .gateway.type
	// - Gateway.URI              -> maps to proto .gateway.uri
	// - ProjectRef               -> KRM hierarchical reference, not in GCP proto
	// - Region                   -> GCP location/region in URL path, not in GCP proto
	// - ResourceID               -> KRM resource name/ID in URL path, not in GCP proto
	// - Type                     -> maps to proto .type

	// Spec fields
	f.SpecField(".application_endpoint")
	f.SpecField(".application_endpoint.host")
	f.SpecField(".application_endpoint.port")
	f.SpecField(".connectors")
	f.SpecField(".display_name")
	f.SpecField(".gateway")
	f.SpecField(".gateway.app_gateway")
	f.SpecField(".gateway.ingress_port")
	f.SpecField(".gateway.type")
	f.SpecField(".gateway.uri")
	f.SpecField(".type")

	// Identity fields
	f.Unimplemented_Identity(".name")

	// Labels / Annotations
	f.Unimplemented_LabelsAnnotations(".labels")

	// Server-managed / untriaged fields
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".uid")
	f.Unimplemented_NotYetTriaged(".state")

	return f
}
