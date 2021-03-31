package tsm1

import (
	"context"

	"github.com/EnnioRC/influxdb/tsdb"
)

func (e *Engine) CreateCursorIterator(ctx context.Context) (tsdb.CursorIterator, error) {
	return &cursorIterator{e: e}, nil
}
