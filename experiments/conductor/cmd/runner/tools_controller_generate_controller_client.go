package runner

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type GenerateControllerClientOptions struct {
	*ToolOptions

	ProtoSvc string

	CRDGroup string
}

func BuildGenerateControllerClientCommand(ToolOptions *ToolOptions) *cobra.Command {
	var opts GenerateControllerClientOptions
	opts.ToolOptions = ToolOptions
	opts.InitDefaults()

	cmd := &cobra.Command{
		Use:   "controller-generate-controller-client",
		Short: "Create a controller.",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := cmd.OutOrStdout()

			results, err := RunGenerateControllerClient(cmd.Context(), opts)
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "  %s\n", results.OutputPath)
			return nil
		},
	}

	// cmd.Flags().StringVar(&opts.ProtoMsg, "proto-message", opts.ProtoMsg, "The protobuf message we are generating a controler for")
	cmd.Flags().StringVar(&opts.ProtoSvc, "proto-service", opts.ProtoSvc, "The protobuf service we are generating a controller for")
	cmd.Flags().StringVar(&opts.CRDGroup, "crd-group", opts.CRDGroup, "The CRD group we are generating a controller for")
	// cmd.Flags().StringVar(&opts.Kind, "crd-kind", opts.Kind, "The CRD kind we are generating we are generating a controller for")
	// cmd.Flags().StringVar(&opts.CRDVersion, "crd-version", opts.CRDVersion, "The CRD version we are generating we are generating a controller for")

	return cmd
}

func (o *GenerateControllerClientOptions) InitDefaults() {
	// o.CRDVersion = "v1alpha1"
}

func (o *GenerateControllerClientOptions) DefaultAndValidate() error {
	if err := o.ToolOptions.DefaultAndValidate(); err != nil {
		return err
	}

	// Check if we have the required fields
	if o.ProtoSvc == "" {
		return fmt.Errorf("missing ProtoSvc field")
	}

	// if o.ProtoMsg == "" {
	// 	return fmt.Errorf("missing ProtoMsg field")
	// }

	if o.CRDGroup == "" {
		return fmt.Errorf("missing Group field")
	}

	// if o.Kind == "" {
	// 	return fmt.Errorf("missing Kind field")
	// }

	// if o.CRDVersion == "" {
	// 	return fmt.Errorf("missing CRDVersion field")
	// }

	return nil
}

// GenerateControllerClientResults holds the results from executing the GenerateControllerClient tool.
type GenerateControllerClientResults struct {
	OutputPath string

	// TODO: This feels like a layering violation
	ExecResults *ExecResults
}

func RunGenerateControllerClient(ctx context.Context, opts GenerateControllerClientOptions) (*GenerateControllerClientResults, error) {
	if err := opts.DefaultAndValidate(); err != nil {
		return nil, err
	}

	// Create the controller
	outputPath := filepath.Join(opts.RepoRoot, "pkg", "controller", "direct", shortNameForCRDGroup(opts.CRDGroup), "client.go")

	// Create the prompt for controllerbuilder
	prompt := fmt.Sprintf("// +tool:controller-client\n// proto.service: %s\n",
		opts.ProtoSvc)

	// Ensure the directory exists
	protoDir := filepath.Join(opts.RepoRoot, ".build", "third_party", "googleapis")
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", filepath.Dir(outputPath), err)
	}

	// Run controllerbuilder command
	cfg := CommandConfig{
		Name:    "Controller Builder",
		Cmd:     "controllerbuilder",
		Args:    []string{"prompt", "--src-dir", opts.RepoRoot, "--proto-dir", protoDir},
		WorkDir: opts.RepoRoot,
		Stdin:   strings.NewReader(prompt),
	}
	runnerOpts := &RunnerOptions{
		verbose: opts.Verbose,
	}
	execResults, err := executeCommand(runnerOpts, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to run controllerbuilder: %w", err)
	}

	// Write the controller file
	if err := os.WriteFile(outputPath, []byte(execResults.Stdout), 0644); err != nil {
		return nil, fmt.Errorf("failed to write client file %q: %w", outputPath, err)
	}

	return &GenerateControllerClientResults{
		OutputPath:  outputPath,
		ExecResults: &execResults,
	}, nil
}

// generateControllerClient creates a controller client for the branch
// This implements the logic from 01-create-controller-client.sh
func generateControllerClient(ctx context.Context, runnerOptions *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Generating controller client for branch %s", branch.Name)

	generateControllerClientOptions := GenerateControllerClientOptions{
		ToolOptions: &ToolOptions{
			RepoRoot:   runnerOptions.branchRepoDir,
			ScratchDir: filepath.Join(runnerOptions.loggingDir, branch.Name),
			Force:      runnerOptions.force,
		},
		ProtoSvc: branch.ProtoSvc,
		CRDGroup: branch.Group,
	}

	results, err := RunGenerateControllerClient(ctx, generateControllerClientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("generating controller client for branch %q: %w", branch.Name, err)
	}

	log.Printf("Successfully generated controller client for %s at %s", branch.Name, results.OutputPath)

	return []string{results.OutputPath}, results.ExecResults, nil
}
