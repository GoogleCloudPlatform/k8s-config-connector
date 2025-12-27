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

package apigee

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/apigee/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ApigeeInstanceGVK, NewApigeeInstanceModel)
}

func NewApigeeInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApigeeInstance{config: config}, nil
}

var _ directbase.Model = &modelApigeeInstance{}

type modelApigeeInstance struct {
	config *config.ControllerConfig
}

func (m *modelApigeeInstance) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ApigeeInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeInstanceIdentity)

	// Get apigee GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &ApigeeInstanceAdapter{
		id:               id,
		k8sClient:        reader,
		instancesClient:  gcpClient.instancesClient(),
		operationsClient: gcpClient.operationsClient(),
		desired:          obj,
	}, nil
}

func (m *modelApigeeInstance) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ApigeeInstanceAdapter struct {
	id               *krm.ApigeeInstanceIdentity
	k8sClient        client.Reader
	instancesClient  *api.OrganizationsInstancesService
	operationsClient *api.OrganizationsOperationsService
	desired          *krm.ApigeeInstance
	actual           *api.GoogleCloudApigeeV1Instance
}

var _ directbase.Adapter = &ApigeeInstanceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ApigeeInstanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ApigeeInstance", "name", a.id)

	googlecloudapigeev1instancepb, err := a.instancesClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApigeeInstance %q: %w", a.id, err)
	}

	a.actual = googlecloudapigeev1instancepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApigeeInstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ApigeeInstance", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	// Resolve references
	if err := ResolveApigeeInstanceRefs(ctx, a.k8sClient, desired); err != nil {
		return err
	}
	// Convert to proto
	resource := ApigeeInstanceSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.ResourceID

	op, err := a.instancesClient.Create(a.id.ParentID.String(), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeInstance %s: %w", a.id, err)
	}
	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("ApigeeInstance %s waiting creation: %w", a.id, err)
	}

	created, err := a.instancesClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created ApigeeInstance: %w", err)
	}

	log.V(2).Info("successfully created ApigeeInstance", "name", a.id)

	status := &krm.ApigeeInstanceStatus{}
	status.ObservedState = ApigeeInstanceObservedState_FromAPI(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApigeeInstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ApigeeInstance", "name", a.id)
	mapCtx := &direct.MapContext{}
	updateMask := fieldmaskpb.FieldMask{}

	desired := a.desired.DeepCopy()

	// Resolve references
	if err := ResolveApigeeInstanceRefs(ctx, a.k8sClient, desired); err != nil {
		return err
	}
	// Convert to proto
	resource := ApigeeInstanceSpec_ToAPI(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	if resource.AccessLoggingConfig != nil {
		if resource.AccessLoggingConfig.Enabled != a.actual.AccessLoggingConfig.Enabled {
			report.AddField("access_logging_config.enabled", a.actual.AccessLoggingConfig.Enabled, resource.AccessLoggingConfig.Enabled)
			updateMask.Paths = append(updateMask.Paths, "access_logging_config.enabled")
		}
		if resource.AccessLoggingConfig.Filter != a.actual.AccessLoggingConfig.Filter {
			report.AddField("access_logging_config.filter", a.actual.AccessLoggingConfig.Filter, resource.AccessLoggingConfig.Filter)
			updateMask.Paths = append(updateMask.Paths, "access_logging_config.filter")
		}
	}
	if resource.ConsumerAcceptList != nil && !reflect.DeepEqual(asSortedCopy(resource.ConsumerAcceptList), asSortedCopy(a.actual.ConsumerAcceptList)) {
		report.AddField("consumer_accept_list", a.actual.ConsumerAcceptList, resource.ConsumerAcceptList)
		updateMask.Paths = append(updateMask.Paths, "consumer_accept_list")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.ApigeeInstanceStatus{}
		status.ObservedState = ApigeeInstanceObservedState_FromAPI(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	structuredreporting.ReportDiff(ctx, report)

	op, err := a.instancesClient.Patch(a.id.String(), resource).UpdateMask(strings.Join(updateMask.Paths, ",")).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating ApigeeInstance %s: %w", a.id, err)
	}
	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("ApigeeInstance %s waiting update: %w", a.id, err)
	}
	updated, err := a.instancesClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting updated ApigeeInstance: %w", err)
	}
	log.V(2).Info("successfully updated ApigeeInstance", "name", a.id)

	status := &krm.ApigeeInstanceStatus{}
	status.ObservedState = ApigeeInstanceObservedState_FromAPI(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ApigeeInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeInstanceSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &krm.ApigeeOrganizationRef{External: a.id.ParentID.String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ResourceID)
	u.SetGroupVersionKind(krm.ApigeeInstanceGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ApigeeInstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ApigeeInstance", "name", a.id)

	op, err := a.instancesClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ApigeeInstance, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeInstance %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ApigeeInstance", "name", a.id)

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("waiting delete ApigeeInstance %s: %w", a.id, err)
	}
	return true, nil
}
