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

package compute

import (
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
)

type MapContext struct {
	errs []error
}

func (c *MapContext) Errorf(msg string, args ...interface{}) {
	c.errs = append(c.errs, fmt.Errorf(msg, args...))
}

func (c *MapContext) Err() error {
	return errors.Join(c.errs...)
}

func Slice_ToProto[T, U any](mapCtx *MapContext, in []T, mapper func(mapCtx *MapContext, in *T) *U) []*U {
	if in == nil {
		return nil
	}

	outSlice := make([]*U, 0, len(in))
	for _, inItem := range in {
		outItem := mapper(mapCtx, &inItem)
		outSlice = append(outSlice, outItem)
	}
	return outSlice
}

func Slice_FromProto[T, U any](mapCtx *MapContext, in []*T, mapper func(mapCtx *MapContext, in *T) *U) []U {
	if in == nil {
		return nil
	}

	outSlice := make([]U, 0, len(in))
	for _, inItem := range in {
		outItem := mapper(mapCtx, inItem)
		outSlice = append(outSlice, *outItem)
	}
	return outSlice
}

func ResourceRef_FromProto(mapCtx *MapContext, in string) *v1alpha1.ResourceRef {
	if in == "" {
		return nil
	}
	return &v1alpha1.ResourceRef{
		External: in,
	}
}

func ResourceRef_ToProto(mapCtx *MapContext, in *v1alpha1.ResourceRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference was not pre-resolved")
	}
	return LazyPtr(in.External)
}

func LazyPtr[V comparable](v V) *V {
	var defaultV V
	if v == defaultV {
		return nil
	}
	return &v
}
