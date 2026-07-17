// Copyright 2022 The kpt Authors
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
	"archive/zip"
	"bytes"
	"context"
	"debug/buildinfo"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

//go:embed modules/*
var modulesFS embed.FS

func main() {
	rootCmd := buildRootCommand()

	err := rootCmd.ExecuteContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func buildRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "licensescan",
	}

	cmd.AddCommand(buildLicenseScanCommand())
	cmd.AddCommand(buildLicenseVerifyCommand())
	cmd.AddCommand(buildLicenseGenerateCommand())

	klog.InitFlags(nil)

	return cmd
}

// buildLicenseGenerateCommand builds the 'generate' command.
// The 'generate' command reads a go.mod file and ensures that a license metadata YAML file
// exists for each dependency. If it doesn't exist, it creates a template with "license: TODO".
func buildLicenseGenerateCommand() *cobra.Command {
	var opts RunLicenseGenerateOptions

	cmd := &cobra.Command{
		Use: "generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunLicenseGenerate(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.GoMod, "gomod", opts.GoMod, "path to go.mod file")
	_ = cmd.MarkFlagRequired("gomod")

	return cmd
}

type RunLicenseGenerateOptions struct {
	GoMod string
}

func RunLicenseGenerate(ctx context.Context, opts RunLicenseGenerateOptions) error {
	b, err := os.ReadFile(opts.GoMod)
	if err != nil {
		return fmt.Errorf("error reading go.mod file %q: %w", opts.GoMod, err)
	}
	f, err := modfile.Parse(opts.GoMod, b, nil)
	if err != nil {
		return fmt.Errorf("error parsing go.mod file %q: %w", opts.GoMod, err)
	}

	for _, require := range f.Require {
		module := &Module{
			Name:    require.Mod.Path,
			Version: require.Mod.Version,
		}

		if err := ensureModuleFile(module); err != nil {
			fmt.Fprintf(os.Stderr, "Error for %s@%s: %v\n", module.Name, module.Version, err)
		}
	}
	return nil
}

func ensureModuleFile(module *Module) error {
	p := filepath.Join("experiments/tools/licensescan/modules", module.Name, module.Version+".yaml")
	if _, err := os.Stat(p); err == nil {
		return nil // Already exists
	}

	// Also check in embed FS
	pFS := filepath.Join("modules", module.Name, module.Version+".yaml")
	if _, err := modulesFS.ReadFile(pFS); err == nil {
		return nil // Already exists in embed
	}

	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return fmt.Errorf("error from Mkdir(%q): %w", filepath.Dir(p), err)
	}
	s := "# https://pkg.go.dev/" + module.Name + "@" + module.Version + "\n"
	s += "license: TODO"
	if err := os.WriteFile(p, []byte(s), 0644); err != nil {
		return fmt.Errorf("error writing %q: %w", p, err)
	}
	fmt.Printf("Created %s\n", p)
	return nil
}

type Module struct {
	Name string `json:"name"`
	Sum  string `json:"sum"`
	//Path    string
	Version string `json:"version"`

	Info         *moduleInfo   `json:"licenseInfo,omitempty"`
	LicenseFiles []LicenseFile `json:"licenseFiles,omitempty"`
}

// buildLicenseScanCommand builds the 'scan' command.
// The 'scan' command analyzes a compiled Go binary to identify its dependencies
// and then looks up their licenses using the embedded or local metadata.
// It can optionally include the full license text in its output.
func buildLicenseScanCommand() *cobra.Command {
	var opts RunLicenseScanOptions

	cmd := &cobra.Command{
		Use: "scan",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunLicenseScan(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.Binary, "binary", opts.Binary, "binary to analyze")
	_ = cmd.MarkFlagRequired("binary")

	cmd.Flags().BoolVar(&opts.IncludeLicenses, "print", opts.IncludeLicenses, "include license text")

	cmd.Flags().StringArrayVar(&opts.IgnorePackage, "ignore", opts.IgnorePackage, "packages to ignore")

	return cmd
}

// buildLicenseVerifyCommand builds the 'verify' command.
// The 'verify' command ensures that all dependencies (from go.mod or the embedded metadata)
// have a valid license specified in their metadata YAML file.
// It fails if any license is marked as "TODO" or is missing.
func buildLicenseVerifyCommand() *cobra.Command {
	var opts RunLicenseVerifyOptions

	cmd := &cobra.Command{
		Use: "verify",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunLicenseVerify(cmd.Context(), opts)
		},
	}

	cmd.Flags().StringVar(&opts.GoMod, "gomod", opts.GoMod, "path to go.mod file")

	return cmd
}

type RunLicenseScanOptions struct {
	Binary string

	// IgnorePackage can be useful for internal libraries
	IgnorePackage []string

	IncludeLicenses bool
}

type RunLicenseVerifyOptions struct {
	GoMod string
}

func RunLicenseVerify(ctx context.Context, opts RunLicenseVerifyOptions) error {
	var errors []error

	if opts.GoMod != "" {
		b, err := os.ReadFile(opts.GoMod)
		if err != nil {
			return fmt.Errorf("error reading go.mod file %q: %w", opts.GoMod, err)
		}
		f, err := modfile.Parse(opts.GoMod, b, nil)
		if err != nil {
			return fmt.Errorf("error parsing go.mod file %q: %w", opts.GoMod, err)
		}
		for _, require := range f.Require {
			module := &Module{
				Name:    require.Mod.Path,
				Version: require.Mod.Version,
			}

			if err := checkModule(module); err != nil {
				errors = append(errors, err)
			}
		}
	} else {
		err := fs.WalkDir(modulesFS, "modules", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				if strings.HasSuffix(path, ".yaml") {
					module := &Module{}
					dir := filepath.Dir(path)
					base := filepath.Base(path)
					module.Name = strings.TrimPrefix(dir, "modules/")
					module.Version = strings.TrimSuffix(base, ".yaml")

					if err := checkModule(module); err != nil {
						errors = append(errors, err)
					}
				}
			}
			return nil
		})
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		return fmt.Errorf("license verification failed")
	}
	return nil
}

var mustShipCodeByLicense = map[string]bool{
	"APACHE-2.0":       false,
	"BSD-3-CLAUSE":     false,
	"BSD-2-CLAUSE":     false,
	"MIT":              false,
	"UNICODE-DFS-2016": false,
	"MPL-2.0":          true,
	"PUBLICDOMAIN":     false,
	"ISC":              false,
	"HPND":             false,
	"CC-BY-3.0":        false,
	"BSL-1.0":          false,
	"NCSA":             false,
	"OPENSSL":          false,
	"ZLIB":             false,
}

func checkModule(module *Module) error {
	p := filepath.Join("experiments/tools/licensescan/modules", module.Name, module.Version+".yaml")
	b, err := os.ReadFile(p)
	if err != nil && os.IsNotExist(err) {
		pFS := filepath.Join("modules", module.Name, module.Version+".yaml")
		b, err = modulesFS.ReadFile(pFS)
	}
	if err != nil {
		return fmt.Errorf("error reading module file for %s@%s: %w", module.Name, module.Version, err)
	}

	info := &moduleInfo{}
	if err := yaml.Unmarshal(b, info); err != nil {
		return fmt.Errorf("error parsing %s@%s: %w", module.Name, module.Version, err)
	}
	info.License = strings.TrimSpace(info.License)
	if info.License == "TODO" || info.License == "" {
		return fmt.Errorf("license not known for %s@%s", module.Name, module.Version)
	}

	licenses := strings.Split(info.License, ",")
	for _, license := range licenses {
		license = strings.TrimSpace(license)
		license = strings.ToUpper(license)

		if _, exists := mustShipCodeByLicense[license]; !exists {
			return fmt.Errorf("unknown license %q (for %s@%s)", license, module.Name, module.Version)
		}
	}

	return nil
}

func RunLicenseScan(ctx context.Context, opts RunLicenseScanOptions) error {
	buildInfo, err := buildinfo.ReadFile(opts.Binary)
	if err != nil {
		return fmt.Errorf("error reading binary info from %q: %w", opts.Binary, err)
	}

	var errors []error

	var modules []*Module
	for _, dep := range buildInfo.Deps {
		ignore := false
		for _, ignorePackage := range opts.IgnorePackage {
			if ignorePackage == dep.Path {
				klog.Infof("ignoring package %s@%s", dep.Path, dep.Version)
				ignore = true
				break
			}
		}
		if ignore {
			continue
		}

		module := &Module{
			Name:    dep.Path,
			Sum:     dep.Sum,
			Version: dep.Version,
		}
		if dep.Replace != nil {
			module.Version = dep.Replace.Version
		}
		modules = append(modules, module)

		p := filepath.Join("experiments/tools/licensescan/modules", module.Name, module.Version+".yaml")
		b, err := os.ReadFile(p)
		if err != nil && os.IsNotExist(err) {
			pFS := filepath.Join("modules", module.Name, module.Version+".yaml")
			b2, err2 := modulesFS.ReadFile(pFS)
			if err2 == nil {
				b = b2
				err = err2
			}
		}
		if err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
					return fmt.Errorf("error from Mkdir(%q): %w", filepath.Dir(p), err)
				}
				s := "# https://pkg.go.dev/" + module.Name + "@" + module.Version + "\n"
				s += "license: TODO"
				b = []byte(s)
				if err := os.WriteFile(p, b, 0644); err != nil {
					return fmt.Errorf("error writing %q: %w", p, err)
				}
			} else {
				return fmt.Errorf("error reading %q: %w", p, err)
			}
		}
		info := &moduleInfo{}
		if err := yaml.Unmarshal(b, info); err != nil {
			return fmt.Errorf("error parsing %q: %w", p, err)
		}
		info.License = strings.TrimSpace(info.License)
		if info.License == "TODO" || info.License == "" {
			errors = append(errors, fmt.Errorf("license not known for %s@%s", module.Name, module.Version))
		}
		module.Info = info
	}

	if len(errors) == 0 && opts.IncludeLicenses {
		for _, module := range modules {
			if licenseFiles, err := includeLicense(ctx, module); err != nil {
				errors = append(errors, fmt.Errorf("error getting license text for %s@%s: %w", module.Name, module.Version, err))
			} else {
				module.LicenseFiles = licenseFiles
			}
		}
	}

	if len(errors) == 0 {
		for _, module := range modules {
			if len(module.Info.LicenseURLs) == 0 {
				licenseURL := "https://pkg.go.dev/" + module.Name + "@" + module.Version + "?tab=licenses"
				module.Info.LicenseURLs = []string{licenseURL}
			}
			licenses := strings.Split(module.Info.License, ",")
			for _, license := range licenses {
				license = strings.TrimSpace(license)
				license = strings.ToUpper(license)

				if mustShipCodeByLicense[license] {
					module.Info.MustShipCode = true
				}
			}
		}
	}

	if len(errors) == 0 {
		j, err := json.Marshal(modules)
		if err != nil {
			return fmt.Errorf("error converting to json: %w", err)
		}
		fmt.Printf("%s\n", j)

		return nil
	}

	for _, err := range errors {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	return fmt.Errorf("could not determine all licenses")
}

type moduleInfo struct {
	License      string   `json:"license,omitempty"`
	LicenseURLs  []string `json:"licenseURLs,omitempty"`
	MustShipCode bool     `json:"mustShipCode"`
}

type goModDownloadInfo struct {
	Zip string `json:"Zip"`
}

type LicenseFile struct {
	Path     string `json:"path,omitempty"`
	Contents string `json:"contents,omitempty"`
}

func includeLicense(ctx context.Context, module *Module) ([]LicenseFile, error) {
	var licenses []LicenseFile

	if len(module.Info.LicenseURLs) != 0 {
		for _, licenseURL := range module.Info.LicenseURLs {
			if err := func() error {
				response, err := http.Get(licenseURL)
				if err != nil {
					return fmt.Errorf("error reading %q: %w", licenseURL, err)
				}
				defer response.Body.Close()
				if response.StatusCode != 200 {
					return fmt.Errorf("unexpected response from GET %q: %v", licenseURL, response.Status)
				}
				b, err := io.ReadAll(response.Body)
				if err != nil {
					return fmt.Errorf("error reading response from %q: %w", licenseURL, err)
				}
				licenses = append(licenses, LicenseFile{
					Path:     licenseURL,
					Contents: string(b),
				})
				return nil
			}(); err != nil {
				return nil, err
			}
		}
		return licenses, nil
	}
	klog.Infof("downloading %s@%s", module.Name, module.Version)
	cmd := exec.CommandContext(ctx, "go", "mod", "download", "-json", module.Name+"@"+module.Version)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("error doing go mod download: %w", err)
	}

	var info goModDownloadInfo
	if err := json.Unmarshal(stdout.Bytes(), &info); err != nil {
		return nil, fmt.Errorf("error parsing go mod download output (%q): %w", stdout.String(), err)
	}

	zipfile, err := zip.OpenReader(info.Zip)
	if err != nil {
		return nil, fmt.Errorf("error opening zip file %q: %w", info.Zip, err)
	}
	defer zipfile.Close()

	for _, f := range zipfile.File {
		isLicense := false

		name := filepath.Base(f.Name)
		name = strings.ToUpper(name)
		switch name {
		case "LICENSE", "LICENSE.TXT", "LICENSE-APACHE-2.0.TXT", "COPYING", "LICENSE.MD", "LICENSE.MIT":
			isLicense = true
		}
		if isLicense {
			if err := func() error {
				r, err := f.Open()
				if err != nil {
					return fmt.Errorf("error opening entry %q: %w", f.Name, err)
				}
				defer r.Close()
				b, err := io.ReadAll(r)
				if err != nil {
					return fmt.Errorf("error reading entry %q: %w", f.Name, err)
				}
				licenses = append(licenses, LicenseFile{
					Path:     f.Name,
					Contents: string(b),
				})
				return nil
			}(); err != nil {
				return nil, err
			}
		}
	}

	if len(licenses) == 0 {
		return nil, fmt.Errorf("unable to find license entries in zipfile %q", info.Zip)
	}
	return licenses, nil
}
