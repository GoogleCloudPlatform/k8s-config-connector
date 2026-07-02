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

package apphub

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gcp "cloud.google.com/go/apphub/apiv1"
	apphubpb "cloud.google.com/go/apphub/apiv1/apphubpb"
)

func init() {
	registry.RegisterModel(krm.AppHubServiceProjectAttachmentGVK, NewServiceProjectAttachmentModel)
}

func NewServiceProjectAttachmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &serviceProjectAttachmentModel{config: config}, nil
}

var _ directbase.Model = &serviceProjectAttachmentModel{}

type serviceProjectAttachmentModel struct {
	config *config.ControllerConfig
}

func (m *serviceProjectAttachmentModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ServiceProjectAttachment client: %w", err)
	}
	return gcpClient, err
}

func (m *serviceProjectAttachmentModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AppHubServiceProjectAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.AppHubServiceProjectAttachmentIdentity)

	// Get apphub client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServiceProjectAttachmentAdapter{
		id:        id,
		inner:     obj,
		gcpClient: gcpClient,
		reader:    reader,
	}, nil
}

func (m *serviceProjectAttachmentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support discovery
	return nil, nil
}

type ServiceProjectAttachmentAdapter struct {
	id        *krm.AppHubServiceProjectAttachmentIdentity
	inner     *krm.AppHubServiceProjectAttachment
	gcpClient *gcp.Client
	reader    client.Reader
}

var _ directbase.Adapter = &ServiceProjectAttachmentAdapter{}

// Find retrieves the GCP resource.
func (a *ServiceProjectAttachmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ServiceProjectAttachment", "name", a.id.String())

	req := &apphubpb.GetServiceProjectAttachmentRequest{
		Name: a.id.String(),
	}
	apiObj, err := a.gcpClient.GetServiceProjectAttachment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ServiceProjectAttachment %q: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	a.inner.Status.ObservedState = AppHubServiceProjectAttachmentObservedState_v1alpha1_FromProto(mapCtx, apiObj)
	if mapCtx.Err() != nil {
		return true, fmt.Errorf("mapping to AppHubServiceProjectAttachmentObservedState: %w", mapCtx.Err())
	}

	return true, nil
}

// Create creates the GCP resource.
func (a *ServiceProjectAttachmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ServiceProjectAttachment", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	apiObj := AppHubServiceProjectAttachmentSpec_v1alpha1_ToProto(mapCtx, &a.inner.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping to ServiceProjectAttachment: %w", mapCtx.Err())
	}

	projectRef, err := refs.ResolveProject(ctx, a.reader, a.inner.Namespace, a.inner.Spec.ServiceProjectRef)
	if err != nil {
		return err
	}
	apiObj.ServiceProject = "projects/" + projectRef.ProjectID

	parentStr := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	req := &apphubpb.CreateServiceProjectAttachmentRequest{
		Parent:                     parentStr,
		ServiceProjectAttachmentId: a.id.ServiceProjectAttachment,
		ServiceProjectAttachment:   apiObj,
	}

	op, err := a.gcpClient.CreateServiceProjectAttachment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ServiceProjectAttachment %q: %w", a.id.String(), err)
	}

	apiObj, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for ServiceProjectAttachment %q to be created: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created ServiceProjectAttachment", "name", a.id.String())

	status := &krm.AppHubServiceProjectAttachmentStatus{}
	status.ObservedState = AppHubServiceProjectAttachmentObservedState_v1alpha1_FromProto(mapCtx, apiObj)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping to AppHubServiceProjectAttachmentObservedState: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(apiObj.GetName())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the GCP resource.
func (a *ServiceProjectAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ServiceProjectAttachment", "name", a.id.String())
	status := &krm.AppHubServiceProjectAttachmentStatus{}
	status.ObservedState = a.inner.Status.ObservedState
	if a.inner.Status.ExternalRef != nil {
		status.ExternalRef = a.inner.Status.ExternalRef
	} else {
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export returns the KRM representation.
func (a *ServiceProjectAttachmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.inner == nil {
		return nil, fmt.Errorf("inner is nil")
	}
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(a.inner)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: u}, nil
}

// Delete deletes the GCP resource.
func (a *ServiceProjectAttachmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ServiceProjectAttachment", "name", a.id.String())

	req := &apphubpb.DeleteServiceProjectAttachmentRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteServiceProjectAttachment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ServiceProjectAttachment %q: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for ServiceProjectAttachment %q to be deleted: %w", a.id.String(), err)
	}

	return true, nil
}
