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

package conversationdataset

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
	registry.RegisterModel(krm.DialogflowConversationDatasetGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.ConversationDatasetsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewConversationDatasetsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ConversationDatasets REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowConversationDataset{}
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
	id := identity.(*krm.DialogflowConversationDatasetIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := DialogflowConversationDatasetSpec_ToProto(mapCtx, &obj.Spec)
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
	id := &krm.DialogflowConversationDatasetIdentity{}
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
	id        *krm.DialogflowConversationDatasetIdentity
	gcpClient *gcp.ConversationDatasetsClient
	desired   *pb.ConversationDataset
	actual    *pb.ConversationDataset
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DialogflowConversationDataset", "name", a.id.String())

	if a.id.ConversationDataset == "" {
		return false, nil
	}

	req := &pb.GetConversationDatasetRequest{Name: a.id.String()}
	dataset, err := a.gcpClient.GetConversationDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DialogflowConversationDataset %q: %w", a.id.String(), err)
	}

	a.actual = dataset
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DialogflowConversationDataset", "parent", a.id.ParentString())

	desired := proto.CloneOf(a.desired)
	desired.Name = ""

	req := &pb.CreateConversationDatasetRequest{
		Parent:              a.id.ParentString(),
		ConversationDataset: desired,
	}
	op, err := a.gcpClient.CreateConversationDataset(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DialogflowConversationDataset %s: %w", a.id.ParentString(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of DialogflowConversationDataset %s: %w", a.id.ParentString(), err)
	}

	if created != nil && created.GetName() != "" {
		if err := a.id.FromExternal(created.GetName()); err != nil {
			return fmt.Errorf("parsing created name %q: %w", created.GetName(), err)
		}
	}

	log.V(2).Info("successfully created DialogflowConversationDataset", "name", a.id.String())

	latest, err := a.gcpClient.GetConversationDataset(ctx, &pb.GetConversationDatasetRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting created DialogflowConversationDataset %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DialogflowConversationDataset", "name", a.id.String())

	a.desired.Name = a.id.String()

	diffs, _, err := compareConversationDataset(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)
	return fmt.Errorf("DialogflowConversationDataset is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
}

func compareConversationDataset(ctx context.Context, actual, desired *pb.ConversationDataset) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DialogflowConversationDatasetSpec_FromProto, DialogflowConversationDatasetSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.ConversationDataset) {
		// Populate defaults if any
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ConversationDataset) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DialogflowConversationDatasetStatus{}
	status.ObservedState = DialogflowConversationDatasetObservedState_FromProto(mapCtx, latest)
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

	obj := &krm.DialogflowConversationDataset{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowConversationDatasetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	if a.id.Location != "" {
		obj.Spec.Location = &a.id.Location
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ConversationDataset)
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ConversationDataset)
	u.SetGroupVersionKind(krm.DialogflowConversationDatasetGVK)

	export.SetProjectID(u, a.id.Project)

	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DialogflowConversationDataset", "name", a.id.String())

	if a.id.ConversationDataset == "" {
		return true, nil
	}

	req := &pb.DeleteConversationDatasetRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteConversationDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DialogflowConversationDataset %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return false, fmt.Errorf("waiting for deletion of DialogflowConversationDataset %s: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully deleted DialogflowConversationDataset", "name", a.id.String())
	return true, nil
}
