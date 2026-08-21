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

package mappers

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// OnlySpecFields filters out non-spec fields from an actual API object by round-tripping
// it through its KRM representation.
func OnlySpecFields[API any, KRM any](
	actual *API,
	fromAPI func(*direct.MapContext, *API) *KRM,
	toAPI func(*direct.MapContext, *KRM) *API,
) (*API, error) {
	mapCtx := &direct.MapContext{}
	spec := fromAPI(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	maskedActual := toAPI(mapCtx, spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	return maskedActual, nil
}
