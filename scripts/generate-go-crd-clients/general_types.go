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
	"os"
	"sort"
	"strings"
)

type GeneralTypes struct {
	*resourceDefinition
	sb strings.Builder
}

func (g *GeneralTypes) WriteToFile(p string) error {
	b := []byte(g.sb.String())
	if err := os.WriteFile(p, b, 0644); err != nil {
		return fmt.Errorf("error writing file %q: %w", p, err)
	}
	return nil
}
func (g *GeneralTypes) Print(msg string, args ...interface{}) {
	if len(args) != 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	g.sb.WriteString(msg)
	g.sb.WriteString("\n")
}

func (g *GeneralTypes) Generate() {
	g.WriteHeader()

	g.Print("package %s", g.Version.Name)

	g.Print("import (\n")
	g.Print("\"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1\"")
	g.Print("metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"")
	g.Print("apiextensionsv1 \"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1\"")
	g.Print(")")
	g.Print("\nvar _ = apiextensionsv1.JSON{}\n")
	for _, structName := range sortedKeys(g.SpecNestedStructs) {
		fields := g.SpecNestedStructs[structName]
		g.Print("type %s struct {", structName)
		for i, f := range fields {
			if i != 0 {
				g.Print("")
			}
			g.structField(f)
		}
		g.Print("}")
		g.Print("")
	}

	g.Print("type %sSpec struct {", g.Name)
	for i, f := range g.SpecFields {
		if i != 0 {
			g.Print("")
		}
		g.structField(f)
	}
	g.Print("}")
	g.Print("")

	for _, structName := range sortedKeys(g.StatusNestedStructs) {
		fields := g.StatusNestedStructs[structName]
		g.Print("type %s struct {", structName)
		for i, f := range fields {
			if i != 0 {
				g.Print("")
			}
			g.structField(f)
		}
		g.Print("}")
		g.Print("")
	}

	g.Print("type %sStatus struct {", g.Name)
	g.Print("\t/* Conditions represent the latest available observations of the")
	g.Print("\t    %s's current state. */", g.Name)
	g.Print("Conditions []v1alpha1.Condition `json:\"conditions,omitempty\"`")

	for i, f := range g.StatusFields {
		if i != 0 {
			g.Print("")
		}
		g.structField(f)
	}
	g.Print("}")

	g.Print("// +genclient")
	g.Print("// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object")
	g.Print("// +kubebuilder:resource:categories=%s,shortName=%s",
		strings.Join(g.CRD.Spec.Names.Categories, ";"),
		strings.Join(g.CRD.Spec.Names.ShortNames, ";"))

	if g.Version.Subresources != nil {
		if g.Version.Subresources.Status != nil {
			g.Print("// +kubebuilder:subresource:status")
		}
	}

	// Output labels for kubebuilder
	{
		labels := make(map[string]string)
		for k, v := range g.CRD.Labels {
			labels[k] = v
		}

		var labelStrings []string
		for k, v := range labels {
			labelStrings = append(labelStrings, fmt.Sprintf("%q", k+"="+v))
		}
		sort.Strings(labelStrings)
		if len(labelStrings) != 0 {
			for _, label := range labelStrings {
				g.Print("// +kubebuilder:metadata:labels=" + label)
			}
		}
	}

	// Output additionalPrinterColumns for kubebuilder
	if additionalPrinterColumns := g.Version.AdditionalPrinterColumns; len(additionalPrinterColumns) != 0 {
		for _, c := range additionalPrinterColumns {
			spec := fmt.Sprintf("name=%q,JSONPath=%q", c.Name, c.JSONPath)
			if c.Type != "" {
				spec += fmt.Sprintf(",type=%q", c.Type)
			}
			if c.Format != "" {
				spec += fmt.Sprintf(",format=%q", c.Format)
			}
			if c.Priority != 0 {
				spec += fmt.Sprintf(",priority=%d", c.Priority)
			}
			if c.Description != "" {
				spec += fmt.Sprintf(",description=%q", c.Description)
			}
			g.Print("// +kubebuilder:printcolumn:" + spec)
		}
	}

	g.Print("")
	g.Print("// %s is the Schema for the %s API", g.Name, g.Service)
	g.Print("// +k8s:openapi-gen=true")
	g.Print("type %s struct {", g.Name)
	g.Print("  metav1.TypeMeta `json:\",inline\"`")
	g.Print("  metav1.ObjectMeta `json:\"metadata,omitempty\"`")
	g.Print("")
	g.Print("  Spec %sSpec `json:\"spec,omitempty\"`", g.Name)
	g.Print("  Status %sStatus `json:\"status,omitempty\"`", g.Name)
	g.Print("}")

	g.Print(" // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object")
	g.Print("")
	g.Print(" // %sList contains a list of %s", g.Name, g.Name)
	g.Print(" type %sList struct {", g.Name)
	g.Print("   metav1.TypeMeta `json:\",inline\"`")
	g.Print("   metav1.ListMeta `json:\"metadata,omitempty\"`")
	g.Print("   Items []%s `json:\"items\"`", g.Name)
	g.Print(" }")

	g.Print(" func init() {")
	g.Print("   SchemeBuilder.Register(&%s{}, &%sList{})", g.Name, g.Name)
	g.Print(" }")
}

func (g *GeneralTypes) structField(f *fieldProperties) {
	description := f.Description
	description = strings.TrimSpace(description)
	if description != "" {
		var lines []string
		for _, line := range strings.Split(description, "\n") {
			lines = append(lines, strings.TrimSpace(line))
		}
		if len(lines) == 1 {
			g.Print("/* %s */", strings.Join(lines, "\n"))
		} else {
			g.Print("\t/* %s */", strings.Join(lines, "\n\t"))
		}
	}
	if f.Optional {
		g.Print("// +optional")
	}
	typeName := f.Type
	if f.UsePointer {
		typeName = "*" + typeName
	}
	g.Print("%s %s `json:\"%s\"`", f.Name, typeName, f.JSONName)
}

func (g *GeneralTypes) WriteHeader() {
	header := `
// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!
`
	g.Print(header)
}

func sortedKeys[V any](m map[string]V) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
