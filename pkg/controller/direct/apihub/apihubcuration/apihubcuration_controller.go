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

package apihubcuration

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/apihub/apiv1"
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

const (
	ctrlName = "apihub-curation-controller"
)

func init() {
	registry.RegisterModel(krm.APIHubCurationGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.ApiHubCurateClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewApiHubCurateRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building APIHubCurate client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.APIHubCuration{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.APIHubCurationIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := apihub.APIHubCurationSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.APIHubCurationIdentity
	gcpClient *gcp.ApiHubCurateClient
	desired   *pb.Curation
	actual    *pb.Curation
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting APIHubCuration", "name", a.id.String())

	req := &pb.GetCurationRequest{Name: a.id.String()}
	pbObj, err := a.gcpClient.GetCuration(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting APIHubCuration %q: %w", a.id.String(), err)
	}

	a.actual = pbObj
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating APIHubCuration")

	req := &pb.CreateCurationRequest{
		Parent:     a.id.ParentString(),
		CurationId: a.id.Curation,
		Curation:   a.desired,
	}

	created, err := a.gcpClient.CreateCuration(ctx, req)
	if err != nil {
		return fmt.Errorf("creating APIHubCuration %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created APIHubCuration", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating APIHubCuration", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	// Mask actual to only contain spec fields for correct diffing
	maskedActualSpec := apihub.APIHubCurationSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := apihub.APIHubCurationSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual.Name = a.id.String()

	clonedDesired := proto.Clone(a.desired).(*pb.Curation)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return nil
	}

	for _, path := range diffs.Fields {
		if path.ID == "endpoint" {
			return fmt.Errorf("field %q is immutable and cannot be updated", path.ID)
		}
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateCurationRequest{
		Curation:   a.desired,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateCuration(ctx, req)
	if err != nil {
		return fmt.Errorf("updating APIHubCuration %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.APIHubCuration{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(apihub.APIHubCurationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ProjectRef = &refs.ProjectRef{External: fmt.Sprintf("projects/%s", a.id.Project)}
	obj.Spec.ResourceID = &a.id.Curation

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting APIHubCuration", "name", a.id.String())

	req := &pb.DeleteCurationRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteCuration(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting APIHubCuration %s: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Curation) error {
	mapCtx := &direct.MapContext{}
	status := &krm.APIHubCurationStatus{}
	status.ObservedState = apihub.APIHubCurationObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := latest.Name
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, status, nil)
}
