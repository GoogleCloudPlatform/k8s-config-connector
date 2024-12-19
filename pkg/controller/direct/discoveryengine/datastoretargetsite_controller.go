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

package discoveryengine

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineDataStoreTargetSiteGVK, NewDataStoreTargetSiteModel)
}

func NewDataStoreTargetSiteModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataStoreTargetSiteModel{config: *config}, nil
}

var _ directbase.Model = &dataStoreTargetSiteModel{}

type dataStoreTargetSiteModel struct {
	config config.ControllerConfig
}

func (m *dataStoreTargetSiteModel) client(ctx context.Context, projectID string) (*gcp.SiteSearchEngineClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewSiteSearchEngineRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine siteSearchEngine client: %w", err)
	}

	return gcpClient, err
}

func (m *dataStoreTargetSiteModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DiscoveryEngineDataStoreTargetSite{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	dataStoreLink, link, err := krm.NewTargetSiteLinkFromObject(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineDataStoreTargetSiteSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, dataStoreLink.ProjectID)
	if err != nil {
		return nil, err
	}

	return &dataStoreTargetSiteAdapter{
		gcpClient:     gcpClient,
		projectMapper: m.config.ProjectMapper,
		parent:        dataStoreLink,
		link:          link,
		desired:       desired,
	}, nil
}

func (m *dataStoreTargetSiteModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		id, err := krm.ParseTargetSiteExternal(url)
		if err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineDataStoreTargetSite format", "url", url)
		} else {
			gcpClient, err := m.client(ctx, id.ProjectID)
			if err != nil {
				return nil, err
			}
			return &dataStoreTargetSiteAdapter{
				gcpClient:     gcpClient,
				projectMapper: m.config.ProjectMapper,
				parent:        id.DiscoveryEngineDataStoreID,
				link:          id,
			}, nil
		}
	}
	return nil, nil
}

type dataStoreTargetSiteAdapter struct {
	gcpClient     *gcp.SiteSearchEngineClient
	projectMapper *projects.ProjectMapper
	parent        *krm.DiscoveryEngineDataStoreID
	link          *krm.TargetSiteLink
	desired       *pb.TargetSite
	actual        *pb.TargetSite
}

var _ directbase.Adapter = &dataStoreTargetSiteAdapter{}

func (a *dataStoreTargetSiteAdapter) Find(ctx context.Context) (bool, error) {
	if a.link == nil {
		// Server-generated ID and no ID set
		return false, nil
	}
	fqn := a.link.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine targetSite", "name", fqn)

	req := &pb.GetTargetSiteRequest{Name: fqn}
	actual, err := a.gcpClient.GetTargetSite(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine targetSite %q from gcp: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *dataStoreTargetSiteAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// fqn := a.link.String()
	parentFQN := a.parent.String() + "/siteSearchEngine"

	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine targetSite", "parent", parentFQN)

	desired := direct.ProtoClone(a.desired)

	req := &pb.CreateTargetSiteRequest{
		Parent:     parentFQN,
		TargetSite: desired,
	}
	op, err := a.gcpClient.CreateTargetSite(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine targetSite in %q: %w", parentFQN, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of discoveryengine targetSite in %q: %w", parentFQN, err)
	}
	fqn := created.Name

	log.V(2).Info("successfully created discoveryengine targetSite in gcp", "name", fqn)

	status := &krm.DiscoveryEngineDataStoreTargetSiteStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	link, err := krm.ParseTargetSiteExternal(created.Name)
	if err != nil {
		return fmt.Errorf("unexpected name %q for created targetSite: %w", created.Name, err)
	}
	if s, err := a.projectMapper.ReplaceProjectNumberWithID(ctx, link.ProjectID); err != nil {
		return err
	} else {
		link.ProjectID = s
	}

	status.ExternalRef = direct.PtrTo(link.String())

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *dataStoreTargetSiteAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	fqn := a.link.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine targetSite", "name", fqn)

	desired := direct.ProtoClone(a.desired)
	desired.Name = fqn

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.ProvidedUriPattern, a.actual.ProvidedUriPattern) {
		updateMask.Paths = append(updateMask.Paths, "provided_uri_pattern")
	}
	if !reflect.DeepEqual(a.desired.Type, a.actual.Type) {
		updateMask.Paths = append(updateMask.Paths, "type")
	}
	if !reflect.DeepEqual(a.desired.ExactMatch, a.actual.ExactMatch) {
		updateMask.Paths = append(updateMask.Paths, "exact_match")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", fqn)
		return nil
	}

	req := &pb.UpdateTargetSiteRequest{
		TargetSite: desired,
	}
	op, err := a.gcpClient.UpdateTargetSite(ctx, req)
	if err != nil {
		return fmt.Errorf("updating discoveryengine targetSite %q: %w", fqn, err)
	}
	if !op.Done() {
		if _, err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for update of discoveryengine targetSite %q: %w", fqn, err)
		}
	}

	updated, err := a.gcpClient.GetTargetSite(ctx, &pb.GetTargetSiteRequest{Name: fqn})
	if err != nil {
		return fmt.Errorf("getting updated discoveryengine targetSite %q: %w", fqn, err)
	}
	log.V(2).Info("successfully updated discoveryengine targetSite", "name", fqn)

	status := &krm.DiscoveryEngineDataStoreTargetSiteStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *dataStoreTargetSiteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineDataStoreTargetSite{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineDataStoreTargetSiteSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	link, err := krm.ParseTargetSiteExternal(a.actual.Name)
	if err != nil {
		return nil, fmt.Errorf("unexpected name %q for targetSite: %w", a.actual.Name, err)
	}
	obj.Spec.DataStoreRef = &krm.DiscoveryEngineDataStoreRef{
		External: link.DiscoveryEngineDataStoreID.String(),
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(link.TargetSite)
	u.SetGroupVersionKind(krm.DiscoveryEngineDataStoreTargetSiteGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *dataStoreTargetSiteAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	fqn := a.link.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine datastoretargetsite", "name", fqn)

	req := &pb.DeleteTargetSiteRequest{Name: fqn}
	op, err := a.gcpClient.DeleteTargetSite(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting discoveryengine targetSite %q: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted discoveryengine targetSite", "name", fqn)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of discoveryengine targetSite %q: %w", fqn, err)
		}
	}
	return true, nil
}
