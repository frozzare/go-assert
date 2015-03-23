package assert

import "testing"

func TestEqual(t *testing.T) {
	//Equal(t, "foo", "foo", "msg!")
	Equal(t, "foo", "boo", "Foo will not match boo")
}

func TestTrue(t *testing.T) {
	True(t, true, "true will match true")
}

func TestFalse(t *testing.T) {
	False(t, true, "false will match false")
}

func TestNotNil(t *testing.T) {
	NotNil(t, nil, "true is not nil")
}
