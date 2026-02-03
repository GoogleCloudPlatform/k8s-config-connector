// Copyright 2025 Google LLC
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

package preview

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

type ReconcileStatus int

// ReconcileStatus is the status of a GKNN object after being reconciled.
const (
	ReconcileStatusIdle ReconcileStatus = iota
	ReconcileStatusHealthy
	ReconcileStatusUnhealthy
)

// reconcileStatusMap is a map of ReconcileStatus to string.
var reconcileStatusMap = map[ReconcileStatus]string{
	ReconcileStatusHealthy:   "RECONCILE_STATUS_HEALTHY",
	ReconcileStatusUnhealthy: "RECONCILE_STATUS_UNHEALTHY",
}

func (s ReconcileStatus) String() string {
	return reconcileStatusMap[s]
}

// GKNNReconciledResult is the result of reconciling a GKNN object with a specific controller type.
type GKNNReconciledResult struct {
	GKNN            GKNN
	ControllerType  k8s.ReconcilerType
	ReconcileStatus ReconcileStatus
	Diffs           *structuredreporting.Diff
}

// FormatGKNNReconciledResult formats the GKNNReconciledResult into a string.
func (r *GKNNReconciledResult) FormatGKNNReconciledResult() string {
	return fmt.Sprintf("ns=\"%s\" name=\"%s\" group=\"%s\" kind=\"%s\" controller_type=\"%s\" diffs=\"%s\" reconcile_status=\"%s\"", r.GKNN.Namespace, r.GKNN.Name, r.GKNN.Group, r.GKNN.Kind, r.ControllerType, FormatFieldIDs(r.Diffs), r.ReconcileStatus.String())
}

func FormatFieldIDs(diffs *structuredreporting.Diff) string {
	return strings.Join(diffs.FieldIDs(), ",")
}

type PreviewReport struct {
	report map[GKNN]*GKNNReport
}

// GKNNReport is the result of reconciling a GKNN object with both default and alternative controller type.
type GKNNReport struct {
	DefaultResult     *GKNNReconciledResult
	AlternativeResult *GKNNReconciledResult
}

func (r *Recorder) GenerateGKNNReconciledResults() map[GKNN]*GKNNReconciledResult {
	result := make(map[GKNN]*GKNNReconciledResult)
	for gknn := range r.objects {
		result[gknn] = &GKNNReconciledResult{
			GKNN:            gknn,
			Diffs:           &structuredreporting.Diff{},
			ReconcileStatus: ReconcileStatusHealthy,
		}
		for _, event := range r.objects[gknn].events {
			switch event.eventType {
			case EventTypeDiff:
				result[gknn].Diffs.AddDiff(event.diff)
			case EventTypeReconcileStart:
				result[gknn].ControllerType = event.reconcilerType
			case EventTypeReconcileEnd:
				result[gknn].ControllerType = event.reconcilerType
			case EventTypeKubeAction:
				// Ignore for now
			case EventTypeGCPAction:
				result[gknn].ReconcileStatus = ReconcileStatusUnhealthy
			default:
				// Ignore for now
			}
		}
	}
	return result
}

// NewPreviewReport creates a new PreviewReport from the given default and alternative recorders.
func NewPreviewReport(defaultRecorder *Recorder, alternativeRecorder *Recorder) *PreviewReport {
	defaultReconciledResults := defaultRecorder.GenerateGKNNReconciledResults()
	alternativeReconciledResults := alternativeRecorder.GenerateGKNNReconciledResults()

	previewReport := &PreviewReport{report: make(map[GKNN]*GKNNReport)}
	previewReport.Initialize(defaultReconciledResults)
	previewReport.Initialize(alternativeReconciledResults)
	reconcilerConfig := resourceconfig.LoadConfig()
	for gknn := range defaultReconciledResults {
		previewReport.report[gknn].DefaultResult = defaultReconciledResults[gknn]
	}
	for gknn := range alternativeReconciledResults {
		alternativeResult := alternativeReconciledResults[gknn]
		if alternativeResult.ControllerType != reconcilerConfig[schema.GroupKind{Group: gknn.Group, Kind: gknn.Kind}].DefaultController {
			previewReport.report[gknn].AlternativeResult = alternativeResult
		}
	}

	return previewReport
}

// Initialize the report with the given recorder.
func (r *PreviewReport) Initialize(results map[GKNN]*GKNNReconciledResult) {
	for gknn := range results {
		r.report[gknn] = &GKNNReport{}
	}
}

// ExportSummary exports a summary of the report to a file.
func (r *PreviewReport) ExportSummary(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", filename, err)
	}
	w := tabwriter.NewWriter(f, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "GROUP\tKIND\tNAMESPACE\tNAME\tDEFAULT CONTROLLER\tDEFAULT STATUS\tALTERNATIVE CONTROLLER\tALTERNATIVE STATUS")
	for gknn, report := range r.report {
		defaultReconciledStatus := formatReconciledStatus(report.DefaultResult)
		defaultControllerType := formatControllerType(report.DefaultResult)
		alternativeReconciledStatus := formatReconciledStatus(report.AlternativeResult)
		alternativeControllerType := formatControllerType(report.AlternativeResult)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", gknn.Group, gknn.Kind, gknn.Namespace, gknn.Name, defaultControllerType, defaultReconciledStatus, alternativeControllerType, alternativeReconciledStatus)
	}
	w.Flush()
	return f.Close()
}

func (r *PreviewReport) ExportFailedResults(filename string) error {
	details := r.GenerateDetailFailedResults()
	if len(details) == 0 {
		return nil
	}
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", filename, err)
	}
	for _, detail := range details {
		fmt.Fprintln(f, detail)
	}
	return f.Close()
}

func (r *PreviewReport) GenerateDetailFailedResults() []string {
	details := []string{}
	for _, report := range r.GetFailures() {
		details = append(details, report.FormatGKNNReconciledResult())
	}
	return details
}

func (r *PreviewReport) GetFailures() []*GKNNReconciledResult {
	result := []*GKNNReconciledResult{}
	for _, report := range r.report {
		defaultResult := report.DefaultResult
		if defaultResult != nil && defaultResult.ReconcileStatus == ReconcileStatusUnhealthy {
			result = append(result, defaultResult)
		}
		alternativeResult := report.AlternativeResult
		if alternativeResult != nil && alternativeResult.ReconcileStatus == ReconcileStatusUnhealthy {
			result = append(result, alternativeResult)
		}
	}
	return result
}

func formatControllerType(result *GKNNReconciledResult) string {
	if result == nil {
		return "N/A"
	}
	return string(result.ControllerType)
}

func formatReconciledStatus(result *GKNNReconciledResult) string {
	if result == nil {
		return "N/A"
	}
	if result.ReconcileStatus == ReconcileStatusUnhealthy {
		return "UNHEALTHY"
	}
	return "HEALTHY"
}

// ExportObjectsEvent writes all captured GKNN and its event to filename.
func (r *Recorder) ExportDetailObjectsEvent(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", filename, err)
	}
	fmt.Fprintf(f, "Total number of resources: %d\n", len(r.objects))
	fmt.Fprintf(f, "Number of resources that have not been fully reconciled: %d\n", r.RemainResourcesCount)
	if r.RemainResourcesCount != 0 {
		fmt.Fprintln(f, "Known Resource that was not reconciled")
		for gknn, reconciled := range r.ReconciledResources {
			if !reconciled {
				fmt.Fprintln(f, "-----------------------------------------------------------------")
				fmt.Fprintf(f, "Not reconciled object %+v\n", gknn)
			}
		}
	}
	for gknn, info := range r.objects {
		fmt.Fprintln(f, "-----------------------------------------------------------------")
		fmt.Fprintf(f, "object %+v\n", gknn)
		for _, event := range info.events {
			switch event.eventType {
			case EventTypeDiff:
				fmt.Fprintf(f, "  diff %+v\n", event.diff)
				for _, field := range event.diff.Fields {
					fmt.Fprintf(f, "  field %s\n", field.ID)
					fmt.Fprintf(f, "    old %+v\n", field.Old)
					fmt.Fprintf(f, "    new %+v\n", field.New)
				}

			case EventTypeReconcileStart:
				fmt.Fprintf(f, "  reconcileStart %+v\n", event.object)

			case EventTypeReconcileEnd:
				fmt.Fprintf(f, "  reconcileEnd %+v\n", event.object)

			case EventTypeKubeAction:
				fmt.Fprintf(f, "  kubeAction %+v\n", event.kubeAction)

			case EventTypeGCPAction:
				fmt.Fprintf(f, "  gcpAction %+v\n", event.gcpAction)

			default:
				fmt.Fprintf(f, "  unknown event: %+v\n", event)
			}
		}
	}
	return f.Close()
}
