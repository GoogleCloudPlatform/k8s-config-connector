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

package common

import (
	"context"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type NormalizeOption func(*refNormalizer)

// SkipPath excludes specific field paths from being normalized
func SkipPath(paths ...string) NormalizeOption {
	return func(r *refNormalizer) {
		if r.skipPaths == nil {
			r.skipPaths = make(map[string]bool)
		}
		for _, p := range paths {
			r.skipPaths[p] = true
		}
	}
}

func NormalizeReferences(ctx context.Context, reader client.Reader, obj client.Object, projectRef *refs.ProjectIdentity, opts ...NormalizeOption) error {
	normalizer := &refNormalizer{
		ctx:     ctx,
		kube:    reader,
		src:     obj,
		project: projectRef,
	}
	for _, opt := range opts {
		opt(normalizer)
	}

	if err := VisitFields(obj, normalizer); err != nil {
		return err
	}
	return nil
}

func normalizeProjectRef(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ProjectRef) (*refs.ProjectRef, error) {
	if ref == nil {
		return nil, nil
	}

	project, err := refs.ResolveProject(ctx, reader, src.GetNamespace(), ref)
	if err != nil {
		return nil, err
	}

	return &refs.ProjectRef{
		External: "projects/" + project.ProjectID,
	}, nil
}

type refNormalizer struct {
	ctx       context.Context
	kube      client.Reader
	src       client.Object
	project   *refs.ProjectIdentity
	skipPaths map[string]bool
}

func (r *refNormalizer) VisitField(path string, v any) error {
	if r.skipPaths != nil {
		if r.skipPaths[path] || r.skipPaths[strings.TrimPrefix(path, ".")] {
			return nil
		}
	}

	if projectRef, ok := v.(*refs.ProjectRef); ok {
		if ref, err := normalizeProjectRef(r.ctx, r.kube, r.src, projectRef); err != nil {
			return err
		} else if ref != nil {
			*projectRef = *ref
		}
	}

	if ref, ok := v.(refs.Ref); ok {
		if err := ref.Normalize(r.ctx, r.kube, r.src.GetNamespace()); err != nil {
			return err
		}
	}

	return nil
}
