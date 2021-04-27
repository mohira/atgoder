package similar

import (
	"fmt"
	"os"
	"sort"
	"testing"
)

// [ABC082B - Two Anagrams](https://atcoder.jp/contests/abc082/tasks/abc082_b)
// 定義に近いかたちでの実装
func AnswerABC082Bその1(s string, t string) string {
	var N, M int
	N = len(s)
	M = len(t)

	var a, b []int
	a = sortCharactersAsc(s)
	b = sortCharactersDesc(t)

	for i := 0; i < N; i++ {
		// ある`i`の存在条件の確認を満たしていない場合はループを抜ける
		if M-1 < i {
			break
		}

		ai := a[i]
		bi := b[i]
		_, _ = fmt.Fprintf(os.Stderr, "a%d='%s' b%d='%s'\n", i, string(int32(ai)), i, string(int32(bi)))

		if ai < bi {
			return "Yes"
		}
		if bi < ai {
			return "No"
		}
	}

	if N < M {
		return "Yes"
	}

	return "No"
}

func sortCharactersDesc(s string) []int {
	var b = make([]int, len(s))

	for i, c := range s {
		b[i] = int(c)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	return b
}

func sortCharactersAsc(s string) []int {
	var a = make([]int, len(s))

	for i, c := range s {
		a[i] = int(c)
	}

	sort.Ints(a)

	return a
}

func TestABC082Bその1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		// N < M の場合
		{"入力例1", "yx", "axy", "Yes"},
		{"入力例4", "w", "ww", "Yes"},
		{"入力例3", "cd", "abc", "No"},

		// N = M の場合
		{"入力例5: 同じ文字列はNo", "zzz", "zzz", "No"},

		// N > M の場合
		{"入力例2", "ratcode", "atlas", "Yes"},
		{"N>Mかつ同一文字の連続はNo", "www", "ww", "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC082Bその1(tt.s, tt.t)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
