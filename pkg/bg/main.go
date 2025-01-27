package bg

import (
	"github.com/eser/ajan/configfx"
	"github.com/eser/blind-guardian/pkg/bg/adapters"
)

func Init() (*adapters.AppConfig, error) {
	cl := configfx.NewConfigManager()

	appConfig := &adapters.AppConfig{} //nolint:exhaustruct

	err := cl.LoadDefaults(appConfig)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return appConfig, nil
}
