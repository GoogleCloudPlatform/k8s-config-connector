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

package documentai

import (
	"context"
	"fmt"
	"time"

	gcp "cloud.google.com/go/documentai/apiv1"
	documentaipb "cloud.google.com/go/documentai/apiv1/documentaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DocumentAIProcessorVersionGVK, NewProcessorVersionModel)
}

func NewProcessorVersionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelProcessorVersion{config: *config}, nil
}

var _ directbase.Model = &modelProcessorVersion{}

type modelProcessorVersion struct {
	config config.ControllerConfig
}

func (m *modelProcessorVersion) client(ctx context.Context) (*gcp.DocumentProcessorClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDocumentProcessorRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ProcessorVersion client: %w", err)
	}
	return gcpClient, err
}

func (m *modelProcessorVersion) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DocumentAIProcessorVersion{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewProcessorVersionIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get documentai GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ProcessorVersionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelProcessorVersion) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ProcessorVersionAdapter struct {
	id        *krm.ProcessorVersionIdentity
	gcpClient *gcp.DocumentProcessorClient
	desired   *krm.DocumentAIProcessorVersion
	actual    *documentaipb.ProcessorVersion
}

var _ directbase.Adapter = &ProcessorVersionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ProcessorVersionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ProcessorVersion", "name", a.id)

	// Check whether Config Connector knows the resource identity.
	// If not, Config Connector saves one GCP GET call, and starts the CREATE call directly.
	// This is mostly for GCP services that do not allow user to specify ID, but assign an ID when creating the object.
	if a.id.ID() == "" {
		return false, nil
	}

	req := &documentaipb.GetProcessorVersionRequest{Name: a.id.String()}
	processorversionpb, err := a.gcpClient.GetProcessorVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ProcessorVersion %q: %w", a.id, err)
	}

	a.actual = processorversionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ProcessorVersionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ProcessorVersion", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DocumentAIProcessorVersionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &documentaipb.TrainProcessorVersionRequest{
		Parent:           a.id.Parent().String(),
		ProcessorVersion: resource,
	}
	op, err := a.gcpClient.TrainProcessorVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ProcessorVersion %s: %w", a.id, err)
	}

	//Get server generated version id from metadata
	metadata, err := op.Metadata()
	if err != nil {
		return fmt.Errorf("waiting create Version %s: %w", a.id, err)
	}

	// op.wait() returns once the version resource exists, but the version's state can still be "creating."
	// It takes a few seconds to reach a stable state like "failed" or "undeployed."
	// Retrieve the resource and check its state to determine if the operation has finished.
	getReq := &documentaipb.GetProcessorVersionRequest{Name: metadata.GetCommonMetadata().GetResource()}
	created, err := a.getCreatedResource(ctx, 5*time.Second, 5*time.Minute, getReq)
	if err != nil {
		return fmt.Errorf("waiting create Version %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ProcessorVersion", "name", a.id)

	status := &krm.DocumentAIProcessorVersionStatus{}
	status.ObservedState = DocumentAIProcessorVersionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := created.GetName()
	status.ExternalRef = direct.LazyPtr(externalRef)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ProcessorVersionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// no-op, just update obj status
	mapCtx := &direct.MapContext{}
	updated := a.actual
	status := &krm.DocumentAIProcessorVersionStatus{}
	status.ObservedState = DocumentAIProcessorVersionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := updated.GetName()
	status.ExternalRef = direct.LazyPtr(externalRef)
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ProcessorVersionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DocumentAIProcessorVersion{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DocumentAIProcessorVersionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.DocumentAIProcessorVersionGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ProcessorVersionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ProcessorVersion", "name", a.id)
	req := &documentaipb.DeleteProcessorVersionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteProcessorVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ProcessorVersion, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ProcessorVersion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ProcessorVersion", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		// todo (b/368419476): similarly, DocumentAI service does not provide a valid response on success.
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return false, fmt.Errorf("waiting delete Version %s: %w", a.id, err)
		}
	}
	return true, nil
}

// Get created resource by making sure its state is a stable state such as FAILED or DEPLOYED or UNDEPLOYED.
// Set a poll interval and a timeout to avoid potential quota issue and excessive GET calls.
// If it still hasn't reached a stable state after a certain time, we'll consider it a "fail".
func (a *ProcessorVersionAdapter) getCreatedResource(ctx context.Context, pollInterval, timeout time.Duration, getReq *documentaipb.GetProcessorVersionRequest) (*documentaipb.ProcessorVersion, error) {
	var createdVersion *documentaipb.ProcessorVersion
	err := wait.PollImmediateWithContext(ctx, pollInterval, timeout, func(ctx context.Context) (bool, error) {
		created, err := a.gcpClient.GetProcessorVersion(ctx, getReq)
		if err != nil {
			return false, fmt.Errorf("getting created version %q: %w", a.id, err)
		}
		switch created.State {
		// Currently we expect resource state to be FAILED after creation as there's no training data
		case documentaipb.ProcessorVersion_FAILED,
			documentaipb.ProcessorVersion_DEPLOYED,
			documentaipb.ProcessorVersion_UNDEPLOYED:
			createdVersion = created
			return true, nil
		// Keep polling if state is CREATING
		case documentaipb.ProcessorVersion_CREATING:
			return false, nil
		// todo: Handle other states, i.e. STATE_UNSPECIFIED, DEPLOYING, etc
		default:
			return false, nil
		}
	})

	if err != nil {
		return nil, err
	}
	return createdVersion, nil
}
