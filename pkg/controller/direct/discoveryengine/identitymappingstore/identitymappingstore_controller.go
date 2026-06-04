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

package identitymappingstore

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/discoveryengine"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineIdentityMappingStoreGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context, projectID string) (*gcp.IdentityMappingStoreClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if config.UserProjectOverride && config.BillingProject == "" {
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewIdentityMappingStoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine identitymappingstore client: %w", err)
	}

	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineIdentityMappingStore{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idObj, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idObj.(*krm.DiscoveryEngineIdentityMappingStoreIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idObj)
	}

	mapCtx := &direct.MapContext{}
	desired := discoveryengine.DiscoveryEngineIdentityMappingStoreSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		parsed := &krm.DiscoveryEngineIdentityMappingStoreIdentity{}
		err := parsed.FromExternal(strings.TrimPrefix(url, "//discoveryengine.googleapis.com/"))
		if err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineIdentityMappingStore format", "url", url)
		} else {
			gcpClient, err := m.client(ctx, parsed.Project)
			if err != nil {
				return nil, err
			}
			return &adapter{
				gcpClient: gcpClient,
				id:        parsed,
			}, nil
		}
	}
	return nil, nil
}

type adapter struct {
	gcpClient *gcp.IdentityMappingStoreClient
	id        *krm.DiscoveryEngineIdentityMappingStoreIdentity
	desired   *pb.IdentityMappingStore
	actual    *pb.IdentityMappingStore
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine identitymappingstore", "name", a.id.String())

	req := &pb.GetIdentityMappingStoreRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetIdentityMappingStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine identitymappingstore %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine identitymappingstore", "name", a.id.String())

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateIdentityMappingStoreRequest{
		Parent:                 parent,
		IdentityMappingStore:   desired,
		IdentityMappingStoreId: a.id.IdentityMappingStore,
	}
	created, err := a.gcpClient.CreateIdentityMappingStore(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine identitymappingstore %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created discoveryengine identitymappingstore in gcp", "name", a.id.String())

	status := &krm.DiscoveryEngineIdentityMappingStoreStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = discoveryengine.DiscoveryEngineIdentityMappingStoreObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine identitymappingstore", "name", a.id.String())

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	paths, report, err := common.CompareProtoMessageStructuredDiff(desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	structuredreporting.ReportDiff(ctx, report)

	if len(paths) > 0 {
		return fmt.Errorf("cannot update DiscoveryEngineIdentityMappingStore %q: fields changed: %v; DiscoveryEngineIdentityMappingStores are immutable after creation", a.id.String(), paths.UnsortedList())
	}

	status := &krm.DiscoveryEngineIdentityMappingStoreStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = discoveryengine.DiscoveryEngineIdentityMappingStoreObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineIdentityMappingStore{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(discoveryengine.DiscoveryEngineIdentityMappingStoreSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.IdentityMappingStore)
	u.SetGroupVersionKind(krm.DiscoveryEngineIdentityMappingStoreGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine identitymappingstore", "name", a.id.String())

	req := &pb.DeleteIdentityMappingStoreRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteIdentityMappingStore(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting discoveryengine identitymappingstore %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted discoveryengine identitymappingstore", "name", a.id.String())

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("waiting for deletion of discoveryengine identitymappingstore %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
