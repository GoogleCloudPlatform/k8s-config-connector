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

package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/benchmark"
	kccio "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/io"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type Options struct {
	predictPath string
	BaseDir     string
}

func run(ctx context.Context) error {
	var o Options

	klog.InitFlags(nil)

	flag.StringVar(&o.predictPath, "predict-path", o.predictPath, "the path to the bash script where the prediction job is run.")
	flag.StringVar(&o.BaseDir, "base-dir", o.BaseDir, "base directory for the project code")

	flag.Parse()

	if o.predictPath == "" {
		return fmt.Errorf("predict-path is required")
	}

	totalCompareResults := &kccio.Results{}
	start := time.Now()

	for _, resource := range benchmark.GroundTruthResources {
		os.Setenv("KIND", resource.Kind)
		fmt.Println("Set env KIND: ", resource.Kind)
		os.Setenv("PROTO", resource.Proto)
		fmt.Println("Set env PROTO: ", resource.Proto)

		tmpFile, err := os.CreateTemp(filepath.Join(o.BaseDir, ".generate"), "benchmark-"+resource.Kind+"-")
		if err != nil {
			return err
		}
		os.Setenv("OUTPUT", tmpFile.Name())
		fmt.Println("Set env OUTPUT: ", tmpFile.Name())

		// Run the bash script to trigger prediction job.

		if _, err := exec.Command("/bin/sh", o.predictPath).CombinedOutput(); err != nil {
			return fmt.Errorf("run prediction script: %w", err)
		}

		tmpOut, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			return fmt.Errorf("read from predict file %s: %w", tmpFile.Name(), err)
		}
		predict := kccio.ExtractGoCode(tmpOut)

		// Get ground truth
		baseUrl, err := url.Parse("https://raw.githubusercontent.com/GoogleCloudPlatform/k8s-config-connector/refs/heads/master/apis/")
		if err != nil {
			return fmt.Errorf("url parse: %w", err)
		}
		rawUrl := baseUrl.ResolveReference(&url.URL{Path: resource.RawFile})
		fmt.Println("URL:", rawUrl.String())

		groundTruth, err := kccio.GetFromRemote(ctx, rawUrl.String())
		if err != nil {
			return fmt.Errorf("get ground truth from remote: %w", err)
		}

		// Compare ground truth and prediction results.
		results, err := kccio.CompareStruct(groundTruth, predict, resource.Kind+"Spec")
		if err != nil {
			return fmt.Errorf("compare struct: %w", err)
		}
		*totalCompareResults = append(*totalCompareResults, *results...)
		fmt.Println("Compared ", resource.Kind, " results: ", results.Accuracy())
	}

	elapsed := time.Since(start)

	fmt.Println("Elapsed time: ", elapsed)
	fmt.Println("Total estimate: ", totalCompareResults.Accuracy())

	return nil
}

// go run ./cmd/benchmark/type/ --predict-path ./tasks/update-type --base-dir $REPO_ROOT
