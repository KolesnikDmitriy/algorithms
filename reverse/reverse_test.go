package reverse

import "testing"

func Test_ReverseInt(t *testing.T) {
	arg := 475
	expected := 574
	actual := ReverseInt(arg)
	if actual != expected {
		t.Errorf("Error in ReverseInt(%v): expected %v but actual %v", arg, expected, actual)
	}
}