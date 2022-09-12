package model

import "github.com/lezi-wiki/lezi-api/pkg/conf"

var defaultSettings = []Setting{
	{Name: "version", Type: SettingTypeSystem, Val: conf.Version},
}
