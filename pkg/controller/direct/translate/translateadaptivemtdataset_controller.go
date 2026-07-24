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

package translate

import (
	"context"
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/translate/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/translate/apiv3"
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.TranslateAdaptiveMtDatasetGVK, NewTranslateAdaptiveMtDatasetModel)
}

func NewTranslateAdaptiveMtDatasetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &translateAdaptiveMtDatasetModel{config: *config}, nil
}

var _ directbase.Model = &translateAdaptiveMtDatasetModel{}

type translateAdaptiveMtDatasetModel struct {
	config config.ControllerConfig
}

func (m *translateAdaptiveMtDatasetModel) client(ctx context.Context) (*gcp.TranslationClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewTranslationRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Translate REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *translateAdaptiveMtDatasetModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.TranslateAdaptiveMtDataset{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idInterface, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idInterface.(*krm.TranslateAdaptiveMtDatasetIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idInterface)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// Normalize references in the spec
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Map to Proto once for the desired state
	mapCtx := &direct.MapContext{}
	desired := TranslateAdaptiveMtDatasetSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	return &translateAdaptiveMtDatasetAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		rawObj:    obj,
		reader:    reader,
	}, nil
}

func (m *translateAdaptiveMtDatasetModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type translateAdaptiveMtDatasetAdapter struct {
	id        *krm.TranslateAdaptiveMtDatasetIdentity
	gcpClient *gcp.TranslationClient
	desired   *pb.AdaptiveMtDataset
	actual    *pb.AdaptiveMtDataset
	rawObj    *krm.TranslateAdaptiveMtDataset
	reader    client.Reader
}

var _ directbase.Adapter = &translateAdaptiveMtDatasetAdapter{}

func (a *translateAdaptiveMtDatasetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.TranslateAdaptiveMtDataset{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(TranslateAdaptiveMtDatasetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.SetName(a.id.Dataset)
	u.SetGroupVersionKind(krm.TranslateAdaptiveMtDatasetGVK)
	u.Object = uObj
	return u, nil
}

func (a *translateAdaptiveMtDatasetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TranslateAdaptiveMtDataset", "name", a.id)

	req := &pb.GetAdaptiveMtDatasetRequest{Name: a.id.String()}
	dataset, err := a.gcpClient.GetAdaptiveMtDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting TranslateAdaptiveMtDataset %q: %w", a.id, err)
	}

	a.actual = dataset
	return true, nil
}

func (a *translateAdaptiveMtDatasetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TranslateAdaptiveMtDataset", "name", a.id)

	req := &pb.CreateAdaptiveMtDatasetRequest{
		Parent:            a.id.ParentString(),
		AdaptiveMtDataset: a.desired,
	}
	created, err := a.gcpClient.CreateAdaptiveMtDataset(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TranslateAdaptiveMtDataset %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created TranslateAdaptiveMtDataset", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *translateAdaptiveMtDatasetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating TranslateAdaptiveMtDataset", "name", a.id)

	// Since TranslateAdaptiveMtDataset is completely immutable on GCP, we check if there are any diffs in the Spec.
	// If there are diffs, we return a descriptive error so that it is surfaced on the status conditions.
	diffs, _, err := compareTranslateAdaptiveMtDataset(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)
	return fmt.Errorf("TranslateAdaptiveMtDataset resource is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
}

func (a *translateAdaptiveMtDatasetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting TranslateAdaptiveMtDataset", "name", a.id)

	req := &pb.DeleteAdaptiveMtDatasetRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAdaptiveMtDataset(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting TranslateAdaptiveMtDataset %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted TranslateAdaptiveMtDataset", "name", a.id)
	return true, nil
}

func (a *translateAdaptiveMtDatasetAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AdaptiveMtDataset) error {
	mapCtx := &direct.MapContext{}
	status := &krm.TranslateAdaptiveMtDatasetStatus{}
	status.ObservedState = TranslateAdaptiveMtDatasetObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, status, nil)
}

func compareTranslateAdaptiveMtDataset(ctx context.Context, actual, desired *pb.AdaptiveMtDataset) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, TranslateAdaptiveMtDatasetSpec_FromProto, TranslateAdaptiveMtDatasetSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.AdaptiveMtDataset) {
		// Define and populate GCP/server defaults here if any exist
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
