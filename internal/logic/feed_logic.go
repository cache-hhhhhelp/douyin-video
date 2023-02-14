package logic

import (
	"context"
	"math"
	"time"

	"github.com/cache-hhhhhelp/douyin-video/internal/svc"
	"github.com/cache-hhhhhelp/douyin-video/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *pb.DouyinFeedRequest) (*pb.DouyinFeedResponse, error) {

	t := int64(time.Now().Second())
	if in.LatestTime != nil {
		t = *(in.LatestTime)
	}
	// get videos
	rawVideos, err := l.svcCtx.VideoModel.ListAllByTimeDesc(l.ctx, t, 30)
	if err != nil {
		return nil, err
	}
	// get next time
	nextTime := int64(math.MaxInt64)
	for _, video := range rawVideos {
		if nextTime > video.CreatedAt {
			nextTime = video.CreatedAt
		}
	}
	// todo: fill other infos

	return &pb.DouyinFeedResponse{
		VideoList: l.svcCtx.Util.FillVideos(rawVideos),
		NextTime:  &nextTime,
	}, nil
}
