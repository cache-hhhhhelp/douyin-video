package logic

import (
	"context"
	"github.com/cache-hhhhhelp/douyin-video/internal/svc"
	"github.com/cache-hhhhhelp/douyin-video/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *pb.DouyinPublishListRequest) (*pb.DouyinPublishListResponse, error) {

	rawVideos, err := l.svcCtx.VideoModel.FindByAuthor(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	// todo: add other infos
	return &pb.DouyinPublishListResponse{
		VideoList: l.svcCtx.Util.FillVideos(rawVideos),
	}, nil
}
