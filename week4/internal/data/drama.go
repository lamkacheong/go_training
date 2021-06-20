package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"week4/internal/biz"
)

type dramaRepo struct {
	data *Data
	log  *log.Helper
}

// NewDramaRepo .
func NewDramaRepo(data *Data, logger log.Logger) biz.DramaRepo {
	return &dramaRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *dramaRepo) CreateDrama(ctx context.Context, g *biz.Drama) error {

	result := r.data.db.Create(&g)

	return result.Error
}

func (r *dramaRepo) UpdateDrama(ctx context.Context, g *biz.Drama) error {
	return nil
}

func (r *dramaRepo) ListDrama(ctx context.Context) (*[]biz.Drama, error) {
	dramas := &[]biz.Drama{}
	result := r.data.db.Find(dramas)
	return dramas,result.Error
}