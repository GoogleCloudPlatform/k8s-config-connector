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
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const (
	examples = `
# Run commands in multiple branches.
conductor runner --branch-repo=/usr/local/google/home/wfender/go/src/github.com/cheftako/k8s-config-connector_branches
`

	// flag names.
	branchConfigurationFlag = "branch-conf"
	branchRepoFlag          = "branch-repo"
	commandFlag             = "command"
	loggingDirFlag          = "logging-dir"
	readFileTypeFlag        = "file-type"

	// Command values
	cmdHelp                = 0
	cmdCheckRepo           = 1
	cmdCreateGitBranch     = 2
	cmdDeleteGitBranch     = 3
	cmdEnableGCPAPIs       = 4
	cmdReadFiles           = 5
	cmdWriteFiles          = 6
	cmdCreateScriptYaml    = 10
	cmdCaptureHttpLog      = 11
	cmdGenerateMockGo      = 12
	cmdAddServiceRoundTrip = 13
	cmdAddProtoMakefile    = 14
	cmdRunMockTests        = 15
	cmdGenerateTypes       = 20
	cmdGenerateCRD         = 21
	cmdGenerateFuzzer      = 22

	typeScriptYaml = "scriptyaml"
	typeHttpLog    = "httplog"
	typeMockGo     = "mockgo"
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
	cmd.Flags().StringVarP(&opts.readFileType, readFileTypeFlag,
		"", "ScriptYaml", "Type of files to read, available values: 'ScriptYaml', 'HttpLog', 'MockGo'.")
	cmd.Flags().Int64VarP(&opts.command, commandFlag,
		"", 0, "Which commands to you on the directory.")
	cmd.Flags().StringVarP(&opts.loggingDir, loggingDirFlag,
		"", "", "dedicated directory for logging, empty for stdout.")
	cmd.Flags().DurationVarP(&opts.timeout, "timeout",
		"t", 5*time.Minute, "Global timeout for commands.")
	cmd.Flags().IntVarP(&opts.defaultRetries, "retries",
		"r", 10, "Default number of retries for failed commands.")

	return cmd
}

type RunnerOptions struct {
	branchConfFile string
	branchRepoDir  string
	command        int64
	loggingDir     string
	timeout        time.Duration
	readFileType   string
	defaultRetries int // Default number of retries for commands
}

func (opts *RunnerOptions) validateFlags() error {
	if opts.defaultRetries < 0 {
		return fmt.Errorf("retries flag cannot be negative, got %d", opts.defaultRetries)
	}
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
	Skip       bool   `yaml:"skip"`       // Skip this branch in processing

	Kind      string `yaml:"kind"`          // AIModel
	Package   string `yaml:"package"`       // google.ai.generativelanguage.v1beta
	Proto     string `yaml:"proto"`         // Model
	ProtoPath string `yaml:"proto-path"`    // google.ai.generativelanguage.v1beta.model_service
	ProtoSvc  string `yaml:"proto-service"` // google.ai.generativelanguage.v1beta.ModelService
	ProtoMsg  string `yaml:"proto-msg"`     // google.ai.generativelanguage.v1beta.Model
	HostName  string `yaml:"host-name"`     // generativelanguage.googleapis.com

	Notes []string `yaml:"notes"` // Observation goes here

	// GCP APIs that need to be enabled for the branch.
	// example:
	// - eventarc.googleapis.com
	// - aiplatform.googleapis.com
	ApisEnabled []string `yaml:"apis-enabled"`
}

// BranchModifier is a function type that takes a Branch and returns a modified Branch
type BranchModifier func(opts *RunnerOptions, branch Branch, workDir string) Branch

func RunRunner(ctx context.Context, opts *RunnerOptions) error {
	log.Printf("Running conductor runner with branch config: %s", opts.branchConfFile)

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
	// Only filter skipped branches if not running metadata commands
	if opts.command >= 0 {
		// Filter out skipped branches
		var filteredBranches []Branch
		for _, branch := range branches.Branches {
			if !branch.Skip {
				filteredBranches = append(filteredBranches, branch)
			}
		}
		branches.Branches = filteredBranches
	}

	switch opts.command {
	case -4:
		fixMetadata(opts, branches, "Skip(with reason)", setSkipOnBranchModifier)
	case -3:
		fixMetadata(opts, branches, "EnableAPIs", addEnableAPIsModifier)
	case -2:
		splitMetadata(opts, branches)
	case -1:
		fixMetadata(opts, branches, "ProtoPath", inferProtoPathModifier)
	case cmdHelp: // 0
		printHelp()
	case cmdCheckRepo: // 1
		checkRepoDir(opts, branches)
	case cmdCreateGitBranch: // 2
		for idx, branch := range branches.Branches {
			/*
				if branch.Controller != "Unknown" {
					// Skipping TF, DCL and Direct controller resources.
					continue
				}
			*/
			log.Printf("Create GitHub Branch: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			createGithubBranch(opts, branch)
		}
	case cmdDeleteGitBranch: // 3
		for idx, branch := range branches.Branches {
			log.Printf("Delete GitHub Branch: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			deleteGithubBranch(opts, branch)
		}
	case cmdEnableGCPAPIs: // 4
		for idx, branch := range branches.Branches {
			log.Printf("Enable GCP APIs: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			enableAPIs(opts, branch)
		}
	case cmdReadFiles: // 5
		readFuncs := map[string]func(*RunnerOptions, Branch){
			typeScriptYaml: readScriptYaml,
			typeHttpLog:    readHttpLog,
			typeMockGo:     readMockGo,
		}

		if readFunc, ok := readFuncs[strings.ToLower(opts.readFileType)]; ok {
			for idx, branch := range branches.Branches {
				log.Printf("Read %s: %d name: %s, branch: %s\r\n", opts.readFileType, idx, branch.Name, branch.Local)
				readFunc(opts, branch)
			}
		}

	case cmdWriteFiles: // 6
		writeFuncs := map[string]func(*RunnerOptions, Branch){
			typeScriptYaml: writeScriptYaml,
		}

		if writeFunc, ok := writeFuncs[strings.ToLower(opts.readFileType)]; ok {
			for idx, branch := range branches.Branches {
				log.Printf("Write %s: %d name: %s, branch: %s\r\n", opts.readFileType, idx, branch.Name, branch.Local)
				writeFunc(opts, branch)
			}
		}

	case cmdCreateScriptYaml: // 10
		for idx, branch := range branches.Branches {
			log.Printf("Create Script YAML: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			createScriptYaml(opts, branch)
		}
	case cmdCaptureHttpLog: // 11
		for idx, branch := range branches.Branches {
			log.Printf("Capture HTTP Log: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			captureHttpLog(opts, branch)
		}
	case cmdGenerateMockGo: // 12
		for idx, branch := range branches.Branches {
			log.Printf("Generate mock Service and Resource go files: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			generateMockGo(opts, branch)
		}
	case cmdAddServiceRoundTrip: // 13
		for idx, branch := range branches.Branches {
			log.Printf("Add service to mock_http_roundtrip.go: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			addServiceToRoundTrip(opts, branch)
		}
	case cmdAddProtoMakefile: // 14
		for idx, branch := range branches.Branches {
			log.Printf("Add proto to makefile: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			addProtoToMakefile(opts, branch)
		}
	case cmdRunMockTests: // 15
		for idx, branch := range branches.Branches {
			log.Printf("Run mockgcptests on generated mocks: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			runMockgcpTests(opts, branch)
		}
	case cmdGenerateTypes: // 20
		for idx, branch := range branches.Branches {
			log.Printf("Generate Types and Mapper: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			generateTypesAndMapper(opts, branch)
		}
	case cmdGenerateCRD: // 21
		for idx, branch := range branches.Branches {
			log.Printf("Generate CRD: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			generateCRD(opts, branch)
			//generateSpecStatus(opts, branch)
		}
	case cmdGenerateFuzzer: // 22
		for idx, branch := range branches.Branches {
			log.Printf("Generate Fuzzer: %d name: %s, branch: %s\r\n", idx, branch.Name, branch.Local)
			generateFuzzer(opts, branch)
		}
	default:
		log.Fatalf("unrecognized command: %d", opts.command)
	}
	return nil
}

func printHelp() {
	log.Println("conductor runner --branch-repo=? --branch-conf=<META> --command=<CMD>")
	log.Println("\t<CMD>")
	log.Println("\t0 - Print help")
	log.Println("\t1 - [Validate] Repo directory and metadata")
	log.Println("\t2 - [Branch] Create the local github branches from the metadata")
	log.Println("\t3 - [Branch] Delete the local github branches from the metadata")
	log.Println("\t4 - [Project] Enable GCP APIs for each branch")
	log.Println("\t5 - [Generated] Read the specific type of generated files in each github branch")
	log.Println("\t6 - [Generated] Write the specific type of files from all_scripts.yaml to each github branch")
	log.Println("\t10 - [Mock] Create script.yaml for mock gcp generation in each github branch")
	log.Println("\t11 - [Mock] Create _http.log for mock gcp generation in each github branch")
	log.Println("\t12 - [Mock] Generate mock Service and Resource go files in each github branch")
	log.Println("\t13 - [Mock] Add service to mock_http_roundtrip.go in each github branch")
	log.Println("\t14 - [Mock] Add proto to makefile in each github branch")
	log.Println("\t15 - [Mock] Run mockgcptests on generated mocks in each github branch")
	log.Println("\t20 - [CRD] Generate Types and Mapper for each branch")
	log.Println("\t21 - [CRD] Generate CRD for each branch")
	log.Println("\t22 - [Fuzzer] Generate fuzzer for each branch")
}

func checkRepoDir(opts *RunnerOptions, branches Branches) {
	stdin, stdout, exit, err := startBash()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()
	defer exit()

	cdRepoBranchDirBash(opts, "", stdin, stdout)

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

const COPYRIGHT_HEADER string = `# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

`

func fixMetadata(opts *RunnerOptions, branches Branches, what string, modifier BranchModifier) {
	workDir := filepath.Join(opts.branchRepoDir, ".build", "third_party", "googleapis")
	var newBranches Branches
	for _, branch := range branches.Branches {
		log.Printf("Fixing Metadata: %s for branch: %s", what, branch.Name)
		newBranch := modifier(opts, branch, workDir)
		newBranches.Branches = append(newBranches.Branches, newBranch)
	}
	data := []byte(COPYRIGHT_HEADER)
	yamlData, err := yaml.Marshal(newBranches)
	if err != nil {
		log.Fatal(err)
	}
	data = append(data, yamlData...)
	err = os.WriteFile("branches-new.yaml", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func inferProtoPathModifier(opts *RunnerOptions, branch Branch, workDir string) Branch {
	if branch.ProtoPath == "" {
		branch.ProtoPath = inferProtoPath(opts, branch, workDir)
		log.Printf("ProtoPath for %s should be %s", branch.Name, branch.ProtoPath)
	}
	return branch
}

func addEnableAPIsModifier(opts *RunnerOptions, branch Branch, workDir string) Branch {
	close := setLoggingWriter(opts, branch)
	defer close()
	if branch.Command == "" {
		return branch
	}
	if branch.ApisEnabled != nil && len(branch.ApisEnabled) > 0 {
		return branch
	}

	cfg := CommandConfig{
		Name:    "API Discovery",
		Cmd:     "codebot",
		Args:    []string{"--prompt=/dev/stdin"},
		WorkDir: workDir,
		Stdin: strings.NewReader(fmt.Sprintf(`Given the gcloud command %q, what Google Cloud APIs need to be enabled to use this command?
Try inferring the apis required from the command.
If needed use gcloud command to instrospect the error message to get the API name.
Please respond with a list of API service names in the format needed for 'gcloud services enable'.
For example, if a command needs the Cloud Storage API, respond with 'storage.googleapis.com'.
If you are using gcloud to create something, make sure you are deleting it after you are done.
Print the list of APIs, one on each line in this format api-required: <api-name>
Only include APIs that are directly needed by this command.
`, branch.Command)),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}

	output, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Printf("Failed to get APIs for %s: %v", branch.Name, err)
		return branch
	}

	// Parse the codebot output to get API list
	var filteredLines []string
	for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "api-required:") {
			filteredLines = append(filteredLines, line)
		}
	}

	// split the structred output into a list of APIs
	var cleanedAPIs []string
	for _, api := range filteredLines {
		api = strings.TrimSpace(api)
		api = strings.TrimPrefix(api, "api-required:")
		api = strings.TrimSpace(api)
		if api != "" {
			cleanedAPIs = append(cleanedAPIs, api)
		}
	}

	branch.ApisEnabled = cleanedAPIs
	return branch
}

func inferProtoPath(opts *RunnerOptions, branch Branch, workDir string) string {
	var protoDir = ""
	var svcNm = ""
	if branch.Proto != "" {
		svcNm = branch.Proto
	} else if branch.ProtoSvc != "" {
		dirs := strings.Split(branch.ProtoSvc, ".")
		svcNm = dirs[len(dirs)-1]
	} else if branch.Resource != "" {
		svcNm = branch.Resource
	} else {
		return ""
	}

	if branch.ProtoPath != "" {
		dirs := strings.Split(branch.ProtoPath, ".")
		protoDir = filepath.Join(dirs[:len(dirs)-1]...)
	} else if branch.ProtoSvc != "" {
		dirs := strings.Split(branch.ProtoSvc, ".")
		protoDir = filepath.Join(dirs[:len(dirs)-1]...)
	} else if branch.Package != "" {
		dirs := strings.Split(branch.Package, ".")
		protoDir = filepath.Join(dirs...)
	} else {
		return ""
	}

	searchPath := filepath.Join(workDir, protoDir, "*.proto")
	files, err := filepath.Glob(searchPath)
	if err != nil {
		log.Printf("Glob error %v", err)
		return ""
	}
	if len(files) == 1 {
		localFile, _ := strings.CutPrefix(files[0], workDir+string(filepath.Separator))
		log.Printf("Glob for %s matched %s", branch.Name, localFile)
		return localFile
	}

	searchList := ""
	first := true
	args := []string{"-iH", fmt.Sprintf("^service %s", svcNm)}
	for _, file := range files {
		localFile, _ := strings.CutPrefix(file, workDir+string(filepath.Separator))
		args = append(args, localFile)
		if first {
			searchList = localFile
			first = false
		} else {
			searchList += " " + localFile
		}
	}

	cfg := CommandConfig{
		Name:       "Search for service",
		Cmd:        "egrep",
		Args:       args,
		WorkDir:    workDir,
		MaxRetries: 2,
	}
	output, errOutput, err := executeCommand(opts, cfg)
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			args := []string{"-iH", "^service"}
			for _, file := range files {
				localFile, _ := strings.CutPrefix(file, workDir+string(filepath.Separator))
				args = append(args, localFile)
				if first {
					searchList = localFile
					first = false
				} else {
					searchList += " " + localFile
				}
			}

			cfg = CommandConfig{
				Name:       "Search for any service",
				Cmd:        "egrep",
				Args:       args,
				WorkDir:    workDir,
				MaxRetries: 2,
			}
			output, errOutput, err = executeCommand(opts, cfg)
			if err != nil {
				log.Printf("Working in directory %s", workDir)
				log.Printf("Got response2 %v", errOutput)
				log.Printf("Find proto file error: %q\n", err)
				return ""
			}
		} else {
			log.Printf("Working in directory %s", workDir)
			log.Printf("Find proto file error: %q\n", err)
			return ""
		}
	}

	vals := strings.Split(output, ":")
	if len(vals) <= 1 {
		log.Printf("ERROR: something wrong with grep response: %q\n", output)
		return ""
	}
	return vals[0]
}

// Trying to put all resources with the same group together while
// keeping the buckets roughly the same size.
func splitMetadata(opts *RunnerOptions, branches Branches) {
	var newBranches [7]Branches
	groupSplitMap := make(map[string]int)
	for _, branch := range branches.Branches {
		if branch.Command == "" {
			continue
		}
		if split, present := groupSplitMap[branch.Group]; present {
			newBranches[split].Branches = append(newBranches[split].Branches, branch)
			continue
		}
		smallest := len(newBranches[0].Branches)
		bucket := 0
		for cntr := 1; cntr < 7; cntr++ {
			if len(newBranches[cntr].Branches) < smallest {
				smallest = len(newBranches[cntr].Branches)
				bucket = cntr
			}
		}
		newBranches[bucket].Branches = append(newBranches[bucket].Branches, branch)
		groupSplitMap[branch.Group] = bucket
	}
	// Hard coding splitting into 7 files
	for cntr := 0; cntr < 7; cntr++ {
		data := []byte(COPYRIGHT_HEADER)
		yamlData, err := yaml.Marshal(newBranches[cntr])
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, yamlData...)
		err = os.WriteFile(fmt.Sprintf("branches-%d.yaml", cntr), data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func setSkipOnBranchModifier(opts *RunnerOptions, branch Branch, workDir string) Branch {
	close := setLoggingWriter(opts, branch)
	defer close()

	if branch.Command == "" {
		branch.Skip = true
		branch.Notes = append(branch.Notes, "no gcloud command")
	}

	// Run gcloud xxx --help to see if it implements CRUD methods
	gcloudCommand := branch.Command
	gcloudCommand = strings.TrimPrefix(gcloudCommand, "gcloud")
	args := strings.Fields(gcloudCommand)

	// Skip if gcloud command has no args
	if len(args) == 0 {
		branch.Skip = true
		branch.Notes = append(branch.Notes, "invalid gcloud command")
	}

	// Keep if any of CRUD is supported
	// args = append(args, "--help")
	cfg := CommandConfig{
		Name:       "Check CRUD Support",
		Cmd:        "gcloud",
		Args:       args,
		WorkDir:    workDir,
		MaxRetries: 1,
	}
	_, errOutput, _ := executeCommand(opts, cfg)
	// todo: shall we keep when gcloud xxx --help returns error?
	// Just keep everything for now
	// todo: Do we need to include 'patch'? I thought 'gcloud' primarily uses 'update' for these operations.
	var operationRegex = regexp.MustCompile(`(create|update|delete)`)
	if !operationRegex.MatchString(errOutput) {
		branch.Skip = true
		branch.Notes = append(branch.Notes, "gcloud command does not contain any of create/update/delete")
	}
	return branch
}
