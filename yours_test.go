package assert

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	// assert equality
	//a.Equal(t, 123, 124, "they should be equal")

	// assert inequality
	a.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	// a.Nil(t, object)

	// assert for not nil (good when you expect something)
	// if a.NotNil(t, object) {

	// now we know that object isn't nil, we are safe to make
	// further assertions without causing any errors
	// a.Equal(t, "Something", object.Value)

	// }

}
