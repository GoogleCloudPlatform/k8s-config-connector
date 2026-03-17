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

package binaryauthorization

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AdmissionRule_RequireAttestationsBy_FromProto(mapCtx *direct.MapContext, in []string) []refsv1beta1.BinaryAuthorizationAttestorRef {
	if in == nil {
		return nil
	}
	out := make([]refsv1beta1.BinaryAuthorizationAttestorRef, len(in))
	for i, s := range in {
		out[i] = refsv1beta1.BinaryAuthorizationAttestorRef{
			External: s,
		}
	}
	return out
}

func AdmissionRule_RequireAttestationsBy_ToProto(mapCtx *direct.MapContext, in []refsv1beta1.BinaryAuthorizationAttestorRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, ref := range in {
		if ref.External != "" {
			out[i] = ref.External
		} else {
			// Typically handled by a generic resolver or just returning Name
			out[i] = ref.Name
		}
	}
	return out
}
