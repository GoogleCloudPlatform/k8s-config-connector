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

package composer

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/orchestration/airflow/service/apiv1"
	composerpb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	pb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComposerEnvironmentGVK, NewEnvironmentModel)
}

func NewEnvironmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEnvironment{config: *config}, nil
}

var _ directbase.Model = &modelEnvironment{}

type modelEnvironment struct {
	config config.ControllerConfig
}

func (m *modelEnvironment) client(ctx context.Context) (*gcp.EnvironmentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewEnvironmentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Environment client: %w", err)
	}
	return gcpClient, err
}

func (m *modelEnvironment) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	copiedu := u.DeepCopy()
	if err := label.ComputeLabels(u.DeepCopy()); err != nil {
		return nil, err
	}
	obj := &krm.ComposerEnvironment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(copiedu.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEnvironmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	copied := obj.DeepCopy()
	desired := ComposerEnvironmentSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()
	// Get composer GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EnvironmentAdapter{
		id:                 id,
		gcpClient:          gcpClient,
		desired:            desired,
		lastModifiedCookie: obj.Status.LastModifiedCookie,
	}, nil
}

func (m *modelEnvironment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EnvironmentAdapter struct {
	id                 *krm.EnvironmentIdentity
	gcpClient          *gcp.EnvironmentsClient
	desired            *composerpb.Environment
	actual             *composerpb.Environment
	lastModifiedCookie *string
}

var _ directbase.Adapter = &EnvironmentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EnvironmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Environment", "name", a.id)

	req := &composerpb.GetEnvironmentRequest{Name: a.id.String()}
	environmentpb, err := a.gcpClient.GetEnvironment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Environment %q: %w", a.id, err)
	}

	a.actual = environmentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EnvironmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Environment", "name", a.id)

	req := &composerpb.CreateEnvironmentRequest{
		Parent:      a.id.Parent().String(),
		Environment: a.desired,
	}
	op, err := a.gcpClient.CreateEnvironment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Environment %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Environment %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Environment", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.ComposerEnvironmentStatus{}
	status.ObservedState = ComposerEnvironmentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	newCookie, err := common.NewCookie(a.desired, created)
	if err != nil {
		return nil
	}
	status.LastModifiedCookie = direct.LazyPtr(newCookie.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *EnvironmentAdapter) updateStatus(ctx context.Context, updated *pb.Environment, updateOp *directbase.UpdateOperation) error {
	status := &krm.ComposerEnvironmentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ComposerEnvironmentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	updatedCookie, err := common.NewCookie(a.desired, updated)
	if err != nil {
		return err
	}
	status.LastModifiedCookie = direct.LazyPtr(updatedCookie.String())
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EnvironmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Environment", "name", a.id)

	currentCookie, err := common.NewCookie(a.desired, a.actual)
	if err != nil {
		return err
	}

	if currentCookie.Equal(a.lastModifiedCookie) {
		log.V(2).Info("resource is up to date", "name", a.id)
		return a.updateStatus(ctx, a.actual, updateOp)
	}

	// The GCP server is non-declarative friendly as it allows many fields to be empty (unset) in creation,
	// but require them to be passed through in update. What makes it worse, the update field mask require to
	// only provide different field and only accept one field update per request.
	// Meanwhile, the Proto does not have field mask to tell whether a field is OUTPUT-ONLY or IMMUTABLE.

	// Here's the failure cases:
	// - Patch returns 400 if optional fields are empty, if not in update field mask --> Default to GCP is not allowed.
	// - Patch returns 400 if optional fields are empty, if is in update field mask --> unset in GCP is not allowed.
	// - Patch returns 400 if optional fields are unchanged, if is in update field mask --> field has to be a different from GCP stored obj.
	// - Patch only allows one field update per request.

	// To simply the logic, we use what actual proto has (including GCP default and OUTPUT-ONLY), and merge it with the desired.
	// Only the different INPUT fields from desired will be passed in update field mask.
	desiredWithGCPDefault := proto.Clone(a.actual).(*composerpb.Environment)
	proto.Merge(desiredWithGCPDefault, a.desired)

	// Note: Environment proto misses the field behavior about OUTPUT-ONLY and IMMUTABLE,
	// so a pure comparison between desired and actual will give wrong fields that contain OUTPUT-ONLY and IMMUTABLE (which is optional with default)
	// But since we have merge the actual OUTPUT-ONLY and IMMUTABLE back to the desiredWithGCPDefault, it will skip those fields.
	paths, err := common.CompareProtoMessage(desiredWithGCPDefault, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	sortedPaths := sets.List(paths)

	if len(sortedPaths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, a.actual, updateOp)
	}

	// GCP server only allows updating one field per request.
	for _, f := range sortedPaths {
		updateMask := &fieldmaskpb.FieldMask{
			Paths: []string{f},
		}
		req := &composerpb.UpdateEnvironmentRequest{
			Name:        a.id.String(),
			UpdateMask:  updateMask,
			Environment: desiredWithGCPDefault,
		}
		op, err := a.gcpClient.UpdateEnvironment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Environment %s: %w", a.id, err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Environment %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Environment", "name", a.id)
		if err := a.updateStatus(ctx, updated, updateOp); err != nil {
			return err
		}
	}
	return nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EnvironmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComposerEnvironment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComposerEnvironmentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ComposerEnvironmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EnvironmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Environment", "name", a.id)

	req := &composerpb.DeleteEnvironmentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEnvironment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Environment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Environment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Environment", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Environment %s: %w", a.id, err)
	}
	return true, nil
}
