package p08_ABC088C

import (
	"testing"
)

// [ABC088C - Half and Half](https://atcoder.jp/contests/abc088/tasks/arc086_a)
func AnswerABC096Cその1(A, B, C, X, Y int) int {
	C = 2 * C // 半分サイズのピザは意味がないし、*2を忘れるので、ピザ1単位になるように補正

	// それぞれのピザの必要枚数(Cピザは2枚分==Aピザ1枚とBピザ1枚分に相当する)
	var nA, nB, nC int

	if C < A+B {
		// なるべくCピザで賄うほうがお得なケース
		nC = min(X, Y)
		nA = X - nC
		nB = Y - nC

		if C <= A {
			// さらにAピザ分もCピザで賄うべき
			nC += nA
			nA = 0
		}
		if C <= B {
			// さらにBピザ分もCピザで賄うべき
			nC += nB
			nB = 0
		}
	} else {
		// Cピザではゴリ押せないので、Aピザ分とBピザ分をそれぞれCピザとコスパ比較する
		if A < C {
			nA = X
		} else {
			nC += X
		}

		if B < C {
			nB = Y
		} else {
			nC += Y
		}
	}

	return A*nA + B*nB + C*nC
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func TestAnswerABC096Cその1(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		C    int
		X    int
		Y    int
		want int
	}{
		{"入力例1", 1500, 2000, 1600, 3, 2, 7900},
		{"入力例2", 1500, 2000, 1900, 3, 2, 8500},
		{"入力例3", 1500, 2000, 500, 90000, 100000, 100000000},
		{"入力例1", 1500, 2000, 1750, 3, 2, 8500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC096Cその1(tt.A, tt.B, tt.C, tt.X, tt.Y)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
