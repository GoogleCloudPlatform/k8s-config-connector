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

package servicedirectory

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/servicedirectory/apiv1beta1"
	pb "cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ServiceDirectoryEndpointGVK, NewEndpointModel)
}

func NewEndpointModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &endpointModel{config: config}, nil
}

var _ directbase.Model = &endpointModel{}

type endpointModel struct {
	config *config.ControllerConfig
}

func (m *endpointModel) client(ctx context.Context) (*gcp.RegistrationClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRegistrationRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ServiceDirectory registration client: %w", err)
	}
	return gcpClient, nil
}

func (m *endpointModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ServiceDirectoryEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve ComputeAddressRef to its actual IP address
	if obj.Spec.AddressRef != nil {
		ip, err := resolveComputeAddressIP(ctx, reader, obj, obj.Spec.AddressRef)
		if err != nil {
			return nil, err
		}
		obj.Spec.AddressRef = &krmcomputev1beta1.ComputeAddressRef{
			External: ip,
		}
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.ServiceDirectoryEndpointIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := ServiceDirectoryEndpointSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Support labels on Metadata
	desired.Metadata = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ServiceDirectoryEndpointAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *endpointModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ServiceDirectoryEndpointAdapter struct {
	id        *krm.ServiceDirectoryEndpointIdentity
	gcpClient *gcp.RegistrationClient
	desired   *pb.Endpoint
	actual    *pb.Endpoint
}

var _ directbase.Adapter = &ServiceDirectoryEndpointAdapter{}

func (a *ServiceDirectoryEndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting ServiceDirectoryEndpoint", "name", fqn)

	req := &pb.GetEndpointRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetEndpoint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ServiceDirectoryEndpoint %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *ServiceDirectoryEndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating ServiceDirectoryEndpoint", "name", fqn)

	req := &pb.CreateEndpointRequest{
		Parent:     parent,
		EndpointId: a.id.Endpoint,
		Endpoint:   a.desired,
	}
	created, err := a.gcpClient.CreateEndpoint(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ServiceDirectoryEndpoint %s: %w", a.id.Endpoint, err)
	}
	log.V(2).Info("successfully created ServiceDirectoryEndpoint", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ServiceDirectoryEndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating ServiceDirectoryEndpoint", "name", fqn)

	diffs, updateMask, err := compareEndpoint(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.Endpoint)
		desired.Name = fqn

		req := &pb.UpdateEndpointRequest{
			Endpoint:   desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateEndpoint(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ServiceDirectoryEndpoint %s: %w", fqn, err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ServiceDirectoryEndpointAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Endpoint) error {
	mapCtx := &direct.MapContext{}
	status := ServiceDirectoryEndpointStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ServiceDirectoryEndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ServiceDirectoryEndpoint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ServiceDirectoryEndpointSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ServiceDirectoryEndpointGVK)
	return u, nil
}

func (a *ServiceDirectoryEndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting ServiceDirectoryEndpoint", "name", fqn)

	req := &pb.DeleteEndpointRequest{Name: fqn}
	err := a.gcpClient.DeleteEndpoint(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ServiceDirectoryEndpoint, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting ServiceDirectoryEndpoint %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted ServiceDirectoryEndpoint", "name", fqn)
	return true, nil
}

func compareEndpoint(ctx context.Context, actual, desired *pb.Endpoint) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ServiceDirectoryEndpointSpec_FromProto, ServiceDirectoryEndpointSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Metadata = actual.Metadata
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func resolveComputeAddressIP(ctx context.Context, reader client.Reader, src client.Object, ref *krmcomputev1beta1.ComputeAddressRef) (string, error) {
	if ref == nil {
		return "", nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return "", fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref.External, nil
	}

	if ref.Name == "" {
		return "", fmt.Errorf("must specify either name or external on reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	gvk := schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeAddress",
	}

	computeAddress := &unstructured.Unstructured{}
	computeAddress.SetGroupVersionKind(gvk)
	if err := reader.Get(ctx, key, computeAddress); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(gvk, key)
		}
		return "", fmt.Errorf("error reading referenced %v %v: %w", gvk.Kind, key, err)
	}

	// Because `spec.address` field is optional, we can't guarantee it always
	// exists in a successfully reconciled ComputeAddress CR, so we should use
	// the `status.address` or `status.observedState.address` instead.
	address, _, err := unstructured.NestedString(computeAddress.Object, "status", "address")
	if err != nil || address == "" {
		address, _, err = unstructured.NestedString(computeAddress.Object, "status", "observedState", "address")
		if err != nil || address == "" {
			return "", fmt.Errorf("cannot get address for referenced %s %v (status.address and status.observedState.address are empty)", computeAddress.GetKind(), computeAddress.GetNamespace())
		}
	}
	return address, nil
}
