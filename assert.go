package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// messageFromMsgAndArgs will return the message
// if it exists.
func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		return msgAndArgs[0].(string)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

// equal will test if the expected and actual value is the same match.
func equal(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if reflect.DeepEqual(expected, actual) {
		return true
	}

	if fmt.Sprintf("%#v", expected) == fmt.Sprintf("%#v", actual) {
		return true
	}

	return false
}

// Fail will print a failing message to the terminal.
func Fail(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	message := messageFromMsgAndArgs(msgAndArgs...)
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("\033[31m✖\033[39m %s (%s:%d) \033[31m%v == %v\033[39m\n",
		message,
		filepath.Base(file),
		line,
		expected,
		actual)
}

// Equal will test if the expected and actual value is the same match.
func Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if !equal(expected, actual) {
		Fail(t, expected, actual, msgAndArgs...)
	}

	return true
}

// True will test the actual value and see if it's true or not.
func True(t *testing.T, actual bool, msgAndArgs ...interface{}) bool {
	if actual != true {
		Fail(t, true, actual, msgAndArgs...)
	}

	return true
}

// False will test the actual value and see if it's false or not.
func False(t *testing.T, actual bool, msgAndArgs ...interface{}) bool {
	if actual != false {
		Fail(t, false, actual, msgAndArgs...)
	}

	return true
}

// Nil will test the actual value and see if it's nil or not.
func Nil(t *testing.T, actual interface{}, msgAndArgs ...interface{}) bool {
	success := true

	if actual == nil {
		success = false
	}

	value := reflect.ValueOf(actual)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		success = false
	}

	if !success {
		Fail(t, nil, actual, msgAndArgs...)
	}

	return success
}

// NotNil test check if the actual value is not nil.
func NotNil(t *testing.T, actual interface{}, msgAndArgs ...interface{}) bool {
	success := true

	if actual == nil {
		success = false
	} else {
		value := reflect.ValueOf(actual)
		kind := value.Kind()
		if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
			success = false
		}
	}

	if !success {
		Fail(t, nil, actual, msgAndArgs...)
	}

	return success
}
