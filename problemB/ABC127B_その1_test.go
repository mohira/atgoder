package problemB

import (
	"fmt"
	"testing"
)

// [ABC127B - Algae](https://atcoder.jp/contests/abc127/tasks/abc127_b)
func AnswerABC127Bその1(r, D, x2000 int) string {
	xi := x2000

	var ans string
	for i := 0; i < 10; i++ {
		xi = r*xi - D
		ans += fmt.Sprintf("%d\n", xi)
	}

	return ans
}

func TestAnswerABC127Bその1(t *testing.T) {
	tests := []struct {
		name        string
		r, D, x2000 int
		want        string
	}{
		{"入力例1", 2, 10, 20, "30\n50\n90\n170\n330\n650\n1290\n2570\n5130\n10250\n"},
		{"入力例2", 4, 40, 60, "200\n760\n3000\n11960\n47800\n191160\n764600\n3058360\n12233400\n48933560\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC127Bその1(tt.r, tt.D, tt.x2000)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
