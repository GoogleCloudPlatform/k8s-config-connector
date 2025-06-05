// Copyright 2025 Google LLC
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

package storagecontrol

import (
	"context"

	control "cloud.google.com/go/storage/control/apiv2"
	controlpb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/googleapis/gax-go/v2"
)

type AnywhereCacheAPI interface {
	CreateAnywhereCache(ctx context.Context, in *controlpb.CreateAnywhereCacheRequest, opts ...gax.CallOption) (*control.CreateAnywhereCacheOperation, error)
	UpdateAnywhereCache(ctx context.Context, in *controlpb.UpdateAnywhereCacheRequest, opts ...gax.CallOption) (*control.UpdateAnywhereCacheOperation, error)
	DisableAnywhereCache(ctx context.Context, in *controlpb.DisableAnywhereCacheRequest, opts ...gax.CallOption) (*controlpb.AnywhereCache, error)
	PauseAnywhereCache(ctx context.Context, in *controlpb.PauseAnywhereCacheRequest, opts ...gax.CallOption) (*controlpb.AnywhereCache, error)
	ResumeAnywhereCache(ctx context.Context, in *controlpb.ResumeAnywhereCacheRequest, opts ...gax.CallOption) (*controlpb.AnywhereCache, error)
	GetAnywhereCache(ctx context.Context, in *controlpb.GetAnywhereCacheRequest, opts ...gax.CallOption) (*controlpb.AnywhereCache, error)
	ListAnywhereCaches(ctx context.Context, in *controlpb.ListAnywhereCachesRequest, opts ...gax.CallOption) *control.AnywhereCacheIterator
}

type DirectBaseUpdateOperation interface {
	UpdateStatus(ctx context.Context, typedStatus any, readyCondition *v1alpha1.Condition) error
	RequestRequeue()
}
