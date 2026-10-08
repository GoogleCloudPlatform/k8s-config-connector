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

package knowledgebase

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/dialogflow/apiv2"
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/dialogflow/generator"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
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
	registry.RegisterModel(krm.DialogflowKnowledgeBaseGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.KnowledgeBasesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewKnowledgeBasesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building KnowledgeBases REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowKnowledgeBase{}
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
	id := identity.(*krm.DialogflowKnowledgeBaseIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := generator.DialogflowKnowledgeBaseSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredProto,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.DialogflowKnowledgeBaseIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type Adapter struct {
	id        *krm.DialogflowKnowledgeBaseIdentity
	gcpClient *gcp.KnowledgeBasesClient
	desired   *pb.KnowledgeBase
	actual    *pb.KnowledgeBase
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DialogflowKnowledgeBase", "name", a.id.String())

	req := &pb.GetKnowledgeBaseRequest{Name: a.id.String()}
	kb, err := a.gcpClient.GetKnowledgeBase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DialogflowKnowledgeBase %q: %w", a.id.String(), err)
	}

	a.actual = kb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DialogflowKnowledgeBase", "name", a.id.String())

	desired := proto.CloneOf(a.desired)
	desired.Name = ""

	req := &pb.CreateKnowledgeBaseRequest{
		Parent:        a.id.ParentString(),
		KnowledgeBase: desired,
	}
	created, err := a.gcpClient.CreateKnowledgeBase(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DialogflowKnowledgeBase %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created DialogflowKnowledgeBase", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DialogflowKnowledgeBase", "name", a.id.String())

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareKnowledgeBase(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateKnowledgeBaseRequest{
			KnowledgeBase: a.desired,
			UpdateMask:    updateMask,
		}
		updated, err := a.gcpClient.UpdateKnowledgeBase(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DialogflowKnowledgeBase %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated DialogflowKnowledgeBase", "name", a.id.String())
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareKnowledgeBase(ctx context.Context, actual, desired *pb.KnowledgeBase) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, generator.DialogflowKnowledgeBaseSpec_FromProto, generator.DialogflowKnowledgeBaseSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.KnowledgeBase) {
		if obj.LanguageCode == "" {
			obj.LanguageCode = "en-US"
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.KnowledgeBase) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DialogflowKnowledgeBaseStatus{}
	status.ObservedState = generator.DialogflowKnowledgeBaseObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowKnowledgeBase{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(generator.DialogflowKnowledgeBaseSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	if a.id.Location != "" {
		obj.Spec.Location = &a.id.Location
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Knowledge_base)
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Knowledge_base)
	u.SetGroupVersionKind(krm.DialogflowKnowledgeBaseGVK)

	export.SetProjectID(u, a.id.Project)

	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DialogflowKnowledgeBase", "name", a.id.String())

	req := &pb.DeleteKnowledgeBaseRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteKnowledgeBase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DialogflowKnowledgeBase %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted DialogflowKnowledgeBase", "name", a.id.String())
	return true, nil
}
