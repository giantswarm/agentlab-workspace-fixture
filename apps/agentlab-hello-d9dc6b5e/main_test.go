package main

import "testing"

func TestGreeting(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "returns expected greeting",
			want: "Hello from agentlab reference run d9dc6b5e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greeting(); got != tt.want {
				t.Errorf("greeting() = %q, want %q", got, tt.want)
			}
		})
	}
}
