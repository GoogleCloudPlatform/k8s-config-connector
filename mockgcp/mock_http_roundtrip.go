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
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/workflows"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockaiplatform"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockalloydb"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockapigateway"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockapigee"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockapikeys"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockapphub"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockartifactregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockasset"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbackupdr"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbatch"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigquery"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigqueryanalyticshub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigquerybiglake"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigqueryconnection"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigquerydatatransfer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigqueryreservation"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbigtable"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockbilling"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcertificatemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudbuild"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockclouddeploy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockclouddms"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudfunctions"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudidentity"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudids"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudquota"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcloudtasks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcomposer"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcompute"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcontainer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockcontaineranalysis"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdatacatalog"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdataflow"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdataform"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdataplex"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdataproc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdatastream"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdiscoveryengine"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockdocumentai"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockedgecontainer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockedgenetwork"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockessentialcontacts"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockeventarc"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockfilestore"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockfirestore"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgkebackup"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgkehub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgkemulticloud"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockiam"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockkms"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocklogging"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockmanagedkafka"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockmetastore"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockmodelarmor"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockmonitoring"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknetapp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknetworkconnectivity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknetworkmanagement"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknetworkservices"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocknotebooks"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockorgpolicy"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockprivateca"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockprivilegedaccessmanager"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockpubsub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockpubsublite"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockrecaptchaenterprise"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockredis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockresourcemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocksecretmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocksecuresourcemanager"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockservicedirectory"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockservicenetworking"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockserviceusage"
	mockspanner "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockspanner/admin"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockspeech"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocksql"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockstorage"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mocktpu"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockvmwareengine"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockvpcaccess"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockworkflowexecution"
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
	mockRoundTripper := &mockRoundTripper{}
	mockHTTPClient := &http.Client{
		Transport: mockRoundTripper,
	}
	env := &common.MockEnvironment{
		KubeClient: k8sClient,
	}

	workflowEngine, err := workflows.NewEngine(mockHTTPClient)
	if err != nil {
		return nil, fmt.Errorf("building workflow engine: %w", err)
	}
	env.Workflows = workflowEngine

	resourcemanagerService := mockresourcemanager.New(env, storage)
	env.Projects = resourcemanagerService.GetProjectStore()

	var serverOpts []grpc.ServerOption
	server := grpc.NewServer(serverOpts...)

	var services []mockgcpregistry.MockService

	services = append(services, resourcemanagerService)

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
	services = append(services, mockbilling.New(env, storage))
	services = append(services, mockcloudidentity.New(env, storage))
	services = append(services, mockcontainer.New(env, storage))
	services = append(services, mockcertificatemanager.New(env, storage))
	services = append(services, mockedgecontainer.New(env, storage))
	services = append(services, mockfirestore.New(env, storage))
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
	services = append(services, mockserviceusage.New(env, storage))
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
	services = append(services, mockessentialcontacts.New(env, storage))
	services = append(services, mockeventarc.New(env, storage))

	services = append(services, mockcloudtasks.New(env, storage))

	services = append(services, mockbatch.New(env, storage))

	services = append(services, mockbigquerybiglake.New(env, storage))
	services = append(services, mocknetapp.New(env, storage))
	services = append(services, mockdataplex.New(env, storage))
	services = append(services, mockclouddms.New(env, storage))
	services = append(services, mockvmwareengine.New(env, storage))
	services = append(services, mockkms.New(env, storage))
	services = append(services, mockgkebackup.New(env, storage))
	services = append(services, mockrecaptchaenterprise.New(env, storage))
	services = append(services, mocknetworkservices.New(env, storage))
	services = append(services, mockspeech.New(env, storage))

	for _, service := range services {
		service.Register(server)
	}

	mockRoundTripper.server = server

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("net.Listen failed: %w", err)
	}
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
