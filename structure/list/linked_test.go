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
		{"T2", Generate(2)},
		{"T3", Generate(1)},
		{"T4", Generate(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s: %s\n", tt.name, tt.list)
		})
	}
}

func TestGenerateSorted(t *testing.T) {
	tests := []struct {
		name     string
		list     *ListNode
	}{
		{"T1", GenerateSorted(10)},
		{"T2", GenerateSorted(2)},
		{"T3", GenerateSorted(1)},
		{"T4", GenerateSorted(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s: %s\n", tt.name, tt.list)
		})
	}
}