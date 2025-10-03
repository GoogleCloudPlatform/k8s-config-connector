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

package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"k8s.io/klog/v2"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/protobuf"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	var opt ProtoPatchOptions

	cmd := cobra.Command{
		Use:   "proto-patch",
		Short: "patches a proto file",
		Long:  `Patches the contents of stdin into a proto file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			patch, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("reading stdin: %w", err)
			}
			opt.Patch = patch

			return RunProtoPatch(ctx, opt)
		},
	}
	cmd.Flags().StringVar(&opt.ProtoPath, "file", opt.ProtoPath, "path to proto file to patch")
	cmd.Flags().VarP(
		enumflag.New(&opt.Mode, "mode", ProtoPatchModeIDs, enumflag.EnumCaseInsensitive),
		"mode", "m",
		"patch mode; can be 'append' or 'replace'")

	cmd.Flags().StringVar(&opt.Message, "message", opt.Message, "message name to patch")
	cmd.Flags().StringVar(&opt.Service, "service", opt.Service, "service name to patch")

	cmd.MarkFlagsOneRequired("message", "service")
	cmd.MarkFlagsMutuallyExclusive("message", "service")

	return cmd.Execute()
}

type ProtoPatchMode enumflag.Flag

const (
	ProtoPatchModeAppend ProtoPatchMode = iota
	ProtoPatchModeReplace
)

var ProtoPatchModeIDs = map[ProtoPatchMode][]string{
	ProtoPatchModeAppend:  {"append"},
	ProtoPatchModeReplace: {"replace"},
}

type ProtoPatchOptions struct {
	ProtoPath string
	Message   string
	Service   string
	Mode      ProtoPatchMode
	Patch     []byte
}

func RunProtoPatch(ctx context.Context, opt ProtoPatchOptions) error {
	x := &patchProto{
		Patch: opt.Patch,
		Mode:  opt.Mode,
	}

	if opt.Message != "" {
		x.Id = ProtoIdentifierMessage
		x.Name = opt.Message
	} else if opt.Service != "" {
		x.Id = ProtoIdentifierService
		x.Name = opt.Service
	}

	protoPath := opt.ProtoPath

	srcBytes, err := os.ReadFile(protoPath)
	if err != nil {
		return fmt.Errorf("reading file %q: %w", protoPath, err)
	}
	x.Source = srcBytes

	parser := sitter.NewParser()
	parser.SetLanguage(protobuf.GetLanguage())

	tree, err := parser.ParseCtx(ctx, nil, x.Source)
	if err != nil {
		return fmt.Errorf("parsing proto: %w", err)
	}
	defer tree.Close()

	cursor := sitter.NewTreeCursor(tree.RootNode())
	x.VisitNode(0, cursor)
	if len(x.Errors) != 0 {
		return errors.Join(x.Errors...)
	}

	if x.Out == nil {
		return fmt.Errorf("identifier %q not found in file %q", x.Name, protoPath)
	}
	if err := os.WriteFile(protoPath, x.Out, 0644); err != nil {
		return fmt.Errorf("writing to file %q: %w", protoPath, err)
	}

	return nil
}

type ProtoIdentifier string

const (
	ProtoIdentifierMessage ProtoIdentifier = "message"
	ProtoIdentifierService ProtoIdentifier = "service"
)

type patchProto struct {
	Mode ProtoPatchMode
	Id   ProtoIdentifier

	Source []byte
	Patch  []byte
	Name   string

	Errors []error

	Out []byte
}

func (x *patchProto) VisitNode(depth int, cursor *sitter.TreeCursor) {
	node := cursor.CurrentNode()

	// fmt.Printf("%s[%d:%s] %s\n", strings.Repeat("  ", depth), node.Symbol(), node.Type(), node.Content(x.Source))
	descend := true
	switch protobuf.GetLanguage().SymbolName(node.Symbol()) {
	case "source_file":
		// ignore
		descend = true
	case "comment":
		// e.g. // Hello world
		descend = false
	case "syntax":
		// e.g. syntax = "proto3";
		descend = false
	case "package":
		// e.g. package google.cloud.service.v1beta;
		descend = false
	case "import":
		// e.g. import "google/api/field_info.proto";
		descend = false
	case "option":
		// e.g. option go_package = "cloud.google.com/go/service/apiv1beta/servicepb;servicepb";
		descend = false
	case "enum":
		// e.g. enum MyEnum { ... }
		descend = false

	case "message":
		// e.g. message MyMessage { ... }
		descend = false
		if x.Id == ProtoIdentifierMessage {
			x.VisitMessage(depth, cursor.CurrentNode())
		}

	case "service":
		// e.g. service MyService { ... }
		descend = false
		if x.Id == ProtoIdentifierService {
			x.VisitService(depth, cursor.CurrentNode())
		}

	default:
		x.Errors = append(x.Errors, fmt.Errorf("unknown top-level node %q", protobuf.GetLanguage().SymbolName(node.Symbol())))
	}

	if descend {
		if cursor.GoToFirstChild() {
			x.VisitNode(depth+1, cursor)
			for cursor.GoToNextSibling() {
				x.VisitNode(depth+1, cursor)
			}
			cursor.GoToParent()
		}
	}
}

func (x *patchProto) VisitMessage(depth int, node *sitter.Node) {
	klog.V(2).Infof("%s[%d:%s] %s\n", strings.Repeat("  ", depth), node.Symbol(), node.Type(), node.Content(x.Source))

	messageName := ""
	var messageBody *sitter.Node
	childCount := int(node.ChildCount())
	for i := 0; i < childCount; i++ {
		child := node.Child(i)
		switch protobuf.GetLanguage().SymbolName(child.Symbol()) {
		case "message":
			// The 'message' literal!!
		case "message_name":
			messageName = child.Content(x.Source)
		case "message_body":
			messageBody = child
		default:
			x.Errors = append(x.Errors, fmt.Errorf("unknown message node %q", protobuf.GetLanguage().SymbolName(child.Symbol())))
		}
	}

	if messageName == x.Name {
		if messageBody == nil {
			x.Errors = append(x.Errors, fmt.Errorf("could not find message definition for message %q", messageName))
			return
		}

		var out bytes.Buffer
		out.Write(x.Source[:messageBody.StartByte()])
		if x.Mode == ProtoPatchModeAppend {
			messageBodyContents := string(x.Source[messageBody.StartByte():messageBody.EndByte()])
			messageBodyContents = strings.TrimSuffix(messageBodyContents, "}")
			out.WriteString(messageBodyContents)
		} else if x.Mode == ProtoPatchModeReplace {
			out.WriteString("{\n")
		}
		out.Write(x.Patch)
		out.WriteString("\n}")
		out.Write(x.Source[messageBody.EndByte():])

		x.Out = out.Bytes()
	}
}

func (x *patchProto) VisitService(depth int, node *sitter.Node) {
	klog.V(2).Infof("%s[%d:%s] %s\n", strings.Repeat("  ", depth), node.Symbol(), node.Type(), node.Content(x.Source))

	childCount := int(node.ChildCount())
	var serviceName *sitter.Node
	var serviceBodyIdx int
	for i := 0; i < childCount; i++ {
		child := node.Child(i)
		if protobuf.GetLanguage().SymbolName(child.Symbol()) == "service_name" {
			serviceName = child
			serviceBodyIdx = i + 1
			break
		}
	}

	if serviceName.Content(x.Source) == x.Name {
		serviceBodyStart := node.Child(serviceBodyIdx)
		lastChild := node.Child(childCount - 1)

		var out bytes.Buffer
		out.Write(x.Source[:serviceBodyStart.EndByte()])
		if x.Mode == ProtoPatchModeAppend {
			out.Write(x.Source[serviceBodyStart.EndByte():lastChild.StartByte()])
		}
		out.WriteString("\n")
		out.Write(x.Patch)
		out.WriteString(string(x.Source[lastChild.StartByte():lastChild.EndByte()]))
		out.Write(x.Source[lastChild.EndByte():])

		x.Out = out.Bytes()
	}
}
