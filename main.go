package main

import (
	"aurora/http"
	"github.com/gin-gonic/gin"
)

func main() {
	//控制台日志颜色
	gin.ForceConsoleColor()

	http.Init()
}
