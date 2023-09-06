package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CompanyModel = (*customCompanyModel)(nil)

type (
	// CompanyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCompanyModel.
	CompanyModel interface {
		companyModel
		FindAll(ctx context.Context) ([]*Company, error)
	}

	customCompanyModel struct {
		*defaultCompanyModel
	}
)

// NewCompanyModel returns a model for the database table.
func NewCompanyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CompanyModel {
	return &customCompanyModel{
		defaultCompanyModel: newCompanyModel(conn, c, opts...),
	}
}

func (m *customCompanyModel) FindAll(ctx context.Context) ([]*Company, error) {

	var resp []*Company
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
