// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/vertexai/genai"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot/ui"

	"k8s.io/klog"

	"context"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type Options struct {
	BaseDir       string
	Group         string
	Version       string
	ProtoResource string
	Feedback      string
}

func run(ctx context.Context) error {
	var o Options

	klog.InitFlags(nil)
	flag.StringVar(&o.BaseDir, "base-dir", o.BaseDir, "base directory for the project code")
	flag.StringVar(&o.Group, "group", o.Group, "the Config Connector service name")
	flag.StringVar(&o.Version, "version", o.Version, "the Config Connector version")
	flag.StringVar(&o.ProtoResource, "proto", o.ProtoResource, "the proto name")
	flag.StringVar(&o.Feedback, "feedback", o.Feedback, "any additional message to tell ai")

	flag.Parse()

	if o.BaseDir == "" {
		return fmt.Errorf("base-dir is required")
	}
	if o.Group == "" {
		return fmt.Errorf("group is required")
	}
	// Target files
	controllerPath := filepath.Join("./pkg/controller/direct/", o.Group, o.ProtoResource+"_controller.go")
	apiPath := filepath.Join("./apis", o.Group, o.Version, o.ProtoResource+"_types.go")
	identityPath := filepath.Join("./apis", o.Group, o.Version, o.ProtoResource+"_identity.go")
	referencePath := filepath.Join("./apis", o.Group, o.Version, o.ProtoResource+"_reference.go")

	// Selectively choose the files to pass to gemini.
	// It maybe better to pass in the entire dir, but it exceeds the token limits.
	includeDirs := []string{
		// "./apis/",
		// "./pkg/controller/direct/",

		"./apis/bigqueryconnection",
		"./apis/cloudbuild",
		"./apis/kms",
		"./pkg/controller/direct/common",
		"./pkg/controller/direct/bigqueryconnection",
		"./pkg/controller/direct/cloudbuild",
		"./pkg/controller/direct/kms",
		"./pkg/controller/direct/directbase",
		"./pkg/controller/direct/registry",
		controllerPath,
		apiPath,
		identityPath,
		referencePath,
	}
	contextFiles, err := readGoSourceFromSubdirs(o.BaseDir, includeDirs)
	if err != nil {
		return fmt.Errorf("read Go source directory %s: %w", o.BaseDir, err)
	}

	u := ui.NewTerminalUI()

	chat, err := codebot.NewChat(ctx, "", contextFiles, u)
	if err != nil {
		return err
	}
	defer chat.Close()

	var errs error

	// "I can't answer this question because the file path is not specific enough for me to understand what file you're referencing.  Please specify the full path to the file"
	if o.Feedback != "" {
		if err := chat.SendMessage(ctx,
			genai.Text("When write or fix the task, try to fix or improve by considering this"),
			genai.Text(o.Feedback),
		); err != nil {
			errs = errors.Join(errs, err)
		}
	}

	if err := chat.SendMessage(ctx,
		genai.Text("write or fix the given file, it should be a compilable Kubernetes controller, similar to the other *_controller.go I gave"),
		genai.Text(fmt.Sprintf("File %q:\n", controllerPath)),
	); err != nil {
		errs = errors.Join(errs, err)
	}
	if err := chat.SendMessage(ctx,
		genai.Text("write or fix the given file, it is kubernetes CRD in golang which is compatible with Kubebuilder"),
		genai.Text(fmt.Sprintf("File %q:\n", apiPath)),
	); err != nil {
		errs = errors.Join(errs, err)
	}

	if err := chat.SendMessage(ctx,
		genai.Text("write or fix the given go code file, whose identity matches the GCP URL path, similar to the other *_identity.go I gave"),
		genai.Text(fmt.Sprintf("File %q:\n", identityPath)),
	); err != nil {
		errs = errors.Join(errs, err)
	}
	if err := chat.SendMessage(ctx,
		genai.Text("write or fix the given go code file with correct NormalizeExternal function, similar to the other *_reference.go I gave"),
		genai.Text(fmt.Sprintf("File %q:\n", referencePath)),
	); err != nil {
		errs = errors.Join(errs, err)
	}
	return nil
}

func readGoSourceFromSubdirs(rootDir string, subdirs []string) (map[string]*codebot.FileInfo, error) {
	contextFiles := make(map[string]*codebot.FileInfo)

	for _, subdir := range subdirs {
		// Construct the full path to the subdirectory
		fullSubdirPath := filepath.Join(rootDir, subdir)

		// Walk through the subdirectory
		err := filepath.WalkDir(fullSubdirPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() && strings.HasSuffix(path, ".go") {
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				contextFiles[path] = &codebot.FileInfo{
					Path:    path,
					Content: string(content),
				}
			}
			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return contextFiles, nil
}
