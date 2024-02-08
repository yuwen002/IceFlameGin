package system

//	SysResponse
//	@Description: 系统标准输出
//
// @Author liuxingyu
// @Date 2024-02-08 23:16:25
type SysResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
