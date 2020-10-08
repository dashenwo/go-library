package empty_test

import (
	"testing"
)

type TestPerson interface {
	Say() string
}
type TestWoman struct {
}

func (woman TestWoman) Say() string {
	return "nice"
}

func TestIsEmpty(t *testing.T) {

}
