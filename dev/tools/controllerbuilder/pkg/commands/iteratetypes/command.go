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

package iteratetypes

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"

	"github.com/spf13/cobra"
)

type ResourceMetadata struct {
	Service   string
	Package   string
	Proto     string
	Kind      string
	ProtoPath string
	Validated *bool
}

func (r *ResourceMetadata) String() string {
	return fmt.Sprintf("%s:%s:%s", r.Service, r.Package, r.Proto)
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(r)) + s[size:]
}

type Options struct {
	ProtoDir   string
	OutputFile string
}

func (o *Options) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	if o.ProtoDir == "" {
		o.ProtoDir = filepath.Join(root, ".build/third_party/googleapis/google")
	}
	if o.OutputFile == "" {
		o.OutputFile = root + "/all-proto.yaml"
	}
	return nil
}

func (o *Options) BindFlags(cmd *cobra.Command) {
	// TODO: Update this flag to accept a file path pointing to the ignored fields YAML file.
	cmd.Flags().StringVar(&o.ProtoDir, "proto-dir", o.ProtoDir, "base directory for proto API definitions")
	cmd.Flags().StringVar(&o.OutputFile, "output-file", o.OutputFile, "")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &Options{}
	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "iterate-types",
		Short: "iterate all protos",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			allproto := new([]ResourceMetadata)
			if err := RunIterate(ctx, opt, allproto); err != nil {
				return err
			}
			if err := writeToFile(ctx, opt, allproto); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

func writeToFile(ctx context.Context, o *Options, allproto *[]ResourceMetadata) error {
	log := klog.FromContext(ctx)
	if allproto == nil {
		log.Info("No proto files found")
		return nil
	}
	f, err := os.OpenFile(o.OutputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("opening file %q: %w", o.OutputFile, err)
	}
	for _, proto := range *allproto {
		b, err := yaml.Marshal(proto)
		if err != nil {
			return fmt.Errorf("marshalling proto %s: %w", proto, err)
		}
		f.WriteString(string(b) + "\n---\n")
	}
	return nil
}

func RunIterate(ctx context.Context, o *Options, allproto *[]ResourceMetadata) error {
	fmt.Println(o.ProtoDir)

	if err := filepath.WalkDir(o.ProtoDir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch filepath.Ext(p) {
		case ".proto":
			// OK
		default:
			return nil
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", p, err)
		}
		r := bytes.NewReader(b)
		br := bufio.NewReader(r)

		packageName := ""
		var protoCandidate *ResourceMetadata

		for {
			line, err := br.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("scanning file %q: %w", p, err)
			}
			line = strings.TrimSuffix(line, "\n")
			tokens := strings.Fields(line)

			if len(tokens) >= 2 && tokens[0] == "package" {
				packageName = strings.TrimSuffix(tokens[1], ";")
			}

			if len(tokens) >= 2 && tokens[0] == "message" {
				proto := tokens[1]
				messageName := packageName + "." + tokens[1]
				indent := 0
				protoCandidate = &ResourceMetadata{
					Package:   packageName,
					Service:   serviceFromPackage(packageName),
					Proto:     proto,
					ProtoPath: messageName}
				protoCandidate.Kind = capitalizeFirstLetter(protoCandidate.Service) + capitalizeFirstLetter(protoCandidate.Proto)
				for {
					for _, r := range line {
						if r == '{' {
							indent++
						}

						if r == '}' {
							indent--
						}
					}
					if indent == 0 {
						break
					}
					line, err = br.ReadString('\n')
					if strings.Contains(line, "(google.api.resource)") {
						*allproto = append(*allproto, *protoCandidate)
					}
					if err != nil {
						if err == io.EOF {
							break
						}
						return fmt.Errorf("scanning file %q: %w", p, err)
					}
					line = strings.TrimSuffix(line, "\n")
				}
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking directory tree: %w", err)
	}

	return nil
}

func serviceFromPackage(p string) string {
	tokens := strings.Split(p, ".")
	if len(tokens) == 0 {
		return ""
	}
	if len(tokens) <= 2 {
		fmt.Println("skip invalid package: ", p)
		return ""
	}
	version := tokens[len(tokens)-1]
	if strings.Contains(version, "alpha") {
		fmt.Println("skip alpha package: ", p)
		return ""
	}

	for _, t := range tokens {
		switch t {

		case "google":
			continue
		case "cloud":
			continue
		case "devtools":
			continue
		default:
			return t
		}
	}
	return ""
}
