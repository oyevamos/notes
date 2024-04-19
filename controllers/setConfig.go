package controllers

import (
	"github.com/oyevamos/notes.git/storage"
)

var appConfig *storage.AppConfig

func SetAppConfig(cfg *storage.AppConfig) {
	appConfig = cfg
}
