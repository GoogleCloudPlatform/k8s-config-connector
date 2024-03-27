package mockkubeapiserver

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Clock interface {
	Now() metav1.Time
}

var _ Clock = &RealClock{}

type RealClock struct {
}

func (c *RealClock) Now() metav1.Time {
	return metav1.Now()
}

var _ Clock = &TestClock{}

type TestClock struct {
	t time.Time
}

func NewTestClock() *TestClock {
	t := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	return &TestClock{t: t}
}
func (c *TestClock) Now() metav1.Time {
	t := c.t
	c.t = t.Add(time.Second)
	return metav1.NewTime(t)
}
