package main

import "testing"

func TestMessage(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "greeting", want: "Hello from agentlab reference run gzf613ee"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := message(); got != tc.want {
				t.Fatalf("message() = %q, want %q", got, tc.want)
			}
		})
	}
}
