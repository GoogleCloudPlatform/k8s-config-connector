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

package modelarmor

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/modelarmor/apiv1"
	pb "cloud.google.com/go/modelarmor/apiv1/modelarmorpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ModelArmorTemplateGVK, NewModelArmorTemplateModel)
}

func NewModelArmorTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelModelArmorTemplate{config: *config}, nil
}

var _ directbase.Model = &modelModelArmorTemplate{}

type modelModelArmorTemplate struct {
	config config.ControllerConfig
}

func (m *modelModelArmorTemplate) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ModelArmorTemplate client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelModelArmorTemplate) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ModelArmorTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := ModelArmorTemplateSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &ModelArmorTemplateAdapter{
		id:        id.(*krm.ModelArmorTemplateIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelModelArmorTemplate) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ModelArmorTemplateIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ModelArmorTemplateAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type ModelArmorTemplateAdapter struct {
	id        *krm.ModelArmorTemplateIdentity
	gcpClient *gcp.Client
	desired   *pb.Template
	actual    *pb.Template
}

var _ directbase.Adapter = &ModelArmorTemplateAdapter{}

func (a *ModelArmorTemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ModelArmorTemplate", "name", a.id)

	req := &pb.GetTemplateRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ModelArmorTemplate %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ModelArmorTemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ModelArmorTemplate", "name", a.id)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	templateID := a.id.Template

	req := &pb.CreateTemplateRequest{
		Parent:     parent,
		Template:   a.desired,
		TemplateId: templateID,
	}
	created, err := a.gcpClient.CreateTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ModelArmorTemplate %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ModelArmorTemplate", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ModelArmorTemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ModelArmorTemplate", "name", a.id.String())

	diffs, updateMask, err := compareModelArmorTemplate(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateTemplateRequest{
			Template:   a.desired,
			UpdateMask: updateMask,
		}
		req.Template.Name = a.id.String()

		updated, err := a.gcpClient.UpdateTemplate(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ModelArmorTemplate %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ModelArmorTemplateAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Template) error {
	mapCtx := &direct.MapContext{}
	status := ModelArmorTemplateStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ModelArmorTemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ModelArmorTemplate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ModelArmorTemplateSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Template)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Template)
	u.SetGroupVersionKind(krm.ModelArmorTemplateGVK)

	export.SetProjectID(u, a.id.Project)

	return u, nil
}

func (a *ModelArmorTemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ModelArmorTemplate", "name", a.id)

	req := &pb.DeleteTemplateRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ModelArmorTemplate %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ModelArmorTemplate", "name", a.id)
	return true, nil
}

func compareModelArmorTemplate(ctx context.Context, actual, desired *pb.Template) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ModelArmorTemplateSpec_FromProto, ModelArmorTemplateSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = actual.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.Clone(desired).(*pb.Template)
	clonedDesired.Name = actual.Name

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
