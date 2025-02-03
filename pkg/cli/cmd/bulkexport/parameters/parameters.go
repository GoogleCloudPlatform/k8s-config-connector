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

package parameters

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/commonparams"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/valutil"
	"golang.org/x/oauth2"
)

type OnErrorOption string
type IAMFormatOption string

const (
	InputParam          = "input"
	OnErrorParam        = "on-error"
	StorageKeyParam     = "storage-key"
	ProjectIDParam      = "project"
	FolderIDParam       = "folder"
	OrganizationIDParam = "organization"

	ContinueOnErrorOption = "continue"
	HaltOnErrorOption     = "halt"
	IgnoreOnErrorOption   = "ignore"
)

var (
	AllErrorOptions = []string{
		ContinueOnErrorOption,
		HaltOnErrorOption,
		IgnoreOnErrorOption,
	}
)

type Parameters struct {
	IAMFormat               string
	FilterDeletedIAMMembers bool
	Input                   string
	Output                  string
	OnError                 string
	StorageKey              string
	ProjectID               string
	FolderID                int
	OrganizationID          int
	OAuth2Token             string
	ResourceFormat          string
	Verbose                 bool
}

func (p *Parameters) ControllerConfig() *config.ControllerConfig {
	c := &config.ControllerConfig{
		UserAgent: gcp.KCCUserAgent(),
	}
	if p.OAuth2Token != "" {
		c.GCPTokenSource = oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: p.OAuth2Token},
		)
	}
	return c
}

// convenience struct used during validation
type param struct {
	Value interface{}
	Name  string
}

func Validate(p *Parameters, stdin *os.File) error {
	inputParam := param{Value: &p.Input, Name: InputParam}
	storageKeyParam := param{Value: &p.StorageKey, Name: StorageKeyParam}
	projectIDParam := param{Value: &p.ProjectID, Name: ProjectIDParam}
	folderIDParam := param{Value: &p.FolderID, Name: FolderIDParam}
	organizationIDParam := param{Value: &p.OrganizationID, Name: OrganizationIDParam}
	if err := validateMutuallyExclusiveParams(inputParam, storageKeyParam, projectIDParam, folderIDParam, organizationIDParam); err != nil {
		return err
	}
	if err := validateMutuallyExclusiveParams(projectIDParam, folderIDParam, organizationIDParam); err != nil {
		return err
	}
	if err := validateMutuallyExclusiveParams(folderIDParam, organizationIDParam); err != nil {
		return err
	}
	if err := validatePipedInput(stdin, inputParam, storageKeyParam, projectIDParam, folderIDParam, organizationIDParam); err != nil {
		return err
	}
	if err := validateStorageKey(p); err != nil {
		return err
	}
	if err := validateOnError(p); err != nil {
		return err
	}
	if err := commonparams.ValidateIAMFormat(p.IAMFormat); err != nil {
		return err
	}
	if err := commonparams.ValidateResourceFormat(p.ResourceFormat, p.IAMFormat); err != nil {
		return err
	}

	return validateOneInput(p, stdin)
}

func IsInputPiped(stdin *os.File) (bool, error) {
	fi, err := stdin.Stat()
	if err != nil {
		return false, fmt.Errorf("error stating stdin: %w", err)
	}
	return (fi.Mode() & os.ModeCharDevice) == 0, nil
}

func validateOneInput(p *Parameters, stdin *os.File) error {
	piped, err := IsInputPiped(stdin)
	if err != nil {
		return err
	}
	if piped {
		return nil
	}
	if !valutil.IsDefaultValue(p.StorageKey) {
		return nil
	}
	if !valutil.IsDefaultValue(p.Input) {
		return nil
	}
	if err := validateCanExport(p); err == nil {
		return nil
	}
	return fmt.Errorf("no input or export parameters supplied, must supply an asset inventory on 'stdin' or the '%v' parameter or supply one of '%v', '%v', or '%v' to perform an export",
		InputParam, ProjectIDParam, FolderIDParam, OrganizationIDParam)
}

func validatePipedInput(stdin *os.File, exclusiveParams ...param) error {
	piped, err := IsInputPiped(stdin)
	if err != nil {
		return err
	}
	if piped {
		for _, p := range exclusiveParams {
			if !valutil.IsDefaultValue(p.Value) {
				return fmt.Errorf("cannot supply input on stdin with the '%v' parameter", p.Name)
			}
		}
	}
	return nil
}

func validateOnError(p *Parameters) error {
	onErrorOptions := []string{ContinueOnErrorOption, HaltOnErrorOption, IgnoreOnErrorOption}
	if valutil.IsDefaultValue(p.OnError) {
		return fmt.Errorf("invalid empty value for %v: must be one of {%v}", OnErrorParam, strings.Join(onErrorOptions, ", "))
	}
	for _, o := range onErrorOptions {
		if p.OnError == o {
			return nil
		}
	}
	return fmt.Errorf("invalid %v value of '%v': must be one of {%v}", OnErrorParam, p.OnError, strings.Join(onErrorOptions, ", "))
}

func validateStorageKey(p *Parameters) error {
	if valutil.IsDefaultValue(p.StorageKey) {
		return nil
	}
	regex := regexp.MustCompile("gs://.*")
	if !regex.MatchString(p.StorageKey) {
		return fmt.Errorf("invalid %v value of '%v': must be a valid cloud storage URI", StorageKeyParam, p.StorageKey)
	}

	return validateCanExport(p)
}

func validateCanExport(p *Parameters) error {
	if p.ProjectID == "" && p.FolderID == 0 && p.OrganizationID == 0 {
		return fmt.Errorf("one of the '%v', '%v', or '%v' parameters must be defined to perform an export",
			ProjectIDParam, FolderIDParam, OrganizationIDParam)
	}
	return nil
}

func validateMutuallyExclusiveParams(verifyParam param, otherParams ...param) error {
	if valutil.IsDefaultValue(verifyParam.Value) {
		return nil
	}
	for _, p := range otherParams {
		if valutil.IsDefaultValue(p.Value) {
			continue
		}
		return fmt.Errorf("cannot supply both '%v' and '%v': the parameters are mutually exclusive", verifyParam.Name, p.Name)
	}
	return nil
}
