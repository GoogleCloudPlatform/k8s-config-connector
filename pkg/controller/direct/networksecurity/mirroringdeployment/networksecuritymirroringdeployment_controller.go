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

package mirroringdeployment

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityMirroringDeploymentGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (pb.MirroringClient, longrunningpb.OperationsClient, error) {
	var opts []option.ClientOption

	config := m.config
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, nil, err
	}

	opts = append(opts, option.WithEndpoint("networksecurity.googleapis.com:443"))

	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("dialing networksecurity mirroring service: %w", err)
	}

	return pb.NewMirroringClient(conn), longrunningpb.NewOperationsClient(conn), nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityMirroringDeployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}
	mapCtx := &direct.MapContext{}
	desired := NetworkSecurityMirroringDeploymentSpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, operationsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		id:               id.(*krm.NetworkSecurityMirroringDeploymentIdentity),
		desired:          desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type adapter struct {
	gcpClient        pb.MirroringClient
	operationsClient longrunningpb.OperationsClient
	id               *krm.NetworkSecurityMirroringDeploymentIdentity
	desired          *pb.MirroringDeployment
	actual           *pb.MirroringDeployment
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting networksecurity mirroring deployment", "name", a.id)

	req := &pb.GetMirroringDeploymentRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetMirroringDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networksecurity mirroring deployment %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating networksecurity mirroring deployment", "name", a.id)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateMirroringDeploymentRequest{
		Parent:                parent,
		MirroringDeploymentId: a.id.Mirroring_deployment,
		MirroringDeployment:   a.desired,
	}
	op, err := a.gcpClient.CreateMirroringDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networksecurity mirroring deployment %s: %w", a.id.String(), err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("networksecurity mirroring deployment %s waiting for creation: %w", a.id.String(), err)
	}

	actual, err := a.gcpClient.GetMirroringDeployment(ctx, &pb.GetMirroringDeploymentRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting networksecurity mirroring deployment after creation: %w", err)
	}

	log.V(2).Info("successfully created networksecurity mirroring deployment", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityMirroringDeploymentStatus{}
	status.ObservedState = NetworkSecurityMirroringDeploymentObservedState_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating networksecurity mirroring deployment", "name", a.id)

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	var paths []string

	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, a.desired.Labels)
		paths = append(paths, "labels")
	}

	if a.desired.Description != a.actual.Description {
		report.AddField("description", a.actual.Description, a.desired.Description)
		paths = append(paths, "description")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, report)

	a.desired.Name = a.id.String()
	updateMask := &fieldmaskpb.FieldMask{Paths: paths}

	req := &pb.UpdateMirroringDeploymentRequest{
		UpdateMask:          updateMask,
		MirroringDeployment: a.desired,
	}

	op, err := a.gcpClient.UpdateMirroringDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("updating networksecurity mirroring deployment %s: %w", a.id.String(), err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("networksecurity mirroring deployment %s waiting for update: %w", a.id.String(), err)
	}

	actual, err := a.gcpClient.GetMirroringDeployment(ctx, &pb.GetMirroringDeploymentRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting networksecurity mirroring deployment after update: %w", err)
	}

	log.V(2).Info("successfully updated networksecurity mirroring deployment", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityMirroringDeploymentStatus{}
	status.ObservedState = NetworkSecurityMirroringDeploymentObservedState_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting networksecurity mirroring deployment", "name", a.id)

	req := &pb.DeleteMirroringDeploymentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMirroringDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting networksecurity mirroring deployment %s: %w", a.id.String(), err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return false, fmt.Errorf("networksecurity mirroring deployment %s waiting for deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *adapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
	if op.Done {
		if op.GetError() != nil {
			return fmt.Errorf("operation failed: %s", op.GetError().GetMessage())
		}
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
			current, err := a.operationsClient.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: op.Name})
			if err != nil {
				return fmt.Errorf("getting operation %q: %w", op.Name, err)
			}
			if current.Done {
				if current.GetError() != nil {
					return fmt.Errorf("operation failed: %s", current.GetError().GetMessage())
				}
				return nil
			}
		}
	}
}
