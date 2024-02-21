package diffs

type FieldPath struct {
	parent *FieldPath
	part   string
}

func (f *FieldPath) With(part string) *FieldPath {
	return &FieldPath{parent: f, part: part}
}

func (f *FieldPath) asSlice() []string {
	n := 0
	pos := f
	for pos != nil {
		pos = pos.parent
		n++
	}
	ret := make([]string, n)
	i := n - 1
	pos = f
	for i >= 0 {
		ret[i] = pos.part
		pos = pos.parent
		i--
	}
	return ret
}
