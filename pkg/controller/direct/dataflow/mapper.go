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

package dataflow

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataflow/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/anypb"
)

func Any_FromProto(mapCtx *direct.MapContext, in *anypb.Any) *krm.Any {
	if in == nil {
		return nil
	}
	return &krm.Any{
		TypeURL: direct.LazyPtr(in.TypeUrl),
		Value:   in.Value,
	}
}

func Any_ToProto(mapCtx *direct.MapContext, in *krm.Any) *anypb.Any {
	if in == nil {
		return nil
	}
	return &anypb.Any{
		TypeUrl: direct.ValueOf(in.TypeURL),
		Value:   in.Value,
	}
}
