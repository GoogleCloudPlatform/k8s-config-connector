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
	"fmt"
	"strings"

	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BackupIdentity{}
	_ identity.Resource   = &BigtableBackup{}
)

var BackupIdentityFormat = gcpurls.Template[BackupIdentity]("bigtableadmin.googleapis.com", "projects/{project}/instances/{instance}/clusters/{cluster}/backups/{backup}")

// +k8s:deepcopy-gen=false
type BackupIdentity struct {
	Project  string
	Instance string
	Cluster  string
	Backup   string
}

func (i *BackupIdentity) String() string {
	return BackupIdentityFormat.ToString(*i)
}

func (i *BackupIdentity) FromExternal(ref string) error {
	parsed, match, err := BackupIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigtableBackup external=%q was not known (use %s): %w", ref, BackupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigtableBackup external=%q was not known (use %s)", ref, BackupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BackupIdentity) Host() string {
	return BackupIdentityFormat.Host()
}

func (i *BackupIdentity) ID() string {
	return i.Backup
}

func (i *BackupIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/instances/%s/clusters/%s", i.Project, i.Instance, i.Cluster)
}

func getIdentityFromBigtableBackupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BackupIdentity, error) {
	backup := &BigtableBackup{}
	if u, ok := obj.(*unstructured.Unstructured); ok {
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, backup); err != nil {
			return nil, fmt.Errorf("failed to convert from unstructured: %w", err)
		}
	} else if typed, ok := obj.(*BigtableBackup); ok {
		backup = typed
	} else {
		return nil, fmt.Errorf("expected BigtableBackup or *unstructured.Unstructured, got %T", obj)
	}

	resourceID := common.ValueOf(backup.Spec.ResourceID)
	if resourceID == "" {
		resourceID = backup.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Resolve clusterRef
	clusterRef := backup.Spec.ClusterRef
	clusterExternal, err := clusterRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	instanceParent, clusterID, err := ParseClusterExternal(clusterExternal)
	if err != nil {
		return nil, err
	}

	return &BackupIdentity{
		Project:  instanceParent.Parent.ProjectID,
		Instance: instanceParent.Id,
		Cluster:  clusterID,
		Backup:   resourceID,
	}, nil
}

func (obj *BigtableBackup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigtableBackupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &BackupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BigtableBackup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (c *BigtableBackup) ExternalIdentifier() *string {
	if c.Status.ExternalRef != nil {
		return c.Status.ExternalRef
	}
	return nil
}

func ParseBackupExternal(external string) (*ClusterIdentity, string, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "instances" || tokens[4] != "clusters" || tokens[6] != "backups" {
		return nil, "", fmt.Errorf("format of BigtableBackup external=%q was not known (use projects/{{projectID}}/instances/{{instanceID}}/clusters/{{clusterID}}/backups/{{backupID}})", external)
	}
	p := &ClusterIdentity{
		parent: &bigtablev1beta1.InstanceIdentity{
			Parent: &parent.ProjectParent{ProjectID: tokens[1]},
			Id:     tokens[3],
		},
		id: tokens[5],
	}
	resourceID := tokens[7]
	return p, resourceID, nil
}

func NewBackupIdentity(ctx context.Context, reader client.Reader, obj *BigtableBackup) (*BackupIdentity, error) {
	return getIdentityFromBigtableBackupSpec(ctx, reader, obj)
}
