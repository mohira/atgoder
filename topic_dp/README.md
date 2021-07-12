# DPまとめ

## ABC129C

### 考察: 床が壊れていないケースを考える

- [AtCoder ABC 129 C - Typical Stairs (茶色, 300 点) - けんちょんの競プロ精進記録](https://drken1215.hatenablog.com/entry/2019/06/10/140000)
  <img width="2048" alt="image.png (5.6 MB)" src="https://img.esa.io/uploads/production/attachments/6586/2021/07/12/21054/4dd9c584-6987-487d-9769-ce518c499c5d.png">

### 壊れているケースを考える

- 壊れている情報を優先させる感じで実装すると楽(ifがややこしくならない)
- 前の床を気にする場合の実装は漸化式が見えにくいと思う

```go
// 前の床を気にする場合のイメージ
if oks[i-1] { dp[i] += dp[i-1]}
if oks[i-2] { dp[i] += dp[i-2]}
```

<img width="2048" alt="image.png (7.1 MB)" src="https://img.esa.io/uploads/production/attachments/6586/2021/07/12/21054/6604eaeb-34ce-4442-92d4-b93c0c22c7a0.png">
