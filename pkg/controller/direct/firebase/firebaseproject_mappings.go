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

package firebase

import (
	"strconv"

	pb "google.golang.org/api/firebase/v1beta1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firebase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FirebaseProjectSpec_FromAPI(mapCtx *direct.MapContext, in *pb.FirebaseProject) *krm.FirebaseProjectSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseProjectSpec{}
	return out
}

func FirebaseProjectSpec_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseProjectSpec) *pb.FirebaseProject {
	if in == nil {
		return nil
	}
	out := &pb.FirebaseProject{}
	return out
}

func FirebaseProjectStatus_FromAPI(mapCtx *direct.MapContext, in *pb.FirebaseProject) *krm.FirebaseProjectStatus {
	if in == nil {
		return nil
	}
	out := &krm.FirebaseProjectStatus{}
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	if in.ProjectNumber != 0 {
		out.ProjectNumber = direct.LazyPtr(strconv.FormatInt(in.ProjectNumber, 10))
	}
	return out
}

func FirebaseProjectStatus_ToAPI(mapCtx *direct.MapContext, in *krm.FirebaseProjectStatus) *pb.FirebaseProject {
	if in == nil {
		return nil
	}
	out := &pb.FirebaseProject{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if in.ProjectNumber != nil {
		n, err := strconv.ParseInt(*in.ProjectNumber, 10, 64)
		if err != nil {
			mapCtx.Errorf("cannot parse projectNumber %q: %w", *in.ProjectNumber, err)
		} else {
			out.ProjectNumber = n
		}
	}
	return out
}
