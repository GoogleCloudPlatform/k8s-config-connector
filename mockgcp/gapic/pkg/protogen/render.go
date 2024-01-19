package protogen

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/klog/v2"
)

type ProtoWriter struct {
	w            io.Writer
	errors       []error
	protoVersion int
}

func NewProtoWriter(w io.Writer) *ProtoWriter {
	return &ProtoWriter{
		w:            w,
		protoVersion: 3,
	}
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

func (p *ProtoWriter) renderField(fd protoreflect.FieldDescriptor) {
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
		b.WriteString(",")
		switch fd.MapValue().Kind() {
		case protoreflect.StringKind:
			b.WriteString("string")
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

		case protoreflect.BoolKind:
			b.WriteString("bool ")

		case protoreflect.Uint32Kind:
			b.WriteString("uint32 ")

		case protoreflect.Int32Kind:
			b.WriteString("int32 ")

		case protoreflect.DoubleKind:
			b.WriteString("double ")

		case protoreflect.MessageKind:
			b.WriteString(string(fd.Message().Name()) + " ")

		default:
			p.errorf("unhandled field Kind %v", fd.Kind())
		}
	}

	fmt.Fprintf(&b, "%s = %d", fd.Name(), fd.Number())
	if fd.JSONName() != "" {
		fmt.Fprintf(&b, " [json_name=%q]", fd.JSONName())
	}
	fmt.Fprintf(&b, ";\n")
	p.writeString(b.String())

}

func (p *ProtoWriter) renderMessage(msg protoreflect.MessageDescriptor) {
	p.printf("message %s {\n", msg.Name())
	fields := msg.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		p.renderField(field)
	}

	p.printf("}\n")
}

func (p *ProtoWriter) renderMethod(md protoreflect.MethodDescriptor) {
	var b bytes.Buffer
	b.WriteString("  rpc ")

	b.WriteString(string(md.Name()))
	b.WriteString("(")
	b.WriteString(string(md.Input().Name()))
	b.WriteString(")")
	b.WriteString(" returns ")
	b.WriteString("(")
	b.WriteString(string(md.Output().Name()))
	b.WriteString(")")
	b.WriteString(";\n")
	p.writeString(b.String())

}

func (p *ProtoWriter) renderService(msg protoreflect.ServiceDescriptor) {
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
	p.printf("syntax = \"proto3\";\n")
	p.printf("package %s;\n", file.Package())
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
