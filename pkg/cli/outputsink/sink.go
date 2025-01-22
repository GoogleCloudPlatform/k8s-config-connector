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

package outputsink

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/outputsink/filename"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type ResourceFormat string

const (
	KRMResourceFormat = "krm"
	HCLResourceFormat = "hcl"
)

type OutputSink interface {
	io.Closer
	Receive(ctx context.Context, bytes []byte, unstructured *unstructured.Unstructured) error
}

type WriterSink struct {
	w io.Writer
}

// NewWriter returns a WriterSink which can be used to write to any io.Writer, all received bytes
// will be written to the io.Writer
func NewWriter(w io.Writer) *WriterSink {
	sink := WriterSink{
		w: w,
	}
	return &sink
}

func (ws *WriterSink) Receive(_ context.Context, bytes []byte, _ *unstructured.Unstructured) error {
	if _, err := ws.w.Write(bytes); err != nil {
		return fmt.Errorf("error writing bytes: %w", err)
	}
	return nil
}

func (ws *WriterSink) Close() error {
	return nil
}

type FileSink struct {
	file   *os.File
	writer *bufio.Writer
}

// NewFile returns a FileSink which will write all received bytes to a single file
func NewFile(filePath string) (*FileSink, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%v': %w", filePath, err)
	}
	sink := FileSink{
		file:   file,
		writer: bufio.NewWriter(file),
	}
	return &sink, nil
}

func (fs *FileSink) Receive(_ context.Context, bytes []byte, _ *unstructured.Unstructured) error {
	// bufio.Writer either writes all the bytes or returns an error so we can ignore the first 'nn' return value
	_, err := fs.writer.Write(bytes)
	if err != nil {
		return fmt.Errorf("error writing to '%v': %w", fs.file.Name(), err)
	}
	return nil
}

func (fs *FileSink) Close() error {
	if fs.file == nil {
		return nil
	}
	if err := fs.writer.Flush(); err != nil {
		return fmt.Errorf("error flushing buffered writes for '%v': %w", fs.file.Name(), err)
	}
	if err := fs.file.Close(); err != nil {
		return fmt.Errorf("error closing '%v': %w", fs.file.Name(), err)
	}
	fs.file = nil
	fs.writer = nil
	return nil
}

// This internal type exists to share some code between the HCL and KRM output streams, namely the KRM output streams
// get a somewhat useful filename with the resource name in it, while the HCL streams do not get this
type DirectorySink struct {
	dir        string
	tfProvider *schema.Provider
}

func (ds *DirectorySink) Close() error {
	return nil
}

func (ds *DirectorySink) receive(ctx context.Context, bytes []byte, unstructured *unstructured.Unstructured, suffix string) error {
	// TODO: need to pass in smloader to the constructor
	smLoader, _ := servicemappingloader.New()
	fileName, err := filename.Get(ctx, unstructured, smLoader, ds.tfProvider)
	if err != nil {
		return fmt.Errorf("error choosing a filename: %w", err)
	}
	slugPath := slugifyPath(fileName)
	filePath := filepath.Join(ds.dir, fmt.Sprintf("%v.%v", slugPath, suffix))
	// ensure the parent directory exists
	dir := path.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error ensuring parent path '%v' exists: %w", dir, err)
	}
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return fmt.Errorf("error writing bytes to '%v': %w", filePath, err)
	}
	return nil
}

func slugifyPath(filePath string) string {
	splits := strings.Split(filePath, string(os.PathSeparator))
	for i, str := range splits {
		splits[i] = filename.MakeSafeFilename(str)
	}
	return filepath.Join(splits...)
}

type HCLDirectorySink struct {
	DirectorySink
}

func NewHCLDirectory(tfProvider *schema.Provider, filepath string) OutputSink {
	sink := HCLDirectorySink{
		DirectorySink: DirectorySink{
			dir:        filepath,
			tfProvider: tfProvider,
		},
	}
	return &sink
}

func (ds *HCLDirectorySink) Receive(ctx context.Context, bytes []byte, unstructured *unstructured.Unstructured) error {
	return ds.receive(ctx, bytes, unstructured, "tf")
}

type KRMYAMLDirectorySink struct {
	DirectorySink
}

// NewKRMYAMLDirectory returns a KRMYAMLDirectorySink. This sink assumes the given filepath parameter is
// a directory. For each received bytes it will yaml.Unmarshal the bytes to map[string]interface{} and
// then use the 'metadata.name' property and write out an associated YAML file using the name.
func NewKRMYAMLDirectory(tfProvider *schema.Provider, filePath string) OutputSink {
	sink := KRMYAMLDirectorySink{
		DirectorySink: DirectorySink{
			dir:        filePath,
			tfProvider: tfProvider,
		},
	}
	return &sink
}

func (ds *KRMYAMLDirectorySink) Close() error {
	return ds.DirectorySink.Close()
}

func (ds *KRMYAMLDirectorySink) Receive(ctx context.Context, bytes []byte, unstructured *unstructured.Unstructured) error {
	if isYAMLTerminator(bytes) {
		return nil
	}
	return ds.DirectorySink.receive(ctx, bytes, unstructured, "yaml")
}

func isYAMLTerminator(bytes []byte) bool {
	if len(bytes) != 3 {
		return false
	}
	return "..." == string(bytes)
}

// New returns a new output sink appropriate for the given resource format.
func New(tfProvider *schema.Provider, outputParam string, resourceFormat ResourceFormat) (OutputSink, error) {
	switch resourceFormat {
	case KRMResourceFormat:
		return newKRM(tfProvider, outputParam)
	case HCLResourceFormat:
		return newHCL(tfProvider, outputParam)
	default:
		return nil, fmt.Errorf("unknown resource format '%v'", resourceFormat)
	}
}

func newKRM(tfProvider *schema.Provider, outputParam string) (OutputSink, error) {
	return newSink(tfProvider, outputParam, NewKRMYAMLDirectory)
}

func newHCL(tfProvider *schema.Provider, outputParam string) (OutputSink, error) {
	return newSink(tfProvider, outputParam, NewHCLDirectory)
}

func newSink(tfProvider *schema.Provider, outputParam string, newDirectoryFunc func(*schema.Provider, string) OutputSink) (OutputSink, error) {
	if outputParam == "" {
		return NewWriter(os.Stdout), nil
	}
	fi, ok := getFileInfo(outputParam)
	if ok {
		if fi.Mode().IsRegular() {
			return NewFile(outputParam)
		}
		if fi.IsDir() {
			return newDirectoryFunc(tfProvider, outputParam), nil
		}
		return nil, fmt.Errorf("cannot use output parameter '%v': is neither a 'regular' file or directory", outputParam)
	}
	// if the output path ends with a '/' then the user intends for a directory to be created
	if strings.HasSuffix(outputParam, string(os.PathSeparator)) {
		if err := os.MkdirAll(outputParam, os.ModePerm); err != nil {
			return nil, fmt.Errorf("error creating directory '%v': %w", outputParam, err)
		}
		return newDirectoryFunc(tfProvider, outputParam), nil
	}
	// if the 'output' parameter is not a file that exists, determine if its parent folder exists
	// if the parent does exist -- then use the output parameter as a single output file
	// if the parent does not exist, then use the output parameters as a new output folder
	dir := filepath.Dir(outputParam)
	fi, ok = getFileInfo(dir)
	if ok {
		if fi.IsDir() {
			return NewFile(outputParam)
		}
		return nil, fmt.Errorf("cannot use output parameter '%v': parent path '%v' exists, but is not a directory", outputParam, dir)
	}
	if err := os.MkdirAll(outputParam, os.ModePerm); err != nil {
		return nil, fmt.Errorf("error creating directory '%v': %w", outputParam, err)
	}
	return newDirectoryFunc(tfProvider, outputParam), nil
}

func getFileInfo(filePath string) (os.FileInfo, bool) {
	fi, err := os.Stat(filePath)
	if err == nil {
		return fi, true
	}
	return nil, false
}
