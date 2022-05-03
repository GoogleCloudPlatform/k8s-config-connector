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
)

func DownloadAndExtractTarballAt(gcsPath, outputDir string) error {
	if err := DownloadObjectFromGCS(gcsPath, outputDir); err != nil {
		return fmt.Errorf("error downloading tarball at '%v': %v", gcsPath, err)
	}
	tarballName := path.Base(gcsPath)
	tarballPath := path.Join(outputDir, tarballName)
	if err := ExtractTarball(tarballPath, outputDir); err != nil {
		return fmt.Errorf("error extracting tarball: %v", err)
	}
	return nil
}

func DownloadObjectFromGCS(gcsPath, outputDir string) error {
	cmd := exec.Command("gsutil", "cp", gcsPath, outputDir)
	return Execute(cmd)
}

func ExtractTarball(tarballPath, outputDir string) error {
	cmd := exec.Command("tar", "-xvzf", tarballPath, "--directory", outputDir)
	return Execute(cmd)
}

func KustomizeBuild(path, output string) error {
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
		return "", fmt.Errorf("%v: %v", err, string(out))
	}
	return string(out), nil
}
