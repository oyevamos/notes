package controllers

import "github.com/oyevamos/notes.git/config"

var appConfig *config.AppConfig

func SetAppConfig(cfg *config.AppConfig) {
	appConfig = cfg
}
