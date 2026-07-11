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

package cloudsecuritycompliance

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsecuritycompliance/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/cloudsecuritycompliance/apiv1"
	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudSecurityComplianceFrameworkDeploymentGVK, NewFrameworkDeploymentModel)
}

func NewFrameworkDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFrameworkDeployment{config: *config}, nil
}

var _ directbase.Model = &modelFrameworkDeployment{}

type modelFrameworkDeployment struct {
	config config.ControllerConfig
}

func (m *modelFrameworkDeployment) client(ctx context.Context) (*gcp.DeploymentClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDeploymentRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CloudSecurityCompliance client: %w", err)
	}
	return gcpClient, err
}

func (m *modelFrameworkDeployment) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudSecurityComplianceFrameworkDeployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}

	copied := obj.DeepCopy()
	desired := CloudSecurityComplianceFrameworkDeploymentSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &FrameworkDeploymentAdapter{
		id:        id.(*krm.CloudSecurityComplianceFrameworkDeploymentIdentity),
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelFrameworkDeployment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type FrameworkDeploymentAdapter struct {
	id        *krm.CloudSecurityComplianceFrameworkDeploymentIdentity
	gcpClient *gcp.DeploymentClient
	desired   *pb.FrameworkDeployment
	actual    *pb.FrameworkDeployment
}

var _ directbase.Adapter = &FrameworkDeploymentAdapter{}

func (a *FrameworkDeploymentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CloudSecurityComplianceFrameworkDeployment", "name", a.id)

	req := &pb.GetFrameworkDeploymentRequest{Name: a.id.String()}
	deployment, err := a.gcpClient.GetFrameworkDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CloudSecurityComplianceFrameworkDeployment %q: %w", a.id, err)
	}

	a.actual = deployment
	return true, nil
}

func (a *FrameworkDeploymentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CloudSecurityComplianceFrameworkDeployment", "name", a.id)

	req := &pb.CreateFrameworkDeploymentRequest{
		Parent:                a.id.ParentString(),
		FrameworkDeployment:   a.desired,
		FrameworkDeploymentId: a.id.FrameworkDeployment,
	}
	op, err := a.gcpClient.CreateFrameworkDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CloudSecurityComplianceFrameworkDeployment %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for CloudSecurityComplianceFrameworkDeployment creation %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created CloudSecurityComplianceFrameworkDeployment", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *FrameworkDeploymentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CloudSecurityComplianceFrameworkDeployment", "name", a.id)

	diffs, err := compareFrameworkDeployment(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	return fmt.Errorf("CloudSecurityComplianceFrameworkDeployment is immutable and cannot be updated")
}

func compareFrameworkDeployment(ctx context.Context, actual, desired *pb.FrameworkDeployment) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CloudSecurityComplianceFrameworkDeploymentSpec_FromProto, CloudSecurityComplianceFrameworkDeploymentSpec_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.FrameworkDeployment)

	populateDefaults := func(obj *pb.FrameworkDeployment) {
		// Populate any GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, _, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func (a *FrameworkDeploymentAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.FrameworkDeployment) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CloudSecurityComplianceFrameworkDeploymentStatus{}
	status.ObservedState = CloudSecurityComplianceFrameworkDeploymentObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *FrameworkDeploymentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CloudSecurityComplianceFrameworkDeployment", "name", a.id)

	req := &pb.DeleteFrameworkDeploymentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteFrameworkDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CloudSecurityComplianceFrameworkDeployment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CloudSecurityComplianceFrameworkDeployment %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for CloudSecurityComplianceFrameworkDeployment deletion %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted CloudSecurityComplianceFrameworkDeployment", "name", a.id)
	return true, nil
}

func (a *FrameworkDeploymentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudSecurityComplianceFrameworkDeployment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudSecurityComplianceFrameworkDeploymentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.FrameworkDeployment)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.FrameworkDeployment)
	u.SetGroupVersionKind(krm.CloudSecurityComplianceFrameworkDeploymentGVK)

	u.Object = uObj
	return u, nil
}
