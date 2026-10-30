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

package validator

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/classifier"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

type httpBindingRule struct {
	host       string
	httpMethod string
	segments   []string
	reqMsg     protoreflect.MessageDescriptor
	respMsg    protoreflect.MessageDescriptor
}

type descriptorRegistry struct {
	files     *protoregistry.Files
	types     *dynamicpb.Types
	httpRules map[string][]httpBindingRule
	packages  map[string]bool
}

var (
	registryCacheMu  sync.Mutex
	registryCache    = make(map[string]*descriptorRegistry)
	registryErrCache = make(map[string]error)
	defaultPBOnce    sync.Once
	defaultPBPath    string
	templateVarRegex = regexp.MustCompile(`\{([^{}=]+)(?:=([^{}]+))?\}`)
)

// ValidateProtoSchemas validates all JSON request and response bodies in entries against the
// protobuf descriptor set at descriptorSetPath (defaulting to .build/googleapis.pb when empty).
// If the descriptor set file does not exist on disk, ValidateProtoSchemas returns an empty slice.
func ValidateProtoSchemas(entries []*classifier.ClassifiedEntry, descriptorSetPath string) []string {
	reg, ok, err := getDescriptorRegistry(descriptorSetPath)
	if err != nil || !ok {
		return []string{}
	}

	var errs []string

	// Infer the primary resource MessageDescriptor from any LRO response.@type on primary entries.
	var inferredPrimaryMsg protoreflect.MessageDescriptor
	for _, ce := range entries {
		if ce.Role != classifier.RolePrimary || ce.LRO == nil || ce.LRO.FinalResponse == nil {
			continue
		}
		if typeURL, isStr := ce.LRO.FinalResponse["@type"].(string); isStr && !strings.HasSuffix(typeURL, "google.protobuf.Empty") {
			if msg := reg.findMessageByFullName(typeURL); msg != nil {
				inferredPrimaryMsg = msg
				break
			}
		}
	}

	for _, ce := range entries {
		if ce == nil || ce.Entry == nil {
			continue
		}
		method := ce.Entry.Request.Method
		expectIntEnums := strings.Contains(ce.Entry.Request.URL, "enum-encoding%3Dint") ||
			strings.Contains(ce.Entry.Request.URL, "enum-encoding=int")

		reqMap, _ := ce.Entry.Request.ParseJSONBody()
		respMap, _ := ce.Entry.Response.ParseJSONBody()

		// 1. Validate all "@type" blocks in request and response JSON payloads.
		if reqMap != nil {
			validateTypedBlocksRecursive(reg, ce.Index, method, "request", reqMap, false, &errs)
		}
		if respMap != nil {
			validateTypedBlocksRecursive(reg, ce.Index, method, "response", respMap, expectIntEnums, &errs)
		}

		// 2. Validate primary resource GET / POST / PATCH / PUT JSON bodies against MessageDescriptor.
		if ce.Role == classifier.RolePrimary && !ce.IsLROPoll {
			matchedRule := reg.matchHTTPRule(ce.Entry.Request.URL, method)

			if reqMap != nil && (method == "POST" || method == "PATCH" || method == "PUT") {
				if _, hasType := reqMap["@type"]; !hasType {
					var reqMsg protoreflect.MessageDescriptor
					if matchedRule != nil && matchedRule.reqMsg != nil {
						reqMsg = matchedRule.reqMsg
					} else if inferredPrimaryMsg != nil {
						reqMsg = inferredPrimaryMsg
					}
					if reqMsg != nil {
						if err := unmarshalDynamicJSON(reg, reqMsg, reqMap); err != nil {
							errs = append(errs, fmt.Sprintf("stanza #%d (`%s`) request body failed protobuf schema `%s`: %v", ce.Index, method, reqMsg.FullName(), err))
						}
					}
				}
			}

			if respMap != nil && ce.Entry.Response.IsSuccess() && !isLROEnvelope(respMap) {
				if _, hasType := respMap["@type"]; !hasType {
					var respMsg protoreflect.MessageDescriptor
					if matchedRule != nil && matchedRule.respMsg != nil {
						respMsg = matchedRule.respMsg
					} else if inferredPrimaryMsg != nil && (method == "GET" || method == "POST" || method == "PATCH" || method == "PUT") {
						respMsg = inferredPrimaryMsg
					}
					if respMsg != nil {
						if err := unmarshalDynamicJSON(reg, respMsg, respMap); err != nil {
							errs = append(errs, fmt.Sprintf("stanza #%d (`%s`) response body failed protobuf schema `%s`: %v", ce.Index, method, respMsg.FullName(), err))
						} else if expectIntEnums {
							checkEnumEncodingRecursive(respMsg, respMap, "", ce.Index, method, &errs)
						}
					}
				}
			}
		}
	}

	return errs
}

// DefaultDescriptorSetPath locates .build/googleapis.pb if present in the workspace.
func DefaultDescriptorSetPath() string {
	defaultPBOnce.Do(func() {
		candidates := []string{
			".build/googleapis.pb",
			"../.build/googleapis.pb",
			"../../.build/googleapis.pb",
			"../../../.build/googleapis.pb",
			"../../../../.build/googleapis.pb",
			"../../../../../.build/googleapis.pb",
		}
		for _, c := range candidates {
			if fixture.FileExists(c) {
				if abs, err := filepath.Abs(c); err == nil {
					defaultPBPath = abs
					return
				}
				defaultPBPath = c
				return
			}
		}
	})
	return defaultPBPath
}

func getDescriptorRegistry(descriptorSetPath string) (*descriptorRegistry, bool, error) {
	if descriptorSetPath == "" {
		descriptorSetPath = DefaultDescriptorSetPath()
	}
	if descriptorSetPath == "" || !fixture.FileExists(descriptorSetPath) {
		return nil, false, nil
	}
	if absPath, err := filepath.Abs(descriptorSetPath); err == nil {
		descriptorSetPath = absPath
	}

	registryCacheMu.Lock()
	defer registryCacheMu.Unlock()
	if reg, ok := registryCache[descriptorSetPath]; ok {
		return reg, true, nil
	}
	if err, failed := registryErrCache[descriptorSetPath]; failed {
		return nil, false, err
	}

	reg, err := loadDescriptorRegistryFromDisk(descriptorSetPath)
	if err != nil {
		registryErrCache[descriptorSetPath] = err
		return nil, false, err
	}
	registryCache[descriptorSetPath] = reg
	return reg, true, nil
}

func loadDescriptorRegistryFromDisk(path string) (*descriptorRegistry, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fds := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(b, fds); err != nil {
		return nil, fmt.Errorf("unmarshaling %s: %w", path, err)
	}

	files, err := protodesc.NewFiles(fds)
	if err != nil {
		return nil, fmt.Errorf("protodesc.NewFiles(%s): %w", path, err)
	}

	if filepath.Base(path) == "googleapis.pb" {
		matches, _ := filepath.Glob(filepath.Join(filepath.Dir(path), "googleapis-*.pb"))
		sort.Slice(matches, func(i, j int) bool {
			fi, errI := os.Stat(matches[i])
			fj, errJ := os.Stat(matches[j])
			if errI == nil && errJ == nil && fi.Size() != fj.Size() {
				return fi.Size() > fj.Size()
			}
			return matches[i] < matches[j]
		})
		if len(matches) > 2 {
			matches = matches[:2]
		}
		for _, m := range matches {
			mb, err := os.ReadFile(m)
			if err != nil {
				continue
			}
			extra := &descriptorpb.FileDescriptorSet{}
			if err := proto.Unmarshal(mb, extra); err != nil {
				continue
			}
			for _, f := range extra.File {
				if _, err := files.FindFileByPath(f.GetName()); err == nil {
					continue
				}
				if fd, err := protodesc.NewFile(f, files); err == nil {
					_ = files.RegisterFile(fd)
				}
			}
		}
	}

	reg := &descriptorRegistry{
		files:     files,
		types:     dynamicpb.NewTypes(files),
		httpRules: make(map[string][]httpBindingRule),
		packages:  make(map[string]bool),
	}
	indexHTTPRules(files, reg.httpRules, reg.packages)
	return reg, nil
}

func indexHTTPRules(files *protoregistry.Files, out map[string][]httpBindingRule, packages map[string]bool) {
	files.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if pkg := string(fd.Package()); pkg != "" {
			packages[pkg] = true
		}
		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			svc := services.Get(i)
			host, _ := proto.GetExtension(svc.Options(), annotations.E_DefaultHost).(string)
			host = strings.ToLower(strings.TrimSpace(host))

			methods := svc.Methods()
			for j := 0; j < methods.Len(); j++ {
				m := methods.Get(j)
				httpRule, ok := proto.GetExtension(m.Options(), annotations.E_Http).(*annotations.HttpRule)
				if !ok || httpRule == nil {
					continue
				}
				collectMethodHTTPRules(host, m, httpRule, out)
			}
		}
		return true
	})
}

func collectMethodHTTPRules(host string, m protoreflect.MethodDescriptor, rule *annotations.HttpRule, out map[string][]httpBindingRule) {
	httpMethod, pattern := extractMethodAndPattern(rule)
	if httpMethod != "" && pattern != "" {
		var reqMsg protoreflect.MessageDescriptor
		bodySpec := strings.TrimSpace(rule.GetBody())
		if bodySpec == "*" {
			reqMsg = m.Input()
		} else if bodySpec != "" {
			if f := m.Input().Fields().ByName(protoreflect.Name(bodySpec)); f != nil && f.Kind() == protoreflect.MessageKind {
				reqMsg = f.Message()
			}
		}

		var respMsg protoreflect.MessageDescriptor
		if m.Output().FullName() != "google.longrunning.Operation" && !strings.HasSuffix(string(m.Output().FullName()), ".Operation") && m.Output().FullName() != "google.protobuf.Empty" {
			respMsg = m.Output()
		}

		out[host] = append(out[host], httpBindingRule{
			host:       host,
			httpMethod: httpMethod,
			segments:   expandHTTPPathTemplate(pattern),
			reqMsg:     reqMsg,
			respMsg:    respMsg,
		})
	}
	for _, add := range rule.GetAdditionalBindings() {
		collectMethodHTTPRules(host, m, add, out)
	}
}

func extractMethodAndPattern(rule *annotations.HttpRule) (string, string) {
	switch p := rule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		return "GET", p.Get
	case *annotations.HttpRule_Post:
		return "POST", p.Post
	case *annotations.HttpRule_Put:
		return "PUT", p.Put
	case *annotations.HttpRule_Patch:
		return "PATCH", p.Patch
	case *annotations.HttpRule_Delete:
		return "DELETE", p.Delete
	case *annotations.HttpRule_Custom:
		if p.Custom != nil {
			return strings.ToUpper(p.Custom.Kind), p.Custom.Path
		}
	}
	return "", ""
}

func expandHTTPPathTemplate(pattern string) []string {
	expanded := templateVarRegex.ReplaceAllStringFunc(pattern, func(m string) string {
		sub := templateVarRegex.FindStringSubmatch(m)
		if len(sub) == 3 && sub[2] != "" {
			return sub[2]
		}
		return "*"
	})
	if colonIdx := strings.Index(expanded, ":"); colonIdx != -1 {
		expanded = expanded[:colonIdx]
	}
	return splitNonEmptySegments(expanded)
}

func splitNonEmptySegments(p string) []string {
	var out []string
	for _, s := range strings.Split(strings.Trim(p, "/"), "/") {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func (r *descriptorRegistry) findMessageByFullName(fullName string) protoreflect.MessageDescriptor {
	fullName = strings.TrimPrefix(strings.TrimSpace(fullName), "type.googleapis.com/")
	fullName = strings.TrimPrefix(fullName, "/")
	if fullName == "" {
		return nil
	}
	desc, err := r.files.FindDescriptorByName(protoreflect.FullName(fullName))
	if err != nil {
		return nil
	}
	msg, ok := desc.(protoreflect.MessageDescriptor)
	if !ok {
		return nil
	}
	return msg
}

func (r *descriptorRegistry) matchHTTPRule(rawURL, httpMethod string) *httpBindingRule {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil
	}
	host := strings.ToLower(parsed.Hostname())
	uPath := parsed.Path
	if colonIdx := strings.Index(uPath, ":"); colonIdx != -1 {
		uPath = uPath[:colonIdx]
	}
	urlSegs := splitNonEmptySegments(uPath)

	for _, candidateHost := range []string{host, ""} {
		rules := r.httpRules[candidateHost]
		for i := range rules {
			rule := &rules[i]
			if rule.httpMethod == httpMethod && matchSegments(rule.segments, urlSegs) {
				return rule
			}
		}
	}
	return nil
}

func matchSegments(patternSegs, urlSegs []string) bool {
	if len(patternSegs) != len(urlSegs) {
		return false
	}
	for i := range patternSegs {
		if patternSegs[i] == "*" || patternSegs[i] == "**" {
			continue
		}
		if patternSegs[i] != urlSegs[i] {
			return false
		}
	}
	return true
}

func (r *descriptorRegistry) hasProtoPackageForType(typeURL string) bool {
	fullName := strings.TrimPrefix(strings.TrimSpace(typeURL), "type.googleapis.com/")
	fullName = strings.TrimPrefix(fullName, "/")
	lastDot := strings.LastIndex(fullName, ".")
	if lastDot <= 0 {
		return false
	}
	return r.packages[fullName[:lastDot]]
}

func validateTypedBlocksRecursive(reg *descriptorRegistry, stanzaIdx int, method, location string, curr map[string]any, expectIntEnums bool, errs *[]string) {
	if typeVal, hasType := curr["@type"]; hasType {
		typeURL, ok := typeVal.(string)
		if !ok || strings.TrimSpace(typeURL) == "" {
			*errs = append(*errs, fmt.Sprintf("stanza #%d (`%s`) %s has non-string or empty `@type`", stanzaIdx, method, location))
		} else {
			msgDesc := reg.findMessageByFullName(typeURL)
			if msgDesc == nil {
				*errs = append(*errs, fmt.Sprintf("stanza #%d (`%s`) %s has unknown protobuf `@type` %q not found in descriptor set", stanzaIdx, method, location, typeURL))
			} else {
				withoutType := make(map[string]any, len(curr))
				for k, v := range curr {
					if k != "@type" {
						withoutType[k] = v
					}
				}
				if err := unmarshalDynamicJSON(reg, msgDesc, withoutType); err != nil {
					*errs = append(*errs, fmt.Sprintf("stanza #%d (`%s`) %s failed protobuf schema for `@type` %q: %v", stanzaIdx, method, location, typeURL, err))
				} else if expectIntEnums {
					checkEnumEncodingRecursive(msgDesc, withoutType, "", stanzaIdx, method, errs)
				}
			}
		}
	}

	var keys []string
	for k := range curr {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		switch child := curr[k].(type) {
		case map[string]any:
			validateTypedBlocksRecursive(reg, stanzaIdx, method, location, child, expectIntEnums, errs)
		case []any:
			for _, elem := range child {
				if m, ok := elem.(map[string]any); ok {
					validateTypedBlocksRecursive(reg, stanzaIdx, method, location, m, expectIntEnums, errs)
				}
			}
		}
	}
}

func unmarshalDynamicJSON(reg *descriptorRegistry, msgDesc protoreflect.MessageDescriptor, obj map[string]any) error {
	rawBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	// Replace KCC's 21-digit normalized numeric ID placeholders ("000000000000000000000",
	// "111111111111111111111") which exceed 64-bit uint64 max (20 digits).
	s := strings.ReplaceAll(string(rawBytes), `"000000000000000000000"`, `"0"`)
	s = strings.ReplaceAll(s, `"111111111111111111111"`, `"1"`)
	dynMsg := dynamicpb.NewMessage(msgDesc)
	opts := protojson.UnmarshalOptions{
		DiscardUnknown: false,
		Resolver:       reg.types,
	}
	return opts.Unmarshal([]byte(s), dynMsg)
}

func checkEnumEncodingRecursive(msgDesc protoreflect.MessageDescriptor, obj map[string]any, prefix string, stanzaIdx int, method string, errs *[]string) {
	var keys []string
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fd := findProtoField(msgDesc, k)
		if fd == nil {
			continue
		}
		fieldPath := k
		if prefix != "" {
			fieldPath = prefix + "." + k
		}
		val := obj[k]

		if fd.Kind() == protoreflect.EnumKind && fd.Enum().FullName() != "google.protobuf.NullValue" {
			if fd.IsList() {
				if arr, ok := val.([]any); ok {
					for _, item := range arr {
						if s, isStr := item.(string); isStr {
							*errs = append(*errs, fmt.Sprintf("stanza #%d (`%s`): enum field `%s` has string value %q despite `enum-encoding=int` in request URL", stanzaIdx, method, fieldPath, s))
							break
						}
					}
				}
			} else if s, isStr := val.(string); isStr {
				*errs = append(*errs, fmt.Sprintf("stanza #%d (`%s`): enum field `%s` has string value %q despite `enum-encoding=int` in request URL", stanzaIdx, method, fieldPath, s))
			}
			continue
		}

		if fd.IsMap() {
			valDesc := fd.MapValue()
			if m, ok := val.(map[string]any); ok {
				if valDesc.Kind() == protoreflect.EnumKind && valDesc.Enum().FullName() != "google.protobuf.NullValue" {
					for mk, mv := range m {
						if s, isStr := mv.(string); isStr {
							*errs = append(*errs, fmt.Sprintf("stanza #%d (`%s`): enum map field `%s[%s]` has string value %q despite `enum-encoding=int` in request URL", stanzaIdx, method, fieldPath, mk, s))
							break
						}
					}
				} else if valDesc.Kind() == protoreflect.MessageKind {
					for mk, mv := range m {
						if childMap, ok := mv.(map[string]any); ok {
							checkEnumEncodingRecursive(valDesc.Message(), childMap, fieldPath+"."+mk, stanzaIdx, method, errs)
						}
					}
				}
			}
			continue
		}

		if fd.Kind() == protoreflect.MessageKind {
			if fd.IsList() {
				if arr, ok := val.([]any); ok {
					for _, item := range arr {
						if childMap, ok := item.(map[string]any); ok {
							checkEnumEncodingRecursive(fd.Message(), childMap, fieldPath+"[]", stanzaIdx, method, errs)
						}
					}
				}
			} else if childMap, ok := val.(map[string]any); ok {
				checkEnumEncodingRecursive(fd.Message(), childMap, fieldPath, stanzaIdx, method, errs)
			}
		}
	}
}

func findProtoField(msgDesc protoreflect.MessageDescriptor, jsonKey string) protoreflect.FieldDescriptor {
	fields := msgDesc.Fields()
	if fd := fields.ByJSONName(jsonKey); fd != nil {
		return fd
	}
	if fd := fields.ByName(protoreflect.Name(jsonKey)); fd != nil {
		return fd
	}
	return nil
}

func isLROEnvelope(m map[string]any) bool {
	if _, hasDone := m["done"]; hasDone {
		return true
	}
	if meta, ok := m["metadata"].(map[string]any); ok {
		if _, hasType := meta["@type"]; hasType {
			return true
		}
	}
	return false
}
