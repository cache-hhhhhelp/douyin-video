package svc

import (
	"douyin-video/internal/config"
	"douyin-video/internal/logic"
	"douyin-video/internal/model"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"net/http"
	"net/url"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
	CosClient  cos.Client
	Util       logic.Util
}

func getCosClient(c config.Config) *cos.Client {
	u, _ := url.Parse("https://" + c.COS.BucketName + "-" + c.COS.AppId + ".cos.ap-" + c.COS.Region + ".myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.COS.SecretId,
			SecretKey: c.COS.SecretKey,
		},
	})
}

func NewServiceContext(c config.Config) *ServiceContext {
	cosClient := *getCosClient(c)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(sqlx.NewMysql(c.Mysql.Datasource), c.Cache),
		CosClient:  cosClient,
		Util:       logic.NewUtilImpl(cosClient),
	}
}
