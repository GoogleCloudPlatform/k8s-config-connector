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

package ccinsightsview

import (
	"context"
	"fmt"

	contactcenterinsights "cloud.google.com/go/contactcenterinsights/apiv1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.CCInsightsViewGVK, newModel)
}

func newModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &model{}

type adapter struct {
	id      *krm.CCInsightsViewIdentity
	desired *pb.View
	actual  *pb.View
	gcp     *contactcenterinsights.Client
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &adapter{}

func (m *model) client(ctx context.Context) (*contactcenterinsights.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := contactcenterinsights.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building contactcenterinsights client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.CCInsightsView{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ViewSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &adapter{
		id:      id.(*krm.CCInsightsViewIdentity),
		desired: desired,
		gcp:     gcp,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	req := &pb.GetViewRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcp.GetView(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting view: %w", err)
	}
	a.actual = actual
	return true, nil
}

// Create implements the Adapter interface.
func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CCInsightsView", "id", a.id)

	view := proto.Clone(a.desired).(*pb.View)
	view.Name = a.id.String()

	req := &pb.CreateViewRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		View:   view,
	}

	created, err := a.gcp.CreateView(ctx, req)
	if err != nil {
		return fmt.Errorf("creating view %s: %w", a.id.String(), err)
	}

	log.V(2).Info("created CCInsightsView", "id", a.id)
	return a.updateStatus(ctx, createOp, created)
}

// Update implements the Adapter interface.
func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CCInsightsView", "id", a.id)

	mapCtx := &direct.MapContext{}
	actualSpec := ViewSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := ViewSpec_ToProto(mapCtx, actualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.View)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "id", a.id)
		return nil
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	view := proto.Clone(clonedDesired).(*pb.View)
	view.Name = a.id.String()

	req := &pb.UpdateViewRequest{
		View:       view,
		UpdateMask: updateMask,
	}

	updated, err := a.gcp.UpdateView(ctx, req)
	if err != nil {
		return fmt.Errorf("updating view %s: %w", a.id.String(), err)
	}

	log.V(2).Info("updated CCInsightsView", "id", a.id)
	return a.updateStatus(ctx, updateOp, updated)
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CCInsightsView", "id", a.id)

	req := &pb.DeleteViewRequest{
		Name: a.id.String(),
	}

	err := a.gcp.DeleteView(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting view %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted CCInsightsView", "id", a.id)
	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	mapCtx := &direct.MapContext{}
	spec := ViewSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.CCInsightsView{}
	obj.Spec = *spec

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{Object: specObj}
	u.SetGroupVersionKind(krm.CCInsightsViewGVK)
	return u, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.View) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CCInsightsViewStatus{}
	status.ObservedState = ViewObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
