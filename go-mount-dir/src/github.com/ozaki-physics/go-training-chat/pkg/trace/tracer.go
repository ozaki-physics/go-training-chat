package trace

// コード内の出来事を記録できるオブジェクトを表すインタフェース
type Tracer interface {
	Trace(...interface{})
}
