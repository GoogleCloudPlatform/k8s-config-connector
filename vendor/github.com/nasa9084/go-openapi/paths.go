package openapi

import (
	"strings"
)

// codebeat:disable[TOO_MANY_IVARS]

// Paths Object
type Paths map[string]*PathItem

// Validate the values of Paths object.
func (paths Paths) Validate() error {
	for path, pathItem := range paths {
		if !strings.HasPrefix(path, "/") {
			return ErrPathFormat
		}
		if err := pathItem.Validate(); err != nil {
			return err
		}
	}
	if paths.hasDuplicatedOperationID() {
		return ErrOperationIDDuplicated
	}
	if paths.hasDuplicatedPaths() {
		return ErrPathsDuplicated
	}
	return nil
}

func (paths Paths) hasDuplicatedOperationID() bool {
	opIDs := map[string]struct{}{}
	for _, pathItem := range paths {
		for _, op := range pathItem.Operations() {
			if _, ok := opIDs[op.OperationID]; ok {
				return true
			}
			opIDs[op.OperationID] = struct{}{}
		}
	}

	return false
}

func (paths Paths) hasDuplicatedPaths() bool {
	pathList := make([]string, len(paths))
	var i int
	for path := range paths {
		pathList[i] = path
		i++
	}
	for i := 0; i < len(pathList); i++ {
		for j := i + 1; j < len(pathList); j++ {
			if isIdenticalPath(pathList[i], pathList[j]) {
				return true
			}
		}
	}
	return false
}

func isIdenticalPath(path1, path2 string) bool {
	p1 := strings.Split(path1, "/")[1:]
	p2 := strings.Split(path2, "/")[1:]
	if len(p1) != len(p2) {
		return false
	}
	for i := range p1 {
		if strings.HasPrefix(p1[i], "{") && strings.HasSuffix(p1[i], "}") && strings.HasPrefix(p2[i], "{") && strings.HasSuffix(p2[i], "}") {
			// both templated
			continue
		}
		if p1[i] == p2[i] {
			continue
		}
		return false
	}
	return true
}

// GetOperationByID returns an operation by operationId.
// If the paths object has two or more operations which matches
// given operationId, this function returns the operation
// matched first. So you should call Validate() before using this
// function.
func (paths Paths) GetOperationByID(operationID string) *Operation {
	for _, pathItem := range paths {
		for _, op := range pathItem.Operations() {
			if op.OperationID == operationID {
				return op
			}
		}
	}
	return nil
}
