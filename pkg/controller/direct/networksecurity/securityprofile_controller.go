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

package networksecurity

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	api "google.golang.org/api/networksecurity/v1"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecuritySecurityProfileGVK, NewSecurityProfileModel)
}

func NewSecurityProfileModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSecurityProfile{config: *config}, nil
}

var _ directbase.Model = &modelSecurityProfile{}

type modelSecurityProfile struct {
	config config.ControllerConfig
}

func (m *modelSecurityProfile) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SecurityProfile client: %w", err)
	}
	return gcpClient, err
}

func (m *modelSecurityProfile) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecuritySecurityProfile{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.SecurityProfileIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	if obj.Spec.CustomInterceptProfile != nil && obj.Spec.CustomInterceptProfile.InterceptEndpointGroupRef != nil {
		if err := obj.Spec.CustomInterceptProfile.InterceptEndpointGroupRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("normalizing InterceptEndpointGroupRef: %w", err)
		}
	}
	if obj.Spec.CustomMirroringProfile != nil && obj.Spec.CustomMirroringProfile.MirroringEndpointGroupRef != nil {
		if err := obj.Spec.CustomMirroringProfile.MirroringEndpointGroupRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("normalizing MirroringEndpointGroupRef: %w", err)
		}
	}

	// Get networksecurity GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SecurityProfileAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelSecurityProfile) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type SecurityProfileAdapter struct {
	id        *krm.SecurityProfileIdentity
	gcpClient *api.Service
	desired   *krm.NetworkSecuritySecurityProfile
	actual    *api.SecurityProfile
}

var _ directbase.Adapter = &SecurityProfileAdapter{}

func (a *SecurityProfileAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecurityProfile", "name", a.id)

	var err error
	a.actual, err = a.gcpClient.Organizations.Locations.SecurityProfiles.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecurityProfile %q: %w", a.id, err)
	}
	return true, nil
}

func (a *SecurityProfileAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SecurityProfile", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkSecuritySecurityProfileSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var op *api.Operation
	var err error
	op, err = a.gcpClient.Organizations.Locations.SecurityProfiles.Create(a.id.Parent(), resource).SecurityProfileId(a.id.ID()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating SecurityProfile %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("SecurityProfile %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created SecurityProfile", "name", a.id)

	// Fetch again to get fully populated object
	_, err = a.Find(ctx)
	if err != nil {
		return err
	}

	status := &krm.NetworkSecuritySecurityProfileStatus{}
	status.ObservedState = NetworkSecuritySecurityProfileObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *SecurityProfileAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SecurityProfile", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := NetworkSecuritySecurityProfileSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])
	{
		if !reflect.DeepEqual(a.desired.Spec.Description, a.actual.Description) {
			paths.Insert("description")
		}
		if !reflect.DeepEqual(a.desired.Spec.Labels, a.actual.Labels) {
			paths.Insert("labels")
		}
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)

		var op *api.Operation
		var err error
		updateMask := strings.Join(sets.List(paths), ",")
		op, err = a.gcpClient.Organizations.Locations.SecurityProfiles.Patch(a.id.String(), desiredPb).UpdateMask(updateMask).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating SecurityProfile %s: %w", a.id, err)
		}
		err = a.waitForOperation(ctx, op)
		if err != nil {
			return fmt.Errorf("SecurityProfile %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated SecurityProfile", "name", a.id)
	}

	// Fetch again to get fully populated object
	_, err := a.Find(ctx)
	if err != nil {
		return err
	}

	status := &krm.NetworkSecuritySecurityProfileStatus{}
	status.ObservedState = NetworkSecuritySecurityProfileObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *SecurityProfileAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecuritySecurityProfile{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *NetworkSecuritySecurityProfileSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if a.id.Project != "" {
		obj.Spec.ProjectRef = &refs.ProjectRef{External: "projects/" + a.id.Project}
	} else if a.id.Organization != "" {
		obj.Spec.OrganizationRef = &refs.OrganizationRef{External: "organizations/" + a.id.Organization}
	}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkSecuritySecurityProfileGVK)

	u.Object = uObj
	return u, nil
}

func (a *SecurityProfileAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SecurityProfile", "name", a.id)

	var op *api.Operation
	var err error
	op, err = a.gcpClient.Organizations.Locations.SecurityProfiles.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent SecurityProfile, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting SecurityProfile %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted SecurityProfile", "name", a.id)

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return false, fmt.Errorf("waiting delete SecurityProfile %s: %w", a.id, err)
	}
	return true, nil
}

func (a *SecurityProfileAdapter) waitForOperation(ctx context.Context, op *api.Operation) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		latest, err := a.gcpClient.Organizations.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation %q: %w", op.Name, err)
		}

		if latest.Done {
			if latest.Error != nil {
				return fmt.Errorf("operation failed: %v", latest.Error.Message)
			}
			return nil
		}
		time.Sleep(2 * time.Second)
	}
}
