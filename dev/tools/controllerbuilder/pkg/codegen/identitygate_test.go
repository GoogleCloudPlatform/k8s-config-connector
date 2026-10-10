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
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
)

// TestIdentityGeneratorOfflineGate validates that IdentityGenerator.Plan reproduces
// the host and URL template pattern of >= 90% of existing canonical gcpurls.Template
// _identity.go files across apis/.
func TestIdentityGeneratorOfflineGate(t *testing.T) {
	repoRoot, err := options.RepoRoot()
	if err != nil {
		t.Skipf("cannot locate repo root: %v", err)
	}
	pbPath := filepath.Join(repoRoot, ".build", "googleapis.pb")
	if _, err := os.Stat(pbPath); err != nil {
		t.Skipf("skipping offline identity gate (%s not present)", pbPath)
	}

	api, err := protoapi.LoadProto(pbPath, "")
	if err != nil {
		t.Fatalf("LoadProto: %v", err)
	}

	apisDir := filepath.Join(repoRoot, "apis")
	caiPath := filepath.Join(repoRoot, "docs", "ai", "metadata", "cloudassetinventory_names.jsonl")

	total := 0
	matched := 0
	var mismatches []string
	for _, tc := range loadGateCases(t, repoRoot, apisDir) {
		plan, ok := planGateCase(t, api, apisDir, caiPath, tc)
		if !ok {
			continue
		}
		total++
		if matchesAnyTemplate(plan, tc.templates) {
			matched++
			continue
		}
		mismatches = append(mismatches, fmt.Sprintf(
			"%s (%s): existing=(%s, %s) vs planned=(%s, %s)",
			tc.kind, tc.relPath, tc.templates[0].host, tc.templates[0].pattern, plan.Host, plan.TemplatePattern,
		))
	}

	if total == 0 {
		t.Fatal("found 0 canonical gcpurls.Template identity files to evaluate")
	}

	rate := float64(matched) * 100.0 / float64(total)
	t.Logf("IdentityGenerator offline gate: %d/%d (%.1f%%) canonical gcpurls.Template resources matched", matched, total, rate)
	for _, mm := range mismatches {
		t.Logf("  mismatch: %s", mm)
	}

	const minGatePct = 90.0
	if rate < minGatePct {
		t.Fatalf("IdentityGenerator offline gate %.1f%% (%d/%d) is below required %.1f%% threshold;\nmismatches:\n%s",
			rate, matched, total, minGatePct, strings.Join(mismatches, "\n"))
	}
}

type gateTemplate struct {
	host    string
	pattern string
}

type gateCase struct {
	kind      string
	group     string
	version   string
	protoName string
	relPath   string
	templates []gateTemplate
}

func planGateCase(t *testing.T, api *protoapi.Proto, apisDir, caiPath string, tc gateCase) (*IdentityPlan, bool) {
	t.Helper()
	service := tc.protoName
	if idx := strings.LastIndexByte(tc.protoName, '.'); idx > 0 {
		service = tc.protoName[:idx]
	}
	gen := NewIdentityGenerator(IdentityGeneratorOptions{
		Group:           tc.group,
		Version:         tc.version,
		ProtoService:    service,
		APIDirectory:    apisDir,
		Proto:           api,
		CAIMetadataPath: caiPath,
	})
	msg, err := gen.ResolveMessage(tc.protoName)
	if err != nil {
		return nil, false
	}
	plan, err := gen.Plan(tc.kind, msg)
	if err != nil {
		t.Fatalf("Plan(%s): %v", tc.kind, err)
	}
	return plan, true
}

func matchesAnyTemplate(plan *IdentityPlan, want []gateTemplate) bool {
	if len(plan.MultiParentVariants) > 1 {
		for _, v := range plan.MultiParentVariants {
			vNorm := tmplVarNormalizeRe.ReplaceAllString(v.TemplatePattern, "{}")
			found := false
			for _, w := range want {
				if plan.Host == w.host && tmplVarNormalizeRe.ReplaceAllString(w.pattern, "{}") == vNorm {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	}
	plannedNorm := tmplVarNormalizeRe.ReplaceAllString(plan.TemplatePattern, "{}")
	for _, w := range want {
		if plan.Host == w.host && plannedNorm == tmplVarNormalizeRe.ReplaceAllString(w.pattern, "{}") {
			return true
		}
	}
	return false
}

var (
	gateTemplateRe  = regexp.MustCompile(`gcpurls\.Template\[[^\]]+\]\(\s*"([^"]*)"\s*,\s*"([^"]+)"\s*\)`)
	gateIdentityRe  = regexp.MustCompile(`func\s+\(\w+\s+\*([A-Za-z0-9_]+)\)\s+GetIdentity\(`)
	gateSpecProtoRe = regexp.MustCompile(`\+kcc:(?:spec:)?proto=([A-Za-z0-9_.]+)\s*\n\s*type\s+([A-Za-z0-9_]+)Spec\s+struct`)
)

func loadGateCases(t *testing.T, repoRoot, apisDir string) []gateCase {
	t.Helper()
	files, err := filepath.Glob(filepath.Join(apisDir, "*", "*", "*_identity.go"))
	if err != nil {
		t.Fatalf("Glob _identity.go: %v", err)
	}
	sort.Strings(files)

	var cases []gateCase
	for _, idFile := range files {
		rel, _ := filepath.Rel(repoRoot, idFile)
		parts := strings.Split(filepath.ToSlash(rel), "/")
		if len(parts) < 4 || parts[1] == "compute" {
			continue
		}
		src, err := os.ReadFile(idFile)
		if err != nil {
			continue
		}
		m := gateIdentityRe.FindSubmatch(src)
		if m == nil {
			continue
		}
		kind := string(m[1])
		var tmpls []gateTemplate
		for _, sub := range gateTemplateRe.FindAllStringSubmatch(string(src), -1) {
			tmpls = append(tmpls, gateTemplate{host: sub[1], pattern: sub[2]})
		}
		if len(tmpls) == 0 {
			continue
		}
		protoName := findSpecProto(filepath.Dir(idFile), kind)
		if protoName == "" {
			continue
		}
		cases = append(cases, gateCase{
			kind:      kind,
			group:     parts[1] + ".cnrm.cloud.google.com",
			version:   parts[2],
			protoName: protoName,
			relPath:   rel,
			templates: tmpls,
		})
	}
	return cases
}

func findSpecProto(pkgDir, kind string) string {
	files, _ := filepath.Glob(filepath.Join(pkgDir, "*.go"))
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		for _, m := range gateSpecProtoRe.FindAllStringSubmatch(string(b), -1) {
			if m[2] == kind {
				return m[1]
			}
		}
	}
	return ""
}
