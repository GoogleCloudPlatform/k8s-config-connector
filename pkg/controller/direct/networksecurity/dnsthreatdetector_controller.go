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

// +tool:controller
// proto.service: google.cloud.networksecurity.v1beta1.DnsThreatDetectorService
// proto.message: google.cloud.networksecurity.v1beta1.DnsThreatDetector
// crd.type: NetworkSecurityDnsThreatDetector
// crd.version: v1alpha1

package networksecurity

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/networksecurity/apiv1beta1"
	pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityDnsThreatDetectorGVK, NewDnsThreatDetectorModel)
}

func NewDnsThreatDetectorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dnsthreatdetectorModel{config: *config}, nil
}

var _ directbase.Model = &dnsthreatdetectorModel{}

type dnsthreatdetectorModel struct {
	config config.ControllerConfig
}

func (m *dnsthreatdetectorModel) client(ctx context.Context) (*gcp.DnsThreatDetectorClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDnsThreatDetectorRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DnsThreatDetector client: %w", err)
	}
	return gcpClient, err
}

func (m *dnsthreatdetectorModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityDnsThreatDetector{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	typedID := id.(*krm.NetworkSecurityDnsThreatDetectorIdentity)

	for i := range obj.Spec.ExcludedNetworks {
		if err := obj.Spec.ExcludedNetworks[i].Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, err
		}
	}

	// Get gcp client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &dnsthreatdetectorAdapter{
		id:        typedID,
		endpoint:  obj.Spec.Location,
		projectID: typedID.Project,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *dnsthreatdetectorModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type dnsthreatdetectorAdapter struct {
	id        *krm.NetworkSecurityDnsThreatDetectorIdentity
	endpoint  *string
	projectID string
	gcpClient *gcp.DnsThreatDetectorClient
	desired   *krm.NetworkSecurityDnsThreatDetector
	actual    *pb.DnsThreatDetector
}

var _ directbase.Adapter = &dnsthreatdetectorAdapter{}

func (a *dnsthreatdetectorAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DnsThreatDetector", "name", a.id.String())

	req := &pb.GetDnsThreatDetectorRequest{
		Name: a.id.String(),
	}
	existing, err := a.gcpClient.GetDnsThreatDetector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DnsThreatDetector %q: %w", a.id.String(), err)
	}

	a.actual = existing
	return true, nil
}

func (a *dnsthreatdetectorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DnsThreatDetector", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkSecurityDnsThreatDetectorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateDnsThreatDetectorRequest{
		Parent:              fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		DnsThreatDetectorId: a.id.DnsThreatDetector,
		DnsThreatDetector:   resource,
	}
	created, err := a.gcpClient.CreateDnsThreatDetector(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DnsThreatDetector %q: %w", a.id.String(), err)
	}

	status := &krm.NetworkSecurityDnsThreatDetectorStatus{}
	status.ObservedState = NetworkSecurityDnsThreatDetectorObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *dnsthreatdetectorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DnsThreatDetector", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkSecurityDnsThreatDetectorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.actual.Labels, resource.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	if !reflect.DeepEqual(a.actual.ExcludedNetworks, resource.ExcludedNetworks) {
		updateMask.Paths = append(updateMask.Paths, "excluded_networks")
	}
	if !reflect.DeepEqual(a.actual.Provider, resource.Provider) {
		updateMask.Paths = append(updateMask.Paths, "provider")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no fields to update for DnsThreatDetector", "name", a.id.String())
		return nil
	}

	req := &pb.UpdateDnsThreatDetectorRequest{
		UpdateMask:        updateMask,
		DnsThreatDetector: resource,
	}
	updated, err := a.gcpClient.UpdateDnsThreatDetector(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DnsThreatDetector %q: %w", a.id.String(), err)
	}

	status := &krm.NetworkSecurityDnsThreatDetectorStatus{}
	status.ObservedState = NetworkSecurityDnsThreatDetectorObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *dnsthreatdetectorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityDnsThreatDetector{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *NetworkSecurityDnsThreatDetectorSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	typedID := &krm.NetworkSecurityDnsThreatDetectorIdentity{}
	if err := typedID.FromExternal(a.actual.Name); err != nil {
		return nil, err
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: typedID.Project}
	obj.Spec.Location = direct.LazyPtr(typedID.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	return u, nil
}

func (a *dnsthreatdetectorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DnsThreatDetector", "name", a.id.String())

	req := &pb.DeleteDnsThreatDetectorRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteDnsThreatDetector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DnsThreatDetector %q: %w", a.id.String(), err)
	}

	return true, nil
}
