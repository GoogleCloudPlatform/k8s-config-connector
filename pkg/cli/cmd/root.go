// Copyright 2022 Google LLC
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

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	golog "log"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/log"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"

	tfversion "github.com/hashicorp/terraform-provider-google-beta/version"
	"github.com/spf13/cobra"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

const (
	VerboseParamName = "verbose"

	commandName    = "config-connector"
	verboseDefault = false
)

var version = "dev"

var (
	description = fmt.Sprintf("The Config Connector CLI, version %v", version)
	rootCmd     = &cobra.Command{
		PersistentPreRunE: rootPreRunE,
		Use:               commandName,
		Short:             description,
		Long:              description,
	}
	verbose bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, VerboseParamName, "v", verboseDefault,
		"enable verbose logging (disabled by default)")
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(bulkExportCmd)
	rootCmd.AddCommand(printResourcesCmd)
	AddVersionCommand(rootCmd)
	AddLicensesCommand(rootCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(NewPreviewCmd())

	powertools.AddCommands(rootCmd)

	rootCmd.SilenceErrors = true
}

func rootPreRunE(_ *cobra.Command, _ []string) error {
	disableGoDefaultLogging()
	log.SetDefault(log.New(verbose))
	setTerraformUserAgent()
	return nil
}

func Execute() {
	if err := recoverExecute(); err != nil {
		log.Error("error in '%v' version '%v': %v", commandName, version, err)
		os.Exit(2)
	}
}

// This function exists to recover from any panics and return the informative error in 'err'
func recoverExecute() (err error) {
	defer execution.RecoverWithGenericError(&err)
	err = execute()
	// IMPORTANT: During a panic recover, the execution.RecoverWithGenericError(...) method will write the result of the
	// recovery into 'err'. For that reason, this explicit return statement must be here.
	return err
}

func execute() error {
	defaultToBulkExport(os.Args)
	return rootCmd.Execute()
}

type TestInvocationOptions struct {
	Stdout bytes.Buffer
	Stderr bytes.Buffer
	Stdin  bytes.Buffer
	Args   []string
}

// ExecuteFromTest allows for invocation of the CLI from a test
func ExecuteFromTest(options *TestInvocationOptions) error {
	rootCmd.SetIn(&options.Stdin)
	rootCmd.SetOut(&options.Stdout)
	rootCmd.SetErr(&options.Stderr)
	rootCmd.SetArgs(options.Args[1:])

	defaultToBulkExport(options.Args)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(&options.Stderr, "%v\n", err)
	}
	return err
}

func defaultToBulkExport(args []string) {
	// previously this command had no sub-commands and effectively defaulted to the bulk-export command for backwards
	// compatibility, if there is no sub-command and the flags appear to be the legacy flags format, default to the
	// bulk-export sub-command
	if isLegacyArgs(args) {
		newArgs := sanitizeArgsForBackwardsCompatibility(args[1:])
		newArgs = append([]string{bulkExportCommandName}, newArgs...)
		rootCmd.SetArgs(newArgs)
	}
}

func isLegacyArgs(args []string) bool {
	args = args[1:]
	if len(args) == 0 {
		// a valid legacy workload was the piping of a list of assets to stdin
		piped, err := parameters.IsInputPiped(os.Stdin)
		if err != nil {
			panic(fmt.Sprintf("unexpected error when verifying if input is piped: %v", err))
		}
		return piped
	}
	if args[0] == "help" {
		return false
	}
	cmd, _, err := rootCmd.Find(args)
	if err == nil && cmd != nil {
		return false
	}
	for i := 0; i < len(args); {
		switch args[i] {
		case "-v":
			fallthrough
		case "-h":
			return false
		case "-verbose":
			i++
		default:
			// if a flag starts with "--" then it is a 'new' workload
			if strings.HasPrefix(args[i], "--") {
				return false
			}
			i += 2
		}
	}
	return true
}

func sanitizeArgsForBackwardsCompatibility(args []string) []string {
	// the command previously used goflags which only requires a single dash in front of a flag
	for i := 0; i < len(args); {
		flag := args[i]
		switch flag {
		case "-verbose":
			args[i] = "--verbose"
			i++
		default:
			// if we encounter a flag add an additional '-'
			if strings.HasPrefix(flag, "-") && !strings.HasPrefix(flag, "--") {
				args[i] = fmt.Sprintf("-%v", flag)
			}
			i += 2
		}
	}
	return args
}

func setTerraformUserAgent() {
	tfversion.ProviderVersion = fmt.Sprintf("%v-%v", commandName, version)
}

func disableGoDefaultLogging() {
	// disable the go logger as the terraform library prints through that log which is undesired
	golog.SetOutput(ioutil.Discard)
}
