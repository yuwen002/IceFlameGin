package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/constants"
	"ice_flame_gin/internal/system"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 文件上传
type sUpload struct {
	uploadPath string
}

// NewUploadService
//
// @Title NewUploadService
// @Description: 创建一个新的 NewUploadService 服务实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 13:41:05
// @return *sUpload
func NewUploadService() *sUpload {

	return &sUpload{
		uploadPath: constants.UploadPath,
	}
}

// CreateDirectoryByYearMonth
//
// @Title CreateDirectoryByYearMonth
// @Description: 根据当前年份和月份创建目录
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 13:49:57
// @receiver svc
func (svc *sUpload) CreateDirectoryByYearMonth() *system.SysResponse {
	year := time.Now().Year()
	month := time.Now().Month()
	dirName := fmt.Sprintf("./%s/%d/%02d/", svc.uploadPath, year, month)

	// 检查目录是否已经存在
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			return &system.SysResponse{
				Code:    1,
				Message: "创建目录错误",
				Data:    nil,
			}
		}

		return &system.SysResponse{
			Code:    0,
			Message: "目录创建成功",
			Data:    dirName,
		}
	} else {
		return &system.SysResponse{
			Code:    0,
			Message: "目录已存在",
			Data:    dirName,
		}
	}
}

// GenerateUniqueFileName
//
// @Title GenerateUniqueFileName
// @Description: 生成不重复的文件名
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 18:27:57
// @receiver svc
// @return *system.SysResponse
func (svc *sUpload) GenerateUniqueFileName() *system.SysResponse {
	t := time.Now()
	// 使用当前时间戳作为文件名的一部分
	fileName := fmt.Sprintf("%d%02d%02d_%02d%02d%02d_", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	/// 生成随机数以确保唯一性
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := randGenerator.Intn(1000)
	fileName += fmt.Sprintf("%03d", randomNum)

	return &system.SysResponse{
		Code:    0,
		Message: "创建文件成功",
		Data:    fileName,
	}
}

// UploadFile
//
// @Title UploadFile
// @Description: 单文件上传
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 18:40:17
// @receiver svc
// @param c
// @param file
// @return *system.SysResponse
func (svc *sUpload) UploadFile(c *gin.Context, file *multipart.FileHeader) *system.SysResponse {
	//  文件目录
	out := svc.GenerateUniqueFileName()
	if out.Code != 0 {
		return out
	}

	dirName := out.Data.(string)

	// 文件名
	out = svc.GenerateUniqueFileName()
	if out.Code != 0 {
		return out
	}
	fileName := out.Data.(string)

	// 获取文件的扩展名
	ext := filepath.Ext(file.Filename)
	// 拼接完整文件路径
	fullFilePath := filepath.Join(dirName, fileName+ext)

	// 保存文件到指定目录
	err := c.SaveUploadedFile(file, "uploads/"+fullFilePath)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: err.Error(),
		Data:    fullFilePath,
	}
}

// UploadFiles
//
// @Title UploadFiles
// @Description: 多文件上传
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 22:13:58
// @receiver svc
// @param c
// @return *system.SysResponse
func (svc *sUpload) UploadFiles(c *gin.Context, files []*multipart.FileHeader) *system.SysResponse {
	var uploadedFiles []string

	for _, file := range files {
		out := svc.GenerateUniqueFileName()
		if out.Code != 0 {
			return out
		}

		dirName := out.Data.(string)

		out = svc.GenerateUniqueFileName()
		if out.Code != 0 {
			return out
		}
		fileName := out.Data.(string)

		ext := filepath.Ext(file.Filename)
		fullFilePath := filepath.Join(dirName, fileName+ext)

		err := c.SaveUploadedFile(file, "uploads/"+fullFilePath)
		if err != nil {
			return &system.SysResponse{
				Code:    1,
				Message: "文件保存失败",
				Data:    nil,
			}
		}

		uploadedFiles = append(uploadedFiles, fullFilePath)
	}

	return &system.SysResponse{
		Code:    0,
		Message: "文件上传成功",
		Data:    uploadedFiles,
	}
}

// UploadSingleImage
//
// @Title UploadSingleImage
// @Description: 上传图片
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 23:01:29
// @receiver svc
// @param c
// @param file
// @return *system.SysResponse
func (svc *sUpload) UploadSingleImage(c *gin.Context, file *multipart.FileHeader) *system.SysResponse {
	// 检查文件类型是否为图片
	image, err := file.Open()
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "无法打开文件",
			Data:    nil,
		}
	}

	buffer := make([]byte, 512) // 读取文件的前 512 个字节用于判断文件类型
	_, err = image.Read(buffer)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "无法读取文件",
			Data:    nil,
		}
	}

	mimeType := http.DetectContentType(buffer)
	if !strings.HasPrefix(mimeType, "image/") {
		return &system.SysResponse{
			Code:    1,
			Message: "文件不是图片",
			Data:    nil,
		}
	}

	// 继续处理文件上传逻辑
	out := svc.GenerateUniqueFileName()
	if out.Code != 0 {
		return out
	}

	dirName := out.Data.(string)

	out = svc.GenerateUniqueFileName()
	if out.Code != 0 {
		return out
	}
	fileName := out.Data.(string)

	ext := filepath.Ext(file.Filename)
	fullFilePath := filepath.Join(dirName, fileName+ext)

	err = c.SaveUploadedFile(file, "uploads/"+fullFilePath)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "文件保存失败",
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "图片上传成功",
		Data:    fullFilePath,
	}
}

// UploadMultipleImages
//
// @Title UploadMultipleImages
// @Description: 多图片上传
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-31 23:08:18
// @receiver svc
// @param c
// @param files
// @return *system.SysResponse
func (svc *sUpload) UploadMultipleImages(c *gin.Context, files []*multipart.FileHeader) *system.SysResponse {
	var uploadedFiles []string

	for _, file := range files {
		// 打开上传的文件
		uploadedFile, err := file.Open()
		if err != nil {
			return &system.SysResponse{
				Code:    1,
				Message: "无法打开文件",
				Data:    nil,
			}
		}
		defer uploadedFile.Close()

		// 检查文件类型是否为图片
		buffer := make([]byte, 512) // 读取文件的前 512 个字节用于判断文件类型
		_, err = uploadedFile.Read(buffer)
		if err != nil {
			return &system.SysResponse{
				Code:    1,
				Message: "无法读取文件",
				Data:    nil,
			}
		}

		mimeType := http.DetectContentType(buffer)
		if !strings.HasPrefix(mimeType, "image/") {
			return &system.SysResponse{
				Code:    1,
				Message: "文件不是图片",
				Data:    nil,
			}
		}

		// 继续处理文件上传逻辑
		out := svc.GenerateUniqueFileName()
		if out.Code != 0 {
			return out
		}

		dirName := out.Data.(string)

		out = svc.GenerateUniqueFileName()
		if out.Code != 0 {
			return out
		}
		fileName := out.Data.(string)

		ext := filepath.Ext(file.Filename)
		fullFilePath := filepath.Join(dirName, fileName+ext)

		// 重新打开文件，因为前面已经读取了一部分内容
		uploadedFile, _ = file.Open()

		// 保存文件到指定目录
		err = c.SaveUploadedFile(file, "uploads/"+fullFilePath)
		if err != nil {
			return &system.SysResponse{
				Code:    1,
				Message: "文件保存失败",
				Data:    nil,
			}
		}

		uploadedFiles = append(uploadedFiles, fullFilePath)
	}

	return &system.SysResponse{
		Code:    0,
		Message: "多图片上传成功",
		Data:    uploadedFiles,
	}
}
