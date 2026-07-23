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

package partnerssegateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityPartnerSSEGatewayGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityPartnerSSEGateway{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}

	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating authenticated HTTP client: %w", err)
	}

	desired := KRMtoGCP(&obj.Spec)

	return &adapter{
		httpClient: httpClient,
		id:         id.(*krm.NetworkSecurityPartnerSSEGatewayIdentity),
		desired:    desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type adapter struct {
	httpClient *http.Client
	id         *krm.NetworkSecurityPartnerSSEGatewayIdentity
	desired    *gcpPartnerSSEGateway
	actual     *gcpPartnerSSEGateway
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

type gcpPartnerSSEGateway struct {
	Name                  string              `json:"name,omitempty"`
	Labels                map[string]string   `json:"labels,omitempty"`
	SseGatewayReferenceId string              `json:"sseGatewayReferenceId,omitempty"`
	PartnerVpcSubnetRange string              `json:"partnerVpcSubnetRange,omitempty"`
	SseSubnetRange        string              `json:"sseSubnetRange,omitempty"`
	PartnerSubnetRange    string              `json:"partnerSubnetRange,omitempty"`
	Vni                   int32               `json:"vni,omitempty"`
	SymantecOptions       *gcpSymantecOptions `json:"symantecOptions,omitempty"`

	// Output fields
	CreateTime            string                           `json:"createTime,omitempty"`
	UpdateTime            string                           `json:"updateTime,omitempty"`
	SseVpcSubnetRange     string                           `json:"sseVpcSubnetRange,omitempty"`
	SseVpcTargetIp        string                           `json:"sseVpcTargetIp,omitempty"`
	SseBgpIps             []string                         `json:"sseBgpIps,omitempty"`
	SseBgpAsn             int32                            `json:"sseBgpAsn,omitempty"`
	PartnerSseRealm       string                           `json:"partnerSseRealm,omitempty"`
	SseTargetIp           string                           `json:"sseTargetIp,omitempty"`
	SymantecOptionsStatus *gcpSymantecOptionsObservedState `json:"symantecOptionsStatus,omitempty"`
	SseProject            string                           `json:"sseProject,omitempty"`
	SseNetwork            string                           `json:"sseNetwork,omitempty"`
	PartnerSseEnvironment string                           `json:"partnerSseEnvironment,omitempty"`
	Country               string                           `json:"country,omitempty"`
	Timezone              string                           `json:"timezone,omitempty"`
	CapacityBps           string                           `json:"capacityBps,omitempty"`
	State                 string                           `json:"state,omitempty"`
	ProberSubnetRanges    []string                         `json:"proberSubnetRanges,omitempty"`
}

type gcpSymantecOptions struct {
	SymantecSiteTargetHost string `json:"symantecSiteTargetHost,omitempty"`
}

type gcpSymantecOptionsObservedState struct {
	SymantecLocationUuid string `json:"symantecLocationUuid,omitempty"`
	SymantecSite         string `json:"symantecSite,omitempty"`
}

type gcpOperation struct {
	Name     string             `json:"name"`
	Done     bool               `json:"done"`
	Error    *gcpOperationError `json:"error,omitempty"`
	Response interface{}        `json:"response,omitempty"`
}

type gcpOperationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func KRMtoGCP(spec *krm.NetworkSecurityPartnerSSEGatewaySpec) *gcpPartnerSSEGateway {
	out := &gcpPartnerSSEGateway{}
	if spec.Labels != nil {
		out.Labels = spec.Labels
	}
	if spec.SseGatewayReferenceID != nil {
		out.SseGatewayReferenceId = *spec.SseGatewayReferenceID
	}
	if spec.PartnerVPCSubnetRange != nil {
		out.PartnerVpcSubnetRange = *spec.PartnerVPCSubnetRange
	}
	if spec.SseSubnetRange != nil {
		out.SseSubnetRange = *spec.SseSubnetRange
	}
	if spec.PartnerSubnetRange != nil {
		out.PartnerSubnetRange = *spec.PartnerSubnetRange
	}
	if spec.Vni != nil {
		out.Vni = *spec.Vni
	}
	if spec.SymantecOptions != nil {
		out.SymantecOptions = &gcpSymantecOptions{}
		if spec.SymantecOptions.SymantecSiteTargetHost != nil {
			out.SymantecOptions.SymantecSiteTargetHost = *spec.SymantecOptions.SymantecSiteTargetHost
		}
	}
	return out
}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting networksecurity partner sse gateway", "name", a.id)

	url := fmt.Sprintf("https://networksecurity.googleapis.com/v1alpha1/%s", a.id.String())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, fmt.Errorf("error creating GET request: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("error getting partner sse gateway %q: %w", a.id.String(), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("getting partner sse gateway %q failed with status %d: %s", a.id.String(), resp.StatusCode, string(body))
	}

	var gateway gcpPartnerSSEGateway
	if err := json.NewDecoder(resp.Body).Decode(&gateway); err != nil {
		return false, fmt.Errorf("error decoding partner sse gateway JSON: %w", err)
	}

	a.actual = &gateway
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating networksecurity partner sse gateway", "name", a.id)

	parent := a.id.ParentString()
	url := fmt.Sprintf("https://networksecurity.googleapis.com/v1alpha1/%s/partnerSSEGateways?partner_sse_gateway_id=%s", parent, a.id.PartnerSSEGateway)

	payload, err := json.Marshal(a.desired)
	if err != nil {
		return fmt.Errorf("error marshaling desired gateway payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("error creating POST request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error posting partner sse gateway: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("creating partner sse gateway failed with status %d: %s", resp.StatusCode, string(body))
	}

	var op gcpOperation
	if err := json.NewDecoder(resp.Body).Decode(&op); err != nil {
		return fmt.Errorf("error decoding operation response: %w", err)
	}

	if err := a.waitForOperation(ctx, op.Name); err != nil {
		return fmt.Errorf("waiting for partner sse gateway creation operation: %w", err)
	}

	// Fetch actual gateway after creation to get fully populated status fields
	found, err := a.Find(ctx)
	if err != nil {
		return fmt.Errorf("fetching partner sse gateway after creation: %w", err)
	}
	if !found {
		return fmt.Errorf("partner sse gateway not found after successful creation")
	}

	status := &krm.NetworkSecurityPartnerSSEGatewayStatus{}
	status.ObservedState = a.buildObservedState()
	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating networksecurity partner sse gateway", "name", a.id)

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	var paths []string

	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) {
		report.AddField("spec.labels", a.actual.Labels, a.desired.Labels)
		paths = append(paths, "labels")
	}
	if a.desired.SseGatewayReferenceId != a.actual.SseGatewayReferenceId {
		report.AddField("spec.sseGatewayReferenceID", a.actual.SseGatewayReferenceId, a.desired.SseGatewayReferenceId)
		paths = append(paths, "sseGatewayReferenceId")
	}
	if a.desired.PartnerVpcSubnetRange != a.actual.PartnerVpcSubnetRange {
		report.AddField("spec.partnerVPCSubnetRange", a.actual.PartnerVpcSubnetRange, a.desired.PartnerVpcSubnetRange)
		paths = append(paths, "partnerVpcSubnetRange")
	}
	if a.desired.SseSubnetRange != a.actual.SseSubnetRange {
		report.AddField("spec.sseSubnetRange", a.actual.SseSubnetRange, a.desired.SseSubnetRange)
		paths = append(paths, "sseSubnetRange")
	}
	if a.desired.PartnerSubnetRange != a.actual.PartnerSubnetRange {
		report.AddField("spec.partnerSubnetRange", a.actual.PartnerSubnetRange, a.desired.PartnerSubnetRange)
		paths = append(paths, "partnerSubnetRange")
	}
	if a.desired.Vni != a.actual.Vni {
		report.AddField("spec.vni", a.actual.Vni, a.desired.Vni)
		paths = append(paths, "vni")
	}
	if !reflect.DeepEqual(a.desired.SymantecOptions, a.actual.SymantecOptions) {
		report.AddField("spec.symantecOptions", a.actual.SymantecOptions, a.desired.SymantecOptions)
		paths = append(paths, "symantecOptions")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, report)

	url := fmt.Sprintf("https://networksecurity.googleapis.com/v1alpha1/%s?updateMask=%s", a.id.String(), strings.Join(paths, ","))

	payload, err := json.Marshal(a.desired)
	if err != nil {
		return fmt.Errorf("error marshaling desired gateway payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("error creating PATCH request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error patching partner sse gateway: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("updating partner sse gateway failed with status %d: %s", resp.StatusCode, string(body))
	}

	var op gcpOperation
	if err := json.NewDecoder(resp.Body).Decode(&op); err != nil {
		return fmt.Errorf("error decoding operation response: %w", err)
	}

	if err := a.waitForOperation(ctx, op.Name); err != nil {
		return fmt.Errorf("waiting for partner sse gateway update operation: %w", err)
	}

	// Fetch actual gateway after update
	found, err := a.Find(ctx)
	if err != nil {
		return fmt.Errorf("fetching partner sse gateway after update: %w", err)
	}
	if !found {
		return fmt.Errorf("partner sse gateway not found after successful update")
	}

	status := &krm.NetworkSecurityPartnerSSEGatewayStatus{}
	status.ObservedState = a.buildObservedState()
	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting networksecurity partner sse gateway", "name", a.id)

	url := fmt.Sprintf("https://networksecurity.googleapis.com/v1alpha1/%s", a.id.String())
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return false, fmt.Errorf("error creating DELETE request: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("error deleting partner sse gateway %q: %w", a.id.String(), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return true, nil
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return false, fmt.Errorf("deleting partner sse gateway %q failed with status %d: %s", a.id.String(), resp.StatusCode, string(body))
	}

	var op gcpOperation
	if err := json.NewDecoder(resp.Body).Decode(&op); err != nil {
		return false, fmt.Errorf("error decoding operation response: %w", err)
	}

	if err := a.waitForOperation(ctx, op.Name); err != nil {
		return false, fmt.Errorf("waiting for partner sse gateway deletion operation: %w", err)
	}

	return true, nil
}

func (a *adapter) waitForOperation(ctx context.Context, opName string) error {
	url := fmt.Sprintf("https://networksecurity.googleapis.com/v1alpha1/%s", opName)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				return fmt.Errorf("error creating operation GET request: %w", err)
			}
			resp, err := a.httpClient.Do(req)
			if err != nil {
				return fmt.Errorf("getting operation %q: %w", opName, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				body, _ := io.ReadAll(resp.Body)
				return fmt.Errorf("getting operation %q failed with status %d: %s", opName, resp.StatusCode, string(body))
			}

			var op gcpOperation
			if err := json.NewDecoder(resp.Body).Decode(&op); err != nil {
				return fmt.Errorf("error decoding operation response: %w", err)
			}

			if op.Done {
				if op.Error != nil {
					return fmt.Errorf("operation failed: %s", op.Error.Message)
				}
				return nil
			}
		}
	}
}

func (a *adapter) buildObservedState() *krm.NetworkSecurityPartnerSSEGatewayObservedState {
	if a.actual == nil {
		return nil
	}
	out := &krm.NetworkSecurityPartnerSSEGatewayObservedState{}
	if a.actual.CreateTime != "" {
		out.CreateTime = &a.actual.CreateTime
	}
	if a.actual.UpdateTime != "" {
		out.UpdateTime = &a.actual.UpdateTime
	}
	if a.actual.SseVpcSubnetRange != "" {
		out.SseVPCSubnetRange = &a.actual.SseVpcSubnetRange
	}
	if a.actual.SseVpcTargetIp != "" {
		out.SseVPCTargetIP = &a.actual.SseVpcTargetIp
	}
	if len(a.actual.SseBgpIps) > 0 {
		out.SseBGPIPs = a.actual.SseBgpIps
	}
	if a.actual.SseBgpAsn != 0 {
		out.SseBGPAsn = &a.actual.SseBgpAsn
	}
	if a.actual.PartnerSseRealm != "" {
		out.PartnerSSERealm = &a.actual.PartnerSseRealm
	}
	if a.actual.SseTargetIp != "" {
		out.SseTargetIP = &a.actual.SseTargetIp
	}
	if a.actual.SymantecOptionsStatus != nil {
		out.SymantecOptions = &krm.PartnerSSEGatewaySymantecOptionsObservedState{}
		if a.actual.SymantecOptionsStatus.SymantecLocationUuid != "" {
			out.SymantecOptions.SymantecLocationUuid = &a.actual.SymantecOptionsStatus.SymantecLocationUuid
		}
		if a.actual.SymantecOptionsStatus.SymantecSite != "" {
			out.SymantecOptions.SymantecSite = &a.actual.SymantecOptionsStatus.SymantecSite
		}
	}
	if a.actual.SseProject != "" {
		out.SseProject = &a.actual.SseProject
	}
	if a.actual.SseNetwork != "" {
		out.SseNetwork = &a.actual.SseNetwork
	}
	if a.actual.PartnerSseEnvironment != "" {
		out.PartnerSSEEnvironment = &a.actual.PartnerSseEnvironment
	}
	if a.actual.Country != "" {
		out.Country = &a.actual.Country
	}
	if a.actual.Timezone != "" {
		out.Timezone = &a.actual.Timezone
	}
	if a.actual.CapacityBps != "" {
		var val int64
		if _, err := fmt.Sscanf(a.actual.CapacityBps, "%d", &val); err == nil {
			out.CapacityBps = &val
		}
	}
	if a.actual.State != "" {
		out.State = &a.actual.State
	}
	if len(a.actual.ProberSubnetRanges) > 0 {
		out.ProberSubnetRanges = a.actual.ProberSubnetRanges
	}
	return out
}
