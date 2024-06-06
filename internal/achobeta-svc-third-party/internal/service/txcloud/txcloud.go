package txcloud

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-third-party/config"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

const (
	// 存储桶
	Url = "https://achobeta-common-1257999694.cos.ap-guangzhou.myqcloud.com"
)

var c *cos.Client

func New() {
	if c == nil {
		u, _ := url.Parse(Url)
		b := &cos.BaseURL{BucketURL: u}
		c = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  config.Get().Txcloud.SecretID,
				SecretKey: config.Get().Txcloud.SecretKey,
			},
		})
	}
}

func PutObject(object io.Reader, objectName string) error {
	tlog.Infof("put object : %s", objectName)
	_, err := c.Object.Put(context.Background(), objectName, object, nil)
	if err != nil {
		tlog.CtxErrorf(context.Background(), "put object error: %v", err)
		return err
	}
	return nil
}
