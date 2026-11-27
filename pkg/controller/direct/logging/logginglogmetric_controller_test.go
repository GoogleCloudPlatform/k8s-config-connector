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

package logging

import (
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
)

func TestFullyQualifiedName(t *testing.T) {
	grid := []struct {
		name    string
		project string
		metric  string
		want    string
	}{
		{
			name:    "plain id is unchanged",
			project: "my-project",
			metric:  "my_metric",
			want:    "projects/my-project/metrics/my_metric",
		},
		{
			// Cloud Logging accepts "/" in a metric ID and Google's own
			// documentation suggests namespaced IDs. The generated client
			// expands the name with the reserved template "v2/{+metricName}",
			// which leaves "/" alone, so an unescaped ID would add a path
			// segment and the API would answer 404.
			name:    "slash in id is escaped so it stays one path segment",
			project: "my-project",
			metric:  "myapp/my_metric",
			want:    "projects/my-project/metrics/myapp%2Fmy_metric",
		},
		{
			name:    "multiple slashes are all escaped",
			project: "my-project",
			metric:  "a/b/c",
			want:    "projects/my-project/metrics/a%2Fb%2Fc",
		},
		{
			// Characters that are legal in a path segment must not be
			// mangled; over-escaping would break existing metrics.
			name:    "dots, dashes and underscores are left alone",
			project: "my-project",
			metric:  "a.b-c_d",
			want:    "projects/my-project/metrics/a.b-c_d",
		},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			a := &logMetricAdapter{
				id: &krm.LoggingLogMetricIdentity{
					Project: g.project,
					Metric:  g.metric,
				},
			}
			got := a.fullyQualifiedName()
			if got != g.want {
				t.Errorf("fullyQualifiedName() = %q, want %q", got, g.want)
			}
		})
	}
}
