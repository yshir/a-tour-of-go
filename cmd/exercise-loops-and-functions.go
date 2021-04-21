package main

import (
	"fmt"
	"math"
)

// 関数とループを使った簡単な練習として、平方根の計算を実装してみましょう: 数値 x が与えられたときに z² が最も x に近い数値 z を求めたいと思います。
// コンピュータは通常ループを使って x の平方根を計算します。 いくつかの z を推測することから始めて、z² がどれほど x に近づいているかに応じて z を調整できます。

// z -= (z*z - x) / (2*z)
// 実際の平方根に近い答えになるまでこの調整を繰り返すことによって、推測はより良いものなります。

// これを func Sqrt に実装してください。 何が入力されても z の適切な開始推測値は 1 です。 まず計算を 10 回繰り返してそれぞれの z を表示します。 x (1, 2, 3, ...) のさまざまな値に対する答えがどれほど近似し、 推測が速くなるかを確認してください。

const N = 10

func Sqrt(x float64) float64 {
	z := 1.0
	seen := map[float64]bool{}

	for i := 0; i < N; i++ {
		diff := (z*z - x) / (2 * z)
		fmt.Println("Search:", i, diff)
		if seen[diff] {
			return z
		}

		seen[diff] = true
		z -= diff
	}

	return z
}

func main() {
	fmt.Println("actual:", Sqrt(2))
	fmt.Println("expected:", math.Sqrt(2))
}
