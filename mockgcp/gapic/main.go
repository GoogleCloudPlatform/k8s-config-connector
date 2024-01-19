package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/gapic/pkg/openapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/gapic/pkg/protogen"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	p := "dns-api.json"
	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("reading %q: %w", p, err)
	}
	doc := &openapi.Document{}
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(doc); err != nil {
		return fmt.Errorf("parsing json %q: %w", p, err)
	}

	c := protogen.NewOpenAPIConverter(doc)
	fileDescriptor, err := c.Convert(ctx)
	if err != nil {
		return fmt.Errorf("convert failed: %w", err)
	}

	filesDescriptorSet := &descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{fileDescriptor},
	}

	files, err := protodesc.NewFiles(filesDescriptorSet)
	if err != nil {
		return fmt.Errorf("protodesc.NewFiles failed: %w", err)
	}

	files.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		pw := protogen.NewProtoWriter(os.Stdout)
		pw.WriteFile(file)
		if err := pw.Error(); err != nil {
			klog.Fatalf("error rendering proto: %v", err)
		}

		return true
	})

	return nil
}

func PtrTo[T any](val T) *T {
	return &val
}
