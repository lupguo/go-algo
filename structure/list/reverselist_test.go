package list

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
	}{
		{"T1", Generate(1)},
		{"T2", Generate(2)},
		{"T3", Generate(5)},
		{"T4", Generate(10)},
		{"T5", Generate(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("origin list: %s", tt.head)
			t.Logf("reverse list: %s", ReverseList(tt.head))
		})
	}
}
