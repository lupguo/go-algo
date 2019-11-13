package sort

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name           string
		args           args
		wantSortedList []int
	}{
		{"T0", args{[]int{}}, []int{}},
		{"T1", args{[]int{1}}, []int{1}},
		{"T2", args{[]int{1, 2}}, []int{1, 2}},
		{"T3-0", args{[]int{1, 2, 3}}, []int{1, 2, 3}},
		{"T3-1", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"T4", args{[]int{4, 3, 2, 1}}, []int{1, 2, 3,4}},
		{"T5", args{[]int{5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5}},
		{"T9", args{[]int{7, 6, 8, 9, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSortedList := Bubble(tt.args.A, len(tt.args.A)); !reflect.DeepEqual(gotSortedList, tt.wantSortedList) {
				t.Errorf("Bubble() = %v, want %v", gotSortedList, tt.wantSortedList)
			}
		})
	}
}
