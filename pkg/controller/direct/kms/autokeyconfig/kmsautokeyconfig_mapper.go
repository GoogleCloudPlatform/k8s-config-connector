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

package kmsautokeyconfig

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func KMSAutokeyConfig_FromFields(mapCtx *direct.MapContext, id *krm.KMSAutokeyConfigIdentity, keyProject *refs.ProjectIdentity, keyProjectResolutionMode *string) *pb.AutokeyConfig {
	out := &pb.AutokeyConfig{}
	out.Name = id.String()
	if keyProject != nil {
		out.KeyProject = "projects/" + keyProject.ProjectID // keyProject expects project of the form `projects/<projectId>` or `projects/<projectNumber>`
	}
	if keyProjectResolutionMode != nil {
		out.KeyProjectResolutionMode = direct.Enum_ToProto[pb.AutokeyConfig_KeyProjectResolutionMode](mapCtx, keyProjectResolutionMode)
	}
	return out
}
