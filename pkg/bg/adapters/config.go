package adapters

import (
	"github.com/eser/ajan/datafx"
	"github.com/eser/ajan/logfx"
	"github.com/eser/blind-guardian/pkg/bg/business/updater"
)

type AppConfig struct {
	Data   datafx.Config `conf:"DATA"`
	AppEnv string        `conf:"ENV"  default:"development"`

	Update updater.UpdateConfig `conf:"UPDATE"`
	Log    logfx.Config         `conf:"LOG"`
}
