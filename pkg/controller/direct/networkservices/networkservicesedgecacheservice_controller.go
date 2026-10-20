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

package networkservices

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkservices/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/googleapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.NetworkServicesEdgeCacheServiceGVK, NewEdgeCacheServiceModel)
}

func NewEdgeCacheServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEdgeCacheService{config: config}, nil
}

type modelEdgeCacheService struct {
	config *config.ControllerConfig
}

func (m *modelEdgeCacheService) client(ctx context.Context) (*edgeCacheServiceClient, error) {
	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("building authenticated http client: %w", err)
	}
	return &edgeCacheServiceClient{
		httpClient: httpClient,
		endpoint:   "https://networkservices.googleapis.com",
		userAgent:  m.config.UserAgent,
	}, nil
}

func (m *modelEdgeCacheService) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkServicesEdgeCacheService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("converting unstructured to %T: %w", obj, err)
	}

	id, err := krm.NewNetworkServicesEdgeCacheServiceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	c, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &EdgeCacheServiceAdapter{
		id:      id,
		client:  c,
		desired: obj,
		reader:  reader,
	}, nil
}

func (m *modelEdgeCacheService) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type EdgeCacheServiceAdapter struct {
	id      *krm.NetworkServicesEdgeCacheServiceIdentity
	client  *edgeCacheServiceClient
	desired *krm.NetworkServicesEdgeCacheService
	reader  client.Reader
	actual  *pb.EdgeCacheService
}

var _ directbase.Adapter = &EdgeCacheServiceAdapter{}

func (a *EdgeCacheServiceAdapter) buildDesired(ctx context.Context, u *unstructured.Unstructured) (*pb.EdgeCacheService, error) {
	if err := common.NormalizeReferences(ctx, a.reader, a.desired, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := EdgeCacheServiceSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(u)

	// If actual already has the server-formatted origin URL matching the desired short name, adopt it
	if a.actual != nil && a.actual.Routing != nil && desired.Routing != nil {
		for _, actualPM := range a.actual.Routing.PathMatchers {
			for _, actualRR := range actualPM.RouteRules {
				for _, desiredPM := range desired.Routing.PathMatchers {
					if desiredPM.Name == actualPM.Name {
						for _, desiredRR := range desiredPM.RouteRules {
							if desiredRR.Priority == actualRR.Priority {
								if strings.HasSuffix(actualRR.Origin, "/"+desiredRR.Origin) {
									desiredRR.Origin = actualRR.Origin
								}
							}
						}
					}
				}
			}
		}
	}

	return desired, nil
}

func (a *EdgeCacheServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting EdgeCacheService", "name", a.id.String())

	actual, err := a.client.Get(ctx, a.id.String())
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting EdgeCacheService %q from GCP: %w", a.id.String(), err)
	}
	a.actual = actual
	return true, nil
}

func (a *EdgeCacheServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating EdgeCacheService", "name", a.id.String())

	u := createOp.GetUnstructured()
	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	op, err := a.client.Create(ctx, a.id.ParentString(), a.id.EdgeCacheService, desired)
	if err != nil {
		return fmt.Errorf("creating EdgeCacheService %q: %w", a.id.String(), err)
	}

	if err := a.client.WaitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for creation of EdgeCacheService %q: %w", a.id.String(), err)
	}

	latest, err := a.client.Get(ctx, a.id.String())
	if err != nil {
		return fmt.Errorf("getting latest EdgeCacheService %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *EdgeCacheServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating EdgeCacheService", "name", a.id.String())

	u := updateOp.GetUnstructured()
	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	diff, updateMask, err := common.CompareBrownfieldSpecAndLabels(
		ctx,
		u,
		&a.desired.Spec,
		a.actual,
		"labels",
		EdgeCacheServiceSpec_FromProto,
		EdgeCacheServiceSpec_ToProto,
		normalizeEdgeCacheService,
	)
	if err != nil {
		return fmt.Errorf("comparing %s: %w", a.id.String(), err)
	}

	if updateMask != nil && len(updateMask.Paths) > 0 {
		structuredreporting.ReportDiff(ctx, diff)

		jsonPaths := toJSONPaths(updateMask, diff)

		op, err := a.client.Patch(ctx, a.id.String(), desired, jsonPaths)
		if err != nil {
			return fmt.Errorf("updating EdgeCacheService %q: %w", a.id.String(), err)
		}

		if err := a.client.WaitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for update of EdgeCacheService %q: %w", a.id.String(), err)
		}
	}

	latest, err := a.client.Get(ctx, a.id.String())
	if err != nil {
		return fmt.Errorf("getting latest EdgeCacheService %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *EdgeCacheServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EdgeCacheService", "name", a.id.String())

	op, err := a.client.Delete(ctx, a.id.String())
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting EdgeCacheService %q: %w", a.id.String(), err)
	}

	if op != nil {
		if err := a.client.WaitForOperation(ctx, op); err != nil {
			return false, fmt.Errorf("waiting for deletion of EdgeCacheService %q: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *EdgeCacheServiceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkServicesEdgeCacheService{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EdgeCacheServiceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *EdgeCacheServiceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.EdgeCacheService) error {
	mapCtx := &direct.MapContext{}
	status := EdgeCacheServiceStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func normalizeEdgeCacheService(ctx context.Context, pb *pb.EdgeCacheService) error {
	if pb == nil {
		return nil
	}
	if pb.Routing != nil {
		for _, pm := range pb.Routing.PathMatchers {
			for _, rr := range pm.RouteRules {
				if rr.Origin != "" && strings.Contains(rr.Origin, "/") {
					rr.Origin = rr.Origin[strings.LastIndex(rr.Origin, "/")+1:]
				}
				if rr.RouteAction != nil && rr.RouteAction.CdnPolicy != nil {
					keyset := rr.RouteAction.CdnPolicy.SignedRequestKeyset
					if keyset != "" && strings.Contains(keyset, "/") {
						rr.RouteAction.CdnPolicy.SignedRequestKeyset = keyset[strings.LastIndex(keyset, "/")+1:]
					}
				}
			}
		}
	}
	return nil
}

func toJSONPaths(mask *fieldmaskpb.FieldMask, diff *structuredreporting.Diff) []string {
	if mask == nil {
		return nil
	}
	var jsonPaths []string
	for _, path := range mask.Paths {
		jsonPath := path
		if diff != nil {
			for _, f := range diff.Fields {
				if f.ID == path && f.ProtoFieldDescriptor != nil {
					jsonPath = f.ProtoFieldDescriptor.JSONName()
					break
				}
			}
		}
		jsonPaths = append(jsonPaths, jsonPath)
	}
	return jsonPaths
}

type edgeCacheServiceClient struct {
	httpClient *http.Client
	endpoint   string
	userAgent  string
}

type operation struct {
	Name     string          `json:"name"`
	Done     bool            `json:"done"`
	Error    *operationError `json:"error,omitempty"`
	Response json.RawMessage `json:"response,omitempty"`
}

type operationError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *edgeCacheServiceClient) do(ctx context.Context, method string, u *url.URL, body []byte) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode >= 400 {
		var gErr struct {
			Error struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			} `json:"error"`
		}
		if jsonErr := json.Unmarshal(respBody, &gErr); jsonErr == nil && gErr.Error.Code != 0 {
			return nil, &googleapi.Error{
				Code:    gErr.Error.Code,
				Message: gErr.Error.Message,
				Body:    string(respBody),
			}
		}
		return nil, &googleapi.Error{
			Code:    resp.StatusCode,
			Message: string(respBody),
			Body:    string(respBody),
		}
	}

	return respBody, nil
}

func (c *edgeCacheServiceClient) Get(ctx context.Context, name string) (*pb.EdgeCacheService, error) {
	u, err := url.Parse(fmt.Sprintf("%s/v1/%s", c.endpoint, name))
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("alt", "json")
	u.RawQuery = q.Encode()

	respBody, err := c.do(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	service := &pb.EdgeCacheService{}
	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(respBody, service); err != nil {
		return nil, fmt.Errorf("unmarshaling EdgeCacheService: %w", err)
	}
	return service, nil
}

func (c *edgeCacheServiceClient) Create(ctx context.Context, parent, edgeCacheServiceID string, desired *pb.EdgeCacheService) (*operation, error) {
	u, err := url.Parse(fmt.Sprintf("%s/v1/%s/edgeCacheServices", c.endpoint, parent))
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("alt", "json")
	q.Set("edgeCacheServiceId", edgeCacheServiceID)
	u.RawQuery = q.Encode()

	body, err := (protojson.MarshalOptions{UseProtoNames: false}).Marshal(desired)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdgeCacheService: %w", err)
	}

	respBody, err := c.do(ctx, http.MethodPost, u, body)
	if err != nil {
		return nil, err
	}

	op := &operation{}
	if err := json.Unmarshal(respBody, op); err != nil {
		return nil, fmt.Errorf("unmarshaling Operation: %w", err)
	}
	return op, nil
}

func (c *edgeCacheServiceClient) Patch(ctx context.Context, name string, desired *pb.EdgeCacheService, updateMask []string) (*operation, error) {
	u, err := url.Parse(fmt.Sprintf("%s/v1/%s", c.endpoint, name))
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("alt", "json")
	if len(updateMask) > 0 {
		q.Set("updateMask", strings.Join(updateMask, ","))
	}
	u.RawQuery = q.Encode()

	body, err := (protojson.MarshalOptions{UseProtoNames: false}).Marshal(desired)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdgeCacheService: %w", err)
	}

	respBody, err := c.do(ctx, http.MethodPatch, u, body)
	if err != nil {
		return nil, err
	}

	op := &operation{}
	if err := json.Unmarshal(respBody, op); err != nil {
		return nil, fmt.Errorf("unmarshaling Operation: %w", err)
	}
	return op, nil
}

func (c *edgeCacheServiceClient) Delete(ctx context.Context, name string) (*operation, error) {
	u, err := url.Parse(fmt.Sprintf("%s/v1/%s", c.endpoint, name))
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("alt", "json")
	u.RawQuery = q.Encode()

	respBody, err := c.do(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	op := &operation{}
	if err := json.Unmarshal(respBody, op); err != nil {
		return nil, fmt.Errorf("unmarshaling Operation: %w", err)
	}
	return op, nil
}

func (c *edgeCacheServiceClient) GetOperation(ctx context.Context, name string) (*operation, error) {
	u, err := url.Parse(fmt.Sprintf("%s/v1/%s", c.endpoint, name))
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("alt", "json")
	u.RawQuery = q.Encode()

	respBody, err := c.do(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	op := &operation{}
	if err := json.Unmarshal(respBody, op); err != nil {
		return nil, fmt.Errorf("unmarshaling Operation: %w", err)
	}
	return op, nil
}

func (c *edgeCacheServiceClient) WaitForOperation(ctx context.Context, op *operation) error {
	if op == nil || op.Done {
		if op != nil && op.Error != nil {
			return fmt.Errorf("operation failed with code %d: %s", op.Error.Code, op.Error.Message)
		}
		return nil
	}

	_, err := common.WaitForOperation(ctx, 2*time.Second, func(o *operation) (bool, error) {
		if o.Error != nil {
			return false, fmt.Errorf("operation failed with code %d: %s", o.Error.Code, o.Error.Message)
		}
		return o.Done, nil
	}, func() (*operation, error) {
		return c.GetOperation(ctx, op.Name)
	})
	return err
}
