package apis

const (
	// State into spec annotation values
	StateMergeIntoSpec               = "merge"
	StateAbsentInSpec                = "absent"
	StateIntoSpecDefaultValueV1Beta1 = StateAbsentInSpec
)

var StateIntoSpecAnnotationValues = []string{
	StateMergeIntoSpec,
	StateAbsentInSpec,
}
