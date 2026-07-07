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
	"sort"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1beta1"
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
	registry.RegisterModel(krm.CloudDeployAutomationGVK, NewAutomationModel)
}

func NewAutomationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAutomation{config: *config}, nil
}

var _ directbase.Model = &modelAutomation{}

type modelAutomation struct {
	config config.ControllerConfig
}

func (m *modelAutomation) client(ctx context.Context) (*gcp.CloudDeployClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudDeployRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Automation client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelAutomation) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudDeployAutomation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Validate that users are not trying to manage system labels (goog- or go-)
	labels := u.GetLabels()
	for k := range labels {
		if isSystemLabel(k) {
			return nil, fmt.Errorf("system label %q is not allowed in metadata.labels", k)
		}
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

	return &AutomationAdapter{
		id:        id.(*krm.CloudDeployAutomationIdentity),
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
		labels:    u.GetLabels(),
	}, nil
}

func (m *modelAutomation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func clearRuleCondition(rule *clouddeploypb.AutomationRule) {
	if rule == nil {
		return
	}
	ruleMsg := rule.ProtoReflect()
	ruleMsg.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Kind() == protoreflect.MessageKind {
			innerMsg := v.Message()
			conditionFD := innerMsg.Descriptor().Fields().ByName("condition")
			if conditionFD != nil && innerMsg.Has(conditionFD) {
				innerMsg.Clear(conditionFD)
			}
		}
		return true
	})
}

type AutomationAdapter struct {
	id        *krm.CloudDeployAutomationIdentity
	gcpClient *gcp.CloudDeployClient
	desiredPb *clouddeploypb.Automation
	actual    *clouddeploypb.Automation

	desired *krm.CloudDeployAutomation
	reader  client.Reader
	labels  map[string]string
}

func (a *AutomationAdapter) resolveReferences(ctx context.Context) error {
	obj := a.desired
	reader := a.reader

	if obj.Spec.ServiceAccountRef != nil {
		if err := obj.Spec.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
			return err
		}
	}
	if obj.Spec.Selector != nil {
		for i := range obj.Spec.Selector.Targets {
			if obj.Spec.Selector.Targets[i].TargetRef != nil {
				if err := obj.Spec.Selector.Targets[i].TargetRef.Normalize(ctx, reader, obj.Namespace); err != nil {
					return err
				}
			}
		}
	}
	for i := range obj.Spec.Rules {
		rule := &obj.Spec.Rules[i]
		if rule.PromoteReleaseRule != nil && rule.PromoteReleaseRule.DestinationTargetRef != nil {
			if err := rule.PromoteReleaseRule.DestinationTargetRef.Normalize(ctx, reader, obj.Namespace); err != nil {
				return err
			}
		}
		if rule.TimedPromoteReleaseRule != nil && rule.TimedPromoteReleaseRule.DestinationTargetRef != nil {
			if err := rule.TimedPromoteReleaseRule.DestinationTargetRef.Normalize(ctx, reader, obj.Namespace); err != nil {
				return err
			}
		}
	}

	mapCtx := &direct.MapContext{}
	a.desiredPb = CloudDeployAutomationSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return nil
}

var _ directbase.Adapter = &AutomationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AutomationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Automation", "name", a.id.String())

	req := &clouddeploypb.GetAutomationRequest{Name: a.id.String()}
	automationpb, err := a.gcpClient.GetAutomation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Automation %q: %w", a.id.String(), err)
	}

	a.actual = automationpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AutomationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Automation", "name", a.id.String())

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	a.desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(a.labels)

	mapCtx := &direct.MapContext{}

	req := &clouddeploypb.CreateAutomationRequest{
		Parent:       a.id.Parent().String(),
		Automation:   a.desiredPb,
		AutomationId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateAutomation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Automation %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Automation %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created Automation", "name", a.id.String())

	status := &krm.CloudDeployAutomationStatus{}
	status.ObservedState = CloudDeployAutomationObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		log.Error(mapCtx.Err(), "error mapping Automation status")
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AutomationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Automation", "name", a.id.String())

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}

	a.desiredPb.Name = a.id.String()

	a.desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(a.labels)

	// Preserve system labels (goog- or go-)
	if a.actual.Labels != nil {
		if a.desiredPb.Labels == nil {
			a.desiredPb.Labels = make(map[string]string)
		}
		for k, v := range a.actual.Labels {
			if isSystemLabel(k) {
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
		// KCC's CompareProtoMessage delegates slice comparisons entirely to proto.Equal,
		// which does not filter OUTPUT_ONLY fields inside list elements. We clone actual
		// and strip output-only Condition structs from rules to prevent infinite update loops.
		if fieldName == "rules" {
			actualClone := proto.Clone(b.(*clouddeploypb.Automation)).(*clouddeploypb.Automation)
			for _, rule := range actualClone.Rules {
				clearRuleCondition(rule)
			}
			desired := a.(*clouddeploypb.Automation)
			if len(desired.Rules) != len(actualClone.Rules) {
				return true, nil
			}

			// The API might return rules in a different order.
			// Sort by ID before comparing to avoid spurious diffs.
			sortRules(desired.Rules)
			sortRules(actualClone.Rules)

			for i := range desired.Rules {
				if !proto.Equal(desired.Rules[i], actualClone.Rules[i]) {
					return true, nil
				}
			}
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

		// Inject the latest etag for optimistic concurrency
		a.desiredPb.Etag = a.actual.Etag

		req := &clouddeploypb.UpdateAutomationRequest{
			UpdateMask: updateMask,
			Automation: a.desiredPb,
		}
		op, err := a.gcpClient.UpdateAutomation(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Automation %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Automation %s waiting update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated Automation", "name", a.id.String())
	} else {
		log.V(2).Info("no field needs update", "name", a.id.String())
	}

	status := &krm.CloudDeployAutomationStatus{}
	status.ObservedState = CloudDeployAutomationObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		log.Error(mapCtx.Err(), "error mapping Automation status")
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AutomationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudDeployAutomation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudDeployAutomationSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = direct.LazyPtr(a.id.Parent().Location)
	obj.Spec.DeliveryPipelineRef = &krmv1beta1.DeliveryPipelineRef{External: a.id.Parent().String()}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(k8s.ValueToDNSSubdomainName(a.id.ID()))
	u.SetGroupVersionKind(krm.CloudDeployAutomationGVK)

	u.Object = uObj

	// Filter out system labels (goog- or go-)
	labels := make(map[string]string)
	for k, v := range a.actual.Labels {
		if !isSystemLabel(k) {
			labels[k] = v
		}
	}
	u.SetLabels(labels)
	return u, nil
}

func isSystemLabel(k string) bool {
	if strings.HasPrefix(k, "goog-") || strings.HasPrefix(k, "go-") {
		return true
	}
	if k == label.CnrmManagedKey {
		return true
	}
	return false
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AutomationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Automation", "name", a.id.String())

	req := &clouddeploypb.DeleteAutomationRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteAutomation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Automation, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Automation %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Automation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Automation", "name", a.id.String())
	return true, nil
}

func getRuleID(rule *clouddeploypb.AutomationRule) string {
	if rule == nil {
		return ""
	}
	if r := rule.GetPromoteReleaseRule(); r != nil {
		return r.GetId()
	}
	if r := rule.GetAdvanceRolloutRule(); r != nil {
		return r.GetId()
	}
	if r := rule.GetRepairRolloutRule(); r != nil {
		return r.GetId()
	}
	if r := rule.GetTimedPromoteReleaseRule(); r != nil {
		return r.GetId()
	}
	return ""
}

func sortRules(rules []*clouddeploypb.AutomationRule) {
	sort.Slice(rules, func(i, j int) bool {
		return getRuleID(rules[i]) < getRuleID(rules[j])
	})
}
