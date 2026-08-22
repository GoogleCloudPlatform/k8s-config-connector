package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var strictMode = false
	flag.BoolVar(&strictMode, "strict", strictMode, "Use strict index-based alignment instead of smart matching")
	flag.Parse()

	args := flag.Args()

	var name1, name2 string
	var reader1, reader2 func() ([]byte, error)

	if len(args) == 0 {
		// Find _http.log in current directory
		matches, err := filepath.Glob("_http.log")
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return fmt.Errorf("no _http.log found in current directory")
		}
		path := matches[0]

		// Compare HEAD:path vs path
		name1 = "HEAD:" + path
		reader1 = func() ([]byte, error) { return getGitContent("HEAD", path) }
		name2 = path
		reader2 = func() ([]byte, error) { return os.ReadFile(path) }
	} else if len(args) == 1 {
		path := args[0]
		name1 = "HEAD:" + path
		reader1 = func() ([]byte, error) { return getGitContent("HEAD", path) }
		name2 = path
		reader2 = func() ([]byte, error) { return os.ReadFile(path) }
	} else {
		name1 = args[0]
		reader1 = func() ([]byte, error) { return os.ReadFile(name1) }
		name2 = args[1]
		reader2 = func() ([]byte, error) { return os.ReadFile(name2) }
	}

	content1, err := reader1()
	if err != nil {
		return fmt.Errorf("error reading %s: %w", name1, err)
	}
	content2, err := reader2()
	if err != nil {
		return fmt.Errorf("error reading %s: %w", name2, err)
	}

	entries1, err := parseLog(content1)
	if err != nil {
		return fmt.Errorf("error parsing %s: %w", name1, err)
	}
	entries2, err := parseLog(content2)
	if err != nil {
		return fmt.Errorf("error parsing %s: %w", name2, err)
	}

	compareLogs(entries1, entries2, strictMode)
	return nil
}

func getGitContent(ref, path string) ([]byte, error) {
	// git show ref:path
	// If path is relative, we might need to resolve it relative to git root,
	// but git show usually works with relative paths if run from subdirs.
	// Let's verify.
	cmd := exec.Command("git", "show", fmt.Sprintf("%s:%s", ref, path))
	return cmd.Output()
}

type LogEntry struct {
	RequestLine  string
	RequestBody  string // JSON
	ResponseBody string // JSON
}

func parseLog(data []byte) ([]LogEntry, error) {
	var entries []LogEntry

	blocks := bytes.Split(data, []byte("\n---\n"))

	for _, block := range blocks {
		block = bytes.TrimSpace(block)
		if len(block) == 0 {
			continue
		}
		entry, err := parseBlock(block)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func parseBlock(block []byte) (LogEntry, error) {
	// Simple state machine or split by lines
	// 1. Request Line
	// 2. Request Headers (until empty line)
	// 3. Request Body (until Response Status Line)
	// 4. Response Status Line
	// 5. Response Headers (until empty line)
	// 6. Response Body

	lines := strings.Split(string(block), "\n")
	if len(lines) == 0 {
		return LogEntry{}, fmt.Errorf("empty block")
	}

	entry := LogEntry{
		RequestLine: lines[0],
	}

	idx := 1
	// Skip Request Headers
	for idx < len(lines) {
		line := strings.TrimSpace(lines[idx])
		idx++
		if line == "" {
			break
		}
	}

	// Now we are at Request Body or Response Status Line
	// Collect lines until we find a Response Status Line
	// A Response Status Line starts with HTTP/ or a 3-digit code

	requestBodyLines := []string{}
	responseStatusLineIdx := -1

	for i := idx; i < len(lines); i++ {
		line := lines[i]

		if isResponseStatusLine(line) {
			responseStatusLineIdx = i
			break
		}
		requestBodyLines = append(requestBodyLines, line)
	}

	if responseStatusLineIdx != -1 {
		entry.RequestBody = strings.TrimSpace(strings.Join(requestBodyLines, "\n"))

		// Parse Response
		idx = responseStatusLineIdx + 1
		// Skip Response Headers
		for idx < len(lines) {
			line := strings.TrimSpace(lines[idx])
			idx++
			if line == "" {
				break
			}
		}

		if idx < len(lines) {
			entry.ResponseBody = strings.TrimSpace(strings.Join(lines[idx:], "\n"))
		}
	}

	return entry, nil
}

func isResponseStatusLine(s string) bool {
	if len(s) < 4 {
		return false
	}
	// Check if starts with 3 digits and a space
	if s[3] != ' ' {
		return false
	}
	for i := 0; i < 3; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	// Optional: check for known status strings or assume format is consistent
	return true
}

func compareLogs(log1, log2 []LogEntry, strictMode bool) {
	if strictMode {
		compareLogsStrict(log1, log2)
	} else {
		compareLogsSmart(log1, log2)
	}
}

func compareLogsStrict(log1, log2 []LogEntry) {
	// Naive alignment by index
	limit := len(log1)
	if len(log2) < limit {
		limit = len(log2)
	}

	for i := 0; i < limit; i++ {
		e1 := log1[i]
		e2 := log2[i]

		// Check if Request matches roughly
		// We compare the request line (Method + URL)
		// Ignoring maybe IDs if they changed?
		// For now, simple string compare.

		// Compare Response Bodies
		diffs := diffJSON(e1.ResponseBody, e2.ResponseBody)
		if len(diffs) > 0 {
			fmt.Printf("Request: %s\n", e1.RequestLine)
			for _, d := range diffs {
				fmt.Printf("  %s\n", d)
			}
			fmt.Println()
		}
	}
}

func compareLogsSmart(log1, log2 []LogEntry) {
	j := 0
	for _, e1 := range log1 {
		// Search for matching request in log2 starting at j
		foundAt := -1
		for k := j; k < len(log2); k++ {
			if areRequestLinesSimilar(log2[k].RequestLine, e1.RequestLine) {
				foundAt = k
				break
			}
		}

		if foundAt != -1 {
			// Found a match
			// Warn about skipped requests in log2
			for k := j; k < foundAt; k++ {
				fmt.Printf("Warning: Extra request in new log:   %s\n\n", log2[k].RequestLine)
			}

			// Compare e1 and log2[foundAt]
			diffs := diffJSON(e1.ResponseBody, log2[foundAt].ResponseBody)
			if len(diffs) > 0 {
				fmt.Printf("Request: %s\n", e1.RequestLine)
				for _, d := range diffs {
					fmt.Printf("  %s\n", d)
				}
				fmt.Println()
			}

			j = foundAt + 1
		} else {
			// Not found in remaining log2
			fmt.Printf("Warning: Request missing in new log: %s\n", e1.RequestLine)
		}
	}

	// Warn about remaining requests in log2
	for k := j; k < len(log2); k++ {
		fmt.Printf("Warning: Extra request in new log: %s\n", log2[k].RequestLine)
	}
}

func areRequestLinesSimilar(l1, l2 string) bool {
	if l1 == l2 {
		return true
	}
	return stripQuery(l1) == stripQuery(l2)
}

func stripQuery(line string) string {
	parts := strings.Split(line, " ")
	for i, part := range parts {
		if strings.Contains(part, "://") {
			if before, _, ok := strings.Cut(part, "?"); ok {
				parts[i] = before
			}
		}
	}
	s := strings.Join(parts, " ")
	s = strings.TrimSuffix(s, "/")
	return s
}

func diffJSON(j1, j2 string) []string {
	var v1, v2 interface{}

	if err := json.Unmarshal([]byte(j1), &v1); err != nil {
		// If not JSON, string compare?
		if j1 != j2 {
			return []string{fmt.Sprintf("Body text changed (not JSON or invalid): %q => %q", truncate(j1), truncate(j2))}
		}
		return nil
	}
	if err := json.Unmarshal([]byte(j2), &v2); err != nil {
		if j1 != j2 {
			return []string{fmt.Sprintf("Body text changed (v2 invalid JSON): %q => %q", truncate(j1), truncate(j2))}
		}
		return nil
	}

	flat1 := flatten(v1)
	flat2 := flatten(v2)

	var diffs []string

	// Keys in flat1
	for k, val1 := range flat1 {
		val2, ok := flat2[k]
		if !ok {
			diffs = append(diffs, fmt.Sprintf("%s: %v => <missing>", k, val1))
		} else if !reflect.DeepEqual(val1, val2) {
			diffs = append(diffs, fmt.Sprintf("%s: %v => %v", k, val1, val2))
		}
	}

	// Keys in flat2 not in flat1
	for k, val2 := range flat2 {
		if _, ok := flat1[k]; !ok {
			diffs = append(diffs, fmt.Sprintf("%s: <missing> => %v", k, val2))
		}
	}

	return diffs
}

func flatten(v interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	flattenRecursive("", v, res)
	return res
}

func flattenRecursive(prefix string, v interface{}, res map[string]interface{}) {
	switch val := v.(type) {
	case map[string]interface{}:
		for k, v2 := range val {
			flattenRecursive(joinPath(prefix, k), v2, res)
		}
	case []interface{}:
		for i, v2 := range val {
			flattenRecursive(fmt.Sprintf("%s[%d]", prefix, i), v2, res)
		}
	default:
		res[prefix] = v
	}
}

func joinPath(prefix, k string) string {
	if prefix == "" {
		return "." + k
	}
	return prefix + "." + k
}

func truncate(s string) string {
	if len(s) > 50 {
		return s[:47] + "..."
	}
	return s
}
