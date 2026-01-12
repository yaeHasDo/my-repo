package main

import "testing"

// テストコードの実行
// input が１０の場合にEvenが変えることをテストしている
// そうじゃないなら返してね

/**
* テストの書き方
* 関数名は必ずTestから始めること
* testing.T は Go の標準テスト用構造体
 */
func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Even" {
		t.Errorf("expected Even, got %s", result)
	}
}
