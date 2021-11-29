package service

import (
	"aurora/config"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
	"io"
	"mime/multipart"
	"os"
)

func (s *Service) UploadFile(files multipart.File, header multipart.FileHeader) (error, string) {
	//创建文件
	f, err := os.Create("./static/file/" + header.Filename)
	if err != nil {
		return err, ""
	}
	//关闭文件，释放资源。
	defer f.Close()
	//文件指针之间直接复制的，不用全读入内存
	_, err = io.Copy(f, files)
	if err != nil {
		return err, ""
	}
	//文件储存路径
	url := "/static/file/" + header.Filename
	return nil, url
}

func (s *Service) GetQiNiuToken() string {
	putPolicy := storage.PutPolicy{
		Scope: config.QiNiu().Bucket,
	}
	mac := auth.New(config.QiNiu().AccessKey, config.QiNiu().SecretKey)
	putPolicy.Expires = 7200 //有效期 两个小时
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
