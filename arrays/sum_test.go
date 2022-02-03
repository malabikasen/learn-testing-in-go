package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("run a slice of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5} //fixed size array
		expectedValue := 15
		returnedValue := Sum(numbers)
		if expectedValue != returnedValue {
			t.Errorf("got '%d', wanted '%d'", returnedValue, expectedValue)
		}
	})

}

func TestSumAll(t *testing.T) {
	returnedValue := SumAll([]int{1, 2}, []int{3, 4})
	expectedValue := []int{3, 7}

	if !reflect.DeepEqual(returnedValue, expectedValue) {
		t.Errorf("got %v, wanted %v", expectedValue, returnedValue)
	}
}
