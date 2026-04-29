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

package clouddeploy

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/deploy/apiv1"
	clouddeploypb "cloud.google.com/go/deploy/apiv1/deploypb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CloudDeployTargetGVK, NewTargetModel)
}

func NewTargetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTarget{config: *config}, nil
}

var _ directbase.Model = &modelTarget{}

type modelTarget struct {
	config config.ControllerConfig
}

func (m *modelTarget) client(ctx context.Context) (*gcp.CloudDeployClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudDeployRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Target client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelTarget) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudDeployTarget{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get clouddeploy GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &TargetAdapter{
		id:        id.(*krm.CloudDeployTargetIdentity),
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
		labels:    u.GetLabels(),
	}, nil
}

func (m *modelTarget) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type TargetAdapter struct {
	id        *krm.CloudDeployTargetIdentity
	gcpClient *gcp.CloudDeployClient
	desiredPb *clouddeploypb.Target
	actual    *clouddeploypb.Target

	desired *krm.CloudDeployTarget
	reader  client.Reader
	labels  map[string]string
}

func (a *TargetAdapter) resolveReferences(ctx context.Context) error {
	obj := a.desired
	reader := a.reader

	// Normalize all reference fields
	if obj.Spec.Gke != nil && obj.Spec.Gke.ClusterRef != nil {
		if _, err := obj.Spec.Gke.ClusterRef.NormalizedExternal(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}
	if obj.Spec.AnthosCluster != nil && obj.Spec.AnthosCluster.MembershipRef != nil {
		if err := obj.Spec.AnthosCluster.MembershipRef.Normalize(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}
	if obj.Spec.MultiTarget != nil {
		for i := range obj.Spec.MultiTarget.TargetRefs {
			if err := obj.Spec.MultiTarget.TargetRefs[i].Normalize(ctx, reader, obj.Namespace); err != nil {
				return err
			}
		}
	}
	if obj.Spec.CustomTarget != nil && obj.Spec.CustomTarget.CustomTargetTypeRef != nil {
		if _, err := obj.Spec.CustomTarget.CustomTargetTypeRef.NormalizedExternal(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}
	for i := range obj.Spec.ExecutionConfigs {
		ec := &obj.Spec.ExecutionConfigs[i]
		if ec.DefaultPool != nil && ec.DefaultPool.ServiceAccountRef != nil {
			if err := ec.DefaultPool.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
				return err
			}
		}
		if ec.PrivatePool != nil {
			if ec.PrivatePool.ServiceAccountRef != nil {
				if err := ec.PrivatePool.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
					return err
				}
			}
			if ec.PrivatePool.WorkerPoolRef != nil {
				if err := refsv1beta1.Normalize(ctx, reader, ec.PrivatePool.WorkerPoolRef, obj.Namespace); err != nil {
					return err
				}
			}
		}
		if ec.WorkerPoolRef != nil {
			if err := refsv1beta1.Normalize(ctx, reader, ec.WorkerPoolRef, obj.Namespace); err != nil {
				return err
			}
		}
		if ec.ServiceAccountRef != nil {
			if err := ec.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
				return err
			}
		}
	}

	mapCtx := &direct.MapContext{}
	a.desiredPb = CloudDeployTargetSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	a.desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(a.labels)
	return nil
}

var _ directbase.Adapter = &TargetAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TargetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Target", "name", a.id.String())

	req := &clouddeploypb.GetTargetRequest{Name: a.id.String()}
	targetpb, err := a.gcpClient.GetTarget(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Target %q: %w", a.id.String(), err)
	}

	a.actual = targetpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TargetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Target", "name", a.id.String())

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}

	req := &clouddeploypb.CreateTargetRequest{
		Parent:   a.id.Parent().String(),
		Target:   a.desiredPb,
		TargetId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateTarget(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Target %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Target %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created Target", "name", a.id.String())

	status := &krm.CloudDeployTargetStatus{}
	status.ObservedState = CloudDeployTargetObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		log.Error(mapCtx.Err(), "error mapping Target status")
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TargetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Target", "name", a.id.String())

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}

	a.desiredPb.Name = a.id.String()

	// Preserve system labels (goog- or go-)
	if a.actual.Labels != nil {
		if a.desiredPb.Labels == nil {
			a.desiredPb.Labels = make(map[string]string)
		}
		for k, v := range a.actual.Labels {
			if strings.HasPrefix(k, "goog-") || strings.HasPrefix(k, "go-") {
				a.desiredPb.Labels[k] = v
			}
		}
	}

	// etag is server-generated, but we use it for optimistic concurrency.
	// We skip the diff when it shows up in path to avoid unnecessary drift.
	paths, err := common.CompareProtoMessage(a.desiredPb, a.actual, func(fieldName protoreflect.Name, a, b proto.Message) (bool, error) {
		if fieldName == "etag" {
			return false, nil
		}
		// If execution_configs is empty in the desired state, we check if the actual state
		// matches the server-side default (a single execution config with a default pool).
		// If it does, we skip the diff to avoid re-reconciliation loops.
		if fieldName == "execution_configs" {
			desired := a.(*clouddeploypb.Target)
			if len(desired.ExecutionConfigs) == 0 {
				actual := b.(*clouddeploypb.Target)
				if len(actual.ExecutionConfigs) == 1 && actual.ExecutionConfigs[0].GetDefaultPool() != nil {
					return false, nil
				}
			}
		}
		return common.BasicDiff(fieldName, a, b)
	})
	if err != nil {
		return err
	}

	updated := a.actual
	if len(paths) != 0 {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)

		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		// Inject the latest etag for optimistic concurrency
		a.desiredPb.Etag = a.actual.Etag

		req := &clouddeploypb.UpdateTargetRequest{
			UpdateMask: updateMask,
			Target:     a.desiredPb,
		}
		op, err := a.gcpClient.UpdateTarget(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Target %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Target %s waiting update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated Target", "name", a.id.String())
	} else {
		log.V(2).Info("no field needs update", "name", a.id.String())
	}

	status := &krm.CloudDeployTargetStatus{}
	status.ObservedState = CloudDeployTargetObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		log.Error(mapCtx.Err(), "error mapping Target status")
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TargetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudDeployTarget{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudDeployTargetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = direct.LazyPtr(a.id.Parent().Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(k8s.ValueToDNSSubdomainName(a.id.ID()))
	u.SetGroupVersionKind(krm.CloudDeployTargetGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TargetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Target", "name", a.id.String())

	req := &clouddeploypb.DeleteTargetRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteTarget(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Target, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Target %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete wait for non-existent Target, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Target %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Target", "name", a.id.String())
	return true, nil
}
