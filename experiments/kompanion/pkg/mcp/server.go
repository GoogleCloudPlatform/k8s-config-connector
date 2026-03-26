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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/yaml"
)

type Server struct {
	mcpServer       *server.MCPServer
	dynamicClient   dynamic.Interface
	discoveryClient discovery.DiscoveryInterface
}

func NewServer(restConfig *rest.Config) (*Server, error) {
	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating dynamic client: %w", err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating discovery client: %w", err)
	}

	mcpServer := server.NewMCPServer(
		"kompanion-kcc",
		"1.0.0",
		server.WithLogging(),
	)

	s := &Server{
		mcpServer:       mcpServer,
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
	}

	s.registerTools()

	return s, nil
}

func (s *Server) registerTools() {
	// Tool 1: get_kcc_crd_schema
	s.mcpServer.AddTool(mcp.NewTool("get_kcc_crd_schema",
		mcp.WithDescription("Get the OpenAPI V3 schema for a KCC CRD."),
		mcp.WithString("kind", mcp.Description("The KCC resource kind (e.g., SQLInstance)."), mcp.Required()),
	), s.getKCCCRDSchema)

	// Tool 2: apply_kcc_yaml
	s.mcpServer.AddTool(mcp.NewTool("apply_kcc_yaml",
		mcp.WithDescription("Apply a KCC resource YAML to the cluster. Uses Server-Side Apply (SSA)."),
		mcp.WithString("yaml", mcp.Description("The KCC resource YAML to apply."), mcp.Required()),
	), s.applyKCCYAML)

	// Tool 3: describe_kcc_resource
	s.mcpServer.AddTool(mcp.NewTool("describe_kcc_resource",
		mcp.WithDescription("Get the status and conditions of a KCC resource."),
		mcp.WithString("kind", mcp.Description("The KCC resource kind."), mcp.Required()),
		mcp.WithString("namespace", mcp.Description("The namespace of the resource."), mcp.Required()),
		mcp.WithString("name", mcp.Description("The name of the resource."), mcp.Required()),
	), s.describeKCCResource)

	// Tool 4: list_kcc_resources
	s.mcpServer.AddTool(mcp.NewTool("list_kcc_resources",
		mcp.WithDescription("List KCC resources of a specific kind, optionally filtered by namespace."),
		mcp.WithString("kind", mcp.Description("The KCC resource kind."), mcp.Required()),
		mcp.WithString("namespace", mcp.Description("The namespace to list resources from (optional).")),
	), s.listKCCResources)
}

func (s *Server) Serve(ctx context.Context) error {
	return server.ServeStdio(s.mcpServer)
}

func (s *Server) findGVR(kind string) (schema.GroupVersionResource, error) {
	_, resources, err := s.discoveryClient.ServerGroupsAndResources()
	if err != nil {
		return schema.GroupVersionResource{}, err
	}

	for _, list := range resources {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			continue
		}
		for _, resource := range list.APIResources {
			if resource.Kind == kind && strings.Contains(gv.Group, "cnrm.cloud.google.com") {
				return gv.WithResource(resource.Name), nil
			}
		}
	}

	return schema.GroupVersionResource{}, fmt.Errorf("KCC kind %s not found", kind)
}

func (s *Server) getKCCCRDSchema(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	kind := request.Params.Arguments["kind"].(string)

	gvr, err := s.findGVR(kind)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	crdName := fmt.Sprintf("%s.%s", gvr.Resource, gvr.Group)
	crdGVR := schema.GroupVersionResource{
		Group:    "apiextensions.k8s.io",
		Version:  "v1",
		Resource: "customresourcedefinitions",
	}

	crd, err := s.dynamicClient.Resource(crdGVR).Get(ctx, crdName, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultError(fmt.Errorf("error fetching CRD %s: %w", crdName, err).Error()), nil
	}

	versions, _, _ := unstructured.NestedSlice(crd.Object, "spec", "versions")
	var schemaData interface{}
	for _, v := range versions {
		versionMap := v.(map[string]interface{})
		if versionMap["name"] == gvr.Version {
			schemaData = versionMap["schema"].(map[string]interface{})["openAPIV3Schema"]
			break
		}
	}

	if schemaData == nil {
		return mcp.NewToolResultError(fmt.Sprintf("schema for %s version %s not found", kind, gvr.Version)), nil
	}

	// Focus on spec and status as per instructions
	schemaMap := schemaData.(map[string]interface{})
	properties, _, _ := unstructured.NestedMap(schemaMap, "properties")
	leanProperties := make(map[string]interface{})
	if spec, ok := properties["spec"]; ok {
		leanProperties["spec"] = spec
	}
	if status, ok := properties["status"]; ok {
		leanProperties["status"] = status
	}

	schemaMap["properties"] = leanProperties

	schemaJSON, _ := json.MarshalIndent(schemaMap, "", "  ")
	return mcp.NewToolResultText(string(schemaJSON)), nil
}

func (s *Server) applyKCCYAML(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	yamlStr := request.Params.Arguments["yaml"].(string)

	var obj unstructured.Unstructured
	if err := yaml.Unmarshal([]byte(yamlStr), &obj); err != nil {
		return mcp.NewToolResultError(fmt.Errorf("error parsing YAML: %w", err).Error()), nil
	}

	gvr, err := s.findGVR(obj.GetKind())
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	data, err := json.Marshal(&obj)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	force := true
	_, err = s.dynamicClient.Resource(gvr).Namespace(obj.GetNamespace()).Patch(ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{
		FieldManager: "kompanion-mcp",
		Force:        &force,
	})

	if err != nil {
		return mcp.NewToolResultError(fmt.Errorf("error applying resource: %w", err).Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Successfully applied %s %s/%s", obj.GetKind(), obj.GetNamespace(), obj.GetName())), nil
}

func (s *Server) describeKCCResource(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	kind := request.Params.Arguments["kind"].(string)
	namespace := request.Params.Arguments["namespace"].(string)
	name := request.Params.Arguments["name"].(string)

	gvr, err := s.findGVR(kind)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	obj, err := s.dynamicClient.Resource(gvr).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultError(fmt.Errorf("error getting resource: %w", err).Error()), nil
	}

	status := obj.Object["status"]
	statusJSON, _ := json.MarshalIndent(status, "", "  ")

	return mcp.NewToolResultText(string(statusJSON)), nil
}

func (s *Server) listKCCResources(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	kind := request.Params.Arguments["kind"].(string)
	namespace, _ := request.Params.Arguments["namespace"].(string)

	gvr, err := s.findGVR(kind)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	var list *unstructured.UnstructuredList
	var listErr error
	if namespace != "" {
		list, listErr = s.dynamicClient.Resource(gvr).Namespace(namespace).List(ctx, metav1.ListOptions{})
	} else {
		list, listErr = s.dynamicClient.Resource(gvr).List(ctx, metav1.ListOptions{})
	}

	if listErr != nil {
		return mcp.NewToolResultError(fmt.Errorf("error listing resources: %w", listErr).Error()), nil
	}

	type resourceRef struct {
		Name        string            `json:"name"`
		Namespace   string            `json:"namespace"`
		Annotations map[string]string `json:"annotations,omitempty"`
	}

	var refs []resourceRef
	for _, item := range list.Items {
		refs = append(refs, resourceRef{
			Name:        item.GetName(),
			Namespace:   item.GetNamespace(),
			Annotations: item.GetAnnotations(),
		})
	}

	refsJSON, _ := json.MarshalIndent(refs, "", "  ")
	return mcp.NewToolResultText(string(refsJSON)), nil
}
