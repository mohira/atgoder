package topic_bit_brute_force

import (
	"testing"
)

// [ABC079C - Train Ticket](https://atcoder.jp/contests/abc079/tasks/abc079_c)
func AnswerABC079Cその2(ABCD string) string {
	n := 3

	var ans string
	for bit := 0; bit <= (1 << n); bit++ {
		//fmt.Fprintf(os.Stderr, "%d -> %04b %v\n", bit, bit, bit&1)

		expression := string(ABCD[0])
		sum := int(ABCD[0] - '0')

		for i := 0; i < n; i++ {
			//fmt.Fprintf(os.Stderr, "%04b>>%d=%04b  (%04b>>%d)&1=%v  \n", bit, i, bit>>i, bit, i, (bit>>i)&1)

			if (bit>>i)&1 == 1 {
				expression += "+" + string(ABCD[i+1])
				sum += int(ABCD[i+1] - '0')
			} else {
				expression += "-" + string(ABCD[i+1])
				sum -= int(ABCD[i+1] - '0')
			}
		}

		if sum == 7 {
			ans = expression + "=7"
		}
	}
	return ans
}

func TestAnswerABC079Cその2(t *testing.T) {
	tests := []struct {
		name string
		ABCD string
		want string
	}{
		{"入力例1", "1222", "1+2+2+2=7"},
		{"入力例2", "0290", "0-2+9+0=7"}, // "0−2+9-0=7" でもOK
		{"入力例3", "3242", "3-2+4+2=7"}, // "3-2+4+2=7" でもOK
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC079Cその2(tt.ABCD)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
