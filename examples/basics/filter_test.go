package examples

import (
	"reflect"
	"testing"
)

func TestFilterEvenNumbers(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	got := filter(arr, func(val int) bool {
		return val%2 == 0
	})

	want := []int{2, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}