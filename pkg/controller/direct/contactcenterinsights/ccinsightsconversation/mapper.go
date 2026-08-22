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

package ccinsightsconversation

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/contactcenterinsights"
)

func ConversationSpec_ToProto(mapCtx *direct.MapContext, in *krm.CCInsightsConversationSpec) *pb.Conversation {
	return parent.CCInsightsConversationSpec_ToProto(mapCtx, in)
}

func ConversationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.CCInsightsConversationSpec {
	return parent.CCInsightsConversationSpec_FromProto(mapCtx, in)
}

func ConversationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Conversation) *krm.CCInsightsConversationObservedState {
	return parent.CCInsightsConversationObservedState_FromProto(mapCtx, in)
}
