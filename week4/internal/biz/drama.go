package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)
type Drama struct {
	gorm.Model
	Id int
	Title string
	Author string
	Category string
}

type DramaRepo interface {
	CreateDrama(context.Context, *Drama) error
	UpdateDrama(context.Context, *Drama) error
	ListDrama(context.Context) (*[]Drama,error)
}

type DramaUserCase struct {
	repo DramaRepo
	log  *log.Helper
}

func NewDramaUserCase(repo DramaRepo, logger log.Logger) *DramaUserCase {
	return &DramaUserCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DramaUserCase) Create(ctx context.Context, g *Drama) error {
	return uc.repo.CreateDrama(ctx, g)
}

func (uc *DramaUserCase) Update(ctx context.Context, g *Drama) error {
	return uc.repo.UpdateDrama(ctx, g)
}

func (uc *DramaUserCase) List(ctx context.Context) (*[]Drama,error) {
	time.Sleep(1 * time.Second)
	return uc.repo.ListDrama(ctx)
}