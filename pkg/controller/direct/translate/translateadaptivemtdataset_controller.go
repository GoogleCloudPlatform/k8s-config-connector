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
// proto.service: google.cloud.translation.v3.TranslationService
// proto.message: google.cloud.translation.v3.AdaptiveMtDataset
// crd.type: TranslateAdaptiveMtDataset
// crd.version: v1alpha1

package translate

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/translate/apiv3"
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/translate/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.TranslateAdaptiveMtDatasetGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.TranslationClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewTranslationRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building translate client: %w", err)
	}

	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.TranslateAdaptiveMtDataset{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.TranslateAdaptiveMtDatasetIdentity)

	mapCtx := &direct.MapContext{}
	desired := TranslateAdaptiveMtDatasetSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//translate.googleapis.com/") {
		trimmed := strings.TrimPrefix(url, "//translate.googleapis.com/")
		id := &krm.TranslateAdaptiveMtDatasetIdentity{}
		if err := id.FromExternal(trimmed); err != nil {
			log.V(2).Error(err, "url did not match TranslateAdaptiveMtDataset format", "url", url)
			return nil, nil
		}
		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}
		return &adapter{
			gcpClient: gcpClient,
			id:        id,
		}, nil
	}
	return nil, nil
}

type adapter struct {
	gcpClient *gcp.TranslationClient
	id        *krm.TranslateAdaptiveMtDatasetIdentity
	desired   *pb.AdaptiveMtDataset
	actual    *pb.AdaptiveMtDataset
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting translate adaptive mt dataset", "name", a.id.String())

	req := &pb.GetAdaptiveMtDatasetRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetAdaptiveMtDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting translate adaptive mt dataset %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating translate adaptive mt dataset", "name", a.id.String())

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	req := &pb.CreateAdaptiveMtDatasetRequest{
		Parent:            a.id.ParentString(),
		AdaptiveMtDataset: desired,
	}
	created, err := a.gcpClient.CreateAdaptiveMtDataset(ctx, req)
	if err != nil {
		return fmt.Errorf("creating translate adaptive mt dataset %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created translate adaptive mt dataset in gcp", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating translate adaptive mt dataset", "name", a.id.String())

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	diffs, err := a.compare(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	// Since TranslateAdaptiveMtDataset is completely immutable on GCP, any differences in mutable spec fields
	// should cause an update error.
	structuredreporting.ReportDiff(ctx, diffs)
	return fmt.Errorf("TranslateAdaptiveMtDataset is immutable and cannot be updated")
}

func (a *adapter) compare(ctx context.Context, actual, desired *pb.AdaptiveMtDataset) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, TranslateAdaptiveMtDatasetSpec_FromProto, TranslateAdaptiveMtDatasetSpec_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.AdaptiveMtDataset)

	// Since exampleCount is read-only / output-only on GCP, copy it from maskedActual to clonedDesired to prevent false diffs.
	clonedDesired.ExampleCount = maskedActual.ExampleCount

	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AdaptiveMtDataset) error {
	status := &krm.TranslateAdaptiveMtDatasetStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = TranslateAdaptiveMtDatasetObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.TranslateAdaptiveMtDataset{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(TranslateAdaptiveMtDatasetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Dataset)
	u.SetGroupVersionKind(krm.TranslateAdaptiveMtDatasetGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting translate adaptive mt dataset", "name", a.id.String())

	req := &pb.DeleteAdaptiveMtDatasetRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAdaptiveMtDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting translate adaptive mt dataset %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted translate adaptive mt dataset in gcp", "name", a.id.String())
	return true, nil
}
