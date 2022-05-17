package single

import "testing"

func TestCounter(t *testing.T) {
	array := []int64{17, 15, 34, 1003, 2, 129359, 800, 9282}
	var result int64 = 4
	var multipleOf int64 = 17

	var testResult int64 = CountMultiples(multipleOf, array)

	if testResult != result {
		t.Error(testResult, " != ", result)
	}
}
