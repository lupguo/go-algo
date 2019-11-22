package sort

import (
	"reflect"
	"testing"
)

func TestQuick(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name           string
		args           args
		wantSortedList []int
	}{
		{"T0", args{[]int{}}, []int{}},
		{"T1-0", args{[]int{1}}, []int{1}},
		{"T1-1", args{[]int{1,1}}, []int{1,1}},
		{"T2", args{[]int{1, 2}}, []int{1, 2}},
		{"T3", args{[]int{1, 2, 3}}, []int{1, 2, 3}},
		{"T4", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"T5", args{[]int{5, 4, 2, 3, 1}}, []int{1, 2, 3, 4, 5}},
		{"T6", args{[]int{5, 4, 6, 2, 3, 1}}, []int{1, 2, 3, 4, 5, 6}},
		{"T7", args{[]int{5, 5, 3, 2, 1}}, []int{1, 2, 3, 5, 5}},
		{"T8", args{[]int{7, 5, 4, 6, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"T9", args{[]int{1,1,1,1}}, []int{1,1,1,1}},
		{"T10", args{[]int{5,5,5,5,5}}, []int{5,5,5,5,5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSortedList := Quick(tt.args.A, len(tt.args.A)); !reflect.DeepEqual(gotSortedList, tt.wantSortedList) {
				t.Errorf("Quick() = %v, want %v", gotSortedList, tt.wantSortedList)
			}
		})
	}
}

func TestPivot(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"T2", args{[]int{1}}},
		{"T3", args{[]int{1,2}}},
		{"T4", args{[]int{1,2,3}}},
		{"T5", args{[]int{1,2,3,4}}},
		{"T6", args{[]int{1,2,3,4,5}}},
		{"T7", args{[]int{1,2,2,5}}},
		{"T8", args{[]int{1,2,2,5,5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("list %v", tt.args.A)
			Pivot(tt.args.A, len(tt.args.A))
			t.Logf("after pivot of %v\n", tt.args.A, )
		})
	}
}