package mockkubeapiserver

import (
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
)

type UIDGenerator interface {
	NewUID() types.UID
}

var _ UIDGenerator = &RandomUIDGenerator{}

type RandomUIDGenerator struct {
}

func (c *RandomUIDGenerator) NewUID() types.UID {
	uid := uuid.NewUUID()
	return uid
}

var _ UIDGenerator = &TestUIDGenerator{}

type TestUIDGenerator struct {
	next int64
}

func NewTestUIDGenerator() *TestUIDGenerator {
	return &TestUIDGenerator{next: 0x1}
}

func (c *TestUIDGenerator) NewUID() types.UID {
	v := c.next
	c.next++
	s := fmt.Sprintf("%012x", v)
	s = "00000000-0000-0000-0000-" + s
	return types.UID(s)
}
