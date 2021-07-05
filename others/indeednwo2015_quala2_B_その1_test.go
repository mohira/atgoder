package others

import (
	"reflect"
	"strings"
	"testing"
)

func AnswerIndeedQuala2Bその1(N int, S []string) string {
	var out []string

	bucket := make(map[rune]int)
	for _, c := range "indeednow" {
		bucket[c]++
	}

	for _, s := range S {
		tmpBucket := make(map[rune]int)
		for _, c := range s {
			tmpBucket[c]++
		}

		if reflect.DeepEqual(bucket, tmpBucket) {
			out = append(out, "YES")
		} else {
			out = append(out, "NO")
		}
	}

	return strings.Join(out, "\n")
}

func TestAnswerIndeedQuala2Bその1(t *testing.T) {
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
			got := AnswerIndeedQuala2Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}

}
