package pairsum

import (
	"reflect"
	"testing"
)

func TestPairSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "target greater than any number",
			nums:   []int{1, 2, 3, 4, 5},
			target: 100,
			want:   nil, // or []int{} depending on your implementation
		},
		{
			name:   "target less than smallest number",
			nums:   []int{1, 2, 3, 4, 5},
			target: -1,
			want:   nil, // or []int{} depending on your implementation
		},
		{
			name:   "target matches first 2",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   []int{1, 2}, // or []int{} depending on your implementation
		},
		{
			name:   "target matches last 2",
			nums:   []int{1, 2, 3, 4, 5},
			target: 9,
			want:   []int{4, 5}, // or []int{} depending on your implementation
		},
		{
			name:   "target matches any 2 numbers",
			nums:   []int{1, 2, 3, 4, 5},
			target: 6,
			want:   []int{1, 5}, // or []int{} depending on your implementation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PairSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
