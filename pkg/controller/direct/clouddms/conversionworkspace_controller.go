// Copyright 2024 Google LLC
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
// proto.service: google.cloud.clouddms.v1.DataMigrationService
// proto.message: google.cloud.clouddms.v1.ConversionWorkspace
// crd.type: CloudDMSConversionWorkspace
// crd.version: v1alpha1

package clouddms

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/clouddms/apiv1"
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CloudDMSConversionWorkspaceGVK, NewConversionWorkspaceModel)
}

func NewConversionWorkspaceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConversionWorkspace{config: config}, nil
}

var _ directbase.Model = &modelConversionWorkspace{}

type modelConversionWorkspace struct {
	config *config.ControllerConfig
}

func (m *modelConversionWorkspace) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.CloudDMSConversionWorkspace{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// conversionWorkspaceId is server-generated.
	id, err := krm.NewConversionWorkspaceIdentity(ctx, reader, obj, directbase.WithAllowEmptyID(true))
	if err != nil {
		return nil, err
	}

	// Get clouddms GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	clouddmsClient, err := gcpClient.newDataMigrationServiceClient(ctx)
	if err != nil {
		return nil, err
	}
	return &conversionWorkspaceAdapter{
		gcpClient: clouddmsClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelConversionWorkspace) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type conversionWorkspaceAdapter struct {
	gcpClient *gcp.DataMigrationClient
	id        *krm.ConversionWorkspaceIdentity
	desired   *krm.CloudDMSConversionWorkspace
	actual    *pb.ConversionWorkspace
	reader    client.Reader
}

var _ directbase.Adapter = &conversionWorkspaceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *conversionWorkspaceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ConversionWorkspace", "name", a.id)

	// Handle server-generated ID: if the resource ID is empty, it means we intend to create.
	if a.id.ID() == "" {
		log.V(2).Info("conversion workspace id is empty, assuming resource does not exist", "name", a.id)
		return false, nil
	}

	req := &pb.GetConversionWorkspaceRequest{Name: a.id.String()}
	conversionworkspacepb, err := a.gcpClient.GetConversionWorkspace(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// If the resource ID was specified but the resource is not found, it's an error.
			// The user might be trying to acquire a resource that doesn't exist.
			return false, fmt.Errorf("conversion workspace %q not found, but was specified in spec.resourceID; remove spec.resourceID to create a new one: %w", a.id.String(), err)
		}
		return false, fmt.Errorf("getting ConversionWorkspace %q: %w", a.id, err)
	}

	a.actual = conversionworkspacepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *conversionWorkspaceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ConversionWorkspace", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudDMSConversionWorkspaceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// The API requires a server-generated ID. We use the k8s object name as the ID hint.
	conversionWorkspaceID := a.id.KRMName()

	req := &pb.CreateConversionWorkspaceRequest{
		Parent:                a.id.Parent().String(),
		ConversionWorkspaceId: conversionWorkspaceID,
		ConversionWorkspace:   resource,
	}
	op, err := a.gcpClient.CreateConversionWorkspace(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ConversionWorkspace %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ConversionWorkspace %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ConversionWorkspace", "name", a.id)

	// Update the resource ID with the server-generated value
	_, actualResourceID, err := krm.ParseConversionWorkspaceExternal(created.Name)
	if err != nil {
		return fmt.Errorf("parsing the resource name %q in the response of CreateConversionWorkspace: %w", created.Name, err)
	}
	a.id.SetID(actualResourceID) // Update the adapter's ID

	status := &krm.CloudDMSConversionWorkspaceStatus{}
	status.ObservedState = CloudDMSConversionWorkspaceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String()) // Use the updated ID string
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *conversionWorkspaceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ConversionWorkspace", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudDMSConversionWorkspaceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String() // Set the name for the update request

	paths := []string{}
	if desired.Spec.DisplayName != nil && !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		paths = append(paths, "display_name")
	}
	if desired.Spec.GlobalSettings != nil && !reflect.DeepEqual(resource.GlobalSettings, a.actual.GlobalSettings) {
		paths = append(paths, "global_settings")
	}
	// source and destination are immutable

	var updated *pb.ConversionWorkspace
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// Even though there is no update, we still want to update KRM status (e.g., for acquisition)
		updated = a.actual
	} else {
		req := &pb.UpdateConversionWorkspaceRequest{
			ConversionWorkspace: resource,
			UpdateMask:          &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateConversionWorkspace(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ConversionWorkspace %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("ConversionWorkspace %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ConversionWorkspace", "name", a.id)
	}

	status := &krm.CloudDMSConversionWorkspaceStatus{}
	status.ObservedState = CloudDMSConversionWorkspaceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *conversionWorkspaceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudDMSConversionWorkspace{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudDMSConversionWorkspaceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID()) // Export the server-generated ID

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.KRMName()) // Set the KRM name
	u.SetGroupVersionKind(krm.CloudDMSConversionWorkspaceGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *conversionWorkspaceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ConversionWorkspace", "name", a.id)

	// Handle server-generated ID: if the resource ID is empty, it likely means Find returned false,
	// and we shouldn't attempt deletion. This scenario should ideally be prevented by the controller flow.
	if a.id.ID() == "" {
		log.V(2).Info("skipping delete for ConversionWorkspace with empty resource ID, resource likely doesn't exist", "name", a.id.KRMName())
		return true, nil // Assume already deleted or never created
	}

	req := &pb.DeleteConversionWorkspaceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteConversionWorkspace(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ConversionWorkspace, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ConversionWorkspace %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ConversionWorkspace", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ConversionWorkspace %s: %w", a.id, err)
	}
	return true, nil
}
