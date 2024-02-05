// Copyright 2023 Google LLC
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

package runtime

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var converterMap = map[string]func() OnePlatformConverter{
	"PubSubTopic": func() OnePlatformConverter { return &PubSubConverter{} },
}

func NewConverter(implType string) OnePlatformConverter {
	if constructor, ok := converterMap[implType]; ok {
		return constructor()
	}
	panic("unknown implementation type")
}

type OnePlatformConverter interface {
	// GetResource() tries to get the corresponding GCP resource back from the
	// OnePlatform API as Unstructured.
	GetResource(ctx context.Context, obj *unstructured.Unstructured) (u *unstructured.Unstructured, err error)
	// CreateResource() tries to create the corresponding GCP resource through
	// the OnePlatform API and returns the created resource as Unstructured.
	CreateResource(ctx context.Context, obj *unstructured.Unstructured) (u *unstructured.Unstructured, err error)
	// UpdateResource() tries to update the corresponding GCP resource through
	// the OnePlatform API and returns the updated resource as Unstructured.
	UpdateResource(ctx context.Context, obj *unstructured.Unstructured) (u *unstructured.Unstructured, err error)
	// DeleteResource() tries to delete the corresponding GCP resource through
	// the OnePlatform API.
	DeleteResource(ctx context.Context, obj *unstructured.Unstructured) error
	// GetDiff() returns true if the local and remote resources are different.
	GetDiff(ctx context.Context, local, remote *unstructured.Unstructured) (bool, error)
}
