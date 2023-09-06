package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PositionModel = (*customPositionModel)(nil)

type (
	// PositionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionModel.
	PositionModel interface {
		positionModel
		InsertAll(ctx context.Context) ([]*Position, error)
	}

	customPositionModel struct {
		*defaultPositionModel
	}
)

// NewPositionModel returns a model for the database table.
func NewPositionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PositionModel {
	return &customPositionModel{
		defaultPositionModel: newPositionModel(conn, c, opts...),
	}
}

func (m *customPositionModel) InsertAll(ctx context.Context) ([]*Position, error) {
	var resp []*Position
	query := fmt.Sprintf("select * from %s", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}
