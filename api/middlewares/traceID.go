package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}

type traceIDKey struct{}

// コンテキストにログ追跡用IDを追加します
func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// コンテキストからログ追跡用IDを取得します
func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}
