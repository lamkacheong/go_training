
package service

import (
	"context"
	"week4/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	pb "week4/api/drama/v1"
)

type DramaService struct {
	pb.UnimplementedDramaServer

	uc *biz.DramaUserCase
	log *log.Helper
}

func NewDramaService(uc *biz.DramaUserCase, logger log.Logger) *DramaService {
	return &DramaService{uc: uc, log: log.NewHelper(logger)}
}

func (s *DramaService) CreateDrama(ctx context.Context, req *pb.CreateDramaRequest) (*pb.CreateDramaReply, error) {
	drama := biz.Drama{
		Title: req.Title,
		Author: req.Author,
		Category: req.Category,
	}
	err := s.uc.Create(ctx, &drama)
	return &pb.CreateDramaReply{Drama: &pb.DramaMessage{
		Id: int64(drama.Id),
		Title: drama.Title,
		Author: drama.Author,
		Category: drama.Category,
	}}, err
}

func (s *DramaService) UpdateDrama(ctx context.Context, req *pb.UpdateDramaRequest) (*pb.UpdateDramaReply, error) {
	return &pb.UpdateDramaReply{}, nil
}
func (s *DramaService) DeleteDrama(ctx context.Context, req *pb.DeleteDramaRequest) (*pb.DeleteDramaReply, error) {
	return &pb.DeleteDramaReply{}, nil
}
func (s *DramaService) GetDrama(ctx context.Context, req *pb.GetDramaRequest) (*pb.GetDramaReply, error) {
	return &pb.GetDramaReply{}, nil
}
func (s *DramaService) ListDrama(ctx context.Context, req *pb.ListDramaRequest) (*pb.ListDramaReply, error) {
	dramas, err := s.uc.List(ctx)
	replay := &pb.ListDramaReply{}
	if err != nil {
		for _, drama := range *dramas{
			replay.Results = append(replay.Results, &pb.DramaMessage{
				Id: int64(drama.Id),
				Title: drama.Title,
				Author: drama.Author,
				Category: drama.Category,
			})
		}
	}else{
		err = errors.Wrap(err, "list drama error")
	}

	return replay, err
}
