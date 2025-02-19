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

package migrationcenter

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/migrationcenter/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	migrationcenterpb "cloud.google.com/go/migrationcenter/v1/migrationcenterpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MigrationcenterPreferenceSetGVK, NewPreferenceSetModel)
}

func NewPreferenceSetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPreferenceSet{config: *config}, nil
}

var _ directbase.Model = &modelPreferenceSet{}

type modelPreferenceSet struct {
	config config.ControllerConfig
}

func (m *modelPreferenceSet) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PreferenceSet client: %w", err)
	}
	return gcpClient, err
}

func (m *modelPreferenceSet) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.MigrationcenterPreferenceSet{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewPreferenceSetIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get migrationcenter GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PreferenceSetAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPreferenceSet) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type PreferenceSetAdapter struct {
	id        *krm.PreferenceSetIdentity
	gcpClient *gcp.Client
	desired   *krm.MigrationcenterPreferenceSet
	actual    *migrationcenterpb.PreferenceSet
}

var _ directbase.Adapter = &PreferenceSetAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PreferenceSetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting PreferenceSet", "name", a.id)

	req := &migrationcenterpb.GetPreferenceSetRequest{Name: a.id.String()}
	preferencesetpb, err := a.gcpClient.GetPreferenceSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting PreferenceSet %q: %w", a.id, err)
	}

	a.actual = preferencesetpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PreferenceSetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PreferenceSet", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := MigrationcenterPreferenceSetSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &migrationcenterpb.CreatePreferenceSetRequest{
		Parent:        a.id.Parent().String(),
		PreferenceSet: resource,
	}
	op, err := a.gcpClient.CreatePreferenceSet(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PreferenceSet %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("PreferenceSet %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created PreferenceSet", "name", a.id)

	status := &krm.MigrationcenterPreferenceSetStatus{}
	status.ObservedState = MigrationcenterPreferenceSetObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PreferenceSetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PreferenceSet", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := MigrationcenterPreferenceSetSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.MigrationcenterPreferenceSetStatus{}
		status.ObservedState = MigrationcenterPreferenceSetObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &migrationcenterpb.UpdatePreferenceSetRequest{
		Name:          a.id,
		UpdateMask:    updateMask,
		PreferenceSet: desiredPb,
	}
	op, err := a.gcpClient.UpdatePreferenceSet(ctx, req)
	if err != nil {
		return fmt.Errorf("updating PreferenceSet %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("PreferenceSet %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated PreferenceSet", "name", a.id)

	status := &krm.MigrationcenterPreferenceSetStatus{}
	status.ObservedState = MigrationcenterPreferenceSetObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PreferenceSetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MigrationcenterPreferenceSet{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MigrationcenterPreferenceSetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Id)
	u.SetGroupVersionKind(krm.MigrationcenterPreferenceSetGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PreferenceSetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PreferenceSet", "name", a.id)

	req := &migrationcenterpb.DeletePreferenceSetRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeletePreferenceSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent PreferenceSet, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting PreferenceSet %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted PreferenceSet", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete PreferenceSet %s: %w", a.id, err)
	}
	return true, nil
}
