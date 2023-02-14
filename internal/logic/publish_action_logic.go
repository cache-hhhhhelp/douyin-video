package logic

import (
	"bytes"
	"context"
	"douyin-video/internal/model"
	"douyin-video/internal/svc"
	"douyin-video/types/pb"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishActionLogic) uploadVideo(data []byte, key string) error {
	fileReader := bytes.NewReader(data)
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/html",
		},
	}
	_, err := l.svcCtx.CosClient.Object.Put(context.Background(), key, fileReader, opt)
	return err
}

func (l *PublishActionLogic) PublishAction(in *pb.DouyinPublishActionRequest) (*pb.DouyinPublishActionResponse, error) {

	video := model.Video{
		AuthorId:  in.UserId,
		Title:     in.Title,
		CreatedAt: time.Now().Unix(),
	}
	// insert metadata into database
	sqlResult, err := l.svcCtx.VideoModel.Insert(l.ctx, &video)
	if err != nil {
		return nil, err
	}
	// get video id and upload file
	id, err := sqlResult.LastInsertId()
	videoKey := strconv.FormatInt(id, 10)
	err = l.uploadVideo(in.Data, videoKey)

	return &pb.DouyinPublishActionResponse{}, err
}
