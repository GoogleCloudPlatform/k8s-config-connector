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

type GenerateControllerOptions struct {
	*ToolOptions

	ProtoSvc string
	ProtoMsg string

	CRDGroup   string
	CRDKind    string
	CRDVersion string
}

func BuildGenerateControllerCommand(ToolOptions *ToolOptions) *cobra.Command {
	var opts GenerateControllerOptions
	opts.ToolOptions = ToolOptions
	opts.InitDefaults()

	cmd := &cobra.Command{
		Use:   "controller-generate-controller",
		Short: "Create a controller.",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := cmd.OutOrStdout()

			results, err := RunGenerateController(cmd.Context(), opts)
			if err != nil {
				return err
			}
			fmt.Fprintf(out, "  %s\n", results.OutputPath)
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.ProtoMsg, "proto-message", opts.ProtoMsg, "The protobuf message we are generating a controler for")
	cmd.Flags().StringVar(&opts.ProtoSvc, "proto-service", opts.ProtoSvc, "The protobuf service we are generating a controller for")
	cmd.Flags().StringVar(&opts.CRDGroup, "crd-group", opts.CRDGroup, "The CRD group we are generating a controller for")
	cmd.Flags().StringVar(&opts.CRDKind, "crd-kind", opts.CRDKind, "The CRD kind we are generating we are generating a controller for")
	cmd.Flags().StringVar(&opts.CRDVersion, "crd-version", opts.CRDVersion, "The CRD version we are generating we are generating a controller for")

	return cmd
}

func (o *GenerateControllerOptions) InitDefaults() {
	o.CRDVersion = "v1alpha1"
}

func (o *GenerateControllerOptions) DefaultAndValidate() error {
	if err := o.ToolOptions.DefaultAndValidate(); err != nil {
		return err
	}

	// Check if we have the required fields
	if o.ProtoSvc == "" {
		return fmt.Errorf("missing ProtoSvc field")
	}

	if o.ProtoMsg == "" {
		return fmt.Errorf("missing ProtoMsg field")
	}

	if o.CRDGroup == "" {
		return fmt.Errorf("missing Group field")
	}

	if o.CRDKind == "" {
		return fmt.Errorf("missing Kind field")
	}

	if o.CRDVersion == "" {
		return fmt.Errorf("missing CRDVersion field")
	}

	return nil
}

// GenerateControllerResults holds the results from executing the GenerateController tool.
type GenerateControllerResults struct {
	OutputPath string

	// TODO: This feels like a layering violation
	ExecResults *ExecResults
}

func shortNameForCRDGroup(s string) string {
	shortName := s
	if ix := strings.Index(shortName, "."); ix != -1 {
		shortName = shortName[:ix]
	}
	return shortName
}

func shortNameForProtoService(s string) string {
	tokens := strings.Split(s, ".")
	if len(tokens) < 3 {
		return tokens[0]
	}
	return tokens[len(tokens)-3]
}

func shortNameForProtoMessage(s string) string {
	shortName := s
	if ix := strings.LastIndex(shortName, "."); ix != -1 {
		shortName = shortName[ix+1:]
	}
	return shortName
}

func RunGenerateController(ctx context.Context, opts GenerateControllerOptions) (*GenerateControllerResults, error) {
	if err := opts.DefaultAndValidate(); err != nil {
		return nil, err
	}

	// Create the controller
	controllerFileName := strings.ToLower(shortNameForProtoMessage(opts.ProtoMsg)) + "_controller.go"
	outputPath := filepath.Join(opts.RepoRoot, "pkg", "controller", "direct", shortNameForCRDGroup(opts.CRDGroup), controllerFileName)

	// Create the prompt for controllerbuilder
	prompt := fmt.Sprintf("// +tool:controller\n// proto.service: %s\n// proto.message: %s\n// crd.type: %s\n// crd.version: %s\n",
		opts.ProtoSvc, opts.ProtoMsg, opts.CRDKind, opts.CRDVersion) // Using v1alpha1 as default CRD version

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
		return nil, fmt.Errorf("failed to write controller file: %w", err)
	}

	return &GenerateControllerResults{
		OutputPath:  outputPath,
		ExecResults: &execResults,
	}, nil
}

// generateController creates a controller for the branch
// This implements the logic from 02-create-controller.sh
func generateController(ctx context.Context, runnerOptions *RunnerOptions, branch Branch, execResults *ExecResults) ([]string, *ExecResults, error) {
	log.Printf("Generating controller for branch %s", branch.Name)

	generateControllerOptions := GenerateControllerOptions{
		ToolOptions: &ToolOptions{
			RepoRoot:   runnerOptions.branchRepoDir,
			ScratchDir: filepath.Join(runnerOptions.loggingDir, branch.Name),
			Force:      runnerOptions.force,
		},
		ProtoSvc: branch.ProtoSvc,
		ProtoMsg: branch.ProtoMsg,
		CRDGroup: branch.Group,
		CRDKind:  branch.Kind,
	}

	results, err := RunGenerateController(ctx, generateControllerOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("generating controller for branch %q: %w", branch.Name, err)
	}
	log.Printf("Successfully generated controller for %s at %s", branch.Name, results.OutputPath)

	return []string{results.OutputPath}, results.ExecResults, nil
}
