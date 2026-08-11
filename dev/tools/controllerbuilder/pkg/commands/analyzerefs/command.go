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

// Package analyzerefs reports how well GCP protos annotate their reference
// fields with (google.api.resource_reference).
//
// This is a read-only measurement tool. It exists to answer whether the
// annotation is a reliable enough signal to drive automatic Ref generation,
// before we build anything that depends on it.
package analyzerefs

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/spf13/cobra"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Options struct {
	*options.GenerateOptions

	// ServiceFilter matches proto packages by substring, e.g. "aiplatform".
	// Empty means all packages.
	ServiceFilter string

	// List prints every annotated field, not just the summary.
	List bool

	// ListUnannotated prints string fields that look reference-like by
	// heuristic but carry no annotation - the gap an annotation-only
	// strategy would miss.
	ListUnannotated bool
}

// fieldInfo is one string field in the descriptor set.
type fieldInfo struct {
	FullName   string // pkg.Message.field
	Package    string
	TargetType string // from resource_reference; "" if unannotated
	Comment    string
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &Options{GenerateOptions: baseOptions}

	cmd := &cobra.Command{
		Use:   "analyze-refs",
		Short: "Report (google.api.resource_reference) annotation coverage on string fields",
		Long: "Measures how many string fields carry the upstream resource_reference " +
			"annotation, and inventories the target resource types they point at.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), opt)
		},
	}

	cmd.Flags().StringVar(&opt.ServiceFilter, "service", opt.ServiceFilter,
		"only consider proto packages containing this substring (e.g. aiplatform)")
	cmd.Flags().BoolVar(&opt.List, "list", opt.List,
		"list every annotated field and its target type")
	cmd.Flags().BoolVar(&opt.ListUnannotated, "list-unannotated", opt.ListUnannotated,
		"list reference-looking fields that carry NO annotation")

	return cmd
}

func run(_ context.Context, opt *Options) error {
	if opt.ProtoSourcePath == "" {
		return fmt.Errorf("`--proto-source-path` is required")
	}
	api, err := protoapi.LoadProto(opt.ProtoSourcePath, opt.ProtoOverlayPath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	files, err := api.GetAllFileDescriptors()
	if err != nil {
		return fmt.Errorf("getting file descriptors: %w", err)
	}

	var fields []fieldInfo
	for _, f := range files {
		pkg := string(f.Package())
		if opt.ServiceFilter != "" && !strings.Contains(pkg, opt.ServiceFilter) {
			continue
		}
		msgs := f.Messages()
		for i := 0; i < msgs.Len(); i++ {
			collectFields(msgs.Get(i), pkg, &fields)
		}
	}

	report(opt, fields)
	return nil
}

// collectFields walks a message and its nested messages, recording string fields.
func collectFields(msg protoreflect.MessageDescriptor, pkg string, out *[]fieldInfo) {
	f := msg.Fields()
	for i := 0; i < f.Len(); i++ {
		fd := f.Get(i)
		if fd.Kind() != protoreflect.StringKind {
			continue
		}
		info := fieldInfo{
			FullName: string(fd.FullName()),
			Package:  pkg,
			Comment:  leadingComments(fd),
		}
		if rr := resourceReference(fd); rr != nil {
			// ChildType is the "parent of" form; Type is the direct reference.
			if rr.GetType() != "" {
				info.TargetType = rr.GetType()
			} else {
				info.TargetType = rr.GetChildType() + " (child_type)"
			}
		}
		*out = append(*out, info)
	}

	nested := msg.Messages()
	for i := 0; i < nested.Len(); i++ {
		n := nested.Get(i)
		if n.IsMapEntry() {
			continue
		}
		collectFields(n, pkg, out)
	}
}

// leadingComments returns the doc comment for a field, matching how
// codegen/typegenerator.go reads descriptions.
func leadingComments(fd protoreflect.FieldDescriptor) string {
	f := fd.ParentFile()
	if f == nil {
		return ""
	}
	return strings.TrimSpace(f.SourceLocations().ByDescriptor(fd).LeadingComments)
}

func resourceReference(fd protoreflect.FieldDescriptor) *annotations.ResourceReference {
	opts := fd.Options()
	if opts == nil {
		return nil
	}
	ext := proto.GetExtension(opts, annotations.E_ResourceReference)
	rr, ok := ext.(*annotations.ResourceReference)
	if !ok || rr == nil {
		return nil
	}
	if rr.GetType() == "" && rr.GetChildType() == "" {
		return nil
	}
	return rr
}

// looksLikeRef mirrors the heuristics used by tests/apichecks TestMissingRefs,
// plus the URI cases that check currently misses. Used only to size the gap
// that an annotation-only strategy would leave behind.
func looksLikeRef(f fieldInfo) bool {
	name := f.FullName
	if i := strings.LastIndex(name, "."); i >= 0 {
		name = name[i+1:]
	}
	lower := strings.ToLower(name)
	c := f.Comment

	// Match uri/url as a whole snake_case token, not a substring: "security"
	// contains "uri" but is not a URI. Catches uris, input_uri,
	// output_uri_prefix - the GcsSource/GcsDestination shape that both the
	// annotation and the current apicheck miss.
	for _, tok := range strings.Split(lower, "_") {
		switch tok {
		case "uri", "uris", "url", "urls":
			return true
		}
	}

	switch {
	case strings.Contains(c, " projects/"), strings.Contains(c, "projects/{"),
		strings.Contains(c, "locations/{"), strings.Contains(c, "organizations/{"),
		strings.Contains(c, "folders/{"):
		return true
	case strings.Contains(c, "gs://"), strings.Contains(c, "bq://"):
		return true
	// Descriptions frequently name the service in prose without any URI scheme.
	case strings.Contains(c, "Cloud Storage"), strings.Contains(c, "BigQuery"):
		return true
	case strings.Contains(lower, "service_account"), strings.Contains(lower, "kms_key"),
		strings.Contains(lower, "network"), strings.Contains(lower, "subnetwork"):
		return true
	}
	return false
}

func report(opt *Options, fields []fieldInfo) {
	total := len(fields)
	annotated := 0
	byTarget := map[string]int{}
	byPackage := map[string][2]int{} // pkg -> [annotated, total]

	var unannotatedLooksLikeRef []fieldInfo

	for _, f := range fields {
		pc := byPackage[f.Package]
		pc[1]++
		if f.TargetType != "" {
			annotated++
			byTarget[f.TargetType]++
			pc[0]++
		} else if looksLikeRef(f) {
			unannotatedLooksLikeRef = append(unannotatedLooksLikeRef, f)
		}
		byPackage[f.Package] = pc
	}

	fmt.Printf("=== resource_reference coverage ===\n")
	if opt.ServiceFilter != "" {
		fmt.Printf("service filter: %q\n", opt.ServiceFilter)
	}
	fmt.Printf("string fields:      %d\n", total)
	fmt.Printf("annotated:          %d (%s)\n", annotated, pct(annotated, total))
	fmt.Printf("unannotated but\n  heuristically ref-like: %d\n\n", len(unannotatedLooksLikeRef))

	// Per-package, worst coverage first: those are the risky services.
	type pkgRow struct {
		pkg              string
		annotated, total int
	}
	var rows []pkgRow
	for k, v := range byPackage {
		rows = append(rows, pkgRow{k, v[0], v[1]})
	}
	sort.Slice(rows, func(i, j int) bool {
		ri, rj := ratio(rows[i].annotated, rows[i].total), ratio(rows[j].annotated, rows[j].total)
		if ri != rj {
			return ri < rj
		}
		return rows[i].pkg < rows[j].pkg
	})
	fmt.Printf("--- per proto package (lowest coverage first) ---\n")
	shown := 0
	for _, r := range rows {
		if r.total == 0 {
			continue
		}
		fmt.Printf("%-58s %4d/%-4d %s\n", r.pkg, r.annotated, r.total, pct(r.annotated, r.total))
		if shown++; shown >= 25 && opt.ServiceFilter == "" {
			fmt.Printf("... (%d more packages)\n", len(rows)-shown)
			break
		}
	}

	fmt.Printf("\n--- target type inventory ---\n")
	type tRow struct {
		t string
		n int
	}
	var trows []tRow
	for k, v := range byTarget {
		trows = append(trows, tRow{k, v})
	}
	sort.Slice(trows, func(i, j int) bool {
		if trows[i].n != trows[j].n {
			return trows[i].n > trows[j].n
		}
		return trows[i].t < trows[j].t
	})
	for _, r := range trows {
		fmt.Printf("%5d  %s\n", r.n, r.t)
	}

	if opt.List {
		fmt.Printf("\n--- annotated fields ---\n")
		sort.Slice(fields, func(i, j int) bool { return fields[i].FullName < fields[j].FullName })
		for _, f := range fields {
			if f.TargetType != "" {
				fmt.Printf("%s -> %s\n", f.FullName, f.TargetType)
			}
		}
	}

	if opt.ListUnannotated {
		fmt.Printf("\n--- reference-like but UNANNOTATED (annotation-only blind spot) ---\n")
		sort.Slice(unannotatedLooksLikeRef, func(i, j int) bool {
			return unannotatedLooksLikeRef[i].FullName < unannotatedLooksLikeRef[j].FullName
		})
		for _, f := range unannotatedLooksLikeRef {
			fmt.Printf("%s\n", f.FullName)
		}
	}
}

func ratio(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

func pct(a, b int) string {
	if b == 0 {
		return "n/a"
	}
	return fmt.Sprintf("%.1f%%", 100*ratio(a, b))
}
