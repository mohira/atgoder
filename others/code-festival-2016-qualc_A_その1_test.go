package others

import (
	"reflect"
	"testing"
)

func AnswerFestival2016Aその1(s string) string {
	cIdx := -1
	for i, c := range s {
		if c == 'C' {
			cIdx = i
			break
		}
	}
	fIdx := -1
	for i, c := range s {
		if c == 'F' {
			fIdx = i
		}
	}

	if cIdx == -1 || fIdx == -1 {
		return "No"
	}

	if cIdx < fIdx {
		return "Yes"
	} else {
		return "No"
	}
}

func TestCodeFestival2016Aその1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"入力例1", "CODEFESTIVAL", "Yes"},
		{"入力例2", "FESTIVALCODE", "No"},
		{"入力例3", "CF", "Yes"},
		{"入力例4", "FCF", "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerFestival2016Aその1(tt.s)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
