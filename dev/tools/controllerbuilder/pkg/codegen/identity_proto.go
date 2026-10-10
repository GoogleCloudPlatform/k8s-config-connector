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
	"strings"

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

func (g *IdentityGenerator) inspectProtoService(msg protoreflect.MessageDescriptor) (defaultHost string, serverGeneratedID bool) {
	if msg == nil {
		return "", false
	}
	wantCreate := "Create" + string(msg.Name())
	for _, fd := range g.serviceFilesForMessage(msg) {
		svcs := fd.Services()
		for i := 0; i < svcs.Len(); i++ {
			svc := svcs.Get(i)
			if defaultHost == "" {
				defaultHost = extractServiceDefaultHost(svc)
			}
			if hasServerGeneratedCreateMethod(svc, wantCreate) {
				serverGeneratedID = true
			}
		}
	}
	return defaultHost, serverGeneratedID
}

func extractServiceDefaultHost(svc protoreflect.ServiceDescriptor) string {
	h, _ := proto.GetExtension(svc.Options(), annotations.E_DefaultHost).(string)
	return h
}

func hasServerGeneratedCreateMethod(svc protoreflect.ServiceDescriptor, wantCreate string) bool {
	meths := svc.Methods()
	for j := 0; j < meths.Len(); j++ {
		m := meths.Get(j)
		if string(m.Name()) == wantCreate && isServerGeneratedCreateInput(m.Input().Fields()) {
			return true
		}
	}
	return false
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

func getResourceDescriptorProto(msg protoreflect.MessageDescriptor) *annotations.ResourceDescriptor {
	if msg == nil {
		return nil
	}
	v := proto.GetExtension(msg.Options(), annotations.E_Resource)
	rd, _ := v.(*annotations.ResourceDescriptor)
	return rd
}
