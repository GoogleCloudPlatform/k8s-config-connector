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

package developerconnect

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/developerconnect/apiv1"
	pb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/developerconnect/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DevConnectInsightsConfigGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.InsightsConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewInsightsConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building InsightsConfig client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DevConnectInsightsConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DevConnectInsightsConfigIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if obj.Spec.AppHubApplicationRef != nil {
		normalizedApp, err := obj.Spec.AppHubApplicationRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return nil, err
		}
		obj.Spec.AppHubApplicationRef.External = normalizedApp
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DevConnectInsightsConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = id.String()
	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	// Default to Projects context tracking current project if AppHubApplicationRef is not set
	if desired.InsightsConfigContext == nil {
		desired.InsightsConfigContext = &pb.InsightsConfig_Projects{
			Projects: &pb.Projects{
				ProjectIds: []string{id.Project},
			},
		}
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		obj:       obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.DevConnectInsightsConfigIdentity
	gcpClient *gcp.InsightsConfigClient
	desired   *pb.InsightsConfig
	obj       *krm.DevConnectInsightsConfig
	actual    *pb.InsightsConfig
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting InsightsConfig", "name", a.id)

	actual, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting InsightsConfig %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *Adapter) get(ctx context.Context) (*pb.InsightsConfig, error) {
	req := &pb.GetInsightsConfigRequest{Name: a.id.String()}
	return a.gcpClient.GetInsightsConfig(ctx, req)
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating InsightsConfig", "name", a.id)

	req := &pb.CreateInsightsConfigRequest{
		Parent:           fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		InsightsConfigId: a.id.Insights_config,
		InsightsConfig:   a.desired,
	}
	op, err := a.gcpClient.CreateInsightsConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating InsightsConfig %s: %w", a.id, err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for InsightsConfig creation %s: %w", a.id, err)
	}

	actual, err := a.get(ctx)
	if err != nil {
		return err
	}
	a.actual = actual

	return a.updateStatus(ctx, createOp, actual)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating InsightsConfig", "name", a.id)

	diffs, err := compareResource(ctx, a.actual, a.desired, a.id.Project)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for InsightsConfig, skipping update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	updateDesired := proto.Clone(a.desired).(*pb.InsightsConfig)
	if a.actual.InsightsConfigContext != nil {
		updateDesired.InsightsConfigContext = a.actual.InsightsConfigContext
	}
	if a.actual.GetProjects() != nil {
		updateDesired.InsightsConfigContext = nil
	}

	req := &pb.UpdateInsightsConfigRequest{
		InsightsConfig: updateDesired,
	}
	op, err := a.gcpClient.UpdateInsightsConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating InsightsConfig %s: %w", a.id, err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for InsightsConfig update %s: %w", a.id, err)
	}

	actual, err := a.get(ctx)
	if err != nil {
		return err
	}
	a.actual = actual

	return a.updateStatus(ctx, updateOp, actual)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting InsightsConfig", "name", a.id)

	req := &pb.DeleteInsightsConfigRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInsightsConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting InsightsConfig %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for InsightsConfig deletion %s: %w", a.id, err)
	}

	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DevConnectInsightsConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DevConnectInsightsConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Insights_config)
	u.SetGroupVersionKind(krm.DevConnectInsightsConfigGVK)

	u.Object = uObj
	return u, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.InsightsConfig) error {
	mapCtx := &direct.MapContext{}
	observedState := DevConnectInsightsConfigObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.DevConnectInsightsConfigStatus{}
	status.ObservedState = observedState
	status.ExternalRef = direct.LazyPtr(a.id.String())

	return op.UpdateStatus(ctx, status, nil)
}

func compareResource(ctx context.Context, actual, desired *pb.InsightsConfig, projectID string) (*structuredreporting.Diff, error) {
	mapCtx := &direct.MapContext{}
	maskedActualSpec := DevConnectInsightsConfigSpec_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	maskedActual := DevConnectInsightsConfigSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.InsightsConfig)

	// Normalize AppHubApplication references to avoid false diffs (including project numbers)
	if actual.GetAppHubApplication() != "" {
		appHub := actual.GetAppHubApplication()
		const prefix = "//apphub.googleapis.com/"
		if len(appHub) > len(prefix) && appHub[:len(prefix)] == prefix {
			appHub = appHub[len(prefix):]
		}
		// If appHub starts with "projects/[project_number]/", replace with "projects/[projectID]/"
		tokens := strings.Split(appHub, "/")
		if len(tokens) >= 2 && tokens[0] == "projects" {
			tokens[1] = projectID
			appHub = strings.Join(tokens, "/")
		}

		if maskedActual.InsightsConfigContext != nil {
			if context, ok := maskedActual.InsightsConfigContext.(*pb.InsightsConfig_AppHubApplication); ok {
				context.AppHubApplication = appHub
			}
		}
		if clonedDesired.InsightsConfigContext != nil {
			if context, ok := clonedDesired.InsightsConfigContext.(*pb.InsightsConfig_AppHubApplication); ok {
				context.AppHubApplication = appHub
			}
		}
	}

	// Since Projects is a KCC-internal default and not exposed to KRM users,
	// copy it from actual to clonedDesired and maskedActual to avoid false diffs.
	if actual.GetProjects() != nil {
		clonedDesired.InsightsConfigContext = actual.InsightsConfigContext
		maskedActual.InsightsConfigContext = actual.InsightsConfigContext
	}

	maskedActual.Labels = actual.Labels
	clonedDesired.Labels = desired.Labels

	maskedActual.Annotations = actual.Annotations
	clonedDesired.Annotations = actual.Annotations

	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
