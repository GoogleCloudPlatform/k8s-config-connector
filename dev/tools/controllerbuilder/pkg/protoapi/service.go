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

// APIStyle names the convention a resource's standard methods were found by.
type APIStyle string

const (
	// StyleAIP means the methods follow the resource-oriented naming of
	// https://google.aip.dev/131 to 135: Get<Resource>, Create<Resource>,
	// Update<Resource>, Delete<Resource> and List<Resources>.
	StyleAIP APIStyle = "aip"
	// StyleCompute is the google.cloud.compute convention: one service per
	// collection, with methods named Get, Insert, Patch or Update, Delete and
	// List, and the resource passed in a <resource>_resource request field.
	StyleCompute APIStyle = "compute"
	// StyleNone means no service in the resource's proto package has a
	// standard method for it under either convention.
	StyleNone APIStyle = "none"
)

// ResourceAPI is what the proto says about managing one resource: the message,
// the service that hosts its standard methods, and the shape of each method.
// It is read from the descriptors alone, so a generator given the same
// descriptor set always sees the same model.
type ResourceAPI struct {
	// Message is the resource message.
	Message protoreflect.MessageDescriptor
	// Metadata is the message's google.api.resource annotation, or nil when
	// the message has none.
	Metadata *ResourceMetadata

	// Service hosts the standard methods below. It is nil when Style is
	// StyleNone.
	Service protoreflect.ServiceDescriptor
	// DefaultHost is Service's google.api.default_host, e.g.
	// "artifactregistry.googleapis.com". It is empty when Service is nil or
	// does not declare one.
	DefaultHost string
	// Style is the convention the standard methods were found by.
	Style APIStyle

	// The standard methods, each nil when Service has no such method. For
	// StyleAIP, they are found by their AIP names alone, so a Get, Create,
	// Update or Delete whose request lacks the fields the AIP gives it is
	// kept and flagged (see StandardMethod.NonstandardRequest). For
	// StyleCompute, Create is the Insert method and Update is Patch, or
	// Update when the service has no Patch.
	Get, Create, Update, Delete, List *StandardMethod

	// ServerGeneratedID reports that Create exists, takes a standard request
	// and has no field for a client-assigned ID, so the server names the
	// resource, or the client does in the body (see
	// ArchetypeServerGeneratedID). It is always false for StyleCompute, where
	// the client names the resource in the body's name field.
	ServerGeneratedID bool
	// HasEtag reports that the resource message has an etag field.
	HasEtag bool
	// HasLabels reports that the resource message has a map<string, string>
	// labels field.
	HasLabels bool
}

// StandardMethod describes one standard method of a resource.
type StandardMethod struct {
	// Descriptor is the method itself.
	Descriptor protoreflect.MethodDescriptor

	// NonstandardRequest reports that the method has the AIP name of a Get,
	// Create, Update or Delete, but its request lacks the fields the AIP
	// gives it (see hasAIPRequest), so code written for AIP-131 to AIP-135
	// cannot call it as is. A GetCluster that takes project_id, region and
	// cluster_name is one, and so is an UpdateCluster whose request holds a
	// ClusterUpdate rather than the Cluster. It is always false for a List,
	// whose request is not checked, and for StyleCompute.
	NonstandardRequest bool

	// HTTPVerb is GET, POST, PUT, PATCH, DELETE or CUSTOM, taken from the
	// method's google.api.http binding. It is empty when the method has no
	// binding.
	HTTPVerb string
	// HTTPPath is the URL template of the binding, e.g.
	// "/v1/{parent=projects/*/locations/*}/repositories".
	HTTPPath string

	// Names of well-known request fields, each empty when the request has no
	// such field.
	//
	// ParentField is "parent" and NameField is "name". ResourceField is the
	// first field whose type is the resource message: the body of a Create or
	// Update. IDField is only set on an AIP Create; it is the field carrying
	// the client-assigned ID, "<resource>_id" (AIP-133), or failing that a
	// string field named after the last words of the resource name, such as
	// "trigger_id" for a BuildTrigger (see clientIDField). The rest are the
	// fields of those names: update_mask (AIP-134), etag (AIP-154),
	// allow_missing, validate_only (AIP-163) and request_id (AIP-155).
	ParentField       string
	NameField         string
	ResourceField     string
	IDField           string
	UpdateMaskField   string
	EtagField         string
	AllowMissingField string
	ValidateOnlyField string
	RequestIDField    string

	// MethodSignatures holds the method's google.api.method_signature
	// options, e.g. "parent,repository,repository_id".
	MethodSignatures []string

	// LRO describes the long-running operation the method starts. It is nil
	// unless the method returns google.longrunning.Operation, so it is always
	// nil for StyleCompute, whose methods return
	// google.cloud.compute.v1.Operation.
	LRO *LROInfo

	// Response is the full name of the message the method returns, e.g.
	// "google.longrunning.Operation" or "google.protobuf.Empty".
	Response string
}

// LROInfo is the google.longrunning.operation_info option of a method that
// returns google.longrunning.Operation.
type LROInfo struct {
	// ResponseType is the full name of the message the finished operation
	// holds, e.g. "google.devtools.artifactregistry.v1.Repository". It is
	// empty when the method declares no operation_info.
	ResponseType string
	// MetadataType is the full name of the operation's metadata message, e.g.
	// "google.devtools.artifactregistry.v1.OperationMetadata". It is empty
	// when the method declares no operation_info.
	MetadataType string
}

const (
	// longRunningOperation is the message an AIP-151 long-running method
	// returns.
	longRunningOperation protoreflect.FullName = "google.longrunning.Operation"

	// operationInfoField is the extension number of
	// google.longrunning.operation_info on google.protobuf.MethodOptions, and
	// operationInfoResponseType and operationInfoMetadataType are the fields
	// of its OperationInfo message. The Go type for the extension lives in
	// cloud.google.com/go/longrunning, which this module does not depend on,
	// so the option is read from the wire format instead.
	operationInfoField        protowire.Number = 1049
	operationInfoResponseType protowire.Number = 1
	operationInfoMetadataType protowire.Number = 2
)

// ResourceAPIByName looks up a message by its full name and builds its
// ResourceAPI. The error wraps protoregistry.NotFound when the descriptor set
// has no message of that name.
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

// ResourceAPIFor builds the API model for a resource message.
//
// The standard methods are looked up in the services of the message's own
// proto package. A service in the StyleCompute convention wins; otherwise the
// methods are found by their AIP names, and those whose requests lack the
// shape the AIPs give them are flagged (see
// StandardMethod.NonstandardRequest). When several services qualify, the one
// with the most standard methods wins, flagged or not. A tie goes to the
// service with the shorter name, then to the one whose full name sorts
// first. All the methods come from that one service.
func (p *Proto) ResourceAPIFor(msg protoreflect.MessageDescriptor) *ResourceAPI {
	if msg == nil {
		return nil
	}
	api := &ResourceAPI{
		Message:   msg,
		Metadata:  GetResourceMetadata(msg),
		Style:     StyleNone,
		HasEtag:   msg.Fields().ByName("etag") != nil,
		HasLabels: isStringMap(msg.Fields().ByName("labels")),
	}

	services := p.servicesInPackage(msg.ParentFile().Package())
	if svc, methods := pickService(services, msg, computeMethods); svc != nil {
		api.Style = StyleCompute
		api.setMethods(svc, methods, false)
	} else if svc, methods := pickService(services, msg, aipMethods); svc != nil {
		api.Style = StyleAIP
		api.setMethods(svc, methods, true)
		api.ServerGeneratedID = api.Create != nil && !api.Create.NonstandardRequest && api.Create.IDField == ""
	}
	return api
}

// setMethods fills in svc and its standard methods. aip means the methods
// were found by their AIP names, so their requests are checked against the
// AIP shape and Create's client-assigned ID field is looked for.
func (api *ResourceAPI) setMethods(svc protoreflect.ServiceDescriptor, methods methodSet, aip bool) {
	api.Service = svc
	api.DefaultHost, _ = proto.GetExtension(svc.Options(), annotations.E_DefaultHost).(string)
	api.Get = newStandardMethod(methods.get, "Get", api.Message, aip)
	api.Create = newStandardMethod(methods.create, "Create", api.Message, aip)
	api.Update = newStandardMethod(methods.update, "Update", api.Message, aip)
	api.Delete = newStandardMethod(methods.delete, "Delete", api.Message, aip)
	api.List = newStandardMethod(methods.list, "List", api.Message, aip)
}

// servicesInPackage returns the services declared in pkg, sorted by full name
// so that every choice made over them is deterministic.
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

// methodSet holds the standard methods found in one service.
type methodSet struct {
	get, create, update, delete, list protoreflect.MethodDescriptor
}

func (s methodSet) count() int {
	n := 0
	for _, m := range []protoreflect.MethodDescriptor{s.get, s.create, s.update, s.delete, s.list} {
		if m != nil {
			n++
		}
	}
	return n
}

// pickService returns the service in which find finds the most standard
// methods for msg, or nil when it finds none. A tie goes to the service with
// the shorter name, so Autoscalers beats RegionAutoscalers, then to the one
// whose full name sorts first, so the order of services does not matter.
func pickService(services []protoreflect.ServiceDescriptor, msg protoreflect.MessageDescriptor,
	find func(protoreflect.ServiceDescriptor, protoreflect.MessageDescriptor) methodSet,
) (protoreflect.ServiceDescriptor, methodSet) {
	var best protoreflect.ServiceDescriptor
	var bestMethods methodSet
	for _, svc := range services {
		methods := find(svc, msg)
		n, bestN := methods.count(), bestMethods.count()
		if n > bestN || (n > 0 && n == bestN && preferService(svc, best)) {
			best, bestMethods = svc, methods
		}
	}
	return best, bestMethods
}

// preferService reports whether a wins a tie against b: the shorter name
// wins, then the full name that sorts first.
func preferService(a, b protoreflect.ServiceDescriptor) bool {
	if len(a.Name()) != len(b.Name()) {
		return len(a.Name()) < len(b.Name())
	}
	return a.FullName() < b.FullName()
}

// aipMethods finds msg's standard methods in svc by their AIP names. The
// requests are not checked here: a GetCluster that takes project_id, region
// and cluster_name is still msg's Get, and newStandardMethod flags it.
func aipMethods(svc protoreflect.ServiceDescriptor, msg protoreflect.MessageDescriptor) methodSet {
	byName := func(verb string) protoreflect.MethodDescriptor {
		return svc.Methods().ByName(protoreflect.Name(verb + string(msg.Name())))
	}
	return methodSet{
		get:    byName("Get"),
		create: byName("Create"),
		update: byName("Update"),
		delete: byName("Delete"),
		list:   aipListMethod(svc, msg),
	}
}

// hasAIPRequest reports whether req, the request of the standard method verb
// for resource, has the fields that code written for AIP-131 to AIP-135
// relies on:
//   - Get and Delete take the resource's name in a name field;
//   - Create takes a parent field, or the resource in a field of its type,
//     since a top-level resource such as a Project or a TagKey has no parent;
//   - Update takes the resource in a field of its type.
//
// Any other verb, List included, passes.
func hasAIPRequest(verb string, req, resource protoreflect.MessageDescriptor) bool {
	switch verb {
	case "Get", "Delete":
		return req.Fields().ByName("name") != nil
	case "Create":
		return req.Fields().ByName("parent") != nil || fieldOfType(req, resource) != ""
	case "Update":
		return fieldOfType(req, resource) != ""
	}
	return true
}

// aipListMethod finds msg's List method in svc: a method named List<Plural>
// whose response has a repeated field of the resource's type, or failing that
// one named List<Resource>.... The response check catches irregular plurals
// such as Repository and ListRepositories, which the name alone misses. A
// method named just List is the compute convention, not AIP-132, so it never
// matches. When several methods qualify, a response match beats a name
// match, then the shorter name wins, so ListInstances beats
// ListInstanceRevisions, then the name that sorts first.
func aipListMethod(svc protoreflect.ServiceDescriptor, msg protoreflect.MessageDescriptor) protoreflect.MethodDescriptor {
	var best protoreflect.MethodDescriptor
	bestByResponse := false
	methods := svc.Methods()
	for i := 0; i < methods.Len(); i++ {
		m := methods.Get(i)
		name := string(m.Name())
		if !strings.HasPrefix(name, "List") || name == "List" {
			continue
		}
		byResponse := returnsRepeated(m, msg)
		byName := strings.HasPrefix(strings.TrimPrefix(name, "List"), string(msg.Name()))
		if !byResponse && !byName {
			continue
		}
		if best == nil || betterListMethod(m, byResponse, best, bestByResponse) {
			best, bestByResponse = m, byResponse
		}
	}
	return best
}

func betterListMethod(m protoreflect.MethodDescriptor, byResponse bool, than protoreflect.MethodDescriptor, thanByResponse bool) bool {
	if byResponse != thanByResponse {
		return byResponse
	}
	a, b := string(m.Name()), string(than.Name())
	if len(a) != len(b) {
		return len(a) < len(b)
	}
	return a < b
}

// returnsRepeated reports whether m's response has a repeated field of msg's
// type, the way a List response carries its page of resources.
func returnsRepeated(m protoreflect.MethodDescriptor, msg protoreflect.MessageDescriptor) bool {
	fields := m.Output().Fields()
	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		if f.Cardinality() == protoreflect.Repeated && !f.IsMap() && f.Message() != nil && f.Message().FullName() == msg.FullName() {
			return true
		}
	}
	return false
}

// computeMethods finds msg's standard methods in svc if svc follows the
// google.cloud.compute convention for msg: it has Get and Insert methods, and
// the Insert request carries the resource in a <resource>_resource field of
// msg's type. It returns an empty set otherwise.
func computeMethods(svc protoreflect.ServiceDescriptor, msg protoreflect.MessageDescriptor) methodSet {
	methods := svc.Methods()
	get, insert := methods.ByName("Get"), methods.ByName("Insert")
	if get == nil || insert == nil {
		return methodSet{}
	}
	body := insert.Input().Fields().ByName(protoreflect.Name(snakeCase(string(msg.Name())) + "_resource"))
	if body == nil || body.Message() == nil || body.Message().FullName() != msg.FullName() {
		return methodSet{}
	}
	update := methods.ByName("Patch")
	if update == nil {
		update = methods.ByName("Update")
	}
	return methodSet{
		get:    get,
		create: insert,
		update: update,
		delete: methods.ByName("Delete"),
		list:   methods.ByName("List"),
	}
}

// newStandardMethod describes m, the standard method verb ("Get", "Create",
// "Update", "Delete" or "List") of resource, or returns nil when m is nil.
// aip means m was found by its AIP name, so its request is checked against
// the AIP shape and, for a Create, IDField is filled in.
func newStandardMethod(m protoreflect.MethodDescriptor, verb string, resource protoreflect.MessageDescriptor, aip bool) *StandardMethod {
	if m == nil {
		return nil
	}
	req := m.Input()
	sm := &StandardMethod{
		Descriptor:        m,
		ParentField:       fieldIfPresent(req, "parent"),
		NameField:         fieldIfPresent(req, "name"),
		ResourceField:     fieldOfType(req, resource),
		UpdateMaskField:   fieldIfPresent(req, "update_mask"),
		EtagField:         fieldIfPresent(req, "etag"),
		AllowMissingField: fieldIfPresent(req, "allow_missing"),
		ValidateOnlyField: fieldIfPresent(req, "validate_only"),
		RequestIDField:    fieldIfPresent(req, "request_id"),
		Response:          string(m.Output().FullName()),
	}
	if aip {
		sm.NonstandardRequest = !hasAIPRequest(verb, req, resource)
		if verb == "Create" {
			sm.IDField = clientIDField(req, resource)
		}
	}
	sm.HTTPVerb, sm.HTTPPath = httpBinding(m)
	sm.MethodSignatures, _ = proto.GetExtension(m.Options(), annotations.E_MethodSignature).([]string)
	if m.Output().FullName() == longRunningOperation {
		sm.LRO = &LROInfo{}
		if responseType, metadataType, ok := operationInfo(m); ok {
			pkg := m.ParentFile().Package()
			sm.LRO.ResponseType = qualifyTypeName(pkg, responseType)
			sm.LRO.MetadataType = qualifyTypeName(pkg, metadataType)
		}
	}
	return sm
}

func fieldIfPresent(msg protoreflect.MessageDescriptor, name protoreflect.Name) string {
	if msg.Fields().ByName(name) == nil {
		return ""
	}
	return string(name)
}

// fieldOfType returns the first field of msg whose type is the message want.
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

// clientIDField returns the field of a Create request that carries the
// client-assigned ID: <resource>_id (AIP-133), with the resource name in
// snake_case. Failing that, it is a string field <suffix>_id, where <suffix>
// is a word-boundary suffix of that name, the longest first: trigger_id for a
// BuildTrigger, config_id for a NotificationConfig. Any other _id field, such
// as project_id, names something else.
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

// httpBinding returns the verb and URL template of m's google.api.http
// option, or empty strings when it has none.
func httpBinding(m protoreflect.MethodDescriptor) (verb, path string) {
	rule, _ := proto.GetExtension(m.Options(), annotations.E_Http).(*annotations.HttpRule)
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
		return "CUSTOM", p.Custom.GetPath()
	}
	return "", ""
}

// operationInfo reads the google.longrunning.operation_info option of m, as
// written in the proto: the type names may be relative to m's package. ok is
// false when m has no such option.
//
// The options are marshalled and scanned for the extension's field number, so
// the option is found whether it was parsed as an unknown field, which is the
// case in this module, or as a registered extension.
func operationInfo(m protoreflect.MethodDescriptor) (responseType, metadataType string, ok bool) {
	opts := m.Options()
	if opts == nil {
		return "", "", false
	}
	b, err := proto.Marshal(opts)
	if err != nil {
		return "", "", false
	}
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return "", "", false
		}
		b = b[n:]
		if num == operationInfoField && typ == protowire.BytesType {
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return "", "", false
			}
			b = b[n:]
			// A message field that occurs more than once is merged, so
			// later values win.
			r, md := parseOperationInfo(v)
			if r != "" {
				responseType = r
			}
			if md != "" {
				metadataType = md
			}
			ok = true
			continue
		}
		n = protowire.ConsumeFieldValue(num, typ, b)
		if n < 0 {
			return "", "", false
		}
		b = b[n:]
	}
	return responseType, metadataType, ok
}

// parseOperationInfo decodes a google.longrunning.OperationInfo message.
func parseOperationInfo(b []byte) (responseType, metadataType string) {
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return responseType, metadataType
		}
		b = b[n:]
		if typ == protowire.BytesType && (num == operationInfoResponseType || num == operationInfoMetadataType) {
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return responseType, metadataType
			}
			b = b[n:]
			if num == operationInfoResponseType {
				responseType = string(v)
			} else {
				metadataType = string(v)
			}
			continue
		}
		n = protowire.ConsumeFieldValue(num, typ, b)
		if n < 0 {
			return responseType, metadataType
		}
		b = b[n:]
	}
	return responseType, metadataType
}

// qualifyTypeName resolves a type name from an operation_info option the way
// the option documents it: a name with no dot is a message in the method's
// own package, and anything else is already fully qualified.
func qualifyTypeName(pkg protoreflect.FullName, name string) string {
	name = strings.TrimPrefix(strings.TrimSpace(name), ".")
	if name == "" || strings.Contains(name, ".") || pkg == "" {
		return name
	}
	return string(pkg) + "." + name
}

// snakeCase converts a message name to the lower_snake_case that request
// fields derived from it use: "Repository" -> "repository",
// "TlsInspectionPolicy" -> "tls_inspection_policy", "SACRealm" ->
// "sac_realm", "Ipv6Range" -> "ipv6_range".
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
