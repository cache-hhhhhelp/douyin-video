package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		FindByAuthor(ctx context.Context, authorId int64) ([]Video, error)
		ListAllByTimeDesc(ctx context.Context, time int64, num int64) ([]Video, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

func (c customVideoModel) ListAllByTimeDesc(ctx context.Context, time int64, num int64) ([]Video, error) {
	var resp []Video
	query := fmt.Sprintf("select %s from %s where `created_at` < ? order by `created_at` desc limit ?", videoRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, time, num)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customVideoModel) FindByAuthor(ctx context.Context, authorId int64) ([]Video, error) {
	var resp []Video
	query := fmt.Sprintf("select %s from %s where `author_id` = ?", videoRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, authorId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c),
	}
}
