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
		Short: "patches a message in a proto file",
		Long:  `Inserts the contents of stdin at the end of the message in the proto file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			insert, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("reading stdin: %w", err)
			}
			opt.Insertion = insert

			return RunProtoPatch(ctx, opt)
		},
	}
	cmd.Flags().StringVar(&opt.ProtoPath, "file", opt.ProtoPath, "path to proto file to patch")
	cmd.Flags().StringVar(&opt.Message, "message", opt.Message, "message file to patch")

	return cmd.Execute()
}

type ProtoPatchOptions struct {
	ProtoPath string
	Message   string
	Insertion []byte
}

func RunProtoPatch(ctx context.Context, opt ProtoPatchOptions) error {
	x := &insertPatchIntoMessage{
		Message:   opt.Message,
		Insertion: opt.Insertion,
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
		return fmt.Errorf("message %q not found in file %q", x.Message, protoPath)
	}
	if err := os.WriteFile(protoPath, x.Out, 0644); err != nil {
		return fmt.Errorf("writing to file %q: %w", protoPath, err)
	}

	return nil
}

type insertPatchIntoMessage struct {
	Source    []byte
	Insertion []byte
	Message   string
	Errors    []error

	Out []byte
}

func (x *insertPatchIntoMessage) VisitNode(depth int, cursor *sitter.TreeCursor) {
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
		x.VisitMessage(depth, cursor.CurrentNode())

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

func (x *insertPatchIntoMessage) VisitMessage(depth int, node *sitter.Node) {
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

	if messageName == x.Message {
		if messageBody == nil {
			x.Errors = append(x.Errors, fmt.Errorf("could not find message definition for message %q", messageName))
			return
		}

		var out bytes.Buffer
		out.Write(x.Source[:messageBody.StartByte()])
		messageBodyContents := string(x.Source[messageBody.StartByte():messageBody.EndByte()])
		messageBodyContents = strings.TrimSuffix(messageBodyContents, "}")
		out.WriteString(messageBodyContents)
		out.Write(x.Insertion)
		out.WriteString("\n}")
		out.Write(x.Source[messageBody.EndByte():])

		x.Out = out.Bytes()
	}
}
