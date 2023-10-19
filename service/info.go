package service

import (
	"Forensics_Equipment_Plugin_Manager/database"
	"Forensics_Equipment_Plugin_Manager/models"
	"fmt"
	"github.com/gin-gonic/gin"
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

func PluginInstall() {}

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
	return true
}

func GenPluginInfo(c *gin.Context) (*models.Plugin, error) {
	var plugin models.Plugin
	var err error
	err = c.BindJSON(&plugin)
	if err != nil {
		return nil, err
	}
	if !IsValidPluginInfo(&plugin) {
		return nil, fmt.Errorf("invalid plugin info")
	}
	return &plugin, nil
}
