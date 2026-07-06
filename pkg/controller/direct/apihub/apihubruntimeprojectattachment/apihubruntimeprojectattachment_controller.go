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

package apihubruntimeprojectattachment

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/apihub/apiv1"
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

const (
	ctrlName = "apihub-runtimeprojectattachment-controller"
)

func init() {
	registry.RegisterModel(krm.APIHubRuntimeProjectAttachmentGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.RuntimeProjectAttachmentClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRuntimeProjectAttachmentRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building APIHubRuntimeProjectAttachment client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.APIHubRuntimeProjectAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.APIHubRuntimeProjectAttachmentIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// NormalizeReferences is called to resolve any resource references:
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := apihub.APIHubRuntimeProjectAttachmentSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		reader:    reader,
		namespace: obj.GetNamespace(),
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.APIHubRuntimeProjectAttachmentIdentity
	gcpClient *gcp.RuntimeProjectAttachmentClient
	desired   *pb.RuntimeProjectAttachment
	actual    *pb.RuntimeProjectAttachment
	reader    client.Reader
	namespace string
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting APIHubRuntimeProjectAttachment", "name", a.id.String())

	req := &pb.GetRuntimeProjectAttachmentRequest{Name: a.id.String()}
	pbObj, err := a.gcpClient.GetRuntimeProjectAttachment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting APIHubRuntimeProjectAttachment %q: %w", a.id.String(), err)
	}

	a.actual = pbObj
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating APIHubRuntimeProjectAttachment", "name", a.id.String())

	req := &pb.CreateRuntimeProjectAttachmentRequest{
		Parent:                     a.id.ParentString(),
		RuntimeProjectAttachmentId: a.id.Runtime_project_attachment,
		RuntimeProjectAttachment:   a.desired,
	}

	_, err := a.gcpClient.CreateRuntimeProjectAttachment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating APIHubRuntimeProjectAttachment %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created APIHubRuntimeProjectAttachment", "name", a.id.String())

	// Fetch the fully populated resource to be safe
	getReq := &pb.GetRuntimeProjectAttachmentRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetRuntimeProjectAttachment(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting created APIHubRuntimeProjectAttachment %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating APIHubRuntimeProjectAttachment", "name", a.id.String())

	diff, err := compareRuntimeProjectAttachment(ctx, a.actual, a.desired, a.reader, a.namespace)
	if err != nil {
		return err
	}

	if diff.HasDiff() {
		diff.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diff)
		return fmt.Errorf("APIHubRuntimeProjectAttachment is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.APIHubRuntimeProjectAttachment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(apihub.APIHubRuntimeProjectAttachmentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = a.id.Location
	obj.Spec.ProjectRef = &refs.ProjectRef{External: fmt.Sprintf("projects/%s", a.id.Project)}
	obj.Spec.ResourceID = &a.id.Runtime_project_attachment

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting APIHubRuntimeProjectAttachment", "name", a.id.String())

	req := &pb.DeleteRuntimeProjectAttachmentRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteRuntimeProjectAttachment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting APIHubRuntimeProjectAttachment %s: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.RuntimeProjectAttachment) error {
	mapCtx := &direct.MapContext{}
	status := &krm.APIHubRuntimeProjectAttachmentStatus{}
	status.ObservedState = apihub.APIHubRuntimeProjectAttachmentObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := latest.Name
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, status, nil)
}

func compareRuntimeProjectAttachment(ctx context.Context, actual, desired *pb.RuntimeProjectAttachment, reader client.Reader, namespace string) (*structuredreporting.Diff, error) {
	if actual.GetRuntimeProject() == desired.GetRuntimeProject() {
		return &structuredreporting.Diff{}, nil
	}

	// Since actual.RuntimeProject might be projects/<number> and desired.RuntimeProject might be projects/<id> (or vice versa),
	// we attempt to resolve both to see if they refer to the same project.
	desiredRef := &refs.ProjectRef{External: desired.GetRuntimeProject()}
	actualRef := &refs.ProjectRef{External: actual.GetRuntimeProject()}

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
	diff.AddField("spec.runtimeProjectRef", actual.GetRuntimeProject(), desired.GetRuntimeProject())
	return diff, nil
}
