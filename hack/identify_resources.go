package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
)

// ResourceInfo represents the output structure for each resource.
type ResourceInfo struct {
	ControllerType     string `json:"controller_type"`
	ResourceName       string `json:"resource_name"`
	ServiceHostName    string `json:"service_hostname,omitempty"`
	APIEndpoint        string `json:"api_endpoint,omitempty"` // Mapped to REST Resource Name
	OfficialRESTAPIURL string `json:"official_rest_api_url,omitempty"`
	Kind               string `json:"kind"`
	Group              string `json:"group"`
}

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	resources := make(map[string]*ResourceInfo)

	// 1. Identify Resources from static_config
	for gvk, config := range resourceconfig.ControllerConfigStatic {
		var controllerType string
		switch config.DefaultController {
		case k8s.ReconcilerTypeTerraform:
			controllerType = "TF"
		case k8s.ReconcilerTypeDCL:
			controllerType = "DCL"
		default:
			continue
		}

		resources[gvk.Kind] = &ResourceInfo{
			ControllerType: controllerType,
			Kind:           gvk.Kind,
			Group:          gvk.Group,
		}
	}

	// 2. Enrich TF resources using servicemappingloader
	if err := enrichTFResources(resources); err != nil {
		fmt.Printf("Error enriching TF resources: %v\n", err)
		os.Exit(1)
	}

	// 3. Enrich DCL resources using metadata
	enrichDCLResources(resources)

	// 4. Enrich with Doc Templates (REST Resource Name & URL)
	if err := enrichWithDocs(rootDir, resources); err != nil {
		fmt.Printf("Error enriching with docs: %v\n", err)
		// We don't exit here, as this is optional enrichment
	}

	// 5. Filter out v1alpha1-only resources by checking CRDs
	if err := filterAlphaOnlyResources(rootDir, resources); err != nil {
		fmt.Printf("Error filtering alpha resources: %v\n", err)
		os.Exit(1)
	}

	// 6. Output JSON to resources.json
	outFile, err := os.Create("resources.json")
	if err != nil {
		fmt.Printf("Error creating resources.json: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	enc := json.NewEncoder(outFile)
	enc.SetIndent("", "  ")
	if err := enc.Encode(resources); err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully identified %d resources and saved to resources.json\n", len(resources))
}

func enrichTFResources(resources map[string]*ResourceInfo) error {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		return err
	}

	serviceMappings := smLoader.GetServiceMappings()
	for _, sm := range serviceMappings {
		for _, r := range sm.Spec.Resources {
			if res, ok := resources[r.Kind]; ok && res.ControllerType == "TF" {
				res.ResourceName = r.Name
				res.ServiceHostName = sm.Spec.ServiceHostName
			}
		}
	}
	return nil
}

func enrichDCLResources(resources map[string]*ResourceInfo) {
	serviceLoader := metadata.New()
	allServices := serviceLoader.GetAllServiceMetadata()

	for _, service := range allServices {
		for _, resource := range service.Resources {
			if res, ok := resources[resource.Kind]; ok && res.ControllerType == "DCL" {
				dclType := resource.DCLType
				if dclType == "" {
					dclType = resource.Kind
				}
				
				res.ResourceName = dclType
				res.ServiceHostName = service.ServiceNameUsedByDCL
			}
		}
	}
}

func enrichWithDocs(rootDir string, resources map[string]*ResourceInfo) error {
	templatesDir := filepath.Join(rootDir, "scripts/generate-google3-docs/resource-reference/templates")

	// Compile regexes once
	// Matches: <td>{{"{{gcp_name_short}}" "}} ... Resource Name</td>\n<td>(.*?)</td>
	// Simplified to look for "Resource Name" to avoid case/typo issues with "REST"
	reRestName := regexp.MustCompile(`(?is)Resource Name.*?<td>(.*?)</td>`)
	// Matches: <td>{{"{{gcp_name_short}}" "}} ... Resource Documentation</td>\n<td><a href="(.*?)">`)
	reRestDoc := regexp.MustCompile(`(?is)Resource Documentation.*?<a href="([^"]+)"`)

	for _, res := range resources {
		groupShort := strings.Split(res.Group, ".")[0]
		filename := fmt.Sprintf("%s_%s.tmpl", groupShort, strings.ToLower(res.Kind))
		path := filepath.Join(templatesDir, filename)

		contentBytes, err := ioutil.ReadFile(path)
		if err != nil {
			continue
		}
		content := string(contentBytes)

		if matches := reRestName.FindStringSubmatch(content); len(matches) > 1 {
			val := strings.TrimSpace(matches[1])
			val = strings.ReplaceAll(val, "<pre>", "")
			val = strings.ReplaceAll(val, "</pre>", "")
			val = strings.Join(strings.Fields(val), ", ") 
			res.APIEndpoint = val
		}

		if matches := reRestDoc.FindStringSubmatch(content); len(matches) > 1 {
			res.OfficialRESTAPIURL = matches[1]
		}
	}
	return nil
}

func filterAlphaOnlyResources(rootDir string, resources map[string]*ResourceInfo) error {
	crdDir := filepath.Join(rootDir, "config/crds/resources")
	files, err := ioutil.ReadDir(crdDir)
	if err != nil {
		return err
	}

	// Helper map to track which resources (by Kind) support v1beta1
	supportsBeta := make(map[string]bool)

	// Regex to extract kind. group is less critical if we assume kind is unique enough, 
	// but strictly we should check both.
	// CRD structure:
	// spec:
	//   group: ...
	//   names:
	//     kind: ...
	//   versions:
	//     - name: v1beta1
	
	reKind := regexp.MustCompile(`kind: (\w+)`)
	reBeta := regexp.MustCompile(`name: v1beta1`)

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".yaml" {
			continue
		}
		
		path := filepath.Join(crdDir, f.Name())
		contentBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(contentBytes)

		// Check if it supports v1beta1
		if !reBeta.MatchString(content) {
			continue
		}

		// Extract Kind to identify the resource
		// We use FindAllStringSubmatch because "kind: CustomResourceDefinition" appears at the top.
		// We want to find the resource kind (e.g. "ComputeInstance") which appears later.
		matches := reKind.FindAllStringSubmatch(content, -1)
		for _, match := range matches {
			if len(match) > 1 {
				kind := match[1]
				supportsBeta[kind] = true
			}
		}
	}

	// Filter resources: delete from map if kind is NOT in supportsBeta
	for kind := range resources {
		if !supportsBeta[kind] {
			delete(resources, kind)
		}
	}

	return nil
}