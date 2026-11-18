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

package geminidataanalytics

import (
	geminidataanalyticspb "cloud.google.com/go/geminidataanalytics/apiv1/geminidataanalyticspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/geminidataanalytics/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GeminiDataAnalyticsDataAgentSpec_ToProto(mapCtx *direct.MapContext, in *krm.GeminiDataAnalyticsDataAgentSpec) *geminidataanalyticspb.DataAgent {
	if in == nil {
		return nil
	}
	out := &geminidataanalyticspb.DataAgent{}
	if in.DataAnalyticsAgent != nil {
		out.Type = &geminidataanalyticspb.DataAgent_DataAnalyticsAgent{DataAnalyticsAgent: DataAnalyticsAgent_ToProto(mapCtx, in.DataAnalyticsAgent)}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	if in.KMSKeyRef != nil && in.KMSKeyRef.External != "" {
		out.KmsKey = direct.LazyPtr(in.KMSKeyRef.External)
	}
	out.BigqueryAgentAnalyticsEnabled = in.BigqueryAgentAnalyticsEnabled
	out.BigqueryAgentAnalyticsTable = BigQueryTableReference_ToProto(mapCtx, in.BigqueryAgentAnalyticsTable)
	return out
}
