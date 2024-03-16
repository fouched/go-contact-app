package apix

import (
	"database/sql"
	"github.com/fouched/go-contact-app/internal/config"
)

var Instance *HtmxApiConfig

type HtmxApiConfig struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewConfig(a *config.AppConfig, db *sql.DB) *HtmxApiConfig {
	return &HtmxApiConfig{
		App: a,
		DB:  db,
	}
}

func NewHandlers(h *HtmxApiConfig) {
	Instance = h
}
