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

// +tool:controller
// proto.service: google.cloud.dataplex.v1.DataTaxonomyService
// proto.message: google.cloud.dataplex.v1.DataAttributeBinding
// crd.type: DataplexDataAttributeBinding
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexDataAttributeBindingGVK, NewDataAttributeBindingModel)
}

func NewDataAttributeBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataAttributeBindingModel{config: config}, nil
}

var _ directbase.Model = &dataAttributeBindingModel{}

type dataAttributeBindingModel struct {
	config *config.ControllerConfig
}

func (m *dataAttributeBindingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexDataAttributeBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	copied := obj.DeepCopy()
	mapCtx := &direct.MapContext{}
	desired := DataplexDataAttributeBindingSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	idI, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idI.(*krm.DataplexDataAttributeBindingIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idI)
	}

	dataAttributeBindingAdapter := &dataAttributeBindingAdapter{
		id:      id,
		desired: desired,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	dtClient, err := gcpClient.dataTaxonomyClient(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	dataAttributeBindingAdapter.gcpClient = dtClient

	return dataAttributeBindingAdapter, nil
}

func (m *dataAttributeBindingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.DataplexDataAttributeBindingIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	dtClient, err := gcpClient.dataTaxonomyClient(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &dataAttributeBindingAdapter{
		id:        id,
		gcpClient: dtClient,
	}, nil
}

type dataAttributeBindingAdapter struct {
	gcpClient *gcp.DataTaxonomyClient
	id        *krm.DataplexDataAttributeBindingIdentity
	desired   *pb.DataAttributeBinding
	actual    *pb.DataAttributeBinding
	reader    client.Reader
}

var _ directbase.Adapter = &dataAttributeBindingAdapter{}

func (a *dataAttributeBindingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex data attribute binding", "name", a.id)

	req := &pb.GetDataAttributeBindingRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetDataAttributeBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = actual
	return true, nil
}

func (a *dataAttributeBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex data attribute binding", "name", a.id)

	req := &pb.CreateDataAttributeBindingRequest{
		Parent:                 a.id.ParentString(),
		DataAttributeBinding:   a.desired,
		DataAttributeBindingId: a.id.Dataattributebinding,
	}
	op, err := a.gcpClient.CreateDataAttributeBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex data attribute binding %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex data attribute binding %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex data attribute binding in gcp", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.DataplexDataAttributeBindingStatus{}
	status.ObservedState = DataplexDataAttributeBindingObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *dataAttributeBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex data attribute binding", "name", a.id)

	a.desired.Name = a.id.String()

	paths, err := common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	var updated *pb.DataAttributeBinding
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)

		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)
		req := &pb.UpdateDataAttributeBindingRequest{
			UpdateMask:           &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
			DataAttributeBinding: a.desired,
		}
		op, err := a.gcpClient.UpdateDataAttributeBinding(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex data attribute binding %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of dataplex data attribute binding %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex data attribute binding", "name", a.id)
	}

	mapCtx := &direct.MapContext{}
	status := &krm.DataplexDataAttributeBindingStatus{}
	status.ObservedState = DataplexDataAttributeBindingObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *dataAttributeBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexDataAttributeBinding{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexDataAttributeBindingSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.Dataattributebinding

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{}
	u.Object = uObj
	u.SetName(a.id.Dataattributebinding)
	u.SetGroupVersionKind(krm.DataplexDataAttributeBindingGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

// Delete implements the Adapter interface.
func (a *dataAttributeBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex data attribute binding", "name", a.id)

	req := &pb.DeleteDataAttributeBindingRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDataAttributeBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex data attribute binding %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex data attribute binding", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of dataplex data attribute binding %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
