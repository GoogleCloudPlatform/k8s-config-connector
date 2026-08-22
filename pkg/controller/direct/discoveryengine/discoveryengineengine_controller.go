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
// proto.service: google.cloud.discoveryengine.v1.EngineService
// proto.message: google.cloud.discoveryengine.v1.Engine
// crd.type: DiscoveryEngineEngine
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineEngineGVK, NewEngineModel)
}

func NewEngineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &engineModel{config: *config}, nil
}

var _ directbase.Model = &engineModel{}

type engineModel struct {
	config config.ControllerConfig
}

func (m *engineModel) client(ctx context.Context, projectID string) (*gcp.EngineClient, error) {
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

	gcpClient, err := gcp.NewEngineRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine engine client: %w", err)
	}

	return gcpClient, err
}

func (m *engineModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineEngine{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewEngineIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineEngineSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Clear the raw/fully-qualified DataStoreIds since we will parse and append the short IDs
	desired.DataStoreIds = nil

	// Manually resolve and map the DataStoreRefs to DataStoreIds
	for _, ref := range obj.Spec.DataStoreRefs {
		normalized, err := ref.NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return nil, fmt.Errorf("resolving dataStoreRef: %w", err)
		}
		dsLink, err := krm.ParseDiscoveryEngineDataStoreExternal(normalized)
		if err != nil {
			return nil, fmt.Errorf("parsing dataStoreRef: %w", err)
		}
		if !krm.IsProjectIDMatch(dsLink.ProjectID, id.Parent().ProjectID) {
			return nil, fmt.Errorf("resolved spec.dataStoreRefs project %q does not match spec.projectRef %q", dsLink.ProjectID, id.Parent().ProjectID)
		}
		if dsLink.Location != id.Parent().Location {
			return nil, fmt.Errorf("resolved spec.dataStoreRefs location %q does not match spec.location %q", dsLink.Location, id.Parent().Location)
		}
		desired.DataStoreIds = append(desired.DataStoreIds, dsLink.DataStore)
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &engineAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *engineModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		trimmed := strings.TrimPrefix(url, "//discoveryengine.googleapis.com/")
		parent, idStr, err := krm.ParseEngineExternal(trimmed)
		if err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineEngine format", "url", url)
			return nil, nil
		}
		obj := &krm.DiscoveryEngineEngine{}
		obj.Spec.Location = parent.Location
		obj.Spec.ProjectRef = &refs.ProjectRef{External: "projects/" + parent.ProjectID}
		obj.Spec.ResourceID = &idStr
		obj.Status.ExternalRef = &trimmed
		id, err := krm.NewEngineIdentity(ctx, nil, obj)
		if err != nil {
			return nil, err
		}
		gcpClient, err := m.client(ctx, parent.ProjectID)
		if err != nil {
			return nil, err
		}
		return &engineAdapter{
			gcpClient: gcpClient,
			id:        id,
		}, nil
	}
	return nil, nil
}

type engineAdapter struct {
	gcpClient *gcp.EngineClient
	id        *krm.EngineIdentity
	desired   *pb.Engine
	actual    *pb.Engine
}

var _ directbase.Adapter = &engineAdapter{}

func (a *engineAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/collections/default_collection/engines/%s", a.id.Parent().ProjectID, a.id.Parent().Location, a.id.ID())
}

func (a *engineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.fullyQualifiedName()
	log.V(2).Info("getting discoveryengine engine", "name", fqn)

	req := &pb.GetEngineRequest{Name: fqn}
	actual, err := a.gcpClient.GetEngine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine engine %q from gcp: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *engineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.fullyQualifiedName()
	log.V(2).Info("creating discoveryengine engine", "name", fqn)

	desired := proto.CloneOf(a.desired)
	desired.Name = fqn

	parent := fmt.Sprintf("projects/%s/locations/%s/collections/default_collection", a.id.Parent().ProjectID, a.id.Parent().Location)
	req := &pb.CreateEngineRequest{
		Parent:   parent,
		Engine:   desired,
		EngineId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateEngine(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine engine %s: %w", fqn, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for discoveryengine engine %s creation: %w", fqn, err)
	}
	log.V(2).Info("successfully created discoveryengine engine in gcp", "name", fqn)

	return a.updateStatus(ctx, createOp, created)
}

func (a *engineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.fullyQualifiedName()
	log.V(2).Info("updating discoveryengine engine", "name", fqn)

	diffs, updateMask, err := compareEngine(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.CloneOf(a.desired)
		desired.Name = fqn

		req := &pb.UpdateEngineRequest{
			Engine:     desired,
			UpdateMask: updateMask,
		}
		updated, err := a.gcpClient.UpdateEngine(ctx, req)
		if err != nil {
			return fmt.Errorf("updating discoveryengine engine %s: %w", fqn, err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareEngine(ctx context.Context, actual, desired *pb.Engine) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DiscoveryEngineEngineSpec_FromProto, DiscoveryEngineEngineSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore name if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Engine) {
		// populate any defaults if necessary
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *engineAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Engine) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DiscoveryEngineEngineStatus{}
	status.ObservedState = DiscoveryEngineEngineObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *engineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("export is not implemented for DiscoveryEngineEngine")
}

// Delete implements the Adapter interface.
func (a *engineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.fullyQualifiedName()
	log.V(2).Info("deleting discoveryengine engine", "name", fqn)

	req := &pb.DeleteEngineRequest{Name: fqn}
	op, err := a.gcpClient.DeleteEngine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting discoveryengine engine %s: %w", fqn, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return false, fmt.Errorf("waiting for discoveryengine engine %s deletion: %w", fqn, err)
		}
	}
	log.V(2).Info("successfully deleted discoveryengine engine", "name", fqn)

	return true, nil
}
