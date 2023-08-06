package handler

import (
	"github.com/kieferdan/Opportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	dB     *gorm.DB
)

func Initialize() {
	logger = config.GetLogger("handler")
	dB = config.GetSQLite()
}
