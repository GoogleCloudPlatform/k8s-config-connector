package main

import (
	"fmt"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"
	"sort"
	"strings"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Resources struct {
	GVK             schema.GroupVersionKind
	ReferenceFields []resourceconfig.ResourceReferenceConfig
}

func main() {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		log.Fatalf("error loading crds: %v", err)
	}

	var v1alpha1Resources, v1beta1Resources []Resources

	for _, crd := range crds {
		group := crd.Spec.Group
		kind := crd.Spec.Names.Kind

		for _, version := range crd.Spec.Versions {
			if version.Schema == nil || version.Schema.OpenAPIV3Schema == nil {
				continue
			}
			spec, ok := version.Schema.OpenAPIV3Schema.Properties["spec"]
			if !ok {
				continue
			}

			refFieldNames := findRefs(spec.Properties, "")
			if len(refFieldNames) > 0 {
				gvk := schema.GroupVersionKind{Group: group, Version: version.Name, Kind: kind}

				var referenceFields []resourceconfig.ResourceReferenceConfig
				for _, fieldName := range refFieldNames {
					referenceFields = append(referenceFields, resourceconfig.ResourceReferenceConfig{ReferenceFieldName: fieldName})
				}
				sort.Slice(referenceFields, func(i, j int) bool {
					return referenceFields[i].ReferenceFieldName < referenceFields[j].ReferenceFieldName
				})

				resources := Resources{
					GVK:             gvk,
					ReferenceFields: referenceFields,
				}
				if version.Name == "v1alpha1" {
					v1alpha1Resources = append(v1alpha1Resources, resources)
				} else if version.Name == "v1beta1" {
					v1beta1Resources = append(v1beta1Resources, resources)
				}
				sort.Slice(v1alpha1Resources, func(i, j int) bool {
					return v1alpha1Resources[i].GVK.Kind < v1alpha1Resources[j].GVK.Kind
				})
				sort.Slice(v1beta1Resources, func(i, j int) bool {
					return v1beta1Resources[i].GVK.Kind < v1beta1Resources[j].GVK.Kind
				})
			}
		}
	}

	generateReport("pkg/controller/resourceconfig/reference_config_alpha.go", v1alpha1Resources)
	generateReport("pkg/controller/resourceconfig/references_config.go", v1beta1Resources)
}

func findRefs(properties map[string]apiextensionsv1.JSONSchemaProps, prefix string) []string {
	var refs []string
	for name, prop := range properties {
		if strings.HasSuffix(name, "Ref") {
			refs = append(refs, prefix+name)
		}
		if strings.HasSuffix(name, "Refs") {
			refs = append(refs, prefix+name+"[]")
		}
		if len(prop.Properties) > 0 {
			refs = append(refs, findRefs(prop.Properties, prefix+name+".")...)
		}
		if prop.Items != nil && prop.Items.Schema != nil && len(prop.Items.Schema.Properties) > 0 {
			refs = append(refs, findRefs(prop.Items.Schema.Properties, prefix+name+"[].")...)
		}
	}
	return refs
}

func generateReport(filename string, resources []Resources) {
	var builder strings.Builder
	builder.WriteString("package resourceconfig\n\n")
	builder.WriteString("// TODO: The ReferenceMeta fields are not populated and need to be filled in.\n")
	if strings.Contains(filename, "alpha") {
		builder.WriteString("var ResourceReferencesAlpha = ResourceReferenceMap{\n")
	} else {
		builder.WriteString("var ResourceReferences = ResourceReferenceMap{\n")
	}

	for _, r := range resources {
		gvkString := fmt.Sprintf(`{Group: "%s", Version: "%s", Kind: "%s"}`, r.GVK.Group, r.GVK.Version, r.GVK.Kind)
		builder.WriteString(fmt.Sprintf("\t%s: {\n", gvkString))
		for _, field := range r.ReferenceFields {
			builder.WriteString("\t\t{\n")
			builder.WriteString(fmt.Sprintf("\t\t\tReferenceFieldName: \"%s\",\n", field.ReferenceFieldName))
			builder.WriteString("\t\t},\n")
		}
		builder.WriteString("\t},\n")
	}

	builder.WriteString("}\n")

	err := ioutil.WriteFile(filename, []byte(builder.String()), 0644)
	if err != nil {
		log.Fatalf("failed to write report to %s: %v", filename, err)
	}
	fmt.Printf("Generated %s\n", filename)
}
