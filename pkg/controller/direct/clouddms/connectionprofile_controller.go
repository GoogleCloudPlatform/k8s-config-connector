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

package clouddms

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/clouddms/apiv1"
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudDMSConnectionProfileGVK, NewConnectionProfileModel)
}

func NewConnectionProfileModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConnectionProfile{config: *config}, nil
}

var _ directbase.Model = &modelConnectionProfile{}

type modelConnectionProfile struct {
	config config.ControllerConfig
}

func (m *modelConnectionProfile) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudDMSConnectionProfile{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idObj, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idObj.(*krm.CloudDMSConnectionProfileIdentity)

	if obj.Spec.Mysql != nil && obj.Spec.Mysql.CloudSQLInstanceRef != nil {
		if err := obj.Spec.Mysql.CloudSQLInstanceRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, err
		}
	}
	if obj.Spec.Postgresql != nil {
		if obj.Spec.Postgresql.CloudSQLInstanceRef != nil {
			if err := obj.Spec.Postgresql.CloudSQLInstanceRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return nil, err
			}
		}
		if obj.Spec.Postgresql.PrivateServiceConnectConnectivity != nil && obj.Spec.Postgresql.PrivateServiceConnectConnectivity.ServiceAttachmentRef != nil {
			if err := obj.Spec.Postgresql.PrivateServiceConnectConnectivity.ServiceAttachmentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return nil, err
			}
		}
	}
	if obj.Spec.Alloydb != nil {
		if obj.Spec.Alloydb.ClusterRef != nil {
			if err := obj.Spec.Alloydb.ClusterRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return nil, err
			}
		}
		if obj.Spec.Alloydb.Settings != nil {
			if obj.Spec.Alloydb.Settings.VPCNetworkRef != nil {
				if err := obj.Spec.Alloydb.Settings.VPCNetworkRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
					return nil, err
				}
			}
			if obj.Spec.Alloydb.Settings.EncryptionConfig != nil && obj.Spec.Alloydb.Settings.EncryptionConfig.KMSKeyRef != nil {
				if err := obj.Spec.Alloydb.Settings.EncryptionConfig.KMSKeyRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
					return nil, err
				}
			}
		}
	}
	if obj.Spec.Cloudsql != nil && obj.Spec.Cloudsql.Settings != nil {
		if obj.Spec.Cloudsql.Settings.SourceRef != nil {
			if err := obj.Spec.Cloudsql.Settings.SourceRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return nil, err
			}
		}
		if obj.Spec.Cloudsql.Settings.KMSKeyRef != nil {
			if err := obj.Spec.Cloudsql.Settings.KMSKeyRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return nil, err
			}
		}
		if obj.Spec.Cloudsql.Settings.IPConfig != nil && obj.Spec.Cloudsql.Settings.IPConfig.PrivateNetworkRef != nil {
			if err := obj.Spec.Cloudsql.Settings.IPConfig.PrivateNetworkRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return nil, err
			}
		}
	}
	if obj.Spec.Oracle != nil && obj.Spec.Oracle.PrivateConnectivity != nil && obj.Spec.Oracle.PrivateConnectivity.PrivateConnectionRef != nil {
		if err := obj.Spec.Oracle.PrivateConnectivity.PrivateConnectionRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, err
		}
	}

	// Get clouddms GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newDataMigrationClient(ctx)
	if err != nil {
		return nil, err
	}
	return &ConnectionProfileAdapter{
		id:        id,
		gcpClient: client,
		desired:   obj,
	}, nil
}

func (m *modelConnectionProfile) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ConnectionProfileAdapter struct {
	id        *krm.CloudDMSConnectionProfileIdentity
	gcpClient *gcp.DataMigrationClient
	desired   *krm.CloudDMSConnectionProfile
	actual    *pb.ConnectionProfile
}

var _ directbase.Adapter = &ConnectionProfileAdapter{}

// Find retrieves the GCP resource.
func (a *ConnectionProfileAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ConnectionProfile", "name", a.id)

	req := &pb.GetConnectionProfileRequest{Name: a.id.String()}
	cp, err := a.gcpClient.GetConnectionProfile(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ConnectionProfile %q: %w", a.id, err)
	}

	a.actual = cp
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConnectionProfileAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ConnectionProfile", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudDMSConnectionProfileSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	req := &pb.CreateConnectionProfileRequest{
		Parent:              "projects/" + a.id.Project + "/locations/" + a.id.Location,
		ConnectionProfileId: a.id.ConnectionProfile,
		ConnectionProfile:   resource,
	}
	op, err := a.gcpClient.CreateConnectionProfile(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ConnectionProfile %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ConnectionProfile %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ConnectionProfile", "name", a.id)

	status := &krm.CloudDMSConnectionProfileStatus{}
	status.ObservedState = CloudDMSConnectionProfileObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConnectionProfileAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ConnectionProfile", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := CloudDMSConnectionProfileSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	paths := make(sets.Set[string])
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	// remove output only fields
	paths = paths.Delete("name")
	paths = paths.Delete("state")
	paths = paths.Delete("create_time")
	paths = paths.Delete("update_time")
	paths = paths.Delete("error")

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}
	desiredPb.Name = a.id.String()
	req := &pb.UpdateConnectionProfileRequest{
		UpdateMask:        updateMask,
		ConnectionProfile: desiredPb,
	}
	op, err := a.gcpClient.UpdateConnectionProfile(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ConnectionProfile %+v: %w", a.id, err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ConnectionProfile %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated ConnectionProfile", "name", a.id)

	status := &krm.CloudDMSConnectionProfileStatus{}
	status.ObservedState = CloudDMSConnectionProfileObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ConnectionProfileAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudDMSConnectionProfile{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudDMSConnectionProfileSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.CloudDMSConnectionProfileGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ConnectionProfileAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ConnectionProfile", "name", a.id)

	req := &pb.DeleteConnectionProfileRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteConnectionProfile(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ConnectionProfile, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ConnectionProfile %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ConnectionProfile %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted ConnectionProfile", "name", a.id)
	return true, nil
}
