package store

import (
	"github.com/qqqasdwx/v2af/common/models"
	"github.com/qqqasdwx/v2af/frontend/config"
)

func InitMysql() {
	cfg := config.Config()
	models.Init(cfg.Database, cfg.ShowSql)
}
