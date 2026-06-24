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

package networkservices

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(edgeCacheServiceFuzzer())
}

type EdgeCacheServiceAPI struct {
	Name   string
	Spec   krm.NetworkServicesEdgeCacheServiceSpec
	Status krm.NetworkServicesEdgeCacheServiceStatus
}

func edgeCacheServiceFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&EdgeCacheServiceAPI{},
		EdgeCacheServiceSpec_FromAPI, EdgeCacheServiceSpec_ToAPI,
		EdgeCacheServiceStatus_FromAPI, EdgeCacheServiceStatus_ToAPI,
	)

	f.SpecField(".Spec")
	f.StatusField(".Status")
	f.IdentityField(".Name")

	return f
}

func EdgeCacheServiceSpec_FromAPI(ctx *direct.MapContext, in *EdgeCacheServiceAPI) *krm.NetworkServicesEdgeCacheServiceSpec {
	if in == nil {
		return nil
	}
	out := in.Spec
	return &out
}

func EdgeCacheServiceSpec_ToAPI(ctx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceSpec) *EdgeCacheServiceAPI {
	if in == nil {
		return nil
	}
	return &EdgeCacheServiceAPI{
		Spec: *in,
	}
}

func EdgeCacheServiceStatus_FromAPI(ctx *direct.MapContext, in *EdgeCacheServiceAPI) *krm.NetworkServicesEdgeCacheServiceStatus {
	if in == nil {
		return nil
	}
	out := in.Status
	return &out
}

func EdgeCacheServiceStatus_ToAPI(ctx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceStatus) *EdgeCacheServiceAPI {
	if in == nil {
		return nil
	}
	return &EdgeCacheServiceAPI{
		Status: *in,
	}
}
