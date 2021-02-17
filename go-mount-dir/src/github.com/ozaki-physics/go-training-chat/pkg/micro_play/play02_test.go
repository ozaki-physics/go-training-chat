package micro_play

import (
	"fmt"
	"testing"
)

func TestPlay02_Subtraction(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, -1},
		{5, 2, -3},
		{-10, -2, -8},
	}
	for idx, pattern := range testCases {
		t.Run(fmt.Sprintf("いけた?"), func(t *testing.T) {
			ans:= Play02_Subtraction(pattern.a, pattern.b)
			// 本当は err の戻り値も付けたいが自作メソッドに実装していない
			// if err != nil {
				// t.Fatal("エラーだよ! エラーが無(nil)じゃないから")
			// }
			if ans != pattern.expected {
				t.Errorf("%d番目が期待した値と違うみたいよ 期待は%d 実際は%d", idx, pattern.expected, ans)
			}
		})
	}
}

// /pkg にて go test -v ./micro_play を実行
