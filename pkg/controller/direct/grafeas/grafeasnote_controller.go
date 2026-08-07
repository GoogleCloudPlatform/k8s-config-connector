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

package grafeas

import (
	"context"
	"fmt"

	grafeas "cloud.google.com/go/grafeas/apiv1"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/grafeas/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/grafeas/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	pkgmappers "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.GrafeasNoteGVK, newGrafeasNoteModel)
}

func newGrafeasNoteModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

var _ directbase.Model = &model{}

type adapter struct {
	id *krm.GrafeasNoteIdentity

	desired *pb.Note
	actual  *pb.Note

	gcp *grafeas.Client
}

var _ directbase.Adapter = &adapter{}

func (m *model) client(ctx context.Context) (*grafeas.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	opts = append(opts, option.WithEndpoint("containeranalysis.googleapis.com:443"))
	opts = append(opts, option.WithScopes("https://www.googleapis.com/auth/cloud-platform"))
	gcpClient, err := grafeas.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building grafeas client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.GrafeasNote{}
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
	desired := GrafeasNoteSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &adapter{
		id:      id.(*krm.GrafeasNoteIdentity),
		desired: desired,
		gcp:     gcp,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *adapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/notes/%s", a.id.Project, a.id.Note)
}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	req := &pb.GetNoteRequest{
		Name: a.fullyQualifiedName(),
	}
	res, err := a.gcp.GetNote(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting note %q: %w", a.fullyQualifiedName(), err)
	}
	a.actual = res
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating note", "name", a.fullyQualifiedName())

	req := &pb.CreateNoteRequest{
		Parent: a.id.ParentString(),
		NoteId: a.id.Note,
		Note:   a.desired,
	}
	res, err := a.gcp.CreateNote(ctx, req)
	if err != nil {
		return fmt.Errorf("creating note %q: %w", a.fullyQualifiedName(), err)
	}
	a.actual = res

	return a.updateStatus(ctx, createOp, a.actual)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating note", "name", a.fullyQualifiedName())

	diffs, updateMask, err := compareNote(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.fullyQualifiedName())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateNoteRequest{
		Name:       a.fullyQualifiedName(),
		Note:       a.desired,
		UpdateMask: updateMask,
	}
	res, err := a.gcp.UpdateNote(ctx, req)
	if err != nil {
		return fmt.Errorf("updating note %q: %w", a.fullyQualifiedName(), err)
	}
	a.actual = res

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting note", "name", a.fullyQualifiedName())

	req := &pb.DeleteNoteRequest{
		Name: a.fullyQualifiedName(),
	}
	err := a.gcp.DeleteNote(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("note already deleted", "name", a.fullyQualifiedName())
			return true, nil
		}
		return false, fmt.Errorf("deleting note %q: %w", a.fullyQualifiedName(), err)
	}
	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GrafeasNote{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *GrafeasNoteSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.ResourceID = &a.id.Note

	raw, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = raw
	u.SetGroupVersionKind(krm.GrafeasNoteGVK)
	return u, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Note) error {
	mapCtx := &direct.MapContext{}
	status := GrafeasNoteStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.fullyQualifiedName())
	return op.UpdateStatus(ctx, status, nil)
}

func compareNote(ctx context.Context, actual, desired *pb.Note) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := pkgmappers.OnlySpecFields(actual, GrafeasNoteSpec_FromProto, GrafeasNoteSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.Clone(desired).(*pb.Note)

	populateDefaults := func(obj *pb.Note) {
		// populate any defaults if necessary
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
