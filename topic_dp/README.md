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
