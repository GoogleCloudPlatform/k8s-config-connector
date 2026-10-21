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

package shadow

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/controllerprofile"
)

// ControllerAgreement represents the agreement analysis for a single existing controller.
type ControllerAgreement struct {
	FilePath       string  `json:"file_path"`
	Kind           string  `json:"kind"`
	Service        string  `json:"service"`
	StructuralScore float64 `json:"structural_score"` // 0.0 to 1.0 (8 standard components)

	// Structural presence in existing controller
	HasRegisterModel     bool `json:"has_register_model"`
	HasModelConstructor  bool `json:"has_model_constructor"`
	HasClientConstructor bool `json:"has_client_constructor"`
	HasAdapterForObject  bool `json:"has_adapter_for_object"`
	HasAdapterForURL     bool `json:"has_adapter_for_url"`
	HasFind              bool `json:"has_find"`
	HasCreate            bool `json:"has_create"`
	HasUpdate            bool `json:"has_update"`
	HasExport            bool `json:"has_export"`
	HasDelete            bool `json:"has_delete"`

	// Behavioral & Pattern Agreement
	HasNormalizeReferences bool `json:"has_normalize_references"`
	HasGetIdentity          bool `json:"has_get_identity"`
	HasIsNotFoundHandling   bool `json:"has_is_not_found_handling"`
	HasStatusUpdate         bool `json:"has_status_update"`
	HasCompareGreenfield    bool `json:"has_compare_greenfield"`
	HasLROWait              bool `json:"has_lro_wait"`

	// Custom / Hand-crafted extensions detected
	CustomFunctions []string `json:"custom_functions,omitempty"`
	GeneratedLines  int      `json:"generated_lines"`
	ExistingLines   int      `json:"existing_lines"`

	ShadowProfile *controllerprofile.ControllerProfile `json:"shadow_profile,omitempty"`
}

// BatchAgreementSummary aggregates agreement results across a collection of direct controllers.
type BatchAgreementSummary struct {
	TotalEvaluated       int                     `json:"total_evaluated"`
	AverageStructuralScore float64               `json:"avg_structural_score"`
	RegisterModelRate    float64                 `json:"register_model_rate"`
	FindAgreementRate    float64                 `json:"find_agreement_rate"`
	CreateAgreementRate  float64                 `json:"create_agreement_rate"`
	UpdateAgreementRate  float64                 `json:"update_agreement_rate"`
	ExportAgreementRate  float64                 `json:"export_agreement_rate"`
	DeleteAgreementRate  float64                 `json:"delete_agreement_rate"`
	Controllers          []ControllerAgreement   `json:"controllers"`
}

// EvaluateControllerFile inspects an existing direct controller file in shadow mode.
func EvaluateControllerFile(path string) (*ControllerAgreement, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file %q: %w", path, err)
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, content, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parse Go file %q: %w", path, err)
	}

	rawText := string(content)
	lines := strings.Split(rawText, "\n")

	relParts := strings.Split(filepath.ToSlash(path), "/")
	service := ""
	for i, part := range relParts {
		if part == "direct" && i+1 < len(relParts) {
			service = relParts[i+1]
			break
		}
	}

	agreement := &ControllerAgreement{
		FilePath:      path,
		Service:       service,
		ExistingLines: len(lines),
	}

	// Scan AST for methods and structure
	var standardMethods = map[string]bool{
		"AdapterForObject": false,
		"AdapterForURL":    false,
		"Find":             false,
		"Create":           false,
		"Update":           false,
		"Export":           false,
		"Delete":           false,
		"client":           false,
	}

	for _, decl := range node.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		name := fn.Name.Name
		if fn.Recv != nil && len(fn.Recv.List) > 0 {
			// Method on Model or Adapter
			if _, exists := standardMethods[name]; exists {
				standardMethods[name] = true
			} else {
				agreement.CustomFunctions = append(agreement.CustomFunctions, name)
			}
		} else {
			if name == "init" {
				agreement.HasRegisterModel = strings.Contains(rawText, "registry.RegisterModel")
			} else if strings.HasPrefix(name, "New") && strings.HasSuffix(name, "Model") {
				agreement.HasModelConstructor = true
			} else {
				agreement.CustomFunctions = append(agreement.CustomFunctions, name)
			}
		}
	}

	agreement.HasAdapterForObject = standardMethods["AdapterForObject"]
	agreement.HasAdapterForURL = standardMethods["AdapterForURL"]
	agreement.HasClientConstructor = standardMethods["client"]
	agreement.HasFind = standardMethods["Find"]
	agreement.HasCreate = standardMethods["Create"]
	agreement.HasUpdate = standardMethods["Update"]
	agreement.HasExport = standardMethods["Export"]
	agreement.HasDelete = standardMethods["Delete"]

	// Check Behavioral traits in source
	agreement.HasNormalizeReferences = strings.Contains(rawText, "NormalizeReferences")
	agreement.HasGetIdentity = strings.Contains(rawText, "GetIdentity")
	agreement.HasIsNotFoundHandling = strings.Contains(rawText, "direct.IsNotFound")
	agreement.HasStatusUpdate = strings.Contains(rawText, "UpdateStatus")
	agreement.HasCompareGreenfield = strings.Contains(rawText, "CompareGreenfieldSpec") || strings.Contains(rawText, "CompareGreenfieldSpecAndLabels")
	agreement.HasLROWait = strings.Contains(rawText, ".Wait(ctx)")

	// Determine Kind
	base := filepath.Base(path)
	kindCandidate := strings.TrimSuffix(base, "_controller.go")
	agreement.Kind = kindCandidate

	// Calculate structural score across 8 primary components:
	// 1: RegisterModel, 2: ModelConstructor, 3: ClientConstructor, 4: AdapterForObject,
	// 5: Find, 6: Create, 7: Update, 8: Delete
	scoreComponents := []bool{
		agreement.HasRegisterModel,
		agreement.HasModelConstructor,
		agreement.HasClientConstructor,
		agreement.HasAdapterForObject,
		agreement.HasFind,
		agreement.HasCreate,
		agreement.HasUpdate,
		agreement.HasDelete,
	}
	matched := 0
	for _, c := range scoreComponents {
		if c {
			matched++
		}
	}
	agreement.StructuralScore = float64(matched) / float64(len(scoreComponents))

	// Synthesize a Shadow ControllerProfile
	protoResource := strings.TrimPrefix(kindCandidate, strings.ToLower(service))
	if protoResource == "" {
		protoResource = kindCandidate
	}
	profile := &controllerprofile.ControllerProfile{
		KCCService:    service,
		KCCVersion:    "v1beta1",
		Kind:          kindCandidate,
		ProtoResource: protoResource,
		Lifecycle: controllerprofile.LifecycleTrait{
			HasGet:            agreement.HasFind,
			CanCreate:         agreement.HasCreate,
			UpdateStyle:       controllerprofile.UpdateStyleMask,
			DeleteStyle:       controllerprofile.DeleteStyleStandard,
			GetMethodName:     "Get" + protoResource,
			GetRequestType:    "Get" + protoResource + "Request",
			GetNameField:      "Name",
			CreateMethodName:  "Create" + protoResource,
			CreateRequestType: "Create" + protoResource + "Request",
			CreateParentField: "Parent",
			CreateBodyField:   protoResource,
			CreateIDField:     protoResource + "Id",
			UpdateMethodName:  "Update" + protoResource,
			UpdateRequestType: "Update" + protoResource + "Request",
			UpdateBodyField:   protoResource,
			UpdateMaskField:   "UpdateMask",
			DeleteMethodName:  "Delete" + protoResource,
			DeleteRequestType: "Delete" + protoResource + "Request",
			DeleteNameField:   "Name",
		},
		Identity: controllerprofile.IdentityTrait{
			Strategy:         controllerprofile.IdentityClientField,
			TypeName:         "krm." + kindCandidate + "Identity",
			UseGetIdentity:   agreement.HasGetIdentity,
			ParentStringExpr: "ParentString()",
			IDExpr:           "ID()",
			NameField:        "Name",
			HasFromExternal:  agreement.HasAdapterForURL,
		},
		Execution: controllerprofile.ExecutionTrait{
			Create: controllerprofile.OpExecution{IsLRO: agreement.HasLROWait},
			Update: controllerprofile.OpExecution{IsLRO: agreement.HasLROWait},
			Delete: controllerprofile.OpExecution{IsLRO: agreement.HasLROWait, IsEmpty: true},
		},
		Transport: controllerprofile.TransportTrait{
			GCPClientImportPath:  fmt.Sprintf("cloud.google.com/go/%s/apiv1", service),
			GCPClientType:        "*gcp.Client",
			GCPClientConstructor: "gcp.NewRESTClient",
			PBImportPath:         fmt.Sprintf("cloud.google.com/go/%s/apiv1/%spb", service, service),
			PBPackageAlias:       service + "pb",
			HasReferences:        agreement.HasNormalizeReferences,
			LabelDiffing:         controllerprofile.LabelsNone,
		},
	}
	agreement.ShadowProfile = profile

	// Render shadow controller to count generated lines
	if generatedBytes, err := controllerprofile.RenderController(profile); err == nil {
		agreement.GeneratedLines = len(strings.Split(string(generatedBytes), "\n"))
	}

	return agreement, nil
}

// BatchEvaluate evaluates all matching controllers in the given directory.
func BatchEvaluate(files []string) *BatchAgreementSummary {
	summary := &BatchAgreementSummary{
		TotalEvaluated: len(files),
	}

	var totalScore float64
	var regCount, findCount, createCount, updateCount, exportCount, deleteCount int

	for _, file := range files {
		agr, err := EvaluateControllerFile(file)
		if err != nil {
			continue
		}
		summary.Controllers = append(summary.Controllers, *agr)
		totalScore += agr.StructuralScore

		if agr.HasRegisterModel {
			regCount++
		}
		if agr.HasFind {
			findCount++
		}
		if agr.HasCreate {
			createCount++
		}
		if agr.HasUpdate {
			updateCount++
		}
		if agr.HasExport {
			exportCount++
		}
		if agr.HasDelete {
			deleteCount++
		}
	}

	if len(summary.Controllers) > 0 {
		n := float64(len(summary.Controllers))
		summary.AverageStructuralScore = totalScore / n
		summary.RegisterModelRate = float64(regCount) / n
		summary.FindAgreementRate = float64(findCount) / n
		summary.CreateAgreementRate = float64(createCount) / n
		summary.UpdateAgreementRate = float64(updateCount) / n
		summary.ExportAgreementRate = float64(exportCount) / n
		summary.DeleteAgreementRate = float64(deleteCount) / n
	}

	return summary
}
