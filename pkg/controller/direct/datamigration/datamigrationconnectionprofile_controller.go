// Copyright 2025 Google LLC
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

// +generated:controller
// krm.group: datamigration.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DatamigrationConnectionProfile

package datamigration

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datamigration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/clouddms/apiv1"
	clouddmspb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DatamigrationConnectionProfileGVK, NewConnectionProfileModel)
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
	obj := &krm.DatamigrationConnectionProfile{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DatamigrationConnectionProfileIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get datamigration GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConnectionProfileAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelConnectionProfile) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//datamigration.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DatamigrationConnectionProfileIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConnectionProfileAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelConnectionProfile) client(ctx context.Context) (*gcp.DataMigrationClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDataMigrationClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ConnectionProfile client: %w", err)
	}
	return gcpClient, err
}

type ConnectionProfileAdapter struct {
	id        *krm.DatamigrationConnectionProfileIdentity
	gcpClient *gcp.DataMigrationClient
	desired   *krm.DatamigrationConnectionProfile
	actual    *clouddmspb.ConnectionProfile
}

var _ directbase.Adapter = &ConnectionProfileAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ConnectionProfileAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ConnectionProfile", "name", a.id)

	req := &clouddmspb.GetConnectionProfileRequest{
		Name: a.id.String(),
	}
	connectionprofilepb, err := a.gcpClient.GetConnectionProfile(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ConnectionProfile %q: %w", a.id, err)
	}

	a.actual = connectionprofilepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConnectionProfileAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ConnectionProfile", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DatamigrationConnectionProfileSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &clouddmspb.CreateConnectionProfileRequest{
		Parent:              a.id.ParentString(),
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

	status := &krm.DatamigrationConnectionProfileStatus{}
	status.ObservedState = DatamigrationConnectionProfileObservedState_FromProto(mapCtx, created)
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

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DatamigrationConnectionProfileSpec_FromProto,
		DatamigrationConnectionProfileSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &clouddmspb.UpdateConnectionProfileRequest{
			UpdateMask:        updateMask,
			ConnectionProfile: resource,
		}
		op, err := a.gcpClient.UpdateConnectionProfile(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ConnectionProfile %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("ConnectionProfile %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ConnectionProfile", "name", a.id)
	}

	status := &krm.DatamigrationConnectionProfileStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DatamigrationConnectionProfileObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ConnectionProfileAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DatamigrationConnectionProfile{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DatamigrationConnectionProfileSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DatamigrationConnectionProfileGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ConnectionProfileAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ConnectionProfile", "name", a.id)

	req := &clouddmspb.DeleteConnectionProfileRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteConnectionProfile(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ConnectionProfile, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ConnectionProfile %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ConnectionProfile", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete ConnectionProfile %s: %w", a.id, err)
	}
	return true, nil
}
