package setup

import (
	"Forensics_Equipment_Plugin_Manager/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

func router() {
	r := gin.Default()
	baseRouter := r.Group("/v0/api")

	pluginRouter := baseRouter.Group("/plugin")

	// kotlin的 run语法糖这里就很舒服了。 稍微封装了一下：
	// 旧： router.GET()...
	// 新： []中添加：{"GET", "/list", controllers.PluginsInfo()}
	{
		pluginRoutes := []*RouteEntity{
			{GET, "/list", controllers.PluginsInfo()},               // 查询插件信息，并且可以进行条件查询
			{POST, "/install", controllers.PluginInstall()},         // 安装新插件
			{POST, "/update", controllers.PluginUpdate()},           // 覆盖现有插件
			{POST, "/uninstall", controllers.PluginUninstall()},     // 移除插件
			{POST, "/testUpdateFile", controllers.TestFileUpload()}, // 移除插件
			{GET, "/modify", controllers.PluginInfoModify()},        // todo 修改插件的配置，待定
		}
		AttachRoute(pluginRouter, pluginRoutes)
	}

	r.Run(":3011")
}

func AttachRoute(group *gin.RouterGroup, entities []*RouteEntity) {
	for _, route := range entities {
		route.Method = strings.ToUpper(route.Method)
		switch route.Method {
		case "GET":
			group.GET(route.RelativePath, route.Handler)
		case "POST":
			group.POST(route.RelativePath, route.Handler)
		case "DELETE":
			group.DELETE(route.RelativePath, route.Handler)
		default:
			fmt.Println("Invalid route method: ", route.Method)
		}

	}
}

type RouteEntity struct {
	Method       string
	RelativePath string
	Handler      gin.HandlerFunc
}
