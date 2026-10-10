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
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ResolveMessage finds a proto MessageDescriptor by short name (within g.opts.ProtoService)
// or full name in the loaded Proto descriptor set.
func (g *IdentityGenerator) ResolveMessage(protoName string) (protoreflect.MessageDescriptor, error) {
	if g.opts.Proto == nil {
		return nil, fmt.Errorf("Proto descriptor set is nil")
	}
	if strings.HasPrefix(protoName, "mockgcp.") || strings.HasPrefix(g.opts.ProtoService, "mockgcp.") {
		return nil, fmt.Errorf("generate-identity does not support mockgcp proto %q", protoName)
	}
	if strings.Contains(protoName, ".") {
		desc, err := g.opts.Proto.Files().FindDescriptorByName(protoreflect.FullName(protoName))
		if err != nil {
			return nil, fmt.Errorf("finding proto message %q: %w", protoName, err)
		}
		msg, ok := desc.(protoreflect.MessageDescriptor)
		if !ok {
			return nil, fmt.Errorf("%q is not a proto message", protoName)
		}
		return msg, nil
	}
	for _, svcPkg := range strings.Split(g.opts.ProtoService, ",") {
		svcPkg = strings.TrimSpace(svcPkg)
		if svcPkg == "" {
			continue
		}
		fullName := protoreflect.FullName(svcPkg + "." + protoName)
		if desc, err := g.opts.Proto.Files().FindDescriptorByName(fullName); err == nil {
			if msg, ok := desc.(protoreflect.MessageDescriptor); ok {
				return msg, nil
			}
		}
	}
	return nil, fmt.Errorf("finding proto message %q", protoName)
}

func (g *IdentityGenerator) inspectProtoService(msg protoreflect.MessageDescriptor) (defaultHost, getHTTPPath string, serverGeneratedID bool) {
	if msg == nil {
		return "", "", false
	}
	wantGet := "Get" + string(msg.Name())
	wantCreate := "Create" + string(msg.Name())
	for _, fd := range g.serviceFilesForMessage(msg) {
		svcs := fd.Services()
		for i := 0; i < svcs.Len(); i++ {
			svc := svcs.Get(i)
			if defaultHost == "" {
				defaultHost = extractServiceDefaultHost(svc)
			}
			getPath, isServerGen := inspectServiceMethods(svc, wantGet, wantCreate)
			if getHTTPPath == "" && getPath != "" {
				getHTTPPath = getPath
			}
			if isServerGen {
				serverGeneratedID = true
			}
		}
	}
	return defaultHost, getHTTPPath, serverGeneratedID
}

func extractServiceDefaultHost(svc protoreflect.ServiceDescriptor) string {
	h, _ := proto.GetExtension(svc.Options(), annotations.E_DefaultHost).(string)
	return h
}

func inspectServiceMethods(svc protoreflect.ServiceDescriptor, wantGet, wantCreate string) (getHTTPPath string, serverGeneratedID bool) {
	meths := svc.Methods()
	for j := 0; j < meths.Len(); j++ {
		m := meths.Get(j)
		if string(m.Name()) == wantGet && getHTTPPath == "" {
			if rule, ok := proto.GetExtension(m.Options(), annotations.E_Http).(*annotations.HttpRule); ok && rule != nil {
				getHTTPPath = rule.GetGet()
			}
		}
		if string(m.Name()) == wantCreate && isServerGeneratedCreateInput(m.Input().Fields()) {
			serverGeneratedID = true
		}
	}
	return getHTTPPath, serverGeneratedID
}

func isServerGeneratedCreateInput(inFields protoreflect.FieldDescriptors) bool {
	if inFields.ByName("parent") == nil {
		return false
	}
	for k := 0; k < inFields.Len(); k++ {
		f := inFields.Get(k)
		if strings.HasSuffix(string(f.Name()), "_id") && f.Kind() == protoreflect.StringKind {
			return false
		}
	}
	return true
}

var httpBindingVarRe = regexp.MustCompile(`\{[a-zA-Z0-9_.]+=([^}]+)\}`)

func (g *IdentityGenerator) rpcPatternsForMessage(msg protoreflect.MessageDescriptor) []string {
	if msg == nil || strings.HasPrefix(string(msg.FullName()), "mockgcp.") {
		return nil
	}
	var out []string
	for _, fd := range g.serviceFilesForMessage(msg) {
		svcs := fd.Services()
		for i := 0; i < svcs.Len(); i++ {
			meths := svcs.Get(i).Methods()
			for j := 0; j < meths.Len(); j++ {
				if pat := extractGetRPCPattern(meths.Get(j), msg.FullName()); pat != "" {
					out = append(out, pat)
				}
			}
		}
	}
	return out
}

func extractGetRPCPattern(m protoreflect.MethodDescriptor, wantOutput protoreflect.FullName) string {
	if !strings.HasPrefix(string(m.Name()), "Get") || m.Output().FullName() != wantOutput {
		return ""
	}
	rule, ok := proto.GetExtension(m.Options(), annotations.E_Http).(*annotations.HttpRule)
	if !ok || rule == nil || rule.GetGet() == "" {
		return ""
	}
	sub := httpBindingVarRe.FindStringSubmatch(rule.GetGet())
	if sub == nil {
		return ""
	}
	return wildcardPathToPattern(sub[1])
}

func (g *IdentityGenerator) serviceFilesForMessage(msg protoreflect.MessageDescriptor) []protoreflect.FileDescriptor {
	var files []protoreflect.FileDescriptor
	if msg.ParentFile() != nil && !strings.HasPrefix(string(msg.ParentFile().Package()), "mockgcp.") {
		files = append(files, msg.ParentFile())
	}
	if g.opts.Proto == nil {
		return files
	}
	wantPkgs := map[string]bool{}
	for _, sp := range strings.Split(g.opts.ProtoService, ",") {
		sp = strings.TrimSpace(sp)
		if sp != "" && !strings.HasPrefix(sp, "mockgcp.") {
			wantPkgs[sp] = true
		}
	}
	groupPrefix := strings.Split(g.opts.Group, ".")[0]
	g.opts.Proto.Files().RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		pkg := string(fd.Package())
		if wantPkgs[pkg] || matchesGroupServicePackage(pkg, groupPrefix) {
			files = append(files, fd)
		}
		return true
	})
	return files
}

func matchesGroupServicePackage(pkg, groupPrefix string) bool {
	if groupPrefix == "" {
		return false
	}
	return strings.HasPrefix(pkg, "google."+groupPrefix+".") || strings.HasPrefix(pkg, "google.cloud."+groupPrefix+".")
}

func wildcardPathToPattern(wc string) string {
	segs := strings.Split(strings.Trim(wc, "/"), "/")
	if len(segs) < 2 {
		return ""
	}
	for i := 0; i < len(segs); i++ {
		if segs[i] != "*" && segs[i] != "**" {
			continue
		}
		if i == 0 || strings.HasPrefix(segs[i-1], "{") || strings.HasPrefix(segs[i-1], "*") {
			return ""
		}
		segs[i] = "{" + snakeCaseIdent(Singular(segs[i-1])) + "}"
	}
	return strings.Join(segs, "/")
}

func snakeCaseIdent(s string) string {
	var b strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func getResourceDescriptorProto(msg protoreflect.MessageDescriptor) *annotations.ResourceDescriptor {
	if msg == nil {
		return nil
	}
	v := proto.GetExtension(msg.Options(), annotations.E_Resource)
	rd, _ := v.(*annotations.ResourceDescriptor)
	return rd
}

var (
	caiVarNormalizeRe = regexp.MustCompile(`\{\{[^}]+\}\}`)
	caiDoubleBraceRe  = regexp.MustCompile(`\{\{([^}]+)\}\}`)
)

func (g *IdentityGenerator) ensureCAILoaded() error {
	if g.caiLoaded {
		return g.caiLoadErr
	}
	g.caiLoaded = true
	g.caiByType = make(map[string][]string)
	g.caiFormats = make(map[string]bool)

	path := g.resolveCAIMetadataPath()
	if path == "" {
		return nil
	}
	f, err := os.Open(path)
	if err != nil {
		g.caiLoadErr = fmt.Errorf("opening CAI metadata %s: %w", path, err)
		return g.caiLoadErr
	}
	defer f.Close()

	type caiEntry struct {
		ResourceType string   `json:"resourceType"`
		NameFormats  []string `json:"nameFormats"`
	}
	sc := bufio.NewScanner(f)
	lineNum := 0
	for sc.Scan() {
		lineNum++
		var e caiEntry
		if err := json.Unmarshal(sc.Bytes(), &e); err != nil {
			g.caiLoadErr = fmt.Errorf("parsing CAI metadata %s line %d: %w", path, lineNum, err)
			return g.caiLoadErr
		}
		g.caiByType[e.ResourceType] = append(g.caiByType[e.ResourceType], e.NameFormats...)
		for _, nf := range e.NameFormats {
			g.caiFormats[caiVarNormalizeRe.ReplaceAllString(nf, "{}")] = true
		}
	}
	if err := sc.Err(); err != nil {
		g.caiLoadErr = fmt.Errorf("scanning CAI metadata %s: %w", path, err)
		return g.caiLoadErr
	}
	return nil
}

func (g *IdentityGenerator) resolveCAIMetadataPath() string {
	path := g.opts.CAIMetadataPath
	if g.opts.APIDirectory == "" {
		return path
	}
	repoRoot := filepath.Dir(strings.TrimRight(g.opts.APIDirectory, "/"))
	candidate := filepath.Join(repoRoot, "docs", "ai", "metadata", "cloudassetinventory_names.jsonl")
	if path == "" {
		if _, err := os.Stat(candidate); err == nil {
			path = candidate
		}
	}
	if path == candidate {
		g.loadRegistryTestAllowlist(repoRoot)
	}
	return path
}

func (g *IdentityGenerator) loadRegistryTestAllowlist(repoRoot string) {
	regBytes, err := os.ReadFile(filepath.Join(repoRoot, "pkg", "gcpurls", "registry_test.go"))
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(regBytes), "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, `"//`) || !strings.HasSuffix(line, `true,`) {
			continue
		}
		if idx := strings.Index(line[1:], `"`); idx > 0 {
			g.caiFormats[line[1:1+idx]] = true
		}
	}
}

func (g *IdentityGenerator) caiPatternsForType(resType string) []string {
	var out []string
	for _, nf := range g.caiByType[resType] {
		if pat := convertCAINameFormatToPattern(nf, true); pat != "" {
			out = append(out, pat)
		}
	}
	return out
}

func (g *IdentityGenerator) fallbackPatterns(msg protoreflect.MessageDescriptor, defaultHost string) []string {
	host := defaultHost
	if host == "" {
		if groupPrefix := strings.TrimSuffix(g.opts.Group, ".cnrm.cloud.google.com"); groupPrefix != "" {
			host = groupPrefix + ".googleapis.com"
		}
	}
	if host == "" {
		return nil
	}
	resType := host + "/" + string(msg.Name())
	var out []string
	for _, nf := range g.caiByType[resType] {
		if pat := convertCAINameFormatToPattern(nf, false); pat != "" {
			out = append(out, pat)
		}
	}
	return out
}

func convertCAINameFormatToPattern(nf string, stripNameSuffix bool) string {
	if !strings.HasPrefix(nf, "//") {
		return ""
	}
	rest := strings.TrimPrefix(nf, "//")
	slash := strings.IndexByte(rest, '/')
	if slash <= 0 {
		return ""
	}
	conv := caiDoubleBraceRe.ReplaceAllStringFunc(rest[slash+1:], func(m string) string {
		inner := strings.ToLower(strings.Trim(m, "{}"))
		inner = strings.TrimSuffix(inner, "_id")
		inner = strings.TrimSuffix(inner, "_number")
		if stripNameSuffix {
			inner = strings.TrimSuffix(inner, "_name")
		}
		return "{" + inner + "}"
	})
	if !strings.Contains(conv, "/") {
		return ""
	}
	return conv
}
