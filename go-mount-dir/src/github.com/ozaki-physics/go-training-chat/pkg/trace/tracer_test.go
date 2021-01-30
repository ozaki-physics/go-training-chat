package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	// 出力されるデータは bytes.Buffer に保持されている
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New からの戻り値が nil です")
	} else {
		// Trace メソッドを使って引数の文字列を tracer に格納する
		tracer.Trace("hello! trace パッケージ")
		// 想定した処理ができているか確認するアサート部分
		if buf.String() != "hello! trace パッケージ\n" {
			t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
		}
	}
}

func TestOff(t *testing) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("データ")
}
