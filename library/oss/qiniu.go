package oss

import (
	"context"
	"fmt"
	"github.com/fast-crud/fast-auth/library/interfaces"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/fast-crud/fast-auth/library/global"
	"github.com/pkg/errors"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"go.uber.org/zap"
)

var _ interfaces.Oss = (*qiniu)(nil)

var Qiniu = new(qiniu)

type qiniu struct {
	filename string
	filesize int64
}

func (q *qiniu) DeleteByKey(key string) error {
	mac := qbox.NewMac(global.Config.Qiniu.AccessKey, global.Config.Qiniu.SecretKey)
	config := global.Config.Qiniu.GetConfig()
	bucketManager := storage.NewBucketManager(mac, config)
	if err := bucketManager.Delete(global.Config.Qiniu.Bucket, key); err != nil {
		return errors.Wrap(err, "删除文件失败!")
	}
	return nil
}

func (q *qiniu) UploadByFile(file multipart.File) (filepath string, filename string, err error) {
	var result storage.PutRet

	mac := qbox.NewMac(global.Config.Qiniu.AccessKey, global.Config.Qiniu.SecretKey)
	putPolicy := storage.PutPolicy{Scope: global.Config.Qiniu.Bucket}
	uploadToken := putPolicy.UploadToken(mac)

	defer func() {
		if err = file.Close(); err != nil {
			zap.L().Error("文件关闭失败!", zap.Error(err))
		}
	}() // 关闭文件流

	formUploader := storage.NewFormUploader(global.Config.Qiniu.GetConfig())
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}
	err = formUploader.Put(context.Background(), &result, uploadToken, q.filename, file, q.filesize, &putExtra)
	if err != nil {
		return filepath, filename, errors.Wrap(err, "上传文件失败!")
	}
	filename = result.Key
	filepath = global.Config.Qiniu.ImgPath + "/" + filename
	return filepath, filename, nil
}

func (q *qiniu) UploadByFilepath(p string) (path string, filename string, err error) {
	var file *os.File
	file, err = os.Open(p)
	if err != nil {
		return path, filename, errors.Wrapf(err, "(%s)文件不存在!", p)
	}
	var info os.FileInfo
	info, err = file.Stat()
	if err != nil {
		return path, filename, errors.Wrapf(err, "(%s)文件信息获取失败!", p)
	}
	q.filesize = info.Size()
	_, q.filename = filepath.Split(path)
	return q.UploadByFile(file)
}

func (q *qiniu) UploadByFileHeader(file *multipart.FileHeader) (filepath string, filename string, err error) {
	var open multipart.File
	open, err = file.Open()
	if err != nil {
		return filepath, filename, err
	}
	q.filename = fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	q.filesize = file.Size
	return q.UploadByFile(open)
}
