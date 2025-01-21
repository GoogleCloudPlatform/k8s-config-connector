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

package outputsink_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/outputsink"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	"github.com/ghodss/yaml" //nolint:depguard
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestDirectorySinkShouldIgnoreTransmissionTerminator(t *testing.T) {
	ctx := context.TODO()

	tmpDir, cleanup := newTmpDir(t)
	defer cleanup()
	dirSink := outputsink.NewKRMYAMLDirectory(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), tmpDir)
	if err := dirSink.Receive(ctx, []byte("..."), nil); err != nil {
		t.Fatalf("unexpected error: got '%v', want 'nil", err)
	}
}

func TestDirectorySink(t *testing.T) {
	testCases := []struct {
		name            string
		sinkConstructor func(*schema.Provider, string) outputsink.OutputSink
		// the file that will be read, with results passed into the sink for testing
		testCaseFile string
		// the yaml file to use to pull unstructured schemas
		unstructuredFile   string
		expectedFileSuffix string
		expectedFilePath   string
	}{
		{
			sinkConstructor:  outputsink.NewKRMYAMLDirectory,
			testCaseFile:     "pubsubtopic.yaml",
			unstructuredFile: "pubsubtopic.yaml",
			expectedFilePath: "projects/my-project-id/PubSubTopic/pubsubtopic-bl5f09tftk5sk8gc0q4g.yaml",
		},
		{
			sinkConstructor:  outputsink.NewKRMYAMLDirectory,
			testCaseFile:     "storagebucket.yaml",
			unstructuredFile: "storagebucket.yaml",
			expectedFilePath: "projects/my-project-id/StorageBucket/US-WEST2/deleteoutofband-0ba21344-d250-11e8-bf9c-dc4a3e7de811.yaml",
		},
		{
			sinkConstructor:  outputsink.NewKRMYAMLDirectory,
			testCaseFile:     "name-with-unicode.yaml",
			unstructuredFile: "name-with-unicode.yaml",
			expectedFilePath: "projects/my-project-id/PubSubTopic/Hell_f6_20W_f6rld_20_445_435_43b_43b_43e_20_432_43e_440_43b_434_20the_20_60quick_20brown_20fox_20jumped_20over_20the_20lazy_20dog_AAAA_60_60_60.yaml",
		},
		{
			sinkConstructor:  outputsink.NewHCLDirectory,
			testCaseFile:     "computeinstance.tf",
			unstructuredFile: "computeinstance.yaml",
			expectedFilePath: "projects/my-project/ComputeInstance/us-central1-f/computetargetpool-dep4.tf",
		},
		{
			sinkConstructor:  outputsink.NewHCLDirectory,
			testCaseFile:     "containercluster.tf",
			unstructuredFile: "containercluster.yaml",
			expectedFilePath: "projects/my-project/ContainerCluster/us-central1-c/twenty-namespaces.tf",
		},
	}
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	for _, tc := range testCases {
		t.Run(tc.testCaseFile, func(t *testing.T) {
			tmpDir, cleanup := newTmpDir(t)
			defer cleanup()
			dirSink := tc.sinkConstructor(tfProvider, tmpDir)
			testDirectorySink(t, dirSink, tmpDir, tc.testCaseFile, tc.unstructuredFile, tc.expectedFilePath)
		})
	}

}

func TestDirectorySinkShouldHandleResourcesWithTheSameName(t *testing.T) {
	tmpDir, cleanup := newTmpDir(t)
	defer cleanup()
	dirSink := outputsink.NewKRMYAMLDirectory(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), tmpDir)
	testDirectorySink(t, dirSink, tmpDir, "pubsubtopic-project1.yaml", "pubsubtopic-project1.yaml", "projects/project1/PubSubTopic/pubsubtopic.yaml")
	testDirectorySink(t, dirSink, tmpDir, "pubsubtopic-project2.yaml", "pubsubtopic-project2.yaml", "projects/project2/PubSubTopic/pubsubtopic.yaml")
}

func findNewFiles(t *testing.T, dirName string, prevFiles []string) []string {
	files := findFilesRecursive(t, dirName)
	results := make([]string, 0)
	fileSet := make(map[string]bool)
	for _, f := range prevFiles {
		fileSet[f] = true
	}
	for _, f := range files {
		if fileSet[f] {
			continue
		}
		results = append(results, f)
	}
	return results
}

func testDirectorySink(t *testing.T, dirSink outputsink.OutputSink, outputDir, testCaseFile, unstructuredFile, expectedFilePath string) {
	ctx := context.TODO()

	unstructured := unstructuredFromYamlFile(t, unstructuredFile)
	initialFiles := findFilesRecursive(t, outputDir)
	expectedBytes := testFileToBytes(t, testCaseFile)
	if err := dirSink.Receive(ctx, expectedBytes, unstructured); err != nil {
		t.Fatalf("error receiving bytes; %v", err)
	}
	newFiles := findNewFiles(t, outputDir, initialFiles)
	if len(newFiles) != 1 {
		t.Fatalf("unexpected number of new files in '%v': got '%v', want '%v'", outputDir, len(newFiles), 1)
	}
	newFile := newFiles[0]
	if !strings.HasSuffix(newFile, expectedFilePath) {
		t.Errorf("new filename: got '%v', want suffix to be '%v'", newFile, expectedFilePath)
	}
	newBytes, err := ioutil.ReadFile(newFile)
	if err != nil {
		t.Fatalf("error reading file '%v': %v", newFile, err)
	}
	if !bytes.Equal(expectedBytes, newBytes) {
		diff := cmp.Diff(string(expectedBytes), string(newBytes))
		t.Fatalf("unexpected byte mismatch for '%v'. Diff:\n%v", newFile, diff)
	}
}

// findFilesRecursive returns an array of all files found
// in dirName, formatted as a path including dirName.
func findFilesRecursive(t *testing.T, dirName string) []string {
	results := make([]string, 0)
	err := filepath.Walk(dirName, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			t.Fatalf("error in walk: %v", err)
		}
		if info.IsDir() {
			return nil
		}
		results = append(results, filePath)
		return nil
	})
	if err != nil {
		t.Fatalf("error walking directory '%v': %v", dirName, err)
	}
	return results
}

func TestFileSink(t *testing.T) {
	ctx := context.TODO()

	f, cleanup := newTmpFile(t)
	defer cleanup()
	fileSink, err := outputsink.NewFile(f.Name())
	if err != nil {
		t.Fatalf("error creating file sink: %v", err)
	}
	// only has an affect when the function results in an error and the explicit call to close below does not occur
	defer fileSink.Close()
	pubSubFile := "pubsubtopic.yaml"
	bytes := testFileToBytes(t, pubSubFile)
	str := test.TrimLicenseHeaderFromYaml(string(bytes))
	unstructured := unstructuredFromYamlFile(t, pubSubFile)
	if err := fileSink.Receive(ctx, []byte(str), unstructured); err != nil {
		t.Fatalf("error receiving bytes: %v", err)
	}
	storageBucketFile := "storagebucket.yaml"
	bytes = testFileToBytes(t, storageBucketFile)
	str = test.TrimLicenseHeaderFromYaml(string(bytes))
	unstructured = unstructuredFromYamlFile(t, storageBucketFile)
	if err := fileSink.Receive(ctx, []byte(str), unstructured); err != nil {
		t.Fatalf("error receiving bytes: %v", err)
	}
	if err := fileSink.Close(); err != nil {
		t.Fatalf("error closing filesink: %v", err)
	}
	expectedBytes := testFileToBytes(t, "expected-combined.yaml")
	expectedStr := test.TrimLicenseHeaderFromYaml(string(expectedBytes))
	actualBytes := fileToBytes(t, f.Name())
	actualStr := string(actualBytes)
	if expectedStr != actualStr {
		diff := cmp.Diff(expectedStr, actualStr)
		t.Fatalf("file sink's output did not match what was expected, diff:\n%v", diff)
	}
}

func TestNewEmptyOutputParamShouldResultInStdoutWriter(t *testing.T) {
	testNewEmptyOutputParamShouldResultInStdoutWriter(t, outputsink.KRMResourceFormat)
	testNewEmptyOutputParamShouldResultInStdoutWriter(t, outputsink.HCLResourceFormat)
}

func testNewEmptyOutputParamShouldResultInStdoutWriter(t *testing.T, resourceFormat outputsink.ResourceFormat) {
	sink, err := outputsink.New(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), "", resourceFormat)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertTypeMatches(t, sink, outputsink.WriterSink{})
}

func TestNewWithExistingFileForOutputParam(t *testing.T) {
	testNewWithExistingFileForOutputParam(t, outputsink.KRMResourceFormat)
	testNewWithExistingFileForOutputParam(t, outputsink.HCLResourceFormat)
}

func testNewWithExistingFileForOutputParam(t *testing.T, resourceFormat outputsink.ResourceFormat) {
	tmpDir, cleanup := newTmpDir(t)
	defer cleanup()
	sink, err := outputsink.New(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), tmpDir, resourceFormat)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertTypeMatches(t, sink, newInstanceOfExpectedDirectorySinkType(t, resourceFormat))
}

func TestNewWithSlashAtEndOfOutputParam(t *testing.T) {
	testNewWithSlashAtEndOfOutputParam(t, outputsink.KRMResourceFormat)
	testNewWithSlashAtEndOfOutputParam(t, outputsink.HCLResourceFormat)
}

func testNewWithSlashAtEndOfOutputParam(t *testing.T, resourceFormat outputsink.ResourceFormat) {
	tmpDir, cleanup := newTmpDir(t)
	defer cleanup()
	output := fmt.Sprintf("%v%v", filepath.Join(tmpDir, "mypathwithaslash"), string(os.PathSeparator))
	sink, err := outputsink.New(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), output, resourceFormat)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertTypeMatches(t, sink, newInstanceOfExpectedDirectorySinkType(t, resourceFormat))
	fi, err := os.Stat(output)
	if err != nil {
		t.Fatalf("expected '%v' to exist instead err: %v", output, err)
	}
	if !fi.IsDir() {
		t.Fatalf("expected '%v' to be a directory", output)
	}
}

func TestNewWithNonExistentFileInExistingDirectoryOutputParam(t *testing.T) {
	testNewWithNonExistentFileInExistingDirectoryOutputParam(t, outputsink.KRMResourceFormat)
	testNewWithNonExistentFileInExistingDirectoryOutputParam(t, outputsink.HCLResourceFormat)
}

func testNewWithNonExistentFileInExistingDirectoryOutputParam(t *testing.T, resourceFormat outputsink.ResourceFormat) {
	tmpDir, cleanup := newTmpDir(t)
	defer cleanup()
	output := filepath.Join(tmpDir, "my-new-file")
	sink, err := outputsink.New(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), output, resourceFormat)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertTypeMatches(t, sink, outputsink.FileSink{})
}

func TestNewWithExistingFileOutputParam(t *testing.T) {
	testNewWithExistingFileOutputParam(t, outputsink.KRMResourceFormat)
	testNewWithExistingFileOutputParam(t, outputsink.HCLResourceFormat)
}

func testNewWithExistingFileOutputParam(t *testing.T, resourceFormat outputsink.ResourceFormat) {
	f, cleanup := newTmpFile(t)
	defer cleanup()
	sink, err := outputsink.New(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), f.Name(), resourceFormat)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertTypeMatches(t, sink, outputsink.FileSink{})
}

func TestNewWithInvalidDeviceFileShouldError(t *testing.T) {
	testNewWithInvalidDeviceFileShouldError(t, outputsink.KRMResourceFormat)
	testNewWithInvalidDeviceFileShouldError(t, outputsink.HCLResourceFormat)
}

func testNewWithInvalidDeviceFileShouldError(t *testing.T, resourceFormat outputsink.ResourceFormat) {
	sink, err := outputsink.New(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), "/dev/null", resourceFormat)
	if err == nil {
		t.Fatalf("expected an error, instead got 'nil'")
	}
	expectedMessage := "cannot use output parameter '/dev/null': is neither a 'regular' file or directory"
	if err.Error() != expectedMessage {
		t.Fatalf("unexpected error message\ngot:\n  %v\nwant:\n  %v", err.Error(), expectedMessage)
	}
	if sink != nil {
		t.Fatalf("expected no value for 'sink', instead got '%v'", sink)
	}
}

func TestParentPathIsFileShouldError(t *testing.T) {
	testParentPathIsFileShouldError(t, outputsink.KRMResourceFormat)
	testParentPathIsFileShouldError(t, outputsink.HCLResourceFormat)
}

func testParentPathIsFileShouldError(t *testing.T, resourceFormat outputsink.ResourceFormat) {
	f, cleanup := newTmpFile(t)
	defer cleanup()
	output := filepath.Join(f.Name(), "my-file")
	sink, err := outputsink.New(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), output, resourceFormat)
	if err == nil {
		t.Fatalf("expected an error, instead got 'nil'")
	}
	expectedMessage := fmt.Sprintf("cannot use output parameter '%v': parent path '%v' exists, but is not a directory",
		output, f.Name())
	if err.Error() != expectedMessage {
		t.Fatalf("unexpected error message\ngot:\n  %v\nwant:\n  %v", err.Error(), expectedMessage)
	}
	if sink != nil {
		t.Fatalf("expected no value for 'sink', instead got '%v'", sink)
	}
}

func TestMultipleNewDirectories(t *testing.T) {
	testMultipleNewDirectories(t, outputsink.KRMResourceFormat)
	testMultipleNewDirectories(t, outputsink.HCLResourceFormat)
}

func testMultipleNewDirectories(t *testing.T, resourceFormat outputsink.ResourceFormat) {
	tmpDir, cleanup := newTmpDir(t)
	defer cleanup()
	output := filepath.Join(tmpDir, "my", "new", "path")
	sink, err := outputsink.New(tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig()), output, resourceFormat)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertTypeMatches(t, sink, newInstanceOfExpectedDirectorySinkType(t, resourceFormat))
	fi, err := os.Stat(output)
	if err != nil {
		t.Fatalf("unexpected error statting '%v': %v", output, err)
	}
	if !fi.IsDir() {
		t.Fatalf("expected '%v' to be a directory", fi.Name())
	}
}

func assertTypeMatches(t *testing.T, sink outputsink.OutputSink, instanceOfExpectedType interface{}) {
	expectedType := reflect.TypeOf(instanceOfExpectedType)
	actualType := reflect.TypeOf(sink).Elem()
	if expectedType != actualType {
		t.Fatalf("unexpected type for sink: got '%v', want '%v'", actualType.Name(), expectedType.Name())
	}
}

func newTmpFile(t *testing.T) (*os.File, func()) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	cleanup := func() {
		deleteFile(t, f)
	}
	return f, cleanup
}

func deleteFile(t *testing.T, f *os.File) {
	if err := os.Remove(f.Name()); err != nil {
		t.Fatalf("error deleting '%v': %v", f.Name(), err)
	}
}

func newTmpDir(t *testing.T) (string, func()) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("error creating temp directory: %v", err)
	}
	cleanup := func() {
		deleteDir(t, tmpDir)
	}
	return tmpDir, cleanup
}

func deleteDir(t *testing.T, dir string) {
	if err := os.RemoveAll(dir); err != nil {
		t.Errorf("error deleting directory '%v': %v", dir, err)
	}
}

func testFileToBytes(t *testing.T, testFileName string) []byte {
	t.Helper()
	path := fmt.Sprintf("testdata/%v", testFileName)
	return fileToBytes(t, path)
}

func fileToBytes(t *testing.T, fileName string) []byte {
	t.Helper()
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatalf("error reading file '%v': %v", fileName, err)
	}
	return bytes
}

func newInstanceOfExpectedDirectorySinkType(t *testing.T, resourceFormat outputsink.ResourceFormat) interface{} {
	switch resourceFormat {
	case outputsink.KRMResourceFormat:
		return outputsink.KRMYAMLDirectorySink{}
	case outputsink.HCLResourceFormat:
		return outputsink.HCLDirectorySink{}
	default:
		t.Fatalf("unexpected resource format '%v'", resourceFormat)
		return nil
	}
}

// unstructuredFromYamlFile returns a unstructured.Unstructured
// struct from a yaml file.
func unstructuredFromYamlFile(t *testing.T, filePath string) *unstructured.Unstructured {
	pathWithPrefix := fmt.Sprintf("testdata/%v", filePath)
	bytes, err := ioutil.ReadFile(pathWithPrefix)
	if err != nil {
		t.Fatalf("error reading file %v: %v", filePath, err)
	}
	var value map[string]interface{}
	if err = yaml.Unmarshal(bytes, &value); err != nil {
		t.Fatalf("error marshalling bytes to yaml: %v", err)
	}
	return &unstructured.Unstructured{Object: value}
}
