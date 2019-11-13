package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name           string
		args           args
		wantSortedList []int
	}{
		{"T0", args{[]int{}, []int{2}}, []int{2}},
		{"T1", args{[]int{1}, []int{2}}, []int{1,2}},
		{"T2", args{[]int{1,2}, []int{3}}, []int{1,2,3}},
		{"T3", args{[]int{1,3,5}, []int{3,4}}, []int{1,3,3,4,5}},
		{"T4", args{[]int{1,3,5}, []int{2,4}}, []int{1,2,3,4,5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSortedList := Merge(tt.args.A, tt.args.B); !reflect.DeepEqual(gotSortedList, tt.wantSortedList) {
				t.Errorf("Merge() = %v, want %v", gotSortedList, tt.wantSortedList)
			}
		})
	}
}

func TestMSort(t *testing.T) {
	type args struct {
		A []int
		n int
	}
	tests := []struct {
		name           string
		args           args
		wantSortedList []int
	}{
		{"T0", args{[]int{}, 0}, []int{}},
		{"T1", args{[]int{1}, 1}, []int{1}},
		{"T2-0", args{[]int{1,2}, 2}, []int{1,2}},
		{"T2-1", args{[]int{2,1}, 2}, []int{1,2}},
		{"T2-2", args{[]int{2,2}, 2}, []int{2,2}},
		{"T3", args{[]int{5,6,4,1,2,3}, 6}, []int{1,2,3,4,5,6}},
		{"T4", args{[]int{1,1,1,1}, 4}, []int{1,1,1,1}},
		{"T5", args{[]int{5,5,5,5,5}, 5}, []int{5,5,5,5,5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSortedList := MSort(tt.args.A, tt.args.n); !reflect.DeepEqual(gotSortedList, tt.wantSortedList) {
				t.Errorf("MSort() = %v, want %v", gotSortedList, tt.wantSortedList)
			}
		})
	}
}