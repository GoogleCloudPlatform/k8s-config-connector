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
	"fmt"
	"strings"

	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
)

func normalizeDashboardProto(ctx context.Context, projectMapper *projects.ProjectMapper, dashboard *pb.Dashboard, dashboardProject string) error {
	if dashboard == nil {
		return nil
	}

	// Process widgets in different layouts
	if grid := dashboard.GetGridLayout(); grid != nil {
		for _, widget := range grid.Widgets {
			if err := normalizeWidget(ctx, projectMapper, widget, dashboardProject); err != nil {
				return err
			}
		}
	}

	if mosaic := dashboard.GetMosaicLayout(); mosaic != nil {
		for _, tile := range mosaic.Tiles {
			if widget := tile.GetWidget(); widget != nil {
				if err := normalizeWidget(ctx, projectMapper, widget, dashboardProject); err != nil {
					return err
				}
			}
		}
	}

	if row := dashboard.GetRowLayout(); row != nil {
		for _, r := range row.Rows {
			for _, widget := range r.Widgets {
				if err := normalizeWidget(ctx, projectMapper, widget, dashboardProject); err != nil {
					return err
				}
			}
		}
	}

	if col := dashboard.GetColumnLayout(); col != nil {
		for _, c := range col.Columns {
			for _, widget := range c.Widgets {
				if err := normalizeWidget(ctx, projectMapper, widget, dashboardProject); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func normalizeWidget(ctx context.Context, projectMapper *projects.ProjectMapper, widget *pb.Widget, dashboardProject string) error {
	if widget == nil {
		return nil
	}

	if alertChart := widget.GetAlertChart(); alertChart != nil {
		if alertChart.Name != "" {
			parts := strings.Split(alertChart.Name, "/")
			if len(parts) == 2 && parts[0] == "alertPolicies" {
				alertChart.Name = "projects/" + dashboardProject + "/" + alertChart.Name
			} else if len(parts) >= 4 && parts[0] == "projects" && parts[2] == "alertPolicies" {
				p1, err := projectMapper.ReplaceProjectNumberWithID(ctx, parts[1])
				if err != nil {
					return err
				}
				p2, err := projectMapper.ReplaceProjectNumberWithID(ctx, dashboardProject)
				if err != nil {
					return err
				}
				if p1 != p2 {
					return fmt.Errorf("alertPolicy reference project %q does not match dashboard project %q", parts[1], dashboardProject)
				}
				parts[1] = dashboardProject
				alertChart.Name = strings.Join(parts, "/")
			}
		}
	}

	if incidentList := widget.GetIncidentList(); incidentList != nil {
		for idx, policyName := range incidentList.PolicyNames {
			if policyName != "" {
				parts := strings.Split(policyName, "/")
				if len(parts) >= 4 && parts[0] == "projects" && parts[2] == "alertPolicies" {
					p1, err := projectMapper.ReplaceProjectNumberWithID(ctx, parts[1])
					if err != nil {
						return err
					}
					p2, err := projectMapper.ReplaceProjectNumberWithID(ctx, dashboardProject)
					if err != nil {
						return err
					}
					if p1 != p2 {
						return fmt.Errorf("incidentList alertPolicy reference project %q does not match dashboard project %q", parts[1], dashboardProject)
					}
					incidentList.PolicyNames[idx] = strings.Join(parts[2:], "/")
				}
			}
		}
	}

	return nil
}
