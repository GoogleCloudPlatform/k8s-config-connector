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

package dialogflowcx

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krmdialogflowcxv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DialogflowCXWebhookSpec_DisplayName_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

func ServiceDirectoryConfig_Service_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

// GenericWebService_v1alpha1_FromProto maps pb.Webhook_GenericWebService to KRM GenericWebService.
// Handcoded due to field name differences (Uri vs URI) and type conversion for allowed CACerts ([]byte vs base64/string).
func GenericWebService_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Webhook_GenericWebService) *krmdialogflowcxv1alpha1.GenericWebService {
	if in == nil {
		return nil
	}
	out := &krmdialogflowcxv1alpha1.GenericWebService{}
	out.URI = in.GetUri()
	out.RequestHeaders = in.GetRequestHeaders()
	if in.GetAllowedCaCerts() != nil {
		out.AllowedCACerts = make([]string, len(in.GetAllowedCaCerts()))
		for i, cert := range in.GetAllowedCaCerts() {
			out.AllowedCACerts[i] = string(cert)
		}
	}
	return out
}

// GenericWebService_v1alpha1_ToProto maps KRM GenericWebService to pb.Webhook_GenericWebService.
// Handcoded due to field name differences (Uri vs URI) and type conversion for allowed CACerts (string vs []byte).
func GenericWebService_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdialogflowcxv1alpha1.GenericWebService) *pb.Webhook_GenericWebService {
	if in == nil {
		return nil
	}
	out := &pb.Webhook_GenericWebService{}
	out.Uri = in.URI
	out.RequestHeaders = in.RequestHeaders
	if in.AllowedCACerts != nil {
		out.AllowedCaCerts = make([][]byte, len(in.AllowedCACerts))
		for i, cert := range in.AllowedCACerts {
			out.AllowedCaCerts[i] = []byte(cert)
		}
	}
	return out
}
