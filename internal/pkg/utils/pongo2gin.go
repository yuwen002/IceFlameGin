package utils

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"ice_flame_gin/internal/system"
	"net/http"
	"path"
)

// PongoRender
//
// @Description: 结构体初始化
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-11 01:17:52
type PongoRender struct {
	TmplDir string
}

// TemplatePath
//
// @Title TemplatePath
// @Description: 返回 PongoRender 实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-11 01:18:24
// @param tmplDir
// @return *PongoRender
func TemplatePath(tmplDir string) *PongoRender {
	return &PongoRender{
		TmplDir: tmplDir,
	}
}

// Instance
//
// @Title Instance
// @Description: 初始化模板并返回渲染器对象
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-11 01:18:37
// @receiver p
// @param name
// @param data
// @return render.Render
func (p *PongoRender) Instance(name string, data interface{}) render.Render {
	var template *pongo2.Template
	fileName := path.Join(p.TmplDir, name)

	if gin.Mode() == gin.DebugMode {
		template = pongo2.Must(pongo2.FromFile(fileName))
	} else {
		template = pongo2.Must(pongo2.FromCache(fileName))
	}

	var ctx pongo2.Context
	if data != nil {
		ctx = data.(pongo2.Context)
	}

	return &PongoHTML{
		Template: template,
		Name:     name,
		Data:     ctx,
	}
}

// PongoHTML
//
// @Description: 结构体
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-11 01:18:49
type PongoHTML struct {
	Template *pongo2.Template
	Name     string
	Data     pongo2.Context
}

// Render
//
// @Title Render
// @Description: gin.Render 接口的实现，用于渲染模板
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-11 01:19:52
// @receiver p
// @param w
// @return error
func (p *PongoHTML) Render(w http.ResponseWriter) error {
	p.WriteContentType(w)
	return p.Template.ExecuteWriter(p.Data, w)
}

// WriteContentType
//
// @Title WriteContentType
// @Description: gin.Render 接口的实现，用于设置响应的内容类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-11 01:20:12
// @receiver p
// @param w
func (p *PongoHTML) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"text/html; charset=utf-8"}
	}
}

// OldFilter
//
// @Title OldFilter
// @Description: 模板Old，类似laravel函数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-12 23:07:16
// @param in
// @param param
// @return out
// @return err
func OldFilter(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	// 获取键名
	key := param.String()

	// 获取 Gin 上下文
	ginContext, ok := in.Interface().(*gin.Context)
	if !ok {
		return nil, &pongo2.Error{
			Sender:    "old",
			OrigError: fmt.Errorf("invalid gin.Context type"),
		}
	}

	// 调用获取旧输入数据的函数
	oldValue := system.GetOldInput(ginContext, key)

	return pongo2.AsValue(oldValue), nil
}
