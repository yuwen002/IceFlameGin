package utils

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
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

	return &PongoHTML{
		Template: template,
		Name:     name,
		Data:     data.(pongo2.Context),
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

type GinContext struct {
	GinContext *gin.Context
}

func Old(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {

	return nil, nil

}
