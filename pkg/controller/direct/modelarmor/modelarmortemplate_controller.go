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
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ModelArmorTemplateGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context, location string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("modelarmor.%s.rep.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ModelArmor client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ModelArmorTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	templateID := id.(*krm.ModelArmorTemplateIdentity)

	gcpClient, err := m.client(ctx, templateID.Location)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ModelArmorTemplateSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        templateID,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.ModelArmorTemplateIdentity
	gcpClient *gcp.Client
	desired   *pb.Template
	actual    *pb.Template
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ModelArmorTemplate", "name", a.id)

	req := &pb.GetTemplateRequest{Name: a.id.String()}
	templatepb, err := a.gcpClient.GetTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ModelArmorTemplate %q: %w", a.id, err)
	}

	a.actual = templatepb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ModelArmorTemplate", "name", a.id)

	req := &pb.CreateTemplateRequest{
		Parent:     a.id.ParentString(),
		TemplateId: a.id.Template,
		Template:   a.desired,
	}
	created, err := a.gcpClient.CreateTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ModelArmorTemplate %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ModelArmorTemplate", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ModelArmorTemplate", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	// Mask actual to only contain spec fields for correct diffing
	maskedActualSpec := ModelArmorTemplateSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := ModelArmorTemplateSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.Template)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateTemplateRequest{
		Template:   a.desired,
		UpdateMask: updateMask,
	}

	// Make sure Name is set on update
	req.Template.Name = a.id.String()

	updated, err := a.gcpClient.UpdateTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ModelArmorTemplate %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Template) error {
	mapCtx := &direct.MapContext{}
	status := &krm.ModelArmorTemplateStatus{}
	status.ObservedState = ModelArmorTemplateObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = direct.PtrTo(a.id.Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Template)
	u.SetGroupVersionKind(krm.ModelArmorTemplateGVK)

	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ModelArmorTemplate", "name", a.id)

	req := &pb.DeleteTemplateRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ModelArmorTemplate, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ModelArmorTemplate %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ModelArmorTemplate", "name", a.id)
	return true, nil
}
