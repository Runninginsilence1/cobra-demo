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
