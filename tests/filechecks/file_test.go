package filechecks

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// Make sure new files have the right year in their license header files.
func TestFilesHaveCurrentYear(t *testing.T) {
	currentYear := time.Now().Year()
	yearString := fmt.Sprintf("Copyright %d Google LLC", currentYear)

	// This is a little fragile because origin/master is assumed to point to master for k8s-config-connector.
	cmd := exec.Command("git", "diff", "--name-only", "--diff-filter=A", "origin/master")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to fetch git diff: %s, error: %v", output, err)
	}

	files := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, file := range files {
		// let's start with go files for now
		if strings.HasSuffix(file, ".go") {
			t.Run(file, func(t *testing.T) {
				// two levels higher is the root repo
				repoFilePath := filepath.Join("..", "..", file)
				absPath, err := filepath.Abs(repoFilePath)
				if err != nil {
					t.Fatalf("Failed to get absolute path for file %s: %v", file, err)
				}

				f, err := os.Open(absPath)
				if err != nil {
					t.Fatalf("Failed to open file %s: %v", absPath, err)
				}
				defer f.Close()

				scanner := bufio.NewScanner(f)
				foundYear := false
				for scanner.Scan() {
					line := scanner.Text()
					if strings.Contains(line, yearString) {
						foundYear = true
						break
					}
				}

				if !foundYear {
					t.Errorf("File %s does not contain the current year (%d) in the copyright header.", file, currentYear)
				}
			})
		}
	}
}
