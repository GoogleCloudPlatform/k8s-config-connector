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

// +generated:controller
// krm.group: netapp.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: NetAppStoragePool

package netapp

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/netapp/apiv1"
	netapppb "cloud.google.com/go/netapp/apiv1/netapppb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetAppStoragePoolGVK, NewStoragePoolModel)
}

func NewStoragePoolModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelStoragePool{config: *config}, nil
}

var _ directbase.Model = &modelStoragePool{}

type modelStoragePool struct {
	config config.ControllerConfig
}

func (m *modelStoragePool) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetAppStoragePool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.NetAppStoragePoolIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get netapp GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &StoragePoolAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelStoragePool) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//netapp.googleapis.com/") {
		return nil, nil
	}
	id := &krm.NetAppStoragePoolIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &StoragePoolAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelStoragePool) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building StoragePool client: %w", err)
	}
	return gcpClient, err
}

type StoragePoolAdapter struct {
	id        *krm.NetAppStoragePoolIdentity
	gcpClient *gcp.Client
	desired   *krm.NetAppStoragePool
	actual    *netapppb.StoragePool
}

var _ directbase.Adapter = &StoragePoolAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *StoragePoolAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting StoragePool", "name", a.id)

	req := &netapppb.GetStoragePoolRequest{
		Name: a.id.String(),
	}
	storagepoolpb, err := a.gcpClient.GetStoragePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting StoragePool %q: %w", a.id, err)
	}

	a.actual = storagepoolpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *StoragePoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating StoragePool", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetAppStoragePoolSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = label.GCPLabels(a.desired)

	req := &netapppb.CreateStoragePoolRequest{
		Parent:        a.id.ParentString(),
		StoragePoolId: a.id.StoragePool,
		StoragePool:   resource,
	}
	op, err := a.gcpClient.CreateStoragePool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating StoragePool %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("StoragePool %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created StoragePool", "name", a.id)

	status := &krm.NetAppStoragePoolStatus{}
	status.ObservedState = NetAppStoragePoolObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *StoragePoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating StoragePool", "name", a.id)

	normalize := func(ctx context.Context, p *netapppb.StoragePool) error {
		if a.actual == nil {
			return nil
		}
		// GCP NetApp Volumes converts network project ID to project number in actual state.
		// Normalize desired network to match actual network if they refer to the same network.
		if p.Network != "" && a.actual.Network != "" {
			desiredParts := strings.Split(p.Network, "/")
			actualParts := strings.Split(a.actual.Network, "/")
			if len(desiredParts) == len(actualParts) && len(desiredParts) >= 4 {
				// Compare network name at end (e.g. "default")
				if desiredParts[len(desiredParts)-1] == actualParts[len(actualParts)-1] {
					p.Network = a.actual.Network
				}
			}
		}
		// GCP NetApp Volumes does not return zone or replicaZone in GET/LIST responses.
		// Because they are write-only / create-only parameters, adopt actual state to prevent false update diffs.
		if a.actual.Zone == "" {
			p.Zone = ""
		}
		if a.actual.ReplicaZone == "" {
			p.ReplicaZone = ""
		}
		return nil
	}

	resource, report, updateMask, err := common.CompareGreenfieldSpecAndLabels(
		ctx,
		updateOp.GetUnstructured(),
		&a.desired.Spec,
		a.actual,
		"labels",
		NetAppStoragePoolSpec_FromProto,
		NetAppStoragePoolSpec_ToProto,
		normalize,
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
		req := &netapppb.UpdateStoragePoolRequest{
			UpdateMask:  updateMask,
			StoragePool: resource,
		}
		op, err := a.gcpClient.UpdateStoragePool(ctx, req)
		if err != nil {
			return fmt.Errorf("updating StoragePool %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("StoragePool %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated StoragePool", "name", a.id)
	}

	status := &krm.NetAppStoragePoolStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = NetAppStoragePoolObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *StoragePoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetAppStoragePool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetAppStoragePoolSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.NetAppStoragePoolGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *StoragePoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting StoragePool", "name", a.id)

	req := &netapppb.DeleteStoragePoolRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteStoragePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent StoragePool, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting StoragePool %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted StoragePool", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete StoragePool %s: %w", a.id, err)
	}
	return true, nil
}
