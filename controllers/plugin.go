package controllers

import (
	"Forensics_Equipment_Plugin_Manager/database"
	"Forensics_Equipment_Plugin_Manager/logger"
	"Forensics_Equipment_Plugin_Manager/models"
	"Forensics_Equipment_Plugin_Manager/models/common"
	"Forensics_Equipment_Plugin_Manager/service"
	"Forensics_Equipment_Plugin_Manager/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

var PluginsRootPath = "./plugins"
var ZipFilePath = "./receive/archive.zip"

// PluginsInfo returns all plugins info to frontend.
func PluginsInfo() gin.HandlerFunc {
	var data any
	var err error

	return func(c *gin.Context) {
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
	return func(c *gin.Context) {
		// 验证参数 ✔
		pluginInfo, err = service.GenValidPluginInfo(c)
		if err != nil {
			common.FailedWithDetail(err.Error(), nil, c)
			return
		}
		logger.Log.Debugf("Print plugin info: \n%v", pluginInfo.FormatString())
		// 执行安装
		// Overwrite Update， 这里应该封装到 service中去
		pluginPath := pluginInfo.Path
		absolutePath := fmt.Sprintf("%v/%v", PluginsRootPath, pluginPath)
		_, err = util.GetFileInfo(absolutePath)
		if err == nil {
			common.FailedWithDetail("该路径已经安装其他插件，请更换路径。", nil, c)
			return
		}
		err = service.PluginUpdate(ZipFilePath, absolutePath)
		if err != nil {
			common.FailedWithDetail(err.Error(), nil, c)
			return
		}
		// 写入到数据库
		database.Db.Save(&pluginInfo)
		common.OKWithDetail("安装成功", pluginInfo, c)
	}
}

// PluginUpdate is similar to PluginInstall.
func PluginUpdate() gin.HandlerFunc {
	var pluginInfo *models.Plugin
	var err error
	return func(c *gin.Context) {
		// 验证参数 ✔
		pluginInfo, err = service.GenValidPluginInfo(c)
		if err != nil {
			common.FailedWithDetail(err.Error(), nil, c)
			return
		}
		logger.Log.Debugf("Print plugin info: \n%v", pluginInfo.FormatString())
		// 执行安装
		// Overwrite Update， 这里应该封装到 service中去
		pluginPath := pluginInfo.Path
		absolutePath := fmt.Sprintf("%v/%v", PluginsRootPath, pluginPath)
		_, err = util.GetFileInfo(absolutePath)
		if err != nil {
			common.FailedWithDetail("不存在该路径，请检查插件更新路径", nil, c)
			return
		}
		err = service.PluginUpdate(ZipFilePath, absolutePath)
		if err != nil {
			common.FailedWithDetail(err.Error(), nil, c)
			return
		}
		// 写入到数据库
		database.Db.Save(&pluginInfo)
		common.OKWithDetail("更新成功", pluginInfo, c)
	}
}

func PluginUninstall() gin.HandlerFunc {
	var err error
	// 验证参数，并且要确认表中包含该主键（或者名称）才可更新插件。
	var pluginInfo models.Plugin

	// 执行卸载

	// 行数据创建
	return func(c *gin.Context) {
		panic("implement me！")
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
		panic("implement me!")
	}
}

func TestFileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pluginInfo models.Plugin
		// test receive file
		file, err := c.FormFile("file")
		if err != nil {
			logger.Log.Error("File Upload, ", err)
			fmt.Println(err)
			common.FailedWithDetail("Bad file argument!", nil, c)
			return
		}
		fmt.Println("file name: ", file.Filename)

		// test receive json data
		jsonValue := c.Request.PostFormValue("data")
		err = json.Unmarshal([]byte(jsonValue), &pluginInfo)
		if err != nil {
			logger.Log.Error("JSON Error: , ", err)
		}
		fmt.Println("Receive Data: ")
		fmt.Println(pluginInfo.FormatString())
		//dst := "./resources/receive/" + file.Filename
		//c.SaveUploadedFile(file, dst)

		common.OK(c)
	}
}
