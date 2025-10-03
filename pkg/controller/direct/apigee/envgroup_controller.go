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
	"sort"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/apigee/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ctrlName = "apigee-envgroup-controller"

func init() {
	registry.RegisterModel(krm.ApigeeEnvgroupGVK, NewApigeeEnvgroupModel)
}

func NewApigeeEnvgroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApigeeEnvgroup{config: config}, nil
}

var _ directbase.Model = &modelApigeeEnvgroup{}

type modelApigeeEnvgroup struct {
	config *config.ControllerConfig
}

func (m *modelApigeeEnvgroup) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	obj := &krm.ApigeeEnvgroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ApigeeEnvgroupIdentity)

	mapCtx := &direct.MapContext{}
	desired := obj
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = id.ResourceID

	return &Adapter{
		id:               id,
		desired:          desired,
		envgroupsClient:  gcpClient.envgroupsClient(),
		operationsClient: gcpClient.operationsClient(),
	}, nil
}

func (m *modelApigeeEnvgroup) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id               *krm.ApigeeEnvgroupIdentity
	desired          *krm.ApigeeEnvgroup
	actual           *api.GoogleCloudApigeeV1EnvironmentGroup
	envgroupsClient  *api.OrganizationsEnvgroupsService
	operationsClient *api.OrganizationsOperationsService
}

var _ directbase.Adapter = &Adapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting ApigeeEnvgroup", "name", a.id)

	envgroup, err := a.envgroupsClient.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ApigeeEnvgroup %q: %w", a.id, err)
	}

	a.actual = envgroup
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating ApigeeEnvgroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	req := ApigeeEnvgroupSpec_ToApi(mapCtx, &a.desired.Spec, a.desired.Name)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	op, err := a.envgroupsClient.Create(a.id.ParentID.String(), req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ApigeeEnvgroup %s: %w", a.fullyQualifiedName(), err)
	}

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("waiting for ApigeeEnvgroup %s creation: %w", a.id, err)
	}

	created, err := a.envgroupsClient.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created ApigeeEnvgroup: %w", err)
	}

	log.V(2).Info("successfully created ApigeeEnvgroup", "ApigeeEnvgroup", created)

	status := &krm.ApigeeEnvgroupStatus{}
	status.ObservedState = ApigeeEnvgroupObservedState_FromApi(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("patching ApigeeEnvgroup", a.fullyQualifiedName())
	mapCtx := &direct.MapContext{}
	updateMask := fieldmaskpb.FieldMask{}

	req := ApigeeEnvgroupSpec_ToApi(mapCtx, &a.desired.Spec, a.desired.Name)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Sorts the Hostname lists so that the comparison is deterministic
	if !reflect.DeepEqual(asSortedCopy(req.Hostnames), asSortedCopy(a.actual.Hostnames)) {
		log.V(2).Info("change detected: hostnames")
		updateMask.Paths = append(updateMask.Paths, "hostnames")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.ApigeeEnvgroupStatus{}
		status.ObservedState = ApigeeEnvgroupObservedState_FromApi(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	clusterName := a.id.String()
	op, err := a.envgroupsClient.Patch(clusterName, req).UpdateMask(strings.Join(updateMask.Paths, ",")).Context(ctx).Do()
	if err != nil {
		return err
	}
	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return fmt.Errorf("waiting for ApigeeEnvgroup update %s: %w", a.id, err)
	}

	updated, err := a.envgroupsClient.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting updated ApigeeEnvgroup: %w", err)
	}
	log.V(2).Info("successfully updated ApigeeEnvgroup", "ApigeeEnvgroup", updated)

	status := &krm.ApigeeEnvgroupStatus{}
	status.ObservedState = ApigeeEnvgroupObservedState_FromApi(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ApigeeEnvgroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ApigeeEnvgroupSpec_FromApi(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Parent.OrganizationRef = &krm.ApigeeOrganizationRef{External: a.id.ParentID.String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ResourceID)
	u.SetGroupVersionKind(krm.ApigeeEnvgroupGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting ApigeeEnvgroup", "name", a.id)

	op, err := a.envgroupsClient.Delete(a.fullyQualifiedName()).Context(ctx).Do()

	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ApigeeEnvgroup, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting ApigeeEnvgroup %s: %w", a.id, err)
	}

	if err := WaitForApigeeOp(ctx, a.operationsClient, op); err != nil {
		return false, fmt.Errorf("ApigeeEnvgroup deletion failed: %w", err)
	}
	return true, nil
}

func (a *Adapter) fullyQualifiedName() string {
	return a.id.String()
}

func asSortedCopy(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	sort.Strings(out)

	return out
}
