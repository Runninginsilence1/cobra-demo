package service

import (
	"Forensics_Equipment_Plugin_Manager/database"
	"Forensics_Equipment_Plugin_Manager/models"
	"fmt"
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

func CreateTestData() {
	plugins := []*models.Plugin{
		{Name: "摄像头取证", Desc: "摄像头取证插件", Version: "0.1.1", RemainingCount: 5, Uuid: "adsfasd"},
		{Name: "Windows取证", Desc: "Windows取证插件", Version: "0.1.1", RemainingCount: 10, Uuid: "ddddslslsl"},
	}

	database.Db.Create(plugins)
}
