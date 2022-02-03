package main

import "testing"

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
