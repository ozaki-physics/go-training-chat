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
		tracer.Trace("hello! trace パッケージ")
		if buf.String() != "hello! trace パッケージ\n" {
			t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
		}
	}
}
