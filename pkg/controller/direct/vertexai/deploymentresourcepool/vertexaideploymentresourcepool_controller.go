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

package deploymentresourcepool

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.VertexAIDeploymentResourcePoolGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context, location string) (*gcp.DeploymentResourcePoolClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewDeploymentResourcePoolRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("error building DeploymentResourcePool client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAIDeploymentResourcePool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.VertexAIDeploymentResourcePoolIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if obj.Spec.ServiceAccountRef != nil {
		if err := obj.Spec.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
			return nil, fmt.Errorf("resolving serviceAccountRef: %w", err)
		}
	}

	// Get vertexai GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := VertexAIDeploymentResourcePoolSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.VertexAIDeploymentResourcePoolIdentity
	gcpClient *gcp.DeploymentResourcePoolClient
	desired   *pb.DeploymentResourcePool
	actual    *pb.DeploymentResourcePool
	reader    client.Reader
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAIDeploymentResourcePool", "name", a.id)

	req := &pb.GetDeploymentResourcePoolRequest{Name: a.id.String()}
	resp, err := a.gcpClient.GetDeploymentResourcePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAIDeploymentResourcePool %q: %w", a.id, err)
	}

	a.actual = resp
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAIDeploymentResourcePool", "name", a.id)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateDeploymentResourcePoolRequest{
		Parent:                   parent,
		DeploymentResourcePool:   a.desired,
		DeploymentResourcePoolId: a.id.DeploymentResourcePool,
	}
	op, err := a.gcpClient.CreateDeploymentResourcePool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAIDeploymentResourcePool %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting VertexAIDeploymentResourcePool %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created VertexAIDeploymentResourcePool", "name", a.id)

	// Fetch fully-populated resource after creation
	latest, err := a.gcpClient.GetDeploymentResourcePool(ctx, &pb.GetDeploymentResourcePoolRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting VertexAIDeploymentResourcePool %s after creation: %w", a.id, err)
	}

	created.SatisfiesPzi = latest.SatisfiesPzi
	created.SatisfiesPzs = latest.SatisfiesPzs

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAIDeploymentResourcePool", "name", a.id)

	diffs, updateMask, err := a.compare(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.DeploymentResourcePool)
		desired.Name = a.id.String()

		req := &pb.UpdateDeploymentResourcePoolRequest{
			DeploymentResourcePool: desired,
			UpdateMask:             updateMask,
		}
		op, err := a.gcpClient.UpdateDeploymentResourcePool(ctx, req)
		if err != nil {
			return fmt.Errorf("updating VertexAIDeploymentResourcePool %s: %w", a.id, err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting VertexAIDeploymentResourcePool %s update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated VertexAIDeploymentResourcePool", "name", a.id)

		// Fetch fully-populated resource after update
		latest, err = a.gcpClient.GetDeploymentResourcePool(ctx, &pb.GetDeploymentResourcePoolRequest{Name: a.id.String()})
		if err != nil {
			return fmt.Errorf("getting VertexAIDeploymentResourcePool %s after update: %w", a.id, err)
		}
		updated.SatisfiesPzi = latest.SatisfiesPzi
		updated.SatisfiesPzs = latest.SatisfiesPzs
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) compare(ctx context.Context, actual, desired *pb.DeploymentResourcePool) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VertexAIDeploymentResourcePoolSpec_FromProto, VertexAIDeploymentResourcePoolSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.DeploymentResourcePool) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAIDeploymentResourcePoolStatus{}
	status.ObservedState = VertexAIDeploymentResourcePoolObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAIDeploymentResourcePool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIDeploymentResourcePoolSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VertexAIDeploymentResourcePoolGVK)

	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAIDeploymentResourcePool", "name", a.id)

	req := &pb.DeleteDeploymentResourcePoolRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDeploymentResourcePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent VertexAIDeploymentResourcePool, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting VertexAIDeploymentResourcePool %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted VertexAIDeploymentResourcePool", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete VertexAIDeploymentResourcePool %s: %w", a.id, err)
	}
	return true, nil
}
