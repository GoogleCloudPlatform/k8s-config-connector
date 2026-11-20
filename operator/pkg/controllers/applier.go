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

package controllers

import (
	"context"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"
)

// NewApplier wraps the default applier so that each Apply works on its own copy
// of the rest.Config. The direct applier hands the config to cli-runtime, which
// rewrites GroupVersion and APIPath on it while building per-group REST clients,
// so sharing the manager config between concurrent reconciles can send requests
// to the wrong API group.
func NewApplier() applier.Applier {
	return &copyConfigApplier{inner: applier.DefaultApplier}
}

type copyConfigApplier struct {
	inner applier.Applier
}

func (a *copyConfigApplier) Apply(ctx context.Context, opt applier.ApplierOptions) error {
	if opt.RESTConfig != nil {
		opt.RESTConfig = rest.CopyConfig(opt.RESTConfig)
	}
	return a.inner.Apply(ctx, opt)
}
