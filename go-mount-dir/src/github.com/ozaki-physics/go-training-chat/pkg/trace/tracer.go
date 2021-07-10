package trace

import (
	"fmt"
	"io"
)

// コード内の出来事を記録できるオブジェクトを表すインタフェース
type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

// 引数を文字列に変換し out フィールドの Writeメソッドに渡す
// fmt.Sprint() は string 型を返すため []byte 型に変換してる
func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct {}

// 何も処理を行わない Trace メソッドを nilTracer struct に属させる
func (t *nilTracer) Trace(a ...interface{}) {}

// Off は Trace メソッドの呼び出しを無視する Tracer を返す
func Off() Tracer {
	return &nilTracer{}
}
