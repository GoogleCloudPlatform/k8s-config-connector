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

package codegen

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
)

// RenderIdentityFile renders <kind>_identity.generated.go, omitting any declaration
// already present in overrides. It returns (_, false, nil) if all declarations are
// already provided by hand-written files.
func (g *IdentityGenerator) RenderIdentityFile(plan *IdentityPlan, overrides *HandWrittenOverrides) ([]byte, bool, error) {
	if plan.RawPattern == "" {
		return nil, false, nil
	}
	ov := overrides
	if ov == nil {
		ov = &HandWrittenOverrides{}
	}
	if !shouldEmitIdentityFile(plan, ov) {
		return nil, false, nil
	}

	kind := plan.Kind
	idType := kind + "Identity"
	formatVar := kind + "IdentityFormat"
	fromSpecFn := "getIdentityFrom" + kind + "Spec"

	var body bytes.Buffer
	renderIdentityInterfaceAssertions(&body, plan, ov, idType)
	renderIdentityStructAndMethods(&body, plan, ov, idType, formatVar)
	if !ov.HasGetIdentityFromSpec {
		renderGetIdentityFromSpec(&body, plan, fromSpecFn, idType)
	}
	if !ov.ResourceMethods["GetIdentity"] {
		renderGetIdentityMethod(&body, plan, fromSpecFn, idType)
	}

	out, err := formatGeneratedGoFile(plan.PackageName, IdentityFileAnnotationKey, plan, body.Bytes())
	return out, err == nil, err
}

// RenderReferenceFile renders <kind>_reference.generated.go, omitting any declaration
// already present in overrides. It returns (_, false, nil) if all declarations are
// already provided by hand-written files.
func (g *IdentityGenerator) RenderReferenceFile(plan *IdentityPlan, overrides *HandWrittenOverrides) ([]byte, bool, error) {
	if plan.RawPattern == "" {
		return nil, false, nil
	}
	ov := overrides
	if ov == nil {
		ov = &HandWrittenOverrides{}
	}
	if !shouldEmitReferenceFile(ov) {
		return nil, false, nil
	}

	var body bytes.Buffer
	renderReferenceStructAndGVK(&body, plan, ov)
	renderReferenceMethods(&body, plan, ov)

	out, err := formatGeneratedGoFile(plan.PackageName, ReferenceFileAnnotationKey, plan, body.Bytes())
	return out, err == nil, err
}

// RenderIdentityTestFile renders <kind>_identity_generated_test.go calling identity.AssertConformance.
func (g *IdentityGenerator) RenderIdentityTestFile(plan *IdentityPlan) ([]byte, error) {
	idType := plan.Kind + "Identity"
	refType := plan.Kind + "Ref"
	sampleExternal := sampleSegmentsString(plan.Segments)
	wantParent := sampleSegmentsString(plan.ParentSegments)

	var body bytes.Buffer
	fmt.Fprintf(&body, "func Test%s_Conformance(t *testing.T) {\n", idType)
	fmt.Fprintf(&body, "\tidentity.AssertConformance(t, identity.ConformanceCase[%s]{\n", idType)
	fmt.Fprintf(&body, "\t\tSampleExternal: %q,\n", sampleExternal)
	fmt.Fprintf(&body, "\t\tWantHost:       %q,\n", plan.Host)
	fmt.Fprintf(&body, "\t\tWantParent:     %q,\n", wantParent)
	if plan.ServerGeneratedID {
		body.WriteString("\t\tExpectServerGenerated: true,\n")
	}
	fmt.Fprintf(&body, "\t\tWantIdentity: &%s{\n", idType)
	for _, f := range plan.Fields {
		fmt.Fprintf(&body, "\t\t\t%s: %q,\n", f.StructField, f.SampleValue)
	}
	body.WriteString("\t\t},\n")
	fmt.Fprintf(&body, "\t\tRef: &%s{},\n", refType)
	body.WriteString("\t})\n")
	body.WriteString("}\n")

	return formatGeneratedGoFile(plan.PackageName, IdentityFileAnnotationKey, plan, body.Bytes())
}

func shouldEmitIdentityFile(plan *IdentityPlan, ov *HandWrittenOverrides) bool {
	if ov.HasIdentityStruct && ov.ResourceMethods["GetIdentity"] {
		return false
	}
	if plan.HasUnresolvedAncestors && !ov.HasGetIdentityFromSpec {
		return false
	}
	needHasID := plan.ServerGeneratedID && !ov.IdentityMethods["HasIdentitySpecified"]
	return !ov.HasIdentityFormat || !ov.HasIdentityStruct || needHasID ||
		!ov.IdentityMethods["String"] || !ov.IdentityMethods["FromExternal"] ||
		!ov.IdentityMethods["Host"] || !ov.HasGetIdentityFromSpec || !ov.ResourceMethods["GetIdentity"]
}

func renderIdentityInterfaceAssertions(w *bytes.Buffer, plan *IdentityPlan, ov *HandWrittenOverrides, idType string) {
	needStruct := !ov.HasIdentityStruct
	needGetIdentity := !ov.ResourceMethods["GetIdentity"]
	if !needStruct && !needGetIdentity {
		return
	}
	w.WriteString("var (\n")
	if needStruct {
		if plan.ServerGeneratedID {
			fmt.Fprintf(w, "\t_ identity.ServerGeneratedIdentity = &%s{}\n", idType)
		} else {
			fmt.Fprintf(w, "\t_ identity.IdentityV2 = &%s{}\n", idType)
		}
	}
	if needGetIdentity {
		fmt.Fprintf(w, "\t_ identity.Resource = &%s{}\n", plan.Kind)
	}
	w.WriteString(")\n\n")
}

func renderIdentityStructAndMethods(w *bytes.Buffer, plan *IdentityPlan, ov *HandWrittenOverrides, idType, formatVar string) {
	kind := plan.Kind
	if !ov.HasIdentityFormat {
		fmt.Fprintf(w, "var %s = gcpurls.Template[%s](%q, %q)\n\n",
			formatVar, idType, plan.Host, plan.TemplatePattern)
	}
	if !ov.HasIdentityStruct {
		fmt.Fprintf(w, "// %s is the identity of a GCP %s resource.\n", idType, kind)
		w.WriteString("// +k8s:deepcopy-gen=false\n")
		fmt.Fprintf(w, "type %s struct {\n", idType)
		for _, f := range plan.Fields {
			fmt.Fprintf(w, "\t%s string\n", f.StructField)
		}
		w.WriteString("}\n\n")
	}
	if plan.ServerGeneratedID && !ov.IdentityMethods["HasIdentitySpecified"] && plan.ResourceIDField != nil {
		fmt.Fprintf(w, "func (i *%s) HasIdentitySpecified() bool {\n", idType)
		fmt.Fprintf(w, "\treturn i.%s != \"\"\n", plan.ResourceIDField.StructField)
		w.WriteString("}\n\n")
	}
	if !ov.IdentityMethods["String"] {
		fmt.Fprintf(w, "func (i *%s) String() string {\n", idType)
		fmt.Fprintf(w, "\treturn %s.ToString(*i)\n", formatVar)
		w.WriteString("}\n\n")
	}
	if !ov.IdentityMethods["FromExternal"] {
		fmt.Fprintf(w, "func (i *%s) FromExternal(ref string) error {\n", idType)
		fmt.Fprintf(w, "\tparsed, match, err := %s.Parse(ref)\n", formatVar)
		w.WriteString("\tif err != nil {\n")
		fmt.Fprintf(w, "\t\treturn fmt.Errorf(\"format of %s external=%%q was not known (use %%s): %%w\", ref, %s.CanonicalForm(), err)\n", kind, formatVar)
		w.WriteString("\t}\n")
		w.WriteString("\tif !match {\n")
		fmt.Fprintf(w, "\t\treturn fmt.Errorf(\"format of %s external=%%q was not known (use %%s)\", ref, %s.CanonicalForm())\n", kind, formatVar)
		w.WriteString("\t}\n\n")
		w.WriteString("\t*i = *parsed\n")
		w.WriteString("\treturn nil\n")
		w.WriteString("}\n\n")
	}
	if !ov.IdentityMethods["Host"] {
		fmt.Fprintf(w, "func (i *%s) Host() string {\n", idType)
		fmt.Fprintf(w, "\treturn %s.Host()\n", formatVar)
		w.WriteString("}\n\n")
	}
	if !ov.IdentityMethods["ParentString"] {
		fmt.Fprintf(w, "func (i *%s) ParentString() string {\n", idType)
		fmt.Fprintf(w, "\treturn %s\n", renderParentStringExpr(plan.ParentSegments))
		w.WriteString("}\n\n")
	}
}

func renderParentStringExpr(segs []IdentitySegment) string {
	if len(segs) == 0 {
		return `""`
	}
	var parts []string
	var litBuf []string
	for idx, seg := range segs {
		if !seg.IsVariable {
			litBuf = append(litBuf, seg.Literal)
			continue
		}
		if len(litBuf) > 0 {
			prefix := strings.Join(litBuf, "/") + "/"
			if idx > len(litBuf) {
				prefix = "/" + prefix
			}
			parts = append(parts, fmt.Sprintf("%q", prefix))
			litBuf = nil
		} else if idx > 0 {
			parts = append(parts, `"/"`)
		}
		parts = append(parts, "i."+seg.Field.StructField)
	}
	if len(litBuf) > 0 {
		suffix := strings.Join(litBuf, "/")
		if len(parts) > 0 {
			suffix = "/" + suffix
		}
		parts = append(parts, fmt.Sprintf("%q", suffix))
	}
	if len(parts) == 0 {
		return `""`
	}
	return strings.Join(parts, " + ")
}

func renderGetIdentityFromSpec(w *bytes.Buffer, plan *IdentityPlan, fnName, idType string) {
	fmt.Fprintf(w, "func %s(ctx context.Context, reader client.Reader, obj *%s) (*%s, error) {\n", fnName, plan.Kind, idType)
	renderResourceIDLookup(w, plan)
	if plan.ParentRef != nil {
		renderParentRefIdentityFromSpec(w, plan, idType)
		return
	}

	fieldVars := map[string]string{}
	if plan.ResourceIDField != nil {
		fieldVars[plan.ResourceIDField.StructField] = "resourceID"
	}
	for _, anc := range plan.Ancestors {
		fieldVars[anc.Field.StructField] = renderSingleAncestorLookup(w, plan, anc)
	}

	fmt.Fprintf(w, "\treturn &%s{\n", idType)
	for _, f := range plan.Fields {
		fmt.Fprintf(w, "\t\t%s: %s,\n", f.StructField, fieldVars[f.StructField])
	}
	w.WriteString("\t}, nil\n")
	w.WriteString("}\n\n")
}

func renderResourceIDLookup(w *bytes.Buffer, plan *IdentityPlan) {
	if plan.ResourceIDField == nil {
		return
	}
	if !plan.ServerGeneratedID {
		w.WriteString("\tresourceID, err := refs.GetResourceID(obj)\n")
		w.WriteString("\tif err != nil {\n")
		w.WriteString("\t\treturn nil, fmt.Errorf(\"cannot resolve resource ID: %w\", err)\n")
		w.WriteString("\t}\n\n")
		return
	}
	if plan.SpecIDIsPointer {
		w.WriteString("\tresourceID := common.ValueOf(obj.Spec.ResourceID)\n\n")
	} else {
		w.WriteString("\tresourceID := obj.Spec.ResourceID\n\n")
	}
}

func renderParentRefIdentityFromSpec(w *bytes.Buffer, plan *IdentityPlan, idType string) {
	pr := plan.ParentRef
	if pr.IsPointer {
		fmt.Fprintf(w, "\tif obj.Spec.%s == nil {\n", pr.SpecFieldName)
		fmt.Fprintf(w, "\t\treturn nil, fmt.Errorf(\"spec.%s is required\")\n", pr.SpecJSONName)
		w.WriteString("\t}\n")
		fmt.Fprintf(w, "\tparentRef := *obj.Spec.%s\n", pr.SpecFieldName)
	} else {
		fmt.Fprintf(w, "\tparentRef := obj.Spec.%s\n", pr.SpecFieldName)
	}
	w.WriteString("\tif err := parentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {\n")
	fmt.Fprintf(w, "\t\treturn nil, fmt.Errorf(\"resolving spec.%s: %%w\", err)\n", pr.SpecJSONName)
	w.WriteString("\t}\n")

	tailExpr := renderTailSuffixExpr(pr.TailSegments, plan.ResourceIDField != nil, plan.ServerGeneratedID)
	fmt.Fprintf(w, "\tparsed := &%s{}\n", idType)
	fmt.Fprintf(w, "\tif err := parsed.FromExternal(strings.TrimSuffix(parentRef.GetExternal(), \"/\") + %s); err != nil {\n", tailExpr)
	fmt.Fprintf(w, "\t\treturn nil, fmt.Errorf(\"parsing spec.%s external=%%q: %%w\", parentRef.GetExternal(), err)\n", pr.SpecJSONName)
	w.WriteString("\t}\n")
	if plan.ResourceIDField != nil {
		fmt.Fprintf(w, "\tparsed.%s = resourceID\n", plan.ResourceIDField.StructField)
	}
	w.WriteString("\treturn parsed, nil\n")
	w.WriteString("}\n\n")
}

func renderTailSuffixExpr(tail []IdentitySegment, hasID, serverGen bool) string {
	if len(tail) == 0 {
		return `""`
	}
	var parts []string
	for i, s := range tail {
		if !s.IsVariable {
			parts = append(parts, fmt.Sprintf("%q", "/"+s.Literal))
			continue
		}
		if i == len(tail)-1 && hasID && !serverGen {
			parts = append(parts, `"/" + resourceID`)
		} else {
			parts = append(parts, `"/_"`)
		}
	}
	return strings.Join(parts, " + ")
}

func renderSingleAncestorLookup(w *bytes.Buffer, plan *IdentityPlan, anc AncestorResolution) string {
	varName := safeAncestorVarName(anc.Field.RawPlaceholder)
	switch anc.Source {
	case AncestorProjectRef:
		varName = "projectID"
		w.WriteString("\tprojectID, err := refs.ResolveProjectID(ctx, reader, obj)\n")
		w.WriteString("\tif err != nil {\n")
		w.WriteString("\t\treturn nil, fmt.Errorf(\"cannot resolve project: %w\", err)\n")
		w.WriteString("\t}\n\n")
	case AncestorOrganizationRef:
		varName = "organizationID"
		renderOrgAncestorLookup(w, anc)
	case AncestorFolderRef:
		varName = "folderID"
		renderFolderAncestorLookup(w, anc)
	case AncestorLocation:
		varName = "location"
		renderLocationAncestorLookup(w, anc)
	case AncestorSpecString:
		if anc.IsPointer {
			fmt.Fprintf(w, "\t%s := common.ValueOf(obj.Spec.%s)\n\n", varName, anc.SpecFieldName)
		} else {
			fmt.Fprintf(w, "\t%s := obj.Spec.%s\n\n", varName, anc.SpecFieldName)
		}
	case AncestorParentRef:
		renderParentRefSegmentLookup(w, anc, varName)
	}
	return varName
}

func renderOrgAncestorLookup(w *bytes.Buffer, anc AncestorResolution) {
	w.WriteString("\tvar organizationID string\n")
	if anc.SpecFieldName == "" {
		w.WriteString("\torg, err := refs.ResolveOrganizationFromAnnotation(ctx, reader, obj)\n")
		w.WriteString("\tif err != nil {\n")
		w.WriteString("\t\treturn nil, fmt.Errorf(\"cannot resolve organization: %w\", err)\n")
		w.WriteString("\t}\n")
		w.WriteString("\torganizationID = org.OrganizationID\n\n")
		return
	}
	if anc.IsPointer {
		fmt.Fprintf(w, "\tif obj.Spec.%s != nil {\n", anc.SpecFieldName)
		fmt.Fprintf(w, "\t\torg, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.%s)\n", anc.SpecFieldName)
	} else {
		fmt.Fprintf(w, "\tif obj.Spec.%s.External != \"\" {\n", anc.SpecFieldName)
		fmt.Fprintf(w, "\t\torg, err := refs.ResolveOrganization(ctx, reader, obj, &obj.Spec.%s)\n", anc.SpecFieldName)
	}
	w.WriteString("\t\tif err != nil {\n")
	w.WriteString("\t\t\treturn nil, fmt.Errorf(\"cannot resolve organization: %w\", err)\n")
	w.WriteString("\t\t}\n")
	w.WriteString("\t\torganizationID = org.OrganizationID\n")
	w.WriteString("\t} else {\n")
	w.WriteString("\t\torg, err := refs.ResolveOrganizationFromAnnotation(ctx, reader, obj)\n")
	w.WriteString("\t\tif err != nil {\n")
	w.WriteString("\t\t\treturn nil, fmt.Errorf(\"cannot resolve organization: %w\", err)\n")
	w.WriteString("\t\t}\n")
	w.WriteString("\t\torganizationID = org.OrganizationID\n")
	w.WriteString("\t}\n\n")
}

func renderFolderAncestorLookup(w *bytes.Buffer, anc AncestorResolution) {
	w.WriteString("\tvar folderID string\n")
	if anc.SpecFieldName == "" {
		w.WriteString("\tfolder, err := refs.ResolveFolderFromAnnotation(ctx, reader, obj)\n")
		w.WriteString("\tif err != nil {\n")
		w.WriteString("\t\treturn nil, fmt.Errorf(\"cannot resolve folder: %w\", err)\n")
		w.WriteString("\t}\n")
		w.WriteString("\tfolderID = folder.FolderID\n\n")
		return
	}
	if anc.IsPointer {
		fmt.Fprintf(w, "\tif obj.Spec.%s != nil {\n", anc.SpecFieldName)
		fmt.Fprintf(w, "\t\tfolder, err := refs.ResolveFolder(ctx, reader, obj, obj.Spec.%s)\n", anc.SpecFieldName)
	} else {
		fmt.Fprintf(w, "\tif obj.Spec.%s.External != \"\" || obj.Spec.%s.Name != \"\" {\n", anc.SpecFieldName, anc.SpecFieldName)
		fmt.Fprintf(w, "\t\tfolder, err := refs.ResolveFolder(ctx, reader, obj, &obj.Spec.%s)\n", anc.SpecFieldName)
	}
	w.WriteString("\t\tif err != nil {\n")
	w.WriteString("\t\t\treturn nil, fmt.Errorf(\"cannot resolve folder: %w\", err)\n")
	w.WriteString("\t\t}\n")
	w.WriteString("\t\tfolderID = folder.FolderID\n")
	w.WriteString("\t} else {\n")
	w.WriteString("\t\tfolder, err := refs.ResolveFolderFromAnnotation(ctx, reader, obj)\n")
	w.WriteString("\t\tif err != nil {\n")
	w.WriteString("\t\t\treturn nil, fmt.Errorf(\"cannot resolve folder: %w\", err)\n")
	w.WriteString("\t\t}\n")
	w.WriteString("\t\tfolderID = folder.FolderID\n")
	w.WriteString("\t}\n\n")
}

func renderLocationAncestorLookup(w *bytes.Buffer, anc AncestorResolution) {
	if anc.SpecFieldName == "" || anc.SpecJSONName == "location" {
		w.WriteString("\tlocation, err := refs.GetLocation(obj)\n")
		w.WriteString("\tif err != nil {\n")
		w.WriteString("\t\treturn nil, fmt.Errorf(\"cannot resolve location: %w\", err)\n")
		w.WriteString("\t}\n\n")
		return
	}
	if anc.IsPointer {
		fmt.Fprintf(w, "\tlocation := common.ValueOf(obj.Spec.%s)\n", anc.SpecFieldName)
	} else {
		fmt.Fprintf(w, "\tlocation := obj.Spec.%s\n", anc.SpecFieldName)
	}
	w.WriteString("\tif location == \"\" {\n")
	fmt.Fprintf(w, "\t\treturn nil, fmt.Errorf(\"spec.%s is required\")\n", anc.SpecJSONName)
	w.WriteString("\t}\n\n")
}

func renderParentRefSegmentLookup(w *bytes.Buffer, anc AncestorResolution, varName string) {
	refVar := varName + "Ref"
	if anc.IsPointer {
		fmt.Fprintf(w, "\tif obj.Spec.%s == nil {\n", anc.SpecFieldName)
		fmt.Fprintf(w, "\t\treturn nil, fmt.Errorf(\"spec.%s is required\")\n", anc.SpecJSONName)
		w.WriteString("\t}\n")
		fmt.Fprintf(w, "\t%s := *obj.Spec.%s\n", refVar, anc.SpecFieldName)
	} else {
		fmt.Fprintf(w, "\t%s := obj.Spec.%s\n", refVar, anc.SpecFieldName)
	}
	fmt.Fprintf(w, "\tif err := %s.Normalize(ctx, reader, obj.GetNamespace()); err != nil {\n", refVar)
	w.WriteString("\t\treturn nil, err\n")
	w.WriteString("\t}\n")
	fmt.Fprintf(w, "\t%sParts := strings.Split(strings.Trim(%s.GetExternal(), \"/\"), \"/\")\n", varName, refVar)
	fmt.Fprintf(w, "\t%s := %sParts[len(%sParts)-1]\n\n", varName, varName, varName)
}

func safeAncestorVarName(rawPlaceholder string) string {
	varName := snakeToLowerCamel(rawPlaceholder)
	switch varName {
	case "type", "func", "range", "map", "default", "select", "case", "interface", "struct", "package", "import":
		return varName + "ID"
	default:
		return varName
	}
}

func renderGetIdentityMethod(w *bytes.Buffer, plan *IdentityPlan, fromSpecFn, idType string) {
	fmt.Fprintf(w, "func (obj *%s) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {\n", plan.Kind)
	fmt.Fprintf(w, "\tspecIdentity, err := %s(ctx, reader, obj)\n", fromSpecFn)
	w.WriteString("\tif err != nil {\n")
	w.WriteString("\t\treturn nil, err\n")
	w.WriteString("\t}\n\n")

	if plan.StatusField != "" {
		renderStatusIdentityCrossCheck(w, plan, idType)
	}
	w.WriteString("\treturn specIdentity, nil\n")
	w.WriteString("}\n")
}

func renderStatusIdentityCrossCheck(w *bytes.Buffer, plan *IdentityPlan, idType string) {
	w.WriteString("\t// Cross-check the identity against the status value, if present.\n")
	if plan.StatusFieldIsPtr {
		fmt.Fprintf(w, "\texternalRef := common.ValueOf(obj.Status.%s)\n", plan.StatusField)
	} else {
		fmt.Fprintf(w, "\texternalRef := obj.Status.%s\n", plan.StatusField)
	}
	w.WriteString("\tif externalRef != \"\" {\n")
	fmt.Fprintf(w, "\t\tstatusIdentity := &%s{}\n", idType)
	w.WriteString("\t\tif err := statusIdentity.FromExternal(externalRef); err != nil {\n")
	w.WriteString("\t\t\treturn nil, err\n")
	w.WriteString("\t\t}\n\n")
	if plan.ServerGeneratedID && plan.ResourceIDField != nil {
		fmt.Fprintf(w, "\t\tif specIdentity.%s == \"\" {\n", plan.ResourceIDField.StructField)
		fmt.Fprintf(w, "\t\t\tspecIdentity.%s = statusIdentity.%s\n", plan.ResourceIDField.StructField, plan.ResourceIDField.StructField)
		w.WriteString("\t\t}\n\n")
	}
	w.WriteString("\t\tif statusIdentity.String() != specIdentity.String() {\n")
	fmt.Fprintf(w, "\t\t\treturn nil, fmt.Errorf(\"cannot change %s identity (old=%%q, new=%%q)\", statusIdentity.String(), specIdentity.String())\n", plan.Kind)
	w.WriteString("\t\t}\n")
	w.WriteString("\t}\n\n")
}

func shouldEmitReferenceFile(ov *HandWrittenOverrides) bool {
	if ov.HasRefStruct && ov.RefMethods["Normalize"] {
		return false
	}
	return !ov.HasRefStruct || !ov.RefMethods["GetGVK"] || !ov.RefMethods["GetNamespacedName"] ||
		!ov.RefMethods["GetExternal"] || !ov.RefMethods["SetExternal"] ||
		!ov.RefMethods["ValidateExternal"] || !ov.RefMethods["ParseExternalToIdentity"] || !ov.RefMethods["Normalize"]
}

func renderReferenceStructAndGVK(w *bytes.Buffer, plan *IdentityPlan, ov *HandWrittenOverrides) {
	kind := plan.Kind
	refType := kind + "Ref"
	gvkVar := kind + "GVK"
	if !ov.HasGVKVar {
		fmt.Fprintf(w, "var %s = schema.GroupVersionKind{\n", gvkVar)
		fmt.Fprintf(w, "\tGroup:   %q,\n", plan.Group)
		fmt.Fprintf(w, "\tVersion: %q,\n", plan.Version)
		fmt.Fprintf(w, "\tKind:    %q,\n", kind)
		w.WriteString("}\n\n")
	}
	if !ov.HasRefStruct {
		fmt.Fprintf(w, "var _ refs.Ref = &%s{}\n\n", refType)
		fmt.Fprintf(w, "// %s is a reference to a GCP %s resource.\n", refType, kind)
		fmt.Fprintf(w, "type %s struct {\n", refType)
		fmt.Fprintf(w, "\t// A reference to an externally managed %s resource.\n", kind)
		fmt.Fprintf(w, "\t// Should be in the format %q.\n", plan.DocFormat)
		w.WriteString("\tExternal string `json:\"external,omitempty\"`\n\n")
		fmt.Fprintf(w, "\t// The name of a %s resource.\n", kind)
		w.WriteString("\tName string `json:\"name,omitempty\"`\n\n")
		fmt.Fprintf(w, "\t// The namespace of a %s resource.\n", kind)
		w.WriteString("\tNamespace string `json:\"namespace,omitempty\"`\n")
		w.WriteString("}\n\n")
	}
	if !ov.HasRefRegister {
		w.WriteString("func init() {\n")
		fmt.Fprintf(w, "\trefs.Register(&%s{}, &%s{})\n", refType, kind)
		w.WriteString("}\n\n")
	}
}

func renderReferenceMethods(w *bytes.Buffer, plan *IdentityPlan, ov *HandWrittenOverrides) {
	refType := plan.Kind + "Ref"
	idType := plan.Kind + "Identity"
	gvkVar := plan.Kind + "GVK"

	if !ov.RefMethods["GetGVK"] {
		fmt.Fprintf(w, "func (r *%s) GetGVK() schema.GroupVersionKind {\n", refType)
		fmt.Fprintf(w, "\treturn %s\n", gvkVar)
		w.WriteString("}\n\n")
	}
	if !ov.RefMethods["GetNamespacedName"] {
		fmt.Fprintf(w, "func (r *%s) GetNamespacedName() types.NamespacedName {\n", refType)
		w.WriteString("\treturn types.NamespacedName{\n\t\tName:      r.Name,\n\t\tNamespace: r.Namespace,\n\t}\n}\n\n")
	}
	if !ov.RefMethods["GetExternal"] {
		fmt.Fprintf(w, "func (r *%s) GetExternal() string {\n\treturn r.External\n}\n\n", refType)
	}
	if !ov.RefMethods["SetExternal"] {
		fmt.Fprintf(w, "func (r *%s) SetExternal(ref string) {\n\tr.External = ref\n\tr.Name = \"\"\n\tr.Namespace = \"\"\n}\n\n", refType)
	}
	if !ov.RefMethods["ValidateExternal"] {
		fmt.Fprintf(w, "func (r *%s) ValidateExternal(ref string) error {\n\tid := &%s{}\n\tif err := id.FromExternal(ref); err != nil {\n\t\treturn err\n\t}\n\treturn nil\n}\n\n", refType, idType)
	}
	if !ov.RefMethods["ParseExternalToIdentity"] {
		fmt.Fprintf(w, "func (r *%s) ParseExternalToIdentity() (identity.Identity, error) {\n\tid := &%s{}\n\tif err := id.FromExternal(r.External); err != nil {\n\t\treturn nil, err\n\t}\n\treturn id, nil\n}\n\n", refType, idType)
	}
	if !ov.RefMethods["Normalize"] {
		fmt.Fprintf(w, "func (r *%s) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {\n\treturn refs.Normalize(ctx, reader, r, defaultNamespace)\n}\n", refType)
	}
}

func sampleSegmentsString(segs []IdentitySegment) string {
	var parts []string
	for _, s := range segs {
		if s.IsVariable {
			parts = append(parts, s.Field.SampleValue)
		} else {
			parts = append(parts, s.Literal)
		}
	}
	return strings.Join(parts, "/")
}

type candidateImport struct {
	alias string
	path  string
	token string
}

var allCandidateImports = []candidateImport{
	{path: "context", token: "context."},
	{path: "fmt", token: "fmt."},
	{path: "strings", token: "strings."},
	{path: "testing", token: "testing."},
	{path: "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common", token: "common."},
	{path: "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity", token: "identity."},
	{alias: "refs", path: "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1", token: "refs."},
	{path: "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls", token: "gcpurls."},
	{path: "k8s.io/apimachinery/pkg/runtime/schema", token: "schema."},
	{path: "k8s.io/apimachinery/pkg/types", token: "types."},
	{path: "sigs.k8s.io/controller-runtime/pkg/client", token: "client."},
}

func formatGeneratedGoFile(pkgName, annotationKey string, plan *IdentityPlan, body []byte) ([]byte, error) {
	stdImports, thirdPartyImports := collectReferencedImports(string(body))

	var out bytes.Buffer
	out.WriteString(`// Copyright 2026 Google LLC
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

// Code generated by dev/tasks/generate-all. DO NOT EDIT.

`)
	fmt.Fprintf(&out, "// %s\n", annotationKey)
	fmt.Fprintf(&out, "// krm.group: %s\n", plan.Group)
	fmt.Fprintf(&out, "// krm.version: %s\n", plan.Version)
	fmt.Fprintf(&out, "// proto.message: %s\n\n", plan.ProtoMessage)
	fmt.Fprintf(&out, "package %s\n\n", pkgName)

	writeImportsBlock(&out, stdImports, thirdPartyImports)
	out.Write(body)

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		return nil, fmt.Errorf("gofmt generated file: %w\n%s", err, out.String())
	}
	return formatted, nil
}

func collectReferencedImports(bodyStr string) (stdImports, thirdPartyImports []candidateImport) {
	for _, imp := range allCandidateImports {
		if !strings.Contains(bodyStr, imp.token) {
			continue
		}
		if strings.Contains(imp.path, ".") {
			thirdPartyImports = append(thirdPartyImports, imp)
		} else {
			stdImports = append(stdImports, imp)
		}
	}
	return stdImports, thirdPartyImports
}

func writeImportsBlock(w *bytes.Buffer, stdImports, thirdPartyImports []candidateImport) {
	if len(stdImports) == 0 && len(thirdPartyImports) == 0 {
		return
	}
	w.WriteString("import (\n")
	writeImportGroup(w, stdImports)
	if len(stdImports) > 0 && len(thirdPartyImports) > 0 {
		w.WriteString("\n")
	}
	writeImportGroup(w, thirdPartyImports)
	w.WriteString(")\n\n")
}

func writeImportGroup(w *bytes.Buffer, imports []candidateImport) {
	for _, imp := range imports {
		if imp.alias != "" {
			fmt.Fprintf(w, "\t%s %q\n", imp.alias, imp.path)
		} else {
			fmt.Fprintf(w, "\t%q\n", imp.path)
		}
	}
}
