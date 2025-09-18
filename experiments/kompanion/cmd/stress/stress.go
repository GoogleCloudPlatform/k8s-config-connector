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

package stress

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/stress"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/stress/iamchurn"
	"github.com/spf13/cobra"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	examples = `
	# run a stress test defined in an instructions file
	kompanion stress -f <path-to-instructions.yaml>
	`
)

type StressOptions struct {
	instructionsFile string
	runID            string
}

func BuildStressCmd() *cobra.Command {
	var opts StressOptions

	cmd := &cobra.Command{
		Use:     "stress",
		Short:   "stress test KCC",
		Example: examples,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunStress(cmd.Context(), &opts)
		},
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().StringVarP(&opts.instructionsFile, "file", "f", "", "path to the instructions.yaml file that defines the stress test parameters.")
	cmd.MarkFlagRequired("file")
	cmd.Flags().StringVar(&opts.runID, "run-id", "", "unique identifier for the stress test run. If not provided, a random ID will be generated.")

	return cmd
}

func RunStress(ctx context.Context, opts *StressOptions) error {
	config, err := stress.LoadConfig(opts.instructionsFile)
	if err != nil {
		return err
	}

	if opts.runID == "" {
		opts.runID = generateRunID()
	}

	log.Printf("Starting stress test with run ID: %s", opts.runID)

	restConfig, err := getRESTConfig(ctx)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}

	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return fmt.Errorf("error creating dynamic client: %w", err)
	}

	var stresser stress.Stresser
	switch config.TypeOfStress {
	case "IAMPolicyMemberChurn":
		stresser = iamchurn.New(opts.runID, config, dynamicClient)
	default:
		return fmt.Errorf("unsupported typeOfStress: %s", config.TypeOfStress)
	}

	if err := stresser.Setup(ctx); err != nil {
		return fmt.Errorf("stress test setup failed: %w", err)
	}

	duration, _ := time.ParseDuration(config.Duration)
	log.Printf("Stress test running for %s. Tearing down afterwards.", duration)

	select {
	case <-ctx.Done():
		log.Println("Context cancelled, starting teardown.")
	case <-time.After(duration):
		log.Println("Duration elapsed, starting teardown.")
	}

	if err := stresser.Teardown(ctx); err != nil {
		return fmt.Errorf("stress test teardown failed: %w", err)
	}

	log.Printf("Stress test with run ID %s completed.", opts.runID)
	return nil
}

func getRESTConfig(ctx context.Context) (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{},
	).ClientConfig()
}

func generateRunID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
