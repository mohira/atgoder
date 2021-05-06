# Go言語でのビット全探索

## 基本パターン

```go
package main

func main() {
  // 2^N 個のループ
  for bit := 0; bit < (1 << N); bit++ {

    // ON-OFFを対応させる
    for i := 0; i < N; i++ {
      // i番目のフラグがYESかどうかチェック
      if bit&(1<<i) > 0 { // (bit>>i) & 1 == 1 でもOK
        // なんかの処理
      }
    }
  }
}
```

## 例題

> [部分和問題]
> N個の正の整数a0,a1,...,aN－1と正の整数Wが与えられます。 a0, a1 , ... , aN－1の中から何個かの整数を選んで総和をWとすることができるかどうかを判定してください。
>
> 問題解決力を鍛える！アルゴリズムとデータ構造 p.36

### コード例

```go
package topic_bit_brute_force

import (
	"testing"
)

// [部分和問題]
// N個の正の整数a0,a1,...,aN－1と正の整数Wが与えられます。
// a0,a1,...,aN－1の中から何個かの整数を選んで総和をWとすることができるかどうかを判定してください。
// 大槻兼資. 問題解決力を鍛える！アルゴリズムとデータ構造 (Japanese Edition) (Kindle の位置No.911-916). Kindle 版.

func Answer部分和(N int, W int, A []int) string {
	var can bool

	for bit := 0; bit < (1 << N); bit++ {
		sum := 0
		for i := 0; i < N; i++ {
			if bit&(1<<i) > 0 {
				sum += A[i]
			}
		}
		if sum == W {
			can = true
		}
	}

	if can {
		return "YES"
	} else {
		return "NO"
	}
}

func TestAnswer部分和(t *testing.T) {
	tests := []struct {
		name string
		N    int
		W    int
		A    []int
		want string
	}{
		{"入力例1", 3, 5, []int{1, 2, 3}, "YES"},
		{"入力例2", 5, 10, []int{1, 2, 4, 5, 11}, "YES"},
		{"入力例3", 4, 10, []int{1, 5, 8, 11}, "NO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Answer部分和(tt.N, tt.W, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
```

## 理解のポイントは5つ

1. ビット列(2進数) と 10進数の変換
2. 左シフト演算結果を2進数と10進数で理解する
3. `1<<N`: `2^N`通りの探索
4. `1 << a`: a桁目(a番目)のフラグが立っているか
5. `bit & (1 << i) > 0`: その要素を含めるのかどうか

## 1. ビット列(2進数) と 10進数の変換

これはかんたん。

## 2. 左シフト演算結果を2進数と10進数で理解する

- `1 << i`
    - `1 << 0`  -> `0001` == `1`
    - `1 << 1`  -> `0010` == `2`
    - `1 << 2`  -> `0100` == `4`
    - `1 << 3`  -> `1000` == `8`
- `2 << i`
    - `2 << 0` -> `00010` = `2`
    - `2 << 1` -> `00100` = `4`
    - `2 << 2` -> `01000` = `8`
    - `2 << 3` -> `10000` = `16`
- `3 << i`
    - `3 << 0` -> `00011` = `3`
    - `3 << 1` -> `00110` = `6`
    - `3 << 2` -> `01100` = `12`
    - `3 << 3` -> `11000` = `24`

### code

```go
package main

import "fmt"

func main() {
	// x << i という 左シフト演算 の動きを確認する
	var x = 3

	for i := 0; i <= 3; i++ {
		fmt.Printf("`%d << %d` -> `%05b` = `%d`\n", x, i, x<<i, x<<i)
	}
}
```

## 3. `1<<N`: `2^N`通りの探索

```go
package main

import "fmt"

func main() {
	// 1 << N は 2^N 通り これをループで表現すると O(2^N)
	var N = 5 // 数列の長さとかそういうの
	for bit := 0; bit < (1 << N); bit++ {
		fmt.Printf("%02d = %08b\n", bit, bit)
	}
}
```

## 4. `1 << a`: a桁目(a番目)のフラグが立っているか

これを数列なり何なりに対応させればよい

- `1 << 0` -> `00000001`
- `1 << 1` -> `00000010`
- `1 << 2` -> `00000100`
- `1 << 3` -> `00001000`
- `1 << 4` -> `00010000`
- `1 << 5` -> `00100000`
- `1 << 6` -> `01000000`
- `1 << 7` -> `10000000`

```go
package main

import "fmt"

func main() {
	// 1 << a は a番目(a桁目)のフラグのみが立っていることを示す
	for a := 0; a <= 7; a++ {
		fmt.Printf("`1 << %d` -> `%08b`\n", a, 1<<a)
	}
}
```

## 5. `bit & (1 << i) > 0`: その要素を含めるのかどうか

- 「部分集合として採用するか？」や「買うか買わないか？」や「ONかOFFか？」の条件判定で使える

```go
package main

import "fmt"

func main() {
	// 例: 部分集合をすべて取り出す
	var A = []int{1, 2, 3}

	var N = len(A) // 数列の長さとかそういうの
	for bit := 0; bit < (1 << N); bit++ {
		fmt.Printf("%04b -> ", bit)
		var subset []int

		// 0番目からN-1番目(==N個)それぞれのフラグチェック
		// 例えば、「部分集合をつくるときに、その要素を採用するか？」みたいなときに使える
      for i := 0; i < N; i++ {
        if (bit & (1 << i)) > 0 {
          // 要素に含める処理
          subset = append(subset, A[i])
        }
      }
      fmt.Println(subset)
    }
}
```

### `(bit>>i)&1 == 1`: もっとスッキリした書き方

- https://youtu.be/ENSOy8u9K9I?t=2163 を参考にするとよい

