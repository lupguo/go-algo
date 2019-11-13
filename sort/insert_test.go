package sort

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
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
		{"T2", args{[]int{1,2}}, []int{1,2}},
		{"T3", args{[]int{1,2,3}}, []int{1,2,3}},
		{"T4", args{[]int{3,2,1}}, []int{1,2,3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSortedList := Insert(tt.args.A); !reflect.DeepEqual(gotSortedList, tt.wantSortedList) {
				t.Errorf("Insert() = %v, want %v", gotSortedList, tt.wantSortedList)
			}
		})
	}
}