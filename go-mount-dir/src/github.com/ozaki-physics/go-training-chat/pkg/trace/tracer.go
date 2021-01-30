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
