// Copyright 2024 Google LLC
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

package protogen

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/klog/v2"
)

type ProtoWriter struct {
	w            io.Writer
	errors       []error
	protoVersion int
	comments     *Comments
}

func NewProtoWriter(w io.Writer) *ProtoWriter {
	return &ProtoWriter{
		w:            w,
		protoVersion: 3,
	}
}

func (p *ProtoWriter) SetProtoVersion(protoVersion int) {
	p.protoVersion = protoVersion
}

func (p *ProtoWriter) SetComments(comments *Comments) {
	p.comments = comments
}

func (p *ProtoWriter) getComment(obj protoreflect.FullName) string {
	if p.comments == nil {
		return ""
	}
	return p.comments.GetComment(string(obj))
}

func (p *ProtoWriter) Error() error {
	if len(p.errors) > 0 {
		return errors.Join(p.errors...)
	}
	return nil
}

func (p *ProtoWriter) writeString(s string) {
	if len(p.errors) != 0 {
		return
	}
	_, err := p.w.Write([]byte(s))
	if err != nil {
		p.errors = append(p.errors, err)
	}
}

func (p *ProtoWriter) errorf(msg string, args ...any) {
	p.errors = append(p.errors, fmt.Errorf(msg, args...))
}

func (p *ProtoWriter) printf(format string, args ...any) {
	p.writeString(fmt.Sprintf(format, args...))
}

func (p *ProtoWriter) nameForMessageType(md protoreflect.MessageDescriptor) string {
	messageName := string(md.FullName())
	if strings.HasPrefix(messageName, "*.") {
		messageName = strings.TrimPrefix(messageName, "*")
	} else {
		messageName = string(md.Name())
	}
	return messageName
}

func (p *ProtoWriter) renderField(fd protoreflect.FieldDescriptor) {
	comment := p.getComment(fd.FullName())
	if comment != "" {
		for _, line := range strings.Split(comment, "\n") {
			p.printf("  // %s\n", line)
		}
	}

	var b bytes.Buffer
	b.WriteString("  ")

	if fd.IsMap() {
		b.WriteString("map<")
		switch fd.MapKey().Kind() {
		case protoreflect.StringKind:
			b.WriteString("string")
		default:
			p.errorf("unhandled MapKey field Kind %v", fd.MapKey())
		}
		b.WriteString(", ")
		switch fd.MapValue().Kind() {
		case protoreflect.StringKind:
			b.WriteString("string")
		case protoreflect.Int32Kind:
			b.WriteString("int32")
		case protoreflect.Int64Kind:
			b.WriteString("int64")
		case protoreflect.MessageKind:
			b.WriteString(p.nameForMessageType(fd.MapValue().Message()))
		case protoreflect.DoubleKind:
			b.WriteString("double")
		default:
			p.errorf("unhandled MapValue field Kind %v", fd.MapValue())
		}
		b.WriteString("> ")
	} else {
		switch fd.Cardinality() {
		case protoreflect.Repeated:
			b.WriteString("repeated ")
		case protoreflect.Optional:
			if p.protoVersion == 2 {
				b.WriteString("optional ")
			}
		default:
			p.errorf("unhandled field Cardinality %v", fd.Cardinality())
		}

		switch fd.Kind() {
		case protoreflect.StringKind:
			b.WriteString("string ")

		case protoreflect.BytesKind:
			b.WriteString("bytes ")

		case protoreflect.BoolKind:
			b.WriteString("bool ")

		case protoreflect.Uint32Kind:
			b.WriteString("uint32 ")

		case protoreflect.Int32Kind:
			b.WriteString("int32 ")

		case protoreflect.Uint64Kind:
			b.WriteString("uint64 ")

		case protoreflect.Int64Kind:
			b.WriteString("int64 ")

		case protoreflect.DoubleKind:
			b.WriteString("double ")

		case protoreflect.FloatKind:
			b.WriteString("float ")

		case protoreflect.MessageKind:
			b.WriteString(p.nameForMessageType(fd.Message()) + " ")

		default:
			p.errorf("unhandled field Kind %v", fd.Kind())
		}
	}

	fmt.Fprintf(&b, "%s = %d", fd.Name(), fd.Number())
	if fd.HasJSONName() && fd.JSONName() != "" {
		fmt.Fprintf(&b, " [json_name=%q]", fd.JSONName())
	}
	fmt.Fprintf(&b, ";\n")
	p.writeString(b.String())

}

func (p *ProtoWriter) renderMessage(msg protoreflect.MessageDescriptor) {
	p.printf("\n")
	comment := p.getComment(msg.FullName())
	if comment != "" {
		for _, line := range strings.Split(comment, "\n") {
			p.printf("// %s\n", line)
		}
	}

	p.printf("message %s {\n", msg.Name())
	fields := msg.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		p.renderField(field)
	}

	p.printf("}\n")
}

func (p *ProtoWriter) renderMethod(md protoreflect.MethodDescriptor) {
	p.printf("\n")
	comment := p.getComment(md.FullName())
	if comment != "" {
		for _, line := range strings.Split(comment, "\n") {
			p.printf("  // %s\n", line)
		}
	}

	var b bytes.Buffer
	b.WriteString("  rpc ")

	b.WriteString(string(md.Name()))
	b.WriteString("(")
	b.WriteString(p.nameForMessageType(md.Input()))
	b.WriteString(")")
	b.WriteString(" returns ")
	b.WriteString("(")
	b.WriteString(p.nameForMessageType(md.Output()))
	b.WriteString(")")

	options := md.Options()
	if options != nil {
		b.WriteString(" {\n")
		proto.RangeExtensions(options, func(xt protoreflect.ExtensionType, v interface{}) bool {
			b.WriteString(fmt.Sprintf("    option (%s) = {\n", xt.TypeDescriptor().FullName()))
			formatted, err := prototext.MarshalOptions{Multiline: true}.Marshal(v.(proto.Message))
			if err != nil {
				p.errors = append(p.errors, err)
			}
			for _, line := range strings.Split(string(formatted), "\n") {
				if line == "" {
					continue
				}
				// Undo the randomization (deliberately) injected by prototext
				line = strings.Replace(line, ":  ", ": ", 1)
				// Add indent (MarshalOptions.Indent just doesn't seem to work...)
				b.WriteString("      ")
				b.WriteString(line)
				b.WriteString("\n")
			}
			b.WriteString("    };\n")
			return true
		})
		b.WriteString("  }")
	}

	b.WriteString(";\n")
	p.writeString(b.String())

}

func (p *ProtoWriter) renderService(msg protoreflect.ServiceDescriptor) {
	p.printf("\n")
	p.printf("service %s {\n", msg.Name())
	methods := msg.Methods()
	for i := 0; i < methods.Len(); i++ {
		field := methods.Get(i)
		p.renderMethod(field)
	}

	p.printf("}\n")
}

func (p *ProtoWriter) WriteFile(file protoreflect.FileDescriptor) {
	klog.Infof("file %v", file.Name())
	p.printf("syntax = \"proto%d\";\n", p.protoVersion)
	p.printf("package %s;\n", file.Package())

	importPaths := []string{
		"google/api/annotations.proto",
	}
	for i := 0; i < file.Imports().Len(); i++ {
		m := file.Imports().Get(i)
		importPaths = append(importPaths, m.FileDescriptor.Path())
	}
	sort.Strings(importPaths)
	for _, importPath := range importPaths {
		p.printf("import %q;\n", importPath)
	}

	fileOptions := file.Options()
	switch fileOptions := fileOptions.(type) {
	case nil:
	case *descriptorpb.FileOptions:
		if fileOptions.GetGoPackage() != "" {
			p.printf("option go_package = \"%s\";\n", fileOptions.GetGoPackage())
		}
	default:
		p.errorf("unhandled fileOptions type %T", fileOptions)
	}

	services := file.Services()
	for i := 0; i < services.Len(); i++ {
		service := services.Get(i)
		p.renderService(service)

	}

	messages := file.Messages()
	for i := 0; i < messages.Len(); i++ {
		msg := messages.Get(i)
		p.renderMessage(msg)

	}
}
