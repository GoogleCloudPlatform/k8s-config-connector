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

package configdeployment

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/config/apiv1"
	pb "cloud.google.com/go/config/apiv1/configpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdeployment/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ConfigDeploymentGroupGVK, NewConfigDeploymentGroupModel)
}

func NewConfigDeploymentGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConfigDeploymentGroup{config: *config}, nil
}

var _ directbase.Model = &modelConfigDeploymentGroup{}

type modelConfigDeploymentGroup struct {
	config config.ControllerConfig
}

func (m *modelConfigDeploymentGroup) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Config client: %w", err)
	}
	return gcpClient, err
}

func (a *ConfigDeploymentGroupAdapter) normalizeReferences(ctx context.Context) error {
	for i := range a.desired.Spec.DeploymentUnits {
		unit := &a.desired.Spec.DeploymentUnits[i]
		if unit.DeploymentRef != nil {
			if err := unit.DeploymentRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
				return fmt.Errorf("normalizing DeploymentRef: %w", err)
			}
		}
	}
	return nil
}

func (m *modelConfigDeploymentGroup) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ConfigDeploymentGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := identity.(*krm.ConfigDeploymentGroupIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", identity)
	}

	// Get config GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConfigDeploymentGroupAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelConfigDeploymentGroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ConfigDeploymentGroupAdapter struct {
	id        *krm.ConfigDeploymentGroupIdentity
	gcpClient *gcp.Client
	desired   *krm.ConfigDeploymentGroup
	actual    *pb.DeploymentGroup
	reader    client.Reader
}

var _ directbase.Adapter = &ConfigDeploymentGroupAdapter{}

// Find retrieves the GCP resource.
func (a *ConfigDeploymentGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DeploymentGroup", "name", a.id)

	req := &pb.GetDeploymentGroupRequest{Name: a.id.String()}
	res, err := a.gcpClient.GetDeploymentGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DeploymentGroup %q: %w", a.id, err)
	}

	a.actual = res
	return true, nil
}

// Create creates the resource in GCP and updates the Config Connector status based on the GCP response.
func (a *ConfigDeploymentGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DeploymentGroup", "name", a.id)

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	resource := ConfigDeploymentGroupSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	req := &pb.CreateDeploymentGroupRequest{
		Parent:            a.id.ParentString(),
		DeploymentGroup:   resource,
		DeploymentGroupId: a.id.DeploymentGroup,
	}
	op, err := a.gcpClient.CreateDeploymentGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DeploymentGroup %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DeploymentGroup %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DeploymentGroup", "name", a.id)

	status := &krm.ConfigDeploymentGroupStatus{}
	status.ObservedState = ConfigDeploymentGroupObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func deploymentUnitsEqual(a, b []*pb.DeploymentUnit) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !proto.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Update updates the resource in GCP and updates the Config Connector status based on the GCP response.
func (a *ConfigDeploymentGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DeploymentGroup", "name", a.id)

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := ConfigDeploymentGroupSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	paths := []string{}
	// For DeploymentGroup, labels, annotations and deployment_units are mutable.
	if !reflect.DeepEqual(desiredPb.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, desiredPb.Labels)
		paths = append(paths, "labels")
	}
	if !reflect.DeepEqual(desiredPb.Annotations, a.actual.Annotations) {
		report.AddField("annotations", a.actual.Annotations, desiredPb.Annotations)
		paths = append(paths, "annotations")
	}
	if !deploymentUnitsEqual(desiredPb.DeploymentUnits, a.actual.DeploymentUnits) {
		report.AddField("deploymentUnits", a.actual.DeploymentUnits, desiredPb.DeploymentUnits)
		paths = append(paths, "deployment_units")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateDeploymentGroupRequest{
		DeploymentGroup: desiredPb,
		UpdateMask:      &fieldmaskpb.FieldMask{Paths: paths},
	}
	desiredPb.Name = a.id.String()

	op, err := a.gcpClient.UpdateDeploymentGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DeploymentGroup %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DeploymentGroup %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated DeploymentGroup", "name", a.id)

	status := &krm.ConfigDeploymentGroupStatus{}
	status.ObservedState = ConfigDeploymentGroupObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ConfigDeploymentGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ConfigDeploymentGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ConfigDeploymentGroupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// Set projectRef and Location from identity
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.DeploymentGroup)
	u.SetGroupVersionKind(krm.ConfigDeploymentGroupGVK)

	u.Object = uObj
	return u, nil
}

// Delete deletes the resource in GCP.
func (a *ConfigDeploymentGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DeploymentGroup", "name", a.id)

	req := &pb.DeleteDeploymentGroupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDeploymentGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DeploymentGroup %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("DeploymentGroup %s waiting deletion: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DeploymentGroup", "name", a.id)
	return true, nil
}
