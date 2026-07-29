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

package gkemulticloud

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/gkemulticloud/apiv1"
	pb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkemulticloud/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

const (
	ctrlName = "gkemulticloudattachedcluster-controller"
)

func init() {
	registry.RegisterModel(krm.GKEMulticloudAttachedClusterGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.AttachedClustersClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAttachedClustersClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AttachedClusters client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.GKEMulticloudAttachedCluster{}
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
	id := identity.(*krm.GKEMulticloudAttachedClusterIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := GKEMulticloudAttachedClusterSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.Name = id.String()

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredProto,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.GKEMulticloudAttachedClusterIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type Adapter struct {
	id        *krm.GKEMulticloudAttachedClusterIdentity
	gcpClient *gcp.AttachedClustersClient
	desired   *pb.AttachedCluster
	actual    *pb.AttachedCluster
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting GKEMulticloudAttachedCluster", "name", a.id.String())

	req := &pb.GetAttachedClusterRequest{Name: a.id.String()}
	attachedClusterpb, err := a.gcpClient.GetAttachedCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GKEMulticloudAttachedCluster %q: %w", a.id.String(), err)
	}

	a.actual = attachedClusterpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GKEMulticloudAttachedCluster", "name", a.id.String())

	req := &pb.CreateAttachedClusterRequest{
		Parent:            a.id.ParentString(),
		AttachedClusterId: a.id.AttachedCluster,
		AttachedCluster:   a.desired,
	}
	op, err := a.gcpClient.CreateAttachedCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating GKEMulticloudAttachedCluster %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("GKEMulticloudAttachedCluster %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created GKEMulticloudAttachedCluster", "name", a.id.String())

	latest, err := a.gcpClient.GetAttachedCluster(ctx, &pb.GetAttachedClusterRequest{Name: a.id.String()})
	if err != nil {
		latest = created
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GKEMulticloudAttachedCluster", "name", a.id.String())

	diffs, updateMask, err := compareGKEMulticloudAttachedCluster(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if len(updateMask.Paths) > 0 {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateAttachedClusterRequest{
			UpdateMask:      updateMask,
			AttachedCluster: a.desired,
		}
		op, err := a.gcpClient.UpdateAttachedCluster(ctx, req)
		if err != nil {
			return fmt.Errorf("updating GKEMulticloudAttachedCluster %s: %w", a.id.String(), err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("GKEMulticloudAttachedCluster %s waiting update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated GKEMulticloudAttachedCluster", "name", a.id.String())

		latest, err = a.gcpClient.GetAttachedCluster(ctx, &pb.GetAttachedClusterRequest{Name: a.id.String()})
		if err != nil {
			latest = updated
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AttachedCluster) error {
	status := &krm.GKEMulticloudAttachedClusterStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = GKEMulticloudAttachedClusterObservedState_FromProto(mapCtx, latest)
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

	obj := &krm.GKEMulticloudAttachedCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GKEMulticloudAttachedClusterSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.AttachedCluster)
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.AttachedCluster)
	u.SetGroupVersionKind(krm.GKEMulticloudAttachedClusterGVK)

	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting GKEMulticloudAttachedCluster", "name", a.id.String())

	req := &pb.DeleteAttachedClusterRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteAttachedCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting GKEMulticloudAttachedCluster %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted GKEMulticloudAttachedCluster", "name", a.id.String())

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete GKEMulticloudAttachedCluster %s: %w", a.id.String(), err)
	}
	return true, nil
}

func compareGKEMulticloudAttachedCluster(ctx context.Context, actual, desired *pb.AttachedCluster) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, GKEMulticloudAttachedClusterSpec_FromProto, GKEMulticloudAttachedClusterSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Annotations = actual.Annotations

	// Normalize defaults for omitted fields in desired to prevent false diffs
	if desired.BinaryAuthorization == nil {
		maskedActual.BinaryAuthorization = nil
	}
	if desired.MonitoringConfig == nil {
		maskedActual.MonitoringConfig = nil
	}
	if desired.LoggingConfig == nil {
		maskedActual.LoggingConfig = nil
	}
	if desired.Authorization == nil {
		maskedActual.Authorization = nil
	}
	if desired.ProxyConfig == nil {
		maskedActual.ProxyConfig = nil
	}
	if desired.SecurityPostureConfig == nil {
		maskedActual.SecurityPostureConfig = nil
	}

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	// Filter updateMask paths to only include valid paths supported by the API and map to exact names expected by MockGCP/GCP API.
	var filteredPaths []string
	for _, p := range updateMask.Paths {
		switch p {
		case "annotations":
			filteredPaths = append(filteredPaths, "annotations")
		case "authorization.admin_users", "authorization.adminUsers":
			filteredPaths = append(filteredPaths, "authorization.admin_users")
		case "binary_authorization.evaluation_mode", "binaryAuthorization.evaluationMode":
			filteredPaths = append(filteredPaths, "binary_authorization.evaluation_mode")
		case "description":
			filteredPaths = append(filteredPaths, "description")
		case "logging_config.component_config.enable_components", "loggingConfig.componentConfig.enableComponents":
			filteredPaths = append(filteredPaths, "logging_config.component_config.enable_components")
		case "monitoring_config.managed_prometheus_config.enabled", "monitoringConfig.managedPrometheusConfig.enabled":
			filteredPaths = append(filteredPaths, "monitoring_config.managed_prometheus_config.enabled")
		case "platform_version", "platformVersion":
			filteredPaths = append(filteredPaths, "platform_version")
		}
	}
	updateMask.Paths = filteredPaths

	return diffs, updateMask, nil
}
