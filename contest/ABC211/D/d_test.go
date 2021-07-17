package D

import (
	"testing"
)

func AnswerD(N, K int) int {
	return 0
}

func TestD(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		want int
	}{
		{"ex1", 13, 2, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerD(tt.N, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
