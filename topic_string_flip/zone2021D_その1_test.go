package topic_bit_brute_force

import (
	"container/list"
	"strings"
	"testing"
)

// [zone2021D - 宇宙人からのメッセージ](https://atcoder.jp/contests/zone2021/tasks/zone2021_d)
func AnswerZone2021Dその1(S string) string {
	var flipped bool
	strList := list.New()

	for _, s := range strings.Split(S, "") {
		if s == "R" {
			flipped = !flipped
		} else {
			if !flipped {
				if strList.Back() != nil && strList.Back().Value == s {
					strList.Remove(strList.Back())
				} else {
					strList.PushBack(s)
				}
			} else {
				if strList.Front() != nil && strList.Front().Value == s {
					strList.Remove(strList.Front())
				} else {
					strList.PushFront(s)
				}
			}
		}
	}

	var ans string

	if !flipped {
		for e := strList.Front(); e != nil; e = e.Next() {
			ans += e.Value.(string)
			// fmt.Print(e.Value) // 提出時
		}
	} else {
		for e := strList.Back(); e != nil; e = e.Prev() {
			ans += e.Value.(string)
			// fmt.Print(e.Value) // 提出時
		}
	}

	return ans
}

func TestAnswerZone2021Dその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "ozRnonnoe", "zone"},
		{"入力例2", "hellospaceRhellospace", ""},
		{"2回反転", "abRcdRe", "dcabe"},
		{"2回反転かつ消失あり", "abRadRe", "dbe"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerZone2021Dその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
