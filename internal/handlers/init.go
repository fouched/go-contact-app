package handlers

import (
	"github.com/fouched/go-contact-app/internal/config"
)

var Instance *HandlerConfig

type HandlerConfig struct {
	App *config.AppConfig
}

func NewConfig(a *config.AppConfig) *HandlerConfig {
	return &HandlerConfig{
		App: a,
	}
}

func NewHandlers(h *HandlerConfig) {
	Instance = h
}

func MakeUpsertMap(t string, a string) map[string]string {
	m := make(map[string]string)
	m["title"] = t
	m["action"] = a

	return m
}
