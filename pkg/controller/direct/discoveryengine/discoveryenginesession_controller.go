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

// +tool:controller
// proto.service: google.cloud.discoveryengine.v1.ConversationalSearchService
// proto.message: google.cloud.discoveryengine.v1.Session
// crd.type: DiscoveryEngineSession
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineSessionGVK, NewSessionModel)
}

func NewSessionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sessionModel{config: *config}, nil
}

var _ directbase.Model = &sessionModel{}

type sessionModel struct {
	config config.ControllerConfig
}

func (m *sessionModel) client(ctx context.Context, projectID string) (*gcp.ConversationalSearchClient, error) {
	var opts []option.ClientOption

	config := m.config

	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewConversationalSearchRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine conversational search client: %w", err)
	}

	return gcpClient, err
}

func (m *sessionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineSession{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Resolve the parent datastore ref to get the correct collection ID
	if obj.Spec.DataStoreRef == nil {
		return nil, fmt.Errorf("spec.dataStoreRef is not set")
	}
	dataStoreRefCopy := *obj.Spec.DataStoreRef
	normalizedDataStore, err := dataStoreRefCopy.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("resolving spec.dataStoreRef: %w", err)
	}
	dataStoreLink, err := krm.ParseDiscoveryEngineDataStoreExternal(normalizedDataStore)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.dataStoreRef external: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DiscoveryEngineSessionIdentity)

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineSessionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &sessionAdapter{
		gcpClient:  gcpClient,
		id:         id,
		collection: dataStoreLink.Collection,
		desired:    desired,
	}, nil
}

func (m *sessionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		trimmed := strings.TrimPrefix(url, "//discoveryengine.googleapis.com/")
		id := &krm.DiscoveryEngineSessionIdentity{}
		if err := id.FromExternal(trimmed); err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineSession format", "url", url)
			return nil, nil
		}
		// Since we don't have the collection in the URL, we default it to "default_collection",
		// which is standard for single-tenant or default setups in Vertex AI Search.
		collection := "default_collection"
		gcpClient, err := m.client(ctx, id.Project)
		if err != nil {
			return nil, err
		}
		return &sessionAdapter{
			gcpClient:  gcpClient,
			id:         id,
			collection: collection,
		}, nil
	}
	return nil, nil
}

type sessionAdapter struct {
	gcpClient  *gcp.ConversationalSearchClient
	id         *krm.DiscoveryEngineSessionIdentity
	collection string
	desired    *pb.Session
	actual     *pb.Session
}

var _ directbase.Adapter = &sessionAdapter{}

func (a *sessionAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s/sessions/%s", a.id.Project, a.id.Location, a.collection, a.id.Datastore, a.id.Session)
}

func (a *sessionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.fullyQualifiedName()
	log.V(2).Info("getting discoveryengine session", "name", fqn)

	req := &pb.GetSessionRequest{Name: fqn}
	actual, err := a.gcpClient.GetSession(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine session %q from gcp: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *sessionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine session", "name", a.id)

	desired := proto.CloneOf(a.desired)
	desired.Name = a.fullyQualifiedName()

	req := &pb.CreateSessionRequest{
		Parent:  fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", a.id.Project, a.id.Location, a.collection, a.id.Datastore),
		Session: desired,
	}
	created, err := a.gcpClient.CreateSession(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine session %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created discoveryengine session in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *sessionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine session", "name", a.id)

	diffs, updateMask, err := compareSession(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.CloneOf(a.desired)
		desired.Name = a.fullyQualifiedName()

		req := &pb.UpdateSessionRequest{
			Session:    desired,
			UpdateMask: updateMask,
		}
		updated, err := a.gcpClient.UpdateSession(ctx, req)
		if err != nil {
			return fmt.Errorf("updating discoveryengine session %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareSession(ctx context.Context, actual, desired *pb.Session) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DiscoveryEngineSessionSpec_FromProto, DiscoveryEngineSessionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore name if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Session) {
		// populate any defaults if necessary
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *sessionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Session) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DiscoveryEngineSessionStatus{}
	status.ObservedState = DiscoveryEngineSessionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *sessionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineSession{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineSessionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.DataStoreRef = &krm.DiscoveryEngineDataStoreRef{
		External: fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", a.id.Project, a.id.Location, a.collection, a.id.Datastore),
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Session)
	u.SetGroupVersionKind(krm.DiscoveryEngineSessionGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *sessionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.fullyQualifiedName()
	log.V(2).Info("deleting discoveryengine session", "name", fqn)

	req := &pb.DeleteSessionRequest{Name: fqn}
	err := a.gcpClient.DeleteSession(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting discoveryengine session %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted discoveryengine session", "name", a.id)

	return true, nil
}
