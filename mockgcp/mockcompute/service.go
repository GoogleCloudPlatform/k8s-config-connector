// Copyright 2022 Google LLC
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

package mockcompute

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func init() {
	mockgcpregistry.Register(New)
}

// MockService represents a mocked compute service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	*computeOperations
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) mockgcpregistry.MockService {
	s := &MockService{
		MockEnvironment:   env,
		storage:           storage,
		computeOperations: newComputeOperationsService(storage),
	}
	return s
}

func (s *MockService) ExpectedHosts() []string {
	// service attachment has host "www.googleapis.com"
	return []string{"compute.googleapis.com", "www.googleapis.com"}
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	pb.RegisterFutureReservationsServer(grpcServer, &FutureReservationsV1{MockService: s})

	pb.RegisterBackendBucketsServer(grpcServer, &backendBuckets{MockService: s})

	pb.RegisterExternalVpnGatewaysServer(grpcServer, &externalVPNGateways{MockService: s})

	pb.RegisterNetworkEdgeSecurityServicesServer(grpcServer, &networkEdgeSecurityServicesV1{MockService: s})

	pb.RegisterSecurityPoliciesServer(grpcServer, &SecurityPoliciesV1{MockService: s})

	pb.RegisterNetworksServer(grpcServer, &NetworksV1{MockService: s})
	pb.RegisterFirewallsServer(grpcServer, &FirewallsV1{MockService: s})
	pb.RegisterNetworkAttachmentsServer(grpcServer, &networkAttachmentsV1{MockService: s})
	pb.RegisterSubnetworksServer(grpcServer, &SubnetsV1{MockService: s})
	pb.RegisterVpnGatewaysServer(grpcServer, &VPNGatewaysV1{MockService: s})
	pb.RegisterTargetVpnGatewaysServer(grpcServer, &TargetVpnGatewaysV1{MockService: s})
	pb.RegisterTargetGrpcProxiesServer(grpcServer, &TargetGrpcProxyV1{MockService: s})

	pb.RegisterTargetHttpProxiesServer(grpcServer, &GlobalTargetHTTPProxiesV1{MockService: s})
	pb.RegisterRegionTargetHttpProxiesServer(grpcServer, &RegionalTargetHTTPProxiesV1{MockService: s})

	pb.RegisterTargetHttpsProxiesServer(grpcServer, &GlobalTargetHTTPSProxiesV1{MockService: s})
	pb.RegisterRegionTargetHttpsProxiesServer(grpcServer, &RegionalTargetHTTPSProxiesV1{MockService: s})

	pb.RegisterUrlMapsServer(grpcServer, &GlobalURLMapsV1{MockService: s})
	pb.RegisterRegionUrlMapsServer(grpcServer, &RegionalURLMapsV1{MockService: s})

	pb.RegisterRegionHealthChecksServer(grpcServer, &RegionalHealthCheckV1{MockService: s})
	pb.RegisterHealthChecksServer(grpcServer, &GlobalHealthCheckV1{MockService: s})

	pb.RegisterBackendServicesServer(grpcServer, &GlobalBackendServicesV1{MockService: s})
	pb.RegisterRegionBackendServicesServer(grpcServer, &RegionalBackendServicesV1{MockService: s})

	pb.RegisterDisksServer(grpcServer, &DisksV1{MockService: s})
	pb.RegisterRegionDisksServer(grpcServer, &RegionalDisksV1{MockService: s})

	pb.RegisterResourcePoliciesServer(grpcServer, &ResourcePoliciesV1{MockService: s})

	pb.RegisterRegionOperationsServer(grpcServer, &RegionalOperationsV1{MockService: s})
	pb.RegisterZoneOperationsServer(grpcServer, &ZonalOperationsV1{MockService: s})
	pb.RegisterGlobalOperationsServer(grpcServer, &GlobalOperationsV1{MockService: s})
	pb.RegisterGlobalOrganizationOperationsServer(grpcServer, &GlobalOrganizationOperationsV1{MockService: s})

	pb.RegisterNodeGroupsServer(grpcServer, &NodeGroupsV1{MockService: s})
	pb.RegisterNodeTemplatesServer(grpcServer, &NodeTemplatesV1{MockService: s})

	pb.RegisterAddressesServer(grpcServer, &RegionalAddressesV1{MockService: s})
	pb.RegisterGlobalAddressesServer(grpcServer, &GlobalAddressesV1{MockService: s})
	pb.RegisterSslCertificatesServer(grpcServer, &GlobalSSLCertificatesV1{MockService: s})
	pb.RegisterRegionSslCertificatesServer(grpcServer, &RegionalSSLCertificatesV1{MockService: s})
	pb.RegisterSslPoliciesServer(grpcServer, &GlobalSslPolicyV1{MockService: s})
	pb.RegisterTargetSslProxiesServer(grpcServer, &TargetSslProxyV1{MockService: s})
	pb.RegisterTargetTcpProxiesServer(grpcServer, &GlobalTargetTcpProxyV1{MockService: s})
	pb.RegisterRegionTargetTcpProxiesServer(grpcServer, &RegionalTargetTcpProxyV1{MockService: s})

	pb.RegisterRegionNetworkEndpointGroupsServer(grpcServer, &RegionNetworkEndpointGroupV1{MockService: s})
	pb.RegisterNetworkEndpointGroupsServer(grpcServer, &NetworkEndpointGroupV1{MockService: s})

	pb.RegisterRoutesServer(grpcServer, &RoutesV1{MockService: s})

	pb.RegisterServiceAttachmentsServer(grpcServer, &RegionalServiceAttachmentV1{MockService: s})

	pb.RegisterFirewallPoliciesServer(grpcServer, &FirewallPoliciesV1{MockService: s})

	pb.RegisterGlobalForwardingRulesServer(grpcServer, &GlobalForwardingRulesV1{MockService: s})
	pb.RegisterForwardingRulesServer(grpcServer, &RegionalForwardingRulesV1{MockService: s})

	pb.RegisterRoutersServer(grpcServer, &RoutersV1{MockService: s})

	pb.RegisterImagesServer(grpcServer, &ImagesV1{MockService: s})

	pb.RegisterInstancesServer(grpcServer, &InstancesV1{MockService: s})
	pb.RegisterInstanceTemplatesServer(grpcServer, &InstanceTemplatesV1{MockService: s})

	pb.RegisterInstanceGroupManagersServer(grpcServer, &instanceGroupManagers{MockService: s})
	pb.RegisterRegionInstanceGroupManagersServer(grpcServer, &regionInstanceGroupManagers{MockService: s})
	pb.RegisterInstanceGroupsServer(grpcServer, &InstanceGroups{MockService: s})

	pb.RegisterZonesServer(grpcServer, &ZonesV1{MockService: s})
	pb.RegisterReservationsServer(grpcServer, &ReservationsV1{MockService: s})
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux, err := httpmux.NewServeMux(ctx, conn, httpmux.Options{},
		pb.RegisterRoutesHandler)
	if err != nil {
		return nil, err
	}

	if err := pb.RegisterFutureReservationsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterInstanceGroupManagersHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterRegionInstanceGroupManagersHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterInstanceGroupsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterReservationsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterBackendBucketsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterBackendServicesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionBackendServicesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterExternalVpnGatewaysHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterNetworkEdgeSecurityServicesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterSecurityPoliciesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterNetworksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterFirewallsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterNetworkAttachmentsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterSubnetworksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterVpnGatewaysHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterTargetVpnGatewaysHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterTargetGrpcProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterTargetHttpProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionTargetHttpProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterTargetHttpsProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionTargetHttpsProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterSslPoliciesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterTargetSslProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterTargetTcpProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionTargetTcpProxiesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterUrlMapsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionUrlMapsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterNodeGroupsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterNodeTemplatesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterDisksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionDisksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterResourcePoliciesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterFirewallPoliciesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterForwardingRulesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterGlobalForwardingRulesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterRoutersHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterRegionOperationsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterZoneOperationsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterGlobalOperationsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterGlobalOrganizationOperationsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterAddressesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterGlobalAddressesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterRegionHealthChecksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterHealthChecksHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	// for ssl certs and the managedsslcerts
	if err := pb.RegisterSslCertificatesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterRegionSslCertificatesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterRegionNetworkEndpointGroupsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterNetworkEndpointGroupsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterServiceAttachmentsHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	if err := pb.RegisterImagesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterInstancesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterInstanceTemplatesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}
	if err := pb.RegisterZonesHandler(ctx, mux.ServeMux, conn); err != nil {
		return nil, err
	}

	// Returns slightly non-standard errors
	mux.RewriteError = func(ctx context.Context, error *httpmux.ErrorResponse) {
		// Does not return status (at least for 404)
		error.Status = ""
	}

	// Does not return Cache-Control header
	mux.RewriteHeaders = func(ctx context.Context, response http.ResponseWriter, payload proto.Message) {
		response.Header().Del("Cache-Control")
	}

	// Terraform uses the /beta/ endpoints, but we have protos only for v1.
	// Also, we probably want to be implementing the newer versions
	// as that makes it easier to move KCC to newer API versions.
	// So far, it seems that all of beta is a direct mapping to v1 - though
	// I'm sure eventually we'll find something that needs special handling.
	rewriteBetaToV1 := func(w http.ResponseWriter, r *http.Request) {
		u := r.URL
		u2 := *u
		changed := false
		if strings.HasPrefix(u.Path, "/compute/beta/") {
			u2.Path = "/compute/v1/" + strings.TrimPrefix(u.Path, "/compute/beta/")
			changed = true
		}
		if changed {
			r = httpmux.RewriteRequest(r, &u2)
		}

		var isLegacyHealthCheck bool
		var isHTTPS bool
		if strings.Contains(r.URL.Path, "/global/httpHealthChecks") {
			isLegacyHealthCheck = true
			isHTTPS = false
		} else if strings.Contains(r.URL.Path, "/global/httpsHealthChecks") {
			isLegacyHealthCheck = true
			isHTTPS = true
		}

		if isLegacyHealthCheck {
			if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
				bodyBytes, err := io.ReadAll(r.Body)
				if err == nil && len(bodyBytes) > 0 {
					translatedBody, err := transformLegacyToModernRequest(bodyBytes, isHTTPS)
					if err == nil {
						r.Body = io.NopCloser(bytes.NewBuffer(translatedBody))
						r.ContentLength = int64(len(translatedBody))
					}
				}
			}

			u2 := *r.URL
			if isHTTPS {
				u2.Path = strings.ReplaceAll(u2.Path, "/global/httpsHealthChecks", "/global/healthChecks")
			} else {
				u2.Path = strings.ReplaceAll(u2.Path, "/global/httpHealthChecks", "/global/healthChecks")
			}
			r = httpmux.RewriteRequest(r, &u2)
		}

		// Merge multiple 'paths' query parameters into a single comma-separated 'paths' parameter.
		// This is needed because the Compute API (and Terraform) can send multiple 'paths' parameters,
		// but our generated proto has 'paths' as a single string field, and grpc-gateway fails
		// if it sees multiple values for a non-repeated field.
		if r.URL.Query().Has("paths") {
			q := r.URL.Query()
			if paths := q["paths"]; len(paths) > 1 {
				u2 := *r.URL
				// We avoid q.Encode() because it sorts query parameters alphabetically,
				// which would cause diffs in the HTTP logs for many tests.
				// Instead we rebuild the RawQuery while maintaining the order.
				var newParts []string
				pathsHandled := false
				for _, part := range strings.Split(u2.RawQuery, "&") {
					key := part
					if i := strings.Index(part, "="); i >= 0 {
						key = part[:i]
					}
					if key == "paths" {
						if !pathsHandled {
							newParts = append(newParts, "paths="+url.QueryEscape(strings.Join(paths, ",")))
							pathsHandled = true
						}
						continue
					}
					newParts = append(newParts, part)
				}
				u2.RawQuery = strings.Join(newParts, "&")
				r = httpmux.RewriteRequest(r, &u2)
			}
		}

		if r.Method == http.MethodPatch && strings.Contains(r.URL.Path, "/routers/") {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				r.ContentLength = int64(len(bodyBytes))

				var data map[string]any
				if err := json.Unmarshal(bodyBytes, &data); err == nil {
					if _, ok := data["nats"]; ok {
						u := r.URL
						u2 := *u
						q := u2.Query()
						q.Set("natsSpecified", "true")
						u2.RawQuery = q.Encode()
						r = httpmux.RewriteRequest(r, &u2)
					}
				}
			}
		}

		isRouter := strings.Contains(r.URL.Path, "/routers")

		var captured *responseCapture
		var originalWriter http.ResponseWriter = w
		if isLegacyHealthCheck || isRouter {
			captured = &responseCapture{ResponseWriter: w}
			w = captured
		}

		mux.ServeHTTP(w, r)

		if captured != nil {
			bodyBytes := captured.body
			if isLegacyHealthCheck {
				if len(bodyBytes) > 0 && captured.code < 400 {
					bodyStr := string(bodyBytes)
					if isHTTPS {
						bodyStr = strings.ReplaceAll(bodyStr, "healthChecks/", "httpsHealthChecks/")
					} else {
						bodyStr = strings.ReplaceAll(bodyStr, "healthChecks/", "httpHealthChecks/")
					}
					bodyBytes = []byte(bodyStr)

					if transformed, err := transformModernToLegacyResponse(bodyBytes, isHTTPS); err == nil {
						bodyBytes = transformed
					}
				}
			}

			if isRouter {
				if len(bodyBytes) > 0 && captured.code < 400 {
					if rewritten, err := rewriteRouterResponse(bodyBytes); err == nil {
						bodyBytes = rewritten
					}
				}
			}

			originalWriter.Header().Set("Content-Length", strconv.Itoa(len(bodyBytes)))
			if captured.code != 0 {
				originalWriter.WriteHeader(captured.code)
			}
			originalWriter.Write(bodyBytes)
		}
	}

	return http.HandlerFunc(rewriteBetaToV1), nil
}

type responseCapture struct {
	http.ResponseWriter
	body []byte
	code int
}

func (w *responseCapture) WriteHeader(statusCode int) {
	w.code = statusCode
}

func (w *responseCapture) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return len(b), nil
}

func transformLegacyToModernRequest(body []byte, isHTTPS bool) ([]byte, error) {
	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	nested := make(map[string]any)
	if host, ok := data["host"]; ok {
		nested["host"] = host
		delete(data, "host")
	}
	if port, ok := data["port"]; ok {
		nested["port"] = port
		delete(data, "port")
	}
	if requestPath, ok := data["requestPath"]; ok {
		nested["requestPath"] = requestPath
		delete(data, "requestPath")
	}

	if isHTTPS {
		data["type"] = "HTTPS"
		if len(nested) > 0 {
			data["httpsHealthCheck"] = nested
		}
	} else {
		data["type"] = "HTTP"
		if len(nested) > 0 {
			data["httpHealthCheck"] = nested
		}
	}

	return json.Marshal(data)
}

func transformModernToLegacyResponse(body []byte, isHTTPS bool) ([]byte, error) {
	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if kind, ok := data["kind"].(string); ok {
		if kind == "compute#healthCheck" {
			if isHTTPS {
				data["kind"] = "compute#httpsHealthCheck"
				if nested, ok := data["httpsHealthCheck"].(map[string]any); ok {
					for k, v := range nested {
						data[k] = v
					}
					delete(data, "httpsHealthCheck")
				}
			} else {
				data["kind"] = "compute#httpHealthCheck"
				if nested, ok := data["httpHealthCheck"].(map[string]any); ok {
					for k, v := range nested {
						data[k] = v
					}
					delete(data, "httpHealthCheck")
				}
			}
			delete(data, "type")
		}
	}

	return json.Marshal(data)
}

func rewriteRouterResponse(bodyBytes []byte) ([]byte, error) {
	var data map[string]any
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return bodyBytes, nil
	}

	populateNatFields := func(router map[string]any) {
		if nats, ok := router["nats"].([]any); ok {
			for _, natVal := range nats {
				if nat, ok := natVal.(map[string]any); ok {
					if _, ok := nat["autoNetworkTier"]; !ok {
						nat["autoNetworkTier"] = "PREMIUM"
					}
					if _, ok := nat["effectiveTcpTimeWaitTimeoutSec"]; !ok {
						nat["effectiveTcpTimeWaitTimeoutSec"] = float64(120)
					}
					if _, ok := nat["endpointTypes"]; !ok {
						nat["endpointTypes"] = []any{"ENDPOINT_TYPE_VM"}
					}
					if _, ok := nat["type"]; !ok {
						nat["type"] = "PUBLIC"
					}
				}
			}
		}
	}

	if kind, ok := data["kind"].(string); ok {
		if kind == "compute#router" {
			populateNatFields(data)
		} else if kind == "compute#routerList" {
			if items, ok := data["items"].([]any); ok {
				for _, item := range items {
					if router, ok := item.(map[string]any); ok {
						populateNatFields(router)
					}
				}
			}
		}
	}

	return json.Marshal(data)
}
