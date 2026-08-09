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

package developerconnect

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/developerconnect/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/developerconnect/apiv1"
	pb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DevConnectAccountConnectorGVK, NewAccountConnectorModel)
}

func NewAccountConnectorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDevConnectAccountConnector{config: *config}, nil
}

var _ directbase.Model = &modelDevConnectAccountConnector{}

type modelDevConnectAccountConnector struct {
	config config.ControllerConfig
}

func (m *modelDevConnectAccountConnector) client(ctx context.Context, location string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	// Add regional endpoint if region is specified
	if location != "" && location != "global" {
		endpoint := fmt.Sprintf("developerconnect.%s.rep.googleapis.com:443", location)
		opts = append(opts, option.WithEndpoint(endpoint))
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Developer Connect client: %w", err)
	}
	return gcpClient, err
}

func (m *modelDevConnectAccountConnector) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DevConnectAccountConnector{}
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

	location := obj.Spec.Location
	if location == nil || *location == "" {
		return nil, fmt.Errorf("location is required on DevConnectAccountConnector")
	}

	gcpClient, err := m.client(ctx, *location)
	if err != nil {
		return nil, err
	}

	return &DevConnectAccountConnectorAdapter{
		id:        id.(*krm.DevConnectAccountConnectorIdentity),
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelDevConnectAccountConnector) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.DevConnectAccountConnectorIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &DevConnectAccountConnectorAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type DevConnectAccountConnectorAdapter struct {
	id        *krm.DevConnectAccountConnectorIdentity
	gcpClient *gcp.Client
	reader    client.Reader
	desired   *krm.DevConnectAccountConnector
	actual    *pb.AccountConnector
}

var _ directbase.Adapter = &DevConnectAccountConnectorAdapter{}

func (a *DevConnectAccountConnectorAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DevConnectAccountConnector", "name", a.id.String())

	req := &pb.GetAccountConnectorRequest{Name: a.id.String()}
	accountConnector, err := a.gcpClient.GetAccountConnector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DevConnectAccountConnector %q from GCP: %w", a.id.String(), err)
	}

	a.actual = accountConnector
	return true, nil
}

func (a *DevConnectAccountConnectorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DevConnectAccountConnector", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DevConnectAccountConnectorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Merge labels in from the metadata labels
	resource.Labels = label.NewGCPLabelsFromK8sLabels(a.desired.GetObjectMeta().GetLabels())
	// Keep specified labels
	for k, v := range desired.Spec.Labels {
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}
		resource.Labels[k] = v
	}

	req := &pb.CreateAccountConnectorRequest{
		Parent:             a.id.ParentString(),
		AccountConnector:   resource,
		AccountConnectorId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateAccountConnector(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DevConnectAccountConnector %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DevConnectAccountConnector %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created DevConnectAccountConnector", "name", a.id.String())

	// Fetch fully populated resource because create response might be incomplete
	reqGet := &pb.GetAccountConnectorRequest{Name: a.id.String()}
	created, err = a.gcpClient.GetAccountConnector(ctx, reqGet)
	if err != nil {
		return fmt.Errorf("getting fully-populated DevConnectAccountConnector %s after creation: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *DevConnectAccountConnectorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DevConnectAccountConnector", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DevConnectAccountConnectorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Merge labels in from the metadata labels
	resource.Labels = label.NewGCPLabelsFromK8sLabels(a.desired.GetObjectMeta().GetLabels())
	for k, v := range desired.Spec.Labels {
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}
		resource.Labels[k] = v
	}

	resource.Name = a.id.String()

	diffs, updateMask, err := compareAccountConnectorResource(ctx, a.actual, resource)
	if err != nil {
		return err
	}

	updated := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
	} else {
		log.V(2).Info("fields need update", "name", a.id.String(), "paths", updateMask.Paths)
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateAccountConnectorRequest{
			UpdateMask:       updateMask,
			AccountConnector: resource,
		}
		op, err := a.gcpClient.UpdateAccountConnector(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DevConnectAccountConnector %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DevConnectAccountConnector %s waiting update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated DevConnectAccountConnector", "name", a.id.String())

		// Fetch fully populated resource because update response might be incomplete
		reqGet := &pb.GetAccountConnectorRequest{Name: a.id.String()}
		updated, err = a.gcpClient.GetAccountConnector(ctx, reqGet)
		if err != nil {
			return fmt.Errorf("getting fully-populated DevConnectAccountConnector %s after update: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func compareAccountConnectorResource(ctx context.Context, actual, desired *pb.AccountConnector) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := DevConnectAccountConnectorSpec_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}

	maskedActual := DevConnectAccountConnectorSpec_ToProto(mapCtx, spec)
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.AccountConnector)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *DevConnectAccountConnectorAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AccountConnector) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DevConnectAccountConnectorStatus{}
	status.ObservedState = DevConnectAccountConnectorObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *DevConnectAccountConnectorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DevConnectAccountConnector{}
	mapCtx := &direct.MapContext{}
	spec := DevConnectAccountConnectorSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if spec == nil {
		spec = &krm.DevConnectAccountConnectorSpec{}
	}
	obj.Spec = *spec
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Account_connector)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Account_connector)
	u.SetGroupVersionKind(krm.DevConnectAccountConnectorGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *DevConnectAccountConnectorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DevConnectAccountConnector", "name", a.id.String())

	req := &pb.DeleteAccountConnectorRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteAccountConnector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DevConnectAccountConnector, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting DevConnectAccountConnector %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete DevConnectAccountConnector %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted DevConnectAccountConnector", "name", a.id.String())
	return true, nil
}
