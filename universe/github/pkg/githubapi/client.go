// Copyright 2023 Google LLC
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

package githubapi

import (
	"context"

	github "github.com/google/go-github/v53/github"
)

type Client struct {
	impl *github.Client
}

func NewClient(ctx context.Context, impl *github.Client) (*Client, error) {
	return &Client{impl: impl}, nil
}

func (c *Client) GetIssue(ctx context.Context, org string, repo string, number int) (*github.Issue, error) {
	obj, _, err := c.impl.Issues.Get(ctx, org, repo, number)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) GetRepo(ctx context.Context, org string, repo string) (*github.Repository, error) {
	obj, _, err := c.impl.Repositories.Get(ctx, org, repo)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) ListRepoTagProtection(ctx context.Context, org string, repo string) ([]*github.TagProtection, error) {
	objs, _, err := c.impl.Repositories.ListTagProtection(ctx, org, repo)
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (c *Client) CreateRepoTagProtection(ctx context.Context, org string, repo string, pattern string) (*github.TagProtection, error) {
	obj, _, err := c.impl.Repositories.CreateTagProtection(ctx, org, repo, pattern)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *Client) DeleteRepoTagProtection(ctx context.Context, org string, repo string, tagProtectionID int64) error {
	_, err := c.impl.Repositories.DeleteTagProtection(ctx, org, repo, tagProtectionID)
	if err != nil {
		return err
	}
	return nil
}
