// Copyright 2026 Google LLC
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
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"sigs.k8s.io/yaml"
)

type Commit struct {
	SHA     string   `json:"sha"`
	Parents []Parent `json:"parents"`
	Commit  struct {
		Committer struct {
			Date time.Time `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
	Author *struct {
		Login string `json:"login"`
	} `json:"author"`
}

type Parent struct {
	SHA string `json:"sha"`
}

type TreeItem struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	SHA  string `json:"sha"`
	Size int    `json:"size"`
}

type TreeResponse struct {
	SHA  string     `json:"sha"`
	Tree []TreeItem `json:"tree"`
}

var (
	weeks            = flag.Int("w", 1, "Number of weeks to go back in time for")
	currentWeek      = flag.Bool("current-week", false, "Generate stats for this week so far (equivalent to -w 0)")
	verbose          = flag.Bool("verbose", false, "Tell you about the commits being fetched and processed")
	showDetails      = flag.Bool("show-details", false, "Name the resources/kinds added and the fields added as a json path like spec.a.b.c")
	showContributors = flag.Bool("show-contributors", false, "Break down metrics by contributor")
	showAllCommits   = flag.Bool("show-all-commits", false, "Show the total number of merge commits (not just positive ones)")
)

type ContributorStats struct {
	CommitCount int
	NewRes      int
	NewFields   int
}

const repo = "GoogleCloudPlatform/k8s-config-connector"
const crdPath = "config/crds/resources"

func main() {
	flag.Parse()

	if *currentWeek {
		*weeks = 0
	}

	now := time.Now().UTC()
	// Calculate the start of the current ISO week (Monday, 00:00:00 UTC)
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7 // Sunday is 7
	}
	startOfCurrentWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, -weekday+1)

	since := startOfCurrentWeek.AddDate(0, 0, -7*(*weeks))
	sinceStr := since.Format(time.RFC3339)

	if *weeks == 0 {
		fmt.Printf("Analyzing velocity for the current week so far (since %s)...\n", sinceStr)
	} else {
		fmt.Printf("Analyzing velocity for the last %d weeks (since %s)...\n", *weeks, sinceStr)
	}

	if *verbose {
		fmt.Println("Fetching commits via gh api...")
	}

	// Fetch commits
	cmd := exec.Command("gh", "api", "-X", "GET", "repos/"+repo+"/commits", "-F", "since="+sinceStr, "--paginate")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to fetch commits: %v\nOutput: %s", err, string(out))
	}

	var allCommits []Commit
	if err := json.Unmarshal(out, &allCommits); err != nil {
		log.Fatalf("Failed to parse commits: %v", err)
	}

	// Filter merge commits
	var mergeCommits []Commit
	for _, c := range allCommits {
		if len(c.Parents) > 1 {
			mergeCommits = append(mergeCommits, c)
		}
	}

	if len(mergeCommits) == 0 {
		fmt.Println("No merge commits found in the given time period.")
		return
	}

	// Sort chronologically (oldest first)
	sort.Slice(mergeCommits, func(i, j int) bool {
		return mergeCommits[i].Commit.Committer.Date.Before(mergeCommits[j].Commit.Committer.Date)
	})

	if *verbose {
		fmt.Printf("Found %d merge commits.\n", len(mergeCommits))
	}

	// Get state map before the first commit
	firstCommit := mergeCommits[0]
	beforeCommitSHA := firstCommit.Parents[0].SHA

	if *verbose {
		fmt.Printf("Fetching initial state at commit %s...\n", beforeCommitSHA)
	}
	stateMap := getTreeState(beforeCommitSHA)

	// Setup caching directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	cacheDir := filepath.Join(homeDir, ".kccstats_cache")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		log.Fatal(err)
	}

	type WeekStats struct {
		Week                int
		Year                int
		NewRes              int
		NewFields           int
		AllCommitCount      int
		PositiveCommitCount int
		Contributors        map[string]*ContributorStats
		DetailsRes          []string
		DetailsFld          []string
	}
	statsMap := make(map[string]*WeekStats)

	for i, mc := range mergeCommits {
		if *verbose && i%10 == 0 {
			fmt.Printf("Processing commit %d/%d (%s)...\n", i+1, len(mergeCommits), mc.SHA)
		}
		year, week := mc.Commit.Committer.Date.ISOWeek()
		weekKey := fmt.Sprintf("%d-W%02d", year, week)
		if statsMap[weekKey] == nil {
			statsMap[weekKey] = &WeekStats{
				Year:         year,
				Week:         week,
				Contributors: make(map[string]*ContributorStats),
			}
		}
		stats := statsMap[weekKey]
		stats.AllCommitCount++

		authorLogin := "unknown"
		if mc.Author != nil && mc.Author.Login != "" {
			authorLogin = mc.Author.Login
		}

		newTree := getTreeState(mc.SHA)

		commitNewRes := 0
		commitNewFields := 0

		for path, newSHA := range newTree {
			if !strings.HasSuffix(path, ".yaml") {
				continue
			}

			oldSHA, exists := stateMap[path]
			if !exists {
				// Added resource
				commitNewRes++
				newContent := getBlob(newSHA, cacheDir)
				kind := getKindFromYaml(newContent)
				stats.DetailsRes = append(stats.DetailsRes, fmt.Sprintf("%s (by %s)", kind, authorLogin))
			} else if oldSHA != newSHA {
				// Modified resource
				oldContent := getBlob(oldSHA, cacheDir)
				newContent := getBlob(newSHA, cacheDir)

				oldPaths := getFieldPaths(oldContent)
				newPaths := getFieldPaths(newContent)

				kind := getKindFromYaml(newContent)

				added := make(map[string]bool)
				for p := range newPaths {
					if _, found := oldPaths[p]; !found {
						added[p] = true
					}
				}

				netFields := 0
				for p := range added {
					isChildOfAdded := false
					curr := p
					for {
						curr = getParentPath(curr)
						if curr == "" {
							break
						}
						if added[curr] {
							isChildOfAdded = true
							break
						}
					}

					if !isChildOfAdded {
						netFields++
						stats.DetailsFld = append(stats.DetailsFld, fmt.Sprintf("%s => %s (by %s)", kind, p, authorLogin))
					}
				}

				if netFields > 0 {
					commitNewFields += netFields
				}
			}
		}

		if commitNewRes > 0 || commitNewFields > 0 {
			stats.PositiveCommitCount++
			stats.NewRes += commitNewRes
			stats.NewFields += commitNewFields

			if stats.Contributors[authorLogin] == nil {
				stats.Contributors[authorLogin] = &ContributorStats{}
			}
			stats.Contributors[authorLogin].CommitCount++
			stats.Contributors[authorLogin].NewRes += commitNewRes
			stats.Contributors[authorLogin].NewFields += commitNewFields
		}

		// Update state map
		stateMap = newTree
	}

	// Report
	fmt.Printf("\n--- KCC Velocity Report ---\n")
	fmt.Printf("Note: Positive Merge Commits only count commits that add a new resource or net new fields.\n\n")

	if *showAllCommits {
		fmt.Printf("%-20s | %-17s | %-22s | %-13s | %-15s\n", "Week/Contributor", "All Merge Commits", "Positive Merge Commits", "New Resources", "Net New Fields")
		fmt.Printf("------------------------------------------------------------------------------------------------------\n")
	} else {
		fmt.Printf("%-20s | %-22s | %-13s | %-15s\n", "Week/Contributor", "Positive Merge Commits", "New Resources", "Net New Fields")
		fmt.Printf("------------------------------------------------------------------------------\n")
	}

	// Sort weeks
	var weeksKeys []string
	for k := range statsMap {
		weeksKeys = append(weeksKeys, k)
	}
	sort.Strings(weeksKeys)

	for _, k := range weeksKeys {
		s := statsMap[k]
		if *showAllCommits {
			fmt.Printf("%-20s | %-17d | %-22d | %-13d | %-15d\n", k, s.AllCommitCount, s.PositiveCommitCount, s.NewRes, s.NewFields)
		} else {
			fmt.Printf("%-20s | %-22d | %-13d | %-15d\n", k, s.PositiveCommitCount, s.NewRes, s.NewFields)
		}

		if *showContributors {
			var authorsList []string
			for author := range s.Contributors {
				authorsList = append(authorsList, author)
			}
			sort.Strings(authorsList)
			for _, author := range authorsList {
				cStats := s.Contributors[author]
				// If author name is too long, truncate it
				displayAuthor := author
				if len(displayAuthor) > 18 {
					displayAuthor = displayAuthor[:15] + "..."
				}

				if *showAllCommits {
					fmt.Printf("  %-18s | %-17s | %-22d | %-13d | %-15d\n", displayAuthor, "-", cStats.CommitCount, cStats.NewRes, cStats.NewFields)
				} else {
					fmt.Printf("  %-18s | %-22d | %-13d | %-15d\n", displayAuthor, cStats.CommitCount, cStats.NewRes, cStats.NewFields)
				}
			}
		}

		if *showDetails {
			if len(s.DetailsRes) > 0 {
				fmt.Printf("  -> Added Kinds:\n")
				for _, res := range s.DetailsRes {
					fmt.Printf("     - %s\n", res)
				}
			}
			if len(s.DetailsFld) > 0 {
				fmt.Printf("  -> Added Fields:\n")
				sort.Strings(s.DetailsFld)
				for _, f := range s.DetailsFld {
					fmt.Printf("     - %s\n", f)
				}
			}
			if len(s.DetailsRes) > 0 || len(s.DetailsFld) > 0 {
				fmt.Println()
			}
		}
	}
}

func getTreeState(commitSHA string) map[string]string {
	cmd := exec.Command("gh", "api", fmt.Sprintf("repos/%s/git/trees/%s:%s", repo, commitSHA, crdPath))
	out, err := cmd.Output()
	if err != nil {
		// Tree might not exist in early commits, just return empty map
		return make(map[string]string)
	}

	var resp TreeResponse
	if err := json.Unmarshal(out, &resp); err != nil {
		log.Fatalf("Failed to parse tree: %v", err)
	}

	state := make(map[string]string)
	for _, item := range resp.Tree {
		if item.Type == "blob" {
			state[item.Path] = item.SHA
		}
	}
	return state
}

func getBlob(sha string, cacheDir string) []byte {
	cachePath := filepath.Join(cacheDir, sha+".yaml")
	if content, err := os.ReadFile(cachePath); err == nil {
		return content
	}

	cmdRaw := exec.Command("gh", "api", "-H", "Accept: application/vnd.github.v3.raw", fmt.Sprintf("repos/%s/git/blobs/%s", repo, sha))
	outRaw, err := cmdRaw.Output()
	if err != nil {
		log.Printf("Failed to fetch raw blob %s: %v", sha, err)
		return nil
	}

	if err := os.WriteFile(cachePath, outRaw, 0644); err != nil {
		log.Printf("Failed to write cache file %s: %v", cachePath, err)
	}
	return outRaw
}

func getKindFromYaml(content []byte) string {
	if content == nil {
		return "Unknown"
	}
	var root map[string]interface{}
	if err := yaml.Unmarshal(content, &root); err != nil {
		return "Unknown"
	}
	spec := getMap(root, "spec")
	if spec == nil {
		return "Unknown"
	}
	names := getMap(spec, "names")
	if names == nil {
		return "Unknown"
	}
	if kind, ok := names["kind"].(string); ok {
		return kind
	}
	return "Unknown"
}

func getFieldPaths(content []byte) map[string]struct{} {
	paths := make(map[string]struct{})
	if content == nil {
		return paths
	}

	var root map[string]interface{}
	if err := yaml.Unmarshal(content, &root); err != nil {
		return paths
	}

	spec := getMap(root, "spec")
	if spec == nil {
		return paths
	}

	versions, ok := spec["versions"].([]interface{})
	if !ok || len(versions) == 0 {
		return paths
	}

	v0, ok := versions[0].(map[string]interface{})
	if !ok {
		return paths
	}

	schema := getMap(v0, "schema")
	if schema == nil {
		return paths
	}

	openAPI := getMap(schema, "openAPIV3Schema")
	if openAPI == nil {
		return paths
	}

	props := getMap(openAPI, "properties")
	if props == nil {
		return paths
	}

	specProps := getMap(props, "spec")
	if specProps == nil {
		return paths
	}

	collectProperties(getMap(specProps, "properties"), "spec", paths)
	return paths
}

func getMap(m map[string]interface{}, key string) map[string]interface{} {
	if m == nil {
		return nil
	}
	if v, ok := m[key]; ok {
		if res, ok := v.(map[string]interface{}); ok {
			return res
		}
	}
	return nil
}

func collectProperties(props map[string]interface{}, currentPath string, paths map[string]struct{}) {
	if props == nil {
		return
	}
	for k, v := range props {
		path := currentPath + "." + k
		paths[path] = struct{}{}
		if m, ok := v.(map[string]interface{}); ok {
			if p := getMap(m, "properties"); p != nil {
				collectProperties(p, path, paths)
			}
			if items := getMap(m, "items"); items != nil {
				if ip := getMap(items, "properties"); ip != nil {
					collectProperties(ip, path+"[]", paths)
				}
			}
		}
	}
}

func getParentPath(p string) string {
	idx := strings.LastIndex(p, ".")
	if idx == -1 {
		return ""
	}
	parent := p[:idx]
	if strings.HasSuffix(parent, "[]") {
		return parent[:len(parent)-2]
	}
	return parent
}
