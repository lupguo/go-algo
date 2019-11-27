package list

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{"T1", args{head: Generate(1), n: 1,}},
		{"T2", args{head: Generate(5), n: 1,}},
		{"T3", args{head: Generate(5), n: 2,}},
		{"T4", args{head: Generate(5), n: 3,}},
		{"T5", args{head: Generate(5), n: 5,}},
		{"T6", args{head: Generate(4), n: 1,}},
		{"T6", args{head: Generate(4), n: 2,}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("origin: %s", tt.args.head)
			t.Logf("after remove %d node: %s",tt.args.n, RemoveNthFromEnd(tt.args.head, tt.args.n))
		})
	}
}
