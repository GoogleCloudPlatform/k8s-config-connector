// Copyright 2024 Google LLC
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

package migrator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/utils"
	"github.com/spf13/pflag"
)

const (
	migrationFileFlag = "migration-file"
)

type MigratorOptions struct {
	utils.ClusterCrawlOptions

	MigrationFile string
}

func (opts *MigratorOptions) AddFlags(flags *pflag.FlagSet) {
	opts.ClusterCrawlAddFlags(flags)

	flags.StringVarP(&opts.MigrationFile, migrationFileFlag, "", opts.MigrationFile, "path of the migration file to create.")
}

func (opts *MigratorOptions) validateFlags() error {
	if err := opts.ValidateClusterCrawlFlags(); err != nil {
		return err
	}
	if opts.MigrationFile == "" {
		return fmt.Errorf("%s is a required field", migrationFileFlag)
	}
	filedir := filepath.Dir(opts.MigrationFile)
	if _, err := os.Stat(filedir); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error checking migration file %s, got %v", opts.MigrationFile, err)
	}
	if _, err := os.Stat(opts.MigrationFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error checking migration file %s, got %v", opts.MigrationFile, err)
	}
	return nil
}

func (opts *MigratorOptions) Print() {
	opts.ClusterCrawlPrint()
	log.Printf("migrationfile set to %q.\n", opts.MigrationFile)
}

func NewMigratorOptions() *MigratorOptions {
	opts := MigratorOptions{
		MigrationFile: "migration.yaml",
	}
	opts.ClusterCrawlOptions = utils.NewClusterCrawlOptions()
	return &opts
}
