// Copyright 2022 Google LLC
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

package configconnectorcontext

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

// SourceLabel returns a fixed label based on the namespace of the ConfigConnectorContext object.
func SourceLabel() declarative.LabelMaker {
	return func(ctx context.Context, o declarative.DeclarativeObject) map[string]string {
		res := map[string]string{
			k8s.ConfigConnectorContextNamespaceLabel: o.GetNamespace(),
		}
		return res
	}
}
