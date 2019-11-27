package list

import (
	"testing"
)

func TestNode_Print(t *testing.T) {
	tests := []struct {
		name string
		list *ListNode
	}{
		{"T1", Generate(10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s: %s\n", tt.name, tt.list)
		})
	}
}
