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

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/fielddesc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

var handwrittenIAMTypes = []string{
	iamv1beta1.IAMPolicyGVK.Kind,
	iamv1beta1.IAMPartialPolicyGVK.Kind,
	iamv1beta1.IAMPolicyMemberGVK.Kind,
	iamv1beta1.IAMAuditConfigGVK.Kind,
}

type fieldProperties struct {
	Name        string
	Type        string
	Description string
	JSONName    string
	FullName    string
	Optional    bool
	UsePointer  bool
}

type resourceDefinition struct {
	Name                string
	Service             string
	Kind                string
	SpecFields          []*fieldProperties
	StatusFields        []*fieldProperties
	SpecNestedStructs   map[string][]*fieldProperties
	StatusNestedStructs map[string][]*fieldProperties

	CRD     *apiextensions.CustomResourceDefinition
	Version *apiextensions.CustomResourceDefinitionVersion
}

type svkMap struct {
	Service string
	Version string
	Kinds   []string
	Defs    []resourceDefinition
}

func main() {
	var resources []*resourceDefinition
	registerKinds := make(map[string]*svkMap)
	crdsDir := repo.GetCRDsPath()
	crdsPath, err := filepath.Abs(crdsDir)
	if err != nil {
		log.Fatalf("error getting the absolute representation of path for directory '%v': %v", crdsDir, err)
	}
	crdFiles, err := ioutil.ReadDir(crdsPath)
	if err != nil {
		log.Fatalf("error reading directory '%v': %v", crdsPath, err)
	}

	for _, crdFile := range crdFiles {
		resources = append(resources, constructResourceDefinitions(crdsPath, crdFile.Name())...)
	}

	for _, rd := range resources {
		// organize resources by service/version, create folder if not present
		serviceVersionString := fmt.Sprintf("%v/%v", rd.Service, rd.Version.Name)
		if v, ok := registerKinds[serviceVersionString]; ok {
			v.Kinds = append(v.Kinds, rd.Name)
			v.Defs = append(v.Defs, *rd)
		} else {
			registerKinds[serviceVersionString] = &svkMap{
				Service: rd.Service,
				Version: rd.Version.Name,
				Kinds:   []string{rd.Name},
				Defs:    []resourceDefinition{*rd},
			}
		}
	}
	makeStructNamesUniquePerKind(registerKinds)

	// clear out all generated types files
	typesDir := repo.GetTypesGeneratedApisPath()
	if err := os.RemoveAll(typesDir); err != nil {
		log.Fatalf("error deleting dir %v: %v", typesDir, err)
	}
	if err := os.MkdirAll(typesDir, 0700); err != nil {
		log.Fatalf("error recreating dir %v: %v", typesDir, err)
	}

	// template execution
	typesTemplateDir := repo.GetTypesTemplatePath()
	// typesTemplateFile := path.Join(typesTemplateDir, "general_types.go.tmpl")
	for _, rd := range resources {
		serviceDir := path.Join(typesDir, rd.Service)
		checkAndCreateFolder(serviceDir)
		serviceVersionDir := path.Join(serviceDir, rd.Version.Name)
		checkAndCreateFolder(serviceVersionDir)
		serviceVersionString := fmt.Sprintf("%v/%v", rd.Service, rd.Version.Name)

		// create new file for generated types file
		typesFileName := fmt.Sprintf("%s_types.go", rd.Kind)

		{
			var gen GeneralTypes
			gen.resourceDefinition = rd
			gen.Generate()
			if err := gen.WriteToFile(path.Join(serviceVersionDir, typesFileName)); err != nil {
				log.Fatalf("error creating %s_types.go file: %v", rd.Kind, err)
			}
		}
		// executeTemplateWithResourceDefinition(f, typesTemplateFile, rd)

		// create doc.go file per service/version directory
		f, err := os.Create(path.Join(serviceVersionDir, "doc.go"))
		if err != nil {
			log.Fatalf("error creating %v doc.go file: %v", serviceVersionString, err)
		}
		docTemplateFile := path.Join(typesTemplateDir, "doc.go.tmpl")
		executeTemplateWithResourceDefinition(f, docTemplateFile, rd)

		// create group.go file per service directory
		f, err = os.Create(path.Join(serviceDir, "group.go"))
		if err != nil {
			log.Fatalf("error creating %v group.go file: %v", rd.Service, err)
		}
		groupTemplateFile := path.Join(typesTemplateDir, "group.go.tmpl")
		executeTemplateWithResourceDefinition(f, groupTemplateFile, rd)
	}

	// create register.go file per service directory
	for _, registerInfo := range registerKinds {
		sort.Strings(registerInfo.Kinds)
		serviceVersionDir := path.Join(typesDir, registerInfo.Service, registerInfo.Version)
		f, err := os.Create(path.Join(serviceVersionDir, "register.go"))
		if err != nil {
			log.Fatalf("error creating %v register.go file: %v", registerInfo.Service, err)
		}
		registerTemplateFile := path.Join(typesTemplateDir, "register.go.tmpl")
		executeTemplateWithResourceDefinition(f, registerTemplateFile, registerInfo)
	}

	// Copy over common package `k8s` to generated/pkg/apis/ folder
	// The `k8s` package has the shared `conditions` type definition
	k8sDir := path.Join(repo.GetClientGenerationPath(), "k8s")
	cmd := exec.Command("cp", "-r", k8sDir, typesDir)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error executing 'cp' command: %v", err)
	}
}

func makeStructNamesUniquePerKind(kindMap map[string]*svkMap) {
	for _, m := range kindMap { // Loop through each service/version entry
		for _, r := range m.Defs { // Loop through each resource kind entry
			// Because we will be modifying the nestedFields map, we need to decouple
			// the existing field names from the map pointer so that fields aren't
			// being renamed multiple times
			nestedSpecFields := getArrayOfNestedFieldKeys(r.SpecNestedStructs)
			nestedStatusFields := getArrayOfNestedFieldKeys(r.StatusNestedStructs)

			// Remove Service from Kind because Kinds will be unique in each service's packages
			resourceKind := strings.TrimPrefix(r.Kind, r.Service)

			for _, s := range nestedSpecFields {
				newStructName := fmt.Sprintf("%v%v", strings.Title(resourceKind), strings.Title(s))
				findAndReplaceInStructField(s, newStructName, r.SpecFields)
				findAndReplaceInNestedFields(s, newStructName, r.SpecNestedStructs)
			}
			for _, s := range nestedStatusFields {
				newStructName := fmt.Sprintf("%v%v%v", strings.Title(resourceKind), strings.Title(s), "Status")
				findAndReplaceInStructField(s, newStructName, r.StatusFields)
				findAndReplaceInNestedFields(s, newStructName, r.StatusNestedStructs)
			}
		}
	}
}

func getArrayOfNestedFieldKeys(m map[string][]*fieldProperties) []string {
	arr := make([]string, 0)
	for k := range m {
		arr = append(arr, k)
	}
	return arr
}

func findAndReplaceInStructField(old, new string, fields []*fieldProperties) {
	for i, f := range fields {
		if f.Name == old {
			switch f.Type {
			case "string", "bool", "int": // is literal, don't replace
				continue
			default:
				if strings.HasPrefix(f.Type, "[]") { // Type is list of object
					f.Type = fmt.Sprintf("[]%v", new)
				} else if strings.HasPrefix(f.Type, "map[string]") {
					f.Type = fmt.Sprintf("map[string]%v", new)
				} else {
					f.Type = new
				}
			}
		}
		fields[i] = f
	}
}

func findAndReplaceInNestedFields(old, new string, fieldMap map[string][]*fieldProperties) {
	for name, children := range fieldMap {
		if name == old {
			fieldMap[new] = children
			delete(fieldMap, old)
		}
		// Replace in field type in nested struct
		findAndReplaceInStructField(old, new, children)
	}
}

func executeTemplateWithResourceDefinition(file *os.File, filePath string, r interface{}) {
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Fatalf("parsing template file failed: %v", err)
	}
	if err := tmpl.Execute(file, r); err != nil {
		log.Fatalf("template execution failed: %v", err)
	}
}

func checkAndCreateFolder(dir string) {
	if _, err := os.Stat(dir); err != nil {
		if err := os.Mkdir(dir, 0700); err != nil {
			log.Fatalf("error creating folder %v: %v", dir, err)
		}
	}
}

func constructResourceDefinitions(crdsPath, crdFile string) []*resourceDefinition {
	crdFilePath, err := filepath.Abs(path.Join(crdsPath, crdFile))
	if err != nil {
		log.Fatalf("error getting the absolute representation of path for directory '%v': %v", crdFile, err)
	}
	crd, err := crdloader.FileToCRD(crdFilePath)
	if err != nil {
		log.Fatalf("error loading crd from filepath %v: %v", crdFilePath, err)
	}

	versionNames := sets.NewString()
	for _, crdVersion := range crd.Spec.Versions {
		versionNames.Insert(crdVersion.Name)
	}

	var resources []*resourceDefinition
	for _, versionName := range versionNames.List() {
		// Don't generate alpha version if we have a beta
		if versionName == "v1alpha1" && versionNames.Has("v1beta1") {
			continue
		}
		crdVersionDefinition := k8s.GetCRDVersionDefinition(crd, versionName)

		r := &resourceDefinition{}
		r.CRD = crd
		r.Name = crd.Spec.Names.Kind
		if err = buildFieldProperties(r, crd, crdVersionDefinition.Name); err != nil {
			log.Fatalf("error building field properties for %v: %v", r.Name, err)
		}
		r.Service = strings.TrimSuffix(crd.Spec.Group, k8s.APIDomainSuffix)
		r.Kind = strings.ToLower(crd.Spec.Names.Kind)

		r.Version = crdVersionDefinition
		resources = append(resources, r)
	}
	return resources
}

func buildFieldProperties(r *resourceDefinition, crd *apiextensions.CustomResourceDefinition, version string) error {
	specDesc := fielddesc.GetSpecDescription(crd, version)
	specDescriptions := dropRootAndFlattenChildrenDescriptions(specDesc)
	r.SpecNestedStructs = make(map[string][]*fieldProperties)
	organizeSpecFieldDescriptions(specDescriptions, r)
	statusDesc, err := fielddesc.GetStatusDescription(crd, version)
	if err != nil {
		return fmt.Errorf("error getting status descriptions: %w", err)
	}
	statusDescriptions := dropRootAndFlattenChildrenDescriptions(statusDesc)
	r.StatusNestedStructs = make(map[string][]*fieldProperties)
	organizeStatusFieldDescriptions(statusDescriptions, r)
	return nil
}

func organizeSpecFieldDescriptions(descriptions []fielddesc.FieldDescription, r *resourceDefinition) {
	for _, d := range descriptions {
		if d.ShortName == "[]" { // Field is just flatted from an array, ignore
			continue
		}

		isRef := isResourceReference(d)
		isSec := isSecretReference(d)
		if !isRef && !isSec { // Field is NOT a resourceRef, add children to nested structs
			if d.Type == "object" { // Field most likely has nested fields
				children := getChildrenFromDescription(d, r)
				r.SpecNestedStructs[strings.Title(d.ShortName)] = children
			}
			if d.Type == "list (object)" {
				children := getChildrenFromDescription(d.Children[0], r)
				r.SpecNestedStructs[strings.Title(d.ShortName)] = children
			}
			if d.Type == "map (key: string, value: object)" {
				additionalProperties := getAdditionalPropertiesFromDescription(d, r)
				r.SpecNestedStructs[strings.Title(d.ShortName)] = additionalProperties
			}
		}
		if len(d.FullName) > 2 {
			continue //field is nested & should not be listed in first-layer spec
		}
		r.SpecFields = append(r.SpecFields, fieldDescriptionToFieldProperties(d, isRef, isSec, r))
	}
}

func isResourceReference(d fielddesc.FieldDescription) bool {
	if d.ShortName == "secretKeyRef" {
		return false
	}
	if strings.HasSuffix(d.ShortName, "Ref") {
		return true
	}
	if len(d.Children) == 3 {
		for _, c := range d.Children {
			r := regexp.MustCompile("external|name|namespace")
			if r.MatchString(c.ShortName) {
				continue
			}

			return false
		}
		return true
	}
	if len(d.Children) == 1 && d.Children[0].ShortName == "[]" {
		return isResourceReference(d.Children[0])
	}
	return false
}

func isSecretReference(d fielddesc.FieldDescription) bool {
	if len(d.Children) == 2 {
		for _, c := range d.Children {
			if c.ShortName == "name" || c.ShortName == "key" {
				continue
			}

			return false
		}
		return d.ShortName == "secretKeyRef"
	}
	return false
}

func getChildrenFromDescription(d fielddesc.FieldDescription, r *resourceDefinition) []*fieldProperties {
	children := make([]*fieldProperties, 0)
	for _, c := range d.Children {
		children = append(children, fieldDescriptionToFieldProperties(c, isResourceReference(c), isSecretReference(c), r))
	}
	return children
}

func getAdditionalPropertiesFromDescription(d fielddesc.FieldDescription, r *resourceDefinition) []*fieldProperties {
	additionalProperties := make([]*fieldProperties, 0)
	for _, c := range d.AdditionalProperties {
		additionalProperties = append(additionalProperties, fieldDescriptionToFieldProperties(c, isResourceReference(c), isSecretReference(c), r))
	}
	return additionalProperties
}

func organizeStatusFieldDescriptions(descriptions []fielddesc.FieldDescription, r *resourceDefinition) {
	for _, d := range descriptions {
		if d.ShortName == "conditions" {
			continue // not defined in types file
		}
		if d.ShortName == "[]" {
			continue
		}
		isRef := isResourceReference(d)
		isSec := isSecretReference(d)
		if !isRef && !isSec {
			if d.Type == "object" {
				children := getChildrenFromDescription(d, r)
				r.StatusNestedStructs[strings.Title(d.ShortName)] = children
			}
			if d.Type == "list (object)" {
				children := getChildrenFromDescription(d.Children[0], r)
				r.StatusNestedStructs[strings.Title(d.ShortName)] = children
			}
			if d.Type == "map (key: string, value: object)" {
				additionalProperties := getAdditionalPropertiesFromDescription(d, r)
				r.StatusNestedStructs[strings.Title(d.ShortName)] = additionalProperties
			}
		}
		if len(d.FullName) > 2 {
			continue // field is nested
		}
		r.StatusFields = append(r.StatusFields, fieldDescriptionToFieldProperties(d, isRef, isSec, r))
	}
}

func fieldDescriptionToFieldProperties(desc fielddesc.FieldDescription, isRef bool, isSec bool, r *resourceDefinition) *fieldProperties {
	var isIAMRef bool
	if isRef {
		// Check if resource is IAMPolicy/PolicyMember/AuditConfig and modify ref to use IAMRef struct
		for _, v := range handwrittenIAMTypes {
			if r.Name == v {
				isIAMRef = true
				break
			}
		}
	}
	fp := &fieldProperties{
		FullName:    formatName(desc),
		Type:        formatType(desc, isRef, isSec, isIAMRef),
		Description: desc.Description,
		Name:        strings.Title(desc.ShortName),     // Field name UpperCamelCase
		JSONName:    fmt.Sprintf("%v", desc.ShortName), // ShortName is default lowerCamelCase, exclude omitempty unless the field is optional
	}
	// If fields are optional, they should use a pointer type in the Go definition (e.g. AwesomeFlag *SomeFlag) or have a built-in nil value (e.g. maps and slices).
	// See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#optional-vs-required for details.
	requirementLevel := formatRequirementLevel(desc)
	if requirementLevel == "Optional" {
		fp.Optional = true
		fp.JSONName = fmt.Sprintf("%v,omitempty", desc.ShortName)
		if !isMapOrSliceType(fp.Type) {
			fp.UsePointer = true
		}
	}
	fixFieldDescriptionContainsClosingTags(fp)
	return fp
}

func fixFieldDescriptionContainsClosingTags(fp *fieldProperties) {
	// If the field description contains "*/" substring, it will break golang comment syntax when placing in block comments /* ${field_description} */
	// TODO(b/191413779): As a temporary hacky workaround, we insert a space into "*/" and need to find a more proper way to address this.
	fp.Description = strings.ReplaceAll(fp.Description, "*/", "* /")
}

func isMapOrSliceType(t string) bool {
	if strings.HasPrefix(t, "[]") {
		return true
	}
	if strings.HasPrefix(t, "map[string]") {
		return true
	}
	return false
}

func formatName(desc fielddesc.FieldDescription) string {
	name := strings.Join(desc.FullName, ".")
	name = strings.TrimPrefix(name, "spec.")
	name = strings.TrimPrefix(name, "status.")
	return name
}

func formatType(desc fielddesc.FieldDescription, isRef, isSec, isIAMRef bool) string {
	switch desc.Type {
	case "boolean":
		return "bool"
	case "integer":
		switch desc.Format {
		case "int64":
			return "int64"
		case "int32":
			return "int32"
		case "":
			// The default is int64 (and not int, we don't want the schema to vary across architectures)
			return "int64"
		default:
			klog.Fatalf("unhandled case in formatType: %+v", desc)
			return ""
		}
	case "float", "number":
		return "float64"
	case "object":
		if isSec {
			return "v1alpha1.SecretKeyRef"
		}
		if isRef {
			if isIAMRef {
				return "v1alpha1.IAMResourceRef"
			}
			return "v1alpha1.ResourceRef"
		}

		return strings.Title(desc.ShortName)
	default:
		if strings.HasPrefix(desc.Type, "list (") {
			listType := strings.TrimSuffix(strings.TrimPrefix(desc.Type, "list ("), ")")
			switch listType {
			case "boolean", "string", "number", "integer":
				return fmt.Sprintf("[]%v", formatToGoLiteral(listType))
			default:
				if isRef {
					return fmt.Sprintf("[]v1alpha1.ResourceRef")
				}
				return fmt.Sprintf("[]%v", strings.Title(desc.ShortName))
			}
		}

		if strings.HasPrefix(desc.Type, "map (") {
			// map (key: string, value: <literal>) ->  map[string]<literal>
			literal := regexp.MustCompile("map \\(key: string, value: (.*?)\\)").FindStringSubmatch(desc.Type)
			valueType := literal[1]
			var goType string
			if valueType == "object" {
				goType = strings.Title(desc.ShortName)
			} else {
				goType = formatToGoLiteral(valueType)
			}
			return fmt.Sprintf("map[string]%v", goType)
		}
		return desc.Type
	}
}

func formatToGoLiteral(t string) string {
	switch t {
	case "string":
		return "string"
	case "boolean":
		return "bool"
	case "integer":
		return "int64"
	case "float", "number":
		return "float64"
	default:
		panic(fmt.Errorf("expected a JSONLiteral but got %v", t))
	}
}

// TODO(kcc-eng): Add feature to expose field requirement level
func formatRequirementLevel(desc fielddesc.FieldDescription) string {
	switch desc.RequirementLevel {
	case fielddesc.RequiredRequirementLevel:
		return "Required"
	case fielddesc.RequiredWhenParentPresentRequirementLevel:
		return "Required*"
	case fielddesc.OptionalRequirementLevel:
		return "Optional"
	default:
		panic(fmt.Errorf("unhandled requirement level: %v", desc.RequirementLevel))
	}
}

func dropRootAndFlattenChildrenDescriptions(rootDesc fielddesc.FieldDescription) []fielddesc.FieldDescription {
	result := flattenChildrenDescription(nil, rootDesc)
	return result[1:]
}

func flattenChildrenDescription(result []fielddesc.FieldDescription, fd fielddesc.FieldDescription) []fielddesc.FieldDescription {
	if result == nil {
		result = make([]fielddesc.FieldDescription, 0, 1)
	}
	result = append(result, fd)
	for _, child := range fd.Children {
		result = flattenChildrenDescription(result, child)
	}
	for _, child := range fd.AdditionalProperties {
		result = flattenChildrenDescription(result, child)
	}
	return result
}
