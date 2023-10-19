package setup

import "Forensics_Equipment_Plugin_Manager/database"

func Run() {
	database.InitDb()
	router()
}
