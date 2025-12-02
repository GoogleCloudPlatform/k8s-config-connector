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
// proto.message: google.cloud.dataplex.v1.Lake
// crd.type: DataplexLake
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexLakeGVK, NewLakeModel)
}

func NewLakeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &lakeModel{config: config}, nil
}

var _ directbase.Model = &lakeModel{}

type lakeModel struct {
	config *config.ControllerConfig
}

func (m *lakeModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataplexLake{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	copied := obj.DeepCopy()
	mapCtx := &direct.MapContext{}
	desired := DataplexLakeSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	idI, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idI.(*krm.LakeIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idI)
	}

	lakeAdapter := &lakeAdapter{
		id:      id,
		desired: desired,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	lakeClient, err := gcpClient.client(ctx)
	if err != nil {
		return nil, err
	}
	lakeAdapter.gcpClient = lakeClient

	return lakeAdapter, nil
}

func (m *lakeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type lakeAdapter struct {
	gcpClient *gcp.Client
	id        *krm.LakeIdentity
	desired   *pb.Lake
	actual    *pb.Lake
	reader    client.Reader
}

var _ directbase.Adapter = &lakeAdapter{}

func (a *lakeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex lake", "name", a.id)

	req := &pb.GetLakeRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetLake(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, nil
	}

	a.actual = actual
	return true, nil
}

func (a *lakeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex lake", "name", a.id)

	req := &pb.CreateLakeRequest{
		Parent: a.id.Parent().String(),
		Lake:   a.desired,
		LakeId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateLake(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex lake %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex lake in gcp", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.DataplexLakeStatus{}
	status.ObservedState = DataplexLakeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *lakeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex lake", "name", a.id)

	a.desired.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(a.desired.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	// a.actual.Metastore != nil
	// default value: metastore: {} (i.e. metastore.service: "")
	if a.desired.Metastore != nil {
		if !reflect.DeepEqual(a.desired.Metastore, a.actual.Metastore) {
			updateMask.Paths = append(updateMask.Paths, "metastore")
		}
	} else if a.actual.Metastore.Service != "" {
		updateMask.Paths = append(updateMask.Paths, "metastore")
	}

	var updated *pb.Lake
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)

		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		req := &pb.UpdateLakeRequest{
			UpdateMask: updateMask,
			Lake:       a.desired,
		}
		op, err := a.gcpClient.UpdateLake(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex lake %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of dataplex lake %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex lake", "name", a.id)
	}

	mapCtx := &direct.MapContext{}
	status := &krm.DataplexLakeStatus{}
	status.ObservedState = DataplexLakeObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *lakeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexLake{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexLakeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ParentRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refs.ProjectRef{External: a.id.Parent().ProjectID},
		Location:   a.id.Parent().Location,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DataplexLakeGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *lakeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex lake", "name", a.id)

	req := &pb.DeleteLakeRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteLake(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting dataplex lake %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex lake", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of dataplex lake %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
