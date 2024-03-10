package handlers

import (
	"database/sql"
	"github.com/fouched/go-contact-app/internal/config"
)

var Instance *HandlerConfig

type HandlerConfig struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewConfig(a *config.AppConfig, db *sql.DB) *HandlerConfig {
	return &HandlerConfig{
		App: a,
		DB:  db,
	}
}

func NewHandlers(h *HandlerConfig) {
	Instance = h
}
