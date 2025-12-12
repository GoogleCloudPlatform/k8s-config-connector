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

package mockgcp

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"regexp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/workflows"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/interceptor"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockapigee"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockapikeys"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockapphub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockasset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbatch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigquery"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigqueryanalyticshub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigquerybiglake"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigqueryconnection"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigquerydatatransfer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigqueryreservation"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcertificatemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudbuild"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockclouddeploy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudfunctions"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudidentity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudquota"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudtasks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcomposer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcontainer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcontaineranalysis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdataform"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdataplex"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdatastream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdocumentai"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockedgenetwork"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockeventarc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgkebackup"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgkehub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgkemulticloud"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockkms"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockmetastore"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockmodelarmor"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknetapp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknetworkconnectivity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknetworkmanagement"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknotebooks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockprivilegedaccessmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockpubsublite"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockrecaptchaenterprise"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockredis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockresourcemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocksecretmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocksecuresourcemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockservicenetworking"
	mockspanner "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockspanner/admin"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockspeech"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockvmwareengine"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockvpcaccess"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockworkflows"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockworkstations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type Interface interface {
	// We support HTTP requests
	http.RoundTripper

	// NewGRPCConnection returns a grpc connection to our mock implementation
	NewGRPCConnection(ctx context.Context) *grpc.ClientConn

	// Run starts the grpc service, until ctx is closed
	Run(ctx context.Context) error

	// We can dispatch test commands
	SupportsTestCommands
}

func NewMockRoundTripper(ctx context.Context, k8sClient client.Client, storage storage.Storage) (Interface, error) {
	log := klog.FromContext(ctx)

	mockRoundTripper := &mockRoundTripper{}
	mockHTTPClient := &http.Client{
		Transport: mockRoundTripper,
	}
	env := &common.MockEnvironment{
		KubeClient: k8sClient,
	}

	env.Projects = mockresourcemanager.NewProjectStore(storage)

	workflowEngine, err := workflows.NewEngine(mockHTTPClient)
	if err != nil {
		return nil, fmt.Errorf("building workflow engine: %w", err)
	}
	env.Workflows = workflowEngine

	var serverOpts []grpc.ServerOption
	serverOpts = append(serverOpts, grpc.UnaryInterceptor(interceptor.LabelValidationInterceptor))
	server := grpc.NewServer(serverOpts...)

	var services []mockgcpregistry.MockService

	registeredServices, err := mockgcpregistry.BuildAllServices(env, storage)
	if err != nil {
		return nil, err
	}
	mockRoundTripper.registeredServices = registeredServices

	for _, service := range registeredServices.Services {
		services = append(services, service)
	}

	services = append(services, mockasset.New(env, storage))
	services = append(services, mockapikeys.New(env, storage))
	services = append(services, mockmetastore.New(env, storage))
	services = append(services, mockbigquery.New(env, storage))
	services = append(services, mockcloudidentity.New(env, storage))
	services = append(services, mockcontainer.New(env, storage))
	services = append(services, mockcertificatemanager.New(env, storage))
	services = append(services, mockgkemulticloud.New(env, storage))
	services = append(services, mockmodelarmor.New(env, storage))
	services = append(services, mocknetworkmanagement.New(env, storage))
	services = append(services, mockclouddeploy.New(env, storage))
	services = append(services, mocksecretmanager.New(env, storage))
	services = append(services, mockspanner.New(env, storage))
	services = append(services, mockpubsublite.New(env, storage))
	services = append(services, mocknetworkconnectivity.New(env, storage))
	services = append(services, mocknotebooks.New(env, storage))
	services = append(services, mockprivilegedaccessmanager.New(env, storage))
	services = append(services, mockredis.New(env, storage))
	services = append(services, mocksecuresourcemanager.New(env, storage))
	services = append(services, mockservicenetworking.New(env, storage))
	services = append(services, mockcloudfunctions.New(env, storage))
	services = append(services, mockedgenetwork.New(env, storage))
	services = append(services, mockgkehub.New(env, storage))

	services = append(services, mockcloudbuild.New(env, storage))
	services = append(services, mockcontaineranalysis.New(env, storage))
	services = append(services, mockdataform.New(env, storage))
	services = append(services, mockbigqueryconnection.New(env, storage))

	services = append(services, mockworkstations.New(env, storage))
	services = append(services, mockbigquerydatatransfer.New(env, storage))
	services = append(services, mockbigqueryanalyticshub.New(env, storage))
	services = append(services, mockvpcaccess.New(env, storage))
	services = append(services, mockapigee.New(env, storage))
	services = append(services, mockbigqueryreservation.New(env, storage))
	services = append(services, mockworkflows.New(env, storage))
	services = append(services, mockcomposer.New(env, storage))
	services = append(services, mockdocumentai.New(env, storage))
	services = append(services, mockapphub.New(env, storage))
	services = append(services, mockcloudquota.New(env, storage))
	services = append(services, mockdatastream.New(env, storage))
	services = append(services, mockeventarc.New(env, storage))

	services = append(services, mockcloudtasks.New(env, storage))

	services = append(services, mockbatch.New(env, storage))

	services = append(services, mockbigquerybiglake.New(env, storage))
	services = append(services, mocknetapp.New(env, storage))
	services = append(services, mockdataplex.New(env, storage))
	services = append(services, mockvmwareengine.New(env, storage))
	services = append(services, mockkms.New(env, storage))
	services = append(services, mockgkebackup.New(env, storage))
	services = append(services, mockrecaptchaenterprise.New(env, storage))
	services = append(services, mockspeech.New(env, storage))

	for _, service := range services {
		service.Register(server)
	}

	mockRoundTripper.server = server

	// We listen on a random port on 127.0.0.2, to avoid conflicts with the webhook server which starts on a random port on "default" localhost
	listener, err := net.Listen("tcp", "127.0.0.2:0")
	if err != nil {
		return nil, fmt.Errorf("net.Listen failed: %w", err)
	}
	log.Info("serving mock gcp grpc server", "address", listener.Addr().String())
	mockRoundTripper.grpcListener = listener

	endpoint := listener.Addr().String()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return nil, fmt.Errorf("error dialing grpc endpoint %q: %v", endpoint, err)
	}
	mockRoundTripper.grpcConnection = conn

	for _, service := range services {
		mux, err := service.NewHTTPMux(ctx, conn)
		if err != nil {
			return nil, fmt.Errorf("error building mux: %v", err)
		}
		var hostRegexes []*regexp.Regexp
		for _, host := range service.ExpectedHosts() {
			hostRegexes = append(hostRegexes, toHostRegex(host))
		}
		mockRoundTripper.services = append(mockRoundTripper.services, registeredService{
			impl:        service,
			hostRegexes: hostRegexes,
			handler:     mux,
		})
	}

	mockRoundTripper.iamPolicies = newMockIAMPolicies()

	return mockRoundTripper, nil
}
