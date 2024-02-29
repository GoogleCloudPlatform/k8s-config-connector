// Copyright 2024 Google LLC
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

package compute

import (
	"context"
	"fmt"

	api "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
)

// AddHealthCheckReconciler creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddHealthCheckReconciler(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.ComputeHealthCheckGVK

	return directbase.Add(mgr, gvk, &model{config: *config})
}

type model struct {
	config controller.Config
}

var healthcheckMapping = NewMapping(&pb.HealthCheck{}, &krm.ComputeHealthCheck{},
	Spec("checkIntervalSec"),
	Spec("description"),
	Spec("healthyThreshold"),
	Ignore("location"),
	Spec("timeoutSec"),
	Spec("unhealthyThreshold"),

	Spec("grpcHealthCheck"),
	Spec("http2HealthCheck"),
	Spec("httpHealthCheck"),
	Spec("httpsHealthCheck"),
	Spec("logConfig"),
	Spec("sslHealthCheck"),
	Spec("tcpHealthCheck"),

	Status("creationTimestamp"),
	Status("selfLink"),
	Status("type"),
	TODO("resourceID"),
).
	MapNested(&pb.TCPHealthCheck{}, &krm.HealthcheckTcpHealthCheck{}, "port", "portName", "portSpecification", "proxyHeader", "request", "response").
	MustBuild()

type globalHealthCheckAdapter struct {
	projectID     string
	location      string
	healthcheckID string

	desired *krm.ComputeHealthCheck
	actual  *krm.ComputeHealthCheck

	gcp *api.HealthChecksClient
}

type regionalHealthCheckAdapter struct {
	projectID     string
	region        string
	healthcheckID string

	desired *krm.ComputeHealthCheck
	actual  *krm.ComputeHealthCheck

	gcp *api.RegionHealthChecksClient
}

func (m *model) client(ctx context.Context) (*api.HealthChecksClient, error) {
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

	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	gcpClient, err := api.NewHealthChecksRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building apihealthchecks client: %w", err)
	}

	return gcpClient, err
}

func (m *model) regionalClient(ctx context.Context) (*api.RegionHealthChecksClient, error) {
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

	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	gcpClient, err := api.NewRegionHealthChecksRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building apihealthchecks client: %w", err)
	}

	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	// TODO: Just fetch this object?
	obj := &krm.ComputeHealthCheck{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.GetAnnotations()[k8s.ProjectIDAnnotation]
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	// TODO: Use name or request resourceID to be set on create?
	healthcheckID := ValueOf(obj.Spec.ResourceID)
	if healthcheckID == "" {
		healthcheckID = obj.GetName()
	}
	if healthcheckID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("unable to determine location")
	}

	if location == "global" {
		gcp, err := m.client(ctx)
		if err != nil {
			return nil, err
		}

		return &globalHealthCheckAdapter{
			projectID:     projectID,
			location:      location,
			healthcheckID: healthcheckID,
			desired:       obj,
			gcp:           gcp,
		}, nil
	} else {
		gcp, err := m.regionalClient(ctx)
		if err != nil {
			return nil, err
		}

		return &regionalHealthCheckAdapter{
			projectID:     projectID,
			region:        location,
			healthcheckID: healthcheckID,
			desired:       obj,
			gcp:           gcp,
		}, nil
	}
}

func (a *globalHealthCheckAdapter) Find(ctx context.Context) (bool, error) {
	if a.healthcheckID == "" {
		return false, nil
	}

	req := &pb.GetHealthCheckRequest{
		Project:     a.projectID,
		HealthCheck: a.healthcheckID,
	}
	healthcheck, err := a.gcp.Get(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("healthcheck was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	u := &krm.ComputeHealthCheck{}
	if err := healthcheckMapping.Map(healthcheck, u, nil); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func (a *globalHealthCheckAdapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteHealthCheckRequest{
		Project:     a.projectID,
		HealthCheck: a.healthcheckID,
	}
	op, err := a.gcp.Delete(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting healthcheck: %w", err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for healthcheck deletion to complete: %w", err)
	}

	return true, nil
}

func (a *globalHealthCheckAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	// You can configure only the `display_name`, `restrictions`, and
	// `annotations` fields.
	desired := &pb.HealthCheck{}
	if err := healthcheckMapping.MapSpec(a.desired, desired); err != nil {
		return err
	}

	desired.Name = PtrTo(a.healthcheckID)
	req := &pb.InsertHealthCheckRequest{
		Project:             a.projectID,
		HealthCheckResource: desired,
	}

	op, err := a.gcp.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating healthcheck: %w", err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for healthcheck creation: %w", err)
	}
	// TODO: Return created object
	return nil
}

func (a *globalHealthCheckAdapter) Update(ctx context.Context) (*unstructured.Unstructured, error) {

	update := &pb.HealthCheck{}
	if err := healthcheckMapping.MapSpec(a.desired, update); err != nil {
		return nil, err
	}

	req := &pb.UpdateHealthCheckRequest{
		Project:             a.projectID,
		HealthCheck:         a.healthcheckID,
		HealthCheckResource: update,
	}

	op, err := a.gcp.Update(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("updating healthcheck: %w", err)
	}
	if err := op.Wait(ctx); err != nil {
		return nil, fmt.Errorf("waiting for healthcheck wait: %w", err)
	}
	// TODO: Return updated object
	return nil, nil
}

func (a *regionalHealthCheckAdapter) Find(ctx context.Context) (bool, error) {
	if a.healthcheckID == "" {
		return false, nil
	}

	req := &pb.GetRegionHealthCheckRequest{
		Project:     a.projectID,
		Region:      a.region,
		HealthCheck: a.healthcheckID,
	}
	healthcheck, err := a.gcp.Get(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("healthcheck was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	u := &krm.ComputeHealthCheck{}
	if err := healthcheckMapping.MapSpec(healthcheck, u); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func (a *regionalHealthCheckAdapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteRegionHealthCheckRequest{
		Project:     a.projectID,
		Region:      a.region,
		HealthCheck: a.healthcheckID,
	}
	op, err := a.gcp.Delete(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting healthcheck: %w", err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for healthcheck deletion to complete: %w", err)
	}

	return true, nil
}

func (a *regionalHealthCheckAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	// You can configure only the `display_name`, `restrictions`, and
	// `annotations` fields.
	desired := &pb.HealthCheck{}
	if err := healthcheckMapping.MapSpec(a.desired, desired); err != nil {
		return err
	}

	desired.Name = PtrTo(a.healthcheckID)
	req := &pb.InsertRegionHealthCheckRequest{
		Project:             a.projectID,
		Region:              a.region,
		HealthCheckResource: desired,
	}

	op, err := a.gcp.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating healthcheck: %w", err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for healthcheck creation: %w", err)
	}
	// TODO: Return created object
	return nil
}

func (a *regionalHealthCheckAdapter) Update(ctx context.Context) (*unstructured.Unstructured, error) {

	update := &pb.HealthCheck{}
	if err := healthcheckMapping.MapSpec(a.desired, update); err != nil {
		return nil, err
	}

	req := &pb.UpdateRegionHealthCheckRequest{
		Project:             a.projectID,
		Region:              a.region,
		HealthCheck:         a.healthcheckID,
		HealthCheckResource: update,
	}

	op, err := a.gcp.Update(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("updating healthcheck: %w", err)
	}
	if err := op.Wait(ctx); err != nil {
		return nil, fmt.Errorf("waiting for healthcheck wait: %w", err)
	}
	// TODO: Return updated object
	return nil, nil
}
