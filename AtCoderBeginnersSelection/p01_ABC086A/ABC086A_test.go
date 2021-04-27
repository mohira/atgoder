package p01_ABC086A

import "testing"

// [ABC086A - Product](https://atcoder.jp/contests/abc086/tasks/abc086_a)
func AnswerABC086A(a int, b int) string {
	product := a * b

	if product%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func TestAnswerABC086A(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want string
	}{
		{"入力例1", 3, 4, "Even"},
		{"入力例2", 1, 21, "Odd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC086A(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}