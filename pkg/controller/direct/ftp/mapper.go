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

package ftp

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ftp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	ftppb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/ftp/pb"
)

func InternalServerConfig_AllowedConsumer_ToProto(mapCtx *direct.MapContext, in *krm.InternalServerConfig_AllowedConsumer) *ftppb.InternalServerConfig_AllowedConsumer {
	if in == nil {
		return nil
	}
	out := &ftppb.InternalServerConfig_AllowedConsumer{}
	if in.ProjectRef != nil {
		out.ConsumerType = &ftppb.InternalServerConfig_AllowedConsumer_Project{Project: in.ProjectRef.External}
	}
	out.ConnectionLimit = direct.ValueOf(in.ConnectionLimit)
	return out
}

func InternalServerConfig_DeniedConsumer_ToProto(mapCtx *direct.MapContext, in *krm.InternalServerConfig_DeniedConsumer) *ftppb.InternalServerConfig_DeniedConsumer {
	if in == nil {
		return nil
	}
	out := &ftppb.InternalServerConfig_DeniedConsumer{}
	if in.ProjectRef != nil {
		out.ConsumerType = &ftppb.InternalServerConfig_DeniedConsumer_Project{Project: in.ProjectRef.External}
	}
	return out
}
