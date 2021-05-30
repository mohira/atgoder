package topic_amari

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

type Student struct {
	Id     int
	People int // 登校時にいた生徒の人数(自身を含む)
}

type Students []Student

func (ps Students) Less(i, j int) bool {
	if ps[i].People < ps[j].People {
		return true
	}
	return false
}

func (ps Students) Len() int {
	return len(ps)
}
func (ps Students) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// [ABC142C - Go to School](https://atcoder.jp/contests/abc142/tasks/abc142_c)
func AnswerABC142Cその1(N int, A []int) string {
	students := make(Students, 0, N)
	for i, a := range A {
		students = append(students, Student{i + 1, a})
	}

	// 人数でソートする
	sort.Stable(students)

	var s = make([]string, 0, N)
	for _, student := range students {
		s = append(s, strconv.Itoa(student.Id))
	}

	return strings.Join(s, " ")
}

func TestAnswerABC142Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 3, []int{2, 3, 1}, "3 1 2"},
		{"入力例1", 5, []int{1, 2, 3, 4, 5}, "1 2 3 4 5"},
		{"入力例1", 8, []int{8, 2, 7, 3, 4, 5, 6, 1}, "8 2 4 5 6 7 3 1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC142Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
