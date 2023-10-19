package controllers

import (
	"Forensics_Equipment_Plugin_Manager/database"
	"Forensics_Equipment_Plugin_Manager/models"
	"Forensics_Equipment_Plugin_Manager/models/common"
	"Forensics_Equipment_Plugin_Manager/service"
	"github.com/gin-gonic/gin"
)

// PluginsInfo returns all plugins info to frontend.
func PluginsInfo() gin.HandlerFunc {
	var data any
	var err error

	return func(c *gin.Context) {
		data = gin.H{"code": 333}
		data, err = service.FetchPluginInfo()
		if err != nil {
			common.Failed(c)
		} else {
			common.OKWithDetail("获取插件信息成功", data, c)
		}
	}
}

// PluginInstall provides an interface for installing new plugins.
func PluginInstall() gin.HandlerFunc {
	var err error
	// 验证参数
	var pluginInfo models.Plugin

	// 执行安装

	// 行数据创建
	return func(c *gin.Context) {
		err = c.BindJSON(&pluginInfo)
		if err != nil {
			common.FailedWithWrongReq(c)
			return
		}
		database.Db.Save(&pluginInfo)
		common.OK(c)
	}
}

func PluginUpdate() gin.HandlerFunc {
	var err error
	// 验证参数，并且要确认表中包含该主键（或者名称）才可更新插件。
	var pluginInfo models.Plugin

	// 执行更新

	// 行数据创建
	return func(c *gin.Context) {
		err = c.BindJSON(&pluginInfo)
		if err != nil {
			common.FailedWithWrongReq(c)
			return
		}
		database.Db.Save(&pluginInfo)
		common.OK(c)
	}
}

func PluginUninstall() gin.HandlerFunc {
	var err error
	// 验证参数，并且要确认表中包含该主键（或者名称）才可更新插件。
	var pluginInfo models.Plugin

	// 执行卸载

	// 行数据创建
	return func(c *gin.Context) {
		err = c.BindJSON(&pluginInfo)
		if err != nil {
			common.FailedWithWrongReq(c)
			return
		}
		database.Db.Save(&pluginInfo)
		common.OK(c)
	}
}
