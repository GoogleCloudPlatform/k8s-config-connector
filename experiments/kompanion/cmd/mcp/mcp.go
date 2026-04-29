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

package mcp

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/version"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func BuildMCPCmd() *cobra.Command {
	var opts Options

	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Run the MCP server for KCC",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunMCP(cmd.Context(), &opts)
		},
	}

	opts.AddFlags(cmd)

	return cmd
}

func getRESTConfig(opts *Options) (*rest.Config, error) {
	var loadingRules clientcmd.ClientConfigLoader
	if opts.Kubeconfig != "" {
		loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: opts.Kubeconfig}
	} else {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	}

	configOverrides := &clientcmd.ConfigOverrides{}
	if opts.Context != "" {
		configOverrides.CurrentContext = opts.Context
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		configOverrides).ClientConfig()
}

type serverContext struct {
	dynamicClient   dynamic.Interface
	discoveryClient discovery.DiscoveryInterface

	mu              sync.RWMutex
	gvrCache        map[string]schema.GroupVersionResource
	gvkCache        map[schema.GroupVersionKind]schema.GroupVersionResource
	kindToGVRCache  map[string]schema.GroupVersionResource
	kccGVRs         []schema.GroupVersionResource
}

func RunMCP(ctx context.Context, opts *Options) error {
	config, err := getRESTConfig(opts)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}

	// We rely more on server-side rate limiting now, so give it a high client-side QPS
	if config.QPS == 0 {
		config.QPS = 100
		config.Burst = 200
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error building dynamic client: %w", err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return fmt.Errorf("error building discovery client: %w", err)
	}

	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	s := server.NewMCPServer(
		"kompanion-kcc",
		version.GetVersion(),
		server.WithLogging(),
	)

	// Tool 1: get_kcc_crd_schema
	s.AddTool(mcp.NewTool("get_kcc_crd_schema",
		mcp.WithDescription("Get the OpenAPI V3 schema for a KCC CRD kind"),
		mcp.WithString("kind", mcp.Required(), mcp.Description("The KRM kind (e.g. SQLInstance)")),
	), sc.handleGetKCCCRDSchema)

	// Tool 2: apply_kcc_yaml
	s.AddTool(mcp.NewTool("apply_kcc_yaml",
		mcp.WithDescription("Apply (Create/Update) KCC resources from YAML"),
		mcp.WithString("yaml", mcp.Required(), mcp.Description("The YAML manifest for the resource(s)")),
	), sc.handleApplyKCCYAML)

	// Tool 3: describe_kcc_resource
	s.AddTool(mcp.NewTool("describe_kcc_resource",
		mcp.WithDescription("Get the status and conditions of a KCC resource"),
		mcp.WithString("kind", mcp.Required(), mcp.Description("The KRM kind (e.g. SQLInstance)")),
		mcp.WithString("namespace", mcp.Description("The namespace of the resource (optional for cluster-scoped resources)")),
		mcp.WithString("name", mcp.Required(), mcp.Description("The name of the resource")),
	), sc.handleDescribeKCCResource)

	// Tool 4: list_kcc_resources
	s.AddTool(mcp.NewTool("list_kcc_resources",
		mcp.WithDescription("List KCC resources"),
		mcp.WithString("kind", mcp.Description("The KRM kind to filter by (optional)")),
		mcp.WithString("namespace", mcp.Description("The namespace to filter by (optional)")),
		mcp.WithNumber("limit", mcp.Description("Maximum number of resources to return (optional, default: 100)")),
	), sc.handleListKCCResources)

	// Tool 5: get_kcc_resource
	s.AddTool(mcp.NewTool("get_kcc_resource",
		mcp.WithDescription("Get the full YAML of a KCC resource"),
		mcp.WithString("kind", mcp.Required(), mcp.Description("The KRM kind (e.g. SQLInstance)")),
		mcp.WithString("namespace", mcp.Description("The namespace of the resource (optional for cluster-scoped resources)")),
		mcp.WithString("name", mcp.Required(), mcp.Description("The name of the resource")),
	), sc.handleGetKCCResource)

	// Tool 6: delete_kcc_resource
	s.AddTool(mcp.NewTool("delete_kcc_resource",
		mcp.WithDescription("Delete a KCC resource"),
		mcp.WithString("kind", mcp.Required(), mcp.Description("The KRM kind (e.g. SQLInstance)")),
		mcp.WithString("namespace", mcp.Description("The namespace of the resource (optional for cluster-scoped resources)")),
		mcp.WithString("name", mcp.Required(), mcp.Description("The name of the resource")),
	), sc.handleDeleteKCCResource)

	// Tool 7: list_kcc_kinds
	s.AddTool(mcp.NewTool("list_kcc_kinds",
		mcp.WithDescription("List all available KCC CRD kinds"),
	), sc.handleListKCCKinds)

	log.Printf("Starting MCP server for KCC on stdio...")
	if err := server.ServeStdio(s); err != nil {
		return fmt.Errorf("error serving mcp: %w", err)
	}
	return nil
}
