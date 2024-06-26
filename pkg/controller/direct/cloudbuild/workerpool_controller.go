/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudbuild

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	gcp "cloud.google.com/go/cloudbuild/apiv1/v2"
	cloudbuildpb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"google.golang.org/api/option"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/references"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.GroupVersionKind, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

const ctrlName = "cloudbuild-controller"

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building cloudbuild client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.CloudBuildWorkerPool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Get ResourceID
	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Get GCP Project
	projectRef, err := references.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Get location
	location := obj.Spec.Location

	// Get computeNetwork
	if obj.Spec.PrivatePoolConfig.NetworkConfig != nil {
		networkRef, err := references.ResolveComputeNetwork(ctx, reader, obj, &obj.Spec.PrivatePoolConfig.NetworkConfig.PeeredNetworkRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.PrivatePoolConfig.NetworkConfig.PeeredNetworkRef.External = networkRef.String()
	}

	// Get CloudBuild GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		resourceID: resourceID,
		projectID:  projectID,
		location:   location,
		gcpClient:  gcpClient,
		desired:    obj,
	}, nil
}

type Adapter struct {
	resourceID string
	projectID  string
	location   string
	gcpClient  *gcp.Client
	desired    *krm.CloudBuildWorkerPool
	actual     *cloudbuildpb.WorkerPool
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	req := &cloudbuildpb.GetWorkerPoolRequest{Name: a.fullyQualifiedName()}
	workerpoolpb, err := a.gcpClient.GetWorkerPool(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting cloudbuildworkerpool %q: %w", a.fullyQualifiedName(), err)
	}

	a.actual = workerpoolpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating object", "u", u)

	projectID := a.projectID
	if projectID == "" {
		return fmt.Errorf("project is empty")
	}
	if a.resourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	desired := a.desired.DeepCopy()
	wp := &cloudbuildpb.WorkerPool{
		Name: a.fullyQualifiedName(),
	}
	err := krm.Convert_WorkerPool_KRM_To_API_v1(desired, wp)
	if err != nil {
		return fmt.Errorf("converting workerpool spec to api: %w", err)
	}
	req := &cloudbuildpb.CreateWorkerPoolRequest{
		Parent:       a.getParent(),
		WorkerPoolId: a.resourceID,
		WorkerPool:   wp,
	}
	op, err := a.gcpClient.CreateWorkerPool(ctx, req)
	if err != nil {
		return fmt.Errorf("cloudbuildworkerpool %s creating failed: %w", wp.Name, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("cloudbuildworkerpool %s waiting creation failed: %w", wp.Name, err)
	}
	status := &krm.CloudBuildWorkerPoolStatus{}
	if err := krm.Convert_WorkerPool_API_v1_To_KRM_status(created, status); err != nil {
		return fmt.Errorf("update workerpool status %w", err)
	}
	status.ObservedState.CreateTime = ToOpenAPIDateTime(created.GetCreateTime())
	status.ObservedState.UpdateTime = ToOpenAPIDateTime(created.GetUpdateTime())
	resRef, err := NewResourceRef(created)
	if err != nil {
		return err
	}
	status.ExternalRef = resRef.GetExternalReference()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	if err := a.ValidateExternalResource(); err != nil {
		return err
	}

	updateMask := &fieldmaskpb.FieldMask{}

	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}

	typedConfig, ok := a.actual.Config.(*cloudbuildpb.WorkerPool_PrivatePoolV1Config)
	if !ok {
		return fmt.Errorf("unable to convert cloudbuildworkerpool %s config to workerpool PrivatePoolV1Config", a.actual.Name)
	}
	actualConfig := typedConfig.PrivatePoolV1Config
	desiredConfig := a.desired.Spec.PrivatePoolConfig

	if desiredConfig.NetworkConfig != nil {
		switch actualConfig.NetworkConfig.EgressOption {
		case cloudbuildpb.PrivatePoolV1Config_NetworkConfig_EGRESS_OPTION_UNSPECIFIED:
			if !reflect.DeepEqual(ValueOf(desiredConfig.NetworkConfig.EgressOption), "UNSPECIFIED") {
				updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.network_config.egress_option")
			}
		case cloudbuildpb.PrivatePoolV1Config_NetworkConfig_NO_PUBLIC_EGRESS:
			if !reflect.DeepEqual(ValueOf(desiredConfig.NetworkConfig.EgressOption), "NO_PUBLIC_EGRESS") {
				updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.network_config.egress_option")
			}
		case cloudbuildpb.PrivatePoolV1Config_NetworkConfig_PUBLIC_EGRESS:
			if !reflect.DeepEqual(ValueOf(desiredConfig.NetworkConfig.EgressOption), "PUBLIC_EGRESS") {
				updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.network_config.egress_option")
			}
		}
		expectedIPRange := ValueOf(desiredConfig.NetworkConfig.PeeredNetworkIPRange)
		if expectedIPRange != "" && !reflect.DeepEqual(expectedIPRange, actualConfig.NetworkConfig.PeeredNetworkIpRange) {
			updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.network_config.peered_network_ip_range")
		}

		// TODO: better handle the network complexity
		// 1. peered_network is an immutable field. whether/when shall we validate
		// 2. the gcp workerpool stores the network with "project_number", different from the spec which uses the "project_id".
		//    * projects/<project_number>/global/networks/<network_id>
		//    * projects/<project_id>/global/networks/<network_id>
		desiredNetwork := strings.Split(desiredConfig.NetworkConfig.PeeredNetworkRef.External, "/")
		actualNetwork := strings.Split(actualConfig.NetworkConfig.PeeredNetwork, "/")
		if len(desiredNetwork) == 5 && len(actualNetwork) == 5 && !reflect.DeepEqual(desiredNetwork[4], actualNetwork[4]) {
			return fmt.Errorf("peered_network is immutable field")
		}
	}
	if !reflect.DeepEqual(desiredConfig.WorkerConfig.DiskSizeGb, actualConfig.WorkerConfig.DiskSizeGb) {
		updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.worker_config.disk_size_gb")
	}
	if !reflect.DeepEqual(desiredConfig.WorkerConfig.MachineType, actualConfig.WorkerConfig.MachineType) {
		updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.worker_config.machine_type")
	}
	if len(updateMask.Paths) == 0 {
		klog.Warningf("unexpected empty update mask, desired: %v, actual: %v", a.desired, a.actual)
		return nil
	}

	wp := &cloudbuildpb.WorkerPool{
		Name: a.fullyQualifiedName(),
		Etag: a.actual.Etag,
	}
	desired := a.desired.DeepCopy()
	err := krm.Convert_WorkerPool_KRM_To_API_v1(desired, wp)
	if err != nil {
		return fmt.Errorf("converting workerpool spec to api: %w", err)
	}
	req := &cloudbuildpb.UpdateWorkerPoolRequest{
		WorkerPool: wp,
		UpdateMask: updateMask,
	}
	op, err := a.gcpClient.UpdateWorkerPool(ctx, req)
	if err != nil {
		return fmt.Errorf("cloudbuildworkerpool %s updating failed: %w", wp.Name, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("cloudbuildworkerpool %s waiting update failed: %w", wp.Name, err)
	}
	status := &krm.CloudBuildWorkerPoolStatus{}
	if err := krm.Convert_WorkerPool_API_v1_To_KRM_status(updated, status); err != nil {
		return fmt.Errorf("update workerpool status %w", err)
	}
	status.ObservedState.CreateTime = ToOpenAPIDateTime(updated.GetCreateTime())
	status.ObservedState.UpdateTime = ToOpenAPIDateTime(updated.GetUpdateTime())
	// This value should not be updated. Just in case.
	resRef, err := NewResourceRef(updated)
	if err != nil {
		return err
	}
	status.ExternalRef = resRef.GetExternalReference()
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

// Delete implements the Adapter interface.
// TODO: Delete can rely on status.externalRef and do not need spec.projectRef.
func (a *Adapter) Delete(ctx context.Context) (bool, error) {
	if err := a.ValidateExternalResource(); err != nil {
		return false, err
	}
	req := &cloudbuildpb.DeleteWorkerPoolRequest{Name: a.fullyQualifiedName(), AllowMissing: true}
	op, err := a.gcpClient.DeleteWorkerPool(ctx, req)
	if err != nil {
		// likely a server bug. worker_pool can be successfully deleted.
		if !strings.Contains(err.Error(), "(line 12:3): missing \"value\" field") {
			return false, fmt.Errorf("deleting cloudbuildworkerpool %s: %w", a.fullyQualifiedName(), err)
		}
	}
	err = op.Wait(ctx)
	if err != nil {
		// likely a server bug. worker_pool can be successfully deleted.
		if !strings.Contains(err.Error(), "(line 12:3): missing \"value\" field") {
			return false, fmt.Errorf("waiting delete cloudbuildworkerpool %s: %w", a.fullyQualifiedName(), err)
		}
	}
	return true, nil
}

// ValidateExternalResource compares the `status.externalRef` with the `spec` Project, Location and
// (external) resourceID to make sure those fields are immutable and matches the previous deployed value.
func (a *Adapter) ValidateExternalResource() error {
	actualResRef, err := NewResourceRef(a.actual)
	if err != nil {
		return err
	}
	desiredExternalRef := "https://cloudbuild.googleapis.com/v1/" + a.fullyQualifiedName()
	if ValueOf(actualResRef.GetExternalReference()) == desiredExternalRef {
		return nil
	}

	// Give user guidance on how to fix the CloudBuildWorkerPool spec.
	if a.desired.Spec.ResourceID != nil && ValueOf(a.desired.Spec.ResourceID) != actualResRef.GetResourceID() {
		return fmt.Errorf("`spec.resourceID` is immutable field, expect %s, got %s",
			actualResRef.GetResourceID(), *a.desired.Spec.ResourceID)
	}
	if a.desired.Spec.Location != actualResRef.GetLocation() {
		return fmt.Errorf("`spec.location` is immutable field, expect %s, got %s",
			actualResRef.GetLocation(), a.desired.Spec.Location)
	}
	// TODO: Some Selflink may change the project from projectID to projectNum.
	/*
		if a.desired.Spec.ProjectRef.Name != "" {
			return fmt.Errorf("`spec.projectRef.name` is immutable field, expect project %s",
				actualExternalRef.GetProject())
		}
		if a.desired.Spec.ProjectRef.External != "" {
			return fmt.Errorf("`spec.projectRef.external` is immutable field, expect project %s",
				actualExternalRef.GetProject())
		}*/
	return nil
}

func (a *Adapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", a.projectID, a.location, a.resourceID)
}

func (a *Adapter) getParent() string {
	return fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location)
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
	}

	u.Object["status"] = status

	return nil
}

func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

// IsNotFound returns true if the given error is an HTTP 404.
func IsNotFound(err error) bool {
	return HasHTTPCode(err, 404)
}

// HasHTTPCode returns true if the given error is an HTTP response with the given code.
func HasHTTPCode(err error, code int) bool {
	if err == nil {
		return false
	}
	apiError := &apierror.APIError{}
	if errors.As(err, &apiError) {
		if apiError.HTTPCode() == code {
			return true
		}
	} else {
		klog.Warningf("unexpected error type %T", err)
	}
	return false
}

// LazyPtr returns a pointer to v, unless it is the empty value, in which case it returns nil.
// It is essentially the inverse of ValueOf, though it is lossy
// because we can't tell nil and empty apart without a pointer.
func LazyPtr[T comparable](v T) *T {
	var defaultValue T
	if v == defaultValue {
		return nil
	}
	return &v
}

func ToOpenAPIDateTime(ts *timestamppb.Timestamp) *string {
	formatted := ts.AsTime().Format(time.RFC3339)
	return &formatted
}
