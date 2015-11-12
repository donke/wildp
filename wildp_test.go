package wildp

import (
	"reflect"
	"testing"
)

func TestExpand(t *testing.T) {
	actual := expand([]string{"a", "b", "c", "[!w]*.go", "[a-z]ild?_*"})
	expected := []string{"a", "b", "c", "wildp_test.go"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v", actual)
	}
}
