package list

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"T1", args{GenerateSorted(1), GenerateSorted(1)}},
		{"T2", args{GenerateSorted(1), GenerateSorted(2)}},
		{"T3", args{GenerateSorted(2), GenerateSorted(1)}},
		{"T4", args{GenerateSorted(1), GenerateSorted(0)}},
		{"T5", args{GenerateSorted(0), GenerateSorted(1)}},
		{"T6", args{GenerateSorted(5), GenerateSorted(5)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("origin l1: %s, l2: %s", tt.args.l1, tt.args.l2)
			t.Logf("after merge: %s", MergeTwoLists(tt.args.l1, tt.args.l2))
		})
	}
}
