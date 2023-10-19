package controllers

import (
	"Forensics_Equipment_Plugin_Manager/database"
	"Forensics_Equipment_Plugin_Manager/models"
	"Forensics_Equipment_Plugin_Manager/models/common"
	"Forensics_Equipment_Plugin_Manager/service"
	"Forensics_Equipment_Plugin_Manager/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

var PluginsRootPath = "./aaa/plugins"

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

// PluginInstall provides an interface for installing new plugin.
func PluginInstall() gin.HandlerFunc {
	var pluginInfo *models.Plugin
	var err error
	// 验证参数

	// 执行安装

	// 行数据创建
	return func(c *gin.Context) {
		pluginInfo, err = service.GenPluginInfo(c)
		if err != nil {
			common.FailedWithWrongReq(c)
			return
		}
		fmt.Println("Print plugin info: ")
		fmt.Println(pluginInfo.FormatString())

		// Overwrite Update

		pluginPath := pluginInfo.Path
		if !util.IsValidFileName(pluginPath) {
			common.FailedWithDetail("请确保插件路径中不要包含空格、中文或者其他特殊字符。", nil, c)
			return
		}
		absolutePath := fmt.Sprintf("%v/%v", PluginsRootPath, pluginPath)
		_, err = util.GetFileInfo(absolutePath)
		if err == nil {
			common.OKWithDetail("该路径已经安装其他插件，请更换路径。", nil, c)
			return
		}

		util.MkdirAll(absolutePath)
		// util.UnzipWithTimeout()
		//
		//database.Db.Save(&pluginInfo)
		common.OKWithDetail("是一个有效的插件路径", nil, c)
	}
}

// PluginUpdate is similar to PluginInstall.
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

func PluginInfoModify() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func TestFileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			common.FailedWithDetail("Bad file argument!", nil, c)
			return
		}
		fmt.Println("file name: ", file.Filename)

		dst := "./resources/receive/" + file.Filename
		c.SaveUploadedFile(file, dst)

		common.OK(c)
	}
}
