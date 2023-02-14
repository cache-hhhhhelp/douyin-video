package logic

import (
	"douyin-video/internal/model"
	"douyin-video/types/pb"
	"github.com/tencentyun/cos-go-sdk-v5"
	"strconv"
)

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
			Id:            rawVideos[i].Id,
			Author:        nil,
			PlayUrl:       u.getUrl(videoKey),
			CoverUrl:      u.getUrl("cover_" + videoKey),
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         rawVideos[i].Title,
		}
	}
	return videos
}

func NewUtilImpl(client cos.Client) *UtilImpl {
	return &UtilImpl{
		Client: client,
	}
}
