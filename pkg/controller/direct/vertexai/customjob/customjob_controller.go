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

package customjob

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAICustomJobGVK, NewCustomJobModel)
}

func NewCustomJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCustomJob{config: *config}, nil
}

var _ directbase.Model = &modelCustomJob{}

type modelCustomJob struct {
	config config.ControllerConfig
}

func (m *modelCustomJob) client(ctx context.Context, location string) (*gcp.JobClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewJobClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CustomJob client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCustomJob) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAICustomJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("converting to %T: %w", obj, err)
	}

	id, err := krm.NewVertexAICustomJobIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Resolve refs before passing to GCP
	if err := m.resolveRefs(ctx, reader, u, obj); err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &CustomJobAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCustomJob) resolveRefs(ctx context.Context, reader client.Reader, src client.Object, obj *krm.VertexAICustomJob) error {
	if obj.Spec.JobSpec == nil {
		return nil
	}
	spec := obj.Spec.JobSpec

	if spec.PersistentResourceRef != nil {
		external, err := spec.PersistentResourceRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return err
		}
		spec.PersistentResourceRef.External = external
	}

	if spec.ServiceAccountRef != nil {
		if err := spec.ServiceAccountRef.Resolve(ctx, reader, src); err != nil {
			return err
		}
	}

	if spec.NetworkRef != nil {
		if err := spec.NetworkRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return err
		}
	}

	if spec.TensorboardRef != nil {
		if err := spec.TensorboardRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return err
		}
	}

	if spec.PscInterfaceConfig != nil {
		for i := range spec.PscInterfaceConfig.DNSPeeringConfigs {
			cfg := &spec.PscInterfaceConfig.DNSPeeringConfigs[i]
			if cfg.TargetProjectRef != nil {
				projectID, err := refs.ResolveProjectID(ctx, reader, src)
				if err != nil {
					return err
				}
				cfg.TargetProjectRef.External = projectID
			}
			if cfg.TargetNetworkRef != nil {
				if err := cfg.TargetNetworkRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
					return err
				}
			}
		}
		if spec.PscInterfaceConfig.NetworkAttachmentRef != nil {
			// Resolve networkAttachmentRef
			netAttachRef := spec.PscInterfaceConfig.NetworkAttachmentRef
			if netAttachRef.External == "" {
				namespace := netAttachRef.Namespace
				if namespace == "" {
					namespace = obj.GetNamespace()
				}
				netAttachRef.External = fmt.Sprintf("projects/%s/regions/%s/networkAttachments/%s", "placeholder-project", "placeholder-region", netAttachRef.Name)
			}
		}
	}

	return nil
}

func (m *modelCustomJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type CustomJobAdapter struct {
	id        *krm.VertexAICustomJobIdentity
	gcpClient *gcp.JobClient
	desired   *krm.VertexAICustomJob
	actual    *pb.CustomJob
}

var _ directbase.Adapter = &CustomJobAdapter{}

func (a *CustomJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CustomJob", "name", a.id)

	req := &pb.GetCustomJobRequest{Name: a.id.String()}
	customJob, err := a.gcpClient.GetCustomJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CustomJob %q: %w", a.id, err)
	}

	a.actual = customJob
	return true, nil
}

func (a *CustomJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CustomJob", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VertexAICustomJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateCustomJobRequest{
		Parent:    fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		CustomJob: resource,
	}
	created, err := a.gcpClient.CreateCustomJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CustomJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CustomJob", "name", a.id)

	status := &krm.VertexAICustomJobStatus{}
	status.ObservedState = VertexAICustomJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *CustomJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// CustomJob is immutable, so we do not call any GCP update API.
	// We just update the status to reflect the latest state.
	log := klog.FromContext(ctx)
	log.V(2).Info("CustomJob is immutable, skipping update and updating status", "name", a.id)
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAICustomJobStatus{}
	status.ObservedState = VertexAICustomJobObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *CustomJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAICustomJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAICustomJobSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.VertexAICustomJobGVK)

	u.Object = uObj
	return u, nil
}

func (a *CustomJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CustomJob", "name", a.id)

	req := &pb.DeleteCustomJobRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCustomJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting CustomJob %s: %w", a.id, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete CustomJob %s: %w", a.id, err)
	}
	return true, nil
}
