package ABC_problemC_gray

import (
	"sort"
	"strings"
	"testing"
)

type Words []string

func (ws Words) Len() int {
	return len(ws)
}

func (ws Words) Less(i, j int) bool {
	if ws[i] < ws[j] {
		return true
	}
	return false
}

func (ws Words) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

// strings.Join のまま
func (ws Words) Join(sep string) string {
	switch len(ws) {
	case 0:
		return ""
	case 1:
		return ws[0]
	}

	n := len(sep) * (len(ws) - 1)
	for i := 0; i < len(ws); i++ {
		n += len(ws[i])
	}

	var builder strings.Builder
	builder.Grow(n)

	builder.WriteString(ws[0])
	for _, w := range ws[1:] {
		builder.WriteString(sep)
		builder.WriteString(w)
	}

	return builder.String()
}

type WordMap map[string]int

func (wm WordMap) MaxFreq() int {
	var maxFreq int
	for _, freq := range wm {
		if maxFreq < freq {
			maxFreq = freq
		}
	}

	return maxFreq
}

// [ABC155C - Poll](https://atcoder.jp/contests/abc155/tasks/abc155_c)
func AnswerABC155Cその1(N int, S []string) string {
	wordMap := make(WordMap)
	for _, s := range S {
		wordMap[s]++
	}
	var maxFreq = wordMap.MaxFreq()

	var modeWords = make(Words, 0, len(wordMap))

	for word, freq := range wordMap {
		if freq == maxFreq {
			modeWords = append(modeWords, word)
		}
	}

	// keysをソートして昇順にしてもいい(というか、そのほうがスッキリ)
	sort.Stable(modeWords)

	return modeWords.Join("\n")
}

func TestAnswerABC155Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want string
	}{
		{
			"入力例1",
			7,
			[]string{"beat", "vet", "beet", "bed", "vet", "bet", "beet"},
			"beet\nvet",
		},
		{
			"入力例2",
			8,
			[]string{"buffalo", "buffalo", "buffalo", "buffalo", "buffalo", "buffalo", "buffalo", "buffalo"},
			"buffalo",
		},
		{
			"入力例3",
			7,
			[]string{"bass", "bass", "kick", "kick", "bass", "kick", "kick"},
			"kick",
		},
		{
			"入力例4",
			4,
			[]string{"ushi", "tapu", "nichia", "kun"},
			"kun\nnichia\ntapu\nushi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC155Cその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
