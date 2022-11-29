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

package manifest

import (
	"context"
	"fmt"
)

func ResolveVersion(ctx context.Context, repo Repository, componentName string, channelName string) (string, error) {
	channel, err := repo.LoadChannel(ctx, channelName)
	if err != nil {
		return "", err
	}

	version, err := channel.Latest(ctx, componentName)
	if err != nil {
		return "", err
	}

	if version == nil {
		return "", fmt.Errorf("could not find latest version in channel %v", channelName)
	}
	return version.Version, nil
}
