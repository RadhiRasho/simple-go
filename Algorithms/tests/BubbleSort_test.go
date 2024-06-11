package sorting

import (
	S "Algorithms/sorting"
	"reflect"
	"testing"
)

type Test struct {
	name string
	arr []int
	want []int
}

func TestBubbleSort(t *testing.T) {
	tests := []Test{
		{
			name: "sorted list",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse list",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "empty list",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "one element",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "random list",
			arr:  []int{3, 1, 4, 1, 5, 9, 2, 6},
			want: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			S.BubbleSort(test.arr)
			if !reflect.DeepEqual(test.arr, test.want) {
				t.Errorf("BubbleSort() = %v, want %v", test.arr, test.want)
			}
		})
	}
}