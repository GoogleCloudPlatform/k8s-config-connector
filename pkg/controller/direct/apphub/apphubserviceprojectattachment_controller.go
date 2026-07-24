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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
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

	// Always call common.NormalizeReferences to resolve any resource references:
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
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

	mapCtx := &direct.MapContext{}
	desired := AppHubServiceProjectAttachmentSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &ServiceProjectAttachmentAdapter{
		id:        id,
		inner:     obj,
		gcpClient: gcpClient,
		reader:    reader,
		desired:   desired,
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

	desired *apphubpb.ServiceProjectAttachment
	actual  *apphubpb.ServiceProjectAttachment
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

	a.actual = apiObj

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

	projectRef, err := refs.ResolveProject(ctx, a.reader, a.inner.Namespace, a.inner.Spec.ServiceProjectRef)
	if err != nil {
		return err
	}
	a.desired.ServiceProject = "projects/" + projectRef.ProjectID

	parentStr := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	req := &apphubpb.CreateServiceProjectAttachmentRequest{
		Parent:                     parentStr,
		ServiceProjectAttachmentId: a.id.ServiceProjectAttachment,
		ServiceProjectAttachment:   a.desired,
	}

	op, err := a.gcpClient.CreateServiceProjectAttachment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ServiceProjectAttachment %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for ServiceProjectAttachment %q to be created: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created ServiceProjectAttachment", "name", a.id.String())

	// Fetch the fully populated resource to be safe and update status
	getReq := &apphubpb.GetServiceProjectAttachmentRequest{
		Name: a.id.String(),
	}
	latest, err := a.gcpClient.GetServiceProjectAttachment(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting created ServiceProjectAttachment %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

// Update updates the GCP resource.
func (a *ServiceProjectAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ServiceProjectAttachment", "name", a.id.String())

	projectRef, err := refs.ResolveProject(ctx, a.reader, a.inner.Namespace, a.inner.Spec.ServiceProjectRef)
	if err != nil {
		return err
	}
	a.desired.ServiceProject = "projects/" + projectRef.ProjectID

	diff, err := compareServiceProjectAttachment(ctx, a.actual, a.desired, a.reader, a.inner.Namespace)
	if err != nil {
		return err
	}

	if diff.HasDiff() {
		diff.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diff)
		return fmt.Errorf("AppHubServiceProjectAttachment is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

// Export returns the KRM representation.
func (a *ServiceProjectAttachmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AppHubServiceProjectAttachment{}
	mapCtx := &direct.MapContext{}
	spec := AppHubServiceProjectAttachmentSpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if spec != nil {
		obj.Spec = *spec
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ProjectRef = &refs.ProjectRef{External: fmt.Sprintf("projects/%s", a.id.Project)}
	obj.Spec.ResourceID = &a.id.ServiceProjectAttachment

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
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

func (a *ServiceProjectAttachmentAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *apphubpb.ServiceProjectAttachment) error {
	mapCtx := &direct.MapContext{}
	status := &krm.AppHubServiceProjectAttachmentStatus{}
	status.ObservedState = AppHubServiceProjectAttachmentObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}

func compareServiceProjectAttachment(ctx context.Context, actual, desired *apphubpb.ServiceProjectAttachment, reader client.Reader, namespace string) (*structuredreporting.Diff, error) {
	if actual.GetServiceProject() == desired.GetServiceProject() {
		return &structuredreporting.Diff{}, nil
	}

	// Since actual.ServiceProject might be projects/<number> and desired.ServiceProject might be projects/<id> (or vice versa),
	// we attempt to resolve both to see if they refer to the same project.
	desiredRef := &refs.ProjectRef{External: desired.GetServiceProject()}
	actualRef := &refs.ProjectRef{External: actual.GetServiceProject()}

	desiredProj, err := refs.ResolveProject(ctx, reader, namespace, desiredRef)
	if err != nil {
		return nil, fmt.Errorf("resolving desired project: %w", err)
	}
	actualProj, err := refs.ResolveProject(ctx, reader, namespace, actualRef)
	if err != nil {
		return nil, fmt.Errorf("resolving actual project: %w", err)
	}

	if desiredProj != nil && actualProj != nil && desiredProj.ProjectID == actualProj.ProjectID {
		return &structuredreporting.Diff{}, nil
	}

	diff := &structuredreporting.Diff{}
	diff.AddField("spec.serviceProjectRef", actual.GetServiceProject(), desired.GetServiceProject())
	return diff, nil
}
