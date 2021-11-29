package http

import (
	"aurora/core"
	"github.com/gin-gonic/gin"
)

// Upload 上传图片保存到本地并返回路径
func Upload(ctx *gin.Context) {
	//从上下文获取上传的file文件
	files, header, err := ctx.Request.FormFile("file")
	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}
	//关闭文件，释放资源
	defer files.Close()
	//上传文件
	err, url := srv.UploadFile(files, *header)

	if err != nil {
		core.ErrorMsg(ctx, err.Error())
		return
	}

	core.SuccessMsg(ctx, url)

}


func GetQiNiuToken(ctx *gin.Context) {
	token := srv.GetQiNiuToken()
	core.SuccessMsg(ctx, token)
}
