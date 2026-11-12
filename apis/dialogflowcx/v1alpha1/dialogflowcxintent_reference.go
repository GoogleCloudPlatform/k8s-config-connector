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

package v1alpha1

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DialogflowCXIntentRef{}

func (r *DialogflowCXIntentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		ready := false
		conditions, found, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
		if found {
			for _, condVal := range conditions {
				cond, ok := condVal.(map[string]interface{})
				if !ok {
					continue
				}
				if cond["type"] == "Ready" && cond["status"] == "True" {
					ready = true
					break
				}
			}
		}
		if !ready {
			return ""
		}
		name, _, _ := unstructured.NestedString(u.Object, "status", "name")
		if name != "" {
			return name
		}
		obj, err := common.ToStructuredType[*DialogflowCXIntent](u)
		if err != nil {
			return ""
		}
		id, err := getIdentityFromDialogflowCXIntentSpec(ctx, reader, obj)
		if err != nil || id.Intent == "" {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
