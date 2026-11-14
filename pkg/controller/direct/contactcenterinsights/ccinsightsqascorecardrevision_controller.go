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

// +generated:controller
// krm.group: contactcenterinsights.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CCInsightsQAScorecardRevision

package contactcenterinsights

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/contactcenterinsights/apiv1"
	contactcenterinsightspb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CCInsightsQAScorecardRevisionGVK, NewQaScorecardRevisionModel)
}

func NewQaScorecardRevisionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelQaScorecardRevision{config: *config}, nil
}

var _ directbase.Model = &modelQaScorecardRevision{}

type modelQaScorecardRevision struct {
	config config.ControllerConfig
}

func (m *modelQaScorecardRevision) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CCInsightsQAScorecardRevision{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CCInsightsQAScorecardRevisionIdentity)

	// Get contactcenterinsights GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &QaScorecardRevisionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelQaScorecardRevision) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//contactcenterinsights.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CCInsightsQAScorecardRevisionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &QaScorecardRevisionAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelQaScorecardRevision) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building QaScorecardRevision client: %w", err)
	}
	return gcpClient, err
}

type QaScorecardRevisionAdapter struct {
	id        *krm.CCInsightsQAScorecardRevisionIdentity
	gcpClient *gcp.Client
	desired   *krm.CCInsightsQAScorecardRevision
	actual    *contactcenterinsightspb.QaScorecardRevision
}

var _ directbase.Adapter = &QaScorecardRevisionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *QaScorecardRevisionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting QaScorecardRevision", "name", a.id)

	req := &contactcenterinsightspb.GetQaScorecardRevisionRequest{
		Name: a.id.String(),
	}
	qascorecardrevisionpb, err := a.gcpClient.GetQaScorecardRevision(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting QaScorecardRevision %q: %w", a.id, err)
	}

	a.actual = qascorecardrevisionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *QaScorecardRevisionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CCInsightsQAScorecardRevisionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	log.V(2).Info("creating QaScorecardRevision", "name", a.id)

	if err := a.prepareParentScorecard(ctx); err != nil {
		return fmt.Errorf("preparing parent scorecard for revision %s: %w", a.id, err)
	}

	req := &contactcenterinsightspb.CreateQaScorecardRevisionRequest{
		Parent:                a.id.ParentString(),
		QaScorecardRevisionId: a.id.Revision,
		QaScorecardRevision:   resource,
	}
	created, err := a.gcpClient.CreateQaScorecardRevision(ctx, req)
	if err != nil {
		return fmt.Errorf("creating QaScorecardRevision %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created QaScorecardRevision", "name", a.id)

	status := &krm.CCInsightsQAScorecardRevisionStatus{}
	status.ObservedState = CCInsightsQAScorecardRevisionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *QaScorecardRevisionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating QaScorecardRevision", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CCInsightsQAScorecardRevisionSpec_FromProto,
		CCInsightsQAScorecardRevisionSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating CCInsightsQAScorecardRevision is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.CCInsightsQAScorecardRevisionStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CCInsightsQAScorecardRevisionObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *QaScorecardRevisionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CCInsightsQAScorecardRevision{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CCInsightsQAScorecardRevisionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CCInsightsQAScorecardRevisionGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *QaScorecardRevisionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting QaScorecardRevision", "name", a.id)

	req := &contactcenterinsightspb.DeleteQaScorecardRevisionRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteQaScorecardRevision(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent QaScorecardRevision, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting QaScorecardRevision %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted QaScorecardRevision", "name", a.id)
	return true, nil
}

// prepareParentScorecard ensures the parent scorecard has no in-flight EDITABLE revisions
// before creating a new revision. A scorecard in GCP can only have one EDITABLE revision at a time;
// any existing EDITABLE revision must be tuned and deployed before a new revision can be created.
func (a *QaScorecardRevisionAdapter) prepareParentScorecard(ctx context.Context) error {
	log := klog.FromContext(ctx)
	it := a.gcpClient.ListQaScorecardRevisions(ctx, &contactcenterinsightspb.ListQaScorecardRevisionsRequest{
		Parent: a.id.ParentString(),
	})
	for {
		resp, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return fmt.Errorf("listing scorecard revisions under %s: %w", a.id.ParentString(), err)
		}
		if resp.GetState() == contactcenterinsightspb.QaScorecardRevision_EDITABLE && resp.GetName() != a.id.String() {
			log.V(2).Info("tuning initial scorecard revision to make it deployable", "name", resp.GetName())
			_, err := a.gcpClient.TuneQaScorecardRevision(ctx, &contactcenterinsightspb.TuneQaScorecardRevisionRequest{
				Parent:       resp.GetName(),
				ValidateOnly: false,
			})
			if err != nil {
				return fmt.Errorf("tuning scorecard revision %s: %w", resp.GetName(), err)
			}
			for {
				rev, err := a.gcpClient.GetQaScorecardRevision(ctx, &contactcenterinsightspb.GetQaScorecardRevisionRequest{
					Name: resp.GetName(),
				})
				if err != nil {
					return fmt.Errorf("getting scorecard revision %s during tuning: %w", resp.GetName(), err)
				}
				if rev.GetState() == contactcenterinsightspb.QaScorecardRevision_READY {
					break
				}
				if rev.GetState() == contactcenterinsightspb.QaScorecardRevision_TRAINING_FAILED || rev.GetState() == contactcenterinsightspb.QaScorecardRevision_TRAINING_CANCELLED {
					return fmt.Errorf("tuning scorecard revision %s failed with state: %v", resp.GetName(), rev.GetState())
				}
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(2 * time.Second):
				}
			}
			log.V(2).Info("deploying scorecard revision", "name", resp.GetName())
			if _, err := a.gcpClient.DeployQaScorecardRevision(ctx, &contactcenterinsightspb.DeployQaScorecardRevisionRequest{
				Name: resp.GetName(),
			}); err != nil {
				return fmt.Errorf("deploying scorecard revision %s: %w", resp.GetName(), err)
			}
		}
	}
	return nil
}
