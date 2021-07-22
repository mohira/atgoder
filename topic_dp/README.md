# DPまとめ

## ABC129C

- [AtCoder ABC 129 C - Typical Stairs (茶色, 300 点) - けんちょんの競プロ精進記録](https://drken1215.hatenablog.com/entry/2019/06/10/140000)

### 考察: 床が壊れていないケースを考える

![](https://i.imgur.com/Gu9LSQb.jpg)

### 壊れているケースを考える

- 壊れている情報を優先させる感じで実装すると楽(ifがややこしくならない)
- 前の床を気にする場合の実装は漸化式が見えにくいと思う

```go
// 前の床を気にする場合のイメージ
if oks[i-1] { dp[i] += dp[i-1]}
if oks[i-2] { dp[i] += dp[i-2]}
```

![](https://i.imgur.com/lZXk8es.jpg)

## [EDPC D - Knapsack 1](https://atcoder.jp/contests/dp/tasks/dp_d)

<img width="1630" alt="image.png (3.3 MB)" src="https://img.esa.io/uploads/production/attachments/6586/2021/07/15/21054/895e282e-b3f1-484f-bd66-2757ec971984.png">

## [EDPC F - LCS](https://atcoder.jp/contests/dp/tasks/dp_f)

- 参考
    - [【EDPC】F問題 LCS【全部やる】 - YouTube](https://www.youtube.com/watch?v=HKI3aTJz4LY)
    - [(1) 情報工学概論（アルゴリズムとデータ構造）09動的計画法02編集距離 - YouTube](https://www.youtube.com/watch?v=VLC6h3GzSYo&list=PLCGPvL9AseS3AcyAYp8UL_CWoIhKyQv2I&index=4)
        - 編集距離の話だけど、愚直にDPテーブルを作って遷移を調べるところはとても参考になる
- 漸化式とテーブルの紐付けに苦労した
    - 具体的なDPテーブルを書き出すことをやるとよい感じ
    - 復元するときにもDPテーブルがある方がわかりやすい

```
// コードから入っていかないほうがいいかも？
// この漸化式はわかるけど、わからない状態になりがちじゃない？
dp[i][j] = max(
    dp[i-1][j],
    dp[i][j-1],
    dp[i-1][j-1]+1,
)
```

### 図解

![](https://i.imgur.com/TBwlg2D.jpg)

## [EPDC G - Longest Path](https://atcoder.jp/contests/dp/tasks/dp_g)

- 参考
  - [【EDPC】G問題 LongestPath【全部やる】 - YouTube](https://www.youtube.com/watch?v=U5geMnL9gGU)
- 発想
  - 自分から見て、「次の頂点」の最長経路 + 1 を再帰的に求める
  - グラフ情報は隣接リストでもっておくとわかりやすい
- メモ化再帰の使い所
  - 良い順番はありそう(今回はDAG)だが、**どの順番で埋めるかわからん**とき！

### 図解

![](https://i.imgur.com/X1WZiR4.jpg)

