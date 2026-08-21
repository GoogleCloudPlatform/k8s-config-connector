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

package utils

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
)

const (
	// this should be the same as KUSTOMIZE_VERSION from scripts/shared-vars-public.sh.
	requiredKustomizeVersion = "v3.5.4"
)

func DownloadAndExtractTarballAt(gcsPath, outputDir string) error {
	if err := DownloadObjectFromGCS(gcsPath, outputDir); err != nil {
		return fmt.Errorf("error downloading tarball at '%v': %w", gcsPath, err)
	}
	tarballName := path.Base(gcsPath)
	tarballPath := path.Join(outputDir, tarballName)
	if err := ExtractTarball(tarballPath, outputDir); err != nil {
		return fmt.Errorf("error extracting tarball: %w", err)
	}
	return nil
}

func DownloadObjectFromGCS(gcsPath, outputDir string) error {
	cmd := exec.Command("gcloud", "storage", "cp", gcsPath, outputDir)
	return Execute(cmd)
}

func ExtractTarball(tarballPath, outputDir string) error {
	cmd := exec.Command("tar", "-xvzf", tarballPath, "--directory", outputDir)
	return Execute(cmd)
}

// getKustomizeVersion returns the version of kustomize or an empty string and an error
// if there is an issue running the kustomize version command or parsing the command output.
func getKustomizeVersion() (string, error) {
	cmd := exec.Command("kustomize", "version", "--short")
	output, err := ExecuteAndCaptureOutput(cmd)
	if err != nil {
		return "", err
	}

	return parseKustomizeVersion(output)
}

// A parseable input should look like {kustomize/v3.5.4  2020-01-11T03:12:59Z  }.
func parseKustomizeVersion(input string) (string, error) {
	versionComponents := strings.Fields(input) // should split as [{kustomize/v3.5.4, 2020-01-11T03:12:59Z, }]
	if len(versionComponents) != 3 {
		return "", fmt.Errorf("unexpected output from kustomize version --short: %s", input)
	}

	return strings.Trim(versionComponents[0], "{kustomize/"), nil
}

func KustomizeBuild(path, output string) error {
	kustomizeVersion, err := getKustomizeVersion()
	if err != nil {
		return err
	}
	if requiredKustomizeVersion != kustomizeVersion {
		return fmt.Errorf("kustomize version mismatch; want: %s, got: %s", requiredKustomizeVersion, kustomizeVersion)
	}
	cmd := exec.Command("kustomize", "build", path, "--output", output)
	return Execute(cmd)
}

func Copy(source, dest string) error {
	cmd := exec.Command("cp", source, dest)
	return Execute(cmd)
}

func Execute(cmd *exec.Cmd) error {
	_, err := ExecuteAndCaptureOutput(cmd)
	return err
}

func ExecuteAndCaptureOutput(cmd *exec.Cmd) (stdout string, err error) {
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %v", err, string(out))
	}
	return string(out), nil
}
