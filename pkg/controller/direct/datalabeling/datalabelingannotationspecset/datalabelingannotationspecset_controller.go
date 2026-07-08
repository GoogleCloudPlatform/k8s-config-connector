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

package datalabelingannotationspecset

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalabeling/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	directcommon "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/datalabeling/apiv1beta1"
	datalabelingpb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DataLabelingAnnotationSpecSetGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context, project string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions(config.WithDefaultQuotaProject(project))
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DataLabeling client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataLabelingAnnotationSpecSet{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := directcommon.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desiredSpecProto := DataLabelingAnnotationSpecSetSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get DataLabeling GCP client
	gcpClient, err := m.client(ctx, id.(*krm.DataLabelingAnnotationSpecSetIdentity).Project)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:               id.(*krm.DataLabelingAnnotationSpecSetIdentity),
		gcpClient:        gcpClient,
		desiredSpecProto: desiredSpecProto,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id               *krm.DataLabelingAnnotationSpecSetIdentity
	gcpClient        *gcp.Client
	desiredSpecProto *datalabelingpb.AnnotationSpecSet
	actual           *datalabelingpb.AnnotationSpecSet
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataLabelingAnnotationSpecSet", "name", a.id.String())

	req := &datalabelingpb.GetAnnotationSpecSetRequest{Name: a.id.String()}
	res, err := a.gcpClient.GetAnnotationSpecSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataLabelingAnnotationSpecSet %q: %w", a.id.String(), err)
	}

	a.actual = res
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataLabelingAnnotationSpecSet", "name", a.id.String())

	parent := fmt.Sprintf("projects/%s", a.id.Project)
	req := &datalabelingpb.CreateAnnotationSpecSetRequest{
		Parent:            parent,
		AnnotationSpecSet: a.desiredSpecProto,
	}
	created, err := a.gcpClient.CreateAnnotationSpecSet(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataLabelingAnnotationSpecSet %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created DataLabelingAnnotationSpecSet", "name", a.id.String())

	status := &krm.DataLabelingAnnotationSpecSetStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DataLabelingAnnotationSpecSetObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := created.GetName()
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataLabelingAnnotationSpecSet", "name", a.id.String())

	// Since DataLabelingAnnotationSpecSet is immutable, any diff in spec fields is an error.
	// We construct a copy of desiredSpecProto with actual's read-only fields for comparison.
	desiredProto := proto.Clone(a.desiredSpecProto).(*datalabelingpb.AnnotationSpecSet)
	desiredProto.Name = a.actual.GetName()
	desiredProto.BlockingResources = a.actual.GetBlockingResources()

	if !proto.Equal(desiredProto, a.actual) {
		return fmt.Errorf("DataLabelingAnnotationSpecSet is immutable and cannot be updated")
	}

	status := &krm.DataLabelingAnnotationSpecSetStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DataLabelingAnnotationSpecSetObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.actual.GetName()
	status.ExternalRef = &externalRef
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataLabelingAnnotationSpecSet{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataLabelingAnnotationSpecSetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataLabelingAnnotationSpecSet", "name", a.id.String())

	req := &datalabelingpb.DeleteAnnotationSpecSetRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAnnotationSpecSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DataLabelingAnnotationSpecSet %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted DataLabelingAnnotationSpecSet", "name", a.id.String())

	return true, nil
}
