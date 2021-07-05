package others

import (
	"atgoder/lib"
	"strings"
	"testing"
)

func AnswerIndeedQuala2Bその2(N int, S []string) string {
	out := make([]string, N)

	indeednow := lib.MySortString("indeednow")

	for i, s := range S {
		if indeednow == lib.MySortString(s) {
			out[i] = "YES"
		} else {
			out[i] = "NO"
		}
	}

	return strings.Join(out, "\n")
}

func TestAnswerIndeedQuala2Bその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want string
	}{
		{
			"入力例1",
			10,
			[]string{"nowindeed", "indeedwow", "windoneed", "indeednow", "wondeedni", "a", "indonow", "ddeennoiw", "indeednoww", "indeow"},
			"YES\nNO\nYES\nYES\nYES\nNO\nNO\nYES\nNO\nNO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerIndeedQuala2Bその2(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}

}
