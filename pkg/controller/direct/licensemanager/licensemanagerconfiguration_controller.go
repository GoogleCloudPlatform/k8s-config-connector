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

package licensemanager

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/licensemanager/apiv1"
	pb "cloud.google.com/go/licensemanager/apiv1/licensemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/licensemanager/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.LicenseManagerConfigurationGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building LicenseManager client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LicenseManagerConfiguration{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := LicenseManagerConfigurationSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &LicenseManagerConfigurationAdapter{
		id:        identity.(*krm.LicenseManagerConfigurationIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.LicenseManagerConfigurationIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &LicenseManagerConfigurationAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type LicenseManagerConfigurationAdapter struct {
	id        *krm.LicenseManagerConfigurationIdentity
	gcpClient *gcp.Client
	desired   *pb.Configuration
	actual    *pb.Configuration
}

var _ directbase.Adapter = &LicenseManagerConfigurationAdapter{}

func (a *LicenseManagerConfigurationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LicenseManagerConfiguration", "name", a.id.String())

	req := &pb.GetConfigurationRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetConfiguration(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LicenseManagerConfiguration %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *LicenseManagerConfigurationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating LicenseManagerConfiguration", "id", fqn)

	parent := a.id.ParentString()

	req := &pb.CreateConfigurationRequest{
		Parent:          parent,
		Configuration:   a.desired,
		ConfigurationId: a.id.Configuration,
	}
	op, err := a.gcpClient.CreateConfiguration(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LicenseManagerConfiguration %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting LicenseManagerConfiguration %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created LicenseManagerConfiguration", "name", a.id.String())

	// Fetch fully-populated resource immediately after LRO completion (per rule: Reconciling Empty or Incomplete LRO Responses)
	getReq := &pb.GetConfigurationRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetConfiguration(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching LicenseManagerConfiguration %s after creation: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *LicenseManagerConfigurationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LicenseManagerConfiguration", "name", a.id.String())

	diffs, updateMask, err := compareLicenseManagerConfiguration(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		clonedDesired := proto.Clone(a.desired).(*pb.Configuration)
		clonedDesired.Name = a.id.String()

		req := &pb.UpdateConfigurationRequest{
			Configuration: clonedDesired,
			UpdateMask:    updateMask,
		}

		op, err := a.gcpClient.UpdateConfiguration(ctx, req)
		if err != nil {
			return fmt.Errorf("updating LicenseManagerConfiguration %s: %w", a.id.String(), err)
		}

		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting LicenseManagerConfiguration %s update: %w", a.id.String(), err)
		}

		// Fetch fully-populated resource immediately after LRO completion (per rule: Reconciling Empty or Incomplete LRO Responses)
		getReq := &pb.GetConfigurationRequest{Name: a.id.String()}
		latest, err = a.gcpClient.GetConfiguration(ctx, getReq)
		if err != nil {
			return fmt.Errorf("fetching LicenseManagerConfiguration %s after update: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *LicenseManagerConfigurationAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Configuration) error {
	mapCtx := &direct.MapContext{}
	status := &krm.LicenseManagerConfigurationStatus{}
	status.ObservedState = LicenseManagerConfigurationObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	identity := a.id.String()
	status.ExternalRef = &identity

	return op.UpdateStatus(ctx, status, nil)
}

func (a *LicenseManagerConfigurationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LicenseManagerConfiguration{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LicenseManagerConfigurationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.Configuration)
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Configuration)
	u.SetGroupVersionKind(krm.LicenseManagerConfigurationGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *LicenseManagerConfigurationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LicenseManagerConfiguration", "name", a.id.String())

	req := &pb.DeleteConfigurationRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteConfiguration(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LicenseManagerConfiguration, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting LicenseManagerConfiguration %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted LicenseManagerConfiguration", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete LicenseManagerConfiguration %s: %w", a.id.String(), err)
	}
	return true, nil
}

func compareLicenseManagerConfiguration(ctx context.Context, actual, desired *pb.Configuration) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, LicenseManagerConfigurationSpec_FromProto, LicenseManagerConfigurationSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
