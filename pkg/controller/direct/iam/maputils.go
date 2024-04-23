package iam

type MapContext struct {
}

func (c *MapContext) Err() error {
	return nil
}

func LazyPtr[T comparable](t T) *T {
	var defaultT T
	if t == defaultT {
		return nil
	}
	return &t
}
