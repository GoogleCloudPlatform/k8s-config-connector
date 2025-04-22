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

// +tool:controller
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Zone
// crd.type: DataplexZone
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"
	"reflect"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexZoneGVK, NewZoneModel)
}

func NewZoneModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &zoneModel{config: config}, nil
}

var _ directbase.Model = &zoneModel{}

type zoneModel struct {
	config *config.ControllerConfig
}

func (m *zoneModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataplexZone{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewZoneIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	zoneAdapter := &zoneAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	zoneClient, err := gcpClient.client(ctx)
	if err != nil {
		return nil, err
	}
	zoneAdapter.gcpClient = zoneClient

	return zoneAdapter, nil
}

func (m *zoneModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type zoneAdapter struct {
	gcpClient *gcp.Client
	id        *krm.ZoneIdentity
	desired   *krm.DataplexZone
	actual    *pb.Zone
	reader    client.Reader
}

var _ directbase.Adapter = &zoneAdapter{}

func (a *zoneAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex zone", "name", a.id)

	req := &pb.GetZoneRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetZone(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		// Returning 'false, nil' means the error is retryable by the controller manager
		return false, nil
	}

	a.actual = actual
	return true, nil
}

func (a *zoneAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex zone", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	desired.Name = a.id.String()

	zone := DataplexZoneSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateZoneRequest{
		Parent: a.id.Parent(),
		Zone:   zone,
		ZoneId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateZone(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex zone %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex zone %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex zone in gcp", "name", a.id)

	status := &krm.DataplexZoneStatus{}
	status.ObservedState = DataplexZoneObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *zoneAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex zone", "name", a.id)

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	zone := DataplexZoneSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	zone.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(zone.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(zone.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(zone.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	// API default value:
	//"discoverySpec": {
	//    "csvOptions": {},
	//    "jsonOptions": {},
	//    "schedule": ""
	//  })
	emptyDiscoverySpec := &pb.Zone_DiscoverySpec{
		CsvOptions:  &pb.Zone_DiscoverySpec_CsvOptions{},
		JsonOptions: &pb.Zone_DiscoverySpec_JsonOptions{},
		Trigger:     &pb.Zone_DiscoverySpec_Schedule{Schedule: ""},
	}
	if zone.DiscoverySpec != nil {
		if !cmp.Equal(zone.DiscoverySpec, a.actual.DiscoverySpec, cmpopts.IgnoreUnexported(pb.Zone_DiscoverySpec{}, pb.Zone_DiscoverySpec_CsvOptions{}, pb.Zone_DiscoverySpec_JsonOptions{})) {
			updateMask.Paths = append(updateMask.Paths, "discovery_spec")
		}
	} else {
		if !cmp.Equal(emptyDiscoverySpec, a.actual.DiscoverySpec, cmpopts.IgnoreUnexported(pb.Zone_DiscoverySpec{}, pb.Zone_DiscoverySpec_CsvOptions{}, pb.Zone_DiscoverySpec_JsonOptions{})) {
			updateMask.Paths = append(updateMask.Paths, "discovery_spec")
		}
	}
	// Type is immutable, no need to check Type
	// ResourceSpec.LocationType is immutable, no need to check ResourceSpec

	var updated *pb.Zone
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)

		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		log.V(2).Info("updating fields", "name", a.id, "paths", updateMask.Paths)
		req := &pb.UpdateZoneRequest{
			UpdateMask: updateMask,
			Zone:       zone,
		}
		op, err := a.gcpClient.UpdateZone(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex zone %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of dataplex zone %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex zone", "name", a.id)
	}

	status := &krm.DataplexZoneStatus{}
	status.ObservedState = DataplexZoneObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *zoneAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexZone{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexZoneSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.String()) // This will likely be overridden by the caller
	u.SetGroupVersionKind(krm.DataplexZoneGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *zoneAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex zone", "name", a.id)

	req := &pb.DeleteZoneRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteZone(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting dataplex zone %s: %w", a.id.String(), err)
	}

	if !op.Done() {
		log.V(2).Info("waiting for deletion of dataplex zone", "name", a.id)
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of dataplex zone %s: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully deleted dataplex zone", "name", a.id)

	return true, nil
}
