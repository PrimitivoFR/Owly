package common_testing

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// CmpAssertEqual allows to test if two objects are equal, using go-cmp.
// opts refer to options to pass to the cmp.Equal() function.
// Can be stuff like cmpopts.IgnoreUnexported() or cmpopts.EquateErrors()
func CmpAssertEqual(t *testing.T, want interface{}, test interface{}, opts ...cmp.Option) {
	if o := cmp.Equal(want, test, opts...); o == false {
		t.Errorf(
			"Assertion failed:\n expected\t %v of (%v)\n got\t\t %v (%v)",
			want, reflect.TypeOf(want), test, reflect.TypeOf(test),
		)
	}
}
