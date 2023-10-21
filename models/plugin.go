package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Plugin struct {
	gorm.Model     `json:"-"`
	Name           string    `json:"name"`
	Uuid           string    `json:"uuid"`
	Desc           string    `json:"desc"`
	Version        string    `json:"version"`
	RemainingCount int       `json:"remaining_count"` // 剩余使用次数
	ReleaseTime    time.Time `json:"release_time"`    // 插件发行时间
	Icon           string    `json:"icon"`
	Port           string    `json:"port"`
	Path           string    `json:"path"` // 插件路径
	// todo 插件级别和分级的需求，参考若依
}

func (l *Plugin) FormatString() string {
	return fmt.Sprintf(`
Plugin UUID: %v
Plugin name: %v
plugin desciption: %v
Plugin version: %v
Plugin path: %v
`, l.Uuid, l.Name, l.Desc, l.Version, l.Path)
}

type Plugins struct {
	ID             int       `json:"id" gorm:"primarykey;"`                             //插件ID
	Name           string    `json:"name" gorm:"column:name"`                           //插件名
	Icon           string    `json:"icon" gorm:"column:icon"`                           //图标
	Version        string    `json:"mod_version" gorm:"column:mod_version"`             //版本
	Path           string    `json:"path" gorm:"column:path"`                           //插件路径
	Enable         int       `json:"enable" gorm:"column:enable;default:1"`             //是否启动
	IPAddress      string    `json:"ip_address" gorm:"column:ipaddress;"`               //IP地址
	IsAuthority    int       `json:"is_authority" gorm:"column:is_authority;default:0"` //权限  使用版 过期  正式版
	ExpireTime     time.Time `json:"expire_time" gorm:"column:expire_time;"`            //过期时间
	RemainingTimes uint      `json:"remaining_times" gorm:"column:remaining_times;"`    //剩余次数
	GroupID        int       `json:"group_id" gorm:"column:group_id"`                   //类别ID
}

func (*Plugins) TableName() string {
	return "plugins"
}

func (l *Plugins) FormatString() string {
	return fmt.Sprintf(`
Plugin name: %v
Plugin version: %v
Plugin path: %v
`, l.Name, l.Version, l.Path)
}

type Group struct {
	ID       int    `json:"id" gorm:"primarykey;"`                       //插件ID
	ParentId int    `json:"parent_id" gorm:"column:parent_id;default:0"` //父类ID
	Name     string `json:"name" gorm:"column:name"`                     //组名
}

func (Group) TableName() string {
	return "Group"
}
