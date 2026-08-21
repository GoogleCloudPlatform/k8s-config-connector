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

package monitoring

import (
	"context"
	"testing"

	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestNormalizeDashboardProto(t *testing.T) {
	ctx := context.Background()
	// Since we are not querying resource manager APIs, we don't need a real API client.
	projectMapper := projects.NewProjectMapper(projects.NewProjectCache(nil, 0))

	tests := []struct {
		name             string
		dashboardProject string
		input            *pb.Dashboard
		want             *pb.Dashboard
		wantErr          bool
	}{
		{
			name:             "normalize AlertChart short relative to canonical",
			dashboardProject: "my-project",
			input: &pb.Dashboard{
				Layout: &pb.Dashboard_GridLayout{
					GridLayout: &pb.GridLayout{
						Widgets: []*pb.Widget{
							{
								Content: &pb.Widget_AlertChart{
									AlertChart: &pb.AlertChart{
										Name: "alertPolicies/my-policy",
									},
								},
							},
						},
					},
				},
			},
			want: &pb.Dashboard{
				Layout: &pb.Dashboard_GridLayout{
					GridLayout: &pb.GridLayout{
						Widgets: []*pb.Widget{
							{
								Content: &pb.Widget_AlertChart{
									AlertChart: &pb.AlertChart{
										Name: "projects/my-project/alertPolicies/my-policy",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:             "normalize IncidentList full to relative short",
			dashboardProject: "my-project",
			input: &pb.Dashboard{
				Layout: &pb.Dashboard_GridLayout{
					GridLayout: &pb.GridLayout{
						Widgets: []*pb.Widget{
							{
								Content: &pb.Widget_IncidentList{
									IncidentList: &pb.IncidentList{
										PolicyNames: []string{
											"projects/my-project/alertPolicies/my-policy",
										},
									},
								},
							},
						},
					},
				},
			},
			want: &pb.Dashboard{
				Layout: &pb.Dashboard_GridLayout{
					GridLayout: &pb.GridLayout{
						Widgets: []*pb.Widget{
							{
								Content: &pb.Widget_IncidentList{
									IncidentList: &pb.IncidentList{
										PolicyNames: []string{
											"alertPolicies/my-policy",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:             "error when IncidentList contains wrong project",
			dashboardProject: "my-project",
			input: &pb.Dashboard{
				Layout: &pb.Dashboard_GridLayout{
					GridLayout: &pb.GridLayout{
						Widgets: []*pb.Widget{
							{
								Content: &pb.Widget_IncidentList{
									IncidentList: &pb.IncidentList{
										PolicyNames: []string{
											"projects/other-project/alertPolicies/my-policy",
										},
									},
								},
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := normalizeDashboardProto(ctx, projectMapper, tt.input, tt.dashboardProject)
			if (err != nil) != tt.wantErr {
				t.Fatalf("normalizeDashboardProto() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, tt.input, protocmp.Transform()); diff != "" {
					t.Errorf("normalizeDashboardProto() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
