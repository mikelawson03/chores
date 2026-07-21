package api

import (
	"github.com/mikelawson03/chores/internal/app"
)

type apiCfg struct {
	App *app.App
}

func NewApiConfig(app *app.App) apiCfg {
	return apiCfg{
		App: app,
	}
}
