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

package vertexai

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewJobRESTClient(ctx, opts...)
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
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewVertexAICustomJobIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err := m.normalize(ctx, reader, obj); err != nil {
		return nil, err
	}

	// Get vertexai GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &CustomJobAdapter{
		id:            id,
		gcpClient:     gcpClient,
		desired:       obj,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

func (m *modelCustomJob) normalize(ctx context.Context, reader client.Reader, obj *krm.VertexAICustomJob) error {
	if obj.Spec.EncryptionSpec != nil && obj.Spec.EncryptionSpec.KMSKeyRef != nil {
		resolved, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, obj, obj.Spec.EncryptionSpec.KMSKeyRef)
		if err != nil {
			return err
		}
		obj.Spec.EncryptionSpec.KMSKeyRef = resolved
	}

	if obj.Spec.JobSpec != nil {
		if obj.Spec.JobSpec.ServiceAccountRef != nil {
			if err := obj.Spec.JobSpec.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
				return err
			}
		}
		if obj.Spec.JobSpec.NetworkRef != nil {
			if err := obj.Spec.JobSpec.NetworkRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
		if obj.Spec.JobSpec.TensorboardRef != nil {
			if err := obj.Spec.JobSpec.TensorboardRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
		if obj.Spec.JobSpec.ExperimentRef != nil {
			if err := obj.Spec.JobSpec.ExperimentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
		if obj.Spec.JobSpec.ExperimentRunRef != nil {
			if err := obj.Spec.JobSpec.ExperimentRunRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
		for i := range obj.Spec.JobSpec.ModelRefs {
			if err := obj.Spec.JobSpec.ModelRefs[i].Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *modelCustomJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type CustomJobAdapter struct {
	id            *krm.VertexAICustomJobIdentity
	gcpClient     *gcp.JobClient
	desired       *krm.VertexAICustomJob
	actual        *pb.CustomJob
	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &CustomJobAdapter{}

func (a *CustomJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CustomJob", "name", a.id)

	req := &pb.GetCustomJobRequest{Name: a.id.String()}
	customjobpb, err := a.gcpClient.GetCustomJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CustomJob %q: %w", a.id, err)
	}

	a.actual = customjobpb
	return true, nil
}

func (a *CustomJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CustomJob", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	if desired.Spec.JobSpec != nil && desired.Spec.JobSpec.NetworkRef != nil {
		if err := desired.Spec.JobSpec.NetworkRef.ConvertToProjectNumber(ctx, a.projectMapper); err != nil {
			return err
		}
	}

	resource := VertexAICustomJobSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateCustomJobRequest{
		Parent:    a.id.ParentString(),
		CustomJob: resource,
	}
	created, err := a.gcpClient.CreateCustomJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CustomJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CustomJob", "name", a.id)

	status := &krm.VertexAICustomJobStatus{}
	status.ObservedState = VertexAICustomJobObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.GetName())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *CustomJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// CustomJob is immutable.
	// We can check if the spec changed and return an error, but usually we just skip update.
	// However, it's better to report if there's a drift.
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CustomJob (immutable, so check for drift)", "name", a.id)

	mapCtx := &direct.MapContext{}

	// For now, let's just update the status with the latest state.
	status := &krm.VertexAICustomJobStatus{}
	status.ObservedState = VertexAICustomJobObservedState_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.actual.GetName())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *CustomJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAICustomJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAICustomJobSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)
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
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent CustomJob, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CustomJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CustomJob", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete CustomJob %s: %w", a.id, err)
	}
	return true, nil
}
