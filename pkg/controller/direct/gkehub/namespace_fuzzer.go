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

package gkehub

import (
	gkehubv1 "google.golang.org/api/gkehub/v1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(fuzzGKEHubNamespace())
}

func fuzzGKEHubNamespace() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&gkehubv1.Namespace{},
		func(ctx *direct.MapContext, in *gkehubv1.Namespace) *krm.GKEHubNamespaceSpec {
			return GKEHubNamespaceSpec_FromAPI(ctx, in, &krm.GKEHubNamespaceIdentity{})
		},
		GKEHubNamespaceSpec_ToAPI,
		func(ctx *direct.MapContext, in *gkehubv1.Namespace) *krm.GKEHubNamespaceStatus {
			return GKEHubNamespaceStatus_FromAPI(ctx, in)
		},
		GKEHubNamespaceStatus_ToAPI,
	)

	f.SpecField(".Labels")
	f.SpecField(".NamespaceLabels")

	f.StatusField(".CreateTime")
	f.StatusField(".UpdateTime")
	f.StatusField(".DeleteTime")
	f.StatusField(".Uid")
	f.StatusField(".State")

	f.Unimplemented_NotYetTriaged(".Name")
	f.Unimplemented_NotYetTriaged(".Scope")
	f.Unimplemented_NotYetTriaged(".ForceSendFields")
	f.Unimplemented_NotYetTriaged(".NullFields")
	f.Unimplemented_NotYetTriaged(".ServerResponse")

	f.Unimplemented_NotYetTriaged(".State.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".State.NullFields")

	return f
}
