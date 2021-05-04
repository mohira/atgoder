# グリッドに関係する問題

## 問題

- [ABC096C - Grid Repainting 2](https://atcoder.jp/contests/abc096/tasks/abc096_c)
- [ABC075B - Minesweeper](https://atcoder.jp/contests/abc075/tasks/abc075_b)
- [第4回PAST C - 隣接カウント](https://atcoder.jp/contests/past202010-open/tasks/past202010_c?lang=ja)

## ポイント

- 隣接したマスを調べる
    - 上下左右4マス
    - 周囲8マス
    - 周囲9マス(自身+周囲8マス)
- 畳み込み演算
    - 実装負荷が高い

## 隣接したマスのパターン

```go
package main

func main() {
  var H int
  var W int
  // あるマス(i,j)の周囲8マスの表現
  // 軸のとり方はこれがオススメらしい
  // -----------------> y軸
  // | 左上(-1,-1) 真上(-1,0) 右上(-1,1)
  // | 左中( 0,-1) 真中( 0,0) 右中( 0,1)
  // | 左下( 1,-1) 真下( 1,0) 右下( 1,1)
  // ↓
  // x軸

  // 12時から時計回りに差分を書いていく
	var dx = [8]int{-1, -1, 0, 1, 1, 1, 0, -1}
	var dy = [8]int{0, 1, 1, 1, 0, -1, -1, -1}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			// (i,j)の周囲8マスをチェックする
			for k := 0; k < 8; k++ {
				x := i + dx[k]
				y := j + dy[k]

				// マス(x,y) が存在するかのチェック
				if (0 <= x && x < H) && (0 <= y && y < W) {
					// やりたい処理
				}
			}
		}
	}
}
```