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

package clouddeploydeploypolicy

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/clouddeploy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
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
	registry.RegisterModel(krm.DeployDeployPolicyGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.CloudDeployClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudDeployRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DeployPolicy client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudDeployDeployPolicy{}
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

	return &Adapter{
		id:        id.(*krm.DeployPolicyIdentity),
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
		labels:    u.GetLabels(),
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.DeployPolicyIdentity
	gcpClient *gcp.CloudDeployClient
	desiredPb *clouddeploypb.DeployPolicy
	actual    *clouddeploypb.DeployPolicy

	desired *krm.CloudDeployDeployPolicy
	reader  client.Reader
	labels  map[string]string
}

func (a *Adapter) resolveReferences(ctx context.Context) error {
	obj := a.desired

	mapCtx := &direct.MapContext{}
	a.desiredPb = clouddeploy.DeployPolicySpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return nil
}

var _ directbase.Adapter = &Adapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	gcpName := a.id.String()
	log.V(2).Info("getting DeployPolicy", "name", gcpName)

	req := &clouddeploypb.GetDeployPolicyRequest{Name: gcpName}
	deploypolicypb, err := a.gcpClient.GetDeployPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DeployPolicy %q: %w", gcpName, err)
	}

	a.actual = deploypolicypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	gcpName := a.id.String()
	log.V(2).Info("creating DeployPolicy", "name", gcpName)

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}

	req := &clouddeploypb.CreateDeployPolicyRequest{
		Parent:         fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		DeployPolicy:   a.desiredPb,
		DeployPolicyId: a.id.DeployPolicy,
	}
	op, err := a.gcpClient.CreateDeployPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DeployPolicy %s: %w", gcpName, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DeployPolicy %s waiting creation: %w", gcpName, err)
	}
	log.V(2).Info("successfully created DeployPolicy", "name", gcpName)

	status := &krm.DeployPolicyStatus{}
	status.ObservedState = clouddeploy.DeployPolicyObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		log.Error(mapCtx.Err(), "error mapping DeployPolicy status")
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	gcpName := a.id.String()
	log.V(2).Info("updating DeployPolicy", "name", gcpName)

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}

	a.desiredPb.Name = gcpName

	// Let's perform diff between desired and actual.
	paths, err := common.CompareProtoMessage(a.desiredPb, a.actual, func(fieldName protoreflect.Name, a, b proto.Message) (bool, error) {
		if fieldName == "etag" {
			return false, nil
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

		req := &clouddeploypb.UpdateDeployPolicyRequest{
			UpdateMask:   updateMask,
			DeployPolicy: a.desiredPb,
		}
		op, err := a.gcpClient.UpdateDeployPolicy(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DeployPolicy %s: %w", gcpName, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DeployPolicy %s waiting update: %w", gcpName, err)
		}
		log.V(2).Info("successfully updated DeployPolicy", "name", gcpName)
	} else {
		log.V(2).Info("no field needs update", "name", gcpName)
	}

	status := &krm.DeployPolicyStatus{}
	status.ObservedState = clouddeploy.DeployPolicyObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		log.Error(mapCtx.Err(), "error mapping DeployPolicy status")
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudDeployDeployPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(clouddeploy.DeployPolicySpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(k8s.ValueToDNSSubdomainName(a.id.DeployPolicy))
	u.SetGroupVersionKind(krm.DeployDeployPolicyGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	gcpName := a.id.String()
	log.V(2).Info("deleting DeployPolicy", "name", gcpName)

	req := &clouddeploypb.DeleteDeployPolicyRequest{Name: gcpName}
	op, err := a.gcpClient.DeleteDeployPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DeployPolicy, assuming it was already deleted", "name", gcpName)
			return true, nil
		}
		return false, fmt.Errorf("deleting DeployPolicy %s: %w", gcpName, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete wait for non-existent DeployPolicy, assuming it was already deleted", "name", gcpName)
			return true, nil
		}
		return false, fmt.Errorf("waiting delete DeployPolicy %s: %w", gcpName, err)
	}
	log.V(2).Info("successfully deleted DeployPolicy", "name", gcpName)
	return true, nil
}
