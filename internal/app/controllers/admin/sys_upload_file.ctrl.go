package admin

import (
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/services"
	"net/http"
)

// cSysUploadFile
//
// @Description: 文件上传
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 23:17:16
type cSysUploadFile struct {
}

// SysUploadFile 初始化文件上传控制器
var SysUploadFile = &cSysUploadFile{}

// UploadFile
//
// @Title UploadFile
// @Description: 单文件上传
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 23:25:53
// @receiver ctrl
// @param c
func (ctrl *cSysUploadFile) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将文件保存到服务器上
	out := services.NewUploadService().UploadFile(c, file)
	if out.Code != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": out.Message})
		return
	}

	// 返回文件路径给Summernote
	c.JSON(http.StatusOK, gin.H{"url": out.Data.(string)})
}

// UploadImage
//
// @Title UploadImage
// @Description: 单图片上传
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 23:46:47
// @receiver ctrl
// @param c
func (ctrl *cSysUploadFile) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将文件保存到服务器上
	out := services.NewUploadService().UploadSingleImage(c, file)
	if out.Code != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": out.Message})
		return
	}

	// 返回文件路径给Summernote
	c.JSON(http.StatusOK, gin.H{"url": out.Data.(string)})
}
