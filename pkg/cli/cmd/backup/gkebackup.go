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

package backup

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/gkebackup/v1"
	"google.golang.org/api/option"
)

// ClusterBackupMetadata captures details of an associated GKE cluster backup.
type ClusterBackupMetadata struct {
	Provider       string `json:"provider"`
	BackupPlan     string `json:"backupPlan,omitempty"`
	BackupName     string `json:"backupName,omitempty"`
	State          string `json:"state,omitempty"`
	ErrorMessage   string `json:"errorMessage,omitempty"`
	IncludeVolumes bool   `json:"includeVolumes"`
	AllNamespaces  bool   `json:"allNamespaces"`
}

// GKEBackupManager interacts with the Backup for GKE API.
type GKEBackupManager struct {
	service *gkebackup.Service
}

// NewGKEBackupManager creates a new GKEBackupManager.
func NewGKEBackupManager(ctx context.Context) (*GKEBackupManager, error) {
	var opts []option.ClientOption
	if httpClient := ctx.Value(oauth2.HTTPClient); httpClient != nil {
		opts = append(opts, option.WithHTTPClient(httpClient.(*http.Client)))
	}
	svc, err := gkebackup.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("creating gkebackup service: %w", err)
	}
	return &GKEBackupManager{service: svc}, nil
}

// EnsureBackupPlan verifies or creates a BackupPlan for the given cluster.
func (m *GKEBackupManager) EnsureBackupPlan(ctx context.Context, project, location, cluster, planName, schedule string, retentionDays int, lockRetention bool) (*gkebackup.BackupPlan, error) {
	parent := fmt.Sprintf("projects/%s/locations/%s", project, location)
	fullPlanName := fmt.Sprintf("%s/backupPlans/%s", parent, planName)

	existing, err := m.service.Projects.Locations.BackupPlans.Get(fullPlanName).Context(ctx).Do()
	if err == nil {
		return existing, nil
	}

	clusterRef := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, location, cluster)
	plan := &gkebackup.BackupPlan{
		Cluster:     clusterRef,
		Description: fmt.Sprintf("Automated holistic cluster backup for %s managed by Config Connector", cluster),
		BackupConfig: &gkebackup.BackupConfig{
			AllNamespaces:     true,
			IncludeSecrets:    true,
			IncludeVolumeData: true,
		},
	}
	if retentionDays > 0 {
		plan.RetentionPolicy = &gkebackup.RetentionPolicy{
			BackupRetainDays: int64(retentionDays),
			Locked:           lockRetention,
		}
	}
	if schedule != "" {
		plan.BackupSchedule = &gkebackup.Schedule{
			CronSchedule: schedule,
		}
	}

	_, err = m.service.Projects.Locations.BackupPlans.Create(parent, plan).BackupPlanId(planName).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("creating BackupPlan %s: %w", planName, err)
	}
	return plan, nil
}

// TriggerBackup initiates a point-in-time backup under the cluster's BackupPlan.
func (m *GKEBackupManager) TriggerBackup(ctx context.Context, project, location, planName string, retentionDays int) (*ClusterBackupMetadata, error) {
	parent := fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", project, location, planName)
	backupID := fmt.Sprintf("kcc-%s", time.Now().UTC().Format("20060102-150405"))

	req := &gkebackup.Backup{
		Description: "Point-in-time cluster backup triggered alongside Config Connector backup",
	}
	if retentionDays > 0 {
		req.RetainDays = int64(retentionDays)
	}

	resp, err := m.service.Projects.Locations.BackupPlans.Backups.Create(parent, req).BackupId(backupID).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("triggering GKE backup under %s: %w", planName, err)
	}

	backupName := resp.Name
	if backupName == "" {
		backupName = fmt.Sprintf("%s/backups/%s", parent, backupID)
	}

	return &ClusterBackupMetadata{
		Provider:       "gkebackup.googleapis.com",
		BackupPlan:     parent,
		BackupName:     backupName,
		State:          "IN_PROGRESS",
		IncludeVolumes: true,
		AllNamespaces:  true,
	}, nil
}
