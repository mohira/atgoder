package ABC_problemB

import (
	"testing"
)

// [ABC168B - ... (Triple Dots)](https://atcoder.jp/contests/abc168/tasks/abc168_b)
func AnswerABC168Bその1(K int, S string) string {
	if len(S) <= K {
		return S
	} else {
		return S[:K] + "..."
	}
}

func TestAnswerABC168Bその1(t *testing.T) {
	tests := []struct {
		name string
		K    int
		S    string
		want string
	}{
		{"入力例1", 7, "nikoandsolstice", "nikoand..."},
		{"入力例1", 40, "ferelibenterhominesidquodvoluntcredunt", "ferelibenterhominesidquodvoluntcredunt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC168Bその1(tt.K, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
