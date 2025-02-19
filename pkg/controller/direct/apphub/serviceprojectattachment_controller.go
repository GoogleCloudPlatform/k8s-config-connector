// Copyright 2025 Google LLC
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

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/apphub/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf

	"google.golang.org/api/option"


	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.AppHubServiceProjectAttachmentGVK, NewServiceProjectAttachmentModel)
}

func NewServiceProjectAttachmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelServiceProjectAttachment{config: *config}, nil
}

var _ directbase.Model = &modelServiceProjectAttachment{}

type modelServiceProjectAttachment struct {
	config config.ControllerConfig
}

func (m *modelServiceProjectAttachment) client(ctx context.Context) (*gcp.Client, error) {
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

func (m *modelServiceProjectAttachment) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.AppHubServiceProjectAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewServiceProjectAttachmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get apphub GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServiceProjectAttachmentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelServiceProjectAttachment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ServiceProjectAttachmentAdapter struct {
	id        *krm.ServiceProjectAttachmentIdentity
	gcpClient *gcp.Client
	desired   *krm.AppHubServiceProjectAttachment
	// actual    *apphubpb.ServiceProjectAttachment
}

var _ directbase.Adapter = &ServiceProjectAttachmentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ServiceProjectAttachmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ServiceProjectAttachment", "name", a.id)

	// req := &apphubpb.GetServiceProjectAttachmentRequest{Name: a.id.String()}
	// serviceprojectattachmentpb, err := a.gcpClient.GetServiceProjectAttachment(ctx, req)
	// if err != nil {
	//	if direct.IsNotFound(err) {
	//		return false, nil
	//	}
	//	return false, fmt.Errorf("getting ServiceProjectAttachment %q: %w", a.id, err)
	// }

	// a.actual = serviceprojectattachmentpb
	return false, nil
}

func (a *ServiceProjectAttachmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	return createOp.UpdateStatus(ctx, nil, nil)
}
// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ServiceProjectAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return updateOp.UpdateStatus(ctx, nil, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ServiceProjectAttachmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	/*if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}*/
	u := &unstructured.Unstructured{}

	obj := &krm.AppHubServiceProjectAttachment{}
	mapCtx := &direct.MapContext{}
	// obj.Spec = direct.ValueOf(AppHubServiceProjectAttachmentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	// u.SetName(a.actual.Id)
	u.SetGroupVersionKind(krm.AppHubServiceProjectAttachmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ServiceProjectAttachmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ServiceProjectAttachment", "name", a.id)

	// req := &apphubpb.DeleteServiceProjectAttachmentRequest{Name: a.id.String()}
	// op, err := a.gcpClient.DeleteServiceProjectAttachment(ctx, req)
	// if err != nil {
	//	if direct.IsNotFound(err) {
	//		// Return success if not found (assume it was already deleted).
	//		log.V(2).Info("skipping delete for non-existent ServiceProjectAttachment, assuming it was already deleted", "name", a.id.String())
	//		return true, nil
	//	}
	//	return false, fmt.Errorf("deleting ServiceProjectAttachment %s: %w", a.id, err)
	// }
	log.V(2).Info("successfully deleted ServiceProjectAttachment", "name", a.id)

	// err = op.Wait(ctx)
	// if err != nil {
	//	return false, fmt.Errorf("waiting delete ServiceProjectAttachment %s: %w", a.id, err)
	// }
	return true, nil
}
