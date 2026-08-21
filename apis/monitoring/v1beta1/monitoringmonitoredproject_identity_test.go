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

package v1beta1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestMonitoringMonitoredProjectIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *MonitoringMonitoredProjectIdentity
	}{
		{
			name: "valid reference",
			ref:  "locations/global/metricsScopes/my-scoping-project/projects/my-monitored-project",
			want: &MonitoringMonitoredProjectIdentity{
				MetricsScope: "my-scoping-project",
				Project:      "my-monitored-project",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://monitoring.googleapis.com/locations/global/metricsScopes/my-scoping-project/projects/my-monitored-project",
			want: &MonitoringMonitoredProjectIdentity{
				MetricsScope: "my-scoping-project",
				Project:      "my-monitored-project",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &MonitoringMonitoredProjectIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestGetIdentityFromMonitoringMonitoredProjectSpec(t *testing.T) {
	tests := []struct {
		name    string
		obj     *MonitoringMonitoredProject
		want    *MonitoringMonitoredProjectIdentity
		wantErr bool
	}{
		{
			name: "valid metricsScope with locations",
			obj: &MonitoringMonitoredProject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-monitored-project",
				},
				Spec: MonitoringMonitoredProjectSpec{
					MetricsScope: "locations/global/metricsScopes/my-scoping-project",
				},
			},
			want: &MonitoringMonitoredProjectIdentity{
				MetricsScope: "my-scoping-project",
				Project:      "my-monitored-project",
			},
			wantErr: false,
		},
		{
			name: "valid metricsScope with singular location",
			obj: &MonitoringMonitoredProject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-monitored-project",
				},
				Spec: MonitoringMonitoredProjectSpec{
					MetricsScope: "location/global/metricsScopes/my-scoping-project",
				},
			},
			want: &MonitoringMonitoredProjectIdentity{
				MetricsScope: "my-scoping-project",
				Project:      "my-monitored-project",
			},
			wantErr: false,
		},
		{
			name: "valid metricsScope with simple project ID",
			obj: &MonitoringMonitoredProject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-monitored-project",
				},
				Spec: MonitoringMonitoredProjectSpec{
					MetricsScope: "my-scoping-project",
				},
			},
			want: &MonitoringMonitoredProjectIdentity{
				MetricsScope: "my-scoping-project",
				Project:      "my-monitored-project",
			},
			wantErr: false,
		},
		{
			name: "invalid metricsScope - empty",
			obj: &MonitoringMonitoredProject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-monitored-project",
				},
				Spec: MonitoringMonitoredProjectSpec{
					MetricsScope: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid metricsScope - wrong first segment",
			obj: &MonitoringMonitoredProject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-monitored-project",
				},
				Spec: MonitoringMonitoredProjectSpec{
					MetricsScope: "wrong/global/metricsScopes/my-scoping-project",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid metricsScope - wrong second segment",
			obj: &MonitoringMonitoredProject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-monitored-project",
				},
				Spec: MonitoringMonitoredProjectSpec{
					MetricsScope: "locations/wrong/metricsScopes/my-scoping-project",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid metricsScope - wrong third segment",
			obj: &MonitoringMonitoredProject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-monitored-project",
				},
				Spec: MonitoringMonitoredProjectSpec{
					MetricsScope: "locations/global/wrong/my-scoping-project",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid metricsScope - empty ID segment",
			obj: &MonitoringMonitoredProject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-monitored-project",
				},
				Spec: MonitoringMonitoredProjectSpec{
					MetricsScope: "locations/global/metricsScopes/",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid metricsScope - wrong number of segments",
			obj: &MonitoringMonitoredProject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-monitored-project",
				},
				Spec: MonitoringMonitoredProjectSpec{
					MetricsScope: "locations/global/metricsScopes/my-scoping-project/extra",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdentityFromMonitoringMonitoredProjectSpec(context.Background(), nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdentityFromMonitoringMonitoredProjectSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("getIdentityFromMonitoringMonitoredProjectSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
