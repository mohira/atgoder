package topic_dp

import (
	"fmt"
	"testing"
)

func AnswerEDPCAFrog1(N int, H []int) int {
	dp := make([]int, N)

	// 初期条件
	dp[0] = 0
	dp[1] = AbsInt(H[1] - H[0])

	for i := 2; i < N; i++ {
		step1 := dp[i-1] + AbsInt(H[i]-H[i-1]) // i+1 の足場へジャンプ
		step2 := dp[i-2] + AbsInt(H[i]-H[i-2]) // i+2 の足場へジャンプ
		cost := Min(step1, step2)
		dp[i] = cost
	}

	return dp[N-1]
}

// https://qiita.com/drken/items/dc53c683d6de8aeacf5a#%E5%AE%9F%E9%9A%9B%E3%81%AB%E6%89%8B%E3%82%92%E5%8B%95%E3%81%8B%E3%81%97%E3%81%A6-dp-%E3%82%92%E4%BD%93%E9%A8%93%E3%81%97%E3%81%A6%E3%81%BF%E3%82%8B
func 愚直に書いてDPに慣れる() {
	N := 7
	H := []int{2, 9, 4, 5, 1, 6, 10}
	dp := make([]int, N)

	dp[0] = 0 // ゼロ値があるけど、わざとらしくかいとく
	fmt.Println(dp)

	dp[1] = H[1] - H[0]
	fmt.Println(dp)

	step1to3 := dp[1] + AbsInt(H[2]-H[1])
	step2to3 := dp[0] + AbsInt(H[2]-H[0])
	dp[2] = Min(step1to3, step2to3)
	fmt.Println(step1to3, step2to3)
	fmt.Println(dp)

	step1to4 := dp[2] + AbsInt(H[3]-H[2])
	step2to4 := dp[1] + AbsInt(H[3]-H[1])
	dp[3] = Min(step1to4, step2to4)
	fmt.Println(step1to4, step2to4)
	fmt.Println(dp)

	step1to5 := dp[3] + AbsInt(H[4]-H[3])
	step2to5 := dp[2] + AbsInt(H[4]-H[2])
	dp[4] = Min(step1to5, step2to5)
	fmt.Println(step1to5, step2to5)
	fmt.Println(dp)

	step1to6 := dp[4] + AbsInt(H[5]-H[4])
	step2to6 := dp[3] + AbsInt(H[5]-H[3])
	dp[5] = Min(step1to6, step2to6)
	fmt.Println(step1to6, step2to6)
	fmt.Println(dp)

	step1to7 := dp[5] + AbsInt(H[6]-H[5])
	step2to7 := dp[4] + AbsInt(H[6]-H[4])
	dp[6] = Min(step1to7, step2to7)
	fmt.Println(step1to7, step2to7)
	fmt.Println(dp)
}

func TestEDPCAFrog1(t *testing.T) {

	tests := []struct {
		name string
		N    int
		H    []int
		want int
	}{
		{"入力例1", 4, []int{10, 30, 40, 20}, 30},
		{"入力例2", 2, []int{10, 10}, 0},
		{"入力例3", 6, []int{30, 10, 60, 10, 60, 50}, 40},
		{"けんちょん本", 7, []int{2, 9, 4, 5, 1, 6, 10}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerEDPCAFrog1(tt.N, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})
	}
}
