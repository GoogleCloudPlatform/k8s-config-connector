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

	networksecurity "cloud.google.com/go/networksecurity/apiv1"
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
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

func (m *model) client(ctx context.Context) (*networksecurity.MirroringClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := networksecurity.NewMirroringRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networksecurity Mirroring REST client: %w", err)
	}
	return gcpClient, nil
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
	desired.Labels = label.GCPLabels(obj)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id.(*krm.NetworkSecurityMirroringDeploymentIdentity),
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.NetworkSecurityMirroringDeploymentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type adapter struct {
	gcpClient *networksecurity.MirroringClient
	id        *krm.NetworkSecurityMirroringDeploymentIdentity
	desired   *pb.MirroringDeployment
	actual    *pb.MirroringDeployment
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networksecurity mirroring deployment", "name", a.id)

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
	log.V(2).Info("creating networksecurity mirroring deployment", "name", a.id)

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

	actual, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("networksecurity mirroring deployment %s waiting for creation: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created networksecurity mirroring deployment", "name", a.id.String())

	return a.updateStatus(ctx, createOp, actual)
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, actual *pb.MirroringDeployment) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityMirroringDeploymentStatus{}
	status.ObservedState = NetworkSecurityMirroringDeploymentObservedState_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating networksecurity mirroring deployment", "name", a.id)

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

	latest := a.actual
	if len(paths) > 0 {
		structuredreporting.ReportDiff(ctx, report)

		desired := proto.Clone(a.desired).(*pb.MirroringDeployment)
		desired.Name = a.id.String()
		updateMask := &fieldmaskpb.FieldMask{Paths: paths}

		req := &pb.UpdateMirroringDeploymentRequest{
			UpdateMask:          updateMask,
			MirroringDeployment: desired,
		}

		op, err := a.gcpClient.UpdateMirroringDeployment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating networksecurity mirroring deployment %s: %w", a.id.String(), err)
		}

		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("networksecurity mirroring deployment %s waiting for update: %w", a.id.String(), err)
		}

		log.V(2).Info("successfully updated networksecurity mirroring deployment", "name", a.id.String())
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityMirroringDeployment{}
	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityMirroringDeploymentSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if spec != nil {
		obj.Spec = *spec
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.Mirroring_deployment

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Mirroring_deployment)
	u.SetGroupVersionKind(krm.NetworkSecurityMirroringDeploymentGVK)
	if len(a.actual.Labels) > 0 {
		u.SetLabels(a.actual.Labels)
	}
	return u, nil
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networksecurity mirroring deployment", "name", a.id)

	req := &pb.DeleteMirroringDeploymentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMirroringDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting networksecurity mirroring deployment %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("networksecurity mirroring deployment %s waiting for deletion: %w", a.id.String(), err)
	}

	return true, nil
}
