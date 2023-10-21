package service

import (
	"Forensics_Equipment_Plugin_Manager/database"
	"Forensics_Equipment_Plugin_Manager/models"
	"Forensics_Equipment_Plugin_Manager/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func FetchPluginInfo() (result []*models.Plugin, err error) {
	scan := database.Db.Raw("select id, name, desc, remaining_count, version, uuid from plugins").Scan(&result)

	if scan.RowsAffected == 0 {
		err = fmt.Errorf("no data in table")
		CreateTestData()
		return nil, err
	}

	return result, nil
}

func PluginInstall() error {
	return nil
}

// zip file path; des path;
func PluginUpdate(zipPath, pluginPath string) error {
	var err error
	if !util.IsZipFileValid(zipPath) {
		return fmt.Errorf("wrong zip file")
	}
	util.MkdirAll(pluginPath)
	err = util.UnzipWithTimeout(zipPath, pluginPath, 10*time.Second, true)
	if err != nil {
		return err
	}
	return nil
}

func CreateTestData() {
	plugins := []*models.Plugin{
		{Name: "摄像头取证", Desc: "摄像头取证插件", Version: "0.1.1", RemainingCount: 5, Uuid: "adsfasd"},
		{Name: "Windows取证", Desc: "Windows取证插件", Version: "0.1.1", RemainingCount: 10, Uuid: "ddddslslsl"},
	}

	database.Db.Create(plugins)
}

// IsValidPluginInfo determines whether the info is valid.
func IsValidPluginInfo(plugin *models.Plugin) bool {
	// 验证每个字段
	pluginPath := plugin.Path
	if !util.IsValidFileName(pluginPath) {
		return false
	}
	return true
}

func GenValidPluginInfo(c *gin.Context) (*models.Plugin, error) {
	// json unmarshal
	var plugin models.Plugin
	var err error
	err = c.BindJSON(&plugin)
	if err != nil {
		return nil, fmt.Errorf("解析JSON失败")
	}

	// check valid info
	pluginPath := plugin.Path
	if !util.IsValidFileName(pluginPath) {
		return nil, fmt.Errorf("请确保插件路径中不要包含空格、中文或者其他特殊字符。")
	}
	return &plugin, nil
}
