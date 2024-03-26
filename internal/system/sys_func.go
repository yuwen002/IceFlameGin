package system

import (
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/services"
	"ice_flame_gin/internal/pkg/utils"
)

// WriteSystemMasterVisitorLogs
//
// @Title SystemMasterVisitorLogs
// @Description: 写入用户访问几率
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 22:46:17
// @param accountID
// @param osCategory
// @param visitCategory
// @param unionID
// @param description
// @param
func WriteSystemMasterVisitorLogs(c *gin.Context, osCategory uint32, visitCategory uint32, unionID uint32, description string) error {
	// 管理员AccountID
	accountID, _, err := GetMasterInfo(c)
	if err != nil {
		return err
	}

	// 获取管理员IP
	ip := c.ClientIP()
	ipLong, err := utils.IPToLong(ip)
	if err != nil {
		return err
	}

	ipLongString, err := utils.ToString(ipLong)
	if err != nil {
		return err
	}

	services.NewUcSystemMasterVisit().CreateVisitorLogs(dto.SystemMasterVisitorLogsInput{
		AccountID:     accountID,
		OsCategory:    osCategory,
		VisitCategory: visitCategory,
		UnionID:       unionID,
		Description:   description,
		IPLong:        ipLongString,
		IP:            ip,
	})

	return nil
}
