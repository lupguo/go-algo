package heap

import (
	"reflect"
	"testing"
)

func TestMinHeap_Insert(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name   string
		fields *MinHeap
		args   args
		want   []int
	}{
		{"T1", NewHeap(10),
			args{[]int{1, 5, 4, 3, 2, 6, 7, 8, 9}},
			[]int{1, 2, 4, 5, 3, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{
				Cap:   tt.fields.Cap,
				Len:   tt.fields.Len,
				Elems: tt.fields.Elems,
			}
			for _, v := range tt.args.x {
				h.Insert(v)
				h.Print()
			}
			// deep check
			if got := h.Elems[1 : h.Len+1]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Delete(t *testing.T) {
	tests := []struct {
		name    string
		fields  *MinHeap
		delElem int
		want    []int
	}{
		{"T3", &MinHeap{10, 9, []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9}}, 1, []int{2, 4, 3, 9, 5, 6, 7, 8},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{
				Cap:   tt.fields.Cap,
				Len:   tt.fields.Len,
				Elems: tt.fields.Elems,
			}
			if _ = h.Delete(); !reflect.DeepEqual(h.Elems[1:h.Len+1], tt.want) {
				t.Errorf("Delete() = %v, want %v", h.Elems[1:h.Len+1], tt.want)
			}
		})
	}
}

func TestMinHeap_Print(t *testing.T) {
	tests := []struct {
		name string
		heap *MinHeap
	}{
		{"T0", &MinHeap{10, 3, []int{-1, 1, 2, 3}}},
		{"T1", &MinHeap{10, 9, []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.Print()
		})
	}
}
