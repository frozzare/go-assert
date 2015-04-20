package assert

import "testing"

func TestEqual(t *testing.T) {
	Equal(t, "foo", "foo")

	type Hello struct {
		Name string
	}

	p := Hello{"Fredrik"}
	p2 := Hello{"Kalle"}

	NotEqual(t, p, p2)
}

func TestTrue(t *testing.T) {
	True(t, true)
}

func TestFalse(t *testing.T) {
	False(t, false)
}

func TestNotNil(t *testing.T) {
	NotNil(t, true)
}
