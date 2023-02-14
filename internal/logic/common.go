package logic

import (
	"github.com/cache-hhhhhelp/douyin-video/internal/model"
	"github.com/cache-hhhhhelp/douyin-video/types/pb"
	"github.com/tencentyun/cos-go-sdk-v5"
	"strconv"
)

var _ Util = (*UtilImpl)(nil)

type (
	Util interface {
		FillVideos(rawVideos []model.Video) []*pb.Video
	}
	UtilImpl struct {
		cos.Client
		Util
	}
)

func (u *UtilImpl) getUrl(key string) string {
	return u.Object.GetObjectURL(key).String()
}

func (u *UtilImpl) FillVideos(rawVideos []model.Video) []*pb.Video {
	videos := make([]*pb.Video, len(rawVideos))
	for i := 0; i < len(videos); i++ {
		videoKey := strconv.FormatInt(rawVideos[i].Id, 10)
		videos[i] = &pb.Video{
			Id:       rawVideos[i].Id,
			AuthorId: rawVideos[i].AuthorId,
			PlayUrl:  u.getUrl(videoKey),
			CoverUrl: u.getUrl("cover_" + videoKey),
			Title:    rawVideos[i].Title,
		}
	}
	return videos
}

func NewUtilImpl(client cos.Client) *UtilImpl {
	return &UtilImpl{
		Client: client,
	}
}
