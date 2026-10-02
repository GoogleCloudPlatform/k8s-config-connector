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

package visionai

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/visionai/apiv1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VisionAIApplicationGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.AppPlatformClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAppPlatformRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AppPlatform client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VisionAIApplication{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := VisionAIApplicationSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	return &ApplicationAdapter{
		id:        identity.(*krm.VisionAIApplicationIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.VisionAIApplicationIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ApplicationAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type ApplicationAdapter struct {
	id        *krm.VisionAIApplicationIdentity
	gcpClient *gcp.AppPlatformClient
	desired   *pb.Application
	actual    *pb.Application
}

var _ directbase.Adapter = &ApplicationAdapter{}

func (a *ApplicationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VisionAIApplication", "name", a.id.String())

	req := &pb.GetApplicationRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetApplication(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VisionAIApplication %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *ApplicationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating VisionAIApplication", "id", fqn)

	req := &pb.CreateApplicationRequest{
		Parent:        a.id.ParentString(),
		Application:   a.desired,
		ApplicationId: a.id.Application,
	}
	op, err := a.gcpClient.CreateApplication(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VisionAIApplication %s: %w", fqn, err)
	}
	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting VisionAIApplication %s creation: %w", fqn, err)
	}
	log.V(2).Info("successfully created VisionAIApplication", "name", fqn)

	created, err := a.gcpClient.GetApplication(ctx, &pb.GetApplicationRequest{Name: fqn})
	if err != nil {
		return fmt.Errorf("getting created VisionAIApplication %s: %w", fqn, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ApplicationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating VisionAIApplication", "name", fqn)

	diffs, updateMask, err := compareVisionAIApplication(ctx, updateOp.GetUnstructured(), a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.Clone(a.desired).(*pb.Application)
		desiredCopy.Name = fqn

		req := &pb.UpdateApplicationRequest{
			Application: desiredCopy,
			UpdateMask:  updateMask,
		}

		op, err := a.gcpClient.UpdateApplication(ctx, req)
		if err != nil {
			return fmt.Errorf("updating VisionAIApplication %s: %w", fqn, err)
		}
		if _, err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting VisionAIApplication %s update: %w", fqn, err)
		}

		latest, err = a.gcpClient.GetApplication(ctx, &pb.GetApplicationRequest{Name: fqn})
		if err != nil {
			return fmt.Errorf("getting updated VisionAIApplication %s: %w", fqn, err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ApplicationAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Application) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VisionAIApplicationStatus{}
	status.ObservedState = VisionAIApplicationObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := latest.GetName()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *ApplicationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VisionAIApplication{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VisionAIApplicationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.Application)
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Application)
	u.SetGroupVersionKind(krm.VisionAIApplicationGVK)

	return u, nil
}

func (a *ApplicationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting VisionAIApplication", "name", fqn)

	req := &pb.DeleteApplicationRequest{Name: fqn}
	op, err := a.gcpClient.DeleteApplication(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent VisionAIApplication, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting VisionAIApplication %s: %w", fqn, err)
	}
	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting VisionAIApplication %s deletion: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted VisionAIApplication", "name", fqn)

	return true, nil
}

func compareVisionAIApplication(ctx context.Context, u *unstructured.Unstructured, actual, desired *pb.Application) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VisionAIApplicationSpec_FromProto, VisionAIApplicationSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.Clone(desired).(*pb.Application)

	if u != nil && actual.GetUpdateTime() != nil {
		obj := &krm.VisionAIApplication{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err == nil {
			ready := false
			for _, condition := range obj.Status.Conditions {
				if condition.Type == "Ready" && condition.Status == "True" {
					ready = true
				}
			}
			obsGen, _, _ := unstructured.NestedInt64(u.Object, "status", "observedGeneration")
			mapCtx := &direct.MapContext{}
			gcpUpdateTime := direct.StringTimestamp_FromProto(mapCtx, actual.GetUpdateTime())
			if ready && gcpUpdateTime != nil && obj.Status.ObservedState != nil && obj.Status.ObservedState.UpdateTime != nil &&
				*gcpUpdateTime == *obj.Status.ObservedState.UpdateTime &&
				u.GetGeneration() == obsGen {
				if clonedDesired.ApplicationConfigs != nil && clonedDesired.ApplicationConfigs.EventDeliveryConfig != nil {
					if maskedActual.ApplicationConfigs == nil {
						maskedActual.ApplicationConfigs = &pb.ApplicationConfigs{}
					}
					if maskedActual.ApplicationConfigs.EventDeliveryConfig == nil {
						maskedActual.ApplicationConfigs.EventDeliveryConfig = clonedDesired.ApplicationConfigs.EventDeliveryConfig
					}
				}
			}
		}
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
