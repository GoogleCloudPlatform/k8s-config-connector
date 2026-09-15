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

package run

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/run/apiv2"
	pb "cloud.google.com/go/run/apiv2/runpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.RunWorkerPoolGVK, NewWorkerPoolModel)
}

func NewWorkerPoolModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWorkerPool{config: *config}, nil
}

var _ directbase.Model = &modelWorkerPool{}

type modelWorkerPool struct {
	config config.ControllerConfig
}

func (m *modelWorkerPool) client(ctx context.Context) (*gcp.WorkerPoolsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewWorkerPoolsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building WorkerPool client: %w", err)
	}
	return gcpClient, err
}

func (m *modelWorkerPool) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.RunWorkerPool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.RunWorkerPoolIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
	if err := ResolveRunWorkerPoolRefs(ctx, reader, obj); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	copied := obj.DeepCopy()
	desired := RunWorkerPoolSpec_v1alpha1_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WorkerPoolAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelWorkerPool) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if s, ok := strings.CutPrefix(url, "//run.googleapis.com/"); ok {
		s = strings.TrimPrefix(s, "v2/")

		var id krm.RunWorkerPoolIdentity
		if err := id.FromExternal(s); err != nil {
			log.V(2).Error(err, "url did not match RunWorkerPool format", "url", url)
			return nil, nil
		}

		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}
		return &WorkerPoolAdapter{
			gcpClient: gcpClient,
			id:        &id,
		}, nil
	}
	return nil, nil
}

type WorkerPoolAdapter struct {
	id        *krm.RunWorkerPoolIdentity
	gcpClient *gcp.WorkerPoolsClient
	desired   *pb.WorkerPool
	actual    *pb.WorkerPool
}

var _ directbase.Adapter = &WorkerPoolAdapter{}

func (a *WorkerPoolAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting WorkerPool", "name", a.id.String())

	req := &pb.GetWorkerPoolRequest{Name: a.id.String()}
	found, err := a.gcpClient.GetWorkerPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting WorkerPool %q: %w", a.id.String(), err)
	}

	a.actual = found
	return true, nil
}

func (a *WorkerPoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating WorkerPool", "name", a.id.String())

	desired := proto.Clone(a.desired).(*pb.WorkerPool)
	desired.Name = ""

	req := &pb.CreateWorkerPoolRequest{
		Parent:       a.id.ParentString(),
		WorkerPool:   desired,
		WorkerPoolId: a.id.WorkerPool,
	}
	op, err := a.gcpClient.CreateWorkerPool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating WorkerPool %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of WorkerPool %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created WorkerPool", "name", a.id.String())

	latest, err := a.gcpClient.GetWorkerPool(ctx, &pb.GetWorkerPoolRequest{Name: created.GetName()})
	if err == nil {
		created = latest
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *WorkerPoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating WorkerPool", "name", a.id.String())

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareWorkerPool(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateWorkerPoolRequest{
			WorkerPool: a.desired,
			UpdateMask: updateMask,
		}
		op, err := a.gcpClient.UpdateWorkerPool(ctx, req)
		if err != nil {
			return fmt.Errorf("updating WorkerPool %s: %w", a.id.String(), err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of WorkerPool %q: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated WorkerPool", "name", a.id.String())
		latest = updated

		latestGet, err := a.gcpClient.GetWorkerPool(ctx, &pb.GetWorkerPoolRequest{Name: latest.GetName()})
		if err == nil {
			latest = latestGet
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareWorkerPool(ctx context.Context, actual, desired *pb.WorkerPool) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, RunWorkerPoolSpec_v1alpha1_FromProto, RunWorkerPoolSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.WorkerPool)

	populateDefaults := func(obj *pb.WorkerPool) {
		// Populate GCP/server defaults here if needed
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	// If desired leaves optional fields unspecified, inherit existing actual values to avoid false diffs.
	if maskedActual != nil {
		if clonedDesired.LaunchStage == 0 && maskedActual.LaunchStage != 0 {
			clonedDesired.LaunchStage = maskedActual.LaunchStage
		}
		if clonedDesired.Scaling == nil && maskedActual.Scaling != nil {
			clonedDesired.Scaling = maskedActual.Scaling
		}
		if len(clonedDesired.InstanceSplits) == 0 && len(maskedActual.InstanceSplits) > 0 {
			clonedDesired.InstanceSplits = maskedActual.InstanceSplits
		}
		if clonedDesired.Template != nil && maskedActual.Template != nil {
			if clonedDesired.Template.ServiceAccount == "" && maskedActual.Template.ServiceAccount != "" {
				clonedDesired.Template.ServiceAccount = maskedActual.Template.ServiceAccount
			}
			if clonedDesired.Template.Revision == "" && maskedActual.Template.Revision != "" {
				clonedDesired.Template.Revision = maskedActual.Template.Revision
			}
			actualContainersByName := make(map[string]*pb.Container)
			for _, actCont := range maskedActual.Template.Containers {
				if actCont.GetName() != "" {
					actualContainersByName[actCont.GetName()] = actCont
				}
			}
			for i, desCont := range clonedDesired.Template.Containers {
				var actCont *pb.Container
				if desCont.GetName() != "" {
					actCont = actualContainersByName[desCont.GetName()]
				}
				if actCont == nil && i < len(maskedActual.Template.Containers) {
					actCont = maskedActual.Template.Containers[i]
				}
				if actCont != nil {
					if desCont.Resources == nil && actCont.Resources != nil {
						desCont.Resources = actCont.Resources
					} else if desCont.Resources != nil && actCont.Resources != nil {
						if desCont.Resources.Limits == nil && actCont.Resources.Limits != nil {
							desCont.Resources.Limits = actCont.Resources.Limits
						} else if desCont.Resources.Limits != nil && actCont.Resources.Limits != nil {
							for k, v := range actCont.Resources.Limits {
								if _, exists := desCont.Resources.Limits[k]; !exists {
									desCont.Resources.Limits[k] = v
								}
							}
						}
					}
				}
			}
		}
	}

	// CustomAudiences is mutable-but-unreadable (missing from the GET response).
	// If the generation matches the observed generation, we can safely copy CustomAudiences
	// from clonedDesired to maskedActual to avoid false diffs.
	if actual != nil && actual.GetGeneration() == actual.GetObservedGeneration() {
		if len(clonedDesired.CustomAudiences) > 0 && len(maskedActual.CustomAudiences) == 0 {
			maskedActual.CustomAudiences = clonedDesired.CustomAudiences
		}
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *WorkerPoolAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.WorkerPool) error {
	mapCtx := &direct.MapContext{}
	status := &krm.RunWorkerPoolStatus{}
	status.ObservedState = RunWorkerPoolObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *WorkerPoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.RunWorkerPool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(RunWorkerPoolSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.WorkerPool)
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.WorkerPool)
	u.SetGroupVersionKind(krm.RunWorkerPoolGVK)

	return u, nil
}

func (a *WorkerPoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting WorkerPool", "name", a.id.String())

	req := &pb.DeleteWorkerPoolRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteWorkerPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent WorkerPool, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting WorkerPool %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted WorkerPool", "name", a.id.String())

	if _, err = op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting delete WorkerPool %s: %w", a.id.String(), err)
	}
	return true, nil
}
