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

package mapmanagement

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/maps/mapmanagement/apiv2beta"
	pb "cloud.google.com/go/maps/mapmanagement/apiv2beta/mapmanagementpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/mapmanagement/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.MapManagementStyleConfigGVK, NewMapManagementStyleConfigModel)
}

func NewMapManagementStyleConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelMapManagementStyleConfig{config: *config}, nil
}

var _ directbase.Model = &modelMapManagementStyleConfig{}

type modelMapManagementStyleConfig struct {
	config config.ControllerConfig
}

func (m *modelMapManagementStyleConfig) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building MapManagement REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelMapManagementStyleConfig) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MapManagementStyleConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	styleConfigID := id.(*krm.MapManagementStyleConfigIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := MapManagementStyleConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &mapManagementStyleConfigAdapter{
		id:        styleConfigID,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelMapManagementStyleConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.MapManagementStyleConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &mapManagementStyleConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type mapManagementStyleConfigAdapter struct {
	id        *krm.MapManagementStyleConfigIdentity
	gcpClient *gcp.Client
	desired   *pb.StyleConfig
	actual    *pb.StyleConfig
}

var _ directbase.Adapter = &mapManagementStyleConfigAdapter{}

func (a *mapManagementStyleConfigAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.StyleConfig == "" {
		return false, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("getting MapManagementStyleConfig", "name", a.id.String())

	req := &pb.GetStyleConfigRequest{Name: a.id.String()}
	styleConfigpb, err := a.gcpClient.GetStyleConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MapManagementStyleConfig %q: %w", a.id, err)
	}

	a.actual = styleConfigpb
	return true, nil
}

func (a *mapManagementStyleConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MapManagementStyleConfig", "parent", a.id.ParentString())

	req := &pb.CreateStyleConfigRequest{
		Parent:      a.id.ParentString(),
		StyleConfig: a.desired,
	}
	created, err := a.gcpClient.CreateStyleConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MapManagementStyleConfig %s: %w", a.id.ParentString(), err)
	}
	log.V(2).Info("successfully created MapManagementStyleConfig", "name", created.GetName())

	if err := a.id.FromExternal(created.GetName()); err != nil {
		return fmt.Errorf("parsing created name %q: %w", created.GetName(), err)
	}

	getReq := &pb.GetStyleConfigRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetStyleConfig(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting MapManagementStyleConfig %s after creation: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *mapManagementStyleConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MapManagementStyleConfig", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	// Mask actual to only contain spec fields for correct diffing
	maskedActualSpec := MapManagementStyleConfigSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := MapManagementStyleConfigSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.StyleConfig)

	populateDefaults := func(obj *pb.StyleConfig) {
		if obj.JsonStyleSheet == "" || obj.JsonStyleSheet == "null" {
			obj.JsonStyleSheet = ""
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	u := updateOp.GetUnstructured()
	gcpUpdateTime := direct.StringTimestamp_FromProto(mapCtx, a.actual.GetUpdateTime())
	updateTimeMatches := false
	if gcpUpdateTime != nil && u != nil {
		statusObservedStateUpdateTime, _, _ := unstructured.NestedString(u.Object, "status", "observedState", "updateTime")
		statusObservedGen, _, _ := unstructured.NestedInt64(u.Object, "status", "observedGeneration")
		if statusObservedStateUpdateTime != "" &&
			*gcpUpdateTime == statusObservedStateUpdateTime &&
			u.GetGeneration() == statusObservedGen {
			updateTimeMatches = true
		}
	}

	if updateTimeMatches {
		// If updateTime matches, the GCP side matches KRM status.
		// JsonStyleSheet is not returned by the GCP API on GET, so we copy it from desired to avoid false diffs.
		maskedActual.JsonStyleSheet = clonedDesired.JsonStyleSheet
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if diffs == nil || !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateStyleConfigRequest{
		StyleConfig: a.desired,
		UpdateMask:  updateMask,
	}

	// Make sure Name is set on update
	req.StyleConfig.Name = a.id.String()

	_, err = a.gcpClient.UpdateStyleConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating MapManagementStyleConfig %s: %w", a.id.String(), err)
	}

	getReq := &pb.GetStyleConfigRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetStyleConfig(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting MapManagementStyleConfig %s after update: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *mapManagementStyleConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.StyleConfig) error {
	mapCtx := &direct.MapContext{}
	status := &krm.MapManagementStyleConfigStatus{}
	status.ObservedState = MapManagementStyleConfigObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	readyCondition := k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage)
	return op.UpdateStatus(ctx, status, &readyCondition)
}

func (a *mapManagementStyleConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MapManagementStyleConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MapManagementStyleConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.StyleConfig)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.StyleConfig)
	u.SetGroupVersionKind(krm.MapManagementStyleConfigGVK)

	return u, nil
}

func (a *mapManagementStyleConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id.StyleConfig == "" {
		return true, nil
	}
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MapManagementStyleConfig", "name", a.id.String())

	req := &pb.DeleteStyleConfigRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteStyleConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent MapManagementStyleConfig, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting MapManagementStyleConfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted MapManagementStyleConfig", "name", a.id.String())
	return true, nil
}
