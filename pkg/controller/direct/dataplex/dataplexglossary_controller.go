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
// proto.service: google.cloud.dataplex.v1.BusinessGlossary
// proto.message: google.cloud.dataplex.v1.Glossary
// crd.type: DataplexGlossary
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexGlossaryGVK, NewGlossaryModel)
}

func NewGlossaryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &glossaryModel{config: config}, nil
}

var _ directbase.Model = &glossaryModel{}

type glossaryModel struct {
	config *config.ControllerConfig
}

func (m *glossaryModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexGlossary{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	copied := obj.DeepCopy()
	mapCtx := &direct.MapContext{}
	desired := DataplexGlossarySpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	idI, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idI.(*krm.GlossaryIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idI)
	}

	glossaryAdapter := &glossaryAdapter{
		id:      id,
		desired: desired,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	gcpGlossaryClient, err := gcpClient.businessGlossaryClient(ctx)
	if err != nil {
		return nil, err
	}
	glossaryAdapter.gcpClient = gcpGlossaryClient

	return glossaryAdapter, nil
}

func (m *glossaryModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type glossaryAdapter struct {
	gcpClient *gcp.BusinessGlossaryClient
	id        *krm.GlossaryIdentity
	desired   *pb.Glossary
	actual    *pb.Glossary
	reader    client.Reader
}

var _ directbase.Adapter = &glossaryAdapter{}

func (a *glossaryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex glossary", "name", a.id)

	req := &pb.GetGlossaryRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetGlossary(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex glossary %s: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *glossaryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex glossary", "name", a.id)

	req := &pb.CreateGlossaryRequest{
		Parent:     a.id.ParentString(),
		Glossary:   a.desired,
		GlossaryId: a.id.Glossary,
	}
	op, err := a.gcpClient.CreateGlossary(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex glossary %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex glossary %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex glossary in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *glossaryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex glossary", "name", a.id)

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareGlossary(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *pb.Glossary
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateGlossaryRequest{
			UpdateMask: updateMask,
			Glossary:   a.desired,
		}
		op, err := a.gcpClient.UpdateGlossary(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex glossary %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of dataplex glossary %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex glossary", "name", a.id)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func compareGlossary(ctx context.Context, actual, desired *pb.Glossary) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	actualSpec := DataplexGlossarySpec_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := DataplexGlossarySpec_ToProto(mapCtx, actualSpec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.Clone(desired).(*pb.Glossary)

	return common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
}

func (a *glossaryAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Glossary) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DataplexGlossaryStatus{}
	status.ObservedState = DataplexGlossaryObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *glossaryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexGlossary{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexGlossarySpec_FromProto(mapCtx, a.actual))
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
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DataplexGlossaryGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *glossaryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex glossary", "name", a.id)

	req := &pb.DeleteGlossaryRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteGlossary(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex glossary %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex glossary", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of dataplex glossary %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
