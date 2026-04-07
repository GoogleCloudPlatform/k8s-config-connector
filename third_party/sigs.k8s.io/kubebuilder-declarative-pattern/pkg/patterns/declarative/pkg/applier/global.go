//go:build !without_direct_applier
// +build !without_direct_applier

package applier

var DefaultApplier = NewDirectApplier()
