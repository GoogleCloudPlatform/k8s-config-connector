// Copyright 2025 Google LLC
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

package runner

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const (
	examples = `
# Run commands in multiple branches.
conductor runner --branch-repo=/usr/local/google/home/wfender/go/src/github.com/cheftako/k8s-config-connector_branches
`
)

func BuildRunnerCmd() *cobra.Command {
	var opts RunnerOptions

	cmd := &cobra.Command{
		Use:     "runner",
		Short:   "runner Run commands in various branches.",
		Example: examples,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunRunner(cmd.Context(), &opts)
		},
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().StringVarP(&opts.branchConfFile, branchConfigurationFlag,
		"", "branches.yaml", "File containing the branch configurations.")
	cmd.Flags().StringVarP(&opts.branchRepoDir, branchRepoFlag,
		"", "", "Directory in which to do the work.")
	cmd.Flags().Int64VarP(&opts.command, commandFlag,
		"", 0, "Which commands to you on the directory.")

	return cmd
}

const (
	// flag names.
	branchConfigurationFlag = "branch-conf"
	branchRepoFlag          = "branch-repo"
	commandFlag             = "command"
)

type RunnerOptions struct {
	branchConfFile string
	branchRepoDir  string
	command        int64
}

func (opts *RunnerOptions) validateFlags() error {
	return nil
}

type Branches struct {
	Branches []Branch `yaml:"branches"`
}

type Proto struct {
	Service   string `yaml:"service"`   // ai
	Package   string `yaml:"package"`   // google.ai.generativelanguage.v1beta
	Proto     string `yaml:"proto"`     // Model
	Kind      string `yaml:"kind"`      // AIModel
	ProtoPath string `yaml:"protopath"` // google.ai.generativelanguage.v1beta.Model
	Validated string `yaml:"validated"` // UNUSED
}

type Branch struct {
	Name       string `yaml:"name"`       // ai-model
	Local      string `yaml:"local"`      // resource-ai-model
	Remote     string `yaml:"remote"`     // resource-ai-model
	Command    string `yaml:"command"`    // gcloud ai models
	Group      string `yaml:"group"`      // ai
	Resource   string `yaml:"resource"`   // model
	Controller string `yaml:"controller"` // Unknown

	Kind      string `yaml:"kind"`       // AIModel
	Package   string `yaml:"package"`    // google.ai.generativelanguage.v1beta
	Proto     string `yaml:"proto"`      // Model
	ProtoPath string `yaml:"proto-path"` // google.ai.generativelanguage.v1beta.Model

	Notes []string `yaml:"notes"` // Observation goes here
}

func RunRunner(ctx context.Context, opts *RunnerOptions) error {
	log.Printf("Running kompanion export with branch config: %s", opts.branchConfFile)

	if err := opts.validateFlags(); err != nil {
		return err
	}

	log.Printf("Starting Runner")

	file, err := os.Open(opts.branchConfFile) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("Read file %s: %v", opts.branchConfFile, string(data))

	var branches Branches

	err = yaml.Unmarshal(data, &branches)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	switch opts.command {
	case 0:
		printHelp()
	case 1:
		checkRepoDir(opts, branches)
	case 2:
		for idx, branch := range branches.Branches {
			/*
				if branch.Controller != "Unknown" {
					// Skipping TF, DCL and Direct controller resources.
					continue
				}
			*/
			log.Printf("%d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			createGithubBranch(opts, branch)
		}
	case 3:
		for idx, branch := range branches.Branches {
			log.Printf("%d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			deleteGithubBranch(opts, branch)
		}
	case 4:
		for idx, branch := range branches.Branches {
			log.Printf("%d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			createScriptYaml(opts, branch)
		}
	default:
		log.Fatalf("unrecognixed command: %d", opts.command)
	}
	return nil
}

func printHelp() {
	log.Println("conductor runner --branch-repo=? --branch-conf=<META> --command=<CMD>")
	log.Println("\t<CMD>")
	log.Println("\t0 - Print help")
	log.Println("\t1 - Check the repo director and metadata")
	log.Println("\t2 - Create the local github branches from the metadata")
	log.Println("\t3 - Delete the local github branches from the metadata")
	log.Println("\t4 - Create script.yaml for mock gcp generation in each github branch")
}

func checkRepoDir(opts *RunnerOptions, branches Branches) {
	stdin, stdout, exit, err := startBash()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()
	defer exit()

	cdRepoBranchDir(opts, "", stdin, stdout)

	log.Println("COMMAND: ls and echo")
	if _, err = stdin.Write([]byte("ls -alh && echo done\n")); err != nil {
		log.Fatal(err)
	}
	done := false
	outBuffer := make([]byte, 1000)
	var msg string
	for !done {
		length, err := stdout.Read(outBuffer)
		if err != nil {
			log.Fatal(err)
		}
		msg += string(outBuffer[:length])
		done = strings.HasSuffix(msg, "done\n")
	}
	log.Printf("LS OUT %s\r\n", msg)

	// Check for uniqueness constraints in the metadata.
	gcloudMap := make(map[string]string)
	nameMap := make(map[string]Branch)
	gitMap := make(map[string]string)
	grMap := make(map[string]Branch)
	for idx, branch := range branches.Branches {
		if branch.Command != "" {
			if existing, ok := gcloudMap[branch.Command]; ok {
				log.Printf("Command uniqueness constraint between %s and (%d)%s\r",
					existing, idx, branch.Name)
			}
			gitMap[branch.Command] = branch.Name
		}
		if existing, ok := nameMap[branch.Name]; ok {
			log.Printf("Name uniqueness constraint between %s at and %s\r",
				branch.Name, existing.Name)
		}
		nameMap[branch.Name] = branch

		if existing, ok := gitMap[branch.Local]; ok {
			log.Printf("Github uniqueness constraint between %s and (%d)%s\r",
				existing, idx, branch.Name)
		}
		gitMap[branch.Local] = branch.Name

		gr := branch.Group + ":" + branch.Resource
		if existing, ok := grMap[gr]; ok {
			log.Printf("Branch:Resource uniqueness constraint between %s and (%d)%s\r",
				existing.Name, idx, branch.Name)
		}
		grMap[gr] = branch
	}

	// Fix the data and write back
	/*
		var newBranches Branches
		// newBranches.Branches = make([]Branch, len(branches.Branches))
		for _, branch := range branches.Branches {
			branch.Name = strings.ToLower(branch.Name)
			branch.Local = strings.ToLower(branch.Local)
			branch.Remote = strings.ToLower(branch.Remote)
			branch.Group = strings.ToLower(branch.Group)
			branch.Resource = strings.ToLower(branch.Resource)
			branch.Notes = make([]string, 1)
			branch.Notes[0] = strings.TrimSpace(branch.Note)
			branch.Note = ""
			branch.Controller = "Unknown"
			newBranches.Branches = append(newBranches.Branches, branch)
		}
		data, err := yaml.Marshal(newBranches)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile("branches-new.yaml", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	*/

	/*
		// Fold in the PROTO file
		file, err := os.Open("proto-list.yaml") // For read access.
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Read file %s: context type %T", "proto-list.yaml", string(data))
		rawProtos := strings.Split(string(data), "---")
		var proto Proto
		var newBranches Branches
		var filledNames []string
		for _, rawProto := range rawProtos {
			err = yaml.Unmarshal([]byte(rawProto), &proto)
			if err != nil {
				log.Printf("Failed to parse %v\r", rawProto)
				log.Fatalf("error: %v", err)
			}
			kindGr := proto.Service + ":" + strings.ToLower(proto.Kind)
			protoGr := proto.Service + ":" + strings.ToLower(proto.Proto)
			kindName := proto.Service + "-" + strings.ToLower(proto.Kind)
			protoName := proto.Service + "-" + strings.ToLower(proto.Proto)
			if branch, ok := grMap[kindGr]; ok {
				branch.Kind = proto.Kind
				branch.Package = proto.Package
				branch.Proto = proto.Proto
				branch.ProtoPath = proto.ProtoPath
				newBranches.Branches = append(newBranches.Branches, branch)
				filledNames = append(filledNames, branch.Name)
			} else if branch, ok := grMap[protoGr]; ok {
				branch.Kind = proto.Kind
				branch.Package = proto.Package
				branch.Proto = proto.Proto
				branch.ProtoPath = proto.ProtoPath
				newBranches.Branches = append(newBranches.Branches, branch)
				filledNames = append(filledNames, branch.Name)
			} else if branch, ok := nameMap[kindName]; ok {
				branch.Kind = proto.Kind
				branch.Package = proto.Package
				branch.Proto = proto.Proto
				branch.ProtoPath = proto.ProtoPath
				newBranches.Branches = append(newBranches.Branches, branch)
				filledNames = append(filledNames, branch.Name)
			} else if branch, ok := nameMap[protoName]; ok {
				branch.Kind = proto.Kind
				branch.Package = proto.Package
				branch.Proto = proto.Proto
				branch.ProtoPath = proto.ProtoPath
				newBranches.Branches = append(newBranches.Branches, branch)
				filledNames = append(filledNames, branch.Name)
			} else {
				log.Printf("No match found for %s or %s\r", kindName, protoName)
				branch.Name = protoName
				branch.Local = "resource-" + protoName
				branch.Remote = "resource-" + protoName
				branch.Group = proto.Service
				branch.Resource = proto.Kind
				branch.Controller = "Unknown"
				branch.Notes = make([]string, 1)
				branch.Notes[0] = "No gcloud command found"
				branch.Kind = proto.Kind
				branch.Package = proto.Package
				branch.Proto = proto.Proto
				branch.ProtoPath = proto.ProtoPath
				newBranches.Branches = append(newBranches.Branches, branch)
				filledNames = append(filledNames, branch.Name)
			}
		}
		for _, branch := range branches.Branches {
			if slices.Contains(filledNames, branch.Name) {
				continue
			}
			newBranches.Branches = append(newBranches.Branches, branch)
			log.Printf("No additional metadata found for %s\r", branch.Name)
		}
		log.Printf("Marshalling %d branches\r", len(branches.Branches))
		data, err = yaml.Marshal(newBranches)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile("branches-new.yaml", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	*/
}
