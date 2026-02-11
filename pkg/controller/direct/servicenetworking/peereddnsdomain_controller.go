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
// proto.message: mockgcp.cloud.servicenetworking.v1.PeeredDnsDomain
// crd.kind: ServiceNetworkingPeeredDnsDomain
// crd.version: v1alpha1

package servicenetworking

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	"google.golang.org/api/option"
	api "google.golang.org/api/servicenetworking/v1"
	gcp "google.golang.org/api/servicenetworking/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicenetworking/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.ServiceNetworkingPeeredDNSDomainGVK, NewPeeredDNSDomainModel)
}

func NewPeeredDNSDomainModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &peeredDnsDomainModel{config: *config}, nil
}

var _ directbase.Model = &peeredDnsDomainModel{}

type peeredDnsDomainModel struct {
	config config.ControllerConfig
}

func (m *peeredDnsDomainModel) client(ctx context.Context) (*gcp.APIService, error) {
	var opts []option.ClientOption

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building servicenetworking client: %w", err)
	}

	return gcpClient, err
}

func (m *peeredDnsDomainModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	obj := &krm.ServiceNetworkingPeeredDNSDomain{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := obj.Spec.NetworkRef.Normalize(ctx, kube, u.GetNamespace()); err != nil {
		return nil, err
	}

	// Make sure we use project number, not project ID
	// We have to do this early to avoid triggering a difference vs status.externalRef
	if err := obj.Spec.NetworkRef.ConvertToProjectNumber(ctx, m.config.ProjectMapper); err != nil {
		return nil, err
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := ServiceNetworkingPeeredDNSDomainSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &peeredDNSDomainAdapter{
		gcpClient: gcpClient,
		id:        id.(*krm.PeeredDNSDomainIdentity),
		desired:   desired,
	}, nil
}

func (m *peeredDnsDomainModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if s, ok := strings.CutPrefix(url, "//servicenetworking.googleapis.com/"); ok {
		s = strings.TrimPrefix(s, "v1/")

		var id krm.PeeredDNSDomainIdentity
		err := id.FromExternal(s)
		if err != nil {
			log.V(2).Error(err, "url did not match ServiceNetworkingPeeredDnsDomain format", "url", url)
			return nil, nil
		}

		if err := id.Network.ConvertToProjectNumber(ctx, m.config.ProjectMapper); err != nil {
			return nil, err
		}
		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}
		return &peeredDNSDomainAdapter{
			gcpClient: gcpClient,
			id:        &id,
		}, nil
	}
	return nil, nil
}

type peeredDNSDomainAdapter struct {
	gcpClient *gcp.APIService
	id        *krm.PeeredDNSDomainIdentity
	desired   *api.PeeredDnsDomain
	actual    *api.PeeredDnsDomain
}

var _ directbase.Adapter = &peeredDNSDomainAdapter{}

func (a *peeredDNSDomainAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting servicenetworking peeredDnsDomain", "name", a.id)

	parent := "services/servicenetworking.googleapis.com/" + a.id.Network.String()

	actualList, err := a.gcpClient.Services.Projects.Global.Networks.PeeredDnsDomains.List(parent).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting servicenetworking peeredDnsDomain %q from gcp: %w", a.id.String(), err)
	}

	for _, actual := range actualList.PeeredDnsDomains {
		if actual.Name == a.id.Name {
			a.actual = actual
			return true, nil
		}
	}
	return false, nil
}

func (a *peeredDNSDomainAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating servicenetworking peeredDnsDomain", "name", a.id)

	desired := ReflectClone(a.desired)
	desired.Name = a.id.Name

	parent := "services/servicenetworking.googleapis.com/" + a.id.Network.String()

	op, err := a.gcpClient.Services.Projects.Global.Networks.PeeredDnsDomains.Create(parent, desired).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating servicenetworking peeredDnsDomain %s: %w", a.id.String(), err)
	}
	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for servicenetworking peeredDnsDomain %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created servicenetworking peeredDnsDomain in gcp", "name", a.id)

	// There's no observed state, easier to just set it to the desired state
	created := desired

	status := &krm.ServiceNetworkingPeeredDNSDomainStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ServiceNetworkingPeeredDNSDomainObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *peeredDNSDomainAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating servicenetworking peeredDnsDomain", "name", a.id)

	desired := ReflectClone(a.desired)
	desired.Name = a.id.Name

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(desired.DnsSuffix, a.actual.DnsSuffix) {
		updateMask.Paths = append(updateMask.Paths, "dns_suffix")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	return fmt.Errorf("cannot update peeredDnsDomain object (values are immutable)")
	// op, err := a.gcpClient.Services.Projects.Global.Networks.PeeredDnsDomains.Create(parent, desired).Context(ctx).Do()
	// if err != nil {
	// 	return fmt.Errorf("updating servicenetworking peeredDnsDomain %s: %w", a.id.String(), err)
	// }
	// if err := waitForOperation(ctx, op); err != nil {
	// 	return fmt.Errorf("waiting for servicenetworking peeredDnsDomain %s update operation: %w", a.id.String(), err)
	// }
	// log.V(2).Info("successfully updated servicenetworking peeredDnsDomain", "name", a.id)

	// status := &krm.DiscoveryEngineDataStoreStatus{}
	// mapCtx := &direct.MapContext{}
	// status.ObservedState = DiscoveryEngineDataStoreObservedState_FromProto(mapCtx, updated)
	// if mapCtx.Err() != nil {
	// 	return mapCtx.Err()
	// }
	// return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *peeredDNSDomainAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.ServiceNetworkingPeeredDNSDomain{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ServiceNetworkingPeeredDNSDomainSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.NetworkRef = &computev1beta1.ComputeNetworkRef{External: a.id.Network.String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Name)
	u.SetGroupVersionKind(krm.ServiceNetworkingPeeredDNSDomainGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *peeredDNSDomainAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting servicenetworking peeredDnsDomain", "name", a.id)

	fqn := a.id.String()
	op, err := a.gcpClient.Services.Projects.Global.Networks.PeeredDnsDomains.Delete(fqn).Context(ctx).Do()
	if err != nil {
		return false, fmt.Errorf("deleting servicenetworking peeredDnsDomain %s: %w", a.id.String(), err)
	}
	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for servicenetworking peeredDnsDomain %s deletion: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted servicenetworking peeredDnsDomain", "name", a.id)

	return true, nil
}

// waitForOperation waits for the given operation to complete.
func (a *peeredDNSDomainAdapter) waitForOperation(ctx context.Context, op *api.Operation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("waiting for operation to complete", "operation", op.Name)

	opName := op.Name
	if opName == "" {
		return fmt.Errorf("operation name is empty")
	}

	timeoutAt := time.Now().Add(5 * time.Minute)
	for {
		op, err := a.gcpClient.Operations.Get(opName).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting servicenetworking operation %s: %w", opName, err)
		}
		if op.Done {
			return nil
		}
		if time.Now().After(timeoutAt) {
			return fmt.Errorf("timeout waiting for servicenetworking operation %s to complete", opName)
		}
		log.V(2).Info("operation still in progress", "operation", op.Name, "done", op.Done)
		// Sleep for a while before checking the operation status again
		time.Sleep(2 * time.Second)
	}
}
