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
// krm.kind: NetAppVolume

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/netapp/apiv1"
	netapppb "cloud.google.com/go/netapp/apiv1/netapppb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetAppVolumeGVK, NewVolumeModel)
}

func NewVolumeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelVolume{config: *config}, nil
}

var _ directbase.Model = &modelVolume{}

type modelVolume struct {
	config config.ControllerConfig
}

func (m *modelVolume) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetAppVolume{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.NetAppVolumeIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get netapp GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &VolumeAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelVolume) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//netapp.googleapis.com/") {
		return nil, nil
	}
	id := &krm.NetAppVolumeIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &VolumeAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelVolume) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Volume client: %w", err)
	}
	return gcpClient, err
}

type VolumeAdapter struct {
	id        *krm.NetAppVolumeIdentity
	gcpClient *gcp.Client
	desired   *krm.NetAppVolume
	actual    *netapppb.Volume
}

var _ directbase.Adapter = &VolumeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *VolumeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Volume", "name", a.id)

	req := &netapppb.GetVolumeRequest{
		Name: a.id.String(),
	}
	volumepb, err := a.gcpClient.GetVolume(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Volume %q: %w", a.id, err)
	}

	a.actual = volumepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *VolumeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Volume", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetAppVolumeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// GCP NetApp Volumes requires StoragePool to be the short ID (e.g. "my-pool")
	// rather than the fully-qualified resource name ("projects/.../locations/.../storagePools/my-pool").
	if resource.StoragePool != "" {
		tokens := strings.Split(resource.StoragePool, "/")
		resource.StoragePool = tokens[len(tokens)-1]
	}

	req := &netapppb.CreateVolumeRequest{
		Parent:   a.id.ParentString(),
		VolumeId: a.id.Volume,
		Volume:   resource,
	}
	op, err := a.gcpClient.CreateVolume(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Volume %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Volume %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Volume", "name", a.id)

	status := &krm.NetAppVolumeStatus{}
	status.ObservedState = NetAppVolumeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *VolumeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Volume", "name", a.id)

	normalize := func(ctx context.Context, p *netapppb.Volume) error {
		if p.GetStoragePool() != "" {
			tokens := strings.Split(p.GetStoragePool(), "/")
			p.StoragePool = tokens[len(tokens)-1]
		}
		return nil
	}

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		NetAppVolumeSpec_FromProto,
		NetAppVolumeSpec_ToProto,
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
		if resource.StoragePool != "" {
			tokens := strings.Split(resource.StoragePool, "/")
			resource.StoragePool = tokens[len(tokens)-1]
		}
		req := &netapppb.UpdateVolumeRequest{
			UpdateMask: updateMask,
			Volume:     resource,
		}
		op, err := a.gcpClient.UpdateVolume(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Volume %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Volume %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Volume", "name", a.id)
	}

	status := &krm.NetAppVolumeStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = NetAppVolumeObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *VolumeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetAppVolume{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetAppVolumeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if a.actual.GetStoragePool() != "" && a.id != nil && !strings.Contains(a.actual.GetStoragePool(), "/") {
		obj.Spec.StoragePoolRef = &krm.NetAppStoragePoolRef{
			External: fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", a.id.Project, a.id.Location, a.actual.GetStoragePool()),
		}
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.NetAppVolumeGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *VolumeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Volume", "name", a.id)

	req := &netapppb.DeleteVolumeRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteVolume(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Volume, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Volume %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Volume", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Volume %s: %w", a.id, err)
	}
	return true, nil
}
