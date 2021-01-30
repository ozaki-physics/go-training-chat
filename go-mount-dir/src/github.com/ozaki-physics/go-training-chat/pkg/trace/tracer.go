package trace

import (
	"io"
)

// コード内の出来事を記録できるオブジェクトを表すインタフェース
type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return nil
}
