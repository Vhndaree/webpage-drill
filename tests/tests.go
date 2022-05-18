package tests

import (
	"reflect"
	"testing"
)

func ExpectEqual(t *testing.T, check, expected interface{}) bool {
	if !reflect.DeepEqual(check, expected) {
		t.Helper()
		t.Errorf("Expected %s, actual %s", expected, check)
		return false
	}
	return true
}
