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
	"io"
	"sort"
	"strings"
	"sync"

	"github.com/mark3labs/mcp-go/mcp"
	"golang.org/x/sync/errgroup"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/apimachinery/pkg/version"
	sigyaml "sigs.k8s.io/yaml"
)

func (sc *serverContext) handleGetKCCCRDSchema(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	kind, err := request.RequireString("kind")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing kind: %v", err)), nil
	}

	gvr, err := sc.findGVRByKind(kind)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to find GVR for kind %s: %v", kind, err)), nil
	}

	crdName := fmt.Sprintf("%s.%s", gvr.Resource, gvr.Group)
	crdGVR := schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}
	targetCRD, err := sc.dynamicClient.Resource(crdGVR).Get(ctx, crdName, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get CRD %s: %v", crdName, err)), nil
	}

	versions, found, _ := unstructured.NestedSlice(targetCRD.Object, "spec", "versions")
	if !found || len(versions) == 0 {
		return mcp.NewToolResultError(fmt.Sprintf("no versions found for CRD %s", kind)), nil
	}

	// Find the served version, preferably the latest one or the storage version
	var schemaObj interface{}
	var bestVersion map[string]interface{}
	for _, v := range versions {
		verMap, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		served, _ := verMap["served"].(bool)
		if !served {
			continue
		}

		if bestVersion == nil {
			bestVersion = verMap
			continue
		}

		isStorage, _ := verMap["storage"].(bool)
		bestIsStorage, _ := bestVersion["storage"].(bool)

		if isStorage && !bestIsStorage {
			bestVersion = verMap
			continue
		}

		if !isStorage && bestIsStorage {
			continue
		}

		// Otherwise prefer v1 over v1beta1, etc. using Kube-aware version strings
		bestName, _ := bestVersion["name"].(string)
		currName, _ := verMap["name"].(string)
		if version.CompareKubeAwareVersionStrings(currName, bestName) > 0 {
			bestVersion = verMap
		}
	}

	if bestVersion != nil {
		schemaObj, _, _ = unstructured.NestedFieldNoCopy(bestVersion, "schema", "openAPIV3Schema")
	}

	if schemaObj == nil {
		return mcp.NewToolResultError(fmt.Sprintf("no schema found for CRD %s", kind)), nil
	}

	// Keep it lean: extract only spec and status properties
	schemaMap, ok := schemaObj.(map[string]interface{})
	if ok {
		properties, found, _ := unstructured.NestedMap(schemaMap, "properties")
		if found {
			leanSchema := make(map[string]interface{})
			if spec, ok := properties["spec"]; ok {
				leanSchema["spec"] = spec
			}
			if status, ok := properties["status"]; ok {
				leanSchema["status"] = status
			}
			schemaObj = leanSchema
		}
	}

	schemaJSON, err := json.MarshalIndent(schemaObj, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal schema: %v", err)), nil
	}

	return mcp.NewToolResultText(string(schemaJSON)), nil
}

func (sc *serverContext) handleApplyKCCYAML(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	yamlStr, err := request.RequireString("yaml")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing yaml: %v", err)), nil
	}

	decoder := yaml.NewYAMLOrJSONDecoder(strings.NewReader(yamlStr), 4096)
	var results []string

	for {
		var obj unstructured.Unstructured
		if err := decoder.Decode(&obj); err != nil {
			if err == io.EOF {
				break
			}
			return mcp.NewToolResultError(fmt.Sprintf("failed to decode YAML: %v", err)), nil
		}

		if len(obj.Object) == 0 {
			continue
		}

		gvk := obj.GroupVersionKind()
		gvr, err := sc.findGVR(gvk)
		if err != nil {
			results = append(results, fmt.Sprintf("Failed to apply (%s/%s): failed to find GVR for %v: %v", obj.GetNamespace(), obj.GetName(), gvk, err))
			continue
		}

		namespace := obj.GetNamespace()
		name := obj.GetName()

		var res interface{}
		var applyErr error

		// Use Server-Side Apply (Patch with SSA) if name is specified.
		// If name is empty (e.g. generateName is used), fall back to Create.
		if name != "" {
			data, err := json.Marshal(obj.Object)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("failed to marshal object: %v", err)), nil
			}

			force := true
			res, applyErr = sc.dynamicClient.Resource(gvr).Namespace(namespace).Patch(ctx, name, types.ApplyPatchType, data, metav1.PatchOptions{
				FieldManager: "kompanion-mcp",
				Force:        &force,
			})
		} else {
			applyErr = fmt.Errorf("name is required for apply, but it is empty (perhaps generateName is used?)")
		}

		if applyErr != nil && (name == "" || apierrors.IsNotFound(applyErr) || strings.Contains(applyErr.Error(), "not found")) {
			res, applyErr = sc.dynamicClient.Resource(gvr).Namespace(namespace).Create(ctx, &obj, metav1.CreateOptions{
				FieldManager: "kompanion-mcp",
			})
		}

		if applyErr != nil {
			results = append(results, fmt.Sprintf("Failed to apply %s/%s (%s): %v", namespace, name, gvk.Kind, applyErr))
		} else {
			appliedObj, ok := res.(*unstructured.Unstructured)
			if ok {
				results = append(results, fmt.Sprintf("Successfully applied %s/%s (%s), resourceVersion: %s", namespace, name, gvk.Kind, appliedObj.GetResourceVersion()))
			} else {
				results = append(results, fmt.Sprintf("Successfully applied %s/%s (%s), but failed to get resource version", namespace, name, gvk.Kind))
			}
		}
	}

	if len(results) == 0 {
		return mcp.NewToolResultText("No resources found in YAML."), nil
	}

	return mcp.NewToolResultText(strings.Join(results, "\n")), nil
}

func (sc *serverContext) handleDescribeKCCResource(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	kind, err := request.RequireString("kind")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing kind: %v", err)), nil
	}
	namespace := request.GetString("namespace", "")
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing name: %v", err)), nil
	}

	gvr, err := sc.findGVRByKind(kind)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to find GVR for kind %s: %v", kind, err)), nil
	}

	obj, err := sc.dynamicClient.Resource(gvr).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get resource: %v", err)), nil
	}

	status, found, _ := unstructured.NestedMap(obj.Object, "status")
	if !found {
		return mcp.NewToolResultText(fmt.Sprintf("Resource %s/%s found but has no status field.", namespace, name)), nil
	}

	conditions, _, _ := unstructured.NestedSlice(status, "conditions")
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Resource: %s/%s (%s)\n", namespace, name, kind))
	if len(conditions) > 0 {
		sb.WriteString("Conditions:\n")
		for _, c := range conditions {
			cond, ok := c.(map[string]interface{})
			if !ok {
				continue
			}
			statusVal, _ := cond["status"].(string)
			reason, _ := cond["reason"].(string)
			message, _ := cond["message"].(string)
			typeVal, _ := cond["type"].(string)
			
			statusPrefix := ""
			if (typeVal == "Ready" || typeVal == "UpToDate" || typeVal == "ManagementFinished") && statusVal == "False" {
				statusPrefix = "⚠️  "
			}
			
			sb.WriteString(fmt.Sprintf("  - %sType: %s, Status: %s, Reason: %s\n", statusPrefix, typeVal, statusVal, reason))
			if message != "" {
				sb.WriteString(fmt.Sprintf("    Message: %s\n", message))
			}
		}
	} else {
		sb.WriteString("No conditions found.\n")
	}

	statusJSON, _ := json.MarshalIndent(status, "", "  ")
	sb.WriteString("\nFull Status:\n")
	sb.WriteString(string(statusJSON))

	return mcp.NewToolResultText(sb.String()), nil
}

func (sc *serverContext) handleGetKCCResource(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	kind, err := request.RequireString("kind")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing kind: %v", err)), nil
	}
	namespace := request.GetString("namespace", "")
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing name: %v", err)), nil
	}

	gvr, err := sc.findGVRByKind(kind)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to find GVR for kind %s: %v", kind, err)), nil
	}

	obj, err := sc.dynamicClient.Resource(gvr).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get resource: %v", err)), nil
	}

	// Remove managedFields to keep the output lean
	unstructured.RemoveNestedField(obj.Object, "metadata", "managedFields")

	yamlData, err := sigyaml.Marshal(obj.Object)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to marshal resource to YAML: %v", err)), nil
	}

	return mcp.NewToolResultText(string(yamlData)), nil
}

func (sc *serverContext) handleDeleteKCCResource(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	kind, err := request.RequireString("kind")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing kind: %v", err)), nil
	}
	namespace := request.GetString("namespace", "")
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Missing name: %v", err)), nil
	}

	gvr, err := sc.findGVRByKind(kind)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to find GVR for kind %s: %v", kind, err)), nil
	}

	err = sc.dynamicClient.Resource(gvr).Namespace(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to delete resource: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Successfully deleted %s/%s (%s)", namespace, name, kind)), nil
}

func (sc *serverContext) handleListKCCKinds(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	_, err := sc.getKCCGVRs()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get KCC GVRs: %v", err)), nil
	}

	sc.mu.RLock()
	var kinds []string
	for kind, gvr := range sc.kindToGVRCache {
		kinds = append(kinds, fmt.Sprintf("- %s (%s)", kind, gvr.Group))
	}
	sc.mu.RUnlock()
	sort.Strings(kinds)

	if len(kinds) == 0 {
		return mcp.NewToolResultText("No KCC CRDs found."), nil
	}

	return mcp.NewToolResultText("Available KCC Kinds:\n" + strings.Join(kinds, "\n")), nil
}

func (sc *serverContext) handleListKCCResources(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	kind := request.GetString("kind", "")
	namespace := request.GetString("namespace", "")
	limit := request.GetInt("limit", 100)
	if limit <= 0 {
		limit = 100
	}

	var gvrs []schema.GroupVersionResource
	if kind != "" {
		gvr, err := sc.findGVRByKind(kind)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to find GVR for kind %s: %v", kind, err)), nil
		}
		gvrs = append(gvrs, gvr)
	} else {
		var err error
		gvrs, err = sc.getKCCGVRs()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to get KCC resources: %v", err)), nil
		}
	}

	var results []string
	var warnings []string
	var mu sync.Mutex
	count := 0

	g, ctx := errgroup.WithContext(ctx)
	// Use a worker pool or just parallelize if there aren't too many GVRs.
	// KCC has ~150-200 GVRs, so parallelizing all at once might be okay but let's be cautious.
	// Actually, let's just use goroutines for each GVR.
	for _, gvr := range gvrs {
		gvr := gvr
		g.Go(func() error {
			mu.Lock()
			if count >= int(limit) {
				mu.Unlock()
				return nil
			}
			currentLimit := int64(limit - count)
			mu.Unlock()

			list, err := sc.dynamicClient.Resource(gvr).Namespace(namespace).List(ctx, metav1.ListOptions{
				Limit: currentLimit,
			})
			if err != nil {
				mu.Lock()
				warnings = append(warnings, fmt.Sprintf("Warning: failed to list %s: %v", gvr.Resource, err))
				mu.Unlock()
				return nil // Continue with other types
			}

			mu.Lock()
			defer mu.Unlock()
			for _, item := range list.Items {
				if count >= int(limit) {
					break
				}
				projectID := item.GetAnnotations()["cnrm.cloud.google.com/project-id"]
				if projectID == "" {
					projectID = "n/a"
				}
				
				statusStr := sc.getResourceStatusShort(&item)
				results = append(results, fmt.Sprintf("- Kind: %s, Namespace: %s, Name: %s, ProjectID: %s, Status: %s", item.GetKind(), item.GetNamespace(), item.GetName(), projectID, statusStr))
				count++
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list resources: %v", err)), nil
	}

	if len(results) == 0 {
		msg := "No KCC resources found."
		if len(warnings) > 0 {
			msg += "\n\nWarnings:\n" + strings.Join(warnings, "\n")
		}
		return mcp.NewToolResultText(msg), nil
	}

	sort.Strings(results)
	output := strings.Join(results, "\n")
	if count >= int(limit) {
		output += fmt.Sprintf("\n\n(Note: results truncated to limit of %d)", int(limit))
	}
	if len(warnings) > 0 {
		output += "\n\nWarnings:\n" + strings.Join(warnings, "\n")
	}

	return mcp.NewToolResultText(output), nil
}

func (sc *serverContext) findGVR(gvk schema.GroupVersionKind) (schema.GroupVersionResource, error) {
	sc.mu.RLock()
	gvr, ok := sc.gvkCache[gvk]
	sc.mu.RUnlock()
	if ok {
		if gvr.Resource == "" {
			return schema.GroupVersionResource{}, fmt.Errorf("GVR not found for %v (cached failure)", gvk)
		}
		return gvr, nil
	}

	// Try loading KCC GVRs if it looks like a KCC resource
	if strings.HasSuffix(gvk.Group, ".cnrm.cloud.google.com") {
		_, _ = sc.getKCCGVRs() // Ignore error, we will try direct discovery if it fails
		sc.mu.RLock()
		gvr, ok = sc.gvkCache[gvk]
		sc.mu.RUnlock()
		if ok {
			if gvr.Resource == "" {
				return schema.GroupVersionResource{}, fmt.Errorf("GVR not found for %v (cached failure)", gvk)
			}
			return gvr, nil
		}
	}

	apiResourceList, err := sc.discoveryClient.ServerResourcesForGroupVersion(gvk.GroupVersion().String())
	if err != nil {
		// Cache negative lookup
		sc.mu.Lock()
		sc.gvkCache[gvk] = schema.GroupVersionResource{}
		sc.mu.Unlock()
		return schema.GroupVersionResource{}, err
	}
	for _, apiResource := range apiResourceList.APIResources {
		if apiResource.Kind == gvk.Kind && !strings.Contains(apiResource.Name, "/") {
			gvr = schema.GroupVersionResource{
				Group:    gvk.Group,
				Version:  gvk.Version,
				Resource: apiResource.Name,
			}
			sc.mu.Lock()
			sc.gvkCache[gvk] = gvr
			sc.mu.Unlock()
			return gvr, nil
		}
	}

	// Cache negative lookup
	sc.mu.Lock()
	sc.gvkCache[gvk] = schema.GroupVersionResource{}
	sc.mu.Unlock()
	return schema.GroupVersionResource{}, fmt.Errorf("GVR not found for %v", gvk)
}

func (sc *serverContext) findGVRByKind(kind string) (schema.GroupVersionResource, error) {
	sc.mu.RLock()
	gvr, ok := sc.kindToGVRCache[kind]
	sc.mu.RUnlock()
	if ok {
		return gvr, nil
	}

	_, err := sc.getKCCGVRs()
	if err != nil {
		return schema.GroupVersionResource{}, err
	}

	sc.mu.RLock()
	gvr, ok = sc.kindToGVRCache[kind]
	sc.mu.RUnlock()
	if ok {
		return gvr, nil
	}

	return schema.GroupVersionResource{}, fmt.Errorf("GVR not found for kind %s", kind)
}

func (sc *serverContext) getResourceStatusShort(obj *unstructured.Unstructured) string {
	conditions, found, _ := unstructured.NestedSlice(obj.Object, "status", "conditions")
	if !found || len(conditions) == 0 {
		return "Unknown"
	}

	for _, c := range conditions {
		cond, ok := c.(map[string]interface{})
		if !ok {
			continue
		}
		typeVal, _ := cond["type"].(string)
		statusVal, _ := cond["status"].(string)
		reason, _ := cond["reason"].(string)

		if typeVal == "Ready" || typeVal == "UpToDate" {
			if statusVal == "True" {
				return "Ready"
			}
			return fmt.Sprintf("⚠️  Not Ready (%s)", reason)
		}
	}
	return "Unknown"
}

func (sc *serverContext) getKCCGVRs() ([]schema.GroupVersionResource, error) {
	sc.mu.RLock()
	if len(sc.kccGVRs) > 0 {
		gvrs := sc.kccGVRs
		sc.mu.RUnlock()
		return gvrs, nil
	}
	sc.mu.RUnlock()

	sc.mu.Lock()
	defer sc.mu.Unlock()

	// Double check
	if len(sc.kccGVRs) > 0 {
		return sc.kccGVRs, nil
	}

	apiResourceLists, err := sc.discoveryClient.ServerPreferredResources()
	if err != nil {
		return nil, err
	}

	var gvrs []schema.GroupVersionResource
	for _, apiResourceList := range apiResourceLists {
		if !strings.Contains(apiResourceList.GroupVersion, ".cnrm.cloud.google.com/") {
			continue
		}
		gv, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			continue
		}
		for _, apiResource := range apiResourceList.APIResources {
			if !strings.Contains(apiResource.Name, "/") { // skip subresources
				gvr := schema.GroupVersionResource{
					Group:    gv.Group,
					Version:  gv.Version,
					Resource: apiResource.Name,
				}
				gvrs = append(gvrs, gvr)
				sc.kindToGVRCache[apiResource.Kind] = gvr
				sc.gvkCache[schema.GroupVersionKind{
					Group:   gv.Group,
					Version: gv.Version,
					Kind:    apiResource.Kind,
				}] = gvr
			}
		}
	}
	sc.kccGVRs = gvrs
	return gvrs, nil
}
