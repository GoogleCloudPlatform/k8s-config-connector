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

package reportarchetypes

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
)

// Mapping is one --resource of a generate-types invocation: a KRM kind and
// the proto message generate.sh maps it to.
type Mapping struct {
	// Kind is the KRM kind, e.g. "ArtifactRegistryRepository".
	Kind string
	// APIVersion is the KRM group and version, e.g.
	// "artifactregistry.cnrm.cloud.google.com/v1beta1".
	APIVersion string
	// Services are the proto packages passed to --service, in order.
	Services []string
	// ProtoName is the message half of --resource as written: relative to a
	// service ("Repository", or "Parent.Child" for a nested message), or
	// fully qualified ("google.api.MetricDescriptor").
	ProtoName string
	// Source is where the invocation starts, relative to the apis directory,
	// e.g. "artifactregistry/generate.sh:36".
	Source string
	// ProtoSource names the descriptor set the invocation reads, from the
	// script text alone: the base name of its --proto-source-path as
	// written, e.g. "googleapis-${PROTO_SHA}.pb", or "HEAD" when it passes
	// no path but the script ran generate-proto.sh HEAD before it. It is
	// empty for the default descriptor set, which generate-proto.sh builds
	// for the googleapis version pinned in apis/git.versions.
	ProtoSource string
}

// ScanGenerateScripts returns every resource mapped by a generate-types
// invocation in apisDir/*/generate.sh, in script order and then file order.
//
// Invocations are recognised by a "generate-types" word, which covers both
// ${CONTROLLERBUILDER} generate-types and go run . generate-types. They read
// --service, --api-version and --resource, or a --config file, the way
// generate-types does, and --proto-source-path for Mapping.ProtoSource.
// ${REPO_ROOT} in a --config path is taken to be the parent of apisDir.
func ScanGenerateScripts(apisDir string) ([]Mapping, error) {
	scripts, err := filepath.Glob(filepath.Join(apisDir, "*", "generate.sh"))
	if err != nil {
		return nil, err
	}
	sort.Strings(scripts)
	repoRoot := filepath.Dir(filepath.Clean(apisDir))

	var mappings []Mapping
	for _, script := range scripts {
		b, err := os.ReadFile(script)
		if err != nil {
			return nil, err
		}
		rel, err := filepath.Rel(apisDir, script)
		if err != nil {
			return nil, err
		}
		found, err := parseScript(string(b), filepath.ToSlash(rel), repoRoot)
		if err != nil {
			return nil, err
		}
		mappings = append(mappings, found...)
	}
	return mappings, nil
}

// parseScript returns the resources mapped by the generate-types invocations
// in one script. source names the script in errors and in Mapping.Source.
func parseScript(src, source, repoRoot string) ([]Mapping, error) {
	var mappings []Mapping
	// vars holds the variables assigned so far, and builtHEAD whether
	// generate-proto.sh HEAD has run; both feed Mapping.ProtoSource.
	vars := map[string]string{}
	builtHEAD := false
	for _, cmd := range shellCommands(src) {
		if name, value, ok := assignment(cmd.words); ok {
			vars[name] = value
			continue
		}
		if buildsHEAD(cmd.words) {
			builtHEAD = true
			continue
		}
		at := slices.Index(cmd.words, "generate-types")
		if at < 0 {
			continue
		}
		where := fmt.Sprintf("%s:%d", source, cmd.line)
		inv, err := parseInvocation(cmd.words[at+1:], repoRoot)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", where, err)
		}
		switch {
		case len(inv.services) == 0:
			return nil, fmt.Errorf("%s: generate-types has no --service", where)
		case inv.apiVersion == "":
			return nil, fmt.Errorf("%s: generate-types has no --api-version", where)
		case len(inv.resources) == 0:
			return nil, fmt.Errorf("%s: generate-types has no --resource", where)
		}
		descriptors := protoSource(inv.protoSourcePath, vars, builtHEAD)
		for _, r := range inv.resources {
			mappings = append(mappings, Mapping{
				Kind:        r.Kind,
				APIVersion:  inv.apiVersion,
				Services:    inv.services,
				ProtoName:   r.ProtoName,
				Source:      where,
				ProtoSource: descriptors,
			})
		}
	}
	return mappings, nil
}

// invocation holds the generate-types flags the report needs.
type invocation struct {
	services        []string
	apiVersion      string
	resources       []options.Resource
	protoSourcePath string
}

// parseInvocation reads the arguments after "generate-types". Flags it does
// not know are skipped; a value following one of them is skipped as a stray
// argument.
func parseInvocation(args []string, repoRoot string) (*invocation, error) {
	inv := &invocation{}
	var config string
	for i := 0; i < len(args); i++ {
		name, value, hasValue := strings.Cut(args[i], "=")
		switch name {
		case "--service", "-s", "--api-version", "-v", "--resource", "--config", "--proto-source-path":
		default:
			continue
		}
		if !hasValue {
			if i+1 >= len(args) {
				return nil, fmt.Errorf("flag %s has no value", name)
			}
			i++
			value = args[i]
		}
		switch name {
		case "--service", "-s":
			inv.services = splitServices(value)
		case "--api-version", "-v":
			inv.apiVersion = value
		case "--resource":
			var r options.Resource
			if err := r.Set(value); err != nil {
				return nil, err
			}
			inv.resources = append(inv.resources, r)
		case "--config":
			config = value
		case "--proto-source-path":
			inv.protoSourcePath = value
		}
	}
	if config != "" {
		if err := inv.applyConfig(config, repoRoot); err != nil {
			return nil, err
		}
	}
	return inv, nil
}

// applyConfig merges a --config file the way generate-types does: the file's
// service and apiVersion replace the flags, and its resources follow any
// given by --resource.
func (inv *invocation) applyConfig(configPath, repoRoot string) error {
	configPath = strings.NewReplacer("${REPO_ROOT}", repoRoot, "$REPO_ROOT", repoRoot).Replace(configPath)
	if strings.Contains(configPath, "$") {
		return fmt.Errorf("cannot expand --config path %q", configPath)
	}
	if !filepath.IsAbs(configPath) {
		// generate.sh runs the controllerbuilder from its own directory.
		configPath = filepath.Join(repoRoot, "dev", "tools", "controllerbuilder", configPath)
	}
	config, err := codegen.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("loading --config %s: %w", configPath, err)
	}
	inv.services = splitServices(config.Service)
	inv.apiVersion = config.APIVersion
	for _, r := range config.Resources {
		inv.resources = append(inv.resources, options.Resource{Kind: r.Kind, ProtoName: r.ProtoName})
	}
	return nil
}

var (
	// shellAssignment matches a word that assigns a shell variable.
	shellAssignment = regexp.MustCompile(`^([A-Za-z_][A-Za-z0-9_]*)=(.*)$`)
	// shellVariable matches a word that is just a variable, $NAME or ${NAME}.
	shellVariable = regexp.MustCompile(`^\$(?:([A-Za-z_][A-Za-z0-9_]*)|\{([A-Za-z_][A-Za-z0-9_]*)\})$`)
)

// assignment returns the variable and value that a command assigns, when
// the command is a single NAME=value word such as
// PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb". The value is
// kept as written: quotes removed, variables unexpanded.
func assignment(words []string) (name, value string, ok bool) {
	if len(words) != 1 {
		return "", "", false
	}
	m := shellAssignment.FindStringSubmatch(words[0])
	if m == nil {
		return "", "", false
	}
	return m[1], m[2], true
}

// buildsHEAD reports whether a command is generate-proto.sh HEAD with no
// output path, which rebuilds the default descriptor set from googleapis
// HEAD instead of the pinned version.
func buildsHEAD(words []string) bool {
	for i, w := range words {
		if path.Base(w) == "generate-proto.sh" {
			return len(words) == i+2 && words[i+1] == "HEAD"
		}
	}
	return false
}

// protoSource returns the Mapping.ProtoSource of an invocation whose
// --proto-source-path is flagValue, or empty when it passes none. A
// flagValue that is just a variable the script assigned earlier stands for
// the value assigned, as written. Nothing else is expanded, so the result
// depends on the script text alone.
func protoSource(flagValue string, vars map[string]string, builtHEAD bool) string {
	if flagValue == "" {
		if builtHEAD {
			return "HEAD"
		}
		return ""
	}
	if m := shellVariable.FindStringSubmatch(flagValue); m != nil {
		if value, ok := vars[m[1]+m[2]]; ok {
			flagValue = value
		}
	}
	return path.Base(flagValue)
}

func splitServices(s string) []string {
	var services []string
	for _, svc := range strings.Split(s, ",") {
		if svc = strings.TrimSpace(svc); svc != "" {
			services = append(services, svc)
		}
	}
	return services
}

// shellCommand is one simple command of a shell script.
type shellCommand struct {
	// words are the command's words, with quotes removed and variables left
	// unexpanded.
	words []string
	// line is the 1-based line the command starts on.
	line int
}

// shellCommands splits a shell script into simple commands. It knows only as
// much shell as generate.sh needs:
//   - a backslash before a newline continues the line;
//   - single and double quotes group a word, and a backslash escapes the next
//     character;
//   - an unquoted # at the start of a word comments out the rest of its line.
//     As in bash, a backslash at the end of a comment does not continue it,
//     so a commented-out line ends the command it sits in;
//   - newlines, ';', '&' and '|' separate commands.
func shellCommands(src string) []shellCommand {
	var (
		cmds      []shellCommand
		words     []string
		word      strings.Builder
		inWord    bool
		line      = 1
		startLine int
	)
	startWord := func() {
		if !inWord {
			inWord = true
			if len(words) == 0 {
				startLine = line
			}
		}
	}
	endWord := func() {
		if inWord {
			words = append(words, word.String())
			word.Reset()
			inWord = false
		}
	}
	endCommand := func() {
		endWord()
		if len(words) > 0 {
			cmds = append(cmds, shellCommand{words: words, line: startLine})
		}
		words = nil
	}

	for i := 0; i < len(src); i++ {
		c := src[i]
		switch c {
		case '\\':
			if i+1 < len(src) {
				i++
				if src[i] == '\n' {
					line++
					continue
				}
				startWord()
				word.WriteByte(src[i])
			}
		case '\'':
			startWord()
			for i++; i < len(src) && src[i] != '\''; i++ {
				if src[i] == '\n' {
					line++
				}
				word.WriteByte(src[i])
			}
		case '"':
			startWord()
			for i++; i < len(src) && src[i] != '"'; i++ {
				if src[i] == '\\' && i+1 < len(src) && strings.IndexByte("\"\\$`\n", src[i+1]) >= 0 {
					i++
					if src[i] == '\n' {
						line++
						continue
					}
				} else if src[i] == '\n' {
					line++
				}
				word.WriteByte(src[i])
			}
		case '#':
			if inWord {
				word.WriteByte(c)
				continue
			}
			for i+1 < len(src) && src[i+1] != '\n' {
				i++
			}
		case '\n':
			endCommand()
			line++
		case ';', '&', '|':
			endCommand()
		case ' ', '\t', '\r':
			endWord()
		default:
			startWord()
			word.WriteByte(c)
		}
	}
	endCommand()
	return cmds
}
