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

// OldTag
//
// @Title OldTag
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-11 23:25:02
// @param doc
// @param start
// @param line
// @return pongo2.Tag
// @return *pongo2.Error
func OldTag(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	// 解析标签参数
	value, err := arguments.ParseExpression()
	if err != nil {
		return nil, &pongo2.Error{
			Sender:    "old",
			OrigError: fmt.Errorf("无法解析标签参数"),
		}
	}

	// 解析标签参数
	valueToken := arguments.Peek(pongo2.String, "")
	if valueToken == nil {
		return nil, &pongo2.Error{
			Sender:    "old",
			OrigError: fmt.Errorf("无法获取标签参数"),
		}
	}
	value := valueToken.Val

	// 进行其他操作...

	return nil, nil

}

type OldTagImpl struct {
	fieldName string
}

func (node *OldTagImpl) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	// 获取 Gin 上下文
	ginContext, ok := ctx.Public["ginContext"].(*gin.Context)
	if !ok {
		return &pongo2.Error{
			Sender:    "old",
			OrigError: fmt.Errorf("invalid gin.Context type"),
		}
	}

	// 调用获取旧输入数据的函数
	oldValue := system.GetOldInput(ginContext, node.fieldName)

	// 将旧输入数据写入模板
	writer.WriteString(oldValue)

	return nil
}
