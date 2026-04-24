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

package securitycenter

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewEventThreatDetectionCustomModuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEventThreatDetectionCustomModule{config: *config}, nil
}

var _ directbase.Model = &modelEventThreatDetectionCustomModule{}

type modelEventThreatDetectionCustomModule struct {
	config config.ControllerConfig
}

func (m *modelEventThreatDetectionCustomModule) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecurityCenterEventThreatDetectionCustomModule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newSecurityCenterClient(ctx)
	if err != nil {
		return nil, err
	}

	return &EventThreatDetectionCustomModuleAdapter{
		id:        id.(*krm.SecurityCenterEventThreatDetectionCustomModuleIdentity),
		gcpClient: client,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelEventThreatDetectionCustomModule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type EventThreatDetectionCustomModuleAdapter struct {
	id        *krm.SecurityCenterEventThreatDetectionCustomModuleIdentity
	gcpClient *securitycenter.Client
	desired   *krm.SecurityCenterEventThreatDetectionCustomModule
	actual    *pb.EventThreatDetectionCustomModule
	reader    client.Reader
}

var _ directbase.Adapter = &EventThreatDetectionCustomModuleAdapter{}

func (a *EventThreatDetectionCustomModuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecurityCenterEventThreatDetectionCustomModule", "name", a.id)

	res, err := a.gcpClient.GetEventThreatDetectionCustomModule(ctx, &pb.GetEventThreatDetectionCustomModuleRequest{Name: a.id.String()})
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecurityCenterEventThreatDetectionCustomModule %q: %w", a.id, err)
	}

	a.actual = res
	return true, nil
}

func (a *EventThreatDetectionCustomModuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SecurityCenterEventThreatDetectionCustomModule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := SecurityCenterEventThreatDetectionCustomModuleSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateEventThreatDetectionCustomModuleRequest{
		Parent:                           a.id.Parent().String(),
		EventThreatDetectionCustomModule: desiredPb,
	}
	created, err := a.gcpClient.CreateEventThreatDetectionCustomModule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SecurityCenterEventThreatDetectionCustomModule %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created SecurityCenterEventThreatDetectionCustomModule", "name", a.id)

	status := &krm.SecurityCenterEventThreatDetectionCustomModuleStatus{}
	status.ObservedState = SecurityCenterEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *EventThreatDetectionCustomModuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SecurityCenterEventThreatDetectionCustomModule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := SecurityCenterEventThreatDetectionCustomModuleSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths, report, err := common.CompareProtoMessageStructuredDiff(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		if a.desired.Status.ExternalRef == nil {
			status := &krm.SecurityCenterEventThreatDetectionCustomModuleStatus{}
			status.ObservedState = SecurityCenterEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx, a.actual)
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	updateOp.RecordUpdatingEvent()

	desiredPb.Name = a.id.String()
	req := &pb.UpdateEventThreatDetectionCustomModuleRequest{
		EventThreatDetectionCustomModule: desiredPb,
		UpdateMask:                       &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
	}
	updated, err := a.gcpClient.UpdateEventThreatDetectionCustomModule(ctx, req)
	if err != nil {
		return fmt.Errorf("updating SecurityCenterEventThreatDetectionCustomModule %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated SecurityCenterEventThreatDetectionCustomModule", "name", a.id)

	status := &krm.SecurityCenterEventThreatDetectionCustomModuleStatus{}
	status.ObservedState = SecurityCenterEventThreatDetectionCustomModuleObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if a.desired.Status.ExternalRef == nil {
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *EventThreatDetectionCustomModuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SecurityCenterEventThreatDetectionCustomModule", "name", a.id)

	err := a.gcpClient.DeleteEventThreatDetectionCustomModule(ctx, &pb.DeleteEventThreatDetectionCustomModuleRequest{Name: a.id.String()})
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting SecurityCenterEventThreatDetectionCustomModule %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted SecurityCenterEventThreatDetectionCustomModule", "name", a.id)
	return true, nil
}

func (a *EventThreatDetectionCustomModuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}
