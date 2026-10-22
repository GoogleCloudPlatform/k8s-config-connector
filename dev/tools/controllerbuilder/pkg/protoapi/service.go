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

package protoapi

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// ResourceAPI captures the proto service and CRUD method descriptors for a
// single resource message.
type ResourceAPI struct {
	// Message is the resource message descriptor.
	Message protoreflect.MessageDescriptor

	// Service hosts the standard CRUD methods below, or is nil when no
	// service in the resource's proto package defines any standard method.
	Service protoreflect.ServiceDescriptor

	// DefaultHost is Service's google.api.default_host (e.g. "cloudkms.googleapis.com").
	DefaultHost string

	// Standard CRUD methods (each nil when Service has no such method).
	Get, List, Create, Update, Delete *StandardMethod

	// HasLabels reports whether the resource message has a map<string, string> labels field.
	HasLabels bool
}

// StandardMethod describes one standard CRUD method of a resource.
type StandardMethod struct {
	Descriptor protoreflect.MethodDescriptor

	// Well-known request field names (empty when absent on the request message).
	ParentField     string
	NameField       string
	ResourceField   string
	IDField         string
	UpdateMaskField string
	ForceField      string

	// LRO is non-nil when the method returns google.longrunning.Operation.
	LRO *LROInfo

	// Response is the full proto name of the method's output message.
	Response string
}

// LROInfo holds the google.longrunning.operation_info response type for an LRO method.
type LROInfo struct {
	ResponseType string
}

const (
	longRunningOperation      protoreflect.FullName = "google.longrunning.Operation"
	operationInfoField        protowire.Number      = 1049
	operationInfoResponseType protowire.Number      = 1
)

// ResourceAPIByName looks up a message by its full name and builds its ResourceAPI.
func (p *Proto) ResourceAPIByName(name protoreflect.FullName) (*ResourceAPI, error) {
	d, err := p.files.FindDescriptorByName(name)
	if err != nil {
		return nil, fmt.Errorf("finding message %q: %w", name, err)
	}
	msg, ok := d.(protoreflect.MessageDescriptor)
	if !ok {
		return nil, fmt.Errorf("%q is not a message: %w", name, protoregistry.NotFound)
	}
	return p.ResourceAPIFor(msg), nil
}

// ResourceAPIFor builds the API model for a resource message from the services
// in its proto package.
func (p *Proto) ResourceAPIFor(msg protoreflect.MessageDescriptor) *ResourceAPI {
	if msg == nil {
		return nil
	}
	api := &ResourceAPI{
		Message:   msg,
		HasLabels: isStringMap(msg.Fields().ByName("labels")),
	}

	services := p.servicesInPackage(msg.ParentFile().Package())
	if svc, methods := pickService(services, msg); svc != nil {
		api.Service = svc
		api.DefaultHost, _ = proto.GetExtension(svc.Options(), annotations.E_DefaultHost).(string)
		api.Get = newStandardMethod(methods.get, "Get", msg)
		api.List = newStandardMethod(methods.list, "List", msg)
		api.Create = newStandardMethod(methods.create, "Create", msg)
		api.Update = newStandardMethod(methods.update, "Update", msg)
		api.Delete = newStandardMethod(methods.delete, "Delete", msg)
	}
	return api
}

func (p *Proto) servicesInPackage(pkg protoreflect.FullName) []protoreflect.ServiceDescriptor {
	var services []protoreflect.ServiceDescriptor
	p.files.RangeFilesByPackage(pkg, func(f protoreflect.FileDescriptor) bool {
		for i := 0; i < f.Services().Len(); i++ {
			services = append(services, f.Services().Get(i))
		}
		return true
	})
	sort.Slice(services, func(i, j int) bool {
		return services[i].FullName() < services[j].FullName()
	})
	return services
}

type methodSet struct {
	get, list, create, update, delete protoreflect.MethodDescriptor
}

func pickService(services []protoreflect.ServiceDescriptor, msg protoreflect.MessageDescriptor) (protoreflect.ServiceDescriptor, methodSet) {
	var best protoreflect.ServiceDescriptor
	var bestMethods methodSet
	for _, svc := range services {
		methods := aipMethods(svc, msg)
		n, bestN := methods.count(), bestMethods.count()
		if n > bestN || (n > 0 && n == bestN && preferService(svc, best)) {
			best, bestMethods = svc, methods
		}
	}
	return best, bestMethods
}

func (s methodSet) count() int {
	n := 0
	for _, m := range []protoreflect.MethodDescriptor{s.get, s.list, s.create, s.update, s.delete} {
		if m != nil {
			n++
		}
	}
	return n
}

func preferService(a, b protoreflect.ServiceDescriptor) bool {
	if len(a.Name()) != len(b.Name()) {
		return len(a.Name()) < len(b.Name())
	}
	return a.FullName() < b.FullName()
}

func aipMethods(svc protoreflect.ServiceDescriptor, msg protoreflect.MessageDescriptor) methodSet {
	byName := func(verb string) protoreflect.MethodDescriptor {
		return svc.Methods().ByName(protoreflect.Name(verb + string(msg.Name())))
	}
	var list protoreflect.MethodDescriptor
	for i := 0; i < svc.Methods().Len(); i++ {
		m := svc.Methods().Get(i)
		if strings.HasPrefix(string(m.Name()), "List") && hasRepeatedFieldOfType(m.Output(), msg) {
			list = m
			break
		}
	}
	return methodSet{
		get:    byName("Get"),
		list:   list,
		create: byName("Create"),
		update: byName("Update"),
		delete: byName("Delete"),
	}
}

func hasRepeatedFieldOfType(msg, target protoreflect.MessageDescriptor) bool {
	for i := 0; i < msg.Fields().Len(); i++ {
		fd := msg.Fields().Get(i)
		if fd.IsList() && fd.Kind() == protoreflect.MessageKind && fd.Message().FullName() == target.FullName() {
			return true
		}
	}
	return false
}

func newStandardMethod(m protoreflect.MethodDescriptor, verb string, resource protoreflect.MessageDescriptor) *StandardMethod {
	if m == nil {
		return nil
	}
	req := m.Input()
	nameField := fieldIfPresent(req, "name")
	if nameField == "" && (verb == "Get" || verb == "Delete") && resource != nil {
		resField := protoreflect.Name(snakeCase(string(resource.Name())))
		if f := req.Fields().ByName(resField); f != nil && f.Kind() == protoreflect.StringKind {
			nameField = string(f.Name())
		}
	}
	sm := &StandardMethod{
		Descriptor:    m,
		ParentField:   fieldIfPresent(req, "parent"),
		NameField:     nameField,
		ResourceField: fieldOfType(req, resource),
		ForceField:    fieldIfPresent(req, "force"),
		Response:      string(m.Output().FullName()),
	}
	if verb == "Update" {
		sm.UpdateMaskField = updateMaskFieldIfPresent(req)
	}
	if verb == "Create" {
		sm.IDField = clientIDField(req, resource)
	}
	if m.Output().FullName() == longRunningOperation {
		sm.LRO = &LROInfo{}
		if responseType, ok := operationResponseType(m); ok {
			sm.LRO.ResponseType = qualifyTypeName(m.ParentFile().Package(), responseType)
		}
	}
	return sm
}

func updateMaskFieldIfPresent(msg protoreflect.MessageDescriptor) string {
	if f := msg.Fields().ByName("update_mask"); isFieldMask(f) {
		return "update_mask"
	}
	fields := msg.Fields()
	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		if isFieldMask(f) && !strings.Contains(string(f.Name()), "read") {
			return string(f.Name())
		}
	}
	return ""
}

func isFieldMask(f protoreflect.FieldDescriptor) bool {
	return f != nil && !f.IsList() && f.Message() != nil && f.Message().FullName() == "google.protobuf.FieldMask"
}

func fieldIfPresent(msg protoreflect.MessageDescriptor, name protoreflect.Name) string {
	if msg.Fields().ByName(name) == nil {
		return ""
	}
	return string(name)
}

func fieldOfType(msg protoreflect.MessageDescriptor, want protoreflect.MessageDescriptor) string {
	fields := msg.Fields()
	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		if f.Message() != nil && !f.IsMap() && f.Message().FullName() == want.FullName() {
			return string(f.Name())
		}
	}
	return ""
}

func clientIDField(req protoreflect.MessageDescriptor, resource protoreflect.MessageDescriptor) string {
	name := snakeCase(string(resource.Name()))
	if f := req.Fields().ByName(protoreflect.Name(name + "_id")); f != nil {
		return string(f.Name())
	}
	words := strings.Split(name, "_")
	for i := 1; i < len(words); i++ {
		f := req.Fields().ByName(protoreflect.Name(strings.Join(words[i:], "_") + "_id"))
		if f != nil && f.Kind() == protoreflect.StringKind {
			return string(f.Name())
		}
	}
	return ""
}

func isStringMap(f protoreflect.FieldDescriptor) bool {
	return f != nil && f.IsMap() &&
		f.MapKey().Kind() == protoreflect.StringKind &&
		f.MapValue().Kind() == protoreflect.StringKind
}

func operationResponseType(m protoreflect.MethodDescriptor) (responseType string, ok bool) {
	opts := m.Options()
	if opts == nil {
		return "", false
	}
	b, err := proto.Marshal(opts)
	if err != nil {
		return "", false
	}
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return "", false
		}
		b = b[n:]
		if num == operationInfoField && typ == protowire.BytesType {
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return "", false
			}
			b = b[n:]
			if r := parseOperationResponseType(v); r != "" {
				responseType = r
			}
			ok = true
			continue
		}
		n = protowire.ConsumeFieldValue(num, typ, b)
		if n < 0 {
			return "", false
		}
		b = b[n:]
	}
	return responseType, ok
}

func parseOperationResponseType(b []byte) string {
	var responseType string
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return responseType
		}
		b = b[n:]
		if typ == protowire.BytesType && num == operationInfoResponseType {
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return responseType
			}
			b = b[n:]
			responseType = string(v)
			continue
		}
		n = protowire.ConsumeFieldValue(num, typ, b)
		if n < 0 {
			return responseType
		}
		b = b[n:]
	}
	return responseType
}

func qualifyTypeName(pkg protoreflect.FullName, name string) string {
	name = strings.TrimPrefix(strings.TrimSpace(name), ".")
	if name == "" || strings.Contains(name, ".") || pkg == "" {
		return name
	}
	return string(pkg) + "." + name
}

func snakeCase(s string) string {
	runes := []rune(s)
	var b strings.Builder
	for i, r := range runes {
		if !unicode.IsUpper(r) {
			b.WriteRune(r)
			continue
		}
		if i > 0 {
			prev := runes[i-1]
			nextIsLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) || (unicode.IsUpper(prev) && nextIsLower) {
				b.WriteByte('_')
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
